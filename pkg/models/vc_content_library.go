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

// VCContentLibrary VCContentLibrary
//
// swagger:model VCContentLibrary
type VCContentLibrary struct {

	// content library Id
	ContentLibraryID string `json:"contentLibraryId,omitempty"`

	// content library items
	ContentLibraryItems []*VCContentLibraryItem `json:"contentLibraryItems"`

	// content library name
	ContentLibraryName string `json:"contentLibraryName,omitempty"`
}

// Validate validates this v c content library
func (m *VCContentLibrary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentLibraryItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VCContentLibrary) validateContentLibraryItems(formats strfmt.Registry) error {

	if swag.IsZero(m.ContentLibraryItems) { // not required
		return nil
	}

	for i := 0; i < len(m.ContentLibraryItems); i++ {
		if swag.IsZero(m.ContentLibraryItems[i]) { // not required
			continue
		}

		if m.ContentLibraryItems[i] != nil {
			if err := m.ContentLibraryItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contentLibraryItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VCContentLibrary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VCContentLibrary) UnmarshalBinary(b []byte) error {
	var res VCContentLibrary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}