// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LockerNodeDTO LockerNodeDTO
//
// swagger:model LockerNodeDTO
type LockerNodeDTO struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this locker node d t o
func (m *LockerNodeDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LockerNodeDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LockerNodeDTO) UnmarshalBinary(b []byte) error {
	var res LockerNodeDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}