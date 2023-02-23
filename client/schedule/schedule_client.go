// Code generated by go-swagger; DO NOT EDIT.

package schedule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new schedule API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for schedule API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddSchedule(params *AddScheduleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddScheduleCreated, error)

	DeleteSchedule(params *DeleteScheduleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteScheduleOK, error)

	GetSchedule(params *GetScheduleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleOK, error)

	ListSchedules(params *ListSchedulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSchedulesOK, error)

	UpdateSchedule(params *UpdateScheduleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateScheduleOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddSchedule creates a schedule
*/
func (a *Client) AddSchedule(params *AddScheduleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddScheduleCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddScheduleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddSchedule",
		Method:             "POST",
		PathPattern:        "/project/{project-slug}/schedule",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddScheduleReader{formats: a.formats},
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
	success, ok := result.(*AddScheduleCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AddSchedule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteSchedule deletes a schedule

Delete a schedule by ID.
*/
func (a *Client) DeleteSchedule(params *DeleteScheduleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteScheduleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteScheduleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteSchedule",
		Method:             "DELETE",
		PathPattern:        "/schedule/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteScheduleReader{formats: a.formats},
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
	success, ok := result.(*DeleteScheduleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteSchedule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSchedule gets a schedule

Get a schedule by ID.
*/
func (a *Client) GetSchedule(params *GetScheduleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScheduleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSchedule",
		Method:             "GET",
		PathPattern:        "/schedule/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScheduleReader{formats: a.formats},
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
	success, ok := result.(*GetScheduleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSchedule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListSchedules lists schedules

Returns all schedules for this project.
*/
func (a *Client) ListSchedules(params *ListSchedulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSchedulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSchedulesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListSchedules",
		Method:             "GET",
		PathPattern:        "/project/{project-slug}/schedule",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSchedulesReader{formats: a.formats},
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
	success, ok := result.(*ListSchedulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListSchedules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateSchedule updates a schedule

Updates a schedule and returns the updatd schedule.
*/
func (a *Client) UpdateSchedule(params *UpdateScheduleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateScheduleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateScheduleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateSchedule",
		Method:             "PATCH",
		PathPattern:        "/schedule/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateScheduleReader{formats: a.formats},
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
	success, ok := result.(*UpdateScheduleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateSchedule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
