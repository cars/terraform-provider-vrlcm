// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ComponentDeploymentSpecification ComponentDeploymentSpecification
//
// swagger:model ComponentDeploymentSpecification
type ComponentDeploymentSpecification struct {

	// component
	Component *Component `json:"component,omitempty"`

	// priority
	Priority int64 `json:"priority,omitempty"`
}

// Validate validates this component deployment specification
func (m *ComponentDeploymentSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComponentDeploymentSpecification) validateComponent(formats strfmt.Registry) error {

	if swag.IsZero(m.Component) { // not required
		return nil
	}

	if m.Component != nil {
		if err := m.Component.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComponentDeploymentSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComponentDeploymentSpecification) UnmarshalBinary(b []byte) error {
	var res ComponentDeploymentSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
