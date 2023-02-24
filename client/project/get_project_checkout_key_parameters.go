// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewGetProjectCheckoutKeyParams creates a new GetProjectCheckoutKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectCheckoutKeyParams() *GetProjectCheckoutKeyParams {
	return &GetProjectCheckoutKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectCheckoutKeyParamsWithTimeout creates a new GetProjectCheckoutKeyParams object
// with the ability to set a timeout on a request.
func NewGetProjectCheckoutKeyParamsWithTimeout(timeout time.Duration) *GetProjectCheckoutKeyParams {
	return &GetProjectCheckoutKeyParams{
		timeout: timeout,
	}
}

// NewGetProjectCheckoutKeyParamsWithContext creates a new GetProjectCheckoutKeyParams object
// with the ability to set a context for a request.
func NewGetProjectCheckoutKeyParamsWithContext(ctx context.Context) *GetProjectCheckoutKeyParams {
	return &GetProjectCheckoutKeyParams{
		Context: ctx,
	}
}

// NewGetProjectCheckoutKeyParamsWithHTTPClient creates a new GetProjectCheckoutKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectCheckoutKeyParamsWithHTTPClient(client *http.Client) *GetProjectCheckoutKeyParams {
	return &GetProjectCheckoutKeyParams{
		HTTPClient: client,
	}
}

/*
GetProjectCheckoutKeyParams contains all the parameters to send to the API endpoint

	for the get project checkout key operation.

	Typically these are written to a http.Request.
*/
type GetProjectCheckoutKeyParams struct {

	/* Fingerprint.

	   An SSH key fingerprint
	*/
	Fingerprint string

	/* ProjectSlug.

	     Project slug in the form vcs-slug/org-name/repo-name.
	The / characters may be URL-escaped.
	Example: gh/CircleCI-Public/api-preview-docs

	*/
	ProjectSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get project checkout key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectCheckoutKeyParams) WithDefaults() *GetProjectCheckoutKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project checkout key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectCheckoutKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get project checkout key params
func (o *GetProjectCheckoutKeyParams) WithTimeout(timeout time.Duration) *GetProjectCheckoutKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project checkout key params
func (o *GetProjectCheckoutKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project checkout key params
func (o *GetProjectCheckoutKeyParams) WithContext(ctx context.Context) *GetProjectCheckoutKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project checkout key params
func (o *GetProjectCheckoutKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project checkout key params
func (o *GetProjectCheckoutKeyParams) WithHTTPClient(client *http.Client) *GetProjectCheckoutKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project checkout key params
func (o *GetProjectCheckoutKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFingerprint adds the fingerprint to the get project checkout key params
func (o *GetProjectCheckoutKeyParams) WithFingerprint(fingerprint string) *GetProjectCheckoutKeyParams {
	o.SetFingerprint(fingerprint)
	return o
}

// SetFingerprint adds the fingerprint to the get project checkout key params
func (o *GetProjectCheckoutKeyParams) SetFingerprint(fingerprint string) {
	o.Fingerprint = fingerprint
}

// WithProjectSlug adds the projectSlug to the get project checkout key params
func (o *GetProjectCheckoutKeyParams) WithProjectSlug(projectSlug string) *GetProjectCheckoutKeyParams {
	o.SetProjectSlug(projectSlug)
	return o
}

// SetProjectSlug adds the projectSlug to the get project checkout key params
func (o *GetProjectCheckoutKeyParams) SetProjectSlug(projectSlug string) {
	o.ProjectSlug = projectSlug
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectCheckoutKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fingerprint
	if err := r.SetPathParam("fingerprint", o.Fingerprint); err != nil {
		return err
	}

	// path param project-slug
	if err := r.SetPathParam("project-slug", o.ProjectSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
