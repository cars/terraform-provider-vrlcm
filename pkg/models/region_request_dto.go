// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RegionRequestDTO RegionRequestDTO
//
// swagger:model RegionRequestDTO
type RegionRequestDTO struct {

	// location
	Location string `json:"location,omitempty"`

	// region name
	RegionName string `json:"regionName,omitempty"`
}

// Validate validates this region request d t o
func (m *RegionRequestDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegionRequestDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegionRequestDTO) UnmarshalBinary(b []byte) error {
	var res RegionRequestDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
