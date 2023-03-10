// Code generated by go-swagger; DO NOT EDIT.

package export

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// ExportDataForAccountReader is a Reader for the ExportDataForAccount structure.
type ExportDataForAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportDataForAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewExportDataForAccountOK()
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

// NewExportDataForAccountOK creates a ExportDataForAccountOK with default headers values
func NewExportDataForAccountOK() *ExportDataForAccountOK {
	return &ExportDataForAccountOK{}
}

/*
ExportDataForAccountOK describes a response with status code 200, with default header values.

Success
*/
type ExportDataForAccountOK struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the export data for account o k response
func (o *ExportDataForAccountOK) Code() int {
	return 200
}

// IsSuccess returns true when this export data for account o k response has a 2xx status code
func (o *ExportDataForAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export data for account o k response has a 3xx status code
func (o *ExportDataForAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export data for account o k response has a 4xx status code
func (o *ExportDataForAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this export data for account o k response has a 5xx status code
func (o *ExportDataForAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this export data for account o k response a status code equal to that given
func (o *ExportDataForAccountOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExportDataForAccountOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/export/{accountId}][%d] exportDataForAccountOK ", 200)
}

func (o *ExportDataForAccountOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/export/{accountId}][%d] exportDataForAccountOK ", 200)
}

func (o *ExportDataForAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportDataForAccountBadRequest creates a ExportDataForAccountBadRequest with default headers values
func NewExportDataForAccountBadRequest() *ExportDataForAccountBadRequest {
	return &ExportDataForAccountBadRequest{}
}

/*
ExportDataForAccountBadRequest describes a response with status code 400, with default header values.

Invalid account id supplied
*/
type ExportDataForAccountBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the export data for account bad request response
func (o *ExportDataForAccountBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this export data for account bad request response has a 2xx status code
func (o *ExportDataForAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export data for account bad request response has a 3xx status code
func (o *ExportDataForAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export data for account bad request response has a 4xx status code
func (o *ExportDataForAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this export data for account bad request response has a 5xx status code
func (o *ExportDataForAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this export data for account bad request response a status code equal to that given
func (o *ExportDataForAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ExportDataForAccountBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/export/{accountId}][%d] exportDataForAccountBadRequest ", 400)
}

func (o *ExportDataForAccountBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/export/{accountId}][%d] exportDataForAccountBadRequest ", 400)
}

func (o *ExportDataForAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportDataForAccountNotFound creates a ExportDataForAccountNotFound with default headers values
func NewExportDataForAccountNotFound() *ExportDataForAccountNotFound {
	return &ExportDataForAccountNotFound{}
}

/*
ExportDataForAccountNotFound describes a response with status code 404, with default header values.

Account not found
*/
type ExportDataForAccountNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the export data for account not found response
func (o *ExportDataForAccountNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this export data for account not found response has a 2xx status code
func (o *ExportDataForAccountNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export data for account not found response has a 3xx status code
func (o *ExportDataForAccountNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export data for account not found response has a 4xx status code
func (o *ExportDataForAccountNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this export data for account not found response has a 5xx status code
func (o *ExportDataForAccountNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this export data for account not found response a status code equal to that given
func (o *ExportDataForAccountNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ExportDataForAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/export/{accountId}][%d] exportDataForAccountNotFound ", 404)
}

func (o *ExportDataForAccountNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/export/{accountId}][%d] exportDataForAccountNotFound ", 404)
}

func (o *ExportDataForAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
