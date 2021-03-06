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

// NewGetAllCofigsUsingGETParams creates a new GetAllCofigsUsingGETParams object
// with the default values initialized.
func NewGetAllCofigsUsingGETParams() *GetAllCofigsUsingGETParams {

	return &GetAllCofigsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllCofigsUsingGETParamsWithTimeout creates a new GetAllCofigsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllCofigsUsingGETParamsWithTimeout(timeout time.Duration) *GetAllCofigsUsingGETParams {

	return &GetAllCofigsUsingGETParams{

		timeout: timeout,
	}
}

// NewGetAllCofigsUsingGETParamsWithContext creates a new GetAllCofigsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllCofigsUsingGETParamsWithContext(ctx context.Context) *GetAllCofigsUsingGETParams {

	return &GetAllCofigsUsingGETParams{

		Context: ctx,
	}
}

// NewGetAllCofigsUsingGETParamsWithHTTPClient creates a new GetAllCofigsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllCofigsUsingGETParamsWithHTTPClient(client *http.Client) *GetAllCofigsUsingGETParams {

	return &GetAllCofigsUsingGETParams{
		HTTPClient: client,
	}
}

/*GetAllCofigsUsingGETParams contains all the parameters to send to the API endpoint
for the get all cofigs using g e t operation typically these are written to a http.Request
*/
type GetAllCofigsUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all cofigs using g e t params
func (o *GetAllCofigsUsingGETParams) WithTimeout(timeout time.Duration) *GetAllCofigsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all cofigs using g e t params
func (o *GetAllCofigsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all cofigs using g e t params
func (o *GetAllCofigsUsingGETParams) WithContext(ctx context.Context) *GetAllCofigsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all cofigs using g e t params
func (o *GetAllCofigsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all cofigs using g e t params
func (o *GetAllCofigsUsingGETParams) WithHTTPClient(client *http.Client) *GetAllCofigsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all cofigs using g e t params
func (o *GetAllCofigsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllCofigsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
