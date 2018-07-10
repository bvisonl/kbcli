// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// NewAddEmailParams creates a new AddEmailParams object
// with the default values initialized.
func NewAddEmailParams() *AddEmailParams {
	var ()
	return &AddEmailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddEmailParamsWithTimeout creates a new AddEmailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddEmailParamsWithTimeout(timeout time.Duration) *AddEmailParams {
	var ()
	return &AddEmailParams{

		timeout: timeout,
	}
}

// NewAddEmailParamsWithContext creates a new AddEmailParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddEmailParamsWithContext(ctx context.Context) *AddEmailParams {
	var ()
	return &AddEmailParams{

		Context: ctx,
	}
}

// NewAddEmailParamsWithHTTPClient creates a new AddEmailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddEmailParamsWithHTTPClient(client *http.Client) *AddEmailParams {
	var ()
	return &AddEmailParams{
		HTTPClient: client,
	}
}

/*AddEmailParams contains all the parameters to send to the API endpoint
for the add email operation typically these are written to a http.Request
*/
type AddEmailParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*AccountID*/
	AccountID strfmt.UUID
	/*Body*/
	Body *kbmodel.AccountEmail

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the add email params
func (o *AddEmailParams) WithTimeout(timeout time.Duration) *AddEmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add email params
func (o *AddEmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add email params
func (o *AddEmailParams) WithContext(ctx context.Context) *AddEmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add email params
func (o *AddEmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add email params
func (o *AddEmailParams) WithHTTPClient(client *http.Client) *AddEmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add email params
func (o *AddEmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the add email params
func (o *AddEmailParams) WithXKillbillAPIKey(xKillbillAPIKey string) *AddEmailParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the add email params
func (o *AddEmailParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the add email params
func (o *AddEmailParams) WithXKillbillAPISecret(xKillbillAPISecret string) *AddEmailParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the add email params
func (o *AddEmailParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the add email params
func (o *AddEmailParams) WithXKillbillComment(xKillbillComment *string) *AddEmailParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the add email params
func (o *AddEmailParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the add email params
func (o *AddEmailParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *AddEmailParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the add email params
func (o *AddEmailParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the add email params
func (o *AddEmailParams) WithXKillbillReason(xKillbillReason *string) *AddEmailParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the add email params
func (o *AddEmailParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the add email params
func (o *AddEmailParams) WithAccountID(accountID strfmt.UUID) *AddEmailParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the add email params
func (o *AddEmailParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithBody adds the body to the add email params
func (o *AddEmailParams) WithBody(body *kbmodel.AccountEmail) *AddEmailParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add email params
func (o *AddEmailParams) SetBody(body *kbmodel.AccountEmail) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Killbill-ApiKey
	if err := r.SetHeaderParam("X-Killbill-ApiKey", o.XKillbillAPIKey); err != nil {
		return err
	}

	// header param X-Killbill-ApiSecret
	if err := r.SetHeaderParam("X-Killbill-ApiSecret", o.XKillbillAPISecret); err != nil {
		return err
	}

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

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
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