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
	"github.com/go-openapi/swag"
)

// NewGetCompatablityMatrixUsingGETParams creates a new GetCompatablityMatrixUsingGETParams object
// with the default values initialized.
func NewGetCompatablityMatrixUsingGETParams() *GetCompatablityMatrixUsingGETParams {
	var ()
	return &GetCompatablityMatrixUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCompatablityMatrixUsingGETParamsWithTimeout creates a new GetCompatablityMatrixUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCompatablityMatrixUsingGETParamsWithTimeout(timeout time.Duration) *GetCompatablityMatrixUsingGETParams {
	var ()
	return &GetCompatablityMatrixUsingGETParams{

		timeout: timeout,
	}
}

// NewGetCompatablityMatrixUsingGETParamsWithContext creates a new GetCompatablityMatrixUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCompatablityMatrixUsingGETParamsWithContext(ctx context.Context) *GetCompatablityMatrixUsingGETParams {
	var ()
	return &GetCompatablityMatrixUsingGETParams{

		Context: ctx,
	}
}

// NewGetCompatablityMatrixUsingGETParamsWithHTTPClient creates a new GetCompatablityMatrixUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCompatablityMatrixUsingGETParamsWithHTTPClient(client *http.Client) *GetCompatablityMatrixUsingGETParams {
	var ()
	return &GetCompatablityMatrixUsingGETParams{
		HTTPClient: client,
	}
}

