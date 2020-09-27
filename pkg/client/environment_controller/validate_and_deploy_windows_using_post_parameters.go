// Code generated by go-swagger; DO NOT EDIT.

package environment_controller

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

// NewValidateAndDeployWindowsUsingPOSTParams creates a new ValidateAndDeployWindowsUsingPOSTParams object
// with the default values initialized.
func NewValidateAndDeployWindowsUsingPOSTParams() *ValidateAndDeployWindowsUsingPOSTParams {
	var ()
	return &ValidateAndDeployWindowsUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateAndDeployWindowsUsingPOSTParamsWithTimeout creates a new ValidateAndDeployWindowsUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateAndDeployWindowsUsingPOSTParamsWithTimeout(timeout time.Duration) *ValidateAndDeployWindowsUsingPOSTParams {
	var ()
	return &ValidateAndDeployWindowsUsingPOSTParams{

		timeout: timeout,
	}
}

// NewValidateAndDeployWindowsUsingPOSTParamsWithContext creates a new ValidateAndDeployWindowsUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateAndDeployWindowsUsingPOSTParamsWithContext(ctx context.Context) *ValidateAndDeployWindowsUsingPOSTParams {
	var ()
	return &ValidateAndDeployWindowsUsingPOSTParams{

		Context: ctx,
	}
}

// NewValidateAndDeployWindowsUsingPOSTParamsWithHTTPClient creates a new ValidateAndDeployWindowsUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateAndDeployWindowsUsingPOSTParamsWithHTTPClient(client *http.Client) *ValidateAndDeployWindowsUsingPOSTParams {
	var ()
	return &ValidateAndDeployWindowsUsingPOSTParams{
		HTTPClient: client,
	}
}

/*ValidateAndDeployWindowsUsingPOSTParams contains all the parameters to send to the API endpoint
for the validate and deploy windows using p o s t operation typically these are written to a http.Request
*/
type ValidateAndDeployWindowsUsingPOSTParams struct {

	/*CreateEnvironmentUIRequest
	  createEnvironmentUiRequest

	*/
	CreateEnvironmentUIRequest *models.CreateEnvironmentUIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate and deploy windows using p o s t params
func (o *ValidateAndDeployWindowsUsingPOSTParams) WithTimeout(timeout time.Duration) *ValidateAndDeployWindowsUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate and deploy windows using p o s t params
func (o *ValidateAndDeployWindowsUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate and deploy windows using p o s t params
func (o *ValidateAndDeployWindowsUsingPOSTParams) WithContext(ctx context.Context) *ValidateAndDeployWindowsUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate and deploy windows using p o s t params
func (o *ValidateAndDeployWindowsUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate and deploy windows using p o s t params
func (o *ValidateAndDeployWindowsUsingPOSTParams) WithHTTPClient(client *http.Client) *ValidateAndDeployWindowsUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate and deploy windows using p o s t params
func (o *ValidateAndDeployWindowsUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateEnvironmentUIRequest adds the createEnvironmentUIRequest to the validate and deploy windows using p o s t params
func (o *ValidateAndDeployWindowsUsingPOSTParams) WithCreateEnvironmentUIRequest(createEnvironmentUIRequest *models.CreateEnvironmentUIRequest) *ValidateAndDeployWindowsUsingPOSTParams {
	o.SetCreateEnvironmentUIRequest(createEnvironmentUIRequest)
	return o
}

// SetCreateEnvironmentUIRequest adds the createEnvironmentUiRequest to the validate and deploy windows using p o s t params
func (o *ValidateAndDeployWindowsUsingPOSTParams) SetCreateEnvironmentUIRequest(createEnvironmentUIRequest *models.CreateEnvironmentUIRequest) {
	o.CreateEnvironmentUIRequest = createEnvironmentUIRequest
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateAndDeployWindowsUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateEnvironmentUIRequest != nil {
		if err := r.SetBodyParam(o.CreateEnvironmentUIRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}