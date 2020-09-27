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

// NewGetUsersByPrincipalNameUsingGETParams creates a new GetUsersByPrincipalNameUsingGETParams object
// with the default values initialized.
func NewGetUsersByPrincipalNameUsingGETParams() *GetUsersByPrincipalNameUsingGETParams {
	var ()
	return &GetUsersByPrincipalNameUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersByPrincipalNameUsingGETParamsWithTimeout creates a new GetUsersByPrincipalNameUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersByPrincipalNameUsingGETParamsWithTimeout(timeout time.Duration) *GetUsersByPrincipalNameUsingGETParams {
	var ()
	return &GetUsersByPrincipalNameUsingGETParams{

		timeout: timeout,
	}
}

// NewGetUsersByPrincipalNameUsingGETParamsWithContext creates a new GetUsersByPrincipalNameUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersByPrincipalNameUsingGETParamsWithContext(ctx context.Context) *GetUsersByPrincipalNameUsingGETParams {
	var ()
	return &GetUsersByPrincipalNameUsingGETParams{

		Context: ctx,
	}
}

// NewGetUsersByPrincipalNameUsingGETParamsWithHTTPClient creates a new GetUsersByPrincipalNameUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsersByPrincipalNameUsingGETParamsWithHTTPClient(client *http.Client) *GetUsersByPrincipalNameUsingGETParams {
	var ()
	return &GetUsersByPrincipalNameUsingGETParams{
		HTTPClient: client,
	}
}

/*GetUsersByPrincipalNameUsingGETParams contains all the parameters to send to the API endpoint
for the get users by principal name using g e t operation typically these are written to a http.Request
*/
type GetUsersByPrincipalNameUsingGETParams struct {

	/*Upn
	  upn

	*/
	Upn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get users by principal name using g e t params
func (o *GetUsersByPrincipalNameUsingGETParams) WithTimeout(timeout time.Duration) *GetUsersByPrincipalNameUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users by principal name using g e t params
func (o *GetUsersByPrincipalNameUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users by principal name using g e t params
func (o *GetUsersByPrincipalNameUsingGETParams) WithContext(ctx context.Context) *GetUsersByPrincipalNameUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users by principal name using g e t params
func (o *GetUsersByPrincipalNameUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users by principal name using g e t params
func (o *GetUsersByPrincipalNameUsingGETParams) WithHTTPClient(client *http.Client) *GetUsersByPrincipalNameUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users by principal name using g e t params
func (o *GetUsersByPrincipalNameUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpn adds the upn to the get users by principal name using g e t params
func (o *GetUsersByPrincipalNameUsingGETParams) WithUpn(upn string) *GetUsersByPrincipalNameUsingGETParams {
	o.SetUpn(upn)
	return o
}

// SetUpn adds the upn to the get users by principal name using g e t params
func (o *GetUsersByPrincipalNameUsingGETParams) SetUpn(upn string) {
	o.Upn = upn
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersByPrincipalNameUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param upn
	if err := r.SetPathParam("upn", o.Upn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
