// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpgradeProductRequestDTO UpgradeProductRequestDTO
//
// swagger:model UpgradeProductRequestDTO
type UpgradeProductRequestDTO struct {

	// environment Id
	EnvironmentID string `json:"environmentId,omitempty"`

	// pre validate
	PreValidate bool `json:"preValidate,omitempty"`

	// product Id
	ProductID string `json:"productId,omitempty"`

	// product version
	ProductVersion string `json:"productVersion,omitempty"`

	// repository type
	RepositoryType string `json:"repositoryType,omitempty"`

	// repository Url
	RepositoryURL string `json:"repositoryUrl,omitempty"`

	// snapshot iaas after va upgrade
	SnapshotIaasAfterVaUpgrade bool `json:"snapshotIaasAfterVaUpgrade,omitempty"`
}

// Validate validates this upgrade product request d t o
func (m *UpgradeProductRequestDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpgradeProductRequestDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpgradeProductRequestDTO) UnmarshalBinary(b []byte) error {
	var res UpgradeProductRequestDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
