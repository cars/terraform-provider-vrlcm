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

// NewGetChildCommandsUsingPOSTParams creates a new GetChildCommandsUsingPOSTParams object
// with the default values initialized.
func NewGetChildCommandsUsingPOSTParams() *GetChildCommandsUsingPOSTParams {
	var ()
	return &GetChildCommandsUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetChildCommandsUsingPOSTParamsWithTimeout creates a new GetChildCommandsUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetChildCommandsUsingPOSTParamsWithTimeout(timeout time.Duration) *GetChildCommandsUsingPOSTParams {
	var ()
	return &GetChildCommandsUsingPOSTParams{

		timeout: timeout,
	}
}

// NewGetChildCommandsUsingPOSTParamsWithContext creates a new GetChildCommandsUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetChildCommandsUsingPOSTParamsWithContext(ctx context.Context) *GetChildCommandsUsingPOSTParams {
	var ()
	return &GetChildCommandsUsingPOSTParams{

		Context: ctx,
	}
}

// NewGetChildCommandsUsingPOSTParamsWithHTTPClient creates a new GetChildCommandsUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetChildCommandsUsingPOSTParamsWithHTTPClient(client *http.Client) *GetChildCommandsUsingPOSTParams {
	var ()
	return &GetChildCommandsUsingPOSTParams{
		HTTPClient: client,
	}
}

/*GetChildCommandsUsingPOSTParams contains all the parameters to send to the API endpoint
for the get child commands using p o s t operation typically these are written to a http.Request
*/
type GetChildCommandsUsingPOSTParams struct {

	/*CommandRequest
	  commandRequest

	*/
	CommandRequest *models.CommandRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get child commands using p o s t params
func (o *GetChildCommandsUsingPOSTParams) WithTimeout(timeout time.Duration) *GetChildCommandsUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get child commands using p o s t params
func (o *GetChildCommandsUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get child commands using p o s t params
func (o *GetChildCommandsUsingPOSTParams) WithContext(ctx context.Context) *GetChildCommandsUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get child commands using p o s t params
func (o *GetChildCommandsUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get child commands using p o s t params
func (o *GetChildCommandsUsingPOSTParams) WithHTTPClient(client *http.Client) *GetChildCommandsUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get child commands using p o s t params
func (o *GetChildCommandsUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommandRequest adds the commandRequest to the get child commands using p o s t params
func (o *GetChildCommandsUsingPOSTParams) WithCommandRequest(commandRequest *models.CommandRequest) *GetChildCommandsUsingPOSTParams {
	o.SetCommandRequest(commandRequest)
	return o
}

// SetCommandRequest adds the commandRequest to the get child commands using p o s t params
func (o *GetChildCommandsUsingPOSTParams) SetCommandRequest(commandRequest *models.CommandRequest) {
	o.CommandRequest = commandRequest
}

// WriteToRequest writes these params to a swagger request
func (o *GetChildCommandsUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CommandRequest != nil {
		if err := r.SetBodyParam(o.CommandRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
