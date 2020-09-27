// Code generated by go-swagger; DO NOT EDIT.

package command_controller

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

// NewUpdateCommandUsingPATCHParams creates a new UpdateCommandUsingPATCHParams object
// with the default values initialized.
func NewUpdateCommandUsingPATCHParams() *UpdateCommandUsingPATCHParams {
	var ()
	return &UpdateCommandUsingPATCHParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCommandUsingPATCHParamsWithTimeout creates a new UpdateCommandUsingPATCHParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCommandUsingPATCHParamsWithTimeout(timeout time.Duration) *UpdateCommandUsingPATCHParams {
	var ()
	return &UpdateCommandUsingPATCHParams{

		timeout: timeout,
	}
}

// NewUpdateCommandUsingPATCHParamsWithContext creates a new UpdateCommandUsingPATCHParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCommandUsingPATCHParamsWithContext(ctx context.Context) *UpdateCommandUsingPATCHParams {
	var ()
	return &UpdateCommandUsingPATCHParams{

		Context: ctx,
	}
}

// NewUpdateCommandUsingPATCHParamsWithHTTPClient creates a new UpdateCommandUsingPATCHParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCommandUsingPATCHParamsWithHTTPClient(client *http.Client) *UpdateCommandUsingPATCHParams {
	var ()
	return &UpdateCommandUsingPATCHParams{
		HTTPClient: client,
	}
}

/*UpdateCommandUsingPATCHParams contains all the parameters to send to the API endpoint
for the update command using p a t c h operation typically these are written to a http.Request
*/
type UpdateCommandUsingPATCHParams struct {

	/*CommandInfoObject
	  commandInfoObject

	*/
	CommandInfoObject *models.CommandInfoObject

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update command using p a t c h params
func (o *UpdateCommandUsingPATCHParams) WithTimeout(timeout time.Duration) *UpdateCommandUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update command using p a t c h params
func (o *UpdateCommandUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update command using p a t c h params
func (o *UpdateCommandUsingPATCHParams) WithContext(ctx context.Context) *UpdateCommandUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update command using p a t c h params
func (o *UpdateCommandUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update command using p a t c h params
func (o *UpdateCommandUsingPATCHParams) WithHTTPClient(client *http.Client) *UpdateCommandUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update command using p a t c h params
func (o *UpdateCommandUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommandInfoObject adds the commandInfoObject to the update command using p a t c h params
func (o *UpdateCommandUsingPATCHParams) WithCommandInfoObject(commandInfoObject *models.CommandInfoObject) *UpdateCommandUsingPATCHParams {
	o.SetCommandInfoObject(commandInfoObject)
	return o
}

// SetCommandInfoObject adds the commandInfoObject to the update command using p a t c h params
func (o *UpdateCommandUsingPATCHParams) SetCommandInfoObject(commandInfoObject *models.CommandInfoObject) {
	o.CommandInfoObject = commandInfoObject
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCommandUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CommandInfoObject != nil {
		if err := r.SetBodyParam(o.CommandInfoObject); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
