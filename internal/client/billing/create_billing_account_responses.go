// Code generated by go-swagger; DO NOT EDIT.

package billing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// CreateBillingAccountReader is a Reader for the CreateBillingAccount structure.
type CreateBillingAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBillingAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateBillingAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateBillingAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateBillingAccountOK creates a CreateBillingAccountOK with default headers values
func NewCreateBillingAccountOK() *CreateBillingAccountOK {
	return &CreateBillingAccountOK{}
}

/*
CreateBillingAccountOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateBillingAccountOK struct {
	Payload *models.BillingAccount
}

// IsSuccess returns true when this create billing account o k response has a 2xx status code
func (o *CreateBillingAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create billing account o k response has a 3xx status code
func (o *CreateBillingAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create billing account o k response has a 4xx status code
func (o *CreateBillingAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create billing account o k response has a 5xx status code
func (o *CreateBillingAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create billing account o k response a status code equal to that given
func (o *CreateBillingAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create billing account o k response
func (o *CreateBillingAccountOK) Code() int {
	return 200
}

func (o *CreateBillingAccountOK) Error() string {
	return fmt.Sprintf("[POST /v1/billing-accounts][%d] createBillingAccountOK  %+v", 200, o.Payload)
}

func (o *CreateBillingAccountOK) String() string {
	return fmt.Sprintf("[POST /v1/billing-accounts][%d] createBillingAccountOK  %+v", 200, o.Payload)
}

func (o *CreateBillingAccountOK) GetPayload() *models.BillingAccount {
	return o.Payload
}

func (o *CreateBillingAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBillingAccountDefault creates a CreateBillingAccountDefault with default headers values
func NewCreateBillingAccountDefault(code int) *CreateBillingAccountDefault {
	return &CreateBillingAccountDefault{
		_statusCode: code,
	}
}

/*
CreateBillingAccountDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateBillingAccountDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this create billing account default response has a 2xx status code
func (o *CreateBillingAccountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create billing account default response has a 3xx status code
func (o *CreateBillingAccountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create billing account default response has a 4xx status code
func (o *CreateBillingAccountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create billing account default response has a 5xx status code
func (o *CreateBillingAccountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create billing account default response a status code equal to that given
func (o *CreateBillingAccountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create billing account default response
func (o *CreateBillingAccountDefault) Code() int {
	return o._statusCode
}

func (o *CreateBillingAccountDefault) Error() string {
	return fmt.Sprintf("[POST /v1/billing-accounts][%d] CreateBillingAccount default  %+v", o._statusCode, o.Payload)
}

func (o *CreateBillingAccountDefault) String() string {
	return fmt.Sprintf("[POST /v1/billing-accounts][%d] CreateBillingAccount default  %+v", o._statusCode, o.Payload)
}

func (o *CreateBillingAccountDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *CreateBillingAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
