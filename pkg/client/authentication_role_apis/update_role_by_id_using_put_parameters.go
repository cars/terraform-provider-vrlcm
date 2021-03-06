// Code generated by go-swagger; DO NOT EDIT.

package authentication_role_a_p_is

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cars/terraform-provider-vrlcm/models"
)

// NewUpdateRoleByIDUsingPUTParams creates a new UpdateRoleByIDUsingPUTParams object
// with the default values initialized.
func NewUpdateRoleByIDUsingPUTParams() *UpdateRoleByIDUsingPUTParams {
	var ()
	return &UpdateRoleByIDUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRoleByIDUsingPUTParamsWithTimeout creates a new UpdateRoleByIDUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateRoleByIDUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateRoleByIDUsingPUTParams {
	var ()
	return &UpdateRoleByIDUsingPUTParams{

		timeout: timeout,
	}
}

// NewUpdateRoleByIDUsingPUTParamsWithContext creates a new UpdateRoleByIDUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateRoleByIDUsingPUTParamsWithContext(ctx context.Context) *UpdateRoleByIDUsingPUTParams {
	var ()
	return &UpdateRoleByIDUsingPUTParams{

		Context: ctx,
	}
}

// NewUpdateRoleByIDUsingPUTParamsWithHTTPClient creates a new UpdateRoleByIDUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateRoleByIDUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateRoleByIDUsingPUTParams {
	var ()
	return &UpdateRoleByIDUsingPUTParams{
		HTTPClient: client,
	}
}

/*UpdateRoleByIDUsingPUTParams contains all the parameters to send to the API endpoint
for the update role by Id using p u t operation typically these are written to a http.Request
*/
type UpdateRoleByIDUsingPUTParams struct {

	/*RoleRequestDTO
	  roleRequestDTO

	*/
	RoleRequestDTO *models.RoleRequestDTO
	/*Vmid
	  vmid

	*/
	Vmid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update role by Id using p u t params
func (o *UpdateRoleByIDUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateRoleByIDUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update role by Id using p u t params
func (o *UpdateRoleByIDUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update role by Id using p u t params
func (o *UpdateRoleByIDUsingPUTParams) WithContext(ctx context.Context) *UpdateRoleByIDUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update role by Id using p u t params
func (o *UpdateRoleByIDUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update role by Id using p u t params
func (o *UpdateRoleByIDUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateRoleByIDUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update role by Id using p u t params
func (o *UpdateRoleByIDUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoleRequestDTO adds the roleRequestDTO to the update role by Id using p u t params
func (o *UpdateRoleByIDUsingPUTParams) WithRoleRequestDTO(roleRequestDTO *models.RoleRequestDTO) *UpdateRoleByIDUsingPUTParams {
	o.SetRoleRequestDTO(roleRequestDTO)
	return o
}

// SetRoleRequestDTO adds the roleRequestDTO to the update role by Id using p u t params
func (o *UpdateRoleByIDUsingPUTParams) SetRoleRequestDTO(roleRequestDTO *models.RoleRequestDTO) {
	o.RoleRequestDTO = roleRequestDTO
}

// WithVmid adds the vmid to the update role by Id using p u t params
func (o *UpdateRoleByIDUsingPUTParams) WithVmid(vmid string) *UpdateRoleByIDUsingPUTParams {
	o.SetVmid(vmid)
	return o
}

// SetVmid adds the vmid to the update role by Id using p u t params
func (o *UpdateRoleByIDUsingPUTParams) SetVmid(vmid string) {
	o.Vmid = vmid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRoleByIDUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RoleRequestDTO != nil {
		if err := r.SetBodyParam(o.RoleRequestDTO); err != nil {
			return err
		}
	}

	// path param vmid
	if err := r.SetPathParam("vmid", o.Vmid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
