// Package userconfig provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package userconfig

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/google/uuid"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// ExtendedGroup defines model for ExtendedGroup.
type ExtendedGroup struct {
	ID      string   `json:"id"`
	Lessons []Lesson `json:"lessons"`
}

// Lesson defines model for Lesson.
type Lesson struct {
	ID   string `json:"id"`
	Kind string `json:"kind"`
	Name string `json:"name"`
}

// UserConfig defines model for UserConfig.
type UserConfig struct {
	Base *struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	} `json:"base"`
	Email          string          `json:"email"`
	ExtendedGroups []ExtendedGroup `json:"extendedGroups"`
	ID             uuid.UUID       `json:"id"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetUsersIdConfig request
	GetUsersIdConfig(ctx context.Context, id uuid.UUID, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PatchUsersIdConfig request
	PatchUsersIdConfig(ctx context.Context, id uuid.UUID, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetUsersIdConfig(ctx context.Context, id uuid.UUID, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUsersIdConfigRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PatchUsersIdConfig(ctx context.Context, id uuid.UUID, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPatchUsersIdConfigRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetUsersIdConfigRequest generates requests for GetUsersIdConfig
func NewGetUsersIdConfigRequest(server string, id uuid.UUID) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s/config", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPatchUsersIdConfigRequest generates requests for PatchUsersIdConfig
func NewPatchUsersIdConfigRequest(server string, id uuid.UUID) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s/config", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetUsersIdConfig request
	GetUsersIdConfigWithResponse(ctx context.Context, id uuid.UUID, reqEditors ...RequestEditorFn) (*GetUsersIdConfigResponse, error)

	// PatchUsersIdConfig request
	PatchUsersIdConfigWithResponse(ctx context.Context, id uuid.UUID, reqEditors ...RequestEditorFn) (*PatchUsersIdConfigResponse, error)
}

type GetUsersIdConfigResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *UserConfig
}

// Status returns HTTPResponse.Status
func (r GetUsersIdConfigResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetUsersIdConfigResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PatchUsersIdConfigResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r PatchUsersIdConfigResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PatchUsersIdConfigResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetUsersIdConfigWithResponse request returning *GetUsersIdConfigResponse
func (c *ClientWithResponses) GetUsersIdConfigWithResponse(ctx context.Context, id uuid.UUID, reqEditors ...RequestEditorFn) (*GetUsersIdConfigResponse, error) {
	rsp, err := c.GetUsersIdConfig(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUsersIdConfigResponse(rsp)
}

// PatchUsersIdConfigWithResponse request returning *PatchUsersIdConfigResponse
func (c *ClientWithResponses) PatchUsersIdConfigWithResponse(ctx context.Context, id uuid.UUID, reqEditors ...RequestEditorFn) (*PatchUsersIdConfigResponse, error) {
	rsp, err := c.PatchUsersIdConfig(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePatchUsersIdConfigResponse(rsp)
}

// ParseGetUsersIdConfigResponse parses an HTTP response from a GetUsersIdConfigWithResponse call
func ParseGetUsersIdConfigResponse(rsp *http.Response) (*GetUsersIdConfigResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetUsersIdConfigResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest UserConfig
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePatchUsersIdConfigResponse parses an HTTP response from a PatchUsersIdConfigWithResponse call
func ParsePatchUsersIdConfigResponse(rsp *http.Response) (*PatchUsersIdConfigResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PatchUsersIdConfigResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
