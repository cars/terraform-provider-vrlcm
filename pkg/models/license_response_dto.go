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

// LicenseResponseDTO LicenseResponseDTO
//
// swagger:model LicenseResponseDTO
type LicenseResponseDTO struct {

	// end date
	EndDate string `json:"End Date,omitempty"`

	// start date
	StartDate string `json:"Start Date,omitempty"`

	// account
	Account string `json:"account,omitempty"`

	// alias
	Alias string `json:"alias,omitempty"`

	// cas Id
	CasID string `json:"casId,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// end date
	EndDate int64 `json:"endDate,omitempty"`

	// expired
	Expired bool `json:"expired,omitempty"`

	// expires in
	ExpiresIn string `json:"expiresIn,omitempty"`

	// expires on
	ExpiresOn string `json:"expiresOn,omitempty"`

	// falcon Id
	FalconID string `json:"falconId,omitempty"`

	// healthy
	Healthy bool `json:"healthy,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// last refreshed
	LastRefreshed int64 `json:"lastRefreshed,omitempty"`

	// lint Id
	LintID string `json:"lintId,omitempty"`

	// on prem allocation
	OnPremAllocation string `json:"onPremAllocation,omitempty"`

	// quantity
	Quantity int32 `json:"quantity,omitempty"`

	// refresh status
	RefreshStatus string `json:"refreshStatus,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// transaction Id
	TransactionID string `json:"transactionId,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// unit
	Unit string `json:"unit,omitempty"`

	// validations
	Validations []*ValidationResultDTO `json:"validations"`

	// vmid
	Vmid string `json:"vmid,omitempty"`
}

// Validate validates this license response d t o
func (m *LicenseResponseDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValidations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LicenseResponseDTO) validateValidations(formats strfmt.Registry) error {

	if swag.IsZero(m.Validations) { // not required
		return nil
	}

	for i := 0; i < len(m.Validations); i++ {
		if swag.IsZero(m.Validations[i]) { // not required
			continue
		}

		if m.Validations[i] != nil {
			if err := m.Validations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LicenseResponseDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicenseResponseDTO) UnmarshalBinary(b []byte) error {
	var res LicenseResponseDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
