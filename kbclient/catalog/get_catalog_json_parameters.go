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

// NewGetCatalogJSONParams creates a new GetCatalogJSONParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCatalogJSONParams() *GetCatalogJSONParams {
	return &GetCatalogJSONParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCatalogJSONParamsWithTimeout creates a new GetCatalogJSONParams object
// with the ability to set a timeout on a request.
func NewGetCatalogJSONParamsWithTimeout(timeout time.Duration) *GetCatalogJSONParams {
	return &GetCatalogJSONParams{
		timeout: timeout,
	}
}

// NewGetCatalogJSONParamsWithContext creates a new GetCatalogJSONParams object
// with the ability to set a context for a request.
func NewGetCatalogJSONParamsWithContext(ctx context.Context) *GetCatalogJSONParams {
	return &GetCatalogJSONParams{
		Context: ctx,
	}
}

// NewGetCatalogJSONParamsWithHTTPClient creates a new GetCatalogJSONParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCatalogJSONParamsWithHTTPClient(client *http.Client) *GetCatalogJSONParams {
	return &GetCatalogJSONParams{
		HTTPClient: client,
	}
}

/*
GetCatalogJSONParams contains all the parameters to send to the API endpoint

	for the get catalog Json operation.

	Typically these are written to a http.Request.
*/
type GetCatalogJSONParams struct {

	// AccountID.
	//
	// Format: uuid
	AccountID *strfmt.UUID

	// RequestedDate.
	//
	// Format: date-time
	RequestedDate *strfmt.DateTime

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the get catalog Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCatalogJSONParams) WithDefaults() *GetCatalogJSONParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get catalog Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCatalogJSONParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get catalog Json params
func (o *GetCatalogJSONParams) WithTimeout(timeout time.Duration) *GetCatalogJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get catalog Json params
func (o *GetCatalogJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get catalog Json params
func (o *GetCatalogJSONParams) WithContext(ctx context.Context) *GetCatalogJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get catalog Json params
func (o *GetCatalogJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get catalog Json params
func (o *GetCatalogJSONParams) WithHTTPClient(client *http.Client) *GetCatalogJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get catalog Json params
func (o *GetCatalogJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get catalog Json params
func (o *GetCatalogJSONParams) WithAccountID(accountID *strfmt.UUID) *GetCatalogJSONParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get catalog Json params
func (o *GetCatalogJSONParams) SetAccountID(accountID *strfmt.UUID) {
	o.AccountID = accountID
}

// WithRequestedDate adds the requestedDate to the get catalog Json params
func (o *GetCatalogJSONParams) WithRequestedDate(requestedDate *strfmt.DateTime) *GetCatalogJSONParams {
	o.SetRequestedDate(requestedDate)
	return o
}

// SetRequestedDate adds the requestedDate to the get catalog Json params
func (o *GetCatalogJSONParams) SetRequestedDate(requestedDate *strfmt.DateTime) {
	o.RequestedDate = requestedDate
}

// WriteToRequest writes these params to a swagger request
func (o *GetCatalogJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountID != nil {

		// query param accountId
		var qrAccountID strfmt.UUID

		if o.AccountID != nil {
			qrAccountID = *o.AccountID
		}
		qAccountID := qrAccountID.String()
		if qAccountID != "" {

			if err := r.SetQueryParam("accountId", qAccountID); err != nil {
				return err
			}
		}
	}

	if o.RequestedDate != nil {

		// query param requestedDate
		var qrRequestedDate strfmt.DateTime

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
