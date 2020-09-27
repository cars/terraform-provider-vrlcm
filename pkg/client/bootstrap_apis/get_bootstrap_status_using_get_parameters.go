// Code generated by go-swagger; DO NOT EDIT.

package bootstrap_a_p_is

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

// NewGetBootstrapStatusUsingGETParams creates a new GetBootstrapStatusUsingGETParams object
// with the default values initialized.
func NewGetBootstrapStatusUsingGETParams() *GetBootstrapStatusUsingGETParams {

	return &GetBootstrapStatusUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBootstrapStatusUsingGETParamsWithTimeout creates a new GetBootstrapStatusUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBootstrapStatusUsingGETParamsWithTimeout(timeout time.Duration) *GetBootstrapStatusUsingGETParams {

	return &GetBootstrapStatusUsingGETParams{

		timeout: timeout,
	}
}

// NewGetBootstrapStatusUsingGETParamsWithContext creates a new GetBootstrapStatusUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBootstrapStatusUsingGETParamsWithContext(ctx context.Context) *GetBootstrapStatusUsingGETParams {

	return &GetBootstrapStatusUsingGETParams{

		Context: ctx,
	}
}

// NewGetBootstrapStatusUsingGETParamsWithHTTPClient creates a new GetBootstrapStatusUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBootstrapStatusUsingGETParamsWithHTTPClient(client *http.Client) *GetBootstrapStatusUsingGETParams {

	return &GetBootstrapStatusUsingGETParams{
		HTTPClient: client,
	}
}

/*GetBootstrapStatusUsingGETParams contains all the parameters to send to the API endpoint
for the get bootstrap status using g e t operation typically these are written to a http.Request
*/
type GetBootstrapStatusUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get bootstrap status using g e t params
func (o *GetBootstrapStatusUsingGETParams) WithTimeout(timeout time.Duration) *GetBootstrapStatusUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bootstrap status using g e t params
func (o *GetBootstrapStatusUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bootstrap status using g e t params
func (o *GetBootstrapStatusUsingGETParams) WithContext(ctx context.Context) *GetBootstrapStatusUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bootstrap status using g e t params
func (o *GetBootstrapStatusUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bootstrap status using g e t params
func (o *GetBootstrapStatusUsingGETParams) WithHTTPClient(client *http.Client) *GetBootstrapStatusUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bootstrap status using g e t params
func (o *GetBootstrapStatusUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetBootstrapStatusUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}