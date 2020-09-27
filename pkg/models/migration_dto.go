// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MigrationDto MigrationDto
//
// swagger:model MigrationDto
type MigrationDto struct {

	// is migrate contents
	IsMigrateContents string `json:"isMigrateContents,omitempty"`

	// is v ID m install
	IsVIDMInstall string `json:"isVIDMInstall,omitempty"`

	// legacy l c m hostname
	LegacyLCMHostname string `json:"legacyLCMHostname,omitempty"`

	// legacy l c m password
	LegacyLCMPassword string `json:"legacyLCMPassword,omitempty"`

	// legacy l c m SSH password
	LegacyLCMSSHPassword string `json:"legacyLCMSSHPassword,omitempty"`

	// legacy l c m username
	LegacyLCMUsername string `json:"legacyLCMUsername,omitempty"`

	// legacyv center hostname
	LegacyvCenterHostname string `json:"legacyvCenterHostname,omitempty"`

	// legacyv center password
	LegacyvCenterPassword string `json:"legacyvCenterPassword,omitempty"`

	// legacyv center username
	LegacyvCenterUsername string `json:"legacyvCenterUsername,omitempty"`

	// new l c m hostname
	NewLCMHostname string `json:"newLCMHostname,omitempty"`

	// new l c m password
	NewLCMPassword string `json:"newLCMPassword,omitempty"`

	// new l c m SSH password
	NewLCMSSHPassword string `json:"newLCMSSHPassword,omitempty"`

	// new l c m username
	NewLCMUsername string `json:"newLCMUsername,omitempty"`

	// newv center hostname
	NewvCenterHostname string `json:"newvCenterHostname,omitempty"`

	// newv center password
	NewvCenterPassword string `json:"newvCenterPassword,omitempty"`

	// newv center username
	NewvCenterUsername string `json:"newvCenterUsername,omitempty"`

	// transaction Id
	TransactionID string `json:"transactionId,omitempty"`

	// vmid
	Vmid string `json:"vmid,omitempty"`
}

// Validate validates this migration dto
func (m *MigrationDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MigrationDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MigrationDto) UnmarshalBinary(b []byte) error {
	var res MigrationDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
