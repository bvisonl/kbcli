// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// New creates a new invoice payment API client.
func New(transport runtime.ClientTransport,
	formats strfmt.Registry,
	authInfo runtime.ClientAuthInfoWriter,
	defaults KillbillDefaults) *Client {

	return &Client{transport: transport, formats: formats, authInfo: authInfo, defaults: defaults}
}

// killbill default values. When a call is made to an operation, these values are used
// if params doesn't specify them.
type KillbillDefaults interface {
	// Default CreatedBy. If not set explicitly in params, this will be used.
	XKillbillCreatedBy() *string
	// Default Comment. If not set explicitly in params, this will be used.
	XKillbillComment() *string
	// Default Reason. If not set explicitly in params, this will be used.
	XKillbillReason() *string
	// Default WithWithProfilingInfo. If not set explicitly in params, this will be used.
	KillbillWithProfilingInfo() *string
	// Default WithStackTrace. If not set explicitly in params, this will be used.
	KillbillWithStackTrace() *bool
}

/*
Client for invoice payment API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
	defaults  KillbillDefaults
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CompleteInvoicePaymentTransaction(ctx context.Context, params *CompleteInvoicePaymentTransactionParams) (*CompleteInvoicePaymentTransactionNoContent, error)

	CreateChargeback(ctx context.Context, params *CreateChargebackParams) (*CreateChargebackCreated, error)

	CreateChargebackReversal(ctx context.Context, params *CreateChargebackReversalParams) (*CreateChargebackReversalCreated, error)

	CreateInvoicePaymentCustomFields(ctx context.Context, params *CreateInvoicePaymentCustomFieldsParams) (*CreateInvoicePaymentCustomFieldsCreated, error)

	CreateInvoicePaymentTags(ctx context.Context, params *CreateInvoicePaymentTagsParams) (*CreateInvoicePaymentTagsCreated, error)

	CreateRefundWithAdjustments(ctx context.Context, params *CreateRefundWithAdjustmentsParams) (*CreateRefundWithAdjustmentsCreated, error)

	DeleteInvoicePaymentCustomFields(ctx context.Context, params *DeleteInvoicePaymentCustomFieldsParams) (*DeleteInvoicePaymentCustomFieldsNoContent, error)

	DeleteInvoicePaymentTags(ctx context.Context, params *DeleteInvoicePaymentTagsParams) (*DeleteInvoicePaymentTagsNoContent, error)

	GetInvoicePayment(ctx context.Context, params *GetInvoicePaymentParams) (*GetInvoicePaymentOK, error)

	GetInvoicePaymentAuditLogsWithHistory(ctx context.Context, params *GetInvoicePaymentAuditLogsWithHistoryParams) (*GetInvoicePaymentAuditLogsWithHistoryOK, error)

	GetInvoicePaymentCustomFields(ctx context.Context, params *GetInvoicePaymentCustomFieldsParams) (*GetInvoicePaymentCustomFieldsOK, error)

	GetInvoicePaymentTags(ctx context.Context, params *GetInvoicePaymentTagsParams) (*GetInvoicePaymentTagsOK, error)

	ModifyInvoicePaymentCustomFields(ctx context.Context, params *ModifyInvoicePaymentCustomFieldsParams) (*ModifyInvoicePaymentCustomFieldsNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CompleteInvoicePaymentTransaction completes an existing transaction
*/
func (a *Client) CompleteInvoicePaymentTransaction(ctx context.Context, params *CompleteInvoicePaymentTransactionParams) (*CompleteInvoicePaymentTransactionNoContent, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewCompleteInvoicePaymentTransactionParams()
	}
	params.Context = ctx
	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "completeInvoicePaymentTransaction",
		Method:             "PUT",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CompleteInvoicePaymentTransactionReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CompleteInvoicePaymentTransactionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for completeInvoicePaymentTransaction: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
CreateChargeback records a chargeback
*/
func (a *Client) CreateChargeback(ctx context.Context, params *CreateChargebackParams) (*CreateChargebackCreated, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewCreateChargebackParams()
	}
	getParams := NewCreateChargebackParams()
	getParams.Context = ctx
	params.Context = ctx
	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}
	getParams.XKillbillComment = params.XKillbillComment
	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}
	getParams.XKillbillCreatedBy = params.XKillbillCreatedBy
	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}
	getParams.XKillbillReason = params.XKillbillReason
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	getParams.WithStackTrace = params.WithStackTrace

	op := &runtime.ClientOperation{
		ID:                 "createChargeback",
		Method:             "POST",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}/chargebacks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateChargebackReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	var location string
	switch value := result.(type) {
	case *CreateChargebackCreated:
		location = kbcommon.ParseLocationHeader(value.HttpResponse.GetHeader("Location"))
		if !params.ProcessLocationHeader || location == "" {
			return value, nil
		}
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createChargeback",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &CreateChargebackReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch value := getResult.(type) {
	case *CreateChargebackCreated:
		return value, nil
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

}

