// Code generated by go-swagger; DO NOT EDIT.

package payment_gateway

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

// NewProcessNotificationParams creates a new ProcessNotificationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProcessNotificationParams() *ProcessNotificationParams {
	return &ProcessNotificationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProcessNotificationParamsWithTimeout creates a new ProcessNotificationParams object
// with the ability to set a timeout on a request.
func NewProcessNotificationParamsWithTimeout(timeout time.Duration) *ProcessNotificationParams {
	return &ProcessNotificationParams{
		timeout: timeout,
	}
}

// NewProcessNotificationParamsWithContext creates a new ProcessNotificationParams object
// with the ability to set a context for a request.
func NewProcessNotificationParamsWithContext(ctx context.Context) *ProcessNotificationParams {
	return &ProcessNotificationParams{
		Context: ctx,
	}
}

// NewProcessNotificationParamsWithHTTPClient creates a new ProcessNotificationParams object
// with the ability to set a custom HTTPClient for a request.
func NewProcessNotificationParamsWithHTTPClient(client *http.Client) *ProcessNotificationParams {
	return &ProcessNotificationParams{
		HTTPClient: client,
	}
}

/*
ProcessNotificationParams contains all the parameters to send to the API endpoint

	for the process notification operation.

	Typically these are written to a http.Request.
*/
type ProcessNotificationParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body string

	// ControlPluginName.
	ControlPluginName []string

	// PluginName.
	PluginName string

	// PluginProperty.
	PluginProperty []string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the process notification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProcessNotificationParams) WithDefaults() *ProcessNotificationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the process notification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProcessNotificationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the process notification params
func (o *ProcessNotificationParams) WithTimeout(timeout time.Duration) *ProcessNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the process notification params
func (o *ProcessNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the process notification params
func (o *ProcessNotificationParams) WithContext(ctx context.Context) *ProcessNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the process notification params
func (o *ProcessNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the process notification params
func (o *ProcessNotificationParams) WithHTTPClient(client *http.Client) *ProcessNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the process notification params
func (o *ProcessNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the process notification params
func (o *ProcessNotificationParams) WithXKillbillComment(xKillbillComment *string) *ProcessNotificationParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the process notification params
func (o *ProcessNotificationParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the process notification params
func (o *ProcessNotificationParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *ProcessNotificationParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the process notification params
func (o *ProcessNotificationParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the process notification params
func (o *ProcessNotificationParams) WithXKillbillReason(xKillbillReason *string) *ProcessNotificationParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the process notification params
func (o *ProcessNotificationParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the process notification params
func (o *ProcessNotificationParams) WithBody(body string) *ProcessNotificationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the process notification params
func (o *ProcessNotificationParams) SetBody(body string) {
	o.Body = body
}

// WithControlPluginName adds the controlPluginName to the process notification params
func (o *ProcessNotificationParams) WithControlPluginName(controlPluginName []string) *ProcessNotificationParams {
	o.SetControlPluginName(controlPluginName)
	return o
}

// SetControlPluginName adds the controlPluginName to the process notification params
func (o *ProcessNotificationParams) SetControlPluginName(controlPluginName []string) {
	o.ControlPluginName = controlPluginName
}

// WithPluginName adds the pluginName to the process notification params
func (o *ProcessNotificationParams) WithPluginName(pluginName string) *ProcessNotificationParams {
	o.SetPluginName(pluginName)
	return o
}

// SetPluginName adds the pluginName to the process notification params
func (o *ProcessNotificationParams) SetPluginName(pluginName string) {
	o.PluginName = pluginName
}

// WithPluginProperty adds the pluginProperty to the process notification params
func (o *ProcessNotificationParams) WithPluginProperty(pluginProperty []string) *ProcessNotificationParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the process notification params
func (o *ProcessNotificationParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *ProcessNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ControlPluginName != nil {

		// binding items for controlPluginName
		joinedControlPluginName := o.bindParamControlPluginName(reg)

		// query array param controlPluginName
		if err := r.SetQueryParam("controlPluginName", joinedControlPluginName...); err != nil {
			return err
		}
	}

	// path param pluginName
	if err := r.SetPathParam("pluginName", o.PluginName); err != nil {
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

// bindParamProcessNotification binds the parameter controlPluginName
func (o *ProcessNotificationParams) bindParamControlPluginName(formats strfmt.Registry) []string {
	controlPluginNameIR := o.ControlPluginName

	var controlPluginNameIC []string
	for _, controlPluginNameIIR := range controlPluginNameIR { // explode []string

		controlPluginNameIIV := controlPluginNameIIR // string as string
		controlPluginNameIC = append(controlPluginNameIC, controlPluginNameIIV)
	}

	// items.CollectionFormat: "multi"
	controlPluginNameIS := swag.JoinByFormat(controlPluginNameIC, "multi")

	return controlPluginNameIS
}

// bindParamProcessNotification binds the parameter pluginProperty
func (o *ProcessNotificationParams) bindParamPluginProperty(formats strfmt.Registry) []string {
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
