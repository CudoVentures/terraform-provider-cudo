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

// DataCenterTimeRevenue data center time revenue
//
// swagger:model DataCenterTimeRevenue
type DataCenterTimeRevenue struct {

	// revenue usd
	// Required: true
	RevenueUsd *Decimal `json:"revenueUsd"`

	// time
	// Required: true
	// Format: date-time
	Time *strfmt.DateTime `json:"time"`
}

// Validate validates this data center time revenue
func (m *DataCenterTimeRevenue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRevenueUsd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataCenterTimeRevenue) validateRevenueUsd(formats strfmt.Registry) error {

	if err := validate.Required("revenueUsd", "body", m.RevenueUsd); err != nil {
		return err
	}

	if m.RevenueUsd != nil {
		if err := m.RevenueUsd.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revenueUsd")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revenueUsd")
			}
			return err
		}
	}

	return nil
}

func (m *DataCenterTimeRevenue) validateTime(formats strfmt.Registry) error {

	if err := validate.Required("time", "body", m.Time); err != nil {
		return err
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this data center time revenue based on the context it is used
func (m *DataCenterTimeRevenue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRevenueUsd(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataCenterTimeRevenue) contextValidateRevenueUsd(ctx context.Context, formats strfmt.Registry) error {

	if m.RevenueUsd != nil {

		if err := m.RevenueUsd.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revenueUsd")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revenueUsd")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataCenterTimeRevenue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataCenterTimeRevenue) UnmarshalBinary(b []byte) error {
	var res DataCenterTimeRevenue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
