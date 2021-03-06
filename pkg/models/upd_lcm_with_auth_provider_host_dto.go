// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdLcmWithAuthProviderHostDTO UpdLcmWithAuthProviderHostDTO
//
// swagger:model UpdLcmWithAuthProviderHostDTO
type UpdLcmWithAuthProviderHostDTO struct {

	// vidm admin password
	VidmAdminPassword string `json:"vidmAdminPassword,omitempty"`

	// vidm hostname
	VidmHostname string `json:"vidmHostname,omitempty"`
}

// Validate validates this upd lcm with auth provider host d t o
func (m *UpdLcmWithAuthProviderHostDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdLcmWithAuthProviderHostDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdLcmWithAuthProviderHostDTO) UnmarshalBinary(b []byte) error {
	var res UpdLcmWithAuthProviderHostDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
