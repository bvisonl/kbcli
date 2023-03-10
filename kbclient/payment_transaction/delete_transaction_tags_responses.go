// Code generated by go-swagger; DO NOT EDIT.

package payment_transaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// DeleteTransactionTagsReader is a Reader for the DeleteTransactionTags structure.
type DeleteTransactionTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTransactionTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteTransactionTagsNoContent()
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

// NewDeleteTransactionTagsNoContent creates a DeleteTransactionTagsNoContent with default headers values
func NewDeleteTransactionTagsNoContent() *DeleteTransactionTagsNoContent {
	return &DeleteTransactionTagsNoContent{}
}

/*
DeleteTransactionTagsNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteTransactionTagsNoContent struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the delete transaction tags no content response
func (o *DeleteTransactionTagsNoContent) Code() int {
	return 204
}

// IsSuccess returns true when this delete transaction tags no content response has a 2xx status code
func (o *DeleteTransactionTagsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete transaction tags no content response has a 3xx status code
func (o *DeleteTransactionTagsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete transaction tags no content response has a 4xx status code
func (o *DeleteTransactionTagsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete transaction tags no content response has a 5xx status code
func (o *DeleteTransactionTagsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete transaction tags no content response a status code equal to that given
func (o *DeleteTransactionTagsNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteTransactionTagsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/paymentTransactions/{transactionId}/tags][%d] deleteTransactionTagsNoContent ", 204)
}

func (o *DeleteTransactionTagsNoContent) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/paymentTransactions/{transactionId}/tags][%d] deleteTransactionTagsNoContent ", 204)
}

func (o *DeleteTransactionTagsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTransactionTagsBadRequest creates a DeleteTransactionTagsBadRequest with default headers values
func NewDeleteTransactionTagsBadRequest() *DeleteTransactionTagsBadRequest {
	return &DeleteTransactionTagsBadRequest{}
}

/*
DeleteTransactionTagsBadRequest describes a response with status code 400, with default header values.

Invalid transaction id supplied
*/
type DeleteTransactionTagsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the delete transaction tags bad request response
func (o *DeleteTransactionTagsBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this delete transaction tags bad request response has a 2xx status code
func (o *DeleteTransactionTagsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete transaction tags bad request response has a 3xx status code
func (o *DeleteTransactionTagsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete transaction tags bad request response has a 4xx status code
func (o *DeleteTransactionTagsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete transaction tags bad request response has a 5xx status code
func (o *DeleteTransactionTagsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete transaction tags bad request response a status code equal to that given
func (o *DeleteTransactionTagsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteTransactionTagsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/paymentTransactions/{transactionId}/tags][%d] deleteTransactionTagsBadRequest ", 400)
}

func (o *DeleteTransactionTagsBadRequest) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/paymentTransactions/{transactionId}/tags][%d] deleteTransactionTagsBadRequest ", 400)
}

func (o *DeleteTransactionTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
