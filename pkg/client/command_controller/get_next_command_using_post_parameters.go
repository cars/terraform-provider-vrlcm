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

// NewGetNextCommandUsingPOSTParams creates a new GetNextCommandUsingPOSTParams object
// with the default values initialized.
func NewGetNextCommandUsingPOSTParams() *GetNextCommandUsingPOSTParams {
	var ()
	return &GetNextCommandUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNextCommandUsingPOSTParamsWithTimeout creates a new GetNextCommandUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNextCommandUsingPOSTParamsWithTimeout(timeout time.Duration) *GetNextCommandUsingPOSTParams {
	var ()
	return &GetNextCommandUsingPOSTParams{

		timeout: timeout,
	}
}

// NewGetNextCommandUsingPOSTParamsWithContext creates a new GetNextCommandUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNextCommandUsingPOSTParamsWithContext(ctx context.Context) *GetNextCommandUsingPOSTParams {
	var ()
	return &GetNextCommandUsingPOSTParams{

		Context: ctx,
	}
}

// NewGetNextCommandUsingPOSTParamsWithHTTPClient creates a new GetNextCommandUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNextCommandUsingPOSTParamsWithHTTPClient(client *http.Client) *GetNextCommandUsingPOSTParams {
	var ()
	return &GetNextCommandUsingPOSTParams{
		HTTPClient: client,
	}
}

/*GetNextCommandUsingPOSTParams contains all the parameters to send to the API endpoint
for the get next command using p o s t operation typically these are written to a http.Request
*/
type GetNextCommandUsingPOSTParams struct {

	/*NodeNextCommandRequest
	  nodeNextCommandRequest

	*/
	NodeNextCommandRequest *models.NodeNextCommandRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get next command using p o s t params
func (o *GetNextCommandUsingPOSTParams) WithTimeout(timeout time.Duration) *GetNextCommandUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get next command using p o s t params
func (o *GetNextCommandUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get next command using p o s t params
func (o *GetNextCommandUsingPOSTParams) WithContext(ctx context.Context) *GetNextCommandUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get next command using p o s t params
func (o *GetNextCommandUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get next command using p o s t params
func (o *GetNextCommandUsingPOSTParams) WithHTTPClient(client *http.Client) *GetNextCommandUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get next command using p o s t params
func (o *GetNextCommandUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNodeNextCommandRequest adds the nodeNextCommandRequest to the get next command using p o s t params
func (o *GetNextCommandUsingPOSTParams) WithNodeNextCommandRequest(nodeNextCommandRequest *models.NodeNextCommandRequest) *GetNextCommandUsingPOSTParams {
	o.SetNodeNextCommandRequest(nodeNextCommandRequest)
	return o
}

// SetNodeNextCommandRequest adds the nodeNextCommandRequest to the get next command using p o s t params
func (o *GetNextCommandUsingPOSTParams) SetNodeNextCommandRequest(nodeNextCommandRequest *models.NodeNextCommandRequest) {
	o.NodeNextCommandRequest = nodeNextCommandRequest
}

// WriteToRequest writes these params to a swagger request
func (o *GetNextCommandUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.NodeNextCommandRequest != nil {
		if err := r.SetBodyParam(o.NodeNextCommandRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
