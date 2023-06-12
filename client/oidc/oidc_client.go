// Code generated by go-swagger; DO NOT EDIT.

package oidc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new oidc API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for oidc API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteOrgLevelClaims(params *DeleteOrgLevelClaimsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrgLevelClaimsOK, error)

	DeleteProjectLevelClaims(params *DeleteProjectLevelClaimsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteProjectLevelClaimsOK, error)

	GetOrgLevelClaims(params *GetOrgLevelClaimsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrgLevelClaimsOK, error)

	GetProjectLevelClaims(params *GetProjectLevelClaimsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProjectLevelClaimsOK, error)

	PatchOrgLevelClaims(params *PatchOrgLevelClaimsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchOrgLevelClaimsOK, error)

	PatchProjectLevelClaims(params *PatchProjectLevelClaimsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchProjectLevelClaimsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteOrgLevelClaims deletes org level claims

Deletes org-level custom claims of OIDC identity tokens.
*/
func (a *Client) DeleteOrgLevelClaims(params *DeleteOrgLevelClaimsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrgLevelClaimsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrgLevelClaimsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteOrgLevelClaims",
		Method:             "DELETE",
		PathPattern:        "/org/{org-id}/oidc-custom-claims",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOrgLevelClaimsReader{formats: a.formats},
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
	success, ok := result.(*DeleteOrgLevelClaimsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteOrgLevelClaims: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteProjectLevelClaims deletes project level claims

Deletes project-level custom claims of OIDC identity tokens.
*/
func (a *Client) DeleteProjectLevelClaims(params *DeleteProjectLevelClaimsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteProjectLevelClaimsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProjectLevelClaimsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteProjectLevelClaims",
		Method:             "DELETE",
		PathPattern:        "/org/{org-id}/projecct/{project-id}/oidc-custom-claims",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteProjectLevelClaimsReader{formats: a.formats},
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
	success, ok := result.(*DeleteProjectLevelClaimsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteProjectLevelClaims: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetOrgLevelClaims gets org level claims

Fetches org-level custom claims of OIDC identity tokens.
*/
func (a *Client) GetOrgLevelClaims(params *GetOrgLevelClaimsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrgLevelClaimsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrgLevelClaimsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetOrgLevelClaims",
		Method:             "GET",
		PathPattern:        "/org/{org-id}/oidc-custom-claims",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrgLevelClaimsReader{formats: a.formats},
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
	success, ok := result.(*GetOrgLevelClaimsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOrgLevelClaims: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetProjectLevelClaims gets project level claims

Fetches project-level custom claims of OIDC identity tokens.
*/
func (a *Client) GetProjectLevelClaims(params *GetProjectLevelClaimsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProjectLevelClaimsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProjectLevelClaimsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProjectLevelClaims",
		Method:             "GET",
		PathPattern:        "/org/{org-id}/projecct/{project-id}/oidc-custom-claims",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProjectLevelClaimsReader{formats: a.formats},
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
	success, ok := result.(*GetProjectLevelClaimsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetProjectLevelClaims: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchOrgLevelClaims patches org level claims

Creates/Updates org-level custom claims of OIDC identity tokens.
*/
func (a *Client) PatchOrgLevelClaims(params *PatchOrgLevelClaimsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchOrgLevelClaimsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchOrgLevelClaimsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchOrgLevelClaims",
		Method:             "PATCH",
		PathPattern:        "/org/{org-id}/oidc-custom-claims",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchOrgLevelClaimsReader{formats: a.formats},
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
	success, ok := result.(*PatchOrgLevelClaimsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchOrgLevelClaims: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchProjectLevelClaims patches project level claims

Creates/Updates project-level custom claims of OIDC identity tokens.
*/
func (a *Client) PatchProjectLevelClaims(params *PatchProjectLevelClaimsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchProjectLevelClaimsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchProjectLevelClaimsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchProjectLevelClaims",
		Method:             "PATCH",
		PathPattern:        "/org/{org-id}/projecct/{project-id}/oidc-custom-claims",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchProjectLevelClaimsReader{formats: a.formats},
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
	success, ok := result.(*PatchProjectLevelClaimsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchProjectLevelClaims: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
