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
)

// UploadInvoiceMPTemplateReader is a Reader for the UploadInvoiceMPTemplate structure.
type UploadInvoiceMPTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadInvoiceMPTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUploadInvoiceMPTemplateOK()
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

// NewUploadInvoiceMPTemplateOK creates a UploadInvoiceMPTemplateOK with default headers values
func NewUploadInvoiceMPTemplateOK() *UploadInvoiceMPTemplateOK {
	return &UploadInvoiceMPTemplateOK{}
}

/*
UploadInvoiceMPTemplateOK describes a response with status code 200, with default header values.

successful operation
*/
type UploadInvoiceMPTemplateOK struct {
	Payload      string
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the upload invoice m p template o k response
func (o *UploadInvoiceMPTemplateOK) Code() int {
	return 200
}

// IsSuccess returns true when this upload invoice m p template o k response has a 2xx status code
func (o *UploadInvoiceMPTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload invoice m p template o k response has a 3xx status code
func (o *UploadInvoiceMPTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload invoice m p template o k response has a 4xx status code
func (o *UploadInvoiceMPTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload invoice m p template o k response has a 5xx status code
func (o *UploadInvoiceMPTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this upload invoice m p template o k response a status code equal to that given
func (o *UploadInvoiceMPTemplateOK) IsCode(code int) bool {
	return code == 200
}

func (o *UploadInvoiceMPTemplateOK) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/manualPayTemplate][%d] uploadInvoiceMPTemplateOK  %+v", 200, o.Payload)
}

func (o *UploadInvoiceMPTemplateOK) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/manualPayTemplate][%d] uploadInvoiceMPTemplateOK  %+v", 200, o.Payload)
}

func (o *UploadInvoiceMPTemplateOK) GetPayload() string {
	return o.Payload
}

func (o *UploadInvoiceMPTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
