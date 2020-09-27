// Code generated by go-swagger; DO NOT EDIT.

package authentication_user_role_mapping_a_p_is

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

// NewGetUserRoleMappingGivenRoleIDUsingGETParams creates a new GetUserRoleMappingGivenRoleIDUsingGETParams object
// with the default values initialized.
func NewGetUserRoleMappingGivenRoleIDUsingGETParams() *GetUserRoleMappingGivenRoleIDUsingGETParams {
	var ()
	return &GetUserRoleMappingGivenRoleIDUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserRoleMappingGivenRoleIDUsingGETParamsWithTimeout creates a new GetUserRoleMappingGivenRoleIDUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserRoleMappingGivenRoleIDUsingGETParamsWithTimeout(timeout time.Duration) *GetUserRoleMappingGivenRoleIDUsingGETParams {
	var ()
	return &GetUserRoleMappingGivenRoleIDUsingGETParams{

		timeout: timeout,
	}
}

// NewGetUserRoleMappingGivenRoleIDUsingGETParamsWithContext creates a new GetUserRoleMappingGivenRoleIDUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserRoleMappingGivenRoleIDUsingGETParamsWithContext(ctx context.Context) *GetUserRoleMappingGivenRoleIDUsingGETParams {
	var ()
	return &GetUserRoleMappingGivenRoleIDUsingGETParams{

		Context: ctx,
	}
}

// NewGetUserRoleMappingGivenRoleIDUsingGETParamsWithHTTPClient creates a new GetUserRoleMappingGivenRoleIDUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserRoleMappingGivenRoleIDUsingGETParamsWithHTTPClient(client *http.Client) *GetUserRoleMappingGivenRoleIDUsingGETParams {
	var ()
	return &GetUserRoleMappingGivenRoleIDUsingGETParams{
		HTTPClient: client,
	}
}

/*GetUserRoleMappingGivenRoleIDUsingGETParams contains all the parameters to send to the API endpoint
for the get user role mapping given role Id using g e t operation typically these are written to a http.Request
*/
type GetUserRoleMappingGivenRoleIDUsingGETParams struct {

	/*Rolevmid
	  rolevmid

	*/
	Rolevmid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user role mapping given role Id using g e t params
func (o *GetUserRoleMappingGivenRoleIDUsingGETParams) WithTimeout(timeout time.Duration) *GetUserRoleMappingGivenRoleIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user role mapping given role Id using g e t params
func (o *GetUserRoleMappingGivenRoleIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user role mapping given role Id using g e t params
func (o *GetUserRoleMappingGivenRoleIDUsingGETParams) WithContext(ctx context.Context) *GetUserRoleMappingGivenRoleIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user role mapping given role Id using g e t params
func (o *GetUserRoleMappingGivenRoleIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user role mapping given role Id using g e t params
func (o *GetUserRoleMappingGivenRoleIDUsingGETParams) WithHTTPClient(client *http.Client) *GetUserRoleMappingGivenRoleIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user role mapping given role Id using g e t params
func (o *GetUserRoleMappingGivenRoleIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRolevmid adds the rolevmid to the get user role mapping given role Id using g e t params
func (o *GetUserRoleMappingGivenRoleIDUsingGETParams) WithRolevmid(rolevmid string) *GetUserRoleMappingGivenRoleIDUsingGETParams {
	o.SetRolevmid(rolevmid)
	return o
}

// SetRolevmid adds the rolevmid to the get user role mapping given role Id using g e t params
func (o *GetUserRoleMappingGivenRoleIDUsingGETParams) SetRolevmid(rolevmid string) {
	o.Rolevmid = rolevmid
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserRoleMappingGivenRoleIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param rolevmid
	if err := r.SetPathParam("rolevmid", o.Rolevmid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
