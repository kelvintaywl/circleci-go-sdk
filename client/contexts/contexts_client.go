// Code generated by go-swagger; DO NOT EDIT.

package contexts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new contexts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for contexts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddContext(params *AddContextParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddContextOK, error)

	DeleteContext(params *DeleteContextParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteContextOK, error)

	GetContext(params *GetContextParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetContextOK, error)

	ListContexts(params *ListContextsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListContextsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddContext creates a context
*/
func (a *Client) AddContext(params *AddContextParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddContextOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddContextParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddContext",
		Method:             "POST",
		PathPattern:        "/context",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddContextReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddContextOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AddContext: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteContext deletes a context
*/
func (a *Client) DeleteContext(params *DeleteContextParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteContextOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteContextParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteContext",
		Method:             "DELETE",
		PathPattern:        "/context/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteContextReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteContextOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteContext: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetContext gets a context

Get a context by ID.
*/
func (a *Client) GetContext(params *GetContextParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetContextOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetContextParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetContext",
		Method:             "GET",
		PathPattern:        "/context/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetContextReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetContextOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetContext: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListContexts lists contexts for an owner

Returns all contexts for this owner.
*/
func (a *Client) ListContexts(params *ListContextsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListContextsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListContextsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListContexts",
		Method:             "GET",
		PathPattern:        "/context",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListContextsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListContextsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListContexts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
