// Code generated by go-swagger; DO NOT EDIT.

package overdue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// New creates a new overdue API client.
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
Client for overdue API
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
	GetOverdueConfigJSON(ctx context.Context, params *GetOverdueConfigJSONParams) (*GetOverdueConfigJSONOK, error)

	GetOverdueConfigXML(ctx context.Context, params *GetOverdueConfigXMLParams) (*GetOverdueConfigXMLOK, error)

	UploadOverdueConfigJSON(ctx context.Context, params *UploadOverdueConfigJSONParams) (*UploadOverdueConfigJSONCreated, error)

	UploadOverdueConfigXML(ctx context.Context, params *UploadOverdueConfigXMLParams) (*UploadOverdueConfigXMLCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetOverdueConfigJSON retrieves the overdue config as JSON
*/
func (a *Client) GetOverdueConfigJSON(ctx context.Context, params *GetOverdueConfigJSONParams) (*GetOverdueConfigJSONOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetOverdueConfigJSONParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getOverdueConfigJson",
		Method:             "GET",
		PathPattern:        "/1.0/kb/overdue",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOverdueConfigJSONReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOverdueConfigJSONOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOverdueConfigJson: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetOverdueConfigXML retrieves the overdue config as XML
*/
func (a *Client) GetOverdueConfigXML(ctx context.Context, params *GetOverdueConfigXMLParams) (*GetOverdueConfigXMLOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetOverdueConfigXMLParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getOverdueConfigXml",
		Method:             "GET",
		PathPattern:        "/1.0/kb/overdue/xml",
		ProducesMediaTypes: []string{"text/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOverdueConfigXMLReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOverdueConfigXMLOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOverdueConfigXml: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
UploadOverdueConfigJSON uploads the full overdue config as JSON
*/
func (a *Client) UploadOverdueConfigJSON(ctx context.Context, params *UploadOverdueConfigJSONParams) (*UploadOverdueConfigJSONCreated, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewUploadOverdueConfigJSONParams()
	}
	getParams := NewUploadOverdueConfigJSONParams()
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
		ID:                 "uploadOverdueConfigJson",
		Method:             "POST",
		PathPattern:        "/1.0/kb/overdue",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UploadOverdueConfigJSONReader{formats: a.formats},
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
	case *UploadOverdueConfigJSONCreated:
		location = kbcommon.ParseLocationHeader(value.HttpResponse.GetHeader("Location"))
		if !params.ProcessLocationHeader || location == "" {
			return value, nil
		}
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "uploadOverdueConfigJson",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &UploadOverdueConfigJSONReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch value := getResult.(type) {
	case *UploadOverdueConfigJSONCreated:
		return value, nil
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

}

/*
UploadOverdueConfigXML uploads the full overdue config as XML
*/
func (a *Client) UploadOverdueConfigXML(ctx context.Context, params *UploadOverdueConfigXMLParams) (*UploadOverdueConfigXMLCreated, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewUploadOverdueConfigXMLParams()
	}
	getParams := NewUploadOverdueConfigXMLParams()
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
		ID:                 "uploadOverdueConfigXml",
		Method:             "POST",
		PathPattern:        "/1.0/kb/overdue/xml",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UploadOverdueConfigXMLReader{formats: a.formats},
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
	case *UploadOverdueConfigXMLCreated:
		location = kbcommon.ParseLocationHeader(value.HttpResponse.GetHeader("Location"))
		if !params.ProcessLocationHeader || location == "" {
			return value, nil
		}
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "uploadOverdueConfigXml",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/xml"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &UploadOverdueConfigXMLReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch value := getResult.(type) {
	case *UploadOverdueConfigXMLCreated:
		return value, nil
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
