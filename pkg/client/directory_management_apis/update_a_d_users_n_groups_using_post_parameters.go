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

// NewUpdateADUsersNGroupsUsingPOSTParams creates a new UpdateADUsersNGroupsUsingPOSTParams object
// with the default values initialized.
func NewUpdateADUsersNGroupsUsingPOSTParams() *UpdateADUsersNGroupsUsingPOSTParams {
	var ()
	return &UpdateADUsersNGroupsUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateADUsersNGroupsUsingPOSTParamsWithTimeout creates a new UpdateADUsersNGroupsUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateADUsersNGroupsUsingPOSTParamsWithTimeout(timeout time.Duration) *UpdateADUsersNGroupsUsingPOSTParams {
	var ()
	return &UpdateADUsersNGroupsUsingPOSTParams{

		timeout: timeout,
	}
}

// NewUpdateADUsersNGroupsUsingPOSTParamsWithContext creates a new UpdateADUsersNGroupsUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateADUsersNGroupsUsingPOSTParamsWithContext(ctx context.Context) *UpdateADUsersNGroupsUsingPOSTParams {
	var ()
	return &UpdateADUsersNGroupsUsingPOSTParams{

		Context: ctx,
	}
}

// NewUpdateADUsersNGroupsUsingPOSTParamsWithHTTPClient creates a new UpdateADUsersNGroupsUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateADUsersNGroupsUsingPOSTParamsWithHTTPClient(client *http.Client) *UpdateADUsersNGroupsUsingPOSTParams {
	var ()
	return &UpdateADUsersNGroupsUsingPOSTParams{
		HTTPClient: client,
	}
}

/*UpdateADUsersNGroupsUsingPOSTParams contains all the parameters to send to the API endpoint
for the update a d users n groups using p o s t operation typically these are written to a http.Request
*/
type UpdateADUsersNGroupsUsingPOSTParams struct {

	/*VidmUpdateUserNGroupsDTO
	  vidmUpdateUserNGroupsDTO

	*/
	VidmUpdateUserNGroupsDTO *models.VidmUpdateUserNGroupsDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update a d users n groups using p o s t params
func (o *UpdateADUsersNGroupsUsingPOSTParams) WithTimeout(timeout time.Duration) *UpdateADUsersNGroupsUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update a d users n groups using p o s t params
func (o *UpdateADUsersNGroupsUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update a d users n groups using p o s t params
func (o *UpdateADUsersNGroupsUsingPOSTParams) WithContext(ctx context.Context) *UpdateADUsersNGroupsUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update a d users n groups using p o s t params
func (o *UpdateADUsersNGroupsUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update a d users n groups using p o s t params
func (o *UpdateADUsersNGroupsUsingPOSTParams) WithHTTPClient(client *http.Client) *UpdateADUsersNGroupsUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update a d users n groups using p o s t params
func (o *UpdateADUsersNGroupsUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVidmUpdateUserNGroupsDTO adds the vidmUpdateUserNGroupsDTO to the update a d users n groups using p o s t params
func (o *UpdateADUsersNGroupsUsingPOSTParams) WithVidmUpdateUserNGroupsDTO(vidmUpdateUserNGroupsDTO *models.VidmUpdateUserNGroupsDTO) *UpdateADUsersNGroupsUsingPOSTParams {
	o.SetVidmUpdateUserNGroupsDTO(vidmUpdateUserNGroupsDTO)
	return o
}

// SetVidmUpdateUserNGroupsDTO adds the vidmUpdateUserNGroupsDTO to the update a d users n groups using p o s t params
func (o *UpdateADUsersNGroupsUsingPOSTParams) SetVidmUpdateUserNGroupsDTO(vidmUpdateUserNGroupsDTO *models.VidmUpdateUserNGroupsDTO) {
	o.VidmUpdateUserNGroupsDTO = vidmUpdateUserNGroupsDTO
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateADUsersNGroupsUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.VidmUpdateUserNGroupsDTO != nil {
		if err := r.SetBodyParam(o.VidmUpdateUserNGroupsDTO); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
