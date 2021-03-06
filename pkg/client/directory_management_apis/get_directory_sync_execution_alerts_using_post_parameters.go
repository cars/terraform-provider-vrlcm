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

// NewGetDirectorySyncExecutionAlertsUsingPOSTParams creates a new GetDirectorySyncExecutionAlertsUsingPOSTParams object
// with the default values initialized.
func NewGetDirectorySyncExecutionAlertsUsingPOSTParams() *GetDirectorySyncExecutionAlertsUsingPOSTParams {
	var ()
	return &GetDirectorySyncExecutionAlertsUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDirectorySyncExecutionAlertsUsingPOSTParamsWithTimeout creates a new GetDirectorySyncExecutionAlertsUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDirectorySyncExecutionAlertsUsingPOSTParamsWithTimeout(timeout time.Duration) *GetDirectorySyncExecutionAlertsUsingPOSTParams {
	var ()
	return &GetDirectorySyncExecutionAlertsUsingPOSTParams{

		timeout: timeout,
	}
}

// NewGetDirectorySyncExecutionAlertsUsingPOSTParamsWithContext creates a new GetDirectorySyncExecutionAlertsUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDirectorySyncExecutionAlertsUsingPOSTParamsWithContext(ctx context.Context) *GetDirectorySyncExecutionAlertsUsingPOSTParams {
	var ()
	return &GetDirectorySyncExecutionAlertsUsingPOSTParams{

		Context: ctx,
	}
}

// NewGetDirectorySyncExecutionAlertsUsingPOSTParamsWithHTTPClient creates a new GetDirectorySyncExecutionAlertsUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDirectorySyncExecutionAlertsUsingPOSTParamsWithHTTPClient(client *http.Client) *GetDirectorySyncExecutionAlertsUsingPOSTParams {
	var ()
	return &GetDirectorySyncExecutionAlertsUsingPOSTParams{
		HTTPClient: client,
	}
}

/*GetDirectorySyncExecutionAlertsUsingPOSTParams contains all the parameters to send to the API endpoint
for the get directory sync execution alerts using p o s t operation typically these are written to a http.Request
*/
type GetDirectorySyncExecutionAlertsUsingPOSTParams struct {

	/*VidmGetDirSyncExecAlertsReqDTO
	  vidmGetDirSyncExecAlertsReqDTO

	*/
	VidmGetDirSyncExecAlertsReqDTO *models.VidmGetDirSyncExecAlertsReqDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get directory sync execution alerts using p o s t params
func (o *GetDirectorySyncExecutionAlertsUsingPOSTParams) WithTimeout(timeout time.Duration) *GetDirectorySyncExecutionAlertsUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get directory sync execution alerts using p o s t params
func (o *GetDirectorySyncExecutionAlertsUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get directory sync execution alerts using p o s t params
func (o *GetDirectorySyncExecutionAlertsUsingPOSTParams) WithContext(ctx context.Context) *GetDirectorySyncExecutionAlertsUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get directory sync execution alerts using p o s t params
func (o *GetDirectorySyncExecutionAlertsUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get directory sync execution alerts using p o s t params
func (o *GetDirectorySyncExecutionAlertsUsingPOSTParams) WithHTTPClient(client *http.Client) *GetDirectorySyncExecutionAlertsUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get directory sync execution alerts using p o s t params
func (o *GetDirectorySyncExecutionAlertsUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVidmGetDirSyncExecAlertsReqDTO adds the vidmGetDirSyncExecAlertsReqDTO to the get directory sync execution alerts using p o s t params
func (o *GetDirectorySyncExecutionAlertsUsingPOSTParams) WithVidmGetDirSyncExecAlertsReqDTO(vidmGetDirSyncExecAlertsReqDTO *models.VidmGetDirSyncExecAlertsReqDTO) *GetDirectorySyncExecutionAlertsUsingPOSTParams {
	o.SetVidmGetDirSyncExecAlertsReqDTO(vidmGetDirSyncExecAlertsReqDTO)
	return o
}

// SetVidmGetDirSyncExecAlertsReqDTO adds the vidmGetDirSyncExecAlertsReqDTO to the get directory sync execution alerts using p o s t params
func (o *GetDirectorySyncExecutionAlertsUsingPOSTParams) SetVidmGetDirSyncExecAlertsReqDTO(vidmGetDirSyncExecAlertsReqDTO *models.VidmGetDirSyncExecAlertsReqDTO) {
	o.VidmGetDirSyncExecAlertsReqDTO = vidmGetDirSyncExecAlertsReqDTO
}

// WriteToRequest writes these params to a swagger request
func (o *GetDirectorySyncExecutionAlertsUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.VidmGetDirSyncExecAlertsReqDTO != nil {
		if err := r.SetBodyParam(o.VidmGetDirSyncExecAlertsReqDTO); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
