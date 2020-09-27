// Code generated by go-swagger; DO NOT EDIT.

package file_browse_controller

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

// NewCreateNodeUsingPOSTParams creates a new CreateNodeUsingPOSTParams object
// with the default values initialized.
func NewCreateNodeUsingPOSTParams() *CreateNodeUsingPOSTParams {
	var ()
	return &CreateNodeUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNodeUsingPOSTParamsWithTimeout creates a new CreateNodeUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateNodeUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateNodeUsingPOSTParams {
	var ()
	return &CreateNodeUsingPOSTParams{

		timeout: timeout,
	}
}

// NewCreateNodeUsingPOSTParamsWithContext creates a new CreateNodeUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateNodeUsingPOSTParamsWithContext(ctx context.Context) *CreateNodeUsingPOSTParams {
	var ()
	return &CreateNodeUsingPOSTParams{

		Context: ctx,
	}
}

// NewCreateNodeUsingPOSTParamsWithHTTPClient creates a new CreateNodeUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateNodeUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateNodeUsingPOSTParams {
	var ()
	return &CreateNodeUsingPOSTParams{
		HTTPClient: client,
	}
}

/*CreateNodeUsingPOSTParams contains all the parameters to send to the API endpoint
for the create node using p o s t operation typically these are written to a http.Request
*/
type CreateNodeUsingPOSTParams struct {

	/*Node
	  node

	*/
	Node *models.NodeDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create node using p o s t params
func (o *CreateNodeUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateNodeUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create node using p o s t params
func (o *CreateNodeUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create node using p o s t params
func (o *CreateNodeUsingPOSTParams) WithContext(ctx context.Context) *CreateNodeUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create node using p o s t params
func (o *CreateNodeUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create node using p o s t params
func (o *CreateNodeUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateNodeUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create node using p o s t params
func (o *CreateNodeUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNode adds the node to the create node using p o s t params
func (o *CreateNodeUsingPOSTParams) WithNode(node *models.NodeDTO) *CreateNodeUsingPOSTParams {
	o.SetNode(node)
	return o
}

// SetNode adds the node to the create node using p o s t params
func (o *CreateNodeUsingPOSTParams) SetNode(node *models.NodeDTO) {
	o.Node = node
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNodeUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Node != nil {
		if err := r.SetBodyParam(o.Node); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
