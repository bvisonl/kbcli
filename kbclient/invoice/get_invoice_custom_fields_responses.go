// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// GetInvoiceCustomFieldsReader is a Reader for the GetInvoiceCustomFields structure.
type GetInvoiceCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoiceCustomFieldsOK()
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

// NewGetInvoiceCustomFieldsOK creates a GetInvoiceCustomFieldsOK with default headers values
func NewGetInvoiceCustomFieldsOK() *GetInvoiceCustomFieldsOK {
	return &GetInvoiceCustomFieldsOK{}
}

/*
GetInvoiceCustomFieldsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetInvoiceCustomFieldsOK struct {
	Payload      []*kbmodel.CustomField
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get invoice custom fields o k response
func (o *GetInvoiceCustomFieldsOK) Code() int {
	return 200
}

// IsSuccess returns true when this get invoice custom fields o k response has a 2xx status code
func (o *GetInvoiceCustomFieldsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get invoice custom fields o k response has a 3xx status code
func (o *GetInvoiceCustomFieldsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice custom fields o k response has a 4xx status code
func (o *GetInvoiceCustomFieldsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get invoice custom fields o k response has a 5xx status code
func (o *GetInvoiceCustomFieldsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice custom fields o k response a status code equal to that given
func (o *GetInvoiceCustomFieldsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetInvoiceCustomFieldsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceId}/customFields][%d] getInvoiceCustomFieldsOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceCustomFieldsOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceId}/customFields][%d] getInvoiceCustomFieldsOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceCustomFieldsOK) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *GetInvoiceCustomFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceCustomFieldsBadRequest creates a GetInvoiceCustomFieldsBadRequest with default headers values
func NewGetInvoiceCustomFieldsBadRequest() *GetInvoiceCustomFieldsBadRequest {
	return &GetInvoiceCustomFieldsBadRequest{}
}

/*
GetInvoiceCustomFieldsBadRequest describes a response with status code 400, with default header values.

Invalid invoice id supplied
*/
type GetInvoiceCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get invoice custom fields bad request response
func (o *GetInvoiceCustomFieldsBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this get invoice custom fields bad request response has a 2xx status code
func (o *GetInvoiceCustomFieldsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoice custom fields bad request response has a 3xx status code
func (o *GetInvoiceCustomFieldsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice custom fields bad request response has a 4xx status code
func (o *GetInvoiceCustomFieldsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoice custom fields bad request response has a 5xx status code
func (o *GetInvoiceCustomFieldsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice custom fields bad request response a status code equal to that given
func (o *GetInvoiceCustomFieldsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetInvoiceCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceId}/customFields][%d] getInvoiceCustomFieldsBadRequest ", 400)
}

func (o *GetInvoiceCustomFieldsBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceId}/customFields][%d] getInvoiceCustomFieldsBadRequest ", 400)
}

func (o *GetInvoiceCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
