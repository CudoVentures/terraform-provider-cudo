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

// RuleProtocol rule protocol
//
// swagger:model Rule.Protocol
type RuleProtocol string

func NewRuleProtocol(value RuleProtocol) *RuleProtocol {
	return &value
}

// Pointer returns a pointer to a freshly-allocated RuleProtocol.
func (m RuleProtocol) Pointer() *RuleProtocol {
	return &m
}

const (

	// RuleProtocolPROTOCOLUNKNOWN captures enum value "PROTOCOL_UNKNOWN"
	RuleProtocolPROTOCOLUNKNOWN RuleProtocol = "PROTOCOL_UNKNOWN"

	// RuleProtocolPROTOCOLALL captures enum value "PROTOCOL_ALL"
	RuleProtocolPROTOCOLALL RuleProtocol = "PROTOCOL_ALL"

	// RuleProtocolPROTOCOLTCP captures enum value "PROTOCOL_TCP"
	RuleProtocolPROTOCOLTCP RuleProtocol = "PROTOCOL_TCP"

	// RuleProtocolPROTOCOLUDP captures enum value "PROTOCOL_UDP"
	RuleProtocolPROTOCOLUDP RuleProtocol = "PROTOCOL_UDP"

	// RuleProtocolPROTOCOLICMP captures enum value "PROTOCOL_ICMP"
	RuleProtocolPROTOCOLICMP RuleProtocol = "PROTOCOL_ICMP"

	// RuleProtocolPROTOCOLICMPv6 captures enum value "PROTOCOL_ICMPv6"
	RuleProtocolPROTOCOLICMPv6 RuleProtocol = "PROTOCOL_ICMPv6"

	// RuleProtocolPROTOCOLIPSEC captures enum value "PROTOCOL_IPSEC"
	RuleProtocolPROTOCOLIPSEC RuleProtocol = "PROTOCOL_IPSEC"
)

// for schema
var ruleProtocolEnum []interface{}

func init() {
	var res []RuleProtocol
	if err := json.Unmarshal([]byte(`["PROTOCOL_UNKNOWN","PROTOCOL_ALL","PROTOCOL_TCP","PROTOCOL_UDP","PROTOCOL_ICMP","PROTOCOL_ICMPv6","PROTOCOL_IPSEC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ruleProtocolEnum = append(ruleProtocolEnum, v)
	}
}

func (m RuleProtocol) validateRuleProtocolEnum(path, location string, value RuleProtocol) error {
	if err := validate.EnumCase(path, location, value, ruleProtocolEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this rule protocol
func (m RuleProtocol) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRuleProtocolEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this rule protocol based on context it is used
func (m RuleProtocol) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
