// Copyright 2018 The LUCI Authors.
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

// Package buildbucket provides access to the Build Bucket Service.
//
// Usage example:
//
//   import "github.com/TriggerMail/luci-go/common/api/buildbucket/buildbucket/v1"
//   ...
//   buildbucketService, err := buildbucket.New(oauthHttpClient)
package buildbucket // import "github.com/TriggerMail/luci-go/common/api/buildbucket/buildbucket/v1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "buildbucket:v1"
const apiName = "buildbucket"
const apiVersion = "v1"
const basePath = "http://localhost:8080/_ah/api/buildbucket/v1"

// OAuth2 scopes used by this API.
const (
	// https://www.googleapis.com/auth/userinfo.email
	UserinfoEmailScope = "https://www.googleapis.com/auth/userinfo.email"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

type ApiBucketMessage struct {
	ConfigFileContent string `json:"config_file_content,omitempty"`

	ConfigFileRev string `json:"config_file_rev,omitempty"`

	ConfigFileUrl string `json:"config_file_url,omitempty"`

	Error *ApiErrorMessage `json:"error,omitempty"`

	Name string `json:"name,omitempty"`

	ProjectId string `json:"project_id,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ConfigFileContent")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ConfigFileContent") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ApiBucketMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ApiBucketMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiBuildResponseMessage struct {
	// Build: Describes model.Build, see its docstring.
	Build *ApiCommonBuildMessage `json:"build,omitempty"`

	Error *ApiErrorMessage `json:"error,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Build") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Build") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApiBuildResponseMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ApiBuildResponseMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiCancelBatchRequestMessage struct {
	BuildIds googleapi.Int64s `json:"build_ids,omitempty"`

	ResultDetailsJson string `json:"result_details_json,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BuildIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BuildIds") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApiCancelBatchRequestMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ApiCancelBatchRequestMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiCancelBatchResponseMessage struct {
	Error *ApiErrorMessage `json:"error,omitempty"`

	Results []*ApiCancelBatchResponseMessageOneResult `json:"results,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Error") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Error") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApiCancelBatchResponseMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ApiCancelBatchResponseMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiCancelBatchResponseMessageOneResult struct {
	// Build: Describes model.Build, see its docstring.
	Build *ApiCommonBuildMessage `json:"build,omitempty"`

	BuildId int64 `json:"build_id,omitempty,string"`

	Error *ApiErrorMessage `json:"error,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Build") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Build") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApiCancelBatchResponseMessageOneResult) MarshalJSON() ([]byte, error) {
	type NoMethod ApiCancelBatchResponseMessageOneResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiCancelRequestBodyMessage struct {
	ResultDetailsJson string `json:"result_details_json,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ResultDetailsJson")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ResultDetailsJson") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ApiCancelRequestBodyMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ApiCancelRequestBodyMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ApiCommonBuildMessage: Describes model.Build, see its docstring.
