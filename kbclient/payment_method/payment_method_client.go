// Code generated by go-swagger; DO NOT EDIT.

package payment_method

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// New creates a new payment method API client.
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
Client for payment method API
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
	CreatePaymentMethodCustomFields(ctx context.Context, params *CreatePaymentMethodCustomFieldsParams) (*CreatePaymentMethodCustomFieldsCreated, error)

	DeletePaymentMethod(ctx context.Context, params *DeletePaymentMethodParams) (*DeletePaymentMethodNoContent, error)

	DeletePaymentMethodCustomFields(ctx context.Context, params *DeletePaymentMethodCustomFieldsParams) (*DeletePaymentMethodCustomFieldsNoContent, error)

	GetPaymentMethod(ctx context.Context, params *GetPaymentMethodParams) (*GetPaymentMethodOK, error)

	GetPaymentMethodAuditLogsWithHistory(ctx context.Context, params *GetPaymentMethodAuditLogsWithHistoryParams) (*GetPaymentMethodAuditLogsWithHistoryOK, error)

	GetPaymentMethodByKey(ctx context.Context, params *GetPaymentMethodByKeyParams) (*GetPaymentMethodByKeyOK, error)

	GetPaymentMethodCustomFields(ctx context.Context, params *GetPaymentMethodCustomFieldsParams) (*GetPaymentMethodCustomFieldsOK, error)

	GetPaymentMethods(ctx context.Context, params *GetPaymentMethodsParams) (*GetPaymentMethodsOK, error)

	ModifyPaymentMethodCustomFields(ctx context.Context, params *ModifyPaymentMethodCustomFieldsParams) (*ModifyPaymentMethodCustomFieldsNoContent, error)

	SearchPaymentMethods(ctx context.Context, params *SearchPaymentMethodsParams) (*SearchPaymentMethodsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreatePaymentMethodCustomFields adds custom fields to payment method
*/
func (a *Client) CreatePaymentMethodCustomFields(ctx context.Context, params *CreatePaymentMethodCustomFieldsParams) (*CreatePaymentMethodCustomFieldsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePaymentMethodCustomFieldsParams()
	}
	getParams := NewCreatePaymentMethodCustomFieldsParams()
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
		ID:                 "createPaymentMethodCustomFields",
		Method:             "POST",
		PathPattern:        "/1.0/kb/paymentMethods/{paymentMethodId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreatePaymentMethodCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	createdResult := result.(*CreatePaymentMethodCustomFieldsCreated)
	location := kbcommon.ParseLocationHeader(createdResult.HttpResponse.GetHeader("Location"))
	if !params.ProcessLocationHeader || location == "" {
		return createdResult, nil
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createPaymentMethodCustomFields",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &CreatePaymentMethodCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return getResult.(*CreatePaymentMethodCustomFieldsCreated), nil

}

/*
DeletePaymentMethod deletes a payment method
*/
func (a *Client) DeletePaymentMethod(ctx context.Context, params *DeletePaymentMethodParams) (*DeletePaymentMethodNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePaymentMethodParams()
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
		ID:                 "deletePaymentMethod",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/paymentMethods/{paymentMethodId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePaymentMethodReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePaymentMethodNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deletePaymentMethod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
DeletePaymentMethodCustomFields removes custom fields from payment method
*/
func (a *Client) DeletePaymentMethodCustomFields(ctx context.Context, params *DeletePaymentMethodCustomFieldsParams) (*DeletePaymentMethodCustomFieldsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePaymentMethodCustomFieldsParams()
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
		ID:                 "deletePaymentMethodCustomFields",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/paymentMethods/{paymentMethodId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePaymentMethodCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePaymentMethodCustomFieldsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deletePaymentMethodCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetPaymentMethod retrieves a payment method by id
*/
func (a *Client) GetPaymentMethod(ctx context.Context, params *GetPaymentMethodParams) (*GetPaymentMethodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentMethodParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getPaymentMethod",
		Method:             "GET",
		PathPattern:        "/1.0/kb/paymentMethods/{paymentMethodId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPaymentMethodReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPaymentMethodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPaymentMethod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetPaymentMethodAuditLogsWithHistory retrieves payment method audit logs with history by id
*/
func (a *Client) GetPaymentMethodAuditLogsWithHistory(ctx context.Context, params *GetPaymentMethodAuditLogsWithHistoryParams) (*GetPaymentMethodAuditLogsWithHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentMethodAuditLogsWithHistoryParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getPaymentMethodAuditLogsWithHistory",
		Method:             "GET",
		PathPattern:        "/1.0/kb/paymentMethods/{paymentMethodId}/auditLogsWithHistory",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPaymentMethodAuditLogsWithHistoryReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPaymentMethodAuditLogsWithHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPaymentMethodAuditLogsWithHistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetPaymentMethodByKey retrieves a payment method by external key
*/
func (a *Client) GetPaymentMethodByKey(ctx context.Context, params *GetPaymentMethodByKeyParams) (*GetPaymentMethodByKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentMethodByKeyParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getPaymentMethodByKey",
		Method:             "GET",
		PathPattern:        "/1.0/kb/paymentMethods",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPaymentMethodByKeyReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPaymentMethodByKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPaymentMethodByKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetPaymentMethodCustomFields retrieves payment method custom fields
*/
func (a *Client) GetPaymentMethodCustomFields(ctx context.Context, params *GetPaymentMethodCustomFieldsParams) (*GetPaymentMethodCustomFieldsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentMethodCustomFieldsParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getPaymentMethodCustomFields",
		Method:             "GET",
		PathPattern:        "/1.0/kb/paymentMethods/{paymentMethodId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPaymentMethodCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPaymentMethodCustomFieldsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPaymentMethodCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetPaymentMethods lists payment methods
*/
func (a *Client) GetPaymentMethods(ctx context.Context, params *GetPaymentMethodsParams) (*GetPaymentMethodsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentMethodsParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getPaymentMethods",
		Method:             "GET",
		PathPattern:        "/1.0/kb/paymentMethods/pagination",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPaymentMethodsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPaymentMethodsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPaymentMethods: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
ModifyPaymentMethodCustomFields modifies custom fields to payment method
*/
func (a *Client) ModifyPaymentMethodCustomFields(ctx context.Context, params *ModifyPaymentMethodCustomFieldsParams) (*ModifyPaymentMethodCustomFieldsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyPaymentMethodCustomFieldsParams()
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
		ID:                 "modifyPaymentMethodCustomFields",
		Method:             "PUT",
		PathPattern:        "/1.0/kb/paymentMethods/{paymentMethodId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifyPaymentMethodCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ModifyPaymentMethodCustomFieldsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for modifyPaymentMethodCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
SearchPaymentMethods searches payment methods
*/
func (a *Client) SearchPaymentMethods(ctx context.Context, params *SearchPaymentMethodsParams) (*SearchPaymentMethodsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchPaymentMethodsParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "searchPaymentMethods",
		Method:             "GET",
		PathPattern:        "/1.0/kb/paymentMethods/search/{searchKey}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SearchPaymentMethodsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchPaymentMethodsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchPaymentMethods: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
