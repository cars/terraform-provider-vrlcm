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
)

// NewShutDownProductVidmUsingPOSTParams creates a new ShutDownProductVidmUsingPOSTParams object
// with the default values initialized.
func NewShutDownProductVidmUsingPOSTParams() *ShutDownProductVidmUsingPOSTParams {
	var ()
	return &ShutDownProductVidmUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewShutDownProductVidmUsingPOSTParamsWithTimeout creates a new ShutDownProductVidmUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewShutDownProductVidmUsingPOSTParamsWithTimeout(timeout time.Duration) *ShutDownProductVidmUsingPOSTParams {
	var ()
	return &ShutDownProductVidmUsingPOSTParams{

		timeout: timeout,
	}
}

// NewShutDownProductVidmUsingPOSTParamsWithContext creates a new ShutDownProductVidmUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewShutDownProductVidmUsingPOSTParamsWithContext(ctx context.Context) *ShutDownProductVidmUsingPOSTParams {
	var ()
	return &ShutDownProductVidmUsingPOSTParams{

		Context: ctx,
	}
}

// NewShutDownProductVidmUsingPOSTParamsWithHTTPClient creates a new ShutDownProductVidmUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewShutDownProductVidmUsingPOSTParamsWithHTTPClient(client *http.Client) *ShutDownProductVidmUsingPOSTParams {
	var ()
	return &ShutDownProductVidmUsingPOSTParams{
		HTTPClient: client,
	}
}

/*ShutDownProductVidmUsingPOSTParams contains all the parameters to send to the API endpoint
for the shut down product vidm using p o s t operation typically these are written to a http.Request
*/
type ShutDownProductVidmUsingPOSTParams struct {

	/*EnvironmentID
	  environmentId

	*/
	EnvironmentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the shut down product vidm using p o s t params
func (o *ShutDownProductVidmUsingPOSTParams) WithTimeout(timeout time.Duration) *ShutDownProductVidmUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the shut down product vidm using p o s t params
func (o *ShutDownProductVidmUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the shut down product vidm using p o s t params
func (o *ShutDownProductVidmUsingPOSTParams) WithContext(ctx context.Context) *ShutDownProductVidmUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the shut down product vidm using p o s t params
func (o *ShutDownProductVidmUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the shut down product vidm using p o s t params
func (o *ShutDownProductVidmUsingPOSTParams) WithHTTPClient(client *http.Client) *ShutDownProductVidmUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the shut down product vidm using p o s t params
func (o *ShutDownProductVidmUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the shut down product vidm using p o s t params
func (o *ShutDownProductVidmUsingPOSTParams) WithEnvironmentID(environmentID string) *ShutDownProductVidmUsingPOSTParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the shut down product vidm using p o s t params
func (o *ShutDownProductVidmUsingPOSTParams) SetEnvironmentID(environmentID string) {
	o.EnvironmentID = environmentID
}

// WriteToRequest writes these params to a swagger request
func (o *ShutDownProductVidmUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param environmentId
	if err := r.SetPathParam("environmentId", o.EnvironmentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
