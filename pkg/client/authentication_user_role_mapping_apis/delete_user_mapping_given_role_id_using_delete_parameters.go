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

// NewDeleteUserMappingGivenRoleIDUsingDELETEParams creates a new DeleteUserMappingGivenRoleIDUsingDELETEParams object
// with the default values initialized.
func NewDeleteUserMappingGivenRoleIDUsingDELETEParams() *DeleteUserMappingGivenRoleIDUsingDELETEParams {
	var ()
	return &DeleteUserMappingGivenRoleIDUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserMappingGivenRoleIDUsingDELETEParamsWithTimeout creates a new DeleteUserMappingGivenRoleIDUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserMappingGivenRoleIDUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteUserMappingGivenRoleIDUsingDELETEParams {
	var ()
	return &DeleteUserMappingGivenRoleIDUsingDELETEParams{

		timeout: timeout,
	}
}

// NewDeleteUserMappingGivenRoleIDUsingDELETEParamsWithContext creates a new DeleteUserMappingGivenRoleIDUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserMappingGivenRoleIDUsingDELETEParamsWithContext(ctx context.Context) *DeleteUserMappingGivenRoleIDUsingDELETEParams {
	var ()
	return &DeleteUserMappingGivenRoleIDUsingDELETEParams{

		Context: ctx,
	}
}

// NewDeleteUserMappingGivenRoleIDUsingDELETEParamsWithHTTPClient creates a new DeleteUserMappingGivenRoleIDUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserMappingGivenRoleIDUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteUserMappingGivenRoleIDUsingDELETEParams {
	var ()
	return &DeleteUserMappingGivenRoleIDUsingDELETEParams{
		HTTPClient: client,
	}
}

/*DeleteUserMappingGivenRoleIDUsingDELETEParams contains all the parameters to send to the API endpoint
for the delete user mapping given role Id using d e l e t e operation typically these are written to a http.Request
*/
type DeleteUserMappingGivenRoleIDUsingDELETEParams struct {

	/*Rolevmid
	  rolevmid

	*/
	Rolevmid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user mapping given role Id using d e l e t e params
func (o *DeleteUserMappingGivenRoleIDUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteUserMappingGivenRoleIDUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user mapping given role Id using d e l e t e params
func (o *DeleteUserMappingGivenRoleIDUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user mapping given role Id using d e l e t e params
func (o *DeleteUserMappingGivenRoleIDUsingDELETEParams) WithContext(ctx context.Context) *DeleteUserMappingGivenRoleIDUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user mapping given role Id using d e l e t e params
func (o *DeleteUserMappingGivenRoleIDUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user mapping given role Id using d e l e t e params
func (o *DeleteUserMappingGivenRoleIDUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteUserMappingGivenRoleIDUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user mapping given role Id using d e l e t e params
func (o *DeleteUserMappingGivenRoleIDUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRolevmid adds the rolevmid to the delete user mapping given role Id using d e l e t e params
func (o *DeleteUserMappingGivenRoleIDUsingDELETEParams) WithRolevmid(rolevmid string) *DeleteUserMappingGivenRoleIDUsingDELETEParams {
	o.SetRolevmid(rolevmid)
	return o
}

// SetRolevmid adds the rolevmid to the delete user mapping given role Id using d e l e t e params
func (o *DeleteUserMappingGivenRoleIDUsingDELETEParams) SetRolevmid(rolevmid string) {
	o.Rolevmid = rolevmid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserMappingGivenRoleIDUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
