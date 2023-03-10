// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewPutInRotationParams creates a new PutInRotationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutInRotationParams() *PutInRotationParams {
	return &PutInRotationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutInRotationParamsWithTimeout creates a new PutInRotationParams object
// with the ability to set a timeout on a request.
func NewPutInRotationParamsWithTimeout(timeout time.Duration) *PutInRotationParams {
	return &PutInRotationParams{
		timeout: timeout,
	}
}

// NewPutInRotationParamsWithContext creates a new PutInRotationParams object
// with the ability to set a context for a request.
func NewPutInRotationParamsWithContext(ctx context.Context) *PutInRotationParams {
	return &PutInRotationParams{
		Context: ctx,
	}
}

// NewPutInRotationParamsWithHTTPClient creates a new PutInRotationParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutInRotationParamsWithHTTPClient(client *http.Client) *PutInRotationParams {
	return &PutInRotationParams{
		HTTPClient: client,
	}
}

/*
PutInRotationParams contains all the parameters to send to the API endpoint

	for the put in rotation operation.

	Typically these are written to a http.Request.
*/
type PutInRotationParams struct {
	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the put in rotation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutInRotationParams) WithDefaults() *PutInRotationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put in rotation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutInRotationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put in rotation params
func (o *PutInRotationParams) WithTimeout(timeout time.Duration) *PutInRotationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put in rotation params
func (o *PutInRotationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put in rotation params
func (o *PutInRotationParams) WithContext(ctx context.Context) *PutInRotationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put in rotation params
func (o *PutInRotationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put in rotation params
func (o *PutInRotationParams) WithHTTPClient(client *http.Client) *PutInRotationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put in rotation params
func (o *PutInRotationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PutInRotationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param WithProfilingInfo
	if o.WithProfilingInfo != nil && len(*o.WithProfilingInfo) > 0 {
		if err := r.SetHeaderParam("X-Killbill-Profiling-Req", *o.WithProfilingInfo); err != nil {
			return err
		}
	}

	// header param withStackTrace
	if o.WithStackTrace != nil && *o.WithStackTrace {
		if err := r.SetQueryParam("withStackTrace", "true"); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
