// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Component Component
//
// swagger:model Component
type Component struct {

	// component version
	ComponentVersion string `json:"componentVersion,omitempty"`

	// properties
	Properties map[string]string `json:"properties,omitempty"`

	// symbolic name
	SymbolicName string `json:"symbolicName,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this component
func (m *Component) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Component) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Component) UnmarshalBinary(b []byte) error {
	var res Component
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
