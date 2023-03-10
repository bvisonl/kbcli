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
)

// NewValidateCatalogXMLParams creates a new ValidateCatalogXMLParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateCatalogXMLParams() *ValidateCatalogXMLParams {
	return &ValidateCatalogXMLParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateCatalogXMLParamsWithTimeout creates a new ValidateCatalogXMLParams object
// with the ability to set a timeout on a request.
func NewValidateCatalogXMLParamsWithTimeout(timeout time.Duration) *ValidateCatalogXMLParams {
	return &ValidateCatalogXMLParams{
		timeout: timeout,
	}
}

// NewValidateCatalogXMLParamsWithContext creates a new ValidateCatalogXMLParams object
// with the ability to set a context for a request.
func NewValidateCatalogXMLParamsWithContext(ctx context.Context) *ValidateCatalogXMLParams {
	return &ValidateCatalogXMLParams{
		Context: ctx,
	}
}

// NewValidateCatalogXMLParamsWithHTTPClient creates a new ValidateCatalogXMLParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateCatalogXMLParamsWithHTTPClient(client *http.Client) *ValidateCatalogXMLParams {
	return &ValidateCatalogXMLParams{
		HTTPClient: client,
	}
}

/*
ValidateCatalogXMLParams contains all the parameters to send to the API endpoint

	for the validate catalog Xml operation.

	Typically these are written to a http.Request.
*/
type ValidateCatalogXMLParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the validate catalog Xml params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateCatalogXMLParams) WithDefaults() *ValidateCatalogXMLParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate catalog Xml params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateCatalogXMLParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate catalog Xml params
func (o *ValidateCatalogXMLParams) WithTimeout(timeout time.Duration) *ValidateCatalogXMLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate catalog Xml params
func (o *ValidateCatalogXMLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate catalog Xml params
func (o *ValidateCatalogXMLParams) WithContext(ctx context.Context) *ValidateCatalogXMLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate catalog Xml params
func (o *ValidateCatalogXMLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate catalog Xml params
func (o *ValidateCatalogXMLParams) WithHTTPClient(client *http.Client) *ValidateCatalogXMLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate catalog Xml params
func (o *ValidateCatalogXMLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the validate catalog Xml params
func (o *ValidateCatalogXMLParams) WithXKillbillComment(xKillbillComment *string) *ValidateCatalogXMLParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the validate catalog Xml params
func (o *ValidateCatalogXMLParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the validate catalog Xml params
func (o *ValidateCatalogXMLParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *ValidateCatalogXMLParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the validate catalog Xml params
func (o *ValidateCatalogXMLParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the validate catalog Xml params
func (o *ValidateCatalogXMLParams) WithXKillbillReason(xKillbillReason *string) *ValidateCatalogXMLParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the validate catalog Xml params
func (o *ValidateCatalogXMLParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the validate catalog Xml params
func (o *ValidateCatalogXMLParams) WithBody(body string) *ValidateCatalogXMLParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate catalog Xml params
func (o *ValidateCatalogXMLParams) SetBody(body string) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateCatalogXMLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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