type ApiCommonBuildMessage struct {
	Bucket string `json:"bucket,omitempty"`

	Canary bool `json:"canary,omitempty"`

	// Possible values:
	//   "AUTO"
	//   "CANARY"
	//   "PROD"
	CanaryPreference string `json:"canary_preference,omitempty"`

	// Possible values:
	//   "CANCELED_EXPLICITLY"
	//   "TIMEOUT"
	CancelationReason string `json:"cancelation_reason,omitempty"`

	CompletedTs int64 `json:"completed_ts,omitempty,string"`

	CreatedBy string `json:"created_by,omitempty"`

	CreatedTs int64 `json:"created_ts,omitempty,string"`

	Experimental bool `json:"experimental,omitempty"`

	// Possible values:
	//   "BUILDBUCKET_FAILURE"
	//   "BUILD_FAILURE"
	//   "INFRA_FAILURE"
	//   "INVALID_BUILD_DEFINITION"
	FailureReason string `json:"failure_reason,omitempty"`

	Id int64 `json:"id,omitempty,string"`

	LeaseExpirationTs int64 `json:"lease_expiration_ts,omitempty,string"`

	LeaseKey int64 `json:"lease_key,omitempty,string"`

	ParametersJson string `json:"parameters_json,omitempty"`

	Project string `json:"project,omitempty"`

	// Possible values:
	//   "CANCELED"
	//   "FAILURE"
	//   "SUCCESS"
	Result string `json:"result,omitempty"`

	ResultDetailsJson string `json:"result_details_json,omitempty"`

	RetryOf int64 `json:"retry_of,omitempty,string"`

	ServiceAccount string `json:"service_account,omitempty"`

	StartedTs int64 `json:"started_ts,omitempty,string"`

	// Possible values:
	//   "COMPLETED"
	//   "SCHEDULED"
	//   "STARTED"
	Status string `json:"status,omitempty"`

	StatusChangedTs int64 `json:"status_changed_ts,omitempty,string"`

	Tags []string `json:"tags,omitempty"`

	UpdatedTs int64 `json:"updated_ts,omitempty,string"`

	Url string `json:"url,omitempty"`

	UtcnowTs int64 `json:"utcnow_ts,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "Bucket") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Bucket") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApiCommonBuildMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ApiCommonBuildMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiDeleteManyBuildsResponse struct {
	Error *ApiErrorMessage `json:"error,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Error") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Error") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApiDeleteManyBuildsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ApiDeleteManyBuildsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiErrorMessage struct {
	Message string `json:"message,omitempty"`

	// Possible values:
	//   "BUILDER_NOT_FOUND"
	//   "BUILD_IS_COMPLETED"
	//   "BUILD_NOT_FOUND"
	//   "CANNOT_LEASE_BUILD"
	//   "INVALID_BUILD_STATE"
	//   "INVALID_INPUT"
	//   "LEASE_EXPIRED"
	Reason string `json:"reason,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Message") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Message") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApiErrorMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ApiErrorMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiFailRequestBodyMessage struct {
	// Possible values:
	//   "BUILDBUCKET_FAILURE"
	//   "BUILD_FAILURE"
	//   "INFRA_FAILURE"
	//   "INVALID_BUILD_DEFINITION"
	FailureReason string `json:"failure_reason,omitempty"`

	LeaseKey int64 `json:"lease_key,omitempty,string"`

	NewTags []string `json:"new_tags,omitempty"`

	ResultDetailsJson string `json:"result_details_json,omitempty"`

	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FailureReason") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FailureReason") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApiFailRequestBodyMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ApiFailRequestBodyMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiHeartbeatBatchRequestMessage struct {
	Heartbeats []*ApiHeartbeatBatchRequestMessageOneHeartbeat `json:"heartbeats,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Heartbeats") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Heartbeats") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApiHeartbeatBatchRequestMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ApiHeartbeatBatchRequestMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiHeartbeatBatchRequestMessageOneHeartbeat struct {
	BuildId int64 `json:"build_id,omitempty,string"`

	LeaseExpirationTs int64 `json:"lease_expiration_ts,omitempty,string"`

	LeaseKey int64 `json:"lease_key,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "BuildId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BuildId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApiHeartbeatBatchRequestMessageOneHeartbeat) MarshalJSON() ([]byte, error) {
	type NoMethod ApiHeartbeatBatchRequestMessageOneHeartbeat
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiHeartbeatBatchResponseMessage struct {
	Error *ApiErrorMessage `json:"error,omitempty"`

	Results []*ApiHeartbeatBatchResponseMessageOneHeartbeatResult `json:"results,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Error") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Error") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApiHeartbeatBatchResponseMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ApiHeartbeatBatchResponseMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiHeartbeatBatchResponseMessageOneHeartbeatResult struct {
	BuildId int64 `json:"build_id,omitempty,string"`

	Error *ApiErrorMessage `json:"error,omitempty"`

	LeaseExpirationTs int64 `json:"lease_expiration_ts,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "BuildId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BuildId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApiHeartbeatBatchResponseMessageOneHeartbeatResult) MarshalJSON() ([]byte, error) {
	type NoMethod ApiHeartbeatBatchResponseMessageOneHeartbeatResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiHeartbeatRequestBodyMessage struct {
	LeaseExpirationTs int64 `json:"lease_expiration_ts,omitempty,string"`

	LeaseKey int64 `json:"lease_key,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "LeaseExpirationTs")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LeaseExpirationTs") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ApiHeartbeatRequestBodyMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ApiHeartbeatRequestBodyMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiLeaseRequestBodyMessage struct {
	LeaseExpirationTs int64 `json:"lease_expiration_ts,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "LeaseExpirationTs")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LeaseExpirationTs") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ApiLeaseRequestBodyMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ApiLeaseRequestBodyMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiPauseResponse struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

type ApiPubSubCallbackMessage struct {
	AuthToken string `json:"auth_token,omitempty"`

	Topic string `json:"topic,omitempty"`

	UserData string `json:"user_data,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AuthToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AuthToken") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApiPubSubCallbackMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ApiPubSubCallbackMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiPutBatchRequestMessage struct {
	Builds []*ApiPutRequestMessage `json:"builds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Builds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Builds") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApiPutBatchRequestMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ApiPutBatchRequestMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiPutBatchResponseMessage struct {
	Error *ApiErrorMessage `json:"error,omitempty"`

	Results []*ApiPutBatchResponseMessageOneResult `json:"results,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Error") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Error") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApiPutBatchResponseMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ApiPutBatchResponseMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiPutBatchResponseMessageOneResult struct {
	// Build: Describes model.Build, see its docstring.
	Build *ApiCommonBuildMessage `json:"build,omitempty"`

	ClientOperationId string `json:"client_operation_id,omitempty"`

	Error *ApiErrorMessage `json:"error,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Build") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Build") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApiPutBatchResponseMessageOneResult) MarshalJSON() ([]byte, error) {
	type NoMethod ApiPutBatchResponseMessageOneResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiPutRequestMessage struct {
	Bucket string `json:"bucket,omitempty"`

	// Possible values:
	//   "AUTO"
	//   "CANARY"
	//   "PROD"
	CanaryPreference string `json:"canary_preference,omitempty"`

	ClientOperationId string `json:"client_operation_id,omitempty"`

	Experimental bool `json:"experimental,omitempty"`

	LeaseExpirationTs int64 `json:"lease_expiration_ts,omitempty,string"`

	ParametersJson string `json:"parameters_json,omitempty"`

	PubsubCallback *ApiPubSubCallbackMessage `json:"pubsub_callback,omitempty"`

	Tags []string `json:"tags,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Bucket") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Bucket") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApiPutRequestMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ApiPutRequestMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiRetryRequestMessage struct {
	ClientOperationId string `json:"client_operation_id,omitempty"`

	LeaseExpirationTs int64 `json:"lease_expiration_ts,omitempty,string"`

	PubsubCallback *ApiPubSubCallbackMessage `json:"pubsub_callback,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ClientOperationId")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ClientOperationId") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ApiRetryRequestMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ApiRetryRequestMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiSearchResponseMessage struct {
	// Builds: Describes model.Build, see its docstring.
	Builds []*ApiCommonBuildMessage `json:"builds,omitempty"`

	Error *ApiErrorMessage `json:"error,omitempty"`

	NextCursor string `json:"next_cursor,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Builds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Builds") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApiSearchResponseMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ApiSearchResponseMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiStartRequestBodyMessage struct {
	Canary bool `json:"canary,omitempty"`

	LeaseKey int64 `json:"lease_key,omitempty,string"`

	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Canary") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Canary") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApiStartRequestBodyMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ApiStartRequestBodyMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ApiSucceedRequestBodyMessage struct {
	LeaseKey int64 `json:"lease_key,omitempty,string"`

	NewTags []string `json:"new_tags,omitempty"`

	ResultDetailsJson string `json:"result_details_json,omitempty"`

	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LeaseKey") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LeaseKey") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApiSucceedRequestBodyMessage) MarshalJSON() ([]byte, error) {
	type NoMethod ApiSucceedRequestBodyMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "buildbucket.backfill_tag_index":

type BackfillTagIndexCall struct {
	s          *Service
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// BackfillTagIndex: Backfills TagIndex entites from builds.
func (s *Service) BackfillTagIndex(tagKey string) *BackfillTagIndexCall {
	c := &BackfillTagIndexCall{s: s, urlParams_: make(gensupport.URLParams)}
	c.urlParams_.Set("tag_key", tagKey)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BackfillTagIndexCall) Fields(s ...googleapi.Field) *BackfillTagIndexCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BackfillTagIndexCall) Context(ctx context.Context) *BackfillTagIndexCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *BackfillTagIndexCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *BackfillTagIndexCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "backfill_tag_index")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "buildbucket.backfill_tag_index" call.
