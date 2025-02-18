// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// New creates a new catalog API client.
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
Client for catalog API
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
	AddSimplePlan(ctx context.Context, params *AddSimplePlanParams) (*AddSimplePlanCreated, error)

	DeleteCatalog(ctx context.Context, params *DeleteCatalogParams) (*DeleteCatalogNoContent, error)

	GetAvailableAddons(ctx context.Context, params *GetAvailableAddonsParams) (*GetAvailableAddonsOK, error)

	GetAvailableBasePlans(ctx context.Context, params *GetAvailableBasePlansParams) (*GetAvailableBasePlansOK, error)

	GetCatalogJSON(ctx context.Context, params *GetCatalogJSONParams) (*GetCatalogJSONOK, error)

	GetCatalogVersions(ctx context.Context, params *GetCatalogVersionsParams) (*GetCatalogVersionsOK, error)

	GetCatalogXML(ctx context.Context, params *GetCatalogXMLParams) (*GetCatalogXMLOK, error)

	GetPhaseForSubscriptionAndDate(ctx context.Context, params *GetPhaseForSubscriptionAndDateParams) (*GetPhaseForSubscriptionAndDateOK, error)

	GetPlanForSubscriptionAndDate(ctx context.Context, params *GetPlanForSubscriptionAndDateParams) (*GetPlanForSubscriptionAndDateOK, error)

	GetPriceListForSubscriptionAndDate(ctx context.Context, params *GetPriceListForSubscriptionAndDateParams) (*GetPriceListForSubscriptionAndDateOK, error)

	GetProductForSubscriptionAndDate(ctx context.Context, params *GetProductForSubscriptionAndDateParams) (*GetProductForSubscriptionAndDateOK, error)

	UploadCatalogXML(ctx context.Context, params *UploadCatalogXMLParams) (*UploadCatalogXMLCreated, error)

	ValidateCatalogXML(ctx context.Context, params *ValidateCatalogXMLParams) (*ValidateCatalogXMLOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddSimplePlan adds a simple plan entry in the current version of the catalog
*/
func (a *Client) AddSimplePlan(ctx context.Context, params *AddSimplePlanParams) (*AddSimplePlanCreated, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewAddSimplePlanParams()
	}
	getParams := NewAddSimplePlanParams()
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
		ID:                 "addSimplePlan",
		Method:             "POST",
		PathPattern:        "/1.0/kb/catalog/simplePlan",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddSimplePlanReader{formats: a.formats},
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
	case *AddSimplePlanCreated:
		location = kbcommon.ParseLocationHeader(value.HttpResponse.GetHeader("Location"))
		if !params.ProcessLocationHeader || location == "" {
			return value, nil
		}
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addSimplePlan",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &AddSimplePlanReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch value := getResult.(type) {
	case *AddSimplePlanCreated:
		return value, nil
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

}

