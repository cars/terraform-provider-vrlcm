// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProductMetaDataDTO ProductMetaDataDTO
//
// swagger:model ProductMetaDataDTO
type ProductMetaDataDTO struct {

	// component list
	ComponentList map[string]ProductComponent `json:"componentList,omitempty"`

	// deployment type
	DeploymentType string `json:"deploymentType,omitempty"`

	// product name
	ProductName string `json:"productName,omitempty"`

	// product nodes
	ProductNodes []*UIMetaDataDTO `json:"productNodes"`

	// product properties
	ProductProperties []*UIMetaDataDTO `json:"productProperties"`

	// product version
	ProductVersion string `json:"productVersion,omitempty"`
}

// Validate validates this product meta data d t o
func (m *ProductMetaDataDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponentList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductNodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductMetaDataDTO) validateComponentList(formats strfmt.Registry) error {

	if swag.IsZero(m.ComponentList) { // not required
		return nil
	}

	for k := range m.ComponentList {

		if err := validate.Required("componentList"+"."+k, "body", m.ComponentList[k]); err != nil {
			return err
		}
		if val, ok := m.ComponentList[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ProductMetaDataDTO) validateProductNodes(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductNodes) { // not required
		return nil
	}

	for i := 0; i < len(m.ProductNodes); i++ {
		if swag.IsZero(m.ProductNodes[i]) { // not required
			continue
		}

		if m.ProductNodes[i] != nil {
			if err := m.ProductNodes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("productNodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProductMetaDataDTO) validateProductProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductProperties) { // not required
		return nil
	}

	for i := 0; i < len(m.ProductProperties); i++ {
		if swag.IsZero(m.ProductProperties[i]) { // not required
			continue
		}

		if m.ProductProperties[i] != nil {
			if err := m.ProductProperties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("productProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductMetaDataDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductMetaDataDTO) UnmarshalBinary(b []byte) error {
	var res ProductMetaDataDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
