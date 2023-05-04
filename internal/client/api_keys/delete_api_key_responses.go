// Code generated by go-swagger; DO NOT EDIT.

package api_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// DeleteAPIKeyReader is a Reader for the DeleteAPIKey structure.
type DeleteAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAPIKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteAPIKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAPIKeyOK creates a DeleteAPIKeyOK with default headers values
func NewDeleteAPIKeyOK() *DeleteAPIKeyOK {
	return &DeleteAPIKeyOK{}
}

/*
	DeleteAPIKeyOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteAPIKeyOK struct {
	Payload interface{}
}

// IsSuccess returns true when this delete Api key o k response has a 2xx status code
func (o *DeleteAPIKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Api key o k response has a 3xx status code
func (o *DeleteAPIKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api key o k response has a 4xx status code
func (o *DeleteAPIKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Api key o k response has a 5xx status code
func (o *DeleteAPIKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api key o k response a status code equal to that given
func (o *DeleteAPIKeyOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteAPIKeyOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/api-keys/{name}][%d] deleteApiKeyOK  %+v", 200, o.Payload)
}

func (o *DeleteAPIKeyOK) String() string {
	return fmt.Sprintf("[DELETE /v1/api-keys/{name}][%d] deleteApiKeyOK  %+v", 200, o.Payload)
}

func (o *DeleteAPIKeyOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteAPIKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIKeyDefault creates a DeleteAPIKeyDefault with default headers values
func NewDeleteAPIKeyDefault(code int) *DeleteAPIKeyDefault {
	return &DeleteAPIKeyDefault{
		_statusCode: code,
	}
}

/*
	DeleteAPIKeyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteAPIKeyDefault struct {
	_statusCode int

	Payload *models.Status
}

// Code gets the status code for the delete Api key default response
func (o *DeleteAPIKeyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete Api key default response has a 2xx status code
func (o *DeleteAPIKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete Api key default response has a 3xx status code
func (o *DeleteAPIKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete Api key default response has a 4xx status code
func (o *DeleteAPIKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete Api key default response has a 5xx status code
func (o *DeleteAPIKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete Api key default response a status code equal to that given
func (o *DeleteAPIKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteAPIKeyDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/api-keys/{name}][%d] DeleteApiKey default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAPIKeyDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/api-keys/{name}][%d] DeleteApiKey default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAPIKeyDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *DeleteAPIKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
