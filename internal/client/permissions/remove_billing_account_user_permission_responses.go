// Code generated by go-swagger; DO NOT EDIT.

package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// RemoveBillingAccountUserPermissionReader is a Reader for the RemoveBillingAccountUserPermission structure.
type RemoveBillingAccountUserPermissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveBillingAccountUserPermissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveBillingAccountUserPermissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRemoveBillingAccountUserPermissionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveBillingAccountUserPermissionOK creates a RemoveBillingAccountUserPermissionOK with default headers values
func NewRemoveBillingAccountUserPermissionOK() *RemoveBillingAccountUserPermissionOK {
	return &RemoveBillingAccountUserPermissionOK{}
}

/*
	RemoveBillingAccountUserPermissionOK describes a response with status code 200, with default header values.

A successful response.
*/
type RemoveBillingAccountUserPermissionOK struct {
	Payload interface{}
}

// IsSuccess returns true when this remove billing account user permission o k response has a 2xx status code
func (o *RemoveBillingAccountUserPermissionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove billing account user permission o k response has a 3xx status code
func (o *RemoveBillingAccountUserPermissionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove billing account user permission o k response has a 4xx status code
func (o *RemoveBillingAccountUserPermissionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove billing account user permission o k response has a 5xx status code
func (o *RemoveBillingAccountUserPermissionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove billing account user permission o k response a status code equal to that given
func (o *RemoveBillingAccountUserPermissionOK) IsCode(code int) bool {
	return code == 200
}

func (o *RemoveBillingAccountUserPermissionOK) Error() string {
	return fmt.Sprintf("[POST /v1/billing-accounts/{billingAccountId}/remove-user-permission][%d] removeBillingAccountUserPermissionOK  %+v", 200, o.Payload)
}

func (o *RemoveBillingAccountUserPermissionOK) String() string {
	return fmt.Sprintf("[POST /v1/billing-accounts/{billingAccountId}/remove-user-permission][%d] removeBillingAccountUserPermissionOK  %+v", 200, o.Payload)
}

func (o *RemoveBillingAccountUserPermissionOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RemoveBillingAccountUserPermissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveBillingAccountUserPermissionDefault creates a RemoveBillingAccountUserPermissionDefault with default headers values
func NewRemoveBillingAccountUserPermissionDefault(code int) *RemoveBillingAccountUserPermissionDefault {
	return &RemoveBillingAccountUserPermissionDefault{
		_statusCode: code,
	}
}

/*
	RemoveBillingAccountUserPermissionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RemoveBillingAccountUserPermissionDefault struct {
	_statusCode int

	Payload *models.Status
}

// Code gets the status code for the remove billing account user permission default response
func (o *RemoveBillingAccountUserPermissionDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this remove billing account user permission default response has a 2xx status code
func (o *RemoveBillingAccountUserPermissionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this remove billing account user permission default response has a 3xx status code
func (o *RemoveBillingAccountUserPermissionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this remove billing account user permission default response has a 4xx status code
func (o *RemoveBillingAccountUserPermissionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this remove billing account user permission default response has a 5xx status code
func (o *RemoveBillingAccountUserPermissionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this remove billing account user permission default response a status code equal to that given
func (o *RemoveBillingAccountUserPermissionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *RemoveBillingAccountUserPermissionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/billing-accounts/{billingAccountId}/remove-user-permission][%d] RemoveBillingAccountUserPermission default  %+v", o._statusCode, o.Payload)
}

func (o *RemoveBillingAccountUserPermissionDefault) String() string {
	return fmt.Sprintf("[POST /v1/billing-accounts/{billingAccountId}/remove-user-permission][%d] RemoveBillingAccountUserPermission default  %+v", o._statusCode, o.Payload)
}

func (o *RemoveBillingAccountUserPermissionDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *RemoveBillingAccountUserPermissionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
RemoveBillingAccountUserPermissionBody remove billing account user permission body
swagger:model RemoveBillingAccountUserPermissionBody
*/
type RemoveBillingAccountUserPermissionBody struct {

	// data center Id
	DataCenterID string `json:"dataCenterId,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// role
	// Required: true
	Role *models.Role `json:"role"`

	// user Id
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this remove billing account user permission body
func (o *RemoveBillingAccountUserPermissionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveBillingAccountUserPermissionBody) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"role", "body", o.Role); err != nil {
		return err
	}

	if err := validate.Required("body"+"."+"role", "body", o.Role); err != nil {
		return err
	}

	if o.Role != nil {
		if err := o.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "role")
			}
			return err
		}
	}

	return nil
}

func (o *RemoveBillingAccountUserPermissionBody) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"userId", "body", o.UserID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this remove billing account user permission body based on the context it is used
func (o *RemoveBillingAccountUserPermissionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveBillingAccountUserPermissionBody) contextValidateRole(ctx context.Context, formats strfmt.Registry) error {

	if o.Role != nil {
		if err := o.Role.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "role")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RemoveBillingAccountUserPermissionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveBillingAccountUserPermissionBody) UnmarshalBinary(b []byte) error {
	var res RemoveBillingAccountUserPermissionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
