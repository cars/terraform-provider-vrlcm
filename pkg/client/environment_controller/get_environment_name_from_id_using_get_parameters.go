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
)

// NewGetEnvironmentNameFromIDUsingGETParams creates a new GetEnvironmentNameFromIDUsingGETParams object
// with the default values initialized.
func NewGetEnvironmentNameFromIDUsingGETParams() *GetEnvironmentNameFromIDUsingGETParams {
	var ()
	return &GetEnvironmentNameFromIDUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEnvironmentNameFromIDUsingGETParamsWithTimeout creates a new GetEnvironmentNameFromIDUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEnvironmentNameFromIDUsingGETParamsWithTimeout(timeout time.Duration) *GetEnvironmentNameFromIDUsingGETParams {
	var ()
	return &GetEnvironmentNameFromIDUsingGETParams{

		timeout: timeout,
	}
}

// NewGetEnvironmentNameFromIDUsingGETParamsWithContext creates a new GetEnvironmentNameFromIDUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEnvironmentNameFromIDUsingGETParamsWithContext(ctx context.Context) *GetEnvironmentNameFromIDUsingGETParams {
	var ()
	return &GetEnvironmentNameFromIDUsingGETParams{

		Context: ctx,
	}
}

// NewGetEnvironmentNameFromIDUsingGETParamsWithHTTPClient creates a new GetEnvironmentNameFromIDUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEnvironmentNameFromIDUsingGETParamsWithHTTPClient(client *http.Client) *GetEnvironmentNameFromIDUsingGETParams {
	var ()
	return &GetEnvironmentNameFromIDUsingGETParams{
		HTTPClient: client,
	}
}

/*GetEnvironmentNameFromIDUsingGETParams contains all the parameters to send to the API endpoint
for the get environment name from Id using g e t operation typically these are written to a http.Request
*/
type GetEnvironmentNameFromIDUsingGETParams struct {

	/*EnvironementID
	  environementId

	*/
	EnvironementID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get environment name from Id using g e t params
func (o *GetEnvironmentNameFromIDUsingGETParams) WithTimeout(timeout time.Duration) *GetEnvironmentNameFromIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get environment name from Id using g e t params
func (o *GetEnvironmentNameFromIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get environment name from Id using g e t params
func (o *GetEnvironmentNameFromIDUsingGETParams) WithContext(ctx context.Context) *GetEnvironmentNameFromIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get environment name from Id using g e t params
func (o *GetEnvironmentNameFromIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get environment name from Id using g e t params
func (o *GetEnvironmentNameFromIDUsingGETParams) WithHTTPClient(client *http.Client) *GetEnvironmentNameFromIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get environment name from Id using g e t params
func (o *GetEnvironmentNameFromIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironementID adds the environementID to the get environment name from Id using g e t params
func (o *GetEnvironmentNameFromIDUsingGETParams) WithEnvironementID(environementID *string) *GetEnvironmentNameFromIDUsingGETParams {
	o.SetEnvironementID(environementID)
	return o
}

// SetEnvironementID adds the environementId to the get environment name from Id using g e t params
func (o *GetEnvironmentNameFromIDUsingGETParams) SetEnvironementID(environementID *string) {
	o.EnvironementID = environementID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEnvironmentNameFromIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EnvironementID != nil {

		// query param environementId
		var qrEnvironementID string
		if o.EnvironementID != nil {
			qrEnvironementID = *o.EnvironementID
		}
		qEnvironementID := qrEnvironementID
		if qEnvironementID != "" {
			if err := r.SetQueryParam("environementId", qEnvironementID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
