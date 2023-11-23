// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DiskStorageClass disk storage class
//
// swagger:model Disk.StorageClass
type DiskStorageClass string

func NewDiskStorageClass(value DiskStorageClass) *DiskStorageClass {
	return &value
}

// Pointer returns a pointer to a freshly-allocated DiskStorageClass.
func (m DiskStorageClass) Pointer() *DiskStorageClass {
	return &m
}

const (

	// DiskStorageClassSTORAGECLASSUNKNOWN captures enum value "STORAGE_CLASS_UNKNOWN"
	DiskStorageClassSTORAGECLASSUNKNOWN DiskStorageClass = "STORAGE_CLASS_UNKNOWN"

	// DiskStorageClassSTORAGECLASSLOCAL captures enum value "STORAGE_CLASS_LOCAL"
	DiskStorageClassSTORAGECLASSLOCAL DiskStorageClass = "STORAGE_CLASS_LOCAL"

	// DiskStorageClassSTORAGECLASSNETWORK captures enum value "STORAGE_CLASS_NETWORK"
	DiskStorageClassSTORAGECLASSNETWORK DiskStorageClass = "STORAGE_CLASS_NETWORK"
)

// for schema
var diskStorageClassEnum []interface{}

func init() {
	var res []DiskStorageClass
	if err := json.Unmarshal([]byte(`["STORAGE_CLASS_UNKNOWN","STORAGE_CLASS_LOCAL","STORAGE_CLASS_NETWORK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		diskStorageClassEnum = append(diskStorageClassEnum, v)
	}
}

func (m DiskStorageClass) validateDiskStorageClassEnum(path, location string, value DiskStorageClass) error {
	if err := validate.EnumCase(path, location, value, diskStorageClassEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this disk storage class
func (m DiskStorageClass) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDiskStorageClassEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this disk storage class based on context it is used
func (m DiskStorageClass) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
