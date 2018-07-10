// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new subscription API client.
func New(transport runtime.ClientTransport,
	formats strfmt.Registry,
	authInfo runtime.ClientAuthInfoWriter,
	defaults KillbillDefaults) *Client {

	return &Client{transport: transport, formats: formats, authInfo: authInfo, defaults: defaults}
}

// killbill default values. When a call is made to an operation, these values are used
// if params doesn't specify them.
type KillbillDefaults interface {
	// Default API Key. If not set explicitly in params, this will be used.
	XKillbillAPIKey() *string
	// Default API Secret. If not set explicitly in params, this will be used.
	XKillbillAPISecret() *string
	// Default CreatedBy. If not set explicitly in params, this will be used.
	XKillbillCreatedBy() *string
	// Default Comment. If not set explicitly in params, this will be used.
	XKillbillComment() *string
	// Default Reason. If not set explicitly in params, this will be used.
	XKillbillReason() *string
	// Default WithStackTrace. If not set explicitly in params, this will be used.
	KillbillWithStackTrace() *bool
}

/*
Client for subscription API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
	defaults  KillbillDefaults
}

/*
AddSubscriptionBlockingState blocks a subscription
*/
func (a *Client) AddSubscriptionBlockingState(ctx context.Context, params *AddSubscriptionBlockingStateParams) (*AddSubscriptionBlockingStateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddSubscriptionBlockingStateParams()
	}
	getParams := NewAddSubscriptionBlockingStateParams()
	getParams.Context = ctx
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}
	getParams.XKillbillAPIKey = params.XKillbillAPIKey

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}
	getParams.XKillbillAPISecret = params.XKillbillAPISecret

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

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	getParams.WithStackTrace = params.WithStackTrace

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addSubscriptionBlockingState",
		Method:             "POST",
		PathPattern:        "/1.0/kb/subscriptions/{subscriptionId}/block",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddSubscriptionBlockingStateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	createdResult := result.(*AddSubscriptionBlockingStateCreated)
	location := kbcommon.ParseLocationHeader(createdResult.HttpResponse.GetHeader("Location"))
	if !params.ProcessLocationHeader || location == "" {
		return createdResult, nil
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addSubscriptionBlockingState",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &AddSubscriptionBlockingStateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return getResult.(*AddSubscriptionBlockingStateCreated), nil

}

/*
CancelSubscriptionPlan cancels an entitlement plan
*/
func (a *Client) CancelSubscriptionPlan(ctx context.Context, params *CancelSubscriptionPlanParams) (*CancelSubscriptionPlanNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelSubscriptionPlanParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "cancelSubscriptionPlan",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/subscriptions/{subscriptionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CancelSubscriptionPlanReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CancelSubscriptionPlanNoContent), nil

}

/*
ChangeSubscriptionPlan changes entitlement plan
*/
func (a *Client) ChangeSubscriptionPlan(ctx context.Context, params *ChangeSubscriptionPlanParams) (*ChangeSubscriptionPlanNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeSubscriptionPlanParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeSubscriptionPlan",
		Method:             "PUT",
		PathPattern:        "/1.0/kb/subscriptions/{subscriptionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ChangeSubscriptionPlanReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeSubscriptionPlanNoContent), nil

}

/*
CreateSubscription creates an subscription
*/
func (a *Client) CreateSubscription(ctx context.Context, params *CreateSubscriptionParams) (*CreateSubscriptionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSubscriptionParams()
	}
	getParams := NewCreateSubscriptionParams()
	getParams.Context = ctx
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}
	getParams.XKillbillAPIKey = params.XKillbillAPIKey

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}
	getParams.XKillbillAPISecret = params.XKillbillAPISecret

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

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	getParams.WithStackTrace = params.WithStackTrace

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSubscription",
		Method:             "POST",
		PathPattern:        "/1.0/kb/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSubscriptionReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	createdResult := result.(*CreateSubscriptionCreated)
	location := kbcommon.ParseLocationHeader(createdResult.HttpResponse.GetHeader("Location"))
	if !params.ProcessLocationHeader || location == "" {
		return createdResult, nil
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSubscription",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &CreateSubscriptionReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return getResult.(*CreateSubscriptionCreated), nil

}

