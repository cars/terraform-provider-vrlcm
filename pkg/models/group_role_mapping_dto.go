// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GroupRoleMappingDTO GroupRoleMappingDTO
//
// swagger:model GroupRoleMappingDTO
type GroupRoleMappingDTO struct {

	// groupvmid
	Groupvmid string `json:"groupvmid,omitempty"`

	// rolevmid
	Rolevmid string `json:"rolevmid,omitempty"`
}

// Validate validates this group role mapping d t o
func (m *GroupRoleMappingDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GroupRoleMappingDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupRoleMappingDTO) UnmarshalBinary(b []byte) error {
	var res GroupRoleMappingDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
