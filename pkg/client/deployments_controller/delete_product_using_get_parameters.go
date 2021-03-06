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

// NewDeleteProductUsingGETParams creates a new DeleteProductUsingGETParams object
// with the default values initialized.
func NewDeleteProductUsingGETParams() *DeleteProductUsingGETParams {
	var ()
	return &DeleteProductUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProductUsingGETParamsWithTimeout creates a new DeleteProductUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteProductUsingGETParamsWithTimeout(timeout time.Duration) *DeleteProductUsingGETParams {
	var ()
	return &DeleteProductUsingGETParams{

		timeout: timeout,
	}
}

// NewDeleteProductUsingGETParamsWithContext creates a new DeleteProductUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteProductUsingGETParamsWithContext(ctx context.Context) *DeleteProductUsingGETParams {
	var ()
	return &DeleteProductUsingGETParams{

		Context: ctx,
	}
}

// NewDeleteProductUsingGETParamsWithHTTPClient creates a new DeleteProductUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteProductUsingGETParamsWithHTTPClient(client *http.Client) *DeleteProductUsingGETParams {
	var ()
	return &DeleteProductUsingGETParams{
		HTTPClient: client,
	}
}

/*DeleteProductUsingGETParams contains all the parameters to send to the API endpoint
for the delete product using g e t operation typically these are written to a http.Request
*/
type DeleteProductUsingGETParams struct {

	/*DeleteWindowsVMs
	  deleteWindowsVMs

	*/
	DeleteWindowsVMs *string
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

// WithTimeout adds the timeout to the delete product using g e t params
func (o *DeleteProductUsingGETParams) WithTimeout(timeout time.Duration) *DeleteProductUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete product using g e t params
func (o *DeleteProductUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete product using g e t params
func (o *DeleteProductUsingGETParams) WithContext(ctx context.Context) *DeleteProductUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete product using g e t params
func (o *DeleteProductUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete product using g e t params
func (o *DeleteProductUsingGETParams) WithHTTPClient(client *http.Client) *DeleteProductUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete product using g e t params
func (o *DeleteProductUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleteWindowsVMs adds the deleteWindowsVMs to the delete product using g e t params
func (o *DeleteProductUsingGETParams) WithDeleteWindowsVMs(deleteWindowsVMs *string) *DeleteProductUsingGETParams {
	o.SetDeleteWindowsVMs(deleteWindowsVMs)
	return o
}

// SetDeleteWindowsVMs adds the deleteWindowsVMs to the delete product using g e t params
func (o *DeleteProductUsingGETParams) SetDeleteWindowsVMs(deleteWindowsVMs *string) {
	o.DeleteWindowsVMs = deleteWindowsVMs
}

// WithEnvironmentID adds the environmentID to the delete product using g e t params
func (o *DeleteProductUsingGETParams) WithEnvironmentID(environmentID string) *DeleteProductUsingGETParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the delete product using g e t params
func (o *DeleteProductUsingGETParams) SetEnvironmentID(environmentID string) {
	o.EnvironmentID = environmentID
}

// WithProductID adds the productID to the delete product using g e t params
func (o *DeleteProductUsingGETParams) WithProductID(productID string) *DeleteProductUsingGETParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the delete product using g e t params
func (o *DeleteProductUsingGETParams) SetProductID(productID string) {
	o.ProductID = productID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProductUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DeleteWindowsVMs != nil {

		// query param deleteWindowsVMs
		var qrDeleteWindowsVMs string
		if o.DeleteWindowsVMs != nil {
			qrDeleteWindowsVMs = *o.DeleteWindowsVMs
		}
		qDeleteWindowsVMs := qrDeleteWindowsVMs
		if qDeleteWindowsVMs != "" {
			if err := r.SetQueryParam("deleteWindowsVMs", qDeleteWindowsVMs); err != nil {
				return err
			}
		}

	}

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
