// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateAdminPasswordRequestDTO UpdateAdminPasswordRequestDTO
//
// swagger:model UpdateAdminPasswordRequestDTO
type UpdateAdminPasswordRequestDTO struct {

	// admin password
	AdminPassword string `json:"adminPassword,omitempty"`

	// current admin password
	CurrentAdminPassword string `json:"currentAdminPassword,omitempty"`
}

// Validate validates this update admin password request d t o
func (m *UpdateAdminPasswordRequestDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateAdminPasswordRequestDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateAdminPasswordRequestDTO) UnmarshalBinary(b []byte) error {
	var res UpdateAdminPasswordRequestDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}