func (c *BackfillTagIndexCall) Do(opts ...googleapi.CallOption) error {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Backfills TagIndex entites from builds.",
	//   "httpMethod": "POST",
	//   "id": "buildbucket.backfill_tag_index",
	//   "parameterOrder": [
	//     "tag_key"
	//   ],
	//   "parameters": {
	//     "tag_key": {
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "backfill_tag_index",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "buildbucket.cancel":

type CancelCall struct {
	s                           *Service
	id                          int64
	apicancelrequestbodymessage *ApiCancelRequestBodyMessage
	urlParams_                  gensupport.URLParams
	ctx_                        context.Context
	header_                     http.Header
}

// Cancel: Cancels a build.
func (s *Service) Cancel(id int64, apicancelrequestbodymessage *ApiCancelRequestBodyMessage) *CancelCall {
	c := &CancelCall{s: s, urlParams_: make(gensupport.URLParams)}
	c.id = id
	c.apicancelrequestbodymessage = apicancelrequestbodymessage
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CancelCall) Fields(s ...googleapi.Field) *CancelCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CancelCall) Context(ctx context.Context) *CancelCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CancelCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CancelCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.apicancelrequestbodymessage)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "builds/{id}/cancel")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"id": strconv.FormatInt(c.id, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "buildbucket.cancel" call.
// Exactly one of *ApiBuildResponseMessage or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ApiBuildResponseMessage.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *CancelCall) Do(opts ...googleapi.CallOption) (*ApiBuildResponseMessage, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ApiBuildResponseMessage{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Cancels a build.",
	//   "httpMethod": "POST",
	//   "id": "buildbucket.cancel",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "builds/{id}/cancel",
	//   "request": {
	//     "$ref": "ApiCancelRequestBodyMessage",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "ApiBuildResponseMessage"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "buildbucket.cancel_batch":

type CancelBatchCall struct {
	s                            *Service
	apicancelbatchrequestmessage *ApiCancelBatchRequestMessage
	urlParams_                   gensupport.URLParams
	ctx_                         context.Context
	header_                      http.Header
}

// CancelBatch: Cancels builds.
func (s *Service) CancelBatch(apicancelbatchrequestmessage *ApiCancelBatchRequestMessage) *CancelBatchCall {
	c := &CancelBatchCall{s: s, urlParams_: make(gensupport.URLParams)}
	c.apicancelbatchrequestmessage = apicancelbatchrequestmessage
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CancelBatchCall) Fields(s ...googleapi.Field) *CancelBatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CancelBatchCall) Context(ctx context.Context) *CancelBatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CancelBatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CancelBatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.apicancelbatchrequestmessage)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "builds/cancel")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "buildbucket.cancel_batch" call.
// Exactly one of *ApiCancelBatchResponseMessage or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ApiCancelBatchResponseMessage.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *CancelBatchCall) Do(opts ...googleapi.CallOption) (*ApiCancelBatchResponseMessage, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ApiCancelBatchResponseMessage{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Cancels builds.",
	//   "httpMethod": "POST",
	//   "id": "buildbucket.cancel_batch",
	//   "path": "builds/cancel",
	//   "request": {
	//     "$ref": "ApiCancelBatchRequestMessage",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "ApiCancelBatchResponseMessage"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "buildbucket.delete_many_builds":

