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

// NewCreateUserUsingPOSTParams creates a new CreateUserUsingPOSTParams object
// with the default values initialized.
func NewCreateUserUsingPOSTParams() *CreateUserUsingPOSTParams {
	var ()
	return &CreateUserUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUserUsingPOSTParamsWithTimeout creates a new CreateUserUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateUserUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateUserUsingPOSTParams {
	var ()
	return &CreateUserUsingPOSTParams{

		timeout: timeout,
	}
}

// NewCreateUserUsingPOSTParamsWithContext creates a new CreateUserUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateUserUsingPOSTParamsWithContext(ctx context.Context) *CreateUserUsingPOSTParams {
	var ()
	return &CreateUserUsingPOSTParams{

		Context: ctx,
	}
}

// NewCreateUserUsingPOSTParamsWithHTTPClient creates a new CreateUserUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateUserUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateUserUsingPOSTParams {
	var ()
	return &CreateUserUsingPOSTParams{
		HTTPClient: client,
	}
}

/*CreateUserUsingPOSTParams contains all the parameters to send to the API endpoint
for the create user using p o s t operation typically these are written to a http.Request
*/
type CreateUserUsingPOSTParams struct {

	/*UserRequestDTO
	  userRequestDTO

	*/
	UserRequestDTO *models.UserRequestDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create user using p o s t params
func (o *CreateUserUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateUserUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create user using p o s t params
func (o *CreateUserUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create user using p o s t params
func (o *CreateUserUsingPOSTParams) WithContext(ctx context.Context) *CreateUserUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create user using p o s t params
func (o *CreateUserUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create user using p o s t params
func (o *CreateUserUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateUserUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create user using p o s t params
func (o *CreateUserUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserRequestDTO adds the userRequestDTO to the create user using p o s t params
func (o *CreateUserUsingPOSTParams) WithUserRequestDTO(userRequestDTO *models.UserRequestDTO) *CreateUserUsingPOSTParams {
	o.SetUserRequestDTO(userRequestDTO)
	return o
}

// SetUserRequestDTO adds the userRequestDTO to the create user using p o s t params
func (o *CreateUserUsingPOSTParams) SetUserRequestDTO(userRequestDTO *models.UserRequestDTO) {
	o.UserRequestDTO = userRequestDTO
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUserUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UserRequestDTO != nil {
		if err := r.SetBodyParam(o.UserRequestDTO); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
