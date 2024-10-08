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

// NewHeadscaleServiceDeleteAPIKeyParams creates a new HeadscaleServiceDeleteAPIKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeadscaleServiceDeleteAPIKeyParams() *HeadscaleServiceDeleteAPIKeyParams {
	return &HeadscaleServiceDeleteAPIKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeadscaleServiceDeleteAPIKeyParamsWithTimeout creates a new HeadscaleServiceDeleteAPIKeyParams object
// with the ability to set a timeout on a request.
func NewHeadscaleServiceDeleteAPIKeyParamsWithTimeout(timeout time.Duration) *HeadscaleServiceDeleteAPIKeyParams {
	return &HeadscaleServiceDeleteAPIKeyParams{
		timeout: timeout,
	}
}

// NewHeadscaleServiceDeleteAPIKeyParamsWithContext creates a new HeadscaleServiceDeleteAPIKeyParams object
// with the ability to set a context for a request.
func NewHeadscaleServiceDeleteAPIKeyParamsWithContext(ctx context.Context) *HeadscaleServiceDeleteAPIKeyParams {
	return &HeadscaleServiceDeleteAPIKeyParams{
		Context: ctx,
	}
}

// NewHeadscaleServiceDeleteAPIKeyParamsWithHTTPClient creates a new HeadscaleServiceDeleteAPIKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeadscaleServiceDeleteAPIKeyParamsWithHTTPClient(client *http.Client) *HeadscaleServiceDeleteAPIKeyParams {
	return &HeadscaleServiceDeleteAPIKeyParams{
		HTTPClient: client,
	}
}

/*
HeadscaleServiceDeleteAPIKeyParams contains all the parameters to send to the API endpoint

	for the headscale service delete Api key operation.

	Typically these are written to a http.Request.
*/
type HeadscaleServiceDeleteAPIKeyParams struct {

	// Prefix.
	Prefix string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the headscale service delete Api key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadscaleServiceDeleteAPIKeyParams) WithDefaults() *HeadscaleServiceDeleteAPIKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the headscale service delete Api key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadscaleServiceDeleteAPIKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the headscale service delete Api key params
func (o *HeadscaleServiceDeleteAPIKeyParams) WithTimeout(timeout time.Duration) *HeadscaleServiceDeleteAPIKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the headscale service delete Api key params
func (o *HeadscaleServiceDeleteAPIKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the headscale service delete Api key params
func (o *HeadscaleServiceDeleteAPIKeyParams) WithContext(ctx context.Context) *HeadscaleServiceDeleteAPIKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the headscale service delete Api key params
func (o *HeadscaleServiceDeleteAPIKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the headscale service delete Api key params
func (o *HeadscaleServiceDeleteAPIKeyParams) WithHTTPClient(client *http.Client) *HeadscaleServiceDeleteAPIKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the headscale service delete Api key params
func (o *HeadscaleServiceDeleteAPIKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPrefix adds the prefix to the headscale service delete Api key params
func (o *HeadscaleServiceDeleteAPIKeyParams) WithPrefix(prefix string) *HeadscaleServiceDeleteAPIKeyParams {
	o.SetPrefix(prefix)
	return o
}

// SetPrefix adds the prefix to the headscale service delete Api key params
func (o *HeadscaleServiceDeleteAPIKeyParams) SetPrefix(prefix string) {
	o.Prefix = prefix
}

// WriteToRequest writes these params to a swagger request
func (o *HeadscaleServiceDeleteAPIKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param prefix
	if err := r.SetPathParam("prefix", o.Prefix); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
