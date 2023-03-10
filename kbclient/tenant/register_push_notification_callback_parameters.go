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

// NewRegisterPushNotificationCallbackParams creates a new RegisterPushNotificationCallbackParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRegisterPushNotificationCallbackParams() *RegisterPushNotificationCallbackParams {
	return &RegisterPushNotificationCallbackParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterPushNotificationCallbackParamsWithTimeout creates a new RegisterPushNotificationCallbackParams object
// with the ability to set a timeout on a request.
func NewRegisterPushNotificationCallbackParamsWithTimeout(timeout time.Duration) *RegisterPushNotificationCallbackParams {
	return &RegisterPushNotificationCallbackParams{
		timeout: timeout,
	}
}

// NewRegisterPushNotificationCallbackParamsWithContext creates a new RegisterPushNotificationCallbackParams object
// with the ability to set a context for a request.
func NewRegisterPushNotificationCallbackParamsWithContext(ctx context.Context) *RegisterPushNotificationCallbackParams {
	return &RegisterPushNotificationCallbackParams{
		Context: ctx,
	}
}

// NewRegisterPushNotificationCallbackParamsWithHTTPClient creates a new RegisterPushNotificationCallbackParams object
// with the ability to set a custom HTTPClient for a request.
func NewRegisterPushNotificationCallbackParamsWithHTTPClient(client *http.Client) *RegisterPushNotificationCallbackParams {
	return &RegisterPushNotificationCallbackParams{
		HTTPClient: client,
	}
}

/*
RegisterPushNotificationCallbackParams contains all the parameters to send to the API endpoint

	for the register push notification callback operation.

	Typically these are written to a http.Request.
*/
type RegisterPushNotificationCallbackParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Cb.
	Cb *string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the register push notification callback params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterPushNotificationCallbackParams) WithDefaults() *RegisterPushNotificationCallbackParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the register push notification callback params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterPushNotificationCallbackParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) WithTimeout(timeout time.Duration) *RegisterPushNotificationCallbackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) WithContext(ctx context.Context) *RegisterPushNotificationCallbackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) WithHTTPClient(client *http.Client) *RegisterPushNotificationCallbackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) WithXKillbillComment(xKillbillComment *string) *RegisterPushNotificationCallbackParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *RegisterPushNotificationCallbackParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) WithXKillbillReason(xKillbillReason *string) *RegisterPushNotificationCallbackParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithCb adds the cb to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) WithCb(cb *string) *RegisterPushNotificationCallbackParams {
	o.SetCb(cb)
	return o
}

// SetCb adds the cb to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) SetCb(cb *string) {
	o.Cb = cb
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterPushNotificationCallbackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Cb != nil {

		// query param cb
		var qrCb string

		if o.Cb != nil {
			qrCb = *o.Cb
		}
		qCb := qrCb
		if qCb != "" {

			if err := r.SetQueryParam("cb", qCb); err != nil {
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
