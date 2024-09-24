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

// NewHeadscaleServiceGetPolicyParams creates a new HeadscaleServiceGetPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeadscaleServiceGetPolicyParams() *HeadscaleServiceGetPolicyParams {
	return &HeadscaleServiceGetPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeadscaleServiceGetPolicyParamsWithTimeout creates a new HeadscaleServiceGetPolicyParams object
// with the ability to set a timeout on a request.
func NewHeadscaleServiceGetPolicyParamsWithTimeout(timeout time.Duration) *HeadscaleServiceGetPolicyParams {
	return &HeadscaleServiceGetPolicyParams{
		timeout: timeout,
	}
}

// NewHeadscaleServiceGetPolicyParamsWithContext creates a new HeadscaleServiceGetPolicyParams object
// with the ability to set a context for a request.
func NewHeadscaleServiceGetPolicyParamsWithContext(ctx context.Context) *HeadscaleServiceGetPolicyParams {
	return &HeadscaleServiceGetPolicyParams{
		Context: ctx,
	}
}

// NewHeadscaleServiceGetPolicyParamsWithHTTPClient creates a new HeadscaleServiceGetPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeadscaleServiceGetPolicyParamsWithHTTPClient(client *http.Client) *HeadscaleServiceGetPolicyParams {
	return &HeadscaleServiceGetPolicyParams{
		HTTPClient: client,
	}
}

/*
HeadscaleServiceGetPolicyParams contains all the parameters to send to the API endpoint

	for the headscale service get policy operation.

	Typically these are written to a http.Request.
*/
type HeadscaleServiceGetPolicyParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the headscale service get policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadscaleServiceGetPolicyParams) WithDefaults() *HeadscaleServiceGetPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the headscale service get policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadscaleServiceGetPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the headscale service get policy params
func (o *HeadscaleServiceGetPolicyParams) WithTimeout(timeout time.Duration) *HeadscaleServiceGetPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the headscale service get policy params
func (o *HeadscaleServiceGetPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the headscale service get policy params
func (o *HeadscaleServiceGetPolicyParams) WithContext(ctx context.Context) *HeadscaleServiceGetPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the headscale service get policy params
func (o *HeadscaleServiceGetPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the headscale service get policy params
func (o *HeadscaleServiceGetPolicyParams) WithHTTPClient(client *http.Client) *HeadscaleServiceGetPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the headscale service get policy params
func (o *HeadscaleServiceGetPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HeadscaleServiceGetPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
