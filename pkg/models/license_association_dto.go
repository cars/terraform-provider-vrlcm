// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LicenseAssociationDTO LicenseAssociationDTO
//
// swagger:model LicenseAssociationDTO
type LicenseAssociationDTO struct {

	// alias map
	AliasMap map[string]string `json:"aliasMap,omitempty"`

	// product Id
	ProductID string `json:"productId,omitempty"`

	// product version
	ProductVersion string `json:"productVersion,omitempty"`
}

// Validate validates this license association d t o
func (m *LicenseAssociationDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LicenseAssociationDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicenseAssociationDTO) UnmarshalBinary(b []byte) error {
	var res LicenseAssociationDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}