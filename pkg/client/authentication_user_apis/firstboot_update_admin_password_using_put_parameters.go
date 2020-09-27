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

// NewFirstbootUpdateAdminPasswordUsingPUTParams creates a new FirstbootUpdateAdminPasswordUsingPUTParams object
// with the default values initialized.
func NewFirstbootUpdateAdminPasswordUsingPUTParams() *FirstbootUpdateAdminPasswordUsingPUTParams {
	var ()
	return &FirstbootUpdateAdminPasswordUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFirstbootUpdateAdminPasswordUsingPUTParamsWithTimeout creates a new FirstbootUpdateAdminPasswordUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFirstbootUpdateAdminPasswordUsingPUTParamsWithTimeout(timeout time.Duration) *FirstbootUpdateAdminPasswordUsingPUTParams {
	var ()
	return &FirstbootUpdateAdminPasswordUsingPUTParams{

		timeout: timeout,
	}
}

// NewFirstbootUpdateAdminPasswordUsingPUTParamsWithContext creates a new FirstbootUpdateAdminPasswordUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewFirstbootUpdateAdminPasswordUsingPUTParamsWithContext(ctx context.Context) *FirstbootUpdateAdminPasswordUsingPUTParams {
	var ()
	return &FirstbootUpdateAdminPasswordUsingPUTParams{

		Context: ctx,
	}
}

// NewFirstbootUpdateAdminPasswordUsingPUTParamsWithHTTPClient creates a new FirstbootUpdateAdminPasswordUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFirstbootUpdateAdminPasswordUsingPUTParamsWithHTTPClient(client *http.Client) *FirstbootUpdateAdminPasswordUsingPUTParams {
	var ()
	return &FirstbootUpdateAdminPasswordUsingPUTParams{
		HTTPClient: client,
	}
}

/*FirstbootUpdateAdminPasswordUsingPUTParams contains all the parameters to send to the API endpoint
for the firstboot update admin password using p u t operation typically these are written to a http.Request
*/
type FirstbootUpdateAdminPasswordUsingPUTParams struct {

	/*PasswordUpdateRequestDTO
	  passwordUpdateRequestDTO

	*/
	PasswordUpdateRequestDTO *models.LocalUserCredsDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the firstboot update admin password using p u t params
func (o *FirstbootUpdateAdminPasswordUsingPUTParams) WithTimeout(timeout time.Duration) *FirstbootUpdateAdminPasswordUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the firstboot update admin password using p u t params
func (o *FirstbootUpdateAdminPasswordUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the firstboot update admin password using p u t params
func (o *FirstbootUpdateAdminPasswordUsingPUTParams) WithContext(ctx context.Context) *FirstbootUpdateAdminPasswordUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the firstboot update admin password using p u t params
func (o *FirstbootUpdateAdminPasswordUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the firstboot update admin password using p u t params
func (o *FirstbootUpdateAdminPasswordUsingPUTParams) WithHTTPClient(client *http.Client) *FirstbootUpdateAdminPasswordUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the firstboot update admin password using p u t params
func (o *FirstbootUpdateAdminPasswordUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPasswordUpdateRequestDTO adds the passwordUpdateRequestDTO to the firstboot update admin password using p u t params
func (o *FirstbootUpdateAdminPasswordUsingPUTParams) WithPasswordUpdateRequestDTO(passwordUpdateRequestDTO *models.LocalUserCredsDTO) *FirstbootUpdateAdminPasswordUsingPUTParams {
	o.SetPasswordUpdateRequestDTO(passwordUpdateRequestDTO)
	return o
}

// SetPasswordUpdateRequestDTO adds the passwordUpdateRequestDTO to the firstboot update admin password using p u t params
func (o *FirstbootUpdateAdminPasswordUsingPUTParams) SetPasswordUpdateRequestDTO(passwordUpdateRequestDTO *models.LocalUserCredsDTO) {
	o.PasswordUpdateRequestDTO = passwordUpdateRequestDTO
}

// WriteToRequest writes these params to a swagger request
func (o *FirstbootUpdateAdminPasswordUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PasswordUpdateRequestDTO != nil {
		if err := r.SetBodyParam(o.PasswordUpdateRequestDTO); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}