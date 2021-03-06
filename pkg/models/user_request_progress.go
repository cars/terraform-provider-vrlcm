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

// UserRequestProgress UserRequestProgress
//
// swagger:model UserRequestProgress
type UserRequestProgress struct {

	// engine request progress details
	EngineRequestProgressDetails []*EngineRequestProgress `json:"engineRequestProgressDetails"`

	// last updated on
	LastUpdatedOn int64 `json:"lastUpdatedOn,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// vmid
	Vmid string `json:"vmid,omitempty"`
}

// Validate validates this user request progress
func (m *UserRequestProgress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEngineRequestProgressDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserRequestProgress) validateEngineRequestProgressDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.EngineRequestProgressDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.EngineRequestProgressDetails); i++ {
		if swag.IsZero(m.EngineRequestProgressDetails[i]) { // not required
			continue
		}

		if m.EngineRequestProgressDetails[i] != nil {
			if err := m.EngineRequestProgressDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("engineRequestProgressDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserRequestProgress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserRequestProgress) UnmarshalBinary(b []byte) error {
	var res UserRequestProgress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