/*
DeleteCatalog deletes all versions for a per tenant catalog
*/
func (a *Client) DeleteCatalog(ctx context.Context, params *DeleteCatalogParams) (*DeleteCatalogNoContent, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewDeleteCatalogParams()
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
		ID:                 "deleteCatalog",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/catalog",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteCatalogReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteCatalogNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteCatalog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetAvailableAddons retrieves available add ons for a given product
*/
func (a *Client) GetAvailableAddons(ctx context.Context, params *GetAvailableAddonsParams) (*GetAvailableAddonsOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetAvailableAddonsParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getAvailableAddons",
		Method:             "GET",
		PathPattern:        "/1.0/kb/catalog/availableAddons",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAvailableAddonsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAvailableAddonsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAvailableAddons: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetAvailableBasePlans retrieves available base plans
*/
func (a *Client) GetAvailableBasePlans(ctx context.Context, params *GetAvailableBasePlansParams) (*GetAvailableBasePlansOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetAvailableBasePlansParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getAvailableBasePlans",
		Method:             "GET",
		PathPattern:        "/1.0/kb/catalog/availableBasePlans",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAvailableBasePlansReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAvailableBasePlansOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAvailableBasePlans: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetCatalogJSON retrieves the catalog as JSON
*/
func (a *Client) GetCatalogJSON(ctx context.Context, params *GetCatalogJSONParams) (*GetCatalogJSONOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetCatalogJSONParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getCatalogJson",
		Method:             "GET",
		PathPattern:        "/1.0/kb/catalog",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCatalogJSONReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCatalogJSONOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCatalogJson: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetCatalogVersions retrieves a list of catalog versions
*/
func (a *Client) GetCatalogVersions(ctx context.Context, params *GetCatalogVersionsParams) (*GetCatalogVersionsOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetCatalogVersionsParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getCatalogVersions",
		Method:             "GET",
		PathPattern:        "/1.0/kb/catalog/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCatalogVersionsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCatalogVersionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCatalogVersions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetCatalogXML retrieves the full catalog as XML
*/
func (a *Client) GetCatalogXML(ctx context.Context, params *GetCatalogXMLParams) (*GetCatalogXMLOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetCatalogXMLParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getCatalogXml",
		Method:             "GET",
		PathPattern:        "/1.0/kb/catalog/xml",
		ProducesMediaTypes: []string{"text/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCatalogXMLReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCatalogXMLOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCatalogXml: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetPhaseForSubscriptionAndDate retrieves phase for a given subscription and date
*/
func (a *Client) GetPhaseForSubscriptionAndDate(ctx context.Context, params *GetPhaseForSubscriptionAndDateParams) (*GetPhaseForSubscriptionAndDateOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetPhaseForSubscriptionAndDateParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getPhaseForSubscriptionAndDate",
		Method:             "GET",
		PathPattern:        "/1.0/kb/catalog/phase",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPhaseForSubscriptionAndDateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPhaseForSubscriptionAndDateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPhaseForSubscriptionAndDate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetPlanForSubscriptionAndDate retrieves plan for a given subscription and date
*/
func (a *Client) GetPlanForSubscriptionAndDate(ctx context.Context, params *GetPlanForSubscriptionAndDateParams) (*GetPlanForSubscriptionAndDateOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetPlanForSubscriptionAndDateParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getPlanForSubscriptionAndDate",
		Method:             "GET",
		PathPattern:        "/1.0/kb/catalog/plan",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPlanForSubscriptionAndDateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPlanForSubscriptionAndDateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPlanForSubscriptionAndDate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetPriceListForSubscriptionAndDate retrieves price list for a given subscription and date
*/
func (a *Client) GetPriceListForSubscriptionAndDate(ctx context.Context, params *GetPriceListForSubscriptionAndDateParams) (*GetPriceListForSubscriptionAndDateOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetPriceListForSubscriptionAndDateParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getPriceListForSubscriptionAndDate",
		Method:             "GET",
		PathPattern:        "/1.0/kb/catalog/priceList",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPriceListForSubscriptionAndDateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPriceListForSubscriptionAndDateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPriceListForSubscriptionAndDate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetProductForSubscriptionAndDate retrieves product for a given subscription and date
*/
func (a *Client) GetProductForSubscriptionAndDate(ctx context.Context, params *GetProductForSubscriptionAndDateParams) (*GetProductForSubscriptionAndDateOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetProductForSubscriptionAndDateParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getProductForSubscriptionAndDate",
		Method:             "GET",
		PathPattern:        "/1.0/kb/catalog/product",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProductForSubscriptionAndDateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProductForSubscriptionAndDateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getProductForSubscriptionAndDate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
UploadCatalogXML uploads the full catalog as XML
*/
func (a *Client) UploadCatalogXML(ctx context.Context, params *UploadCatalogXMLParams) (*UploadCatalogXMLCreated, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewUploadCatalogXMLParams()
	}
	getParams := NewUploadCatalogXMLParams()
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
		ID:                 "uploadCatalogXml",
		Method:             "POST",
		PathPattern:        "/1.0/kb/catalog/xml",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UploadCatalogXMLReader{formats: a.formats},
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
	case *UploadCatalogXMLCreated:
		location = kbcommon.ParseLocationHeader(value.HttpResponse.GetHeader("Location"))
		if !params.ProcessLocationHeader || location == "" {
			return value, nil
		}
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "uploadCatalogXml",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/xml"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &UploadCatalogXMLReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch value := getResult.(type) {
	case *UploadCatalogXMLCreated:
		return value, nil
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

}

/*
ValidateCatalogXML validates a XML catalog
*/
func (a *Client) ValidateCatalogXML(ctx context.Context, params *ValidateCatalogXMLParams) (*ValidateCatalogXMLOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewValidateCatalogXMLParams()
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
		ID:                 "validateCatalogXml",
		Method:             "POST",
		PathPattern:        "/1.0/kb/catalog/xml/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ValidateCatalogXMLReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidateCatalogXMLOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for validateCatalogXml: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
