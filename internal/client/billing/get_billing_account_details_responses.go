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

// GetBillingAccountDetailsReader is a Reader for the GetBillingAccountDetails structure.
type GetBillingAccountDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBillingAccountDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBillingAccountDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBillingAccountDetailsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBillingAccountDetailsOK creates a GetBillingAccountDetailsOK with default headers values
func NewGetBillingAccountDetailsOK() *GetBillingAccountDetailsOK {
	return &GetBillingAccountDetailsOK{}
}

/*
GetBillingAccountDetailsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetBillingAccountDetailsOK struct {
	Payload *models.GetBillingAccountDetailsResponse
}

// IsSuccess returns true when this get billing account details o k response has a 2xx status code
func (o *GetBillingAccountDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get billing account details o k response has a 3xx status code
func (o *GetBillingAccountDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get billing account details o k response has a 4xx status code
func (o *GetBillingAccountDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get billing account details o k response has a 5xx status code
func (o *GetBillingAccountDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get billing account details o k response a status code equal to that given
func (o *GetBillingAccountDetailsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get billing account details o k response
func (o *GetBillingAccountDetailsOK) Code() int {
	return 200
}

func (o *GetBillingAccountDetailsOK) Error() string {
	return fmt.Sprintf("[GET /v1/billing-accounts/{id}/details][%d] getBillingAccountDetailsOK  %+v", 200, o.Payload)
}

func (o *GetBillingAccountDetailsOK) String() string {
	return fmt.Sprintf("[GET /v1/billing-accounts/{id}/details][%d] getBillingAccountDetailsOK  %+v", 200, o.Payload)
}

func (o *GetBillingAccountDetailsOK) GetPayload() *models.GetBillingAccountDetailsResponse {
	return o.Payload
}

func (o *GetBillingAccountDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetBillingAccountDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBillingAccountDetailsDefault creates a GetBillingAccountDetailsDefault with default headers values
func NewGetBillingAccountDetailsDefault(code int) *GetBillingAccountDetailsDefault {
	return &GetBillingAccountDetailsDefault{
		_statusCode: code,
	}
}

/*
GetBillingAccountDetailsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetBillingAccountDetailsDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this get billing account details default response has a 2xx status code
func (o *GetBillingAccountDetailsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get billing account details default response has a 3xx status code
func (o *GetBillingAccountDetailsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get billing account details default response has a 4xx status code
func (o *GetBillingAccountDetailsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get billing account details default response has a 5xx status code
func (o *GetBillingAccountDetailsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get billing account details default response a status code equal to that given
func (o *GetBillingAccountDetailsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get billing account details default response
func (o *GetBillingAccountDetailsDefault) Code() int {
	return o._statusCode
}

func (o *GetBillingAccountDetailsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/billing-accounts/{id}/details][%d] GetBillingAccountDetails default  %+v", o._statusCode, o.Payload)
}

func (o *GetBillingAccountDetailsDefault) String() string {
	return fmt.Sprintf("[GET /v1/billing-accounts/{id}/details][%d] GetBillingAccountDetails default  %+v", o._statusCode, o.Payload)
}

func (o *GetBillingAccountDetailsDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *GetBillingAccountDetailsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
