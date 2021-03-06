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
)

// NewGetChildrenUsingGETParams creates a new GetChildrenUsingGETParams object
// with the default values initialized.
func NewGetChildrenUsingGETParams() *GetChildrenUsingGETParams {
	var ()
	return &GetChildrenUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetChildrenUsingGETParamsWithTimeout creates a new GetChildrenUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetChildrenUsingGETParamsWithTimeout(timeout time.Duration) *GetChildrenUsingGETParams {
	var ()
	return &GetChildrenUsingGETParams{

		timeout: timeout,
	}
}

// NewGetChildrenUsingGETParamsWithContext creates a new GetChildrenUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetChildrenUsingGETParamsWithContext(ctx context.Context) *GetChildrenUsingGETParams {
	var ()
	return &GetChildrenUsingGETParams{

		Context: ctx,
	}
}

// NewGetChildrenUsingGETParamsWithHTTPClient creates a new GetChildrenUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetChildrenUsingGETParamsWithHTTPClient(client *http.Client) *GetChildrenUsingGETParams {
	var ()
	return &GetChildrenUsingGETParams{
		HTTPClient: client,
	}
}

/*GetChildrenUsingGETParams contains all the parameters to send to the API endpoint
for the get children using g e t operation typically these are written to a http.Request
*/
type GetChildrenUsingGETParams struct {

	/*ParentID
	  parentId

	*/
	ParentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get children using g e t params
func (o *GetChildrenUsingGETParams) WithTimeout(timeout time.Duration) *GetChildrenUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get children using g e t params
func (o *GetChildrenUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get children using g e t params
func (o *GetChildrenUsingGETParams) WithContext(ctx context.Context) *GetChildrenUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get children using g e t params
func (o *GetChildrenUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get children using g e t params
func (o *GetChildrenUsingGETParams) WithHTTPClient(client *http.Client) *GetChildrenUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get children using g e t params
func (o *GetChildrenUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParentID adds the parentID to the get children using g e t params
func (o *GetChildrenUsingGETParams) WithParentID(parentID string) *GetChildrenUsingGETParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the get children using g e t params
func (o *GetChildrenUsingGETParams) SetParentID(parentID string) {
	o.ParentID = parentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetChildrenUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param parentId
	if err := r.SetPathParam("parentId", o.ParentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
