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

// MachineRequestProgress MachineRequestProgress
//
// swagger:model MachineRequestProgress
type MachineRequestProgress struct {

	// current state
	CurrentState string `json:"currentState,omitempty"`

	// current state txt
	CurrentStateTxt string `json:"currentStateTxt,omitempty"`

	// event status details
	EventStatusDetails []*EventStatusMessage `json:"eventStatusDetails"`

	// last updated on
	LastUpdatedOn int64 `json:"lastUpdatedOn,omitempty"`

	// machine instance Id
	MachineInstanceID string `json:"machineInstanceId,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this machine request progress
func (m *MachineRequestProgress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventStatusDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MachineRequestProgress) validateEventStatusDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.EventStatusDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.EventStatusDetails); i++ {
		if swag.IsZero(m.EventStatusDetails[i]) { // not required
			continue
		}

		if m.EventStatusDetails[i] != nil {
			if err := m.EventStatusDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("eventStatusDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MachineRequestProgress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MachineRequestProgress) UnmarshalBinary(b []byte) error {
	var res MachineRequestProgress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}