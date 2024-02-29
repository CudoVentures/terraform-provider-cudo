// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Invoice invoice
//
// swagger:model Invoice
type Invoice struct {

	// amount due
	AmountDue string `json:"amountDue,omitempty"`

	// amount paid
	AmountPaid string `json:"amountPaid,omitempty"`

	// amount remaining
	AmountRemaining string `json:"amountRemaining,omitempty"`

	// auto advance
	AutoAdvance bool `json:"autoAdvance,omitempty"`

	// billing reason
	BillingReason string `json:"billingReason,omitempty"`

	// created
	Created string `json:"created,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// due date
	DueDate string `json:"dueDate,omitempty"`

	// hosted invoice Url
	HostedInvoiceURL string `json:"hostedInvoiceUrl,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// number
	Number string `json:"number,omitempty"`

	// paid date
	PaidDate string `json:"paidDate,omitempty"`

	// period end
	PeriodEnd string `json:"periodEnd,omitempty"`

	// period start
	PeriodStart string `json:"periodStart,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this invoice
func (m *Invoice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this invoice based on context it is used
func (m *Invoice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Invoice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Invoice) UnmarshalBinary(b []byte) error {
	var res Invoice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
