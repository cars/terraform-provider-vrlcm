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

// NewSyncProductIntegrationsUsingPOSTParams creates a new SyncProductIntegrationsUsingPOSTParams object
// with the default values initialized.
func NewSyncProductIntegrationsUsingPOSTParams() *SyncProductIntegrationsUsingPOSTParams {
	var ()
	return &SyncProductIntegrationsUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSyncProductIntegrationsUsingPOSTParamsWithTimeout creates a new SyncProductIntegrationsUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSyncProductIntegrationsUsingPOSTParamsWithTimeout(timeout time.Duration) *SyncProductIntegrationsUsingPOSTParams {
	var ()
	return &SyncProductIntegrationsUsingPOSTParams{

		timeout: timeout,
	}
}

// NewSyncProductIntegrationsUsingPOSTParamsWithContext creates a new SyncProductIntegrationsUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewSyncProductIntegrationsUsingPOSTParamsWithContext(ctx context.Context) *SyncProductIntegrationsUsingPOSTParams {
	var ()
	return &SyncProductIntegrationsUsingPOSTParams{

		Context: ctx,
	}
}

// NewSyncProductIntegrationsUsingPOSTParamsWithHTTPClient creates a new SyncProductIntegrationsUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSyncProductIntegrationsUsingPOSTParamsWithHTTPClient(client *http.Client) *SyncProductIntegrationsUsingPOSTParams {
	var ()
	return &SyncProductIntegrationsUsingPOSTParams{
		HTTPClient: client,
	}
}

/*SyncProductIntegrationsUsingPOSTParams contains all the parameters to send to the API endpoint
for the sync product integrations using p o s t operation typically these are written to a http.Request
*/
type SyncProductIntegrationsUsingPOSTParams struct {

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

// WithTimeout adds the timeout to the sync product integrations using p o s t params
func (o *SyncProductIntegrationsUsingPOSTParams) WithTimeout(timeout time.Duration) *SyncProductIntegrationsUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sync product integrations using p o s t params
func (o *SyncProductIntegrationsUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sync product integrations using p o s t params
func (o *SyncProductIntegrationsUsingPOSTParams) WithContext(ctx context.Context) *SyncProductIntegrationsUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sync product integrations using p o s t params
func (o *SyncProductIntegrationsUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sync product integrations using p o s t params
func (o *SyncProductIntegrationsUsingPOSTParams) WithHTTPClient(client *http.Client) *SyncProductIntegrationsUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sync product integrations using p o s t params
func (o *SyncProductIntegrationsUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the sync product integrations using p o s t params
func (o *SyncProductIntegrationsUsingPOSTParams) WithEnvironmentID(environmentID string) *SyncProductIntegrationsUsingPOSTParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the sync product integrations using p o s t params
func (o *SyncProductIntegrationsUsingPOSTParams) SetEnvironmentID(environmentID string) {
	o.EnvironmentID = environmentID
}

// WithProductID adds the productID to the sync product integrations using p o s t params
func (o *SyncProductIntegrationsUsingPOSTParams) WithProductID(productID string) *SyncProductIntegrationsUsingPOSTParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the sync product integrations using p o s t params
func (o *SyncProductIntegrationsUsingPOSTParams) SetProductID(productID string) {
	o.ProductID = productID
}

// WriteToRequest writes these params to a swagger request
func (o *SyncProductIntegrationsUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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