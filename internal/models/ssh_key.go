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

// SSHKey Ssh key
//
// swagger:model SshKey
type SSHKey struct {

	// comment
	// Read Only: true
	Comment string `json:"comment,omitempty"`

	// create time
	// Read Only: true
	// Format: date-time
	CreateTime strfmt.DateTime `json:"createTime,omitempty"`

	// fingerprint
	// Read Only: true
	Fingerprint string `json:"fingerprint,omitempty"`

	// id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// public key
	// Required: true
	PublicKey *string `json:"publicKey"`

	// type
	// Read Only: true
	Type string `json:"type,omitempty"`
}

// Validate validates this Ssh key
func (m *SSHKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SSHKey) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("createTime", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SSHKey) validatePublicKey(formats strfmt.Registry) error {

	if err := validate.Required("publicKey", "body", m.PublicKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this Ssh key based on the context it is used
func (m *SSHKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFingerprint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SSHKey) contextValidateComment(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "comment", "body", string(m.Comment)); err != nil {
		return err
	}

	return nil
}

func (m *SSHKey) contextValidateCreateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createTime", "body", strfmt.DateTime(m.CreateTime)); err != nil {
		return err
	}

	return nil
}

func (m *SSHKey) contextValidateFingerprint(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "fingerprint", "body", string(m.Fingerprint)); err != nil {
		return err
	}

	return nil
}

func (m *SSHKey) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *SSHKey) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SSHKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SSHKey) UnmarshalBinary(b []byte) error {
	var res SSHKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
