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

// GetProjectSpendDetailsReader is a Reader for the GetProjectSpendDetails structure.
type GetProjectSpendDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectSpendDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectSpendDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetProjectSpendDetailsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProjectSpendDetailsOK creates a GetProjectSpendDetailsOK with default headers values
func NewGetProjectSpendDetailsOK() *GetProjectSpendDetailsOK {
	return &GetProjectSpendDetailsOK{}
}

/*
	GetProjectSpendDetailsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetProjectSpendDetailsOK struct {
	Payload *models.GetProjectSpendDetailsResponse
}

// IsSuccess returns true when this get project spend details o k response has a 2xx status code
func (o *GetProjectSpendDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get project spend details o k response has a 3xx status code
func (o *GetProjectSpendDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project spend details o k response has a 4xx status code
func (o *GetProjectSpendDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project spend details o k response has a 5xx status code
func (o *GetProjectSpendDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get project spend details o k response a status code equal to that given
func (o *GetProjectSpendDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetProjectSpendDetailsOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/spend/{spendId}/details][%d] getProjectSpendDetailsOK  %+v", 200, o.Payload)
}

func (o *GetProjectSpendDetailsOK) String() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/spend/{spendId}/details][%d] getProjectSpendDetailsOK  %+v", 200, o.Payload)
}

func (o *GetProjectSpendDetailsOK) GetPayload() *models.GetProjectSpendDetailsResponse {
	return o.Payload
}

func (o *GetProjectSpendDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetProjectSpendDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectSpendDetailsDefault creates a GetProjectSpendDetailsDefault with default headers values
func NewGetProjectSpendDetailsDefault(code int) *GetProjectSpendDetailsDefault {
	return &GetProjectSpendDetailsDefault{
		_statusCode: code,
	}
}

/*
	GetProjectSpendDetailsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetProjectSpendDetailsDefault struct {
	_statusCode int

	Payload *models.Status
}

// Code gets the status code for the get project spend details default response
func (o *GetProjectSpendDetailsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get project spend details default response has a 2xx status code
func (o *GetProjectSpendDetailsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get project spend details default response has a 3xx status code
func (o *GetProjectSpendDetailsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get project spend details default response has a 4xx status code
func (o *GetProjectSpendDetailsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get project spend details default response has a 5xx status code
func (o *GetProjectSpendDetailsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get project spend details default response a status code equal to that given
func (o *GetProjectSpendDetailsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetProjectSpendDetailsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/spend/{spendId}/details][%d] GetProjectSpendDetails default  %+v", o._statusCode, o.Payload)
}

func (o *GetProjectSpendDetailsDefault) String() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/spend/{spendId}/details][%d] GetProjectSpendDetails default  %+v", o._statusCode, o.Payload)
}

func (o *GetProjectSpendDetailsDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *GetProjectSpendDetailsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
