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

// LockerProductDTO LockerProductDTO
//
// swagger:model LockerProductDTO
type LockerProductDTO struct {

	// consumption
	Consumption string `json:"consumption,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// nodes
	Nodes []*LockerNodeDTO `json:"nodes"`
}

// Validate validates this locker product d t o
func (m *LockerProductDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LockerProductDTO) validateNodes(formats strfmt.Registry) error {

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
func (m *LockerProductDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LockerProductDTO) UnmarshalBinary(b []byte) error {
	var res LockerProductDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
