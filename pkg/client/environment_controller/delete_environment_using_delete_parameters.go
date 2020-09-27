// Code generated by go-swagger; DO NOT EDIT.

package environment_controller

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

	"github.com/cars/terraform-provider-vrlcm/models"
)

// NewDeleteEnvironmentUsingDELETEParams creates a new DeleteEnvironmentUsingDELETEParams object
// with the default values initialized.
func NewDeleteEnvironmentUsingDELETEParams() *DeleteEnvironmentUsingDELETEParams {
	var ()
	return &DeleteEnvironmentUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEnvironmentUsingDELETEParamsWithTimeout creates a new DeleteEnvironmentUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteEnvironmentUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteEnvironmentUsingDELETEParams {
	var ()
	return &DeleteEnvironmentUsingDELETEParams{

		timeout: timeout,
	}
}

// NewDeleteEnvironmentUsingDELETEParamsWithContext creates a new DeleteEnvironmentUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteEnvironmentUsingDELETEParamsWithContext(ctx context.Context) *DeleteEnvironmentUsingDELETEParams {
	var ()
	return &DeleteEnvironmentUsingDELETEParams{

		Context: ctx,
	}
}

// NewDeleteEnvironmentUsingDELETEParamsWithHTTPClient creates a new DeleteEnvironmentUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteEnvironmentUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteEnvironmentUsingDELETEParams {
	var ()
	return &DeleteEnvironmentUsingDELETEParams{
		HTTPClient: client,
	}
}

/*DeleteEnvironmentUsingDELETEParams contains all the parameters to send to the API endpoint
for the delete environment using d e l e t e operation typically these are written to a http.Request
*/
type DeleteEnvironmentUsingDELETEParams struct {

	/*DeleteEnviornmentRequestDTO
	  deleteEnviornmentRequestDTO

	*/
	DeleteEnviornmentRequestDTO *models.DeleteEnvironmentRequestDTO
	/*EnvironmentID
	  environmentId

	*/
	EnvironmentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete environment using d e l e t e params
func (o *DeleteEnvironmentUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteEnvironmentUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete environment using d e l e t e params
func (o *DeleteEnvironmentUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete environment using d e l e t e params
func (o *DeleteEnvironmentUsingDELETEParams) WithContext(ctx context.Context) *DeleteEnvironmentUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete environment using d e l e t e params
func (o *DeleteEnvironmentUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete environment using d e l e t e params
func (o *DeleteEnvironmentUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteEnvironmentUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete environment using d e l e t e params
func (o *DeleteEnvironmentUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleteEnviornmentRequestDTO adds the deleteEnviornmentRequestDTO to the delete environment using d e l e t e params
func (o *DeleteEnvironmentUsingDELETEParams) WithDeleteEnviornmentRequestDTO(deleteEnviornmentRequestDTO *models.DeleteEnvironmentRequestDTO) *DeleteEnvironmentUsingDELETEParams {
	o.SetDeleteEnviornmentRequestDTO(deleteEnviornmentRequestDTO)
	return o
}

// SetDeleteEnviornmentRequestDTO adds the deleteEnviornmentRequestDTO to the delete environment using d e l e t e params
func (o *DeleteEnvironmentUsingDELETEParams) SetDeleteEnviornmentRequestDTO(deleteEnviornmentRequestDTO *models.DeleteEnvironmentRequestDTO) {
	o.DeleteEnviornmentRequestDTO = deleteEnviornmentRequestDTO
}

// WithEnvironmentID adds the environmentID to the delete environment using d e l e t e params
func (o *DeleteEnvironmentUsingDELETEParams) WithEnvironmentID(environmentID string) *DeleteEnvironmentUsingDELETEParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the delete environment using d e l e t e params
func (o *DeleteEnvironmentUsingDELETEParams) SetEnvironmentID(environmentID string) {
	o.EnvironmentID = environmentID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEnvironmentUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DeleteEnviornmentRequestDTO != nil {
		if err := r.SetBodyParam(o.DeleteEnviornmentRequestDTO); err != nil {
			return err
		}
	}

	// path param environmentId
	if err := r.SetPathParam("environmentId", o.EnvironmentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}