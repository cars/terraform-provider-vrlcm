// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SyncProfileGroupDTO SyncProfileGroupDTO
//
// swagger:model SyncProfileGroupDTO
type SyncProfileGroupDTO struct {

	// exclude nested group members
	ExcludeNestedGroupMembers bool `json:"excludeNestedGroupMembers,omitempty"`

	// identity group info
	IdentityGroupInfo map[string]ADIdentityGroupInfoDTO `json:"identityGroupInfo,omitempty"`
}

// Validate validates this sync profile group d t o
func (m *SyncProfileGroupDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdentityGroupInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SyncProfileGroupDTO) validateIdentityGroupInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.IdentityGroupInfo) { // not required
		return nil
	}

	for k := range m.IdentityGroupInfo {

		if err := validate.Required("identityGroupInfo"+"."+k, "body", m.IdentityGroupInfo[k]); err != nil {
			return err
		}
		if val, ok := m.IdentityGroupInfo[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SyncProfileGroupDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyncProfileGroupDTO) UnmarshalBinary(b []byte) error {
	var res SyncProfileGroupDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
