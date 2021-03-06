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

// NewRetrustVidmWithProductUsingPOSTParams creates a new RetrustVidmWithProductUsingPOSTParams object
// with the default values initialized.
func NewRetrustVidmWithProductUsingPOSTParams() *RetrustVidmWithProductUsingPOSTParams {
	var ()
	return &RetrustVidmWithProductUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrustVidmWithProductUsingPOSTParamsWithTimeout creates a new RetrustVidmWithProductUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrustVidmWithProductUsingPOSTParamsWithTimeout(timeout time.Duration) *RetrustVidmWithProductUsingPOSTParams {
	var ()
	return &RetrustVidmWithProductUsingPOSTParams{

		timeout: timeout,
	}
}

// NewRetrustVidmWithProductUsingPOSTParamsWithContext creates a new RetrustVidmWithProductUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetrustVidmWithProductUsingPOSTParamsWithContext(ctx context.Context) *RetrustVidmWithProductUsingPOSTParams {
	var ()
	return &RetrustVidmWithProductUsingPOSTParams{

		Context: ctx,
	}
}

// NewRetrustVidmWithProductUsingPOSTParamsWithHTTPClient creates a new RetrustVidmWithProductUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrustVidmWithProductUsingPOSTParamsWithHTTPClient(client *http.Client) *RetrustVidmWithProductUsingPOSTParams {
	var ()
	return &RetrustVidmWithProductUsingPOSTParams{
		HTTPClient: client,
	}
}

/*RetrustVidmWithProductUsingPOSTParams contains all the parameters to send to the API endpoint
for the retrust vidm with product using p o s t operation typically these are written to a http.Request
*/
type RetrustVidmWithProductUsingPOSTParams struct {

	/*EnvironmentID
	  environmentId

	*/
	EnvironmentID string
	/*ProductID
	  productId

	*/
	ProductID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retrust vidm with product using p o s t params
func (o *RetrustVidmWithProductUsingPOSTParams) WithTimeout(timeout time.Duration) *RetrustVidmWithProductUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrust vidm with product using p o s t params
func (o *RetrustVidmWithProductUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrust vidm with product using p o s t params
func (o *RetrustVidmWithProductUsingPOSTParams) WithContext(ctx context.Context) *RetrustVidmWithProductUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrust vidm with product using p o s t params
func (o *RetrustVidmWithProductUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrust vidm with product using p o s t params
func (o *RetrustVidmWithProductUsingPOSTParams) WithHTTPClient(client *http.Client) *RetrustVidmWithProductUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrust vidm with product using p o s t params
func (o *RetrustVidmWithProductUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the retrust vidm with product using p o s t params
func (o *RetrustVidmWithProductUsingPOSTParams) WithEnvironmentID(environmentID string) *RetrustVidmWithProductUsingPOSTParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the retrust vidm with product using p o s t params
func (o *RetrustVidmWithProductUsingPOSTParams) SetEnvironmentID(environmentID string) {
	o.EnvironmentID = environmentID
}

// WithProductID adds the productID to the retrust vidm with product using p o s t params
func (o *RetrustVidmWithProductUsingPOSTParams) WithProductID(productID string) *RetrustVidmWithProductUsingPOSTParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the retrust vidm with product using p o s t params
func (o *RetrustVidmWithProductUsingPOSTParams) SetProductID(productID string) {
	o.ProductID = productID
}

// WriteToRequest writes these params to a swagger request
func (o *RetrustVidmWithProductUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param environmentId
	if err := r.SetPathParam("environmentId", o.EnvironmentID); err != nil {
		return err
	}

	// path param productId
	if err := r.SetPathParam("productId", o.ProductID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