/*
CreateSubscriptionCustomFields adds custom fields to subscription
*/
func (a *Client) CreateSubscriptionCustomFields(ctx context.Context, params *CreateSubscriptionCustomFieldsParams) (*CreateSubscriptionCustomFieldsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSubscriptionCustomFieldsParams()
	}
	getParams := NewCreateSubscriptionCustomFieldsParams()
	getParams.Context = ctx
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}
	getParams.XKillbillAPIKey = params.XKillbillAPIKey

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}
	getParams.XKillbillAPISecret = params.XKillbillAPISecret

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

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	getParams.WithStackTrace = params.WithStackTrace

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSubscriptionCustomFields",
		Method:             "POST",
		PathPattern:        "/1.0/kb/subscriptions/{subscriptionId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSubscriptionCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	createdResult := result.(*CreateSubscriptionCustomFieldsCreated)
	location := kbcommon.ParseLocationHeader(createdResult.HttpResponse.GetHeader("Location"))
	if !params.ProcessLocationHeader || location == "" {
		return createdResult, nil
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSubscriptionCustomFields",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &CreateSubscriptionCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return getResult.(*CreateSubscriptionCustomFieldsCreated), nil

}

/*
CreateSubscriptionTags create subscription tags API
*/
func (a *Client) CreateSubscriptionTags(ctx context.Context, params *CreateSubscriptionTagsParams) (*CreateSubscriptionTagsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSubscriptionTagsParams()
	}
	getParams := NewCreateSubscriptionTagsParams()
	getParams.Context = ctx
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}
	getParams.XKillbillAPIKey = params.XKillbillAPIKey

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}
	getParams.XKillbillAPISecret = params.XKillbillAPISecret

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

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	getParams.WithStackTrace = params.WithStackTrace

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSubscriptionTags",
		Method:             "POST",
		PathPattern:        "/1.0/kb/subscriptions/{subscriptionId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSubscriptionTagsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	createdResult := result.(*CreateSubscriptionTagsCreated)
	location := kbcommon.ParseLocationHeader(createdResult.HttpResponse.GetHeader("Location"))
	if !params.ProcessLocationHeader || location == "" {
		return createdResult, nil
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSubscriptionTags",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &CreateSubscriptionTagsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return getResult.(*CreateSubscriptionTagsCreated), nil

}

/*
CreateSubscriptionWithAddOns creates an entitlement with add on products
*/
func (a *Client) CreateSubscriptionWithAddOns(ctx context.Context, params *CreateSubscriptionWithAddOnsParams) (*CreateSubscriptionWithAddOnsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSubscriptionWithAddOnsParams()
	}
	getParams := NewCreateSubscriptionWithAddOnsParams()
	getParams.Context = ctx
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}
	getParams.XKillbillAPIKey = params.XKillbillAPIKey

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}
	getParams.XKillbillAPISecret = params.XKillbillAPISecret

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

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	getParams.WithStackTrace = params.WithStackTrace

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSubscriptionWithAddOns",
		Method:             "POST",
		PathPattern:        "/1.0/kb/subscriptions/createSubscriptionWithAddOns",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSubscriptionWithAddOnsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	createdResult := result.(*CreateSubscriptionWithAddOnsCreated)
	location := kbcommon.ParseLocationHeader(createdResult.HttpResponse.GetHeader("Location"))
	if !params.ProcessLocationHeader || location == "" {
		return createdResult, nil
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSubscriptionWithAddOns",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &CreateSubscriptionWithAddOnsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return getResult.(*CreateSubscriptionWithAddOnsCreated), nil

}

/*
CreateSubscriptionsWithAddOns creates multiple entitlements with add on products
*/
func (a *Client) CreateSubscriptionsWithAddOns(ctx context.Context, params *CreateSubscriptionsWithAddOnsParams) (*CreateSubscriptionsWithAddOnsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSubscriptionsWithAddOnsParams()
	}
	getParams := NewCreateSubscriptionsWithAddOnsParams()
	getParams.Context = ctx
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}
	getParams.XKillbillAPIKey = params.XKillbillAPIKey

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}
	getParams.XKillbillAPISecret = params.XKillbillAPISecret

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

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	getParams.WithStackTrace = params.WithStackTrace

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSubscriptionsWithAddOns",
		Method:             "POST",
		PathPattern:        "/1.0/kb/subscriptions/createSubscriptionsWithAddOns",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSubscriptionsWithAddOnsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	createdResult := result.(*CreateSubscriptionsWithAddOnsCreated)
	location := kbcommon.ParseLocationHeader(createdResult.HttpResponse.GetHeader("Location"))
	if !params.ProcessLocationHeader || location == "" {
		return createdResult, nil
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSubscriptionsWithAddOns",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &CreateSubscriptionsWithAddOnsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return getResult.(*CreateSubscriptionsWithAddOnsCreated), nil

}

