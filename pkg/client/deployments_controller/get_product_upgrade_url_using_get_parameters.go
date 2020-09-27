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

// NewGetProductUpgradeURLUsingGETParams creates a new GetProductUpgradeURLUsingGETParams object
// with the default values initialized.
func NewGetProductUpgradeURLUsingGETParams() *GetProductUpgradeURLUsingGETParams {
	var ()
	return &GetProductUpgradeURLUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProductUpgradeURLUsingGETParamsWithTimeout creates a new GetProductUpgradeURLUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProductUpgradeURLUsingGETParamsWithTimeout(timeout time.Duration) *GetProductUpgradeURLUsingGETParams {
	var ()
	return &GetProductUpgradeURLUsingGETParams{

		timeout: timeout,
	}
}

// NewGetProductUpgradeURLUsingGETParamsWithContext creates a new GetProductUpgradeURLUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProductUpgradeURLUsingGETParamsWithContext(ctx context.Context) *GetProductUpgradeURLUsingGETParams {
	var ()
	return &GetProductUpgradeURLUsingGETParams{

		Context: ctx,
	}
}

// NewGetProductUpgradeURLUsingGETParamsWithHTTPClient creates a new GetProductUpgradeURLUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProductUpgradeURLUsingGETParamsWithHTTPClient(client *http.Client) *GetProductUpgradeURLUsingGETParams {
	var ()
	return &GetProductUpgradeURLUsingGETParams{
		HTTPClient: client,
	}
}

/*GetProductUpgradeURLUsingGETParams contains all the parameters to send to the API endpoint
for the get product upgrade Url using g e t operation typically these are written to a http.Request
*/
type GetProductUpgradeURLUsingGETParams struct {

	/*ProductID
	  productId

	*/
	ProductID string
	/*Version
	  version

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get product upgrade Url using g e t params
func (o *GetProductUpgradeURLUsingGETParams) WithTimeout(timeout time.Duration) *GetProductUpgradeURLUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get product upgrade Url using g e t params
func (o *GetProductUpgradeURLUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get product upgrade Url using g e t params
func (o *GetProductUpgradeURLUsingGETParams) WithContext(ctx context.Context) *GetProductUpgradeURLUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get product upgrade Url using g e t params
func (o *GetProductUpgradeURLUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get product upgrade Url using g e t params
func (o *GetProductUpgradeURLUsingGETParams) WithHTTPClient(client *http.Client) *GetProductUpgradeURLUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get product upgrade Url using g e t params
func (o *GetProductUpgradeURLUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProductID adds the productID to the get product upgrade Url using g e t params
func (o *GetProductUpgradeURLUsingGETParams) WithProductID(productID string) *GetProductUpgradeURLUsingGETParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the get product upgrade Url using g e t params
func (o *GetProductUpgradeURLUsingGETParams) SetProductID(productID string) {
	o.ProductID = productID
}

// WithVersion adds the version to the get product upgrade Url using g e t params
func (o *GetProductUpgradeURLUsingGETParams) WithVersion(version string) *GetProductUpgradeURLUsingGETParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get product upgrade Url using g e t params
func (o *GetProductUpgradeURLUsingGETParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductUpgradeURLUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param productId
	if err := r.SetPathParam("productId", o.ProductID); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}