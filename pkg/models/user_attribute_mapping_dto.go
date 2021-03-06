// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserAttributeMappingDTO UserAttributeMappingDTO
//
// swagger:model UserAttributeMappingDTO
type UserAttributeMappingDTO struct {

	// attr name
	AttrName string `json:"attrName,omitempty"`

	// directory name
	DirectoryName string `json:"directoryName,omitempty"`

	// required
	Required bool `json:"required,omitempty"`
}

// Validate validates this user attribute mapping d t o
func (m *UserAttributeMappingDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserAttributeMappingDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserAttributeMappingDTO) UnmarshalBinary(b []byte) error {
	var res UserAttributeMappingDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
