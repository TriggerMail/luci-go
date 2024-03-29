// Copyright 2015 The LUCI Authors.
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

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/TriggerMail/luci-go/examples/appengine/helloworld_standard/proto"
	"github.com/TriggerMail/luci-go/grpc/prpc"
)

func main() {
	ctx := context.Background()
	greeter := helloworld.NewGreeterPRPCClient(&prpc.Client{Host: "https://helloworld-dot-prpc-talk.appspot.com"})

	req := &helloworld.HelloRequest{}
	if len(os.Args) > 1 {
		req.Name = os.Args[1]
	}

	res, err := greeter.SayHello(ctx, req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(res.Message)
}