/*
CreateChargebackReversal records a chargeback reversal
*/
func (a *Client) CreateChargebackReversal(ctx context.Context, params *CreateChargebackReversalParams) (*CreateChargebackReversalCreated, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewCreateChargebackReversalParams()
	}
	getParams := NewCreateChargebackReversalParams()
	getParams.Context = ctx
	params.Context = ctx
	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}
	getParams.XKillbillComment = params.XKillbillComment
	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}
	getParams.XKillbillCreatedBy = params.XKillbillCreatedBy
	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}
	getParams.XKillbillReason = params.XKillbillReason
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	getParams.WithStackTrace = params.WithStackTrace

	op := &runtime.ClientOperation{
		ID:                 "createChargebackReversal",
		Method:             "POST",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}/chargebackReversals",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateChargebackReversalReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	var location string
	switch value := result.(type) {
	case *CreateChargebackReversalCreated:
		location = kbcommon.ParseLocationHeader(value.HttpResponse.GetHeader("Location"))
		if !params.ProcessLocationHeader || location == "" {
			return value, nil
		}
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createChargebackReversal",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &CreateChargebackReversalReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch value := getResult.(type) {
	case *CreateChargebackReversalCreated:
		return value, nil
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

}

/*
CreateInvoicePaymentCustomFields adds custom fields to payment
*/
func (a *Client) CreateInvoicePaymentCustomFields(ctx context.Context, params *CreateInvoicePaymentCustomFieldsParams) (*CreateInvoicePaymentCustomFieldsCreated, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewCreateInvoicePaymentCustomFieldsParams()
	}
	getParams := NewCreateInvoicePaymentCustomFieldsParams()
	getParams.Context = ctx
	params.Context = ctx
	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}
	getParams.XKillbillComment = params.XKillbillComment
	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}
	getParams.XKillbillCreatedBy = params.XKillbillCreatedBy
	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}
	getParams.XKillbillReason = params.XKillbillReason
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	getParams.WithStackTrace = params.WithStackTrace

	op := &runtime.ClientOperation{
		ID:                 "createInvoicePaymentCustomFields",
		Method:             "POST",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateInvoicePaymentCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	var location string
	switch value := result.(type) {
	case *CreateInvoicePaymentCustomFieldsCreated:
		location = kbcommon.ParseLocationHeader(value.HttpResponse.GetHeader("Location"))
		if !params.ProcessLocationHeader || location == "" {
			return value, nil
		}
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createInvoicePaymentCustomFields",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &CreateInvoicePaymentCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch value := getResult.(type) {
	case *CreateInvoicePaymentCustomFieldsCreated:
		return value, nil
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

}

/*
CreateInvoicePaymentTags adds tags to payment
*/
func (a *Client) CreateInvoicePaymentTags(ctx context.Context, params *CreateInvoicePaymentTagsParams) (*CreateInvoicePaymentTagsCreated, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewCreateInvoicePaymentTagsParams()
	}
	getParams := NewCreateInvoicePaymentTagsParams()
	getParams.Context = ctx
	params.Context = ctx
	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}
	getParams.XKillbillComment = params.XKillbillComment
	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}
	getParams.XKillbillCreatedBy = params.XKillbillCreatedBy
	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}
	getParams.XKillbillReason = params.XKillbillReason
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	getParams.WithStackTrace = params.WithStackTrace

	op := &runtime.ClientOperation{
		ID:                 "createInvoicePaymentTags",
		Method:             "POST",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateInvoicePaymentTagsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	var location string
	switch value := result.(type) {
	case *CreateInvoicePaymentTagsCreated:
		location = kbcommon.ParseLocationHeader(value.HttpResponse.GetHeader("Location"))
		if !params.ProcessLocationHeader || location == "" {
			return value, nil
		}
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createInvoicePaymentTags",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &CreateInvoicePaymentTagsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch value := getResult.(type) {
	case *CreateInvoicePaymentTagsCreated:
		return value, nil
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

}

/*
CreateRefundWithAdjustments refunds a payment and adjust the invoice if needed
*/
func (a *Client) CreateRefundWithAdjustments(ctx context.Context, params *CreateRefundWithAdjustmentsParams) (*CreateRefundWithAdjustmentsCreated, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewCreateRefundWithAdjustmentsParams()
	}
	getParams := NewCreateRefundWithAdjustmentsParams()
	getParams.Context = ctx
	params.Context = ctx
	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}
	getParams.XKillbillComment = params.XKillbillComment
	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}
	getParams.XKillbillCreatedBy = params.XKillbillCreatedBy
	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}
	getParams.XKillbillReason = params.XKillbillReason
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	getParams.WithStackTrace = params.WithStackTrace

	op := &runtime.ClientOperation{
		ID:                 "createRefundWithAdjustments",
		Method:             "POST",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}/refunds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateRefundWithAdjustmentsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	var location string
	switch value := result.(type) {
	case *CreateRefundWithAdjustmentsCreated:
		location = kbcommon.ParseLocationHeader(value.HttpResponse.GetHeader("Location"))
		if !params.ProcessLocationHeader || location == "" {
			return value, nil
		}
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createRefundWithAdjustments",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &CreateRefundWithAdjustmentsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch value := getResult.(type) {
	case *CreateRefundWithAdjustmentsCreated:
		return value, nil
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

}

