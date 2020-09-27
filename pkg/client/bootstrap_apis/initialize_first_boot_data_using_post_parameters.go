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

// NewInitializeFirstBootDataUsingPOSTParams creates a new InitializeFirstBootDataUsingPOSTParams object
// with the default values initialized.
func NewInitializeFirstBootDataUsingPOSTParams() *InitializeFirstBootDataUsingPOSTParams {

	return &InitializeFirstBootDataUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInitializeFirstBootDataUsingPOSTParamsWithTimeout creates a new InitializeFirstBootDataUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInitializeFirstBootDataUsingPOSTParamsWithTimeout(timeout time.Duration) *InitializeFirstBootDataUsingPOSTParams {

	return &InitializeFirstBootDataUsingPOSTParams{

		timeout: timeout,
	}
}

// NewInitializeFirstBootDataUsingPOSTParamsWithContext creates a new InitializeFirstBootDataUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewInitializeFirstBootDataUsingPOSTParamsWithContext(ctx context.Context) *InitializeFirstBootDataUsingPOSTParams {

	return &InitializeFirstBootDataUsingPOSTParams{

		Context: ctx,
	}
}

// NewInitializeFirstBootDataUsingPOSTParamsWithHTTPClient creates a new InitializeFirstBootDataUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInitializeFirstBootDataUsingPOSTParamsWithHTTPClient(client *http.Client) *InitializeFirstBootDataUsingPOSTParams {

	return &InitializeFirstBootDataUsingPOSTParams{
		HTTPClient: client,
	}
}

/*InitializeFirstBootDataUsingPOSTParams contains all the parameters to send to the API endpoint
for the initialize first boot data using p o s t operation typically these are written to a http.Request
*/
type InitializeFirstBootDataUsingPOSTParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the initialize first boot data using p o s t params
func (o *InitializeFirstBootDataUsingPOSTParams) WithTimeout(timeout time.Duration) *InitializeFirstBootDataUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the initialize first boot data using p o s t params
func (o *InitializeFirstBootDataUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the initialize first boot data using p o s t params
func (o *InitializeFirstBootDataUsingPOSTParams) WithContext(ctx context.Context) *InitializeFirstBootDataUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the initialize first boot data using p o s t params
func (o *InitializeFirstBootDataUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the initialize first boot data using p o s t params
func (o *InitializeFirstBootDataUsingPOSTParams) WithHTTPClient(client *http.Client) *InitializeFirstBootDataUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the initialize first boot data using p o s t params
func (o *InitializeFirstBootDataUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *InitializeFirstBootDataUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}