// Code generated by go-swagger; DO NOT EDIT.

package authentication_user_a_p_is

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

// NewGetUsersByPrincipalNameAndDomainUsingGETParams creates a new GetUsersByPrincipalNameAndDomainUsingGETParams object
// with the default values initialized.
func NewGetUsersByPrincipalNameAndDomainUsingGETParams() *GetUsersByPrincipalNameAndDomainUsingGETParams {
	var ()
	return &GetUsersByPrincipalNameAndDomainUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersByPrincipalNameAndDomainUsingGETParamsWithTimeout creates a new GetUsersByPrincipalNameAndDomainUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersByPrincipalNameAndDomainUsingGETParamsWithTimeout(timeout time.Duration) *GetUsersByPrincipalNameAndDomainUsingGETParams {
	var ()
	return &GetUsersByPrincipalNameAndDomainUsingGETParams{

		timeout: timeout,
	}
}

// NewGetUsersByPrincipalNameAndDomainUsingGETParamsWithContext creates a new GetUsersByPrincipalNameAndDomainUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersByPrincipalNameAndDomainUsingGETParamsWithContext(ctx context.Context) *GetUsersByPrincipalNameAndDomainUsingGETParams {
	var ()
	return &GetUsersByPrincipalNameAndDomainUsingGETParams{

		Context: ctx,
	}
}

// NewGetUsersByPrincipalNameAndDomainUsingGETParamsWithHTTPClient creates a new GetUsersByPrincipalNameAndDomainUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsersByPrincipalNameAndDomainUsingGETParamsWithHTTPClient(client *http.Client) *GetUsersByPrincipalNameAndDomainUsingGETParams {
	var ()
	return &GetUsersByPrincipalNameAndDomainUsingGETParams{
		HTTPClient: client,
	}
}

/*GetUsersByPrincipalNameAndDomainUsingGETParams contains all the parameters to send to the API endpoint
for the get users by principal name and domain using g e t operation typically these are written to a http.Request
*/
type GetUsersByPrincipalNameAndDomainUsingGETParams struct {

	/*Domain
	  domain

	*/
	Domain string
	/*Upn
	  upn

	*/
	Upn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get users by principal name and domain using g e t params
func (o *GetUsersByPrincipalNameAndDomainUsingGETParams) WithTimeout(timeout time.Duration) *GetUsersByPrincipalNameAndDomainUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users by principal name and domain using g e t params
func (o *GetUsersByPrincipalNameAndDomainUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users by principal name and domain using g e t params
func (o *GetUsersByPrincipalNameAndDomainUsingGETParams) WithContext(ctx context.Context) *GetUsersByPrincipalNameAndDomainUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users by principal name and domain using g e t params
func (o *GetUsersByPrincipalNameAndDomainUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users by principal name and domain using g e t params
func (o *GetUsersByPrincipalNameAndDomainUsingGETParams) WithHTTPClient(client *http.Client) *GetUsersByPrincipalNameAndDomainUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users by principal name and domain using g e t params
func (o *GetUsersByPrincipalNameAndDomainUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomain adds the domain to the get users by principal name and domain using g e t params
func (o *GetUsersByPrincipalNameAndDomainUsingGETParams) WithDomain(domain string) *GetUsersByPrincipalNameAndDomainUsingGETParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the get users by principal name and domain using g e t params
func (o *GetUsersByPrincipalNameAndDomainUsingGETParams) SetDomain(domain string) {
	o.Domain = domain
}

// WithUpn adds the upn to the get users by principal name and domain using g e t params
func (o *GetUsersByPrincipalNameAndDomainUsingGETParams) WithUpn(upn string) *GetUsersByPrincipalNameAndDomainUsingGETParams {
	o.SetUpn(upn)
	return o
}

// SetUpn adds the upn to the get users by principal name and domain using g e t params
func (o *GetUsersByPrincipalNameAndDomainUsingGETParams) SetUpn(upn string) {
	o.Upn = upn
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersByPrincipalNameAndDomainUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domain
	if err := r.SetPathParam("domain", o.Domain); err != nil {
		return err
	}

	// path param upn
	if err := r.SetPathParam("upn", o.Upn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
