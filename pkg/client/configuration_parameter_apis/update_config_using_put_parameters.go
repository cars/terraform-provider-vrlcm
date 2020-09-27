// Code generated by go-swagger; DO NOT EDIT.

package configuration_parameter_a_p_is

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

// NewUpdateConfigUsingPUTParams creates a new UpdateConfigUsingPUTParams object
// with the default values initialized.
func NewUpdateConfigUsingPUTParams() *UpdateConfigUsingPUTParams {
	var ()
	return &UpdateConfigUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateConfigUsingPUTParamsWithTimeout creates a new UpdateConfigUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateConfigUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateConfigUsingPUTParams {
	var ()
	return &UpdateConfigUsingPUTParams{

		timeout: timeout,
	}
}

// NewUpdateConfigUsingPUTParamsWithContext creates a new UpdateConfigUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateConfigUsingPUTParamsWithContext(ctx context.Context) *UpdateConfigUsingPUTParams {
	var ()
	return &UpdateConfigUsingPUTParams{

		Context: ctx,
	}
}

// NewUpdateConfigUsingPUTParamsWithHTTPClient creates a new UpdateConfigUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateConfigUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateConfigUsingPUTParams {
	var ()
	return &UpdateConfigUsingPUTParams{
		HTTPClient: client,
	}
}

/*UpdateConfigUsingPUTParams contains all the parameters to send to the API endpoint
for the update config using p u t operation typically these are written to a http.Request
*/
type UpdateConfigUsingPUTParams struct {

	/*ConfigParamsDTO
	  configParamsDTO

	*/
	ConfigParamsDTO *models.ConfigParamsDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update config using p u t params
func (o *UpdateConfigUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateConfigUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update config using p u t params
func (o *UpdateConfigUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update config using p u t params
func (o *UpdateConfigUsingPUTParams) WithContext(ctx context.Context) *UpdateConfigUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update config using p u t params
func (o *UpdateConfigUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update config using p u t params
func (o *UpdateConfigUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateConfigUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update config using p u t params
func (o *UpdateConfigUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigParamsDTO adds the configParamsDTO to the update config using p u t params
func (o *UpdateConfigUsingPUTParams) WithConfigParamsDTO(configParamsDTO *models.ConfigParamsDTO) *UpdateConfigUsingPUTParams {
	o.SetConfigParamsDTO(configParamsDTO)
	return o
}

// SetConfigParamsDTO adds the configParamsDTO to the update config using p u t params
func (o *UpdateConfigUsingPUTParams) SetConfigParamsDTO(configParamsDTO *models.ConfigParamsDTO) {
	o.ConfigParamsDTO = configParamsDTO
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateConfigUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ConfigParamsDTO != nil {
		if err := r.SetBodyParam(o.ConfigParamsDTO); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
