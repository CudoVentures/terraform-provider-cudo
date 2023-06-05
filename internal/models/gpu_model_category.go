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

// GpuModelCategory gpu model category
//
// swagger:model GpuModelCategory
type GpuModelCategory struct {

	// count Vm available
	// Required: true
	CountVMAvailable *int32 `json:"countVmAvailable"`

	// min price hr
	// Required: true
	MinPriceHr *Decimal `json:"minPriceHr"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this gpu model category
func (m *GpuModelCategory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountVMAvailable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinPriceHr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GpuModelCategory) validateCountVMAvailable(formats strfmt.Registry) error {

	if err := validate.Required("countVmAvailable", "body", m.CountVMAvailable); err != nil {
		return err
	}

	return nil
}

func (m *GpuModelCategory) validateMinPriceHr(formats strfmt.Registry) error {

	if err := validate.Required("minPriceHr", "body", m.MinPriceHr); err != nil {
		return err
	}

	if m.MinPriceHr != nil {
		if err := m.MinPriceHr.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("minPriceHr")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("minPriceHr")
			}
			return err
		}
	}

	return nil
}

func (m *GpuModelCategory) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this gpu model category based on the context it is used
func (m *GpuModelCategory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMinPriceHr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GpuModelCategory) contextValidateMinPriceHr(ctx context.Context, formats strfmt.Registry) error {

	if m.MinPriceHr != nil {
		if err := m.MinPriceHr.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("minPriceHr")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("minPriceHr")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GpuModelCategory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GpuModelCategory) UnmarshalBinary(b []byte) error {
	var res GpuModelCategory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
