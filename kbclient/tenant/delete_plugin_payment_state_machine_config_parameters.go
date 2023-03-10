// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// NewDeletePluginPaymentStateMachineConfigParams creates a new DeletePluginPaymentStateMachineConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePluginPaymentStateMachineConfigParams() *DeletePluginPaymentStateMachineConfigParams {
	return &DeletePluginPaymentStateMachineConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePluginPaymentStateMachineConfigParamsWithTimeout creates a new DeletePluginPaymentStateMachineConfigParams object
// with the ability to set a timeout on a request.
func NewDeletePluginPaymentStateMachineConfigParamsWithTimeout(timeout time.Duration) *DeletePluginPaymentStateMachineConfigParams {
	return &DeletePluginPaymentStateMachineConfigParams{
		timeout: timeout,
	}
}

// NewDeletePluginPaymentStateMachineConfigParamsWithContext creates a new DeletePluginPaymentStateMachineConfigParams object
// with the ability to set a context for a request.
func NewDeletePluginPaymentStateMachineConfigParamsWithContext(ctx context.Context) *DeletePluginPaymentStateMachineConfigParams {
	return &DeletePluginPaymentStateMachineConfigParams{
		Context: ctx,
	}
}

// NewDeletePluginPaymentStateMachineConfigParamsWithHTTPClient creates a new DeletePluginPaymentStateMachineConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePluginPaymentStateMachineConfigParamsWithHTTPClient(client *http.Client) *DeletePluginPaymentStateMachineConfigParams {
	return &DeletePluginPaymentStateMachineConfigParams{
		HTTPClient: client,
	}
}

/*
DeletePluginPaymentStateMachineConfigParams contains all the parameters to send to the API endpoint

	for the delete plugin payment state machine config operation.

	Typically these are written to a http.Request.
*/
type DeletePluginPaymentStateMachineConfigParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// PluginName.
	PluginName string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the delete plugin payment state machine config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePluginPaymentStateMachineConfigParams) WithDefaults() *DeletePluginPaymentStateMachineConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete plugin payment state machine config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePluginPaymentStateMachineConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) WithTimeout(timeout time.Duration) *DeletePluginPaymentStateMachineConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) WithContext(ctx context.Context) *DeletePluginPaymentStateMachineConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) WithHTTPClient(client *http.Client) *DeletePluginPaymentStateMachineConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) WithXKillbillComment(xKillbillComment *string) *DeletePluginPaymentStateMachineConfigParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeletePluginPaymentStateMachineConfigParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) WithXKillbillReason(xKillbillReason *string) *DeletePluginPaymentStateMachineConfigParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithPluginName adds the pluginName to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) WithPluginName(pluginName string) *DeletePluginPaymentStateMachineConfigParams {
	o.SetPluginName(pluginName)
	return o
}

// SetPluginName adds the pluginName to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) SetPluginName(pluginName string) {
	o.PluginName = pluginName
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePluginPaymentStateMachineConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param pluginName
	if err := r.SetPathParam("pluginName", o.PluginName); err != nil {
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