type DeleteManyBuildsCall struct {
	s          *Service
	bucket     string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// DeleteManyBuilds: Deletes scheduled or started builds in a bucket.
func (s *Service) DeleteManyBuilds(bucket string, status string) *DeleteManyBuildsCall {
	c := &DeleteManyBuildsCall{s: s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.urlParams_.Set("status", status)
	return c
}

// CreatedBy sets the optional parameter "created_by":
func (c *DeleteManyBuildsCall) CreatedBy(createdBy string) *DeleteManyBuildsCall {
	c.urlParams_.Set("created_by", createdBy)
	return c
}

// Tag sets the optional parameter "tag":
func (c *DeleteManyBuildsCall) Tag(tag ...string) *DeleteManyBuildsCall {
	c.urlParams_.SetMulti("tag", append([]string{}, tag...))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DeleteManyBuildsCall) Fields(s ...googleapi.Field) *DeleteManyBuildsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DeleteManyBuildsCall) Context(ctx context.Context) *DeleteManyBuildsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DeleteManyBuildsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *DeleteManyBuildsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "bucket/{bucket}/delete")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "buildbucket.delete_many_builds" call.
// Exactly one of *ApiDeleteManyBuildsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ApiDeleteManyBuildsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *DeleteManyBuildsCall) Do(opts ...googleapi.CallOption) (*ApiDeleteManyBuildsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ApiDeleteManyBuildsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes scheduled or started builds in a bucket.",
	//   "httpMethod": "POST",
	//   "id": "buildbucket.delete_many_builds",
	//   "parameterOrder": [
	//     "bucket",
	//     "status"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "created_by": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "status": {
	//       "enum": [
	//         "COMPLETED",
	//         "SCHEDULED",
	//         "STARTED"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tag": {
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "bucket/{bucket}/delete",
	//   "response": {
	//     "$ref": "ApiDeleteManyBuildsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "buildbucket.fail":

type FailCall struct {
	s                         *Service
	id                        int64
	apifailrequestbodymessage *ApiFailRequestBodyMessage
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
	header_                   http.Header
}

// Fail: Marks a build as failed.
func (s *Service) Fail(id int64, apifailrequestbodymessage *ApiFailRequestBodyMessage) *FailCall {
	c := &FailCall{s: s, urlParams_: make(gensupport.URLParams)}
	c.id = id
	c.apifailrequestbodymessage = apifailrequestbodymessage
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FailCall) Fields(s ...googleapi.Field) *FailCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *FailCall) Context(ctx context.Context) *FailCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *FailCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *FailCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.apifailrequestbodymessage)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "builds/{id}/fail")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"id": strconv.FormatInt(c.id, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "buildbucket.fail" call.
// Exactly one of *ApiBuildResponseMessage or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ApiBuildResponseMessage.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *FailCall) Do(opts ...googleapi.CallOption) (*ApiBuildResponseMessage, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ApiBuildResponseMessage{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Marks a build as failed.",
	//   "httpMethod": "POST",
	//   "id": "buildbucket.fail",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "builds/{id}/fail",
	//   "request": {
	//     "$ref": "ApiFailRequestBodyMessage",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "ApiBuildResponseMessage"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "buildbucket.get":

type GetCall struct {
	s            *Service
	id           int64
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Returns a build by id.
func (s *Service) Get(id int64) *GetCall {
	c := &GetCall{s: s, urlParams_: make(gensupport.URLParams)}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GetCall) Fields(s ...googleapi.Field) *GetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *GetCall) IfNoneMatch(entityTag string) *GetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GetCall) Context(ctx context.Context) *GetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "builds/{id}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"id": strconv.FormatInt(c.id, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "buildbucket.get" call.
// Exactly one of *ApiBuildResponseMessage or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ApiBuildResponseMessage.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *GetCall) Do(opts ...googleapi.CallOption) (*ApiBuildResponseMessage, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ApiBuildResponseMessage{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a build by id.",
	//   "httpMethod": "GET",
	//   "id": "buildbucket.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "builds/{id}",
	//   "response": {
	//     "$ref": "ApiBuildResponseMessage"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "buildbucket.get_bucket":

type GetBucketCall struct {
	s            *Service
	bucket       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetBucket: Returns bucket information.
func (s *Service) GetBucket(bucket string) *GetBucketCall {
	c := &GetBucketCall{s: s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GetBucketCall) Fields(s ...googleapi.Field) *GetBucketCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *GetBucketCall) IfNoneMatch(entityTag string) *GetBucketCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GetBucketCall) Context(ctx context.Context) *GetBucketCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GetBucketCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GetBucketCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "buckets/{bucket}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "buildbucket.get_bucket" call.
// Exactly one of *ApiBucketMessage or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ApiBucketMessage.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *GetBucketCall) Do(opts ...googleapi.CallOption) (*ApiBucketMessage, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ApiBucketMessage{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns bucket information.",
	//   "httpMethod": "GET",
	//   "id": "buildbucket.get_bucket",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "buckets/{bucket}",
	//   "response": {
	//     "$ref": "ApiBucketMessage"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "buildbucket.heartbeat":

