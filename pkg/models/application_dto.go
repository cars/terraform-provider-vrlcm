// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApplicationDTO ApplicationDTO
//
// swagger:model ApplicationDTO
type ApplicationDTO struct {

	// access order
	AccessOrder int32 `json:"accessOrder,omitempty"`

	// app Url
	AppURL string `json:"appUrl,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// detailed description
	DetailedDescription string `json:"detailedDescription,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// icon
	Icon string `json:"icon,omitempty"`

	// licensed
	Licensed bool `json:"licensed,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// transaction Id
	TransactionID string `json:"transactionId,omitempty"`

	// vmid
	Vmid string `json:"vmid,omitempty"`
}

// Validate validates this application d t o
func (m *ApplicationDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationDTO) UnmarshalBinary(b []byte) error {
	var res ApplicationDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}