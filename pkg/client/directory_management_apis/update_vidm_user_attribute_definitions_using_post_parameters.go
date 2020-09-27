// Code generated by go-swagger; DO NOT EDIT.

package directory_management_a_p_is

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

// NewUpdateVidmUserAttributeDefinitionsUsingPOSTParams creates a new UpdateVidmUserAttributeDefinitionsUsingPOSTParams object
// with the default values initialized.
func NewUpdateVidmUserAttributeDefinitionsUsingPOSTParams() *UpdateVidmUserAttributeDefinitionsUsingPOSTParams {
	var ()
	return &UpdateVidmUserAttributeDefinitionsUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVidmUserAttributeDefinitionsUsingPOSTParamsWithTimeout creates a new UpdateVidmUserAttributeDefinitionsUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateVidmUserAttributeDefinitionsUsingPOSTParamsWithTimeout(timeout time.Duration) *UpdateVidmUserAttributeDefinitionsUsingPOSTParams {
	var ()
	return &UpdateVidmUserAttributeDefinitionsUsingPOSTParams{

		timeout: timeout,
	}
}

// NewUpdateVidmUserAttributeDefinitionsUsingPOSTParamsWithContext creates a new UpdateVidmUserAttributeDefinitionsUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateVidmUserAttributeDefinitionsUsingPOSTParamsWithContext(ctx context.Context) *UpdateVidmUserAttributeDefinitionsUsingPOSTParams {
	var ()
	return &UpdateVidmUserAttributeDefinitionsUsingPOSTParams{

		Context: ctx,
	}
}

// NewUpdateVidmUserAttributeDefinitionsUsingPOSTParamsWithHTTPClient creates a new UpdateVidmUserAttributeDefinitionsUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateVidmUserAttributeDefinitionsUsingPOSTParamsWithHTTPClient(client *http.Client) *UpdateVidmUserAttributeDefinitionsUsingPOSTParams {
	var ()
	return &UpdateVidmUserAttributeDefinitionsUsingPOSTParams{
		HTTPClient: client,
	}
}

/*UpdateVidmUserAttributeDefinitionsUsingPOSTParams contains all the parameters to send to the API endpoint
for the update vidm user attribute definitions using p o s t operation typically these are written to a http.Request
*/
type UpdateVidmUserAttributeDefinitionsUsingPOSTParams struct {

	/*UserAttrDefnUpdateReqDTO
	  userAttrDefnUpdateReqDTO

	*/
	UserAttrDefnUpdateReqDTO *models.UserAttrDefnUpdateReqDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update vidm user attribute definitions using p o s t params
func (o *UpdateVidmUserAttributeDefinitionsUsingPOSTParams) WithTimeout(timeout time.Duration) *UpdateVidmUserAttributeDefinitionsUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update vidm user attribute definitions using p o s t params
func (o *UpdateVidmUserAttributeDefinitionsUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update vidm user attribute definitions using p o s t params
func (o *UpdateVidmUserAttributeDefinitionsUsingPOSTParams) WithContext(ctx context.Context) *UpdateVidmUserAttributeDefinitionsUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update vidm user attribute definitions using p o s t params
func (o *UpdateVidmUserAttributeDefinitionsUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update vidm user attribute definitions using p o s t params
func (o *UpdateVidmUserAttributeDefinitionsUsingPOSTParams) WithHTTPClient(client *http.Client) *UpdateVidmUserAttributeDefinitionsUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update vidm user attribute definitions using p o s t params
func (o *UpdateVidmUserAttributeDefinitionsUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAttrDefnUpdateReqDTO adds the userAttrDefnUpdateReqDTO to the update vidm user attribute definitions using p o s t params
func (o *UpdateVidmUserAttributeDefinitionsUsingPOSTParams) WithUserAttrDefnUpdateReqDTO(userAttrDefnUpdateReqDTO *models.UserAttrDefnUpdateReqDTO) *UpdateVidmUserAttributeDefinitionsUsingPOSTParams {
	o.SetUserAttrDefnUpdateReqDTO(userAttrDefnUpdateReqDTO)
	return o
}

// SetUserAttrDefnUpdateReqDTO adds the userAttrDefnUpdateReqDTO to the update vidm user attribute definitions using p o s t params
func (o *UpdateVidmUserAttributeDefinitionsUsingPOSTParams) SetUserAttrDefnUpdateReqDTO(userAttrDefnUpdateReqDTO *models.UserAttrDefnUpdateReqDTO) {
	o.UserAttrDefnUpdateReqDTO = userAttrDefnUpdateReqDTO
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVidmUserAttributeDefinitionsUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UserAttrDefnUpdateReqDTO != nil {
		if err := r.SetBodyParam(o.UserAttrDefnUpdateReqDTO); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}