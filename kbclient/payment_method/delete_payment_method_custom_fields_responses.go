// Code generated by go-swagger; DO NOT EDIT.

package payment_method

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// DeletePaymentMethodCustomFieldsReader is a Reader for the DeletePaymentMethodCustomFields structure.
type DeletePaymentMethodCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePaymentMethodCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeletePaymentMethodCustomFieldsNoContent()
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

// NewDeletePaymentMethodCustomFieldsNoContent creates a DeletePaymentMethodCustomFieldsNoContent with default headers values
func NewDeletePaymentMethodCustomFieldsNoContent() *DeletePaymentMethodCustomFieldsNoContent {
	return &DeletePaymentMethodCustomFieldsNoContent{}
}

/*
DeletePaymentMethodCustomFieldsNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeletePaymentMethodCustomFieldsNoContent struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the delete payment method custom fields no content response
func (o *DeletePaymentMethodCustomFieldsNoContent) Code() int {
	return 204
}

// IsSuccess returns true when this delete payment method custom fields no content response has a 2xx status code
func (o *DeletePaymentMethodCustomFieldsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete payment method custom fields no content response has a 3xx status code
func (o *DeletePaymentMethodCustomFieldsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete payment method custom fields no content response has a 4xx status code
func (o *DeletePaymentMethodCustomFieldsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete payment method custom fields no content response has a 5xx status code
func (o *DeletePaymentMethodCustomFieldsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete payment method custom fields no content response a status code equal to that given
func (o *DeletePaymentMethodCustomFieldsNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeletePaymentMethodCustomFieldsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/paymentMethods/{paymentMethodId}/customFields][%d] deletePaymentMethodCustomFieldsNoContent ", 204)
}

func (o *DeletePaymentMethodCustomFieldsNoContent) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/paymentMethods/{paymentMethodId}/customFields][%d] deletePaymentMethodCustomFieldsNoContent ", 204)
}

func (o *DeletePaymentMethodCustomFieldsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePaymentMethodCustomFieldsBadRequest creates a DeletePaymentMethodCustomFieldsBadRequest with default headers values
func NewDeletePaymentMethodCustomFieldsBadRequest() *DeletePaymentMethodCustomFieldsBadRequest {
	return &DeletePaymentMethodCustomFieldsBadRequest{}
}

/*
DeletePaymentMethodCustomFieldsBadRequest describes a response with status code 400, with default header values.

Invalid payment method id supplied
*/
type DeletePaymentMethodCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the delete payment method custom fields bad request response
func (o *DeletePaymentMethodCustomFieldsBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this delete payment method custom fields bad request response has a 2xx status code
func (o *DeletePaymentMethodCustomFieldsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete payment method custom fields bad request response has a 3xx status code
func (o *DeletePaymentMethodCustomFieldsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete payment method custom fields bad request response has a 4xx status code
func (o *DeletePaymentMethodCustomFieldsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete payment method custom fields bad request response has a 5xx status code
func (o *DeletePaymentMethodCustomFieldsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete payment method custom fields bad request response a status code equal to that given
func (o *DeletePaymentMethodCustomFieldsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeletePaymentMethodCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/paymentMethods/{paymentMethodId}/customFields][%d] deletePaymentMethodCustomFieldsBadRequest ", 400)
}

func (o *DeletePaymentMethodCustomFieldsBadRequest) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/paymentMethods/{paymentMethodId}/customFields][%d] deletePaymentMethodCustomFieldsBadRequest ", 400)
}

func (o *DeletePaymentMethodCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