/*GetCompatablityMatrixUsingGETParams contains all the parameters to send to the API endpoint
for the get compatablity matrix using g e t operation typically these are written to a http.Request
*/
type GetCompatablityMatrixUsingGETParams struct {

	/*IsProxyEnabled*/
	IsProxyEnabled *string
	/*Product
	  product

	*/
	Product string
	/*ProxyHost*/
	ProxyHost *string
	/*ProxyPassword*/
	ProxyPassword *string
	/*ProxyPort*/
	ProxyPort *int32
	/*ProxyUserName*/
	ProxyUserName *string
	/*Version
	  version

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get compatablity matrix using g e t params
func (o *GetCompatablityMatrixUsingGETParams) WithTimeout(timeout time.Duration) *GetCompatablityMatrixUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get compatablity matrix using g e t params
func (o *GetCompatablityMatrixUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get compatablity matrix using g e t params
func (o *GetCompatablityMatrixUsingGETParams) WithContext(ctx context.Context) *GetCompatablityMatrixUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get compatablity matrix using g e t params
func (o *GetCompatablityMatrixUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get compatablity matrix using g e t params
func (o *GetCompatablityMatrixUsingGETParams) WithHTTPClient(client *http.Client) *GetCompatablityMatrixUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get compatablity matrix using g e t params
func (o *GetCompatablityMatrixUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIsProxyEnabled adds the isProxyEnabled to the get compatablity matrix using g e t params
func (o *GetCompatablityMatrixUsingGETParams) WithIsProxyEnabled(isProxyEnabled *string) *GetCompatablityMatrixUsingGETParams {
	o.SetIsProxyEnabled(isProxyEnabled)
	return o
}

// SetIsProxyEnabled adds the isProxyEnabled to the get compatablity matrix using g e t params
func (o *GetCompatablityMatrixUsingGETParams) SetIsProxyEnabled(isProxyEnabled *string) {
	o.IsProxyEnabled = isProxyEnabled
}

// WithProduct adds the product to the get compatablity matrix using g e t params
func (o *GetCompatablityMatrixUsingGETParams) WithProduct(product string) *GetCompatablityMatrixUsingGETParams {
	o.SetProduct(product)
	return o
}

// SetProduct adds the product to the get compatablity matrix using g e t params
func (o *GetCompatablityMatrixUsingGETParams) SetProduct(product string) {
	o.Product = product
}

// WithProxyHost adds the proxyHost to the get compatablity matrix using g e t params
func (o *GetCompatablityMatrixUsingGETParams) WithProxyHost(proxyHost *string) *GetCompatablityMatrixUsingGETParams {
	o.SetProxyHost(proxyHost)
	return o
}

// SetProxyHost adds the proxyHost to the get compatablity matrix using g e t params
func (o *GetCompatablityMatrixUsingGETParams) SetProxyHost(proxyHost *string) {
	o.ProxyHost = proxyHost
}

// WithProxyPassword adds the proxyPassword to the get compatablity matrix using g e t params
func (o *GetCompatablityMatrixUsingGETParams) WithProxyPassword(proxyPassword *string) *GetCompatablityMatrixUsingGETParams {
	o.SetProxyPassword(proxyPassword)
	return o
}

// SetProxyPassword adds the proxyPassword to the get compatablity matrix using g e t params
func (o *GetCompatablityMatrixUsingGETParams) SetProxyPassword(proxyPassword *string) {
	o.ProxyPassword = proxyPassword
}

// WithProxyPort adds the proxyPort to the get compatablity matrix using g e t params
func (o *GetCompatablityMatrixUsingGETParams) WithProxyPort(proxyPort *int32) *GetCompatablityMatrixUsingGETParams {
	o.SetProxyPort(proxyPort)
	return o
}

// SetProxyPort adds the proxyPort to the get compatablity matrix using g e t params
func (o *GetCompatablityMatrixUsingGETParams) SetProxyPort(proxyPort *int32) {
	o.ProxyPort = proxyPort
}

// WithProxyUserName adds the proxyUserName to the get compatablity matrix using g e t params
func (o *GetCompatablityMatrixUsingGETParams) WithProxyUserName(proxyUserName *string) *GetCompatablityMatrixUsingGETParams {
	o.SetProxyUserName(proxyUserName)
	return o
}

// SetProxyUserName adds the proxyUserName to the get compatablity matrix using g e t params
func (o *GetCompatablityMatrixUsingGETParams) SetProxyUserName(proxyUserName *string) {
	o.ProxyUserName = proxyUserName
}

// WithVersion adds the version to the get compatablity matrix using g e t params
func (o *GetCompatablityMatrixUsingGETParams) WithVersion(version string) *GetCompatablityMatrixUsingGETParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get compatablity matrix using g e t params
func (o *GetCompatablityMatrixUsingGETParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetCompatablityMatrixUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IsProxyEnabled != nil {

		// query param isProxyEnabled
		var qrIsProxyEnabled string
		if o.IsProxyEnabled != nil {
			qrIsProxyEnabled = *o.IsProxyEnabled
		}
		qIsProxyEnabled := qrIsProxyEnabled
		if qIsProxyEnabled != "" {
			if err := r.SetQueryParam("isProxyEnabled", qIsProxyEnabled); err != nil {
				return err
			}
		}

	}

	// path param product
	if err := r.SetPathParam("product", o.Product); err != nil {
		return err
	}

	if o.ProxyHost != nil {

		// query param proxyHost
		var qrProxyHost string
		if o.ProxyHost != nil {
			qrProxyHost = *o.ProxyHost
		}
		qProxyHost := qrProxyHost
		if qProxyHost != "" {
			if err := r.SetQueryParam("proxyHost", qProxyHost); err != nil {
				return err
			}
		}

	}

	if o.ProxyPassword != nil {

		// query param proxyPassword
		var qrProxyPassword string
		if o.ProxyPassword != nil {
			qrProxyPassword = *o.ProxyPassword
		}
		qProxyPassword := qrProxyPassword
		if qProxyPassword != "" {
			if err := r.SetQueryParam("proxyPassword", qProxyPassword); err != nil {
				return err
			}
		}

	}

	if o.ProxyPort != nil {

		// query param proxyPort
		var qrProxyPort int32
		if o.ProxyPort != nil {
			qrProxyPort = *o.ProxyPort
		}
		qProxyPort := swag.FormatInt32(qrProxyPort)
		if qProxyPort != "" {
			if err := r.SetQueryParam("proxyPort", qProxyPort); err != nil {
				return err
			}
		}

	}

	if o.ProxyUserName != nil {

		// query param proxyUserName
		var qrProxyUserName string
		if o.ProxyUserName != nil {
			qrProxyUserName = *o.ProxyUserName
		}
		qProxyUserName := qrProxyUserName
		if qProxyUserName != "" {
			if err := r.SetQueryParam("proxyUserName", qProxyUserName); err != nil {
				return err
			}
		}

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
