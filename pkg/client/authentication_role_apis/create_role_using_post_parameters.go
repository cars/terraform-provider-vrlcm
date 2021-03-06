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

// NewCreateRoleUsingPOSTParams creates a new CreateRoleUsingPOSTParams object
// with the default values initialized.
func NewCreateRoleUsingPOSTParams() *CreateRoleUsingPOSTParams {
	var ()
	return &CreateRoleUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRoleUsingPOSTParamsWithTimeout creates a new CreateRoleUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRoleUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateRoleUsingPOSTParams {
	var ()
	return &CreateRoleUsingPOSTParams{

		timeout: timeout,
	}
}

// NewCreateRoleUsingPOSTParamsWithContext creates a new CreateRoleUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRoleUsingPOSTParamsWithContext(ctx context.Context) *CreateRoleUsingPOSTParams {
	var ()
	return &CreateRoleUsingPOSTParams{

		Context: ctx,
	}
}

// NewCreateRoleUsingPOSTParamsWithHTTPClient creates a new CreateRoleUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRoleUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateRoleUsingPOSTParams {
	var ()
	return &CreateRoleUsingPOSTParams{
		HTTPClient: client,
	}
}

/*CreateRoleUsingPOSTParams contains all the parameters to send to the API endpoint
for the create role using p o s t operation typically these are written to a http.Request
*/
type CreateRoleUsingPOSTParams struct {

	/*RoleRequestDTO
	  roleRequestDTO

	*/
	RoleRequestDTO *models.RoleRequestDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create role using p o s t params
func (o *CreateRoleUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateRoleUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create role using p o s t params
func (o *CreateRoleUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create role using p o s t params
func (o *CreateRoleUsingPOSTParams) WithContext(ctx context.Context) *CreateRoleUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create role using p o s t params
func (o *CreateRoleUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create role using p o s t params
func (o *CreateRoleUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateRoleUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create role using p o s t params
func (o *CreateRoleUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoleRequestDTO adds the roleRequestDTO to the create role using p o s t params
func (o *CreateRoleUsingPOSTParams) WithRoleRequestDTO(roleRequestDTO *models.RoleRequestDTO) *CreateRoleUsingPOSTParams {
	o.SetRoleRequestDTO(roleRequestDTO)
	return o
}

// SetRoleRequestDTO adds the roleRequestDTO to the create role using p o s t params
func (o *CreateRoleUsingPOSTParams) SetRoleRequestDTO(roleRequestDTO *models.RoleRequestDTO) {
	o.RoleRequestDTO = roleRequestDTO
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRoleUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RoleRequestDTO != nil {
		if err := r.SetBodyParam(o.RoleRequestDTO); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
