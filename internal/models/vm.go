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
	"github.com/go-openapi/validate"
)

// VM VM
//
// swagger:model VM
type VM struct {

	// active state
	ActiveState string `json:"activeState,omitempty"`

	// boot disk
	BootDisk *Disk `json:"bootDisk,omitempty"`

	// boot disk size gib
	BootDiskSizeGib int64 `json:"bootDiskSizeGib,omitempty"`

	// cpu model
	CPUModel string `json:"cpuModel,omitempty"`

	// create by
	// Read Only: true
	CreateBy string `json:"createBy,omitempty"`

	// datacenter Id
	// Read Only: true
	DatacenterID string `json:"datacenterId,omitempty"`

	// external Ip address
	ExternalIPAddress string `json:"externalIpAddress,omitempty"`

	// gpu model
	GpuModel string `json:"gpuModel,omitempty"`

	// gpu quantity
	GpuQuantity int64 `json:"gpuQuantity,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// image Id
	ImageID string `json:"imageId,omitempty"`

	// image name
	ImageName string `json:"imageName,omitempty"`

	// init state
	InitState string `json:"initState,omitempty"`

	// internal Ip address
	InternalIPAddress string `json:"internalIpAddress,omitempty"`

	// lcm state
	LcmState string `json:"lcmState,omitempty"`

	// machine type
	MachineType string `json:"machineType,omitempty"`

	// memory
	Memory int64 `json:"memory,omitempty"`

	// nics
	Nics []*VMNIC `json:"nics"`

	// one state
	OneState string `json:"oneState,omitempty"`

	// price hr
	PriceHr float32 `json:"priceHr,omitempty"`

	// private image Id
	PrivateImageID string `json:"privateImageId,omitempty"`

	// public image Id
	PublicImageID string `json:"publicImageId,omitempty"`

	// public image name
	PublicImageName string `json:"publicImageName,omitempty"`

	// public Ip address
	PublicIPAddress string `json:"publicIpAddress,omitempty"`

	// region Id
	// Read Only: true
	RegionID string `json:"regionId,omitempty"`

	// region name
	// Read Only: true
	RegionName string `json:"regionName,omitempty"`

	// renewable energy
	RenewableEnergy bool `json:"renewableEnergy,omitempty"`

	// rules
	Rules []*SecurityGroupRule `json:"rules"`

	// security group ids
	SecurityGroupIds []string `json:"securityGroupIds"`

	// short state
	ShortState string `json:"shortState,omitempty"`

	// storage disks
	StorageDisks []*Disk `json:"storageDisks"`

	// vcpus
	Vcpus int64 `json:"vcpus,omitempty"`
}

// Validate validates this VM
func (m *VM) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBootDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageDisks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VM) validateBootDisk(formats strfmt.Registry) error {
	if swag.IsZero(m.BootDisk) { // not required
		return nil
	}

	if m.BootDisk != nil {
		if err := m.BootDisk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bootDisk")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bootDisk")
			}
			return err
		}
	}

	return nil
}

func (m *VM) validateNics(formats strfmt.Registry) error {
	if swag.IsZero(m.Nics) { // not required
		return nil
	}

	for i := 0; i < len(m.Nics); i++ {
		if swag.IsZero(m.Nics[i]) { // not required
			continue
		}

		if m.Nics[i] != nil {
			if err := m.Nics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VM) validateRules(formats strfmt.Registry) error {
	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	for i := 0; i < len(m.Rules); i++ {
		if swag.IsZero(m.Rules[i]) { // not required
			continue
		}

		if m.Rules[i] != nil {
			if err := m.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VM) validateStorageDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageDisks) { // not required
		return nil
	}

	for i := 0; i < len(m.StorageDisks); i++ {
		if swag.IsZero(m.StorageDisks[i]) { // not required
			continue
		}

		if m.StorageDisks[i] != nil {
			if err := m.StorageDisks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("storageDisks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("storageDisks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this VM based on the context it is used
func (m *VM) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBootDisk(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreateBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatacenterID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegionID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegionName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VM) contextValidateBootDisk(ctx context.Context, formats strfmt.Registry) error {

	if m.BootDisk != nil {
		if err := m.BootDisk.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bootDisk")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bootDisk")
			}
			return err
		}
	}

	return nil
}

func (m *VM) contextValidateCreateBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createBy", "body", string(m.CreateBy)); err != nil {
		return err
	}

	return nil
}

func (m *VM) contextValidateDatacenterID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "datacenterId", "body", string(m.DatacenterID)); err != nil {
		return err
	}

	return nil
}

func (m *VM) contextValidateNics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Nics); i++ {

		if m.Nics[i] != nil {
			if err := m.Nics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VM) contextValidateRegionID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "regionId", "body", string(m.RegionID)); err != nil {
		return err
	}

	return nil
}

func (m *VM) contextValidateRegionName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "regionName", "body", string(m.RegionName)); err != nil {
		return err
	}

	return nil
}

func (m *VM) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Rules); i++ {

		if m.Rules[i] != nil {
			if err := m.Rules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VM) contextValidateStorageDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StorageDisks); i++ {

		if m.StorageDisks[i] != nil {
			if err := m.StorageDisks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("storageDisks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("storageDisks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VM) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VM) UnmarshalBinary(b []byte) error {
	var res VM
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
