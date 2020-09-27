// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VidmDirConfigRequestDTO VidmDirConfigRequestDTO
//
// swagger:model VidmDirConfigRequestDTO
type VidmDirConfigRequestDTO struct {

	// directory config
	DirectoryConfig interface{} `json:"directoryConfig,omitempty"`

	// directory type
	DirectoryType string `json:"directoryType,omitempty"`

	// vidm admin password
	VidmAdminPassword string `json:"vidmAdminPassword,omitempty"`

	// vidm admin user
	VidmAdminUser string `json:"vidmAdminUser,omitempty"`

	// vidm host
	VidmHost string `json:"vidmHost,omitempty"`

	// vidm tenant
	VidmTenant string `json:"vidmTenant,omitempty"`
}

// Validate validates this vidm dir config request d t o
func (m *VidmDirConfigRequestDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VidmDirConfigRequestDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VidmDirConfigRequestDTO) UnmarshalBinary(b []byte) error {
	var res VidmDirConfigRequestDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}