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

// NewGetAllUserRoleMappingsUsingGETParams creates a new GetAllUserRoleMappingsUsingGETParams object
// with the default values initialized.
func NewGetAllUserRoleMappingsUsingGETParams() *GetAllUserRoleMappingsUsingGETParams {

	return &GetAllUserRoleMappingsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllUserRoleMappingsUsingGETParamsWithTimeout creates a new GetAllUserRoleMappingsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllUserRoleMappingsUsingGETParamsWithTimeout(timeout time.Duration) *GetAllUserRoleMappingsUsingGETParams {

	return &GetAllUserRoleMappingsUsingGETParams{

		timeout: timeout,
	}
}

// NewGetAllUserRoleMappingsUsingGETParamsWithContext creates a new GetAllUserRoleMappingsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllUserRoleMappingsUsingGETParamsWithContext(ctx context.Context) *GetAllUserRoleMappingsUsingGETParams {

	return &GetAllUserRoleMappingsUsingGETParams{

		Context: ctx,
	}
}

// NewGetAllUserRoleMappingsUsingGETParamsWithHTTPClient creates a new GetAllUserRoleMappingsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllUserRoleMappingsUsingGETParamsWithHTTPClient(client *http.Client) *GetAllUserRoleMappingsUsingGETParams {

	return &GetAllUserRoleMappingsUsingGETParams{
		HTTPClient: client,
	}
}

/*GetAllUserRoleMappingsUsingGETParams contains all the parameters to send to the API endpoint
for the get all user role mappings using g e t operation typically these are written to a http.Request
*/
type GetAllUserRoleMappingsUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all user role mappings using g e t params
func (o *GetAllUserRoleMappingsUsingGETParams) WithTimeout(timeout time.Duration) *GetAllUserRoleMappingsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all user role mappings using g e t params
func (o *GetAllUserRoleMappingsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all user role mappings using g e t params
func (o *GetAllUserRoleMappingsUsingGETParams) WithContext(ctx context.Context) *GetAllUserRoleMappingsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all user role mappings using g e t params
func (o *GetAllUserRoleMappingsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all user role mappings using g e t params
func (o *GetAllUserRoleMappingsUsingGETParams) WithHTTPClient(client *http.Client) *GetAllUserRoleMappingsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all user role mappings using g e t params
func (o *GetAllUserRoleMappingsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllUserRoleMappingsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
