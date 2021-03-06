// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VCenterRequestDTO VCenterRequestDTO
//
// swagger:model VCenterRequestDTO
type VCenterRequestDTO struct {

	// v center host
	VCenterHost string `json:"vCenterHost,omitempty"`

	// v center name
	VCenterName string `json:"vCenterName,omitempty"`

	// vc data collection status
	VcDataCollectionStatus string `json:"vcDataCollectionStatus,omitempty"`

	// vc password
	VcPassword string `json:"vcPassword,omitempty"`

	// vc used as
	VcUsedAs string `json:"vcUsedAs,omitempty"`

	// vc username
	VcUsername string `json:"vcUsername,omitempty"`
}

// Validate validates this v center request d t o
func (m *VCenterRequestDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VCenterRequestDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VCenterRequestDTO) UnmarshalBinary(b []byte) error {
	var res VCenterRequestDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
