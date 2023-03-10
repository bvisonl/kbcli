// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// PutOutOfRotationReader is a Reader for the PutOutOfRotation structure.
type PutOutOfRotationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOutOfRotationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutOutOfRotationNoContent()
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

// NewPutOutOfRotationNoContent creates a PutOutOfRotationNoContent with default headers values
func NewPutOutOfRotationNoContent() *PutOutOfRotationNoContent {
	return &PutOutOfRotationNoContent{}
}

/*
PutOutOfRotationNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type PutOutOfRotationNoContent struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the put out of rotation no content response
func (o *PutOutOfRotationNoContent) Code() int {
	return 204
}

// IsSuccess returns true when this put out of rotation no content response has a 2xx status code
func (o *PutOutOfRotationNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put out of rotation no content response has a 3xx status code
func (o *PutOutOfRotationNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put out of rotation no content response has a 4xx status code
func (o *PutOutOfRotationNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put out of rotation no content response has a 5xx status code
func (o *PutOutOfRotationNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put out of rotation no content response a status code equal to that given
func (o *PutOutOfRotationNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *PutOutOfRotationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/admin/healthcheck][%d] putOutOfRotationNoContent ", 204)
}

func (o *PutOutOfRotationNoContent) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/admin/healthcheck][%d] putOutOfRotationNoContent ", 204)
}

func (o *PutOutOfRotationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
