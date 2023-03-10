// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// NewGetInvoiceByNumberParams creates a new GetInvoiceByNumberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInvoiceByNumberParams() *GetInvoiceByNumberParams {
	return &GetInvoiceByNumberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvoiceByNumberParamsWithTimeout creates a new GetInvoiceByNumberParams object
// with the ability to set a timeout on a request.
func NewGetInvoiceByNumberParamsWithTimeout(timeout time.Duration) *GetInvoiceByNumberParams {
	return &GetInvoiceByNumberParams{
		timeout: timeout,
	}
}

// NewGetInvoiceByNumberParamsWithContext creates a new GetInvoiceByNumberParams object
// with the ability to set a context for a request.
func NewGetInvoiceByNumberParamsWithContext(ctx context.Context) *GetInvoiceByNumberParams {
	return &GetInvoiceByNumberParams{
		Context: ctx,
	}
}

// NewGetInvoiceByNumberParamsWithHTTPClient creates a new GetInvoiceByNumberParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInvoiceByNumberParamsWithHTTPClient(client *http.Client) *GetInvoiceByNumberParams {
	return &GetInvoiceByNumberParams{
		HTTPClient: client,
	}
}

/*
GetInvoiceByNumberParams contains all the parameters to send to the API endpoint

	for the get invoice by number operation.

	Typically these are written to a http.Request.
*/
type GetInvoiceByNumberParams struct {

	// Audit.
	//
	// Default: "NONE"
	Audit *string

	// InvoiceNumber.
	//
	// Format: int32
	InvoiceNumber int32

	// WithChildrenItems.
	WithChildrenItems *bool

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the get invoice by number params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInvoiceByNumberParams) WithDefaults() *GetInvoiceByNumberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get invoice by number params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInvoiceByNumberParams) SetDefaults() {
	var (
		auditDefault = string("NONE")

		withChildrenItemsDefault = bool(false)
	)

	val := GetInvoiceByNumberParams{
		Audit:             &auditDefault,
		WithChildrenItems: &withChildrenItemsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get invoice by number params
func (o *GetInvoiceByNumberParams) WithTimeout(timeout time.Duration) *GetInvoiceByNumberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invoice by number params
func (o *GetInvoiceByNumberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invoice by number params
func (o *GetInvoiceByNumberParams) WithContext(ctx context.Context) *GetInvoiceByNumberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invoice by number params
func (o *GetInvoiceByNumberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invoice by number params
func (o *GetInvoiceByNumberParams) WithHTTPClient(client *http.Client) *GetInvoiceByNumberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invoice by number params
func (o *GetInvoiceByNumberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the get invoice by number params
func (o *GetInvoiceByNumberParams) WithAudit(audit *string) *GetInvoiceByNumberParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get invoice by number params
func (o *GetInvoiceByNumberParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithInvoiceNumber adds the invoiceNumber to the get invoice by number params
func (o *GetInvoiceByNumberParams) WithInvoiceNumber(invoiceNumber int32) *GetInvoiceByNumberParams {
	o.SetInvoiceNumber(invoiceNumber)
	return o
}

// SetInvoiceNumber adds the invoiceNumber to the get invoice by number params
func (o *GetInvoiceByNumberParams) SetInvoiceNumber(invoiceNumber int32) {
	o.InvoiceNumber = invoiceNumber
}

// WithWithChildrenItems adds the withChildrenItems to the get invoice by number params
func (o *GetInvoiceByNumberParams) WithWithChildrenItems(withChildrenItems *bool) *GetInvoiceByNumberParams {
	o.SetWithChildrenItems(withChildrenItems)
	return o
}

// SetWithChildrenItems adds the withChildrenItems to the get invoice by number params
func (o *GetInvoiceByNumberParams) SetWithChildrenItems(withChildrenItems *bool) {
	o.WithChildrenItems = withChildrenItems
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoiceByNumberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Audit != nil {

		// query param audit
		var qrAudit string

		if o.Audit != nil {
			qrAudit = *o.Audit
		}
		qAudit := qrAudit
		if qAudit != "" {

			if err := r.SetQueryParam("audit", qAudit); err != nil {
				return err
			}
		}
	}

	// path param invoiceNumber
	if err := r.SetPathParam("invoiceNumber", swag.FormatInt32(o.InvoiceNumber)); err != nil {
		return err
	}

	if o.WithChildrenItems != nil {

		// query param withChildrenItems
		var qrWithChildrenItems bool

		if o.WithChildrenItems != nil {
			qrWithChildrenItems = *o.WithChildrenItems
		}
		qWithChildrenItems := swag.FormatBool(qrWithChildrenItems)
		if qWithChildrenItems != "" {

			if err := r.SetQueryParam("withChildrenItems", qWithChildrenItems); err != nil {
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
