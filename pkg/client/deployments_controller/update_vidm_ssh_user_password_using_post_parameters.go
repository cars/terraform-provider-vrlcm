// Code generated by go-swagger; DO NOT EDIT.

package deployments_controller

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

// NewUpdateVidmSSHUserPasswordUsingPOSTParams creates a new UpdateVidmSSHUserPasswordUsingPOSTParams object
// with the default values initialized.
func NewUpdateVidmSSHUserPasswordUsingPOSTParams() *UpdateVidmSSHUserPasswordUsingPOSTParams {
	var ()
	return &UpdateVidmSSHUserPasswordUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVidmSSHUserPasswordUsingPOSTParamsWithTimeout creates a new UpdateVidmSSHUserPasswordUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateVidmSSHUserPasswordUsingPOSTParamsWithTimeout(timeout time.Duration) *UpdateVidmSSHUserPasswordUsingPOSTParams {
	var ()
	return &UpdateVidmSSHUserPasswordUsingPOSTParams{

		timeout: timeout,
	}
}

// NewUpdateVidmSSHUserPasswordUsingPOSTParamsWithContext creates a new UpdateVidmSSHUserPasswordUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateVidmSSHUserPasswordUsingPOSTParamsWithContext(ctx context.Context) *UpdateVidmSSHUserPasswordUsingPOSTParams {
	var ()
	return &UpdateVidmSSHUserPasswordUsingPOSTParams{

		Context: ctx,
	}
}

// NewUpdateVidmSSHUserPasswordUsingPOSTParamsWithHTTPClient creates a new UpdateVidmSSHUserPasswordUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateVidmSSHUserPasswordUsingPOSTParamsWithHTTPClient(client *http.Client) *UpdateVidmSSHUserPasswordUsingPOSTParams {
	var ()
	return &UpdateVidmSSHUserPasswordUsingPOSTParams{
		HTTPClient: client,
	}
}

/*UpdateVidmSSHUserPasswordUsingPOSTParams contains all the parameters to send to the API endpoint
for the update vidm Ssh user password using p o s t operation typically these are written to a http.Request
*/
type UpdateVidmSSHUserPasswordUsingPOSTParams struct {

	/*EnvironmentID
	  environmentId

	*/
	EnvironmentID string
	/*NodeType
	  nodeType

	*/
	NodeType string
	/*ProductID
	  productId

	*/
	ProductID string
	/*UpdateNodePasswordRequestDTO
	  updateNodePasswordRequestDTO

	*/
	UpdateNodePasswordRequestDTO *models.UpdateNodePasswordRequestDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update vidm Ssh user password using p o s t params
func (o *UpdateVidmSSHUserPasswordUsingPOSTParams) WithTimeout(timeout time.Duration) *UpdateVidmSSHUserPasswordUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update vidm Ssh user password using p o s t params
func (o *UpdateVidmSSHUserPasswordUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update vidm Ssh user password using p o s t params
func (o *UpdateVidmSSHUserPasswordUsingPOSTParams) WithContext(ctx context.Context) *UpdateVidmSSHUserPasswordUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update vidm Ssh user password using p o s t params
func (o *UpdateVidmSSHUserPasswordUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update vidm Ssh user password using p o s t params
func (o *UpdateVidmSSHUserPasswordUsingPOSTParams) WithHTTPClient(client *http.Client) *UpdateVidmSSHUserPasswordUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update vidm Ssh user password using p o s t params
func (o *UpdateVidmSSHUserPasswordUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the update vidm Ssh user password using p o s t params
func (o *UpdateVidmSSHUserPasswordUsingPOSTParams) WithEnvironmentID(environmentID string) *UpdateVidmSSHUserPasswordUsingPOSTParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the update vidm Ssh user password using p o s t params
func (o *UpdateVidmSSHUserPasswordUsingPOSTParams) SetEnvironmentID(environmentID string) {
	o.EnvironmentID = environmentID
}

// WithNodeType adds the nodeType to the update vidm Ssh user password using p o s t params
func (o *UpdateVidmSSHUserPasswordUsingPOSTParams) WithNodeType(nodeType string) *UpdateVidmSSHUserPasswordUsingPOSTParams {
	o.SetNodeType(nodeType)
	return o
}

// SetNodeType adds the nodeType to the update vidm Ssh user password using p o s t params
func (o *UpdateVidmSSHUserPasswordUsingPOSTParams) SetNodeType(nodeType string) {
	o.NodeType = nodeType
}

// WithProductID adds the productID to the update vidm Ssh user password using p o s t params
func (o *UpdateVidmSSHUserPasswordUsingPOSTParams) WithProductID(productID string) *UpdateVidmSSHUserPasswordUsingPOSTParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the update vidm Ssh user password using p o s t params
func (o *UpdateVidmSSHUserPasswordUsingPOSTParams) SetProductID(productID string) {
	o.ProductID = productID
}

// WithUpdateNodePasswordRequestDTO adds the updateNodePasswordRequestDTO to the update vidm Ssh user password using p o s t params
func (o *UpdateVidmSSHUserPasswordUsingPOSTParams) WithUpdateNodePasswordRequestDTO(updateNodePasswordRequestDTO *models.UpdateNodePasswordRequestDTO) *UpdateVidmSSHUserPasswordUsingPOSTParams {
	o.SetUpdateNodePasswordRequestDTO(updateNodePasswordRequestDTO)
	return o
}

// SetUpdateNodePasswordRequestDTO adds the updateNodePasswordRequestDTO to the update vidm Ssh user password using p o s t params
func (o *UpdateVidmSSHUserPasswordUsingPOSTParams) SetUpdateNodePasswordRequestDTO(updateNodePasswordRequestDTO *models.UpdateNodePasswordRequestDTO) {
	o.UpdateNodePasswordRequestDTO = updateNodePasswordRequestDTO
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVidmSSHUserPasswordUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param environmentId
	if err := r.SetPathParam("environmentId", o.EnvironmentID); err != nil {
		return err
	}

	// path param nodeType
	if err := r.SetPathParam("nodeType", o.NodeType); err != nil {
		return err
	}

	// path param productId
	if err := r.SetPathParam("productId", o.ProductID); err != nil {
		return err
	}

	if o.UpdateNodePasswordRequestDTO != nil {
		if err := r.SetBodyParam(o.UpdateNodePasswordRequestDTO); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