/*
DeleteSubscriptionCustomFields removes custom fields from subscription
*/
func (a *Client) DeleteSubscriptionCustomFields(ctx context.Context, params *DeleteSubscriptionCustomFieldsParams) (*DeleteSubscriptionCustomFieldsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSubscriptionCustomFieldsParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSubscriptionCustomFields",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/subscriptions/{subscriptionId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSubscriptionCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSubscriptionCustomFieldsNoContent), nil

}

/*
DeleteSubscriptionTags removes tags from subscription
*/
func (a *Client) DeleteSubscriptionTags(ctx context.Context, params *DeleteSubscriptionTagsParams) (*DeleteSubscriptionTagsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSubscriptionTagsParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSubscriptionTags",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/subscriptions/{subscriptionId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSubscriptionTagsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSubscriptionTagsNoContent), nil

}

/*
GetSubscription retrieves a subscription by id
*/
func (a *Client) GetSubscription(ctx context.Context, params *GetSubscriptionParams) (*GetSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSubscriptionParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSubscription",
		Method:             "GET",
		PathPattern:        "/1.0/kb/subscriptions/{subscriptionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSubscriptionReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSubscriptionOK), nil

}

/*
GetSubscriptionCustomFields retrieves subscription custom fields
*/
func (a *Client) GetSubscriptionCustomFields(ctx context.Context, params *GetSubscriptionCustomFieldsParams) (*GetSubscriptionCustomFieldsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSubscriptionCustomFieldsParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSubscriptionCustomFields",
		Method:             "GET",
		PathPattern:        "/1.0/kb/subscriptions/{subscriptionId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSubscriptionCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSubscriptionCustomFieldsOK), nil

}

/*
GetSubscriptionTags retrieves subscription tags
*/
func (a *Client) GetSubscriptionTags(ctx context.Context, params *GetSubscriptionTagsParams) (*GetSubscriptionTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSubscriptionTagsParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSubscriptionTags",
		Method:             "GET",
		PathPattern:        "/1.0/kb/subscriptions/{subscriptionId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSubscriptionTagsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSubscriptionTagsOK), nil

}

/*
ModifySubscriptionCustomFields modifies custom fields to subscription
*/
func (a *Client) ModifySubscriptionCustomFields(ctx context.Context, params *ModifySubscriptionCustomFieldsParams) (*ModifySubscriptionCustomFieldsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifySubscriptionCustomFieldsParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "modifySubscriptionCustomFields",
		Method:             "PUT",
		PathPattern:        "/1.0/kb/subscriptions/{subscriptionId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifySubscriptionCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifySubscriptionCustomFieldsNoContent), nil

}

/*
UncancelSubscriptionPlan uns cancel an entitlement
*/
func (a *Client) UncancelSubscriptionPlan(ctx context.Context, params *UncancelSubscriptionPlanParams) (*UncancelSubscriptionPlanNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUncancelSubscriptionPlanParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "uncancelSubscriptionPlan",
		Method:             "PUT",
		PathPattern:        "/1.0/kb/subscriptions/{subscriptionId}/uncancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UncancelSubscriptionPlanReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UncancelSubscriptionPlanNoContent), nil

}

/*
UndoChangeSubscriptionPlan undos a pending change plan on an entitlement
*/
func (a *Client) UndoChangeSubscriptionPlan(ctx context.Context, params *UndoChangeSubscriptionPlanParams) (*UndoChangeSubscriptionPlanNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUndoChangeSubscriptionPlanParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "undoChangeSubscriptionPlan",
		Method:             "PUT",
		PathPattern:        "/1.0/kb/subscriptions/{subscriptionId}/undoChangePlan",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UndoChangeSubscriptionPlanReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UndoChangeSubscriptionPlanNoContent), nil

}

/*
UpdateSubscriptionBCD updates the b c d associated to a subscription
*/
func (a *Client) UpdateSubscriptionBCD(ctx context.Context, params *UpdateSubscriptionBCDParams) (*UpdateSubscriptionBCDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSubscriptionBCDParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSubscriptionBCD",
		Method:             "PUT",
		PathPattern:        "/1.0/kb/subscriptions/{subscriptionId}/bcd",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSubscriptionBCDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateSubscriptionBCDNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}