type HeartbeatCall struct {
	s                              *Service
	id                             int64
	apiheartbeatrequestbodymessage *ApiHeartbeatRequestBodyMessage
	urlParams_                     gensupport.URLParams
	ctx_                           context.Context
	header_                        http.Header
}

// Heartbeat: Updates build lease.
func (s *Service) Heartbeat(id int64, apiheartbeatrequestbodymessage *ApiHeartbeatRequestBodyMessage) *HeartbeatCall {
	c := &HeartbeatCall{s: s, urlParams_: make(gensupport.URLParams)}
	c.id = id
	c.apiheartbeatrequestbodymessage = apiheartbeatrequestbodymessage
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *HeartbeatCall) Fields(s ...googleapi.Field) *HeartbeatCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *HeartbeatCall) Context(ctx context.Context) *HeartbeatCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *HeartbeatCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *HeartbeatCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.apiheartbeatrequestbodymessage)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "builds/{id}/heartbeat")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"id": strconv.FormatInt(c.id, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "buildbucket.heartbeat" call.
// Exactly one of *ApiBuildResponseMessage or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ApiBuildResponseMessage.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *HeartbeatCall) Do(opts ...googleapi.CallOption) (*ApiBuildResponseMessage, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ApiBuildResponseMessage{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates build lease.",
	//   "httpMethod": "POST",
	//   "id": "buildbucket.heartbeat",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "builds/{id}/heartbeat",
	//   "request": {
	//     "$ref": "ApiHeartbeatRequestBodyMessage",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "ApiBuildResponseMessage"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "buildbucket.heartbeat_batch":

type HeartbeatBatchCall struct {
	s                               *Service
	apiheartbeatbatchrequestmessage *ApiHeartbeatBatchRequestMessage
	urlParams_                      gensupport.URLParams
	ctx_                            context.Context
	header_                         http.Header
}

// HeartbeatBatch: Updates multiple build leases.
func (s *Service) HeartbeatBatch(apiheartbeatbatchrequestmessage *ApiHeartbeatBatchRequestMessage) *HeartbeatBatchCall {
	c := &HeartbeatBatchCall{s: s, urlParams_: make(gensupport.URLParams)}
	c.apiheartbeatbatchrequestmessage = apiheartbeatbatchrequestmessage
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *HeartbeatBatchCall) Fields(s ...googleapi.Field) *HeartbeatBatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *HeartbeatBatchCall) Context(ctx context.Context) *HeartbeatBatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *HeartbeatBatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *HeartbeatBatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.apiheartbeatbatchrequestmessage)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "heartbeat")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "buildbucket.heartbeat_batch" call.
// Exactly one of *ApiHeartbeatBatchResponseMessage or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ApiHeartbeatBatchResponseMessage.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *HeartbeatBatchCall) Do(opts ...googleapi.CallOption) (*ApiHeartbeatBatchResponseMessage, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ApiHeartbeatBatchResponseMessage{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates multiple build leases.",
	//   "httpMethod": "POST",
	//   "id": "buildbucket.heartbeat_batch",
	//   "path": "heartbeat",
	//   "request": {
	//     "$ref": "ApiHeartbeatBatchRequestMessage",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "ApiHeartbeatBatchResponseMessage"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "buildbucket.lease":

type LeaseCall struct {
	s                          *Service
	id                         int64
	apileaserequestbodymessage *ApiLeaseRequestBodyMessage
	urlParams_                 gensupport.URLParams
	ctx_                       context.Context
	header_                    http.Header
}

