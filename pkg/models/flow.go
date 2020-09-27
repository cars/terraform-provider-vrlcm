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

// Flow Flow
//
// swagger:model Flow
type Flow struct {

	// edges
	Edges []*Edge `json:"edges"`

	// error state
	ErrorState string `json:"errorState,omitempty"`

	// finish state
	FinishState string `json:"finishState,omitempty"`

	// nodes
	Nodes []*Node `json:"nodes"`

	// properties
	Properties map[string]string `json:"properties,omitempty"`

	// start state
	StartState string `json:"startState,omitempty"`

	// symbolic name
	SymbolicName string `json:"symbolicName,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// ui properties
	UIProperties map[string]string `json:"uiProperties,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this flow
func (m *Flow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Flow) validateEdges(formats strfmt.Registry) error {

	if swag.IsZero(m.Edges) { // not required
		return nil
	}

	for i := 0; i < len(m.Edges); i++ {
		if swag.IsZero(m.Edges[i]) { // not required
			continue
		}

		if m.Edges[i] != nil {
			if err := m.Edges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("edges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Flow) validateNodes(formats strfmt.Registry) error {

	if swag.IsZero(m.Nodes) { // not required
		return nil
	}

	for i := 0; i < len(m.Nodes); i++ {
		if swag.IsZero(m.Nodes[i]) { // not required
			continue
		}

		if m.Nodes[i] != nil {
			if err := m.Nodes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Flow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Flow) UnmarshalBinary(b []byte) error {
	var res Flow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}