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

// NewGetUsersByTypeUsingGETParams creates a new GetUsersByTypeUsingGETParams object
// with the default values initialized.
func NewGetUsersByTypeUsingGETParams() *GetUsersByTypeUsingGETParams {
	var ()
	return &GetUsersByTypeUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersByTypeUsingGETParamsWithTimeout creates a new GetUsersByTypeUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersByTypeUsingGETParamsWithTimeout(timeout time.Duration) *GetUsersByTypeUsingGETParams {
	var ()
	return &GetUsersByTypeUsingGETParams{

		timeout: timeout,
	}
}

// NewGetUsersByTypeUsingGETParamsWithContext creates a new GetUsersByTypeUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersByTypeUsingGETParamsWithContext(ctx context.Context) *GetUsersByTypeUsingGETParams {
	var ()
	return &GetUsersByTypeUsingGETParams{

		Context: ctx,
	}
}

// NewGetUsersByTypeUsingGETParamsWithHTTPClient creates a new GetUsersByTypeUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsersByTypeUsingGETParamsWithHTTPClient(client *http.Client) *GetUsersByTypeUsingGETParams {
	var ()
	return &GetUsersByTypeUsingGETParams{
		HTTPClient: client,
	}
}

/*GetUsersByTypeUsingGETParams contains all the parameters to send to the API endpoint
for the get users by type using g e t operation typically these are written to a http.Request
*/
type GetUsersByTypeUsingGETParams struct {

	/*UserType
	  userType

	*/
	UserType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get users by type using g e t params
func (o *GetUsersByTypeUsingGETParams) WithTimeout(timeout time.Duration) *GetUsersByTypeUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users by type using g e t params
func (o *GetUsersByTypeUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users by type using g e t params
func (o *GetUsersByTypeUsingGETParams) WithContext(ctx context.Context) *GetUsersByTypeUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users by type using g e t params
func (o *GetUsersByTypeUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users by type using g e t params
func (o *GetUsersByTypeUsingGETParams) WithHTTPClient(client *http.Client) *GetUsersByTypeUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users by type using g e t params
func (o *GetUsersByTypeUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserType adds the userType to the get users by type using g e t params
func (o *GetUsersByTypeUsingGETParams) WithUserType(userType string) *GetUsersByTypeUsingGETParams {
	o.SetUserType(userType)
	return o
}

// SetUserType adds the userType to the get users by type using g e t params
func (o *GetUsersByTypeUsingGETParams) SetUserType(userType string) {
	o.UserType = userType
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersByTypeUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param userType
	if err := r.SetPathParam("userType", o.UserType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
