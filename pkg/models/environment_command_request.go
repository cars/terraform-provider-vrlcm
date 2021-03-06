// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EnvironmentCommandRequest EnvironmentCommandRequest
//
// swagger:model EnvironmentCommandRequest
type EnvironmentCommandRequest struct {

	// command type
	CommandType string `json:"commandType,omitempty"`

	// environment Id
	EnvironmentID string `json:"environmentId,omitempty"`
}

// Validate validates this environment command request
func (m *EnvironmentCommandRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentCommandRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentCommandRequest) UnmarshalBinary(b []byte) error {
	var res EnvironmentCommandRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
