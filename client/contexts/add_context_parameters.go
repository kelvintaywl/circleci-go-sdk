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

	"github.com/kelvintaywl/circleci-go-sdk/models"
)

// NewAddContextParams creates a new AddContextParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddContextParams() *AddContextParams {
	return &AddContextParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddContextParamsWithTimeout creates a new AddContextParams object
// with the ability to set a timeout on a request.
func NewAddContextParamsWithTimeout(timeout time.Duration) *AddContextParams {
	return &AddContextParams{
		timeout: timeout,
	}
}

// NewAddContextParamsWithContext creates a new AddContextParams object
// with the ability to set a context for a request.
func NewAddContextParamsWithContext(ctx context.Context) *AddContextParams {
	return &AddContextParams{
		Context: ctx,
	}
}

// NewAddContextParamsWithHTTPClient creates a new AddContextParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddContextParamsWithHTTPClient(client *http.Client) *AddContextParams {
	return &AddContextParams{
		HTTPClient: client,
	}
}

/*
AddContextParams contains all the parameters to send to the API endpoint

	for the add context operation.

	Typically these are written to a http.Request.
*/
type AddContextParams struct {

	/* Body.

	   Context information (payload)
	*/
	Body *models.ContextPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add context params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddContextParams) WithDefaults() *AddContextParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add context params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddContextParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add context params
func (o *AddContextParams) WithTimeout(timeout time.Duration) *AddContextParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add context params
func (o *AddContextParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add context params
func (o *AddContextParams) WithContext(ctx context.Context) *AddContextParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add context params
func (o *AddContextParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add context params
func (o *AddContextParams) WithHTTPClient(client *http.Client) *AddContextParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add context params
func (o *AddContextParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add context params
func (o *AddContextParams) WithBody(body *models.ContextPayload) *AddContextParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add context params
func (o *AddContextParams) SetBody(body *models.ContextPayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddContextParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}