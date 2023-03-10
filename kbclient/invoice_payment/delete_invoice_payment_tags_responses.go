// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// DeleteInvoicePaymentTagsReader is a Reader for the DeleteInvoicePaymentTags structure.
type DeleteInvoicePaymentTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInvoicePaymentTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteInvoicePaymentTagsNoContent()
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

// NewDeleteInvoicePaymentTagsNoContent creates a DeleteInvoicePaymentTagsNoContent with default headers values
func NewDeleteInvoicePaymentTagsNoContent() *DeleteInvoicePaymentTagsNoContent {
	return &DeleteInvoicePaymentTagsNoContent{}
}

/*
DeleteInvoicePaymentTagsNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteInvoicePaymentTagsNoContent struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the delete invoice payment tags no content response
func (o *DeleteInvoicePaymentTagsNoContent) Code() int {
	return 204
}

// IsSuccess returns true when this delete invoice payment tags no content response has a 2xx status code
func (o *DeleteInvoicePaymentTagsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete invoice payment tags no content response has a 3xx status code
func (o *DeleteInvoicePaymentTagsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete invoice payment tags no content response has a 4xx status code
func (o *DeleteInvoicePaymentTagsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete invoice payment tags no content response has a 5xx status code
func (o *DeleteInvoicePaymentTagsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete invoice payment tags no content response a status code equal to that given
func (o *DeleteInvoicePaymentTagsNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteInvoicePaymentTagsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/invoicePayments/{paymentId}/tags][%d] deleteInvoicePaymentTagsNoContent ", 204)
}

func (o *DeleteInvoicePaymentTagsNoContent) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/invoicePayments/{paymentId}/tags][%d] deleteInvoicePaymentTagsNoContent ", 204)
}

func (o *DeleteInvoicePaymentTagsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInvoicePaymentTagsBadRequest creates a DeleteInvoicePaymentTagsBadRequest with default headers values
func NewDeleteInvoicePaymentTagsBadRequest() *DeleteInvoicePaymentTagsBadRequest {
	return &DeleteInvoicePaymentTagsBadRequest{}
}

/*
DeleteInvoicePaymentTagsBadRequest describes a response with status code 400, with default header values.

Invalid payment id supplied
*/
type DeleteInvoicePaymentTagsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the delete invoice payment tags bad request response
func (o *DeleteInvoicePaymentTagsBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this delete invoice payment tags bad request response has a 2xx status code
func (o *DeleteInvoicePaymentTagsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete invoice payment tags bad request response has a 3xx status code
func (o *DeleteInvoicePaymentTagsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete invoice payment tags bad request response has a 4xx status code
func (o *DeleteInvoicePaymentTagsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete invoice payment tags bad request response has a 5xx status code
func (o *DeleteInvoicePaymentTagsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete invoice payment tags bad request response a status code equal to that given
func (o *DeleteInvoicePaymentTagsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteInvoicePaymentTagsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/invoicePayments/{paymentId}/tags][%d] deleteInvoicePaymentTagsBadRequest ", 400)
}

func (o *DeleteInvoicePaymentTagsBadRequest) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/invoicePayments/{paymentId}/tags][%d] deleteInvoicePaymentTagsBadRequest ", 400)
}

func (o *DeleteInvoicePaymentTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
