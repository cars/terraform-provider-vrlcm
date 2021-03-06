// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LockerRefBulkPostResponseDTO LockerRefBulkPostResponseDTO
//
// swagger:model LockerRefBulkPostResponseDTO
type LockerRefBulkPostResponseDTO struct {

	// destination name
	DestinationName string `json:"destinationName,omitempty"`

	// env Id
	EnvID string `json:"envId,omitempty"`

	// locker Id
	LockerID string `json:"lockerId,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// transaction Id
	TransactionID string `json:"transactionId,omitempty"`

	// vmid
	Vmid string `json:"vmid,omitempty"`
}

// Validate validates this locker ref bulk post response d t o
func (m *LockerRefBulkPostResponseDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LockerRefBulkPostResponseDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LockerRefBulkPostResponseDTO) UnmarshalBinary(b []byte) error {
	var res LockerRefBulkPostResponseDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
