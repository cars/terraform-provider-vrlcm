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

// NewUpdateUserByIDUsingPATCHParams creates a new UpdateUserByIDUsingPATCHParams object
// with the default values initialized.
func NewUpdateUserByIDUsingPATCHParams() *UpdateUserByIDUsingPATCHParams {
	var ()
	return &UpdateUserByIDUsingPATCHParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserByIDUsingPATCHParamsWithTimeout creates a new UpdateUserByIDUsingPATCHParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateUserByIDUsingPATCHParamsWithTimeout(timeout time.Duration) *UpdateUserByIDUsingPATCHParams {
	var ()
	return &UpdateUserByIDUsingPATCHParams{

		timeout: timeout,
	}
}

// NewUpdateUserByIDUsingPATCHParamsWithContext creates a new UpdateUserByIDUsingPATCHParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateUserByIDUsingPATCHParamsWithContext(ctx context.Context) *UpdateUserByIDUsingPATCHParams {
	var ()
	return &UpdateUserByIDUsingPATCHParams{

		Context: ctx,
	}
}

// NewUpdateUserByIDUsingPATCHParamsWithHTTPClient creates a new UpdateUserByIDUsingPATCHParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateUserByIDUsingPATCHParamsWithHTTPClient(client *http.Client) *UpdateUserByIDUsingPATCHParams {
	var ()
	return &UpdateUserByIDUsingPATCHParams{
		HTTPClient: client,
	}
}

/*UpdateUserByIDUsingPATCHParams contains all the parameters to send to the API endpoint
for the update user by Id using p a t c h operation typically these are written to a http.Request
*/
type UpdateUserByIDUsingPATCHParams struct {

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

// WithTimeout adds the timeout to the update user by Id using p a t c h params
func (o *UpdateUserByIDUsingPATCHParams) WithTimeout(timeout time.Duration) *UpdateUserByIDUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user by Id using p a t c h params
func (o *UpdateUserByIDUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user by Id using p a t c h params
func (o *UpdateUserByIDUsingPATCHParams) WithContext(ctx context.Context) *UpdateUserByIDUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user by Id using p a t c h params
func (o *UpdateUserByIDUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user by Id using p a t c h params
func (o *UpdateUserByIDUsingPATCHParams) WithHTTPClient(client *http.Client) *UpdateUserByIDUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user by Id using p a t c h params
func (o *UpdateUserByIDUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserRequestDTO adds the userRequestDTO to the update user by Id using p a t c h params
func (o *UpdateUserByIDUsingPATCHParams) WithUserRequestDTO(userRequestDTO *models.UserRequestDTO) *UpdateUserByIDUsingPATCHParams {
	o.SetUserRequestDTO(userRequestDTO)
	return o
}

// SetUserRequestDTO adds the userRequestDTO to the update user by Id using p a t c h params
func (o *UpdateUserByIDUsingPATCHParams) SetUserRequestDTO(userRequestDTO *models.UserRequestDTO) {
	o.UserRequestDTO = userRequestDTO
}

// WithVmid adds the vmid to the update user by Id using p a t c h params
func (o *UpdateUserByIDUsingPATCHParams) WithVmid(vmid string) *UpdateUserByIDUsingPATCHParams {
	o.SetVmid(vmid)
	return o
}

// SetVmid adds the vmid to the update user by Id using p a t c h params
func (o *UpdateUserByIDUsingPATCHParams) SetVmid(vmid string) {
	o.Vmid = vmid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserByIDUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
