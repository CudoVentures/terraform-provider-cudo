// Code generated by go-swagger; DO NOT EDIT.

package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// ListUserPermissionsReader is a Reader for the ListUserPermissions structure.
type ListUserPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUserPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListUserPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListUserPermissionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListUserPermissionsOK creates a ListUserPermissionsOK with default headers values
func NewListUserPermissionsOK() *ListUserPermissionsOK {
	return &ListUserPermissionsOK{}
}

/*
	ListUserPermissionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListUserPermissionsOK struct {
	Payload *models.ListUserPermissionsResponse
}

// IsSuccess returns true when this list user permissions o k response has a 2xx status code
func (o *ListUserPermissionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list user permissions o k response has a 3xx status code
func (o *ListUserPermissionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list user permissions o k response has a 4xx status code
func (o *ListUserPermissionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list user permissions o k response has a 5xx status code
func (o *ListUserPermissionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list user permissions o k response a status code equal to that given
func (o *ListUserPermissionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListUserPermissionsOK) Error() string {
	return fmt.Sprintf("[GET /v1/auth/permissions][%d] listUserPermissionsOK  %+v", 200, o.Payload)
}

func (o *ListUserPermissionsOK) String() string {
	return fmt.Sprintf("[GET /v1/auth/permissions][%d] listUserPermissionsOK  %+v", 200, o.Payload)
}

func (o *ListUserPermissionsOK) GetPayload() *models.ListUserPermissionsResponse {
	return o.Payload
}

func (o *ListUserPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListUserPermissionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserPermissionsDefault creates a ListUserPermissionsDefault with default headers values
func NewListUserPermissionsDefault(code int) *ListUserPermissionsDefault {
	return &ListUserPermissionsDefault{
		_statusCode: code,
	}
}

/*
	ListUserPermissionsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListUserPermissionsDefault struct {
	_statusCode int

	Payload *models.Status
}

// Code gets the status code for the list user permissions default response
func (o *ListUserPermissionsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list user permissions default response has a 2xx status code
func (o *ListUserPermissionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list user permissions default response has a 3xx status code
func (o *ListUserPermissionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list user permissions default response has a 4xx status code
func (o *ListUserPermissionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list user permissions default response has a 5xx status code
func (o *ListUserPermissionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list user permissions default response a status code equal to that given
func (o *ListUserPermissionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListUserPermissionsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/auth/permissions][%d] ListUserPermissions default  %+v", o._statusCode, o.Payload)
}

func (o *ListUserPermissionsDefault) String() string {
	return fmt.Sprintf("[GET /v1/auth/permissions][%d] ListUserPermissions default  %+v", o._statusCode, o.Payload)
}

func (o *ListUserPermissionsDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *ListUserPermissionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
