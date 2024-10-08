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

// NewHeadscaleServiceListNodesParams creates a new HeadscaleServiceListNodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeadscaleServiceListNodesParams() *HeadscaleServiceListNodesParams {
	return &HeadscaleServiceListNodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeadscaleServiceListNodesParamsWithTimeout creates a new HeadscaleServiceListNodesParams object
// with the ability to set a timeout on a request.
func NewHeadscaleServiceListNodesParamsWithTimeout(timeout time.Duration) *HeadscaleServiceListNodesParams {
	return &HeadscaleServiceListNodesParams{
		timeout: timeout,
	}
}

// NewHeadscaleServiceListNodesParamsWithContext creates a new HeadscaleServiceListNodesParams object
// with the ability to set a context for a request.
func NewHeadscaleServiceListNodesParamsWithContext(ctx context.Context) *HeadscaleServiceListNodesParams {
	return &HeadscaleServiceListNodesParams{
		Context: ctx,
	}
}

// NewHeadscaleServiceListNodesParamsWithHTTPClient creates a new HeadscaleServiceListNodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeadscaleServiceListNodesParamsWithHTTPClient(client *http.Client) *HeadscaleServiceListNodesParams {
	return &HeadscaleServiceListNodesParams{
		HTTPClient: client,
	}
}

/*
HeadscaleServiceListNodesParams contains all the parameters to send to the API endpoint

	for the headscale service list nodes operation.

	Typically these are written to a http.Request.
*/
type HeadscaleServiceListNodesParams struct {

	// User.
	User *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the headscale service list nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadscaleServiceListNodesParams) WithDefaults() *HeadscaleServiceListNodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the headscale service list nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadscaleServiceListNodesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the headscale service list nodes params
func (o *HeadscaleServiceListNodesParams) WithTimeout(timeout time.Duration) *HeadscaleServiceListNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the headscale service list nodes params
func (o *HeadscaleServiceListNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the headscale service list nodes params
func (o *HeadscaleServiceListNodesParams) WithContext(ctx context.Context) *HeadscaleServiceListNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the headscale service list nodes params
func (o *HeadscaleServiceListNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the headscale service list nodes params
func (o *HeadscaleServiceListNodesParams) WithHTTPClient(client *http.Client) *HeadscaleServiceListNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the headscale service list nodes params
func (o *HeadscaleServiceListNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUser adds the user to the headscale service list nodes params
func (o *HeadscaleServiceListNodesParams) WithUser(user *string) *HeadscaleServiceListNodesParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the headscale service list nodes params
func (o *HeadscaleServiceListNodesParams) SetUser(user *string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *HeadscaleServiceListNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.User != nil {

		// query param user
		var qrUser string

		if o.User != nil {
			qrUser = *o.User
		}
		qUser := qrUser
		if qUser != "" {

			if err := r.SetQueryParam("user", qUser); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
