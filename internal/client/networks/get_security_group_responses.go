// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// GetSecurityGroupReader is a Reader for the GetSecurityGroup structure.
type GetSecurityGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecurityGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSecurityGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSecurityGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSecurityGroupOK creates a GetSecurityGroupOK with default headers values
func NewGetSecurityGroupOK() *GetSecurityGroupOK {
	return &GetSecurityGroupOK{}
}

/*
	GetSecurityGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetSecurityGroupOK struct {
	Payload *models.GetSecurityGroupResponse
}

// IsSuccess returns true when this get security group o k response has a 2xx status code
func (o *GetSecurityGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get security group o k response has a 3xx status code
func (o *GetSecurityGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get security group o k response has a 4xx status code
func (o *GetSecurityGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get security group o k response has a 5xx status code
func (o *GetSecurityGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get security group o k response a status code equal to that given
func (o *GetSecurityGroupOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSecurityGroupOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/networks/security-groups/{id}][%d] getSecurityGroupOK  %+v", 200, o.Payload)
}

func (o *GetSecurityGroupOK) String() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/networks/security-groups/{id}][%d] getSecurityGroupOK  %+v", 200, o.Payload)
}

func (o *GetSecurityGroupOK) GetPayload() *models.GetSecurityGroupResponse {
	return o.Payload
}

func (o *GetSecurityGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetSecurityGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecurityGroupDefault creates a GetSecurityGroupDefault with default headers values
func NewGetSecurityGroupDefault(code int) *GetSecurityGroupDefault {
	return &GetSecurityGroupDefault{
		_statusCode: code,
	}
}

/*
	GetSecurityGroupDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetSecurityGroupDefault struct {
	_statusCode int

	Payload *models.Status
}

// Code gets the status code for the get security group default response
func (o *GetSecurityGroupDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get security group default response has a 2xx status code
func (o *GetSecurityGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get security group default response has a 3xx status code
func (o *GetSecurityGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get security group default response has a 4xx status code
func (o *GetSecurityGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get security group default response has a 5xx status code
func (o *GetSecurityGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get security group default response a status code equal to that given
func (o *GetSecurityGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetSecurityGroupDefault) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/networks/security-groups/{id}][%d] GetSecurityGroup default  %+v", o._statusCode, o.Payload)
}

func (o *GetSecurityGroupDefault) String() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/networks/security-groups/{id}][%d] GetSecurityGroup default  %+v", o._statusCode, o.Payload)
}

func (o *GetSecurityGroupDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *GetSecurityGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
