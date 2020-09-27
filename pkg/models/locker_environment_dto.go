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

// LockerEnvironmentDTO LockerEnvironmentDTO
//
// swagger:model LockerEnvironmentDTO
type LockerEnvironmentDTO struct {

	// env Id
	EnvID string `json:"envId,omitempty"`

	// env name
	EnvName string `json:"envName,omitempty"`

	// products
	Products []*LockerProductDTO `json:"products"`
}

// Validate validates this locker environment d t o
func (m *LockerEnvironmentDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProducts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LockerEnvironmentDTO) validateProducts(formats strfmt.Registry) error {

	if swag.IsZero(m.Products) { // not required
		return nil
	}

	for i := 0; i < len(m.Products); i++ {
		if swag.IsZero(m.Products[i]) { // not required
			continue
		}

		if m.Products[i] != nil {
			if err := m.Products[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("products" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LockerEnvironmentDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LockerEnvironmentDTO) UnmarshalBinary(b []byte) error {
	var res LockerEnvironmentDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
