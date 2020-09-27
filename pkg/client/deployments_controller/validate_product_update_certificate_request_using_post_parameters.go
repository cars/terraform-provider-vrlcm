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

// NewValidateProductUpdateCertificateRequestUsingPOSTParams creates a new ValidateProductUpdateCertificateRequestUsingPOSTParams object
// with the default values initialized.
func NewValidateProductUpdateCertificateRequestUsingPOSTParams() *ValidateProductUpdateCertificateRequestUsingPOSTParams {
	var ()
	return &ValidateProductUpdateCertificateRequestUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateProductUpdateCertificateRequestUsingPOSTParamsWithTimeout creates a new ValidateProductUpdateCertificateRequestUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateProductUpdateCertificateRequestUsingPOSTParamsWithTimeout(timeout time.Duration) *ValidateProductUpdateCertificateRequestUsingPOSTParams {
	var ()
	return &ValidateProductUpdateCertificateRequestUsingPOSTParams{

		timeout: timeout,
	}
}

// NewValidateProductUpdateCertificateRequestUsingPOSTParamsWithContext creates a new ValidateProductUpdateCertificateRequestUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateProductUpdateCertificateRequestUsingPOSTParamsWithContext(ctx context.Context) *ValidateProductUpdateCertificateRequestUsingPOSTParams {
	var ()
	return &ValidateProductUpdateCertificateRequestUsingPOSTParams{

		Context: ctx,
	}
}

// NewValidateProductUpdateCertificateRequestUsingPOSTParamsWithHTTPClient creates a new ValidateProductUpdateCertificateRequestUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateProductUpdateCertificateRequestUsingPOSTParamsWithHTTPClient(client *http.Client) *ValidateProductUpdateCertificateRequestUsingPOSTParams {
	var ()
	return &ValidateProductUpdateCertificateRequestUsingPOSTParams{
		HTTPClient: client,
	}
}

/*ValidateProductUpdateCertificateRequestUsingPOSTParams contains all the parameters to send to the API endpoint
for the validate product update certificate request using p o s t operation typically these are written to a http.Request
*/
type ValidateProductUpdateCertificateRequestUsingPOSTParams struct {

	/*CertificateUpdateProductRequestDTO
	  certificateUpdateProductRequestDTO

	*/
	CertificateUpdateProductRequestDTO *models.UpdateProductCertificateRequestDTO
	/*EnvironmentID
	  environmentId

	*/
	EnvironmentID string
	/*ProductID
	  productId

	*/
	ProductID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate product update certificate request using p o s t params
func (o *ValidateProductUpdateCertificateRequestUsingPOSTParams) WithTimeout(timeout time.Duration) *ValidateProductUpdateCertificateRequestUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate product update certificate request using p o s t params
func (o *ValidateProductUpdateCertificateRequestUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate product update certificate request using p o s t params
func (o *ValidateProductUpdateCertificateRequestUsingPOSTParams) WithContext(ctx context.Context) *ValidateProductUpdateCertificateRequestUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate product update certificate request using p o s t params
func (o *ValidateProductUpdateCertificateRequestUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate product update certificate request using p o s t params
func (o *ValidateProductUpdateCertificateRequestUsingPOSTParams) WithHTTPClient(client *http.Client) *ValidateProductUpdateCertificateRequestUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate product update certificate request using p o s t params
func (o *ValidateProductUpdateCertificateRequestUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCertificateUpdateProductRequestDTO adds the certificateUpdateProductRequestDTO to the validate product update certificate request using p o s t params
func (o *ValidateProductUpdateCertificateRequestUsingPOSTParams) WithCertificateUpdateProductRequestDTO(certificateUpdateProductRequestDTO *models.UpdateProductCertificateRequestDTO) *ValidateProductUpdateCertificateRequestUsingPOSTParams {
	o.SetCertificateUpdateProductRequestDTO(certificateUpdateProductRequestDTO)
	return o
}

// SetCertificateUpdateProductRequestDTO adds the certificateUpdateProductRequestDTO to the validate product update certificate request using p o s t params
func (o *ValidateProductUpdateCertificateRequestUsingPOSTParams) SetCertificateUpdateProductRequestDTO(certificateUpdateProductRequestDTO *models.UpdateProductCertificateRequestDTO) {
	o.CertificateUpdateProductRequestDTO = certificateUpdateProductRequestDTO
}

// WithEnvironmentID adds the environmentID to the validate product update certificate request using p o s t params
func (o *ValidateProductUpdateCertificateRequestUsingPOSTParams) WithEnvironmentID(environmentID string) *ValidateProductUpdateCertificateRequestUsingPOSTParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the validate product update certificate request using p o s t params
func (o *ValidateProductUpdateCertificateRequestUsingPOSTParams) SetEnvironmentID(environmentID string) {
	o.EnvironmentID = environmentID
}

// WithProductID adds the productID to the validate product update certificate request using p o s t params
func (o *ValidateProductUpdateCertificateRequestUsingPOSTParams) WithProductID(productID string) *ValidateProductUpdateCertificateRequestUsingPOSTParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the validate product update certificate request using p o s t params
func (o *ValidateProductUpdateCertificateRequestUsingPOSTParams) SetProductID(productID string) {
	o.ProductID = productID
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateProductUpdateCertificateRequestUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CertificateUpdateProductRequestDTO != nil {
		if err := r.SetBodyParam(o.CertificateUpdateProductRequestDTO); err != nil {
			return err
		}
	}

	// path param environmentId
	if err := r.SetPathParam("environmentId", o.EnvironmentID); err != nil {
		return err
	}

	// path param productId
	if err := r.SetPathParam("productId", o.ProductID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
