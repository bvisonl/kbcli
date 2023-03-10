// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"

	"github.com/killbill/kbcli/v3/kbmodel"
)

// UploadPluginPaymentStateMachineConfigReader is a Reader for the UploadPluginPaymentStateMachineConfig structure.
type UploadPluginPaymentStateMachineConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadPluginPaymentStateMachineConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewUploadPluginPaymentStateMachineConfigCreated()
		result.HttpResponse = response
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		errorResult := kbcommon.NewKillbillError(response.Code())
		if err := consumer.Consume(response.Body(), &errorResult); err != nil && err != io.EOF {
			return nil, err
		}
		return nil, errorResult
	}
}

// NewUploadPluginPaymentStateMachineConfigCreated creates a UploadPluginPaymentStateMachineConfigCreated with default headers values
func NewUploadPluginPaymentStateMachineConfigCreated() *UploadPluginPaymentStateMachineConfigCreated {
	return &UploadPluginPaymentStateMachineConfigCreated{}
}

/*
UploadPluginPaymentStateMachineConfigCreated describes a response with status code 201, with default header values.

Per tenant state machine uploaded successfully
*/
type UploadPluginPaymentStateMachineConfigCreated struct {
	Payload      *kbmodel.TenantKeyValue
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the upload plugin payment state machine config created response
func (o *UploadPluginPaymentStateMachineConfigCreated) Code() int {
	return 201
}

// IsSuccess returns true when this upload plugin payment state machine config created response has a 2xx status code
func (o *UploadPluginPaymentStateMachineConfigCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload plugin payment state machine config created response has a 3xx status code
func (o *UploadPluginPaymentStateMachineConfigCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload plugin payment state machine config created response has a 4xx status code
func (o *UploadPluginPaymentStateMachineConfigCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload plugin payment state machine config created response has a 5xx status code
func (o *UploadPluginPaymentStateMachineConfigCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this upload plugin payment state machine config created response a status code equal to that given
func (o *UploadPluginPaymentStateMachineConfigCreated) IsCode(code int) bool {
	return code == 201
}

func (o *UploadPluginPaymentStateMachineConfigCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/uploadPluginPaymentStateMachineConfig/{pluginName}][%d] uploadPluginPaymentStateMachineConfigCreated  %+v", 201, o.Payload)
}

func (o *UploadPluginPaymentStateMachineConfigCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/uploadPluginPaymentStateMachineConfig/{pluginName}][%d] uploadPluginPaymentStateMachineConfigCreated  %+v", 201, o.Payload)
}

func (o *UploadPluginPaymentStateMachineConfigCreated) GetPayload() *kbmodel.TenantKeyValue {
	return o.Payload
}

func (o *UploadPluginPaymentStateMachineConfigCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.TenantKeyValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadPluginPaymentStateMachineConfigBadRequest creates a UploadPluginPaymentStateMachineConfigBadRequest with default headers values
func NewUploadPluginPaymentStateMachineConfigBadRequest() *UploadPluginPaymentStateMachineConfigBadRequest {
	return &UploadPluginPaymentStateMachineConfigBadRequest{}
}

/*
UploadPluginPaymentStateMachineConfigBadRequest describes a response with status code 400, with default header values.

Invalid tenantId supplied
*/
type UploadPluginPaymentStateMachineConfigBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the upload plugin payment state machine config bad request response
func (o *UploadPluginPaymentStateMachineConfigBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this upload plugin payment state machine config bad request response has a 2xx status code
func (o *UploadPluginPaymentStateMachineConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload plugin payment state machine config bad request response has a 3xx status code
func (o *UploadPluginPaymentStateMachineConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload plugin payment state machine config bad request response has a 4xx status code
func (o *UploadPluginPaymentStateMachineConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload plugin payment state machine config bad request response has a 5xx status code
func (o *UploadPluginPaymentStateMachineConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this upload plugin payment state machine config bad request response a status code equal to that given
func (o *UploadPluginPaymentStateMachineConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UploadPluginPaymentStateMachineConfigBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/uploadPluginPaymentStateMachineConfig/{pluginName}][%d] uploadPluginPaymentStateMachineConfigBadRequest ", 400)
}

func (o *UploadPluginPaymentStateMachineConfigBadRequest) String() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/uploadPluginPaymentStateMachineConfig/{pluginName}][%d] uploadPluginPaymentStateMachineConfigBadRequest ", 400)
}

func (o *UploadPluginPaymentStateMachineConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