// Lease: Leases a build. Response may contain an error.
func (s *Service) Lease(id int64, apileaserequestbodymessage *ApiLeaseRequestBodyMessage) *LeaseCall {
	c := &LeaseCall{s: s, urlParams_: make(gensupport.URLParams)}
	c.id = id
	c.apileaserequestbodymessage = apileaserequestbodymessage
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LeaseCall) Fields(s ...googleapi.Field) *LeaseCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *LeaseCall) Context(ctx context.Context) *LeaseCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *LeaseCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *LeaseCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.apileaserequestbodymessage)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "builds/{id}/lease")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"id": strconv.FormatInt(c.id, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "buildbucket.lease" call.
// Exactly one of *ApiBuildResponseMessage or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ApiBuildResponseMessage.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *LeaseCall) Do(opts ...googleapi.CallOption) (*ApiBuildResponseMessage, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ApiBuildResponseMessage{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Leases a build. Response may contain an error.",
	//   "httpMethod": "POST",
	//   "id": "buildbucket.lease",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "builds/{id}/lease",
	//   "request": {
	//     "$ref": "ApiLeaseRequestBodyMessage",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "ApiBuildResponseMessage"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "buildbucket.pause":

type PauseCall struct {
	s          *Service
	bucket     string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Pause: Pauses or unpause a bucket.
func (s *Service) Pause(bucket string, isPaused bool) *PauseCall {
	c := &PauseCall{s: s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.urlParams_.Set("is_paused", fmt.Sprint(isPaused))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PauseCall) Fields(s ...googleapi.Field) *PauseCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PauseCall) Context(ctx context.Context) *PauseCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PauseCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PauseCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "buckets/{bucket}/pause")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "buildbucket.pause" call.
// Exactly one of *ApiPauseResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ApiPauseResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *PauseCall) Do(opts ...googleapi.CallOption) (*ApiPauseResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ApiPauseResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Pauses or unpause a bucket.",
	//   "httpMethod": "POST",
	//   "id": "buildbucket.pause",
	//   "parameterOrder": [
	//     "bucket",
	//     "is_paused"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "is_paused": {
	//       "location": "query",
	//       "required": true,
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "buckets/{bucket}/pause",
	//   "response": {
	//     "$ref": "ApiPauseResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "buildbucket.peek":

type PeekCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Peek: Returns available builds.
func (s *Service) Peek() *PeekCall {
	c := &PeekCall{s: s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Bucket sets the optional parameter "bucket":
func (c *PeekCall) Bucket(bucket ...string) *PeekCall {
	c.urlParams_.SetMulti("bucket", append([]string{}, bucket...))
	return c
}

// MaxBuilds sets the optional parameter "max_builds":
func (c *PeekCall) MaxBuilds(maxBuilds int64) *PeekCall {
	c.urlParams_.Set("max_builds", fmt.Sprint(maxBuilds))
	return c
}

// StartCursor sets the optional parameter "start_cursor":
func (c *PeekCall) StartCursor(startCursor string) *PeekCall {
	c.urlParams_.Set("start_cursor", startCursor)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PeekCall) Fields(s ...googleapi.Field) *PeekCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *PeekCall) IfNoneMatch(entityTag string) *PeekCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PeekCall) Context(ctx context.Context) *PeekCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PeekCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PeekCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "peek")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "buildbucket.peek" call.
// Exactly one of *ApiSearchResponseMessage or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ApiSearchResponseMessage.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *PeekCall) Do(opts ...googleapi.CallOption) (*ApiSearchResponseMessage, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ApiSearchResponseMessage{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns available builds.",
	//   "httpMethod": "GET",
	//   "id": "buildbucket.peek",
	//   "parameters": {
	//     "bucket": {
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "max_builds": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "start_cursor": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "peek",
	//   "response": {
	//     "$ref": "ApiSearchResponseMessage"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "buildbucket.put":

type PutCall struct {
	s                    *Service
	apiputrequestmessage *ApiPutRequestMessage
	urlParams_           gensupport.URLParams
	ctx_                 context.Context
	header_              http.Header
}

// Put: Creates a new build.
func (s *Service) Put(apiputrequestmessage *ApiPutRequestMessage) *PutCall {
	c := &PutCall{s: s, urlParams_: make(gensupport.URLParams)}
	c.apiputrequestmessage = apiputrequestmessage
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PutCall) Fields(s ...googleapi.Field) *PutCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PutCall) Context(ctx context.Context) *PutCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PutCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PutCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.apiputrequestmessage)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "builds")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "buildbucket.put" call.
// Exactly one of *ApiBuildResponseMessage or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ApiBuildResponseMessage.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *PutCall) Do(opts ...googleapi.CallOption) (*ApiBuildResponseMessage, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ApiBuildResponseMessage{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new build.",
	//   "httpMethod": "PUT",
	//   "id": "buildbucket.put",
	//   "path": "builds",
	//   "request": {
	//     "$ref": "ApiPutRequestMessage",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "ApiBuildResponseMessage"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "buildbucket.put_batch":

type PutBatchCall struct {
	s                         *Service
	apiputbatchrequestmessage *ApiPutBatchRequestMessage
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
	header_                   http.Header
}

// PutBatch: Creates builds.
func (s *Service) PutBatch(apiputbatchrequestmessage *ApiPutBatchRequestMessage) *PutBatchCall {
	c := &PutBatchCall{s: s, urlParams_: make(gensupport.URLParams)}
	c.apiputbatchrequestmessage = apiputbatchrequestmessage
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PutBatchCall) Fields(s ...googleapi.Field) *PutBatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PutBatchCall) Context(ctx context.Context) *PutBatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PutBatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PutBatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.apiputbatchrequestmessage)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "builds/batch")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "buildbucket.put_batch" call.
// Exactly one of *ApiPutBatchResponseMessage or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ApiPutBatchResponseMessage.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *PutBatchCall) Do(opts ...googleapi.CallOption) (*ApiPutBatchResponseMessage, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ApiPutBatchResponseMessage{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates builds.",
	//   "httpMethod": "PUT",
	//   "id": "buildbucket.put_batch",
	//   "path": "builds/batch",
	//   "request": {
	//     "$ref": "ApiPutBatchRequestMessage",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "ApiPutBatchResponseMessage"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "buildbucket.reput_builds":

type ReputBuildsCall struct {
	s          *Service
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// ReputBuilds: Reputs every build, recomputing its properties.
func (s *Service) ReputBuilds() *ReputBuildsCall {
	c := &ReputBuildsCall{s: s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReputBuildsCall) Fields(s ...googleapi.Field) *ReputBuildsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ReputBuildsCall) Context(ctx context.Context) *ReputBuildsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ReputBuildsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ReputBuildsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "reput_builds")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "buildbucket.reput_builds" call.
