// Code generated by go-swagger; DO NOT EDIT.

package contexts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteContextEnvVarParams creates a new DeleteContextEnvVarParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteContextEnvVarParams() *DeleteContextEnvVarParams {
	return &DeleteContextEnvVarParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteContextEnvVarParamsWithTimeout creates a new DeleteContextEnvVarParams object
// with the ability to set a timeout on a request.
func NewDeleteContextEnvVarParamsWithTimeout(timeout time.Duration) *DeleteContextEnvVarParams {
	return &DeleteContextEnvVarParams{
		timeout: timeout,
	}
}

// NewDeleteContextEnvVarParamsWithContext creates a new DeleteContextEnvVarParams object
// with the ability to set a context for a request.
func NewDeleteContextEnvVarParamsWithContext(ctx context.Context) *DeleteContextEnvVarParams {
	return &DeleteContextEnvVarParams{
		Context: ctx,
	}
}

// NewDeleteContextEnvVarParamsWithHTTPClient creates a new DeleteContextEnvVarParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteContextEnvVarParamsWithHTTPClient(client *http.Client) *DeleteContextEnvVarParams {
	return &DeleteContextEnvVarParams{
		HTTPClient: client,
	}
}

/*
DeleteContextEnvVarParams contains all the parameters to send to the API endpoint

	for the delete context env var operation.

	Typically these are written to a http.Request.
*/
type DeleteContextEnvVarParams struct {

	/* ID.

	   ID of the context

	   Format: uuid
	*/
	ID strfmt.UUID

	/* Name.

	   The name of the environment variable
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete context env var params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteContextEnvVarParams) WithDefaults() *DeleteContextEnvVarParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete context env var params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteContextEnvVarParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete context env var params
func (o *DeleteContextEnvVarParams) WithTimeout(timeout time.Duration) *DeleteContextEnvVarParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete context env var params
func (o *DeleteContextEnvVarParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete context env var params
func (o *DeleteContextEnvVarParams) WithContext(ctx context.Context) *DeleteContextEnvVarParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete context env var params
func (o *DeleteContextEnvVarParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete context env var params
func (o *DeleteContextEnvVarParams) WithHTTPClient(client *http.Client) *DeleteContextEnvVarParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete context env var params
func (o *DeleteContextEnvVarParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete context env var params
func (o *DeleteContextEnvVarParams) WithID(id strfmt.UUID) *DeleteContextEnvVarParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete context env var params
func (o *DeleteContextEnvVarParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithName adds the name to the delete context env var params
func (o *DeleteContextEnvVarParams) WithName(name string) *DeleteContextEnvVarParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete context env var params
func (o *DeleteContextEnvVarParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteContextEnvVarParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}