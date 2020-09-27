// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RoleRequestDTO RoleRequestDTO
//
// swagger:model RoleRequestDTO
type RoleRequestDTO struct {

	// authorities
	Authorities string `json:"authorities,omitempty"`

	// is internal
	IsInternal bool `json:"isInternal,omitempty"`

	// role description
	RoleDescription string `json:"roleDescription,omitempty"`

	// role name
	RoleName string `json:"roleName,omitempty"`
}

// Validate validates this role request d t o
func (m *RoleRequestDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RoleRequestDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleRequestDTO) UnmarshalBinary(b []byte) error {
	var res RoleRequestDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
