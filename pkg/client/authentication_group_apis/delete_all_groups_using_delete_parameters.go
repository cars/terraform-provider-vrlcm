// Code generated by go-swagger; DO NOT EDIT.

package authentication_group_a_p_is

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

// NewDeleteAllGroupsUsingDELETEParams creates a new DeleteAllGroupsUsingDELETEParams object
// with the default values initialized.
func NewDeleteAllGroupsUsingDELETEParams() *DeleteAllGroupsUsingDELETEParams {

	return &DeleteAllGroupsUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAllGroupsUsingDELETEParamsWithTimeout creates a new DeleteAllGroupsUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAllGroupsUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteAllGroupsUsingDELETEParams {

	return &DeleteAllGroupsUsingDELETEParams{

		timeout: timeout,
	}
}

// NewDeleteAllGroupsUsingDELETEParamsWithContext creates a new DeleteAllGroupsUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAllGroupsUsingDELETEParamsWithContext(ctx context.Context) *DeleteAllGroupsUsingDELETEParams {

	return &DeleteAllGroupsUsingDELETEParams{

		Context: ctx,
	}
}

// NewDeleteAllGroupsUsingDELETEParamsWithHTTPClient creates a new DeleteAllGroupsUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAllGroupsUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteAllGroupsUsingDELETEParams {

	return &DeleteAllGroupsUsingDELETEParams{
		HTTPClient: client,
	}
}

/*DeleteAllGroupsUsingDELETEParams contains all the parameters to send to the API endpoint
for the delete all groups using d e l e t e operation typically these are written to a http.Request
*/
type DeleteAllGroupsUsingDELETEParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete all groups using d e l e t e params
func (o *DeleteAllGroupsUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteAllGroupsUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete all groups using d e l e t e params
func (o *DeleteAllGroupsUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete all groups using d e l e t e params
func (o *DeleteAllGroupsUsingDELETEParams) WithContext(ctx context.Context) *DeleteAllGroupsUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete all groups using d e l e t e params
func (o *DeleteAllGroupsUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete all groups using d e l e t e params
func (o *DeleteAllGroupsUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteAllGroupsUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete all groups using d e l e t e params
func (o *DeleteAllGroupsUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAllGroupsUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
