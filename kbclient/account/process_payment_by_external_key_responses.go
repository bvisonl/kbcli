// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// ProcessPaymentByExternalKeyReader is a Reader for the ProcessPaymentByExternalKey structure.
type ProcessPaymentByExternalKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProcessPaymentByExternalKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewProcessPaymentByExternalKeyCreated()
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

// NewProcessPaymentByExternalKeyCreated creates a ProcessPaymentByExternalKeyCreated with default headers values
func NewProcessPaymentByExternalKeyCreated() *ProcessPaymentByExternalKeyCreated {
	return &ProcessPaymentByExternalKeyCreated{}
}

/*ProcessPaymentByExternalKeyCreated handles this case with default header values.

Payment transaction created successfully
*/
type ProcessPaymentByExternalKeyCreated struct {
	Payload *kbmodel.Payment

	HttpResponse runtime.ClientResponse
}

func (o *ProcessPaymentByExternalKeyCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/payments][%d] processPaymentByExternalKeyCreated  %+v", 201, o.Payload)
}

func (o *ProcessPaymentByExternalKeyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Payment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProcessPaymentByExternalKeyBadRequest creates a ProcessPaymentByExternalKeyBadRequest with default headers values
func NewProcessPaymentByExternalKeyBadRequest() *ProcessPaymentByExternalKeyBadRequest {
	return &ProcessPaymentByExternalKeyBadRequest{}
}

/*ProcessPaymentByExternalKeyBadRequest handles this case with default header values.

Invalid account external key supplied
*/
type ProcessPaymentByExternalKeyBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *ProcessPaymentByExternalKeyBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/payments][%d] processPaymentByExternalKeyBadRequest ", 400)
}

func (o *ProcessPaymentByExternalKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProcessPaymentByExternalKeyPaymentRequired creates a ProcessPaymentByExternalKeyPaymentRequired with default headers values
func NewProcessPaymentByExternalKeyPaymentRequired() *ProcessPaymentByExternalKeyPaymentRequired {
	return &ProcessPaymentByExternalKeyPaymentRequired{}
}

/*ProcessPaymentByExternalKeyPaymentRequired handles this case with default header values.

Transaction declined by gateway
*/
type ProcessPaymentByExternalKeyPaymentRequired struct {
	HttpResponse runtime.ClientResponse
}

func (o *ProcessPaymentByExternalKeyPaymentRequired) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/payments][%d] processPaymentByExternalKeyPaymentRequired ", 402)
}

func (o *ProcessPaymentByExternalKeyPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProcessPaymentByExternalKeyNotFound creates a ProcessPaymentByExternalKeyNotFound with default headers values
func NewProcessPaymentByExternalKeyNotFound() *ProcessPaymentByExternalKeyNotFound {
	return &ProcessPaymentByExternalKeyNotFound{}
}

/*ProcessPaymentByExternalKeyNotFound handles this case with default header values.

Account not found
*/
type ProcessPaymentByExternalKeyNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *ProcessPaymentByExternalKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/payments][%d] processPaymentByExternalKeyNotFound ", 404)
}

func (o *ProcessPaymentByExternalKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProcessPaymentByExternalKeyUnprocessableEntity creates a ProcessPaymentByExternalKeyUnprocessableEntity with default headers values
func NewProcessPaymentByExternalKeyUnprocessableEntity() *ProcessPaymentByExternalKeyUnprocessableEntity {
	return &ProcessPaymentByExternalKeyUnprocessableEntity{}
}

/*ProcessPaymentByExternalKeyUnprocessableEntity handles this case with default header values.

Payment is aborted by a control plugin
*/
type ProcessPaymentByExternalKeyUnprocessableEntity struct {
	HttpResponse runtime.ClientResponse
}

func (o *ProcessPaymentByExternalKeyUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/payments][%d] processPaymentByExternalKeyUnprocessableEntity ", 422)
}

func (o *ProcessPaymentByExternalKeyUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProcessPaymentByExternalKeyBadGateway creates a ProcessPaymentByExternalKeyBadGateway with default headers values
func NewProcessPaymentByExternalKeyBadGateway() *ProcessPaymentByExternalKeyBadGateway {
	return &ProcessPaymentByExternalKeyBadGateway{}
}

/*ProcessPaymentByExternalKeyBadGateway handles this case with default header values.

Failed to submit payment transaction
*/
type ProcessPaymentByExternalKeyBadGateway struct {
	HttpResponse runtime.ClientResponse
}

func (o *ProcessPaymentByExternalKeyBadGateway) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/payments][%d] processPaymentByExternalKeyBadGateway ", 502)
}

func (o *ProcessPaymentByExternalKeyBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProcessPaymentByExternalKeyServiceUnavailable creates a ProcessPaymentByExternalKeyServiceUnavailable with default headers values
func NewProcessPaymentByExternalKeyServiceUnavailable() *ProcessPaymentByExternalKeyServiceUnavailable {
	return &ProcessPaymentByExternalKeyServiceUnavailable{}
}

/*ProcessPaymentByExternalKeyServiceUnavailable handles this case with default header values.

Payment in unknown status, failed to receive gateway response
*/
type ProcessPaymentByExternalKeyServiceUnavailable struct {
	HttpResponse runtime.ClientResponse
}

func (o *ProcessPaymentByExternalKeyServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/payments][%d] processPaymentByExternalKeyServiceUnavailable ", 503)
}

func (o *ProcessPaymentByExternalKeyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProcessPaymentByExternalKeyGatewayTimeout creates a ProcessPaymentByExternalKeyGatewayTimeout with default headers values
func NewProcessPaymentByExternalKeyGatewayTimeout() *ProcessPaymentByExternalKeyGatewayTimeout {
	return &ProcessPaymentByExternalKeyGatewayTimeout{}
}

/*ProcessPaymentByExternalKeyGatewayTimeout handles this case with default header values.

Payment operation timeout
*/
type ProcessPaymentByExternalKeyGatewayTimeout struct {
	HttpResponse runtime.ClientResponse
}

func (o *ProcessPaymentByExternalKeyGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/payments][%d] processPaymentByExternalKeyGatewayTimeout ", 504)
}

func (o *ProcessPaymentByExternalKeyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}