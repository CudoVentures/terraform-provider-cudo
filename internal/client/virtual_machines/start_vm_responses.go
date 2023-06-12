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

// StartVMReader is a Reader for the StartVM structure.
type StartVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartVMDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartVMOK creates a StartVMOK with default headers values
func NewStartVMOK() *StartVMOK {
	return &StartVMOK{}
}

/*
StartVMOK describes a response with status code 200, with default header values.

A successful response.
*/
type StartVMOK struct {
	Payload models.StartVMResponse
}

// IsSuccess returns true when this start Vm o k response has a 2xx status code
func (o *StartVMOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start Vm o k response has a 3xx status code
func (o *StartVMOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start Vm o k response has a 4xx status code
func (o *StartVMOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this start Vm o k response has a 5xx status code
func (o *StartVMOK) IsServerError() bool {
	return false
}

// IsCode returns true when this start Vm o k response a status code equal to that given
func (o *StartVMOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the start Vm o k response
func (o *StartVMOK) Code() int {
	return 200
}

func (o *StartVMOK) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/vms/{id}/start][%d] startVmOK  %+v", 200, o.Payload)
}

func (o *StartVMOK) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/vms/{id}/start][%d] startVmOK  %+v", 200, o.Payload)
}

func (o *StartVMOK) GetPayload() models.StartVMResponse {
	return o.Payload
}

func (o *StartVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartVMDefault creates a StartVMDefault with default headers values
func NewStartVMDefault(code int) *StartVMDefault {
	return &StartVMDefault{
		_statusCode: code,
	}
}

/*
StartVMDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StartVMDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this start VM default response has a 2xx status code
func (o *StartVMDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this start VM default response has a 3xx status code
func (o *StartVMDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this start VM default response has a 4xx status code
func (o *StartVMDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this start VM default response has a 5xx status code
func (o *StartVMDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this start VM default response a status code equal to that given
func (o *StartVMDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the start VM default response
func (o *StartVMDefault) Code() int {
	return o._statusCode
}

func (o *StartVMDefault) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/vms/{id}/start][%d] StartVM default  %+v", o._statusCode, o.Payload)
}

func (o *StartVMDefault) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/vms/{id}/start][%d] StartVM default  %+v", o._statusCode, o.Payload)
}

func (o *StartVMDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *StartVMDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}