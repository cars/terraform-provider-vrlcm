// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PasswordRequestResponseDTO PasswordRequestResponseDTO
//
// swagger:model PasswordRequestResponseDTO
type PasswordRequestResponseDTO struct {

	// alias
	Alias string `json:"alias,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// password description
	PasswordDescription string `json:"passwordDescription,omitempty"`

	// principal
	Principal string `json:"principal,omitempty"`

	// transaction Id
	TransactionID string `json:"transactionId,omitempty"`

	// user name
	UserName string `json:"userName,omitempty"`

	// vmid
	Vmid string `json:"vmid,omitempty"`
}

// Validate validates this password request response d t o
func (m *PasswordRequestResponseDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PasswordRequestResponseDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PasswordRequestResponseDTO) UnmarshalBinary(b []byte) error {
	var res PasswordRequestResponseDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