func (c *ReputBuildsCall) Do(opts ...googleapi.CallOption) error {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Reputs every build, recomputing its properties.",
	//   "httpMethod": "POST",
	//   "id": "buildbucket.reput_builds",
	//   "path": "reput_builds",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "buildbucket.reset":

type ResetCall struct {
	s          *Service
	id         int64
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Reset: Forcibly unleases a build and resets its state to SCHEDULED.
func (s *Service) Reset(id int64) *ResetCall {
	c := &ResetCall{s: s, urlParams_: make(gensupport.URLParams)}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ResetCall) Fields(s ...googleapi.Field) *ResetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ResetCall) Context(ctx context.Context) *ResetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ResetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ResetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "builds/{id}/reset")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"id": strconv.FormatInt(c.id, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "buildbucket.reset" call.
// Exactly one of *ApiBuildResponseMessage or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ApiBuildResponseMessage.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ResetCall) Do(opts ...googleapi.CallOption) (*ApiBuildResponseMessage, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ApiBuildResponseMessage{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Forcibly unleases a build and resets its state to SCHEDULED.",
	//   "httpMethod": "POST",
	//   "id": "buildbucket.reset",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "builds/{id}/reset",
	//   "response": {
	//     "$ref": "ApiBuildResponseMessage"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "buildbucket.retry":

type RetryCall struct {
	s                      *Service
	id                     int64
	apiretryrequestmessage *ApiRetryRequestMessage
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// Retry: Retries an existing build.
func (s *Service) Retry(id int64, apiretryrequestmessage *ApiRetryRequestMessage) *RetryCall {
	c := &RetryCall{s: s, urlParams_: make(gensupport.URLParams)}
	c.id = id
	c.apiretryrequestmessage = apiretryrequestmessage
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RetryCall) Fields(s ...googleapi.Field) *RetryCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RetryCall) Context(ctx context.Context) *RetryCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *RetryCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *RetryCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.apiretryrequestmessage)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "builds/{id}/retry")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"id": strconv.FormatInt(c.id, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "buildbucket.retry" call.
// Exactly one of *ApiBuildResponseMessage or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ApiBuildResponseMessage.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *RetryCall) Do(opts ...googleapi.CallOption) (*ApiBuildResponseMessage, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ApiBuildResponseMessage{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retries an existing build.",
	//   "httpMethod": "PUT",
	//   "id": "buildbucket.retry",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "builds/{id}/retry",
	//   "request": {
	//     "$ref": "ApiRetryRequestMessage",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "ApiBuildResponseMessage"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "buildbucket.search":

type SearchCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Search: Searches for builds.
func (s *Service) Search() *SearchCall {
	c := &SearchCall{s: s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Bucket sets the optional parameter "bucket":
func (c *SearchCall) Bucket(bucket ...string) *SearchCall {
	c.urlParams_.SetMulti("bucket", append([]string{}, bucket...))
	return c
}

// Canary sets the optional parameter "canary":
func (c *SearchCall) Canary(canary bool) *SearchCall {
	c.urlParams_.Set("canary", fmt.Sprint(canary))
	return c
}

// CancelationReason sets the optional parameter "cancelation_reason":
//
// Possible values:
//   "CANCELED_EXPLICITLY"
//   "TIMEOUT"
func (c *SearchCall) CancelationReason(cancelationReason string) *SearchCall {
	c.urlParams_.Set("cancelation_reason", cancelationReason)
	return c
}

// CreatedBy sets the optional parameter "created_by":
func (c *SearchCall) CreatedBy(createdBy string) *SearchCall {
	c.urlParams_.Set("created_by", createdBy)
	return c
}

// CreationTsHigh sets the optional parameter "creation_ts_high":
func (c *SearchCall) CreationTsHigh(creationTsHigh int64) *SearchCall {
	c.urlParams_.Set("creation_ts_high", fmt.Sprint(creationTsHigh))
	return c
}

// CreationTsLow sets the optional parameter "creation_ts_low":
func (c *SearchCall) CreationTsLow(creationTsLow int64) *SearchCall {
	c.urlParams_.Set("creation_ts_low", fmt.Sprint(creationTsLow))
	return c
}

// FailureReason sets the optional parameter "failure_reason":
//
// Possible values:
//   "BUILDBUCKET_FAILURE"
//   "BUILD_FAILURE"
//   "INFRA_FAILURE"
//   "INVALID_BUILD_DEFINITION"
func (c *SearchCall) FailureReason(failureReason string) *SearchCall {
	c.urlParams_.Set("failure_reason", failureReason)
	return c
}

// IncludeExperimental sets the optional parameter
// "include_experimental":
func (c *SearchCall) IncludeExperimental(includeExperimental bool) *SearchCall {
	c.urlParams_.Set("include_experimental", fmt.Sprint(includeExperimental))
	return c
}

// MaxBuilds sets the optional parameter "max_builds":
func (c *SearchCall) MaxBuilds(maxBuilds int64) *SearchCall {
	c.urlParams_.Set("max_builds", fmt.Sprint(maxBuilds))
	return c
}

// Result sets the optional parameter "result":
//
// Possible values:
//   "CANCELED"
//   "FAILURE"
//   "SUCCESS"
func (c *SearchCall) Result(result string) *SearchCall {
	c.urlParams_.Set("result", result)
	return c
}

// RetryOf sets the optional parameter "retry_of":
func (c *SearchCall) RetryOf(retryOf int64) *SearchCall {
	c.urlParams_.Set("retry_of", fmt.Sprint(retryOf))
	return c
}

// StartCursor sets the optional parameter "start_cursor":
func (c *SearchCall) StartCursor(startCursor string) *SearchCall {
	c.urlParams_.Set("start_cursor", startCursor)
	return c
}

// Status sets the optional parameter "status":
//
// Possible values:
//   "COMPLETED"
//   "INCOMPLETE"
//   "SCHEDULED"
//   "STARTED"
func (c *SearchCall) Status(status string) *SearchCall {
	c.urlParams_.Set("status", status)
	return c
}

// Tag sets the optional parameter "tag":
func (c *SearchCall) Tag(tag ...string) *SearchCall {
	c.urlParams_.SetMulti("tag", append([]string{}, tag...))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SearchCall) Fields(s ...googleapi.Field) *SearchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *SearchCall) IfNoneMatch(entityTag string) *SearchCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SearchCall) Context(ctx context.Context) *SearchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SearchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SearchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "search")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "buildbucket.search" call.
