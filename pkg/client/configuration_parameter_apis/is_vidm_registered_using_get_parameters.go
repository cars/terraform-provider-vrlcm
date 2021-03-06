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
)

// NewIsVidmRegisteredUsingGETParams creates a new IsVidmRegisteredUsingGETParams object
// with the default values initialized.
func NewIsVidmRegisteredUsingGETParams() *IsVidmRegisteredUsingGETParams {

	return &IsVidmRegisteredUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIsVidmRegisteredUsingGETParamsWithTimeout creates a new IsVidmRegisteredUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIsVidmRegisteredUsingGETParamsWithTimeout(timeout time.Duration) *IsVidmRegisteredUsingGETParams {

	return &IsVidmRegisteredUsingGETParams{

		timeout: timeout,
	}
}

// NewIsVidmRegisteredUsingGETParamsWithContext creates a new IsVidmRegisteredUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewIsVidmRegisteredUsingGETParamsWithContext(ctx context.Context) *IsVidmRegisteredUsingGETParams {

	return &IsVidmRegisteredUsingGETParams{

		Context: ctx,
	}
}

// NewIsVidmRegisteredUsingGETParamsWithHTTPClient creates a new IsVidmRegisteredUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIsVidmRegisteredUsingGETParamsWithHTTPClient(client *http.Client) *IsVidmRegisteredUsingGETParams {

	return &IsVidmRegisteredUsingGETParams{
		HTTPClient: client,
	}
}

/*IsVidmRegisteredUsingGETParams contains all the parameters to send to the API endpoint
for the is vidm registered using g e t operation typically these are written to a http.Request
*/
type IsVidmRegisteredUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the is vidm registered using g e t params
func (o *IsVidmRegisteredUsingGETParams) WithTimeout(timeout time.Duration) *IsVidmRegisteredUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the is vidm registered using g e t params
func (o *IsVidmRegisteredUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the is vidm registered using g e t params
func (o *IsVidmRegisteredUsingGETParams) WithContext(ctx context.Context) *IsVidmRegisteredUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the is vidm registered using g e t params
func (o *IsVidmRegisteredUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the is vidm registered using g e t params
func (o *IsVidmRegisteredUsingGETParams) WithHTTPClient(client *http.Client) *IsVidmRegisteredUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the is vidm registered using g e t params
func (o *IsVidmRegisteredUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *IsVidmRegisteredUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
