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

// UpdateBillingAccountReader is a Reader for the UpdateBillingAccount structure.
type UpdateBillingAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBillingAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateBillingAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateBillingAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateBillingAccountOK creates a UpdateBillingAccountOK with default headers values
func NewUpdateBillingAccountOK() *UpdateBillingAccountOK {
	return &UpdateBillingAccountOK{}
}

/*
UpdateBillingAccountOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateBillingAccountOK struct {
	Payload *models.BillingAccount
}

// IsSuccess returns true when this update billing account o k response has a 2xx status code
func (o *UpdateBillingAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update billing account o k response has a 3xx status code
func (o *UpdateBillingAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update billing account o k response has a 4xx status code
func (o *UpdateBillingAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update billing account o k response has a 5xx status code
func (o *UpdateBillingAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update billing account o k response a status code equal to that given
func (o *UpdateBillingAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update billing account o k response
func (o *UpdateBillingAccountOK) Code() int {
	return 200
}

func (o *UpdateBillingAccountOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/billing-accounts/{billingAccount.id}][%d] updateBillingAccountOK  %+v", 200, o.Payload)
}

func (o *UpdateBillingAccountOK) String() string {
	return fmt.Sprintf("[PATCH /v1/billing-accounts/{billingAccount.id}][%d] updateBillingAccountOK  %+v", 200, o.Payload)
}

func (o *UpdateBillingAccountOK) GetPayload() *models.BillingAccount {
	return o.Payload
}

func (o *UpdateBillingAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBillingAccountDefault creates a UpdateBillingAccountDefault with default headers values
func NewUpdateBillingAccountDefault(code int) *UpdateBillingAccountDefault {
	return &UpdateBillingAccountDefault{
		_statusCode: code,
	}
}

/*
UpdateBillingAccountDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateBillingAccountDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this update billing account default response has a 2xx status code
func (o *UpdateBillingAccountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update billing account default response has a 3xx status code
func (o *UpdateBillingAccountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update billing account default response has a 4xx status code
func (o *UpdateBillingAccountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update billing account default response has a 5xx status code
func (o *UpdateBillingAccountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update billing account default response a status code equal to that given
func (o *UpdateBillingAccountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update billing account default response
func (o *UpdateBillingAccountDefault) Code() int {
	return o._statusCode
}

func (o *UpdateBillingAccountDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1/billing-accounts/{billingAccount.id}][%d] UpdateBillingAccount default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateBillingAccountDefault) String() string {
	return fmt.Sprintf("[PATCH /v1/billing-accounts/{billingAccount.id}][%d] UpdateBillingAccount default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateBillingAccountDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *UpdateBillingAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
