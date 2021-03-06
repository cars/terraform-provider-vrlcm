// Code generated by go-swagger; DO NOT EDIT.

package authentication_group_role_mapping_a_p_is

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

// NewGetGroupRoleMappingGivenRoleIDUsingGETParams creates a new GetGroupRoleMappingGivenRoleIDUsingGETParams object
// with the default values initialized.
func NewGetGroupRoleMappingGivenRoleIDUsingGETParams() *GetGroupRoleMappingGivenRoleIDUsingGETParams {
	var ()
	return &GetGroupRoleMappingGivenRoleIDUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGroupRoleMappingGivenRoleIDUsingGETParamsWithTimeout creates a new GetGroupRoleMappingGivenRoleIDUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGroupRoleMappingGivenRoleIDUsingGETParamsWithTimeout(timeout time.Duration) *GetGroupRoleMappingGivenRoleIDUsingGETParams {
	var ()
	return &GetGroupRoleMappingGivenRoleIDUsingGETParams{

		timeout: timeout,
	}
}

// NewGetGroupRoleMappingGivenRoleIDUsingGETParamsWithContext creates a new GetGroupRoleMappingGivenRoleIDUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGroupRoleMappingGivenRoleIDUsingGETParamsWithContext(ctx context.Context) *GetGroupRoleMappingGivenRoleIDUsingGETParams {
	var ()
	return &GetGroupRoleMappingGivenRoleIDUsingGETParams{

		Context: ctx,
	}
}

// NewGetGroupRoleMappingGivenRoleIDUsingGETParamsWithHTTPClient creates a new GetGroupRoleMappingGivenRoleIDUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGroupRoleMappingGivenRoleIDUsingGETParamsWithHTTPClient(client *http.Client) *GetGroupRoleMappingGivenRoleIDUsingGETParams {
	var ()
	return &GetGroupRoleMappingGivenRoleIDUsingGETParams{
		HTTPClient: client,
	}
}

/*GetGroupRoleMappingGivenRoleIDUsingGETParams contains all the parameters to send to the API endpoint
for the get group role mapping given role Id using g e t operation typically these are written to a http.Request
*/
type GetGroupRoleMappingGivenRoleIDUsingGETParams struct {

	/*Rolevmid
	  rolevmid

	*/
	Rolevmid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get group role mapping given role Id using g e t params
func (o *GetGroupRoleMappingGivenRoleIDUsingGETParams) WithTimeout(timeout time.Duration) *GetGroupRoleMappingGivenRoleIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get group role mapping given role Id using g e t params
func (o *GetGroupRoleMappingGivenRoleIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get group role mapping given role Id using g e t params
func (o *GetGroupRoleMappingGivenRoleIDUsingGETParams) WithContext(ctx context.Context) *GetGroupRoleMappingGivenRoleIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get group role mapping given role Id using g e t params
func (o *GetGroupRoleMappingGivenRoleIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get group role mapping given role Id using g e t params
func (o *GetGroupRoleMappingGivenRoleIDUsingGETParams) WithHTTPClient(client *http.Client) *GetGroupRoleMappingGivenRoleIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get group role mapping given role Id using g e t params
func (o *GetGroupRoleMappingGivenRoleIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRolevmid adds the rolevmid to the get group role mapping given role Id using g e t params
func (o *GetGroupRoleMappingGivenRoleIDUsingGETParams) WithRolevmid(rolevmid string) *GetGroupRoleMappingGivenRoleIDUsingGETParams {
	o.SetRolevmid(rolevmid)
	return o
}

// SetRolevmid adds the rolevmid to the get group role mapping given role Id using g e t params
func (o *GetGroupRoleMappingGivenRoleIDUsingGETParams) SetRolevmid(rolevmid string) {
	o.Rolevmid = rolevmid
}

// WriteToRequest writes these params to a swagger request
func (o *GetGroupRoleMappingGivenRoleIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
