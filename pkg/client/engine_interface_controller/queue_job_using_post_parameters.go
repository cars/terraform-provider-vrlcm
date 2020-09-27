// Code generated by go-swagger; DO NOT EDIT.

package engine_interface_controller

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

// NewQueueJobUsingPOSTParams creates a new QueueJobUsingPOSTParams object
// with the default values initialized.
func NewQueueJobUsingPOSTParams() *QueueJobUsingPOSTParams {
	var ()
	return &QueueJobUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQueueJobUsingPOSTParamsWithTimeout creates a new QueueJobUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueueJobUsingPOSTParamsWithTimeout(timeout time.Duration) *QueueJobUsingPOSTParams {
	var ()
	return &QueueJobUsingPOSTParams{

		timeout: timeout,
	}
}

// NewQueueJobUsingPOSTParamsWithContext creates a new QueueJobUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueueJobUsingPOSTParamsWithContext(ctx context.Context) *QueueJobUsingPOSTParams {
	var ()
	return &QueueJobUsingPOSTParams{

		Context: ctx,
	}
}

// NewQueueJobUsingPOSTParamsWithHTTPClient creates a new QueueJobUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueueJobUsingPOSTParamsWithHTTPClient(client *http.Client) *QueueJobUsingPOSTParams {
	var ()
	return &QueueJobUsingPOSTParams{
		HTTPClient: client,
	}
}

/*QueueJobUsingPOSTParams contains all the parameters to send to the API endpoint
for the queue job using p o s t operation typically these are written to a http.Request
*/
type QueueJobUsingPOSTParams struct {

	/*Suite
	  suite

	*/
	Suite *models.StateMachineInvocationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the queue job using p o s t params
func (o *QueueJobUsingPOSTParams) WithTimeout(timeout time.Duration) *QueueJobUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the queue job using p o s t params
func (o *QueueJobUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the queue job using p o s t params
func (o *QueueJobUsingPOSTParams) WithContext(ctx context.Context) *QueueJobUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the queue job using p o s t params
func (o *QueueJobUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the queue job using p o s t params
func (o *QueueJobUsingPOSTParams) WithHTTPClient(client *http.Client) *QueueJobUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the queue job using p o s t params
func (o *QueueJobUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSuite adds the suite to the queue job using p o s t params
func (o *QueueJobUsingPOSTParams) WithSuite(suite *models.StateMachineInvocationRequest) *QueueJobUsingPOSTParams {
	o.SetSuite(suite)
	return o
}

// SetSuite adds the suite to the queue job using p o s t params
func (o *QueueJobUsingPOSTParams) SetSuite(suite *models.StateMachineInvocationRequest) {
	o.Suite = suite
}

// WriteToRequest writes these params to a swagger request
func (o *QueueJobUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Suite != nil {
		if err := r.SetBodyParam(o.Suite); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}