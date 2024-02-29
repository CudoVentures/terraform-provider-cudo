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

// DataCenterMachineType data center machine type
//
// swagger:model DataCenterMachineType
type DataCenterMachineType struct {

	// count clusters
	// Required: true
	CountClusters *int32 `json:"countClusters"`

	// count hosts
	// Required: true
	CountHosts *int32 `json:"countHosts"`

	// count hosts active
	// Required: true
	CountHostsActive *int32 `json:"countHostsActive"`

	// count hosts inactive
	// Required: true
	CountHostsInactive *int32 `json:"countHostsInactive"`

	// machine type
	// Required: true
	MachineType *string `json:"machineType"`
}

// Validate validates this data center machine type
func (m *DataCenterMachineType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountHostsActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountHostsInactive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataCenterMachineType) validateCountClusters(formats strfmt.Registry) error {

	if err := validate.Required("countClusters", "body", m.CountClusters); err != nil {
		return err
	}

	return nil
}

func (m *DataCenterMachineType) validateCountHosts(formats strfmt.Registry) error {

	if err := validate.Required("countHosts", "body", m.CountHosts); err != nil {
		return err
	}

	return nil
}

func (m *DataCenterMachineType) validateCountHostsActive(formats strfmt.Registry) error {

	if err := validate.Required("countHostsActive", "body", m.CountHostsActive); err != nil {
		return err
	}

	return nil
}

func (m *DataCenterMachineType) validateCountHostsInactive(formats strfmt.Registry) error {

	if err := validate.Required("countHostsInactive", "body", m.CountHostsInactive); err != nil {
		return err
	}

	return nil
}

func (m *DataCenterMachineType) validateMachineType(formats strfmt.Registry) error {

	if err := validate.Required("machineType", "body", m.MachineType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this data center machine type based on context it is used
func (m *DataCenterMachineType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataCenterMachineType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataCenterMachineType) UnmarshalBinary(b []byte) error {
	var res DataCenterMachineType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
