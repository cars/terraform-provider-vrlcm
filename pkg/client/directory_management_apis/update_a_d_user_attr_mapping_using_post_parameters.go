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

// NewUpdateADUserAttrMappingUsingPOSTParams creates a new UpdateADUserAttrMappingUsingPOSTParams object
// with the default values initialized.
func NewUpdateADUserAttrMappingUsingPOSTParams() *UpdateADUserAttrMappingUsingPOSTParams {
	var ()
	return &UpdateADUserAttrMappingUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateADUserAttrMappingUsingPOSTParamsWithTimeout creates a new UpdateADUserAttrMappingUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateADUserAttrMappingUsingPOSTParamsWithTimeout(timeout time.Duration) *UpdateADUserAttrMappingUsingPOSTParams {
	var ()
	return &UpdateADUserAttrMappingUsingPOSTParams{

		timeout: timeout,
	}
}

// NewUpdateADUserAttrMappingUsingPOSTParamsWithContext creates a new UpdateADUserAttrMappingUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateADUserAttrMappingUsingPOSTParamsWithContext(ctx context.Context) *UpdateADUserAttrMappingUsingPOSTParams {
	var ()
	return &UpdateADUserAttrMappingUsingPOSTParams{

		Context: ctx,
	}
}

// NewUpdateADUserAttrMappingUsingPOSTParamsWithHTTPClient creates a new UpdateADUserAttrMappingUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateADUserAttrMappingUsingPOSTParamsWithHTTPClient(client *http.Client) *UpdateADUserAttrMappingUsingPOSTParams {
	var ()
	return &UpdateADUserAttrMappingUsingPOSTParams{
		HTTPClient: client,
	}
}

/*UpdateADUserAttrMappingUsingPOSTParams contains all the parameters to send to the API endpoint
for the update a d user attr mapping using p o s t operation typically these are written to a http.Request
*/
type UpdateADUserAttrMappingUsingPOSTParams struct {

	/*VidmUpdUserAttributeMappingDTO
	  vidmUpdUserAttributeMappingDTO

	*/
	VidmUpdUserAttributeMappingDTO *models.VidmUpdUserAttributeMappingDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update a d user attr mapping using p o s t params
func (o *UpdateADUserAttrMappingUsingPOSTParams) WithTimeout(timeout time.Duration) *UpdateADUserAttrMappingUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update a d user attr mapping using p o s t params
func (o *UpdateADUserAttrMappingUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update a d user attr mapping using p o s t params
func (o *UpdateADUserAttrMappingUsingPOSTParams) WithContext(ctx context.Context) *UpdateADUserAttrMappingUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update a d user attr mapping using p o s t params
func (o *UpdateADUserAttrMappingUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update a d user attr mapping using p o s t params
func (o *UpdateADUserAttrMappingUsingPOSTParams) WithHTTPClient(client *http.Client) *UpdateADUserAttrMappingUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update a d user attr mapping using p o s t params
func (o *UpdateADUserAttrMappingUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVidmUpdUserAttributeMappingDTO adds the vidmUpdUserAttributeMappingDTO to the update a d user attr mapping using p o s t params
func (o *UpdateADUserAttrMappingUsingPOSTParams) WithVidmUpdUserAttributeMappingDTO(vidmUpdUserAttributeMappingDTO *models.VidmUpdUserAttributeMappingDTO) *UpdateADUserAttrMappingUsingPOSTParams {
	o.SetVidmUpdUserAttributeMappingDTO(vidmUpdUserAttributeMappingDTO)
	return o
}

// SetVidmUpdUserAttributeMappingDTO adds the vidmUpdUserAttributeMappingDTO to the update a d user attr mapping using p o s t params
func (o *UpdateADUserAttrMappingUsingPOSTParams) SetVidmUpdUserAttributeMappingDTO(vidmUpdUserAttributeMappingDTO *models.VidmUpdUserAttributeMappingDTO) {
	o.VidmUpdUserAttributeMappingDTO = vidmUpdUserAttributeMappingDTO
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateADUserAttrMappingUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.VidmUpdUserAttributeMappingDTO != nil {
		if err := r.SetBodyParam(o.VidmUpdUserAttributeMappingDTO); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
