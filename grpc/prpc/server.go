// Copyright 2016 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prpc

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"sync"

	"github.com/golang/protobuf/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/TriggerMail/luci-go/common/retry/transient"
	"github.com/TriggerMail/luci-go/grpc/grpcutil"
	"github.com/TriggerMail/luci-go/server/router"
)

var (
	// Describe the permitted Access Control requests.
	allowHeaders = strings.Join([]string{"Origin", "Content-Type", "Accept", "Authorization"}, ", ")
	allowMethods = strings.Join([]string{"OPTIONS", "POST"}, ", ")

	// allowPreflightCacheAgeSecs is the amount of time to enable the browser to
	// cache the preflight access control response, in seconds.
	//
	// 600 seconds is 10 minutes.
	allowPreflightCacheAgeSecs = "600"

	// exposeHeaders lists the whitelisted non-standard response headers that the
	// client may accept.
	exposeHeaders = strings.Join([]string{HeaderGRPCCode}, ", ")

	// NoAuthentication can be used in place of an Authenticator to explicitly
	// specify that your Server will skip authentication.
	//
	// Use it with Server.Authenticator or RegisterDefaultAuth.
	NoAuthentication Authenticator = nullAuthenticator{}
)

// Server is a pRPC server to serve RPC requests.
// Zero value is valid.
type Server struct {
	// Authenticator, if not nil, specifies how to authenticate requests.
	//
	// If nil, the default authenticator set by RegisterDefaultAuth will be used.
	// If the default authenticator is also nil, all request handlers will panic.
	//
	// If you want to disable the authentication (e.g for tests), explicitly set
	// Authenticator to NoAuthentication.
	Authenticator Authenticator

	// AccessControl, if not nil, is a callback that is invoked per request to
	// determine if permissive access control headers should be added to the
	// response.
	//
	// This callback includes the request Context and the origin header supplied
	// by the client. If nil, or if it returns false, no headers will be written.
	// Otherwise, access control headers for the specified origin will be
	// included in the response.
	AccessControl func(c context.Context, origin string) bool

	// UnaryServerInterceptor provides a hook to intercept the execution of
	// a unary RPC on the server. It is the responsibility of the interceptor to
	// invoke handler to complete the RPC.
	UnaryServerInterceptor grpc.UnaryServerInterceptor

	mu       sync.Mutex
	services map[string]*service
}

type service struct {
	methods map[string]grpc.MethodDesc
	impl    interface{}
}

// RegisterService registers a service implementation.
// Called from the generated code.
//
// desc must contain description of the service, its message types
// and all transitive dependencies.
//
// Panics if a service of the same name is already registered.
func (s *Server) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	serv := &service{
		impl:    impl,
		methods: make(map[string]grpc.MethodDesc, len(desc.Methods)),
	}
	for _, m := range desc.Methods {
		serv.methods[m.MethodName] = m
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.services == nil {
		s.services = map[string]*service{}
	} else if _, ok := s.services[desc.ServiceName]; ok {
		panic(fmt.Errorf("service %q is already registered", desc.ServiceName))
	}

	s.services[desc.ServiceName] = serv
}

// authenticate forces authentication set by RegisterDefaultAuth.
func (s *Server) authenticate() router.Middleware {
	a := s.Authenticator
	if a == nil {
		a = GetDefaultAuth()
		if a == nil {
			panic("prpc: no custom Authenticator was provided and default authenticator was not registered.\n" +
				"Either explicitly set `Server.Authenticator = NoAuthentication`, or use RegisterDefaultAuth()")
		}
	}

	return func(c *router.Context, next router.Handler) {
		switch ctx, err := a.Authenticate(c.Context, c.Request); {
		case transient.Tag.In(err):
			writeError(c.Context, c.Writer, withCode(err, codes.Internal))
		case err != nil:
			writeError(c.Context, c.Writer, withCode(err, codes.Unauthenticated))
		default:
			c.Context = ctx
			next(c)
		}
	}
}

