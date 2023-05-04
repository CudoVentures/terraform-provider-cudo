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

// StopInstanceReader is a Reader for the StopInstance structure.
type StopInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStopInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStopInstanceOK creates a StopInstanceOK with default headers values
func NewStopInstanceOK() *StopInstanceOK {
	return &StopInstanceOK{}
}

/*
	StopInstanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type StopInstanceOK struct {
	Payload models.StopInstanceResponse
}

// IsSuccess returns true when this stop instance o k response has a 2xx status code
func (o *StopInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop instance o k response has a 3xx status code
func (o *StopInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop instance o k response has a 4xx status code
func (o *StopInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop instance o k response has a 5xx status code
func (o *StopInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stop instance o k response a status code equal to that given
func (o *StopInstanceOK) IsCode(code int) bool {
	return code == 200
}

func (o *StopInstanceOK) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/instances/{instanceId}/stop][%d] stopInstanceOK  %+v", 200, o.Payload)
}

func (o *StopInstanceOK) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/instances/{instanceId}/stop][%d] stopInstanceOK  %+v", 200, o.Payload)
}

func (o *StopInstanceOK) GetPayload() models.StopInstanceResponse {
	return o.Payload
}

func (o *StopInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopInstanceDefault creates a StopInstanceDefault with default headers values
func NewStopInstanceDefault(code int) *StopInstanceDefault {
	return &StopInstanceDefault{
		_statusCode: code,
	}
}

/*
	StopInstanceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StopInstanceDefault struct {
	_statusCode int

	Payload *models.Status
}

// Code gets the status code for the stop instance default response
func (o *StopInstanceDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this stop instance default response has a 2xx status code
func (o *StopInstanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this stop instance default response has a 3xx status code
func (o *StopInstanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this stop instance default response has a 4xx status code
func (o *StopInstanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this stop instance default response has a 5xx status code
func (o *StopInstanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this stop instance default response a status code equal to that given
func (o *StopInstanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *StopInstanceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/instances/{instanceId}/stop][%d] StopInstance default  %+v", o._statusCode, o.Payload)
}

func (o *StopInstanceDefault) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/instances/{instanceId}/stop][%d] StopInstance default  %+v", o._statusCode, o.Payload)
}

func (o *StopInstanceDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *StopInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
