// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateProductCertificateRequestDTO UpdateProductCertificateRequestDTO
//
// swagger:model UpdateProductCertificateRequestDTO
type UpdateProductCertificateRequestDTO struct {

	// certificate chain
	CertificateChain string `json:"certificateChain,omitempty"`

	// certificate Vm Id
	CertificateVMID string `json:"certificateVmId,omitempty"`

	// components
	Components []string `json:"components"`

	// private key
	PrivateKey string `json:"privateKey,omitempty"`

	// retrust prod meta
	RetrustProdMeta *RetrustProductMetaDTO `json:"retrustProdMeta,omitempty"`

	// snapshot enabled
	SnapshotEnabled bool `json:"snapshotEnabled,omitempty"`
}

// Validate validates this update product certificate request d t o
func (m *UpdateProductCertificateRequestDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRetrustProdMeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateProductCertificateRequestDTO) validateRetrustProdMeta(formats strfmt.Registry) error {

	if swag.IsZero(m.RetrustProdMeta) { // not required
		return nil
	}

	if m.RetrustProdMeta != nil {
		if err := m.RetrustProdMeta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retrustProdMeta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateProductCertificateRequestDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateProductCertificateRequestDTO) UnmarshalBinary(b []byte) error {
	var res UpdateProductCertificateRequestDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
