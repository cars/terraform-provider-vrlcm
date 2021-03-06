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

// CertificatePageDTO CertificatePageDTO
//
// swagger:model CertificatePageDTO
type CertificatePageDTO struct {

	// certificates
	Certificates []*CertificateMetaResponseDTO `json:"certificates"`

	// page
	Page int64 `json:"page,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this certificate page d t o
func (m *CertificatePageDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificatePageDTO) validateCertificates(formats strfmt.Registry) error {

	if swag.IsZero(m.Certificates) { // not required
		return nil
	}

	for i := 0; i < len(m.Certificates); i++ {
		if swag.IsZero(m.Certificates[i]) { // not required
			continue
		}

		if m.Certificates[i] != nil {
			if err := m.Certificates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("certificates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificatePageDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificatePageDTO) UnmarshalBinary(b []byte) error {
	var res CertificatePageDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
