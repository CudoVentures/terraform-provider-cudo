// Code generated by go-swagger; DO NOT EDIT.

package virtual_machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// StartInstanceReader is a Reader for the StartInstance structure.
type StartInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartInstanceOK creates a StartInstanceOK with default headers values
func NewStartInstanceOK() *StartInstanceOK {
	return &StartInstanceOK{}
}

/*
	StartInstanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type StartInstanceOK struct {
	Payload models.StartInstanceResponse
}

// IsSuccess returns true when this start instance o k response has a 2xx status code
func (o *StartInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start instance o k response has a 3xx status code
func (o *StartInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start instance o k response has a 4xx status code
func (o *StartInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this start instance o k response has a 5xx status code
func (o *StartInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this start instance o k response a status code equal to that given
func (o *StartInstanceOK) IsCode(code int) bool {
	return code == 200
}

func (o *StartInstanceOK) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/instances/{instanceId}/start][%d] startInstanceOK  %+v", 200, o.Payload)
}

func (o *StartInstanceOK) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/instances/{instanceId}/start][%d] startInstanceOK  %+v", 200, o.Payload)
}

func (o *StartInstanceOK) GetPayload() models.StartInstanceResponse {
	return o.Payload
}

func (o *StartInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartInstanceDefault creates a StartInstanceDefault with default headers values
func NewStartInstanceDefault(code int) *StartInstanceDefault {
	return &StartInstanceDefault{
		_statusCode: code,
	}
}

/*
	StartInstanceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StartInstanceDefault struct {
	_statusCode int

	Payload *models.Status
}

// Code gets the status code for the start instance default response
func (o *StartInstanceDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this start instance default response has a 2xx status code
func (o *StartInstanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this start instance default response has a 3xx status code
func (o *StartInstanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this start instance default response has a 4xx status code
func (o *StartInstanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this start instance default response has a 5xx status code
func (o *StartInstanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this start instance default response a status code equal to that given
func (o *StartInstanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *StartInstanceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/instances/{instanceId}/start][%d] StartInstance default  %+v", o._statusCode, o.Payload)
}

func (o *StartInstanceDefault) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/instances/{instanceId}/start][%d] StartInstance default  %+v", o._statusCode, o.Payload)
}

func (o *StartInstanceDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *StartInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
