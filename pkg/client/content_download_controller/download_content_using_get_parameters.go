// Code generated by go-swagger; DO NOT EDIT.

package content_download_controller

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

// NewDownloadContentUsingGETParams creates a new DownloadContentUsingGETParams object
// with the default values initialized.
func NewDownloadContentUsingGETParams() *DownloadContentUsingGETParams {
	var ()
	return &DownloadContentUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadContentUsingGETParamsWithTimeout creates a new DownloadContentUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDownloadContentUsingGETParamsWithTimeout(timeout time.Duration) *DownloadContentUsingGETParams {
	var ()
	return &DownloadContentUsingGETParams{

		timeout: timeout,
	}
}

// NewDownloadContentUsingGETParamsWithContext creates a new DownloadContentUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewDownloadContentUsingGETParamsWithContext(ctx context.Context) *DownloadContentUsingGETParams {
	var ()
	return &DownloadContentUsingGETParams{

		Context: ctx,
	}
}

// NewDownloadContentUsingGETParamsWithHTTPClient creates a new DownloadContentUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDownloadContentUsingGETParamsWithHTTPClient(client *http.Client) *DownloadContentUsingGETParams {
	var ()
	return &DownloadContentUsingGETParams{
		HTTPClient: client,
	}
}

/*DownloadContentUsingGETParams contains all the parameters to send to the API endpoint
for the download content using g e t operation typically these are written to a http.Request
*/
type DownloadContentUsingGETParams struct {

	/*ContentPath
	  contentPath

	*/
	ContentPath string
	/*RepoName
	  repoName

	*/
	RepoName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the download content using g e t params
func (o *DownloadContentUsingGETParams) WithTimeout(timeout time.Duration) *DownloadContentUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download content using g e t params
func (o *DownloadContentUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download content using g e t params
func (o *DownloadContentUsingGETParams) WithContext(ctx context.Context) *DownloadContentUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download content using g e t params
func (o *DownloadContentUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download content using g e t params
func (o *DownloadContentUsingGETParams) WithHTTPClient(client *http.Client) *DownloadContentUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download content using g e t params
func (o *DownloadContentUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentPath adds the contentPath to the download content using g e t params
func (o *DownloadContentUsingGETParams) WithContentPath(contentPath string) *DownloadContentUsingGETParams {
	o.SetContentPath(contentPath)
	return o
}

// SetContentPath adds the contentPath to the download content using g e t params
func (o *DownloadContentUsingGETParams) SetContentPath(contentPath string) {
	o.ContentPath = contentPath
}

// WithRepoName adds the repoName to the download content using g e t params
func (o *DownloadContentUsingGETParams) WithRepoName(repoName string) *DownloadContentUsingGETParams {
	o.SetRepoName(repoName)
	return o
}

// SetRepoName adds the repoName to the download content using g e t params
func (o *DownloadContentUsingGETParams) SetRepoName(repoName string) {
	o.RepoName = repoName
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadContentUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contentPath
	if err := r.SetPathParam("contentPath", o.ContentPath); err != nil {
		return err
	}

	// path param repoName
	if err := r.SetPathParam("repoName", o.RepoName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
