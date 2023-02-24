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

// NewListProjectCheckoutKeysParams creates a new ListProjectCheckoutKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectCheckoutKeysParams() *ListProjectCheckoutKeysParams {
	return &ListProjectCheckoutKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectCheckoutKeysParamsWithTimeout creates a new ListProjectCheckoutKeysParams object
// with the ability to set a timeout on a request.
func NewListProjectCheckoutKeysParamsWithTimeout(timeout time.Duration) *ListProjectCheckoutKeysParams {
	return &ListProjectCheckoutKeysParams{
		timeout: timeout,
	}
}

// NewListProjectCheckoutKeysParamsWithContext creates a new ListProjectCheckoutKeysParams object
// with the ability to set a context for a request.
func NewListProjectCheckoutKeysParamsWithContext(ctx context.Context) *ListProjectCheckoutKeysParams {
	return &ListProjectCheckoutKeysParams{
		Context: ctx,
	}
}

// NewListProjectCheckoutKeysParamsWithHTTPClient creates a new ListProjectCheckoutKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectCheckoutKeysParamsWithHTTPClient(client *http.Client) *ListProjectCheckoutKeysParams {
	return &ListProjectCheckoutKeysParams{
		HTTPClient: client,
	}
}

/*
ListProjectCheckoutKeysParams contains all the parameters to send to the API endpoint

	for the list project checkout keys operation.

	Typically these are written to a http.Request.
*/
type ListProjectCheckoutKeysParams struct {

	/* PageToken.

	   A token to retrieve the next page of results.
	*/
	PageToken string

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

// WithDefaults hydrates default values in the list project checkout keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectCheckoutKeysParams) WithDefaults() *ListProjectCheckoutKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project checkout keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectCheckoutKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project checkout keys params
func (o *ListProjectCheckoutKeysParams) WithTimeout(timeout time.Duration) *ListProjectCheckoutKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project checkout keys params
func (o *ListProjectCheckoutKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project checkout keys params
func (o *ListProjectCheckoutKeysParams) WithContext(ctx context.Context) *ListProjectCheckoutKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project checkout keys params
func (o *ListProjectCheckoutKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project checkout keys params
func (o *ListProjectCheckoutKeysParams) WithHTTPClient(client *http.Client) *ListProjectCheckoutKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project checkout keys params
func (o *ListProjectCheckoutKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageToken adds the pageToken to the list project checkout keys params
func (o *ListProjectCheckoutKeysParams) WithPageToken(pageToken string) *ListProjectCheckoutKeysParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the list project checkout keys params
func (o *ListProjectCheckoutKeysParams) SetPageToken(pageToken string) {
	o.PageToken = pageToken
}

// WithProjectSlug adds the projectSlug to the list project checkout keys params
func (o *ListProjectCheckoutKeysParams) WithProjectSlug(projectSlug string) *ListProjectCheckoutKeysParams {
	o.SetProjectSlug(projectSlug)
	return o
}

// SetProjectSlug adds the projectSlug to the list project checkout keys params
func (o *ListProjectCheckoutKeysParams) SetProjectSlug(projectSlug string) {
	o.ProjectSlug = projectSlug
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectCheckoutKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param page-token
	qrPageToken := o.PageToken
	qPageToken := qrPageToken
	if qPageToken != "" {

		if err := r.SetQueryParam("page-token", qPageToken); err != nil {
			return err
		}
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
