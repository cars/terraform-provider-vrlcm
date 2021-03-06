// Code generated by go-swagger; DO NOT EDIT.

package deployments_controller

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

// NewPrevalidateProductUpgradeRequestUsingPOSTParams creates a new PrevalidateProductUpgradeRequestUsingPOSTParams object
// with the default values initialized.
func NewPrevalidateProductUpgradeRequestUsingPOSTParams() *PrevalidateProductUpgradeRequestUsingPOSTParams {
	var ()
	return &PrevalidateProductUpgradeRequestUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPrevalidateProductUpgradeRequestUsingPOSTParamsWithTimeout creates a new PrevalidateProductUpgradeRequestUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPrevalidateProductUpgradeRequestUsingPOSTParamsWithTimeout(timeout time.Duration) *PrevalidateProductUpgradeRequestUsingPOSTParams {
	var ()
	return &PrevalidateProductUpgradeRequestUsingPOSTParams{

		timeout: timeout,
	}
}

// NewPrevalidateProductUpgradeRequestUsingPOSTParamsWithContext creates a new PrevalidateProductUpgradeRequestUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewPrevalidateProductUpgradeRequestUsingPOSTParamsWithContext(ctx context.Context) *PrevalidateProductUpgradeRequestUsingPOSTParams {
	var ()
	return &PrevalidateProductUpgradeRequestUsingPOSTParams{

		Context: ctx,
	}
}

// NewPrevalidateProductUpgradeRequestUsingPOSTParamsWithHTTPClient creates a new PrevalidateProductUpgradeRequestUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPrevalidateProductUpgradeRequestUsingPOSTParamsWithHTTPClient(client *http.Client) *PrevalidateProductUpgradeRequestUsingPOSTParams {
	var ()
	return &PrevalidateProductUpgradeRequestUsingPOSTParams{
		HTTPClient: client,
	}
}

/*PrevalidateProductUpgradeRequestUsingPOSTParams contains all the parameters to send to the API endpoint
for the prevalidate product upgrade request using p o s t operation typically these are written to a http.Request
*/
type PrevalidateProductUpgradeRequestUsingPOSTParams struct {

	/*UpgradeEnvironmentUIRequest
	  upgradeEnvironmentUiRequest

	*/
	UpgradeEnvironmentUIRequest *models.UpgradeEnvironmentUIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the prevalidate product upgrade request using p o s t params
func (o *PrevalidateProductUpgradeRequestUsingPOSTParams) WithTimeout(timeout time.Duration) *PrevalidateProductUpgradeRequestUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the prevalidate product upgrade request using p o s t params
func (o *PrevalidateProductUpgradeRequestUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the prevalidate product upgrade request using p o s t params
func (o *PrevalidateProductUpgradeRequestUsingPOSTParams) WithContext(ctx context.Context) *PrevalidateProductUpgradeRequestUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the prevalidate product upgrade request using p o s t params
func (o *PrevalidateProductUpgradeRequestUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the prevalidate product upgrade request using p o s t params
func (o *PrevalidateProductUpgradeRequestUsingPOSTParams) WithHTTPClient(client *http.Client) *PrevalidateProductUpgradeRequestUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the prevalidate product upgrade request using p o s t params
func (o *PrevalidateProductUpgradeRequestUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpgradeEnvironmentUIRequest adds the upgradeEnvironmentUIRequest to the prevalidate product upgrade request using p o s t params
func (o *PrevalidateProductUpgradeRequestUsingPOSTParams) WithUpgradeEnvironmentUIRequest(upgradeEnvironmentUIRequest *models.UpgradeEnvironmentUIRequest) *PrevalidateProductUpgradeRequestUsingPOSTParams {
	o.SetUpgradeEnvironmentUIRequest(upgradeEnvironmentUIRequest)
	return o
}

// SetUpgradeEnvironmentUIRequest adds the upgradeEnvironmentUiRequest to the prevalidate product upgrade request using p o s t params
func (o *PrevalidateProductUpgradeRequestUsingPOSTParams) SetUpgradeEnvironmentUIRequest(upgradeEnvironmentUIRequest *models.UpgradeEnvironmentUIRequest) {
	o.UpgradeEnvironmentUIRequest = upgradeEnvironmentUIRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PrevalidateProductUpgradeRequestUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UpgradeEnvironmentUIRequest != nil {
		if err := r.SetBodyParam(o.UpgradeEnvironmentUIRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
