// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MarketplaceAppInstallDTO MarketplaceAppInstallDTO
//
// swagger:model MarketplaceAppInstallDTO
type MarketplaceAppInstallDTO struct {

	// datacenter
	Datacenter string `json:"datacenter,omitempty"`

	// environment Id
	EnvironmentID string `json:"environmentId,omitempty"`

	// product type
	ProductType string `json:"productType,omitempty"`

	// product version
	ProductVersion string `json:"productVersion,omitempty"`

	// properties
	Properties map[string]string `json:"properties,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`
}

// Validate validates this marketplace app install d t o
func (m *MarketplaceAppInstallDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MarketplaceAppInstallDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarketplaceAppInstallDTO) UnmarshalBinary(b []byte) error {
	var res MarketplaceAppInstallDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
