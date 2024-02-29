// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetBillingAccountDetailsResponse get billing account details response
//
// swagger:model GetBillingAccountDetailsResponse
type GetBillingAccountDetailsResponse struct {

	// billing account
	// Required: true
	BillingAccount *BillingAccount `json:"billingAccount"`

	// stripe customer
	// Required: true
	StripeCustomer *StripeCustomer `json:"stripeCustomer"`
}

// Validate validates this get billing account details response
func (m *GetBillingAccountDetailsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStripeCustomer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetBillingAccountDetailsResponse) validateBillingAccount(formats strfmt.Registry) error {

	if err := validate.Required("billingAccount", "body", m.BillingAccount); err != nil {
		return err
	}

	if m.BillingAccount != nil {
		if err := m.BillingAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billingAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billingAccount")
			}
			return err
		}
	}

	return nil
}

func (m *GetBillingAccountDetailsResponse) validateStripeCustomer(formats strfmt.Registry) error {

	if err := validate.Required("stripeCustomer", "body", m.StripeCustomer); err != nil {
		return err
	}

	if m.StripeCustomer != nil {
		if err := m.StripeCustomer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stripeCustomer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stripeCustomer")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get billing account details response based on the context it is used
func (m *GetBillingAccountDetailsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBillingAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStripeCustomer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetBillingAccountDetailsResponse) contextValidateBillingAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.BillingAccount != nil {

		if err := m.BillingAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billingAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billingAccount")
			}
			return err
		}
	}

	return nil
}

func (m *GetBillingAccountDetailsResponse) contextValidateStripeCustomer(ctx context.Context, formats strfmt.Registry) error {

	if m.StripeCustomer != nil {

		if err := m.StripeCustomer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stripeCustomer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stripeCustomer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetBillingAccountDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBillingAccountDetailsResponse) UnmarshalBinary(b []byte) error {
	var res GetBillingAccountDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
