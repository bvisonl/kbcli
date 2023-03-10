// Code generated by go-swagger; DO NOT EDIT.

package payment

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

// NewCreatePaymentTagsParams creates a new CreatePaymentTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePaymentTagsParams() *CreatePaymentTagsParams {
	return &CreatePaymentTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePaymentTagsParamsWithTimeout creates a new CreatePaymentTagsParams object
// with the ability to set a timeout on a request.
func NewCreatePaymentTagsParamsWithTimeout(timeout time.Duration) *CreatePaymentTagsParams {
	return &CreatePaymentTagsParams{
		timeout: timeout,
	}
}

// NewCreatePaymentTagsParamsWithContext creates a new CreatePaymentTagsParams object
// with the ability to set a context for a request.
func NewCreatePaymentTagsParamsWithContext(ctx context.Context) *CreatePaymentTagsParams {
	return &CreatePaymentTagsParams{
		Context: ctx,
	}
}

// NewCreatePaymentTagsParamsWithHTTPClient creates a new CreatePaymentTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePaymentTagsParamsWithHTTPClient(client *http.Client) *CreatePaymentTagsParams {
	return &CreatePaymentTagsParams{
		HTTPClient: client,
	}
}

/*
CreatePaymentTagsParams contains all the parameters to send to the API endpoint

	for the create payment tags operation.

	Typically these are written to a http.Request.
*/
type CreatePaymentTagsParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body []strfmt.UUID

	// PaymentID.
	//
	// Format: uuid
	PaymentID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the create payment tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePaymentTagsParams) WithDefaults() *CreatePaymentTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create payment tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePaymentTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create payment tags params
func (o *CreatePaymentTagsParams) WithTimeout(timeout time.Duration) *CreatePaymentTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create payment tags params
func (o *CreatePaymentTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create payment tags params
func (o *CreatePaymentTagsParams) WithContext(ctx context.Context) *CreatePaymentTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create payment tags params
func (o *CreatePaymentTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create payment tags params
func (o *CreatePaymentTagsParams) WithHTTPClient(client *http.Client) *CreatePaymentTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create payment tags params
func (o *CreatePaymentTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the create payment tags params
func (o *CreatePaymentTagsParams) WithXKillbillComment(xKillbillComment *string) *CreatePaymentTagsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create payment tags params
func (o *CreatePaymentTagsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create payment tags params
func (o *CreatePaymentTagsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreatePaymentTagsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create payment tags params
func (o *CreatePaymentTagsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create payment tags params
func (o *CreatePaymentTagsParams) WithXKillbillReason(xKillbillReason *string) *CreatePaymentTagsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create payment tags params
func (o *CreatePaymentTagsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the create payment tags params
func (o *CreatePaymentTagsParams) WithBody(body []strfmt.UUID) *CreatePaymentTagsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create payment tags params
func (o *CreatePaymentTagsParams) SetBody(body []strfmt.UUID) {
	o.Body = body
}

// WithPaymentID adds the paymentID to the create payment tags params
func (o *CreatePaymentTagsParams) WithPaymentID(paymentID strfmt.UUID) *CreatePaymentTagsParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the create payment tags params
func (o *CreatePaymentTagsParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePaymentTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID.String()); err != nil {
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
