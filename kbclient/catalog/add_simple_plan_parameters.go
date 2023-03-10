// Code generated by go-swagger; DO NOT EDIT.

package catalog

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

	"github.com/killbill/kbcli/v3/kbmodel"
)

// NewAddSimplePlanParams creates a new AddSimplePlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddSimplePlanParams() *AddSimplePlanParams {
	return &AddSimplePlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddSimplePlanParamsWithTimeout creates a new AddSimplePlanParams object
// with the ability to set a timeout on a request.
func NewAddSimplePlanParamsWithTimeout(timeout time.Duration) *AddSimplePlanParams {
	return &AddSimplePlanParams{
		timeout: timeout,
	}
}

// NewAddSimplePlanParamsWithContext creates a new AddSimplePlanParams object
// with the ability to set a context for a request.
func NewAddSimplePlanParamsWithContext(ctx context.Context) *AddSimplePlanParams {
	return &AddSimplePlanParams{
		Context: ctx,
	}
}

// NewAddSimplePlanParamsWithHTTPClient creates a new AddSimplePlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddSimplePlanParamsWithHTTPClient(client *http.Client) *AddSimplePlanParams {
	return &AddSimplePlanParams{
		HTTPClient: client,
	}
}

/*
AddSimplePlanParams contains all the parameters to send to the API endpoint

	for the add simple plan operation.

	Typically these are written to a http.Request.
*/
type AddSimplePlanParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body *kbmodel.SimplePlan

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the add simple plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddSimplePlanParams) WithDefaults() *AddSimplePlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add simple plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddSimplePlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add simple plan params
func (o *AddSimplePlanParams) WithTimeout(timeout time.Duration) *AddSimplePlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add simple plan params
func (o *AddSimplePlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add simple plan params
func (o *AddSimplePlanParams) WithContext(ctx context.Context) *AddSimplePlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add simple plan params
func (o *AddSimplePlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add simple plan params
func (o *AddSimplePlanParams) WithHTTPClient(client *http.Client) *AddSimplePlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add simple plan params
func (o *AddSimplePlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the add simple plan params
func (o *AddSimplePlanParams) WithXKillbillComment(xKillbillComment *string) *AddSimplePlanParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the add simple plan params
func (o *AddSimplePlanParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the add simple plan params
func (o *AddSimplePlanParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *AddSimplePlanParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the add simple plan params
func (o *AddSimplePlanParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the add simple plan params
func (o *AddSimplePlanParams) WithXKillbillReason(xKillbillReason *string) *AddSimplePlanParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the add simple plan params
func (o *AddSimplePlanParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the add simple plan params
func (o *AddSimplePlanParams) WithBody(body *kbmodel.SimplePlan) *AddSimplePlanParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add simple plan params
func (o *AddSimplePlanParams) SetBody(body *kbmodel.SimplePlan) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddSimplePlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XKillbillComment != nil {

		// header param X-Killbill-Comment
		if err := r.SetHeaderParam("X-Killbill-Comment", *o.XKillbillComment); err != nil {
			return err
		}
	}

	// header param X-Killbill-CreatedBy
	if err := r.SetHeaderParam("X-Killbill-CreatedBy", o.XKillbillCreatedBy); err != nil {
		return err
	}

	if o.XKillbillReason != nil {

		// header param X-Killbill-Reason
		if err := r.SetHeaderParam("X-Killbill-Reason", *o.XKillbillReason); err != nil {
			return err
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
