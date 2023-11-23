// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetDiskResponse get disk response
//
// swagger:model GetDiskResponse
type GetDiskResponse struct {

	// disk
	Disk *Disk `json:"disk,omitempty"`

	// snapshots
	Snapshots []*Snapshot `json:"snapshots"`
}

// Validate validates this get disk response
func (m *GetDiskResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshots(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetDiskResponse) validateDisk(formats strfmt.Registry) error {
	if swag.IsZero(m.Disk) { // not required
		return nil
	}

	if m.Disk != nil {
		if err := m.Disk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disk")
			}
			return err
		}
	}

	return nil
}

func (m *GetDiskResponse) validateSnapshots(formats strfmt.Registry) error {
	if swag.IsZero(m.Snapshots) { // not required
		return nil
	}

	for i := 0; i < len(m.Snapshots); i++ {
		if swag.IsZero(m.Snapshots[i]) { // not required
			continue
		}

		if m.Snapshots[i] != nil {
			if err := m.Snapshots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("snapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get disk response based on the context it is used
func (m *GetDiskResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisk(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetDiskResponse) contextValidateDisk(ctx context.Context, formats strfmt.Registry) error {

	if m.Disk != nil {

		if swag.IsZero(m.Disk) { // not required
			return nil
		}

		if err := m.Disk.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disk")
			}
			return err
		}
	}

	return nil
}

func (m *GetDiskResponse) contextValidateSnapshots(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Snapshots); i++ {

		if m.Snapshots[i] != nil {

			if swag.IsZero(m.Snapshots[i]) { // not required
				return nil
			}

			if err := m.Snapshots[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("snapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetDiskResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetDiskResponse) UnmarshalBinary(b []byte) error {
	var res GetDiskResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
