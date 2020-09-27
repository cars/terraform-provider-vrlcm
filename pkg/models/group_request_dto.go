// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GroupRequestDTO GroupRequestDTO
//
// swagger:model GroupRequestDTO
type GroupRequestDTO struct {

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// domain
	Domain string `json:"domain,omitempty"`

	// group metadata
	GroupMetadata *GroupMetadataDTO `json:"groupMetadata,omitempty"`

	// group type
	GroupType string `json:"groupType,omitempty"`

	// is disabled
	IsDisabled bool `json:"isDisabled,omitempty"`

	// mapped roles
	MappedRoles []string `json:"mappedRoles"`

	// provider identifier
	ProviderIdentifier string `json:"providerIdentifier,omitempty"`
}

// Validate validates this group request d t o
func (m *GroupRequestDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupRequestDTO) validateGroupMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupMetadata) { // not required
		return nil
	}

	if m.GroupMetadata != nil {
		if err := m.GroupMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("groupMetadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupRequestDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupRequestDTO) UnmarshalBinary(b []byte) error {
	var res GroupRequestDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
