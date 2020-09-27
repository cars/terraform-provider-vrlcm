// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CertificateImportRequestDTO CertificateImportRequestDTO
//
// swagger:model CertificateImportRequestDTO
type CertificateImportRequestDTO struct {

	// alias
	Alias string `json:"alias,omitempty"`

	// certificate chain
	CertificateChain string `json:"certificateChain,omitempty"`

	// passphrase
	Passphrase string `json:"passphrase,omitempty"`

	// private key
	PrivateKey string `json:"privateKey,omitempty"`
}

// Validate validates this certificate import request d t o
func (m *CertificateImportRequestDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificateImportRequestDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateImportRequestDTO) UnmarshalBinary(b []byte) error {
	var res CertificateImportRequestDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}