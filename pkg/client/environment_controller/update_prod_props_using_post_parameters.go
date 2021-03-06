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

// NewUpdateProdPropsUsingPOSTParams creates a new UpdateProdPropsUsingPOSTParams object
// with the default values initialized.
func NewUpdateProdPropsUsingPOSTParams() *UpdateProdPropsUsingPOSTParams {
	var ()
	return &UpdateProdPropsUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProdPropsUsingPOSTParamsWithTimeout creates a new UpdateProdPropsUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateProdPropsUsingPOSTParamsWithTimeout(timeout time.Duration) *UpdateProdPropsUsingPOSTParams {
	var ()
	return &UpdateProdPropsUsingPOSTParams{

		timeout: timeout,
	}
}

// NewUpdateProdPropsUsingPOSTParamsWithContext creates a new UpdateProdPropsUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateProdPropsUsingPOSTParamsWithContext(ctx context.Context) *UpdateProdPropsUsingPOSTParams {
	var ()
	return &UpdateProdPropsUsingPOSTParams{

		Context: ctx,
	}
}

// NewUpdateProdPropsUsingPOSTParamsWithHTTPClient creates a new UpdateProdPropsUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateProdPropsUsingPOSTParamsWithHTTPClient(client *http.Client) *UpdateProdPropsUsingPOSTParams {
	var ()
	return &UpdateProdPropsUsingPOSTParams{
		HTTPClient: client,
	}
}

/*UpdateProdPropsUsingPOSTParams contains all the parameters to send to the API endpoint
for the update prod props using p o s t operation typically these are written to a http.Request
*/
type UpdateProdPropsUsingPOSTParams struct {

	/*EnvironmentID
	  environmentId

	*/
	EnvironmentID string
	/*ProductID
	  productId

	*/
	ProductID string
	/*Props
	  props

	*/
	Props map[string]string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update prod props using p o s t params
func (o *UpdateProdPropsUsingPOSTParams) WithTimeout(timeout time.Duration) *UpdateProdPropsUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update prod props using p o s t params
func (o *UpdateProdPropsUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update prod props using p o s t params
func (o *UpdateProdPropsUsingPOSTParams) WithContext(ctx context.Context) *UpdateProdPropsUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update prod props using p o s t params
func (o *UpdateProdPropsUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update prod props using p o s t params
func (o *UpdateProdPropsUsingPOSTParams) WithHTTPClient(client *http.Client) *UpdateProdPropsUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update prod props using p o s t params
func (o *UpdateProdPropsUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the update prod props using p o s t params
func (o *UpdateProdPropsUsingPOSTParams) WithEnvironmentID(environmentID string) *UpdateProdPropsUsingPOSTParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the update prod props using p o s t params
func (o *UpdateProdPropsUsingPOSTParams) SetEnvironmentID(environmentID string) {
	o.EnvironmentID = environmentID
}

// WithProductID adds the productID to the update prod props using p o s t params
func (o *UpdateProdPropsUsingPOSTParams) WithProductID(productID string) *UpdateProdPropsUsingPOSTParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the update prod props using p o s t params
func (o *UpdateProdPropsUsingPOSTParams) SetProductID(productID string) {
	o.ProductID = productID
}

// WithProps adds the props to the update prod props using p o s t params
func (o *UpdateProdPropsUsingPOSTParams) WithProps(props map[string]string) *UpdateProdPropsUsingPOSTParams {
	o.SetProps(props)
	return o
}

// SetProps adds the props to the update prod props using p o s t params
func (o *UpdateProdPropsUsingPOSTParams) SetProps(props map[string]string) {
	o.Props = props
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProdPropsUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Props != nil {
		if err := r.SetBodyParam(o.Props); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