/*
DeleteInvoicePaymentCustomFields removes custom fields from payment
*/
func (a *Client) DeleteInvoicePaymentCustomFields(ctx context.Context, params *DeleteInvoicePaymentCustomFieldsParams) (*DeleteInvoicePaymentCustomFieldsNoContent, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewDeleteInvoicePaymentCustomFieldsParams()
	}
	params.Context = ctx
	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "deleteInvoicePaymentCustomFields",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteInvoicePaymentCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteInvoicePaymentCustomFieldsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteInvoicePaymentCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
DeleteInvoicePaymentTags removes tags from payment
*/
func (a *Client) DeleteInvoicePaymentTags(ctx context.Context, params *DeleteInvoicePaymentTagsParams) (*DeleteInvoicePaymentTagsNoContent, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewDeleteInvoicePaymentTagsParams()
	}
	params.Context = ctx
	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "deleteInvoicePaymentTags",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteInvoicePaymentTagsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteInvoicePaymentTagsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteInvoicePaymentTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetInvoicePayment retrieves a payment by id
*/
func (a *Client) GetInvoicePayment(ctx context.Context, params *GetInvoicePaymentParams) (*GetInvoicePaymentOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetInvoicePaymentParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getInvoicePayment",
		Method:             "GET",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInvoicePaymentReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInvoicePaymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInvoicePayment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetInvoicePaymentAuditLogsWithHistory retrieves invoice payment audit logs with history by id
*/
func (a *Client) GetInvoicePaymentAuditLogsWithHistory(ctx context.Context, params *GetInvoicePaymentAuditLogsWithHistoryParams) (*GetInvoicePaymentAuditLogsWithHistoryOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetInvoicePaymentAuditLogsWithHistoryParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getInvoicePaymentAuditLogsWithHistory",
		Method:             "GET",
		PathPattern:        "/1.0/kb/invoicePayments/{invoicePaymentId}/auditLogsWithHistory",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInvoicePaymentAuditLogsWithHistoryReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInvoicePaymentAuditLogsWithHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInvoicePaymentAuditLogsWithHistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetInvoicePaymentCustomFields retrieves payment custom fields
*/
func (a *Client) GetInvoicePaymentCustomFields(ctx context.Context, params *GetInvoicePaymentCustomFieldsParams) (*GetInvoicePaymentCustomFieldsOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetInvoicePaymentCustomFieldsParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getInvoicePaymentCustomFields",
		Method:             "GET",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInvoicePaymentCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInvoicePaymentCustomFieldsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInvoicePaymentCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetInvoicePaymentTags retrieves payment tags
*/
func (a *Client) GetInvoicePaymentTags(ctx context.Context, params *GetInvoicePaymentTagsParams) (*GetInvoicePaymentTagsOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetInvoicePaymentTagsParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getInvoicePaymentTags",
		Method:             "GET",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInvoicePaymentTagsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInvoicePaymentTagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInvoicePaymentTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
ModifyInvoicePaymentCustomFields modifies custom fields to payment
*/
func (a *Client) ModifyInvoicePaymentCustomFields(ctx context.Context, params *ModifyInvoicePaymentCustomFieldsParams) (*ModifyInvoicePaymentCustomFieldsNoContent, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewModifyInvoicePaymentCustomFieldsParams()
	}
	params.Context = ctx
	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "modifyInvoicePaymentCustomFields",
		Method:             "PUT",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifyInvoicePaymentCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ModifyInvoicePaymentCustomFieldsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for modifyInvoicePaymentCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
