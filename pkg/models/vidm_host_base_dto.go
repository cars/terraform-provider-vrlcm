// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VidmHostBaseDTO VidmHostBaseDTO
//
// swagger:model VidmHostBaseDTO
type VidmHostBaseDTO struct {

	// vidm admin password
	VidmAdminPassword string `json:"vidmAdminPassword,omitempty"`

	// vidm admin user
	VidmAdminUser string `json:"vidmAdminUser,omitempty"`

	// vidm host
	VidmHost string `json:"vidmHost,omitempty"`

	// vidm tenant
	VidmTenant string `json:"vidmTenant,omitempty"`
}

// Validate validates this vidm host base d t o
func (m *VidmHostBaseDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VidmHostBaseDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VidmHostBaseDTO) UnmarshalBinary(b []byte) error {
	var res VidmHostBaseDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
