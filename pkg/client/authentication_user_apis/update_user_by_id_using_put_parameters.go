// Code generated by go-swagger; DO NOT EDIT.

package authentication_user_a_p_is

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

// NewUpdateUserByIDUsingPUTParams creates a new UpdateUserByIDUsingPUTParams object
// with the default values initialized.
func NewUpdateUserByIDUsingPUTParams() *UpdateUserByIDUsingPUTParams {
	var ()
	return &UpdateUserByIDUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserByIDUsingPUTParamsWithTimeout creates a new UpdateUserByIDUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateUserByIDUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateUserByIDUsingPUTParams {
	var ()
	return &UpdateUserByIDUsingPUTParams{

		timeout: timeout,
	}
}

// NewUpdateUserByIDUsingPUTParamsWithContext creates a new UpdateUserByIDUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateUserByIDUsingPUTParamsWithContext(ctx context.Context) *UpdateUserByIDUsingPUTParams {
	var ()
	return &UpdateUserByIDUsingPUTParams{

		Context: ctx,
	}
}

// NewUpdateUserByIDUsingPUTParamsWithHTTPClient creates a new UpdateUserByIDUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateUserByIDUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateUserByIDUsingPUTParams {
	var ()
	return &UpdateUserByIDUsingPUTParams{
		HTTPClient: client,
	}
}

/*UpdateUserByIDUsingPUTParams contains all the parameters to send to the API endpoint
for the update user by Id using p u t operation typically these are written to a http.Request
*/
type UpdateUserByIDUsingPUTParams struct {

	/*UserRequestDTO
	  userRequestDTO

	*/
	UserRequestDTO *models.UserRequestDTO
	/*Vmid
	  vmid

	*/
	Vmid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update user by Id using p u t params
func (o *UpdateUserByIDUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateUserByIDUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user by Id using p u t params
func (o *UpdateUserByIDUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user by Id using p u t params
func (o *UpdateUserByIDUsingPUTParams) WithContext(ctx context.Context) *UpdateUserByIDUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user by Id using p u t params
func (o *UpdateUserByIDUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user by Id using p u t params
func (o *UpdateUserByIDUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateUserByIDUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user by Id using p u t params
func (o *UpdateUserByIDUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserRequestDTO adds the userRequestDTO to the update user by Id using p u t params
func (o *UpdateUserByIDUsingPUTParams) WithUserRequestDTO(userRequestDTO *models.UserRequestDTO) *UpdateUserByIDUsingPUTParams {
	o.SetUserRequestDTO(userRequestDTO)
	return o
}

// SetUserRequestDTO adds the userRequestDTO to the update user by Id using p u t params
func (o *UpdateUserByIDUsingPUTParams) SetUserRequestDTO(userRequestDTO *models.UserRequestDTO) {
	o.UserRequestDTO = userRequestDTO
}

// WithVmid adds the vmid to the update user by Id using p u t params
func (o *UpdateUserByIDUsingPUTParams) WithVmid(vmid string) *UpdateUserByIDUsingPUTParams {
	o.SetVmid(vmid)
	return o
}

// SetVmid adds the vmid to the update user by Id using p u t params
func (o *UpdateUserByIDUsingPUTParams) SetVmid(vmid string) {
	o.Vmid = vmid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserByIDUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UserRequestDTO != nil {
		if err := r.SetBodyParam(o.UserRequestDTO); err != nil {
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
