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

// NewQueueCommandUsingPOSTParams creates a new QueueCommandUsingPOSTParams object
// with the default values initialized.
func NewQueueCommandUsingPOSTParams() *QueueCommandUsingPOSTParams {
	var ()
	return &QueueCommandUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQueueCommandUsingPOSTParamsWithTimeout creates a new QueueCommandUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueueCommandUsingPOSTParamsWithTimeout(timeout time.Duration) *QueueCommandUsingPOSTParams {
	var ()
	return &QueueCommandUsingPOSTParams{

		timeout: timeout,
	}
}

// NewQueueCommandUsingPOSTParamsWithContext creates a new QueueCommandUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueueCommandUsingPOSTParamsWithContext(ctx context.Context) *QueueCommandUsingPOSTParams {
	var ()
	return &QueueCommandUsingPOSTParams{

		Context: ctx,
	}
}

// NewQueueCommandUsingPOSTParamsWithHTTPClient creates a new QueueCommandUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueueCommandUsingPOSTParamsWithHTTPClient(client *http.Client) *QueueCommandUsingPOSTParams {
	var ()
	return &QueueCommandUsingPOSTParams{
		HTTPClient: client,
	}
}

/*QueueCommandUsingPOSTParams contains all the parameters to send to the API endpoint
for the queue command using p o s t operation typically these are written to a http.Request
*/
type QueueCommandUsingPOSTParams struct {

	/*CommandInfoObject
	  commandInfoObject

	*/
	CommandInfoObject *models.CommandInfoObject

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the queue command using p o s t params
func (o *QueueCommandUsingPOSTParams) WithTimeout(timeout time.Duration) *QueueCommandUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the queue command using p o s t params
func (o *QueueCommandUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the queue command using p o s t params
func (o *QueueCommandUsingPOSTParams) WithContext(ctx context.Context) *QueueCommandUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the queue command using p o s t params
func (o *QueueCommandUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the queue command using p o s t params
func (o *QueueCommandUsingPOSTParams) WithHTTPClient(client *http.Client) *QueueCommandUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the queue command using p o s t params
func (o *QueueCommandUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommandInfoObject adds the commandInfoObject to the queue command using p o s t params
func (o *QueueCommandUsingPOSTParams) WithCommandInfoObject(commandInfoObject *models.CommandInfoObject) *QueueCommandUsingPOSTParams {
	o.SetCommandInfoObject(commandInfoObject)
	return o
}

// SetCommandInfoObject adds the commandInfoObject to the queue command using p o s t params
func (o *QueueCommandUsingPOSTParams) SetCommandInfoObject(commandInfoObject *models.CommandInfoObject) {
	o.CommandInfoObject = commandInfoObject
}

// WriteToRequest writes these params to a swagger request
func (o *QueueCommandUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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