// Code generated by go-swagger; DO NOT EDIT.

package data_centers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// CountHostsReader is a Reader for the CountHosts structure.
type CountHostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CountHostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCountHostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCountHostsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCountHostsOK creates a CountHostsOK with default headers values
func NewCountHostsOK() *CountHostsOK {
	return &CountHostsOK{}
}

/*
CountHostsOK describes a response with status code 200, with default header values.

A successful response.
*/
type CountHostsOK struct {
	Payload *models.CountHostsResponse
}

// IsSuccess returns true when this count hosts o k response has a 2xx status code
func (o *CountHostsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this count hosts o k response has a 3xx status code
func (o *CountHostsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this count hosts o k response has a 4xx status code
func (o *CountHostsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this count hosts o k response has a 5xx status code
func (o *CountHostsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this count hosts o k response a status code equal to that given
func (o *CountHostsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the count hosts o k response
func (o *CountHostsOK) Code() int {
	return 200
}

func (o *CountHostsOK) Error() string {
	return fmt.Sprintf("[GET /v1/data-centers/{dataCenterId}/host-count][%d] countHostsOK  %+v", 200, o.Payload)
}

func (o *CountHostsOK) String() string {
	return fmt.Sprintf("[GET /v1/data-centers/{dataCenterId}/host-count][%d] countHostsOK  %+v", 200, o.Payload)
}

func (o *CountHostsOK) GetPayload() *models.CountHostsResponse {
	return o.Payload
}

func (o *CountHostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CountHostsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCountHostsDefault creates a CountHostsDefault with default headers values
func NewCountHostsDefault(code int) *CountHostsDefault {
	return &CountHostsDefault{
		_statusCode: code,
	}
}

/*
CountHostsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CountHostsDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this count hosts default response has a 2xx status code
func (o *CountHostsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this count hosts default response has a 3xx status code
func (o *CountHostsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this count hosts default response has a 4xx status code
func (o *CountHostsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this count hosts default response has a 5xx status code
func (o *CountHostsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this count hosts default response a status code equal to that given
func (o *CountHostsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the count hosts default response
func (o *CountHostsDefault) Code() int {
	return o._statusCode
}

func (o *CountHostsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/data-centers/{dataCenterId}/host-count][%d] CountHosts default  %+v", o._statusCode, o.Payload)
}

func (o *CountHostsDefault) String() string {
	return fmt.Sprintf("[GET /v1/data-centers/{dataCenterId}/host-count][%d] CountHosts default  %+v", o._statusCode, o.Payload)
}

func (o *CountHostsDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *CountHostsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
