// Code generated by go-swagger; DO NOT EDIT.

package headscale_service

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

// NewHeadscaleServiceGetUserParams creates a new HeadscaleServiceGetUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeadscaleServiceGetUserParams() *HeadscaleServiceGetUserParams {
	return &HeadscaleServiceGetUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeadscaleServiceGetUserParamsWithTimeout creates a new HeadscaleServiceGetUserParams object
// with the ability to set a timeout on a request.
func NewHeadscaleServiceGetUserParamsWithTimeout(timeout time.Duration) *HeadscaleServiceGetUserParams {
	return &HeadscaleServiceGetUserParams{
		timeout: timeout,
	}
}

// NewHeadscaleServiceGetUserParamsWithContext creates a new HeadscaleServiceGetUserParams object
// with the ability to set a context for a request.
func NewHeadscaleServiceGetUserParamsWithContext(ctx context.Context) *HeadscaleServiceGetUserParams {
	return &HeadscaleServiceGetUserParams{
		Context: ctx,
	}
}

// NewHeadscaleServiceGetUserParamsWithHTTPClient creates a new HeadscaleServiceGetUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeadscaleServiceGetUserParamsWithHTTPClient(client *http.Client) *HeadscaleServiceGetUserParams {
	return &HeadscaleServiceGetUserParams{
		HTTPClient: client,
	}
}

/*
HeadscaleServiceGetUserParams contains all the parameters to send to the API endpoint

	for the headscale service get user operation.

	Typically these are written to a http.Request.
*/
type HeadscaleServiceGetUserParams struct {

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the headscale service get user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadscaleServiceGetUserParams) WithDefaults() *HeadscaleServiceGetUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the headscale service get user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadscaleServiceGetUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the headscale service get user params
func (o *HeadscaleServiceGetUserParams) WithTimeout(timeout time.Duration) *HeadscaleServiceGetUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the headscale service get user params
func (o *HeadscaleServiceGetUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the headscale service get user params
func (o *HeadscaleServiceGetUserParams) WithContext(ctx context.Context) *HeadscaleServiceGetUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the headscale service get user params
func (o *HeadscaleServiceGetUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the headscale service get user params
func (o *HeadscaleServiceGetUserParams) WithHTTPClient(client *http.Client) *HeadscaleServiceGetUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the headscale service get user params
func (o *HeadscaleServiceGetUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the headscale service get user params
func (o *HeadscaleServiceGetUserParams) WithName(name string) *HeadscaleServiceGetUserParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the headscale service get user params
func (o *HeadscaleServiceGetUserParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *HeadscaleServiceGetUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
