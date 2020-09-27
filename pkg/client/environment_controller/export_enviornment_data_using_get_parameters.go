// Code generated by go-swagger; DO NOT EDIT.

package environment_controller

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
)

// NewExportEnviornmentDataUsingGETParams creates a new ExportEnviornmentDataUsingGETParams object
// with the default values initialized.
func NewExportEnviornmentDataUsingGETParams() *ExportEnviornmentDataUsingGETParams {
	var ()
	return &ExportEnviornmentDataUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExportEnviornmentDataUsingGETParamsWithTimeout creates a new ExportEnviornmentDataUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExportEnviornmentDataUsingGETParamsWithTimeout(timeout time.Duration) *ExportEnviornmentDataUsingGETParams {
	var ()
	return &ExportEnviornmentDataUsingGETParams{

		timeout: timeout,
	}
}

// NewExportEnviornmentDataUsingGETParamsWithContext creates a new ExportEnviornmentDataUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewExportEnviornmentDataUsingGETParamsWithContext(ctx context.Context) *ExportEnviornmentDataUsingGETParams {
	var ()
	return &ExportEnviornmentDataUsingGETParams{

		Context: ctx,
	}
}

// NewExportEnviornmentDataUsingGETParamsWithHTTPClient creates a new ExportEnviornmentDataUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExportEnviornmentDataUsingGETParamsWithHTTPClient(client *http.Client) *ExportEnviornmentDataUsingGETParams {
	var ()
	return &ExportEnviornmentDataUsingGETParams{
		HTTPClient: client,
	}
}

/*ExportEnviornmentDataUsingGETParams contains all the parameters to send to the API endpoint
for the export enviornment data using g e t operation typically these are written to a http.Request
*/
type ExportEnviornmentDataUsingGETParams struct {

	/*EnvironmentID
	  environmentId

	*/
	EnvironmentID string
	/*Type
	  type

	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the export enviornment data using g e t params
func (o *ExportEnviornmentDataUsingGETParams) WithTimeout(timeout time.Duration) *ExportEnviornmentDataUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export enviornment data using g e t params
func (o *ExportEnviornmentDataUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export enviornment data using g e t params
func (o *ExportEnviornmentDataUsingGETParams) WithContext(ctx context.Context) *ExportEnviornmentDataUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export enviornment data using g e t params
func (o *ExportEnviornmentDataUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export enviornment data using g e t params
func (o *ExportEnviornmentDataUsingGETParams) WithHTTPClient(client *http.Client) *ExportEnviornmentDataUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export enviornment data using g e t params
func (o *ExportEnviornmentDataUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the export enviornment data using g e t params
func (o *ExportEnviornmentDataUsingGETParams) WithEnvironmentID(environmentID string) *ExportEnviornmentDataUsingGETParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the export enviornment data using g e t params
func (o *ExportEnviornmentDataUsingGETParams) SetEnvironmentID(environmentID string) {
	o.EnvironmentID = environmentID
}

// WithType adds the typeVar to the export enviornment data using g e t params
func (o *ExportEnviornmentDataUsingGETParams) WithType(typeVar string) *ExportEnviornmentDataUsingGETParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the export enviornment data using g e t params
func (o *ExportEnviornmentDataUsingGETParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *ExportEnviornmentDataUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param environmentId
	if err := r.SetPathParam("environmentId", o.EnvironmentID); err != nil {
		return err
	}

	// query param type
	qrType := o.Type
	qType := qrType
	if qType != "" {
		if err := r.SetQueryParam("type", qType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
