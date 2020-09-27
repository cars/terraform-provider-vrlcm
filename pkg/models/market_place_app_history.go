// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MarketPlaceAppHistory MarketPlaceAppHistory
//
// swagger:model MarketPlaceAppHistory
type MarketPlaceAppHistory struct {

	// app version
	AppVersion string `json:"appVersion,omitempty"`

	// modified action
	ModifiedAction string `json:"modifiedAction,omitempty"`

	// modified by
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// modified date
	ModifiedDate string `json:"modifiedDate,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`
}

// Validate validates this market place app history
func (m *MarketPlaceAppHistory) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MarketPlaceAppHistory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarketPlaceAppHistory) UnmarshalBinary(b []byte) error {
	var res MarketPlaceAppHistory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
