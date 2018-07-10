// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// ModifyInvoiceCustomFieldsReader is a Reader for the ModifyInvoiceCustomFields structure.
type ModifyInvoiceCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyInvoiceCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewModifyInvoiceCustomFieldsNoContent()
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

// NewModifyInvoiceCustomFieldsNoContent creates a ModifyInvoiceCustomFieldsNoContent with default headers values
func NewModifyInvoiceCustomFieldsNoContent() *ModifyInvoiceCustomFieldsNoContent {
	return &ModifyInvoiceCustomFieldsNoContent{}
}

/*ModifyInvoiceCustomFieldsNoContent handles this case with default header values.

Successful operation
*/
type ModifyInvoiceCustomFieldsNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *ModifyInvoiceCustomFieldsNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoices/{invoiceId}/customFields][%d] modifyInvoiceCustomFieldsNoContent ", 204)
}

func (o *ModifyInvoiceCustomFieldsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyInvoiceCustomFieldsBadRequest creates a ModifyInvoiceCustomFieldsBadRequest with default headers values
func NewModifyInvoiceCustomFieldsBadRequest() *ModifyInvoiceCustomFieldsBadRequest {
	return &ModifyInvoiceCustomFieldsBadRequest{}
}

/*ModifyInvoiceCustomFieldsBadRequest handles this case with default header values.

Invalid invoice id supplied
*/
type ModifyInvoiceCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *ModifyInvoiceCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoices/{invoiceId}/customFields][%d] modifyInvoiceCustomFieldsBadRequest ", 400)
}

func (o *ModifyInvoiceCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}