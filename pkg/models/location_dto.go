// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LocationDTO LocationDTO
//
// swagger:model LocationDTO
type LocationDTO struct {

	// city
	City string `json:"city,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// latitude
	Latitude string `json:"latitude,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// longitude
	Longitude string `json:"longitude,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this location d t o
func (m *LocationDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LocationDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocationDTO) UnmarshalBinary(b []byte) error {
	var res LocationDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
