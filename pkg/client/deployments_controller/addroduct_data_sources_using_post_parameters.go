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

	"github.com/cars/terraform-provider-vrlcm/models"
)

// NewAddroductDataSourcesUsingPOSTParams creates a new AddroductDataSourcesUsingPOSTParams object
// with the default values initialized.
func NewAddroductDataSourcesUsingPOSTParams() *AddroductDataSourcesUsingPOSTParams {
	var ()
	return &AddroductDataSourcesUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddroductDataSourcesUsingPOSTParamsWithTimeout creates a new AddroductDataSourcesUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddroductDataSourcesUsingPOSTParamsWithTimeout(timeout time.Duration) *AddroductDataSourcesUsingPOSTParams {
	var ()
	return &AddroductDataSourcesUsingPOSTParams{

		timeout: timeout,
	}
}

// NewAddroductDataSourcesUsingPOSTParamsWithContext creates a new AddroductDataSourcesUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddroductDataSourcesUsingPOSTParamsWithContext(ctx context.Context) *AddroductDataSourcesUsingPOSTParams {
	var ()
	return &AddroductDataSourcesUsingPOSTParams{

		Context: ctx,
	}
}

// NewAddroductDataSourcesUsingPOSTParamsWithHTTPClient creates a new AddroductDataSourcesUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddroductDataSourcesUsingPOSTParamsWithHTTPClient(client *http.Client) *AddroductDataSourcesUsingPOSTParams {
	var ()
	return &AddroductDataSourcesUsingPOSTParams{
		HTTPClient: client,
	}
}

/*AddroductDataSourcesUsingPOSTParams contains all the parameters to send to the API endpoint
for the addroduct data sources using p o s t operation typically these are written to a http.Request
*/
type AddroductDataSourcesUsingPOSTParams struct {

	/*EnvironmentID
	  environmentId

	*/
	EnvironmentID string
	/*Nodes
	  nodes

	*/
	Nodes []*models.Node
	/*ProductID
	  productId

	*/
	ProductID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the addroduct data sources using p o s t params
func (o *AddroductDataSourcesUsingPOSTParams) WithTimeout(timeout time.Duration) *AddroductDataSourcesUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the addroduct data sources using p o s t params
func (o *AddroductDataSourcesUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the addroduct data sources using p o s t params
func (o *AddroductDataSourcesUsingPOSTParams) WithContext(ctx context.Context) *AddroductDataSourcesUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the addroduct data sources using p o s t params
func (o *AddroductDataSourcesUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the addroduct data sources using p o s t params
func (o *AddroductDataSourcesUsingPOSTParams) WithHTTPClient(client *http.Client) *AddroductDataSourcesUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the addroduct data sources using p o s t params
func (o *AddroductDataSourcesUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the addroduct data sources using p o s t params
func (o *AddroductDataSourcesUsingPOSTParams) WithEnvironmentID(environmentID string) *AddroductDataSourcesUsingPOSTParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the addroduct data sources using p o s t params
func (o *AddroductDataSourcesUsingPOSTParams) SetEnvironmentID(environmentID string) {
	o.EnvironmentID = environmentID
}

// WithNodes adds the nodes to the addroduct data sources using p o s t params
func (o *AddroductDataSourcesUsingPOSTParams) WithNodes(nodes []*models.Node) *AddroductDataSourcesUsingPOSTParams {
	o.SetNodes(nodes)
	return o
}

// SetNodes adds the nodes to the addroduct data sources using p o s t params
func (o *AddroductDataSourcesUsingPOSTParams) SetNodes(nodes []*models.Node) {
	o.Nodes = nodes
}

// WithProductID adds the productID to the addroduct data sources using p o s t params
func (o *AddroductDataSourcesUsingPOSTParams) WithProductID(productID string) *AddroductDataSourcesUsingPOSTParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the addroduct data sources using p o s t params
func (o *AddroductDataSourcesUsingPOSTParams) SetProductID(productID string) {
	o.ProductID = productID
}

// WriteToRequest writes these params to a swagger request
func (o *AddroductDataSourcesUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param environmentId
	if err := r.SetPathParam("environmentId", o.EnvironmentID); err != nil {
		return err
	}

	if o.Nodes != nil {
		if err := r.SetBodyParam(o.Nodes); err != nil {
			return err
		}
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
