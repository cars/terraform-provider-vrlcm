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

// NewGetAllProviderMappedUsersUsingGETParams creates a new GetAllProviderMappedUsersUsingGETParams object
// with the default values initialized.
func NewGetAllProviderMappedUsersUsingGETParams() *GetAllProviderMappedUsersUsingGETParams {

	return &GetAllProviderMappedUsersUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllProviderMappedUsersUsingGETParamsWithTimeout creates a new GetAllProviderMappedUsersUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllProviderMappedUsersUsingGETParamsWithTimeout(timeout time.Duration) *GetAllProviderMappedUsersUsingGETParams {

	return &GetAllProviderMappedUsersUsingGETParams{

		timeout: timeout,
	}
}

// NewGetAllProviderMappedUsersUsingGETParamsWithContext creates a new GetAllProviderMappedUsersUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllProviderMappedUsersUsingGETParamsWithContext(ctx context.Context) *GetAllProviderMappedUsersUsingGETParams {

	return &GetAllProviderMappedUsersUsingGETParams{

		Context: ctx,
	}
}

// NewGetAllProviderMappedUsersUsingGETParamsWithHTTPClient creates a new GetAllProviderMappedUsersUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllProviderMappedUsersUsingGETParamsWithHTTPClient(client *http.Client) *GetAllProviderMappedUsersUsingGETParams {

	return &GetAllProviderMappedUsersUsingGETParams{
		HTTPClient: client,
	}
}

/*GetAllProviderMappedUsersUsingGETParams contains all the parameters to send to the API endpoint
for the get all provider mapped users using g e t operation typically these are written to a http.Request
*/
type GetAllProviderMappedUsersUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all provider mapped users using g e t params
func (o *GetAllProviderMappedUsersUsingGETParams) WithTimeout(timeout time.Duration) *GetAllProviderMappedUsersUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all provider mapped users using g e t params
func (o *GetAllProviderMappedUsersUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all provider mapped users using g e t params
func (o *GetAllProviderMappedUsersUsingGETParams) WithContext(ctx context.Context) *GetAllProviderMappedUsersUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all provider mapped users using g e t params
func (o *GetAllProviderMappedUsersUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all provider mapped users using g e t params
func (o *GetAllProviderMappedUsersUsingGETParams) WithHTTPClient(client *http.Client) *GetAllProviderMappedUsersUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all provider mapped users using g e t params
func (o *GetAllProviderMappedUsersUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllProviderMappedUsersUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
