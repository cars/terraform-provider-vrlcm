// Code generated by go-swagger; DO NOT EDIT.

package command_controller

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

// NewGetCommandsForEnvironmentUsingPOSTParams creates a new GetCommandsForEnvironmentUsingPOSTParams object
// with the default values initialized.
func NewGetCommandsForEnvironmentUsingPOSTParams() *GetCommandsForEnvironmentUsingPOSTParams {
	var ()
	return &GetCommandsForEnvironmentUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCommandsForEnvironmentUsingPOSTParamsWithTimeout creates a new GetCommandsForEnvironmentUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCommandsForEnvironmentUsingPOSTParamsWithTimeout(timeout time.Duration) *GetCommandsForEnvironmentUsingPOSTParams {
	var ()
	return &GetCommandsForEnvironmentUsingPOSTParams{

		timeout: timeout,
	}
}

// NewGetCommandsForEnvironmentUsingPOSTParamsWithContext creates a new GetCommandsForEnvironmentUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCommandsForEnvironmentUsingPOSTParamsWithContext(ctx context.Context) *GetCommandsForEnvironmentUsingPOSTParams {
	var ()
	return &GetCommandsForEnvironmentUsingPOSTParams{

		Context: ctx,
	}
}

// NewGetCommandsForEnvironmentUsingPOSTParamsWithHTTPClient creates a new GetCommandsForEnvironmentUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCommandsForEnvironmentUsingPOSTParamsWithHTTPClient(client *http.Client) *GetCommandsForEnvironmentUsingPOSTParams {
	var ()
	return &GetCommandsForEnvironmentUsingPOSTParams{
		HTTPClient: client,
	}
}

/*GetCommandsForEnvironmentUsingPOSTParams contains all the parameters to send to the API endpoint
for the get commands for environment using p o s t operation typically these are written to a http.Request
*/
type GetCommandsForEnvironmentUsingPOSTParams struct {

	/*EnvironmentCommandRequest
	  environmentCommandRequest

	*/
	EnvironmentCommandRequest *models.EnvironmentCommandRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get commands for environment using p o s t params
func (o *GetCommandsForEnvironmentUsingPOSTParams) WithTimeout(timeout time.Duration) *GetCommandsForEnvironmentUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get commands for environment using p o s t params
func (o *GetCommandsForEnvironmentUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get commands for environment using p o s t params
func (o *GetCommandsForEnvironmentUsingPOSTParams) WithContext(ctx context.Context) *GetCommandsForEnvironmentUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get commands for environment using p o s t params
func (o *GetCommandsForEnvironmentUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get commands for environment using p o s t params
func (o *GetCommandsForEnvironmentUsingPOSTParams) WithHTTPClient(client *http.Client) *GetCommandsForEnvironmentUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get commands for environment using p o s t params
func (o *GetCommandsForEnvironmentUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentCommandRequest adds the environmentCommandRequest to the get commands for environment using p o s t params
func (o *GetCommandsForEnvironmentUsingPOSTParams) WithEnvironmentCommandRequest(environmentCommandRequest *models.EnvironmentCommandRequest) *GetCommandsForEnvironmentUsingPOSTParams {
	o.SetEnvironmentCommandRequest(environmentCommandRequest)
	return o
}

// SetEnvironmentCommandRequest adds the environmentCommandRequest to the get commands for environment using p o s t params
func (o *GetCommandsForEnvironmentUsingPOSTParams) SetEnvironmentCommandRequest(environmentCommandRequest *models.EnvironmentCommandRequest) {
	o.EnvironmentCommandRequest = environmentCommandRequest
}

// WriteToRequest writes these params to a swagger request
func (o *GetCommandsForEnvironmentUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EnvironmentCommandRequest != nil {
		if err := r.SetBodyParam(o.EnvironmentCommandRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
