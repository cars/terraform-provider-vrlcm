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

// NewGetUIEnvironmentRequestUsingGETParams creates a new GetUIEnvironmentRequestUsingGETParams object
// with the default values initialized.
func NewGetUIEnvironmentRequestUsingGETParams() *GetUIEnvironmentRequestUsingGETParams {
	var ()
	return &GetUIEnvironmentRequestUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUIEnvironmentRequestUsingGETParamsWithTimeout creates a new GetUIEnvironmentRequestUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUIEnvironmentRequestUsingGETParamsWithTimeout(timeout time.Duration) *GetUIEnvironmentRequestUsingGETParams {
	var ()
	return &GetUIEnvironmentRequestUsingGETParams{

		timeout: timeout,
	}
}

// NewGetUIEnvironmentRequestUsingGETParamsWithContext creates a new GetUIEnvironmentRequestUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUIEnvironmentRequestUsingGETParamsWithContext(ctx context.Context) *GetUIEnvironmentRequestUsingGETParams {
	var ()
	return &GetUIEnvironmentRequestUsingGETParams{

		Context: ctx,
	}
}

// NewGetUIEnvironmentRequestUsingGETParamsWithHTTPClient creates a new GetUIEnvironmentRequestUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUIEnvironmentRequestUsingGETParamsWithHTTPClient(client *http.Client) *GetUIEnvironmentRequestUsingGETParams {
	var ()
	return &GetUIEnvironmentRequestUsingGETParams{
		HTTPClient: client,
	}
}

/*GetUIEnvironmentRequestUsingGETParams contains all the parameters to send to the API endpoint
for the get Ui environment request using g e t operation typically these are written to a http.Request
*/
type GetUIEnvironmentRequestUsingGETParams struct {

	/*RequestID
	  requestId

	*/
	RequestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Ui environment request using g e t params
func (o *GetUIEnvironmentRequestUsingGETParams) WithTimeout(timeout time.Duration) *GetUIEnvironmentRequestUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Ui environment request using g e t params
func (o *GetUIEnvironmentRequestUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Ui environment request using g e t params
func (o *GetUIEnvironmentRequestUsingGETParams) WithContext(ctx context.Context) *GetUIEnvironmentRequestUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Ui environment request using g e t params
func (o *GetUIEnvironmentRequestUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Ui environment request using g e t params
func (o *GetUIEnvironmentRequestUsingGETParams) WithHTTPClient(client *http.Client) *GetUIEnvironmentRequestUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Ui environment request using g e t params
func (o *GetUIEnvironmentRequestUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestID adds the requestID to the get Ui environment request using g e t params
func (o *GetUIEnvironmentRequestUsingGETParams) WithRequestID(requestID string) *GetUIEnvironmentRequestUsingGETParams {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the get Ui environment request using g e t params
func (o *GetUIEnvironmentRequestUsingGETParams) SetRequestID(requestID string) {
	o.RequestID = requestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUIEnvironmentRequestUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param requestId
	if err := r.SetPathParam("requestId", o.RequestID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
