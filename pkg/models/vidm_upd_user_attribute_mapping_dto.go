// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VidmUpdUserAttributeMappingDTO VidmUpdUserAttributeMappingDTO
//
// swagger:model VidmUpdUserAttributeMappingDTO
type VidmUpdUserAttributeMappingDTO struct {

	// directory config Id
	DirectoryConfigID string `json:"directoryConfigId,omitempty"`

	// user attribute mappings
	UserAttributeMappings []*UserAttributeMappingDTO `json:"userAttributeMappings"`

	// vidm admin password
	VidmAdminPassword string `json:"vidmAdminPassword,omitempty"`

	// vidm admin user
	VidmAdminUser string `json:"vidmAdminUser,omitempty"`

	// vidm host
	VidmHost string `json:"vidmHost,omitempty"`

	// vidm tenant
	VidmTenant string `json:"vidmTenant,omitempty"`
}

// Validate validates this vidm upd user attribute mapping d t o
func (m *VidmUpdUserAttributeMappingDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserAttributeMappings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VidmUpdUserAttributeMappingDTO) validateUserAttributeMappings(formats strfmt.Registry) error {

	if swag.IsZero(m.UserAttributeMappings) { // not required
		return nil
	}

	for i := 0; i < len(m.UserAttributeMappings); i++ {
		if swag.IsZero(m.UserAttributeMappings[i]) { // not required
			continue
		}

		if m.UserAttributeMappings[i] != nil {
			if err := m.UserAttributeMappings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userAttributeMappings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VidmUpdUserAttributeMappingDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VidmUpdUserAttributeMappingDTO) UnmarshalBinary(b []byte) error {
	var res VidmUpdUserAttributeMappingDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}