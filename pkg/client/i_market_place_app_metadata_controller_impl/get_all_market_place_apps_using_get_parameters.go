// Code generated by go-swagger; DO NOT EDIT.

package i_market_place_app_metadata_controller_impl

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

// NewGetAllMarketPlaceAppsUsingGETParams creates a new GetAllMarketPlaceAppsUsingGETParams object
// with the default values initialized.
func NewGetAllMarketPlaceAppsUsingGETParams() *GetAllMarketPlaceAppsUsingGETParams {

	return &GetAllMarketPlaceAppsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllMarketPlaceAppsUsingGETParamsWithTimeout creates a new GetAllMarketPlaceAppsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllMarketPlaceAppsUsingGETParamsWithTimeout(timeout time.Duration) *GetAllMarketPlaceAppsUsingGETParams {

	return &GetAllMarketPlaceAppsUsingGETParams{

		timeout: timeout,
	}
}

// NewGetAllMarketPlaceAppsUsingGETParamsWithContext creates a new GetAllMarketPlaceAppsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllMarketPlaceAppsUsingGETParamsWithContext(ctx context.Context) *GetAllMarketPlaceAppsUsingGETParams {

	return &GetAllMarketPlaceAppsUsingGETParams{

		Context: ctx,
	}
}

// NewGetAllMarketPlaceAppsUsingGETParamsWithHTTPClient creates a new GetAllMarketPlaceAppsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllMarketPlaceAppsUsingGETParamsWithHTTPClient(client *http.Client) *GetAllMarketPlaceAppsUsingGETParams {

	return &GetAllMarketPlaceAppsUsingGETParams{
		HTTPClient: client,
	}
}

/*GetAllMarketPlaceAppsUsingGETParams contains all the parameters to send to the API endpoint
for the get all market place apps using g e t operation typically these are written to a http.Request
*/
type GetAllMarketPlaceAppsUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all market place apps using g e t params
func (o *GetAllMarketPlaceAppsUsingGETParams) WithTimeout(timeout time.Duration) *GetAllMarketPlaceAppsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all market place apps using g e t params
func (o *GetAllMarketPlaceAppsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all market place apps using g e t params
func (o *GetAllMarketPlaceAppsUsingGETParams) WithContext(ctx context.Context) *GetAllMarketPlaceAppsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all market place apps using g e t params
func (o *GetAllMarketPlaceAppsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all market place apps using g e t params
func (o *GetAllMarketPlaceAppsUsingGETParams) WithHTTPClient(client *http.Client) *GetAllMarketPlaceAppsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all market place apps using g e t params
func (o *GetAllMarketPlaceAppsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllMarketPlaceAppsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
