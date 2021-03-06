// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VirtualMachineTemplateSpec VirtualMachineTemplateSpec
//
// swagger:model VirtualMachineTemplateSpec
type VirtualMachineTemplateSpec struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this virtual machine template spec
func (m *VirtualMachineTemplateSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VirtualMachineTemplateSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualMachineTemplateSpec) UnmarshalBinary(b []byte) error {
	var res VirtualMachineTemplateSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