// Exactly one of *ApiSearchResponseMessage or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ApiSearchResponseMessage.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *SearchCall) Do(opts ...googleapi.CallOption) (*ApiSearchResponseMessage, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ApiSearchResponseMessage{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Searches for builds.",
	//   "httpMethod": "GET",
	//   "id": "buildbucket.search",
	//   "parameters": {
	//     "bucket": {
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "canary": {
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "cancelation_reason": {
	//       "enum": [
	//         "CANCELED_EXPLICITLY",
	//         "TIMEOUT"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "created_by": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "creation_ts_high": {
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "creation_ts_low": {
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "failure_reason": {
	//       "enum": [
	//         "BUILDBUCKET_FAILURE",
	//         "BUILD_FAILURE",
	//         "INFRA_FAILURE",
	//         "INVALID_BUILD_DEFINITION"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "include_experimental": {
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "max_builds": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "result": {
	//       "enum": [
	//         "CANCELED",
	//         "FAILURE",
	//         "SUCCESS"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "retry_of": {
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "start_cursor": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "status": {
	//       "enum": [
	//         "COMPLETED",
	//         "INCOMPLETE",
	//         "SCHEDULED",
	//         "STARTED"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "tag": {
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "search",
	//   "response": {
	//     "$ref": "ApiSearchResponseMessage"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "buildbucket.start":

type StartCall struct {
	s                          *Service
	id                         int64
	apistartrequestbodymessage *ApiStartRequestBodyMessage
	urlParams_                 gensupport.URLParams
	ctx_                       context.Context
	header_                    http.Header
}

// Start: Marks a build as started.
func (s *Service) Start(id int64, apistartrequestbodymessage *ApiStartRequestBodyMessage) *StartCall {
	c := &StartCall{s: s, urlParams_: make(gensupport.URLParams)}
	c.id = id
	c.apistartrequestbodymessage = apistartrequestbodymessage
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StartCall) Fields(s ...googleapi.Field) *StartCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *StartCall) Context(ctx context.Context) *StartCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *StartCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *StartCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.apistartrequestbodymessage)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "builds/{id}/start")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"id": strconv.FormatInt(c.id, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "buildbucket.start" call.
// Exactly one of *ApiBuildResponseMessage or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ApiBuildResponseMessage.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *StartCall) Do(opts ...googleapi.CallOption) (*ApiBuildResponseMessage, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ApiBuildResponseMessage{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Marks a build as started.",
	//   "httpMethod": "POST",
	//   "id": "buildbucket.start",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "builds/{id}/start",
	//   "request": {
	//     "$ref": "ApiStartRequestBodyMessage",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "ApiBuildResponseMessage"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "buildbucket.succeed":

type SucceedCall struct {
	s                            *Service
	id                           int64
	apisucceedrequestbodymessage *ApiSucceedRequestBodyMessage
	urlParams_                   gensupport.URLParams
	ctx_                         context.Context
	header_                      http.Header
}

// Succeed: Marks a build as succeeded.
func (s *Service) Succeed(id int64, apisucceedrequestbodymessage *ApiSucceedRequestBodyMessage) *SucceedCall {
	c := &SucceedCall{s: s, urlParams_: make(gensupport.URLParams)}
	c.id = id
	c.apisucceedrequestbodymessage = apisucceedrequestbodymessage
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SucceedCall) Fields(s ...googleapi.Field) *SucceedCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SucceedCall) Context(ctx context.Context) *SucceedCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SucceedCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SucceedCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.apisucceedrequestbodymessage)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "builds/{id}/succeed")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"id": strconv.FormatInt(c.id, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "buildbucket.succeed" call.
// Exactly one of *ApiBuildResponseMessage or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ApiBuildResponseMessage.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *SucceedCall) Do(opts ...googleapi.CallOption) (*ApiBuildResponseMessage, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ApiBuildResponseMessage{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Marks a build as succeeded.",
	//   "httpMethod": "POST",
	//   "id": "buildbucket.succeed",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "builds/{id}/succeed",
	//   "request": {
	//     "$ref": "ApiSucceedRequestBodyMessage",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "ApiBuildResponseMessage"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}
