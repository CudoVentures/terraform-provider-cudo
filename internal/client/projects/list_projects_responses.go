// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// ListProjectsReader is a Reader for the ListProjects structure.
type ListProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectsOK creates a ListProjectsOK with default headers values
func NewListProjectsOK() *ListProjectsOK {
	return &ListProjectsOK{}
}

/*
	ListProjectsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListProjectsOK struct {
	Payload *models.ListProjectsResponse
}

// IsSuccess returns true when this list projects o k response has a 2xx status code
func (o *ListProjectsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list projects o k response has a 3xx status code
func (o *ListProjectsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list projects o k response has a 4xx status code
func (o *ListProjectsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list projects o k response has a 5xx status code
func (o *ListProjectsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list projects o k response a status code equal to that given
func (o *ListProjectsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListProjectsOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects][%d] listProjectsOK  %+v", 200, o.Payload)
}

func (o *ListProjectsOK) String() string {
	return fmt.Sprintf("[GET /v1/projects][%d] listProjectsOK  %+v", 200, o.Payload)
}

func (o *ListProjectsOK) GetPayload() *models.ListProjectsResponse {
	return o.Payload
}

func (o *ListProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListProjectsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectsDefault creates a ListProjectsDefault with default headers values
func NewListProjectsDefault(code int) *ListProjectsDefault {
	return &ListProjectsDefault{
		_statusCode: code,
	}
}

/*
	ListProjectsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListProjectsDefault struct {
	_statusCode int

	Payload *models.Status
}

// Code gets the status code for the list projects default response
func (o *ListProjectsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list projects default response has a 2xx status code
func (o *ListProjectsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list projects default response has a 3xx status code
func (o *ListProjectsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list projects default response has a 4xx status code
func (o *ListProjectsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list projects default response has a 5xx status code
func (o *ListProjectsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list projects default response a status code equal to that given
func (o *ListProjectsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListProjectsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/projects][%d] ListProjects default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectsDefault) String() string {
	return fmt.Sprintf("[GET /v1/projects][%d] ListProjects default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectsDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *ListProjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