// InstallHandlers installs HTTP handlers at /prpc/:service/:method.
//
// See https://godoc.org/github.com/TriggerMail/luci-go/grpc/prpc#hdr-Protocol
// for pRPC protocol.
//
// The authenticator in 'base' is always replaced by pRPC specific one. For more
// details about the authentication see Server.Authenticator doc.
func (s *Server) InstallHandlers(r *router.Router, base router.MiddlewareChain) {
	s.mu.Lock()
	defer s.mu.Unlock()

	rr := r.Subrouter("/prpc/:service/:method")
	rr.Use(base.Extend(s.authenticate()))

	rr.POST("", router.MiddlewareChain{}, s.handlePOST)
	rr.OPTIONS("", router.MiddlewareChain{}, s.handleOPTIONS)
}

// handle handles RPCs.
// See https://godoc.org/github.com/TriggerMail/luci-go/grpc/prpc#hdr-Protocol
// for pRPC protocol.
func (s *Server) handlePOST(c *router.Context) {
	serviceName := c.Params.ByName("service")
	methodName := c.Params.ByName("method")
	s.setAccessControlHeaders(c, false)
	res := s.call(c, serviceName, methodName)
	if res.err != nil {
		writeError(c.Context, c.Writer, res.err)
		return
	}
	writeMessage(c.Context, c.Writer, res.out, res.fmt)
}

func (s *Server) handleOPTIONS(c *router.Context) {
	s.setAccessControlHeaders(c, true)
	c.Writer.WriteHeader(http.StatusOK)
}

type response struct {
	out proto.Message
	fmt Format
	err error
}

func (s *Server) call(c *router.Context, serviceName, methodName string) (r response) {
	service := s.services[serviceName]
	if service == nil {
		r.err = status.Errorf(
			codes.Unimplemented,
			"service %q is not implemented",
			serviceName)
		return
	}

	method, ok := service.methods[methodName]
	if !ok {
		r.err = status.Errorf(
			codes.Unimplemented,
			"method %q in service %q is not implemented",
			methodName, serviceName)
		return
	}

	var perr *protocolError
	r.fmt, perr = responseFormat(c.Request.Header.Get(headerAccept))
	if perr != nil {
		r.err = perr
		return
	}

	methodCtx, err := parseHeader(c.Context, c.Request.Header)
	if err != nil {
		r.err = withStatus(err, http.StatusBadRequest)
		return
	}

	out, err := method.Handler(service.impl, methodCtx, func(in interface{}) error {
		if in == nil {
			return grpcutil.Errf(codes.Internal, "input message is nil")
		}
		// Do not collapse it to one line. There is implicit err type conversion.
		if perr := readMessage(c.Request, in.(proto.Message)); perr != nil {
			return perr
		}
		return nil
	}, s.UnaryServerInterceptor)

	switch {
	case err != nil:
		r.err = err
	case out == nil:
		r.err = status.Error(codes.Internal, "service returned nil message")
	default:
		r.out = out.(proto.Message)
	}
	return
}

func (s *Server) setAccessControlHeaders(c *router.Context, preflight bool) {
	// Don't write out access control headers if the origin is unspecified.
	const originHeader = "Origin"
	origin := c.Request.Header.Get(originHeader)
	if origin == "" || s.AccessControl == nil || !s.AccessControl(c.Context, origin) {
		return
	}

	h := c.Writer.Header()
	h.Add("Access-Control-Allow-Origin", origin)
	h.Add("Vary", originHeader)
	h.Add("Access-Control-Allow-Credentials", "true")

	if preflight {
		h.Add("Access-Control-Allow-Headers", allowHeaders)
		h.Add("Access-Control-Allow-Methods", allowMethods)
		h.Add("Access-Control-Max-Age", allowPreflightCacheAgeSecs)
	} else {
		h.Add("Access-Control-Expose-Headers", exposeHeaders)
	}
}

// ServiceNames returns a sorted list of full names of all registered services.
func (s *Server) ServiceNames() []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	names := make([]string, 0, len(s.services))
	for name := range s.services {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
