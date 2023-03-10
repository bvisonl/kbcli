// Code generated by go-swagger; DO NOT EDIT.

package bundle

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
	"github.com/go-openapi/swag"
)

// NewResumeBundleParams creates a new ResumeBundleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResumeBundleParams() *ResumeBundleParams {
	return &ResumeBundleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResumeBundleParamsWithTimeout creates a new ResumeBundleParams object
// with the ability to set a timeout on a request.
func NewResumeBundleParamsWithTimeout(timeout time.Duration) *ResumeBundleParams {
	return &ResumeBundleParams{
		timeout: timeout,
	}
}

// NewResumeBundleParamsWithContext creates a new ResumeBundleParams object
// with the ability to set a context for a request.
func NewResumeBundleParamsWithContext(ctx context.Context) *ResumeBundleParams {
	return &ResumeBundleParams{
		Context: ctx,
	}
}

// NewResumeBundleParamsWithHTTPClient creates a new ResumeBundleParams object
// with the ability to set a custom HTTPClient for a request.
func NewResumeBundleParamsWithHTTPClient(client *http.Client) *ResumeBundleParams {
	return &ResumeBundleParams{
		HTTPClient: client,
	}
}

/*
ResumeBundleParams contains all the parameters to send to the API endpoint

	for the resume bundle operation.

	Typically these are written to a http.Request.
*/
type ResumeBundleParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// BundleID.
	//
	// Format: uuid
	BundleID strfmt.UUID

	// PluginProperty.
	PluginProperty []string

	// RequestedDate.
	//
	// Format: date
	RequestedDate *strfmt.Date

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the resume bundle params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResumeBundleParams) WithDefaults() *ResumeBundleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resume bundle params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResumeBundleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the resume bundle params
func (o *ResumeBundleParams) WithTimeout(timeout time.Duration) *ResumeBundleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resume bundle params
func (o *ResumeBundleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resume bundle params
func (o *ResumeBundleParams) WithContext(ctx context.Context) *ResumeBundleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resume bundle params
func (o *ResumeBundleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resume bundle params
func (o *ResumeBundleParams) WithHTTPClient(client *http.Client) *ResumeBundleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resume bundle params
func (o *ResumeBundleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the resume bundle params
func (o *ResumeBundleParams) WithXKillbillComment(xKillbillComment *string) *ResumeBundleParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the resume bundle params
func (o *ResumeBundleParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the resume bundle params
func (o *ResumeBundleParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *ResumeBundleParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the resume bundle params
func (o *ResumeBundleParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the resume bundle params
func (o *ResumeBundleParams) WithXKillbillReason(xKillbillReason *string) *ResumeBundleParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the resume bundle params
func (o *ResumeBundleParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBundleID adds the bundleID to the resume bundle params
func (o *ResumeBundleParams) WithBundleID(bundleID strfmt.UUID) *ResumeBundleParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the resume bundle params
func (o *ResumeBundleParams) SetBundleID(bundleID strfmt.UUID) {
	o.BundleID = bundleID
}

// WithPluginProperty adds the pluginProperty to the resume bundle params
func (o *ResumeBundleParams) WithPluginProperty(pluginProperty []string) *ResumeBundleParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the resume bundle params
func (o *ResumeBundleParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithRequestedDate adds the requestedDate to the resume bundle params
func (o *ResumeBundleParams) WithRequestedDate(requestedDate *strfmt.Date) *ResumeBundleParams {
	o.SetRequestedDate(requestedDate)
	return o
}

// SetRequestedDate adds the requestedDate to the resume bundle params
func (o *ResumeBundleParams) SetRequestedDate(requestedDate *strfmt.Date) {
	o.RequestedDate = requestedDate
}

// WriteToRequest writes these params to a swagger request
func (o *ResumeBundleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param bundleId
	if err := r.SetPathParam("bundleId", o.BundleID.String()); err != nil {
		return err
	}

	if o.PluginProperty != nil {

		// binding items for pluginProperty
		joinedPluginProperty := o.bindParamPluginProperty(reg)

		// query array param pluginProperty
		if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
			return err
		}
	}

	if o.RequestedDate != nil {

		// query param requestedDate
		var qrRequestedDate strfmt.Date

		if o.RequestedDate != nil {
			qrRequestedDate = *o.RequestedDate
		}
		qRequestedDate := qrRequestedDate.String()
		if qRequestedDate != "" {

			if err := r.SetQueryParam("requestedDate", qRequestedDate); err != nil {
				return err
			}
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

// bindParamResumeBundle binds the parameter pluginProperty
func (o *ResumeBundleParams) bindParamPluginProperty(formats strfmt.Registry) []string {
	pluginPropertyIR := o.PluginProperty

	var pluginPropertyIC []string
	for _, pluginPropertyIIR := range pluginPropertyIR { // explode []string

		pluginPropertyIIV := pluginPropertyIIR // string as string
		pluginPropertyIC = append(pluginPropertyIC, pluginPropertyIIV)
	}

	// items.CollectionFormat: "multi"
	pluginPropertyIS := swag.JoinByFormat(pluginPropertyIC, "multi")

	return pluginPropertyIS
}
