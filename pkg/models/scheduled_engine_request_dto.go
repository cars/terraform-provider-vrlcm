// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ScheduledEngineRequestDTO ScheduledEngineRequestDTO
//
// swagger:model ScheduledEngineRequestDTO
type ScheduledEngineRequestDTO struct {

	// interval
	Interval int64 `json:"interval,omitempty"`

	// is wait for previous request
	IsWaitForPreviousRequest string `json:"isWaitForPreviousRequest,omitempty"`

	// last execution
	LastExecution int64 `json:"lastExecution,omitempty"`

	// next execution
	NextExecution int64 `json:"nextExecution,omitempty"`

	// request
	Request *StateMachineInvocationRequest `json:"request,omitempty"`

	// request data
	RequestData *StateMachineInvocationRequest `json:"requestData,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// target Id
	TargetID string `json:"targetId,omitempty"`

	// transaction Id
	TransactionID string `json:"transactionId,omitempty"`

	// vmid
	Vmid string `json:"vmid,omitempty"`
}

// Validate validates this scheduled engine request d t o
func (m *ScheduledEngineRequestDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduledEngineRequestDTO) validateRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.Request) { // not required
		return nil
	}

	if m.Request != nil {
		if err := m.Request.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduledEngineRequestDTO) validateRequestData(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestData) { // not required
		return nil
	}

	if m.RequestData != nil {
		if err := m.RequestData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestData")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScheduledEngineRequestDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduledEngineRequestDTO) UnmarshalBinary(b []byte) error {
	var res ScheduledEngineRequestDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
