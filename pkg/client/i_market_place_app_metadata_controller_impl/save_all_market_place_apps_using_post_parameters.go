// Code generated by go-swagger; DO NOT EDIT.





package i_market_place_app_metadata_controller_impl

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
  "github.com/go-openapi/swag"
  "github.com/go-openapi/validate"

  	"github.com/cars/terraform-provider-vrlcm/client/application_controller"
	"github.com/cars/terraform-provider-vrlcm/client/authentication_group_a_p_is"
	"github.com/cars/terraform-provider-vrlcm/client/authentication_group_role_mapping_a_p_is"
	"github.com/cars/terraform-provider-vrlcm/client/authentication_provider_a_p_is"
	"github.com/cars/terraform-provider-vrlcm/client/authentication_role_a_p_is"
	"github.com/cars/terraform-provider-vrlcm/client/authentication_user_a_p_is"
	"github.com/cars/terraform-provider-vrlcm/client/authentication_user_role_mapping_a_p_is"
	"github.com/cars/terraform-provider-vrlcm/client/authzn_login_logout_a_p_is"
	"github.com/cars/terraform-provider-vrlcm/client/authzn_sample_a_p_is"
	"github.com/cars/terraform-provider-vrlcm/client/blackstone_proxy_service_controller"
	"github.com/cars/terraform-provider-vrlcm/client/bootstrap_a_p_is"
	"github.com/cars/terraform-provider-vrlcm/client/command_controller"
	"github.com/cars/terraform-provider-vrlcm/client/configuration_parameter_a_p_is"
	"github.com/cars/terraform-provider-vrlcm/client/content_download_controller"
	"github.com/cars/terraform-provider-vrlcm/client/content_lease_controller"
	"github.com/cars/terraform-provider-vrlcm/client/content_repository_controller"
	"github.com/cars/terraform-provider-vrlcm/client/data_center_controller"
	"github.com/cars/terraform-provider-vrlcm/client/deployments_controller"
	"github.com/cars/terraform-provider-vrlcm/client/directory_management_a_p_is"
	"github.com/cars/terraform-provider-vrlcm/client/engine_interface_controller"
	"github.com/cars/terraform-provider-vrlcm/client/engine_message_controller"
	"github.com/cars/terraform-provider-vrlcm/client/environment_controller"
	"github.com/cars/terraform-provider-vrlcm/client/file_browse_controller"
	"github.com/cars/terraform-provider-vrlcm/client/geo_location_controller"
	"github.com/cars/terraform-provider-vrlcm/client/health_a_p_is"
	"github.com/cars/terraform-provider-vrlcm/client/i_market_place_app_metadata_controller_impl"
	"github.com/cars/terraform-provider-vrlcm/client/inventory_query_controller"
	"github.com/cars/terraform-provider-vrlcm/client/inventory_schema_query_controller"
	"github.com/cars/terraform-provider-vrlcm/client/lcm_management_controller"
	"github.com/cars/terraform-provider-vrlcm/client/lcm_update_controller"
	"github.com/cars/terraform-provider-vrlcm/client/locker_certificates_api"
	"github.com/cars/terraform-provider-vrlcm/client/locker_password_and_license_api"
	"github.com/cars/terraform-provider-vrlcm/client/locker_references_api"
	"github.com/cars/terraform-provider-vrlcm/client/metadata_action_controller"
	"github.com/cars/terraform-provider-vrlcm/client/metadata_controller"
	"github.com/cars/terraform-provider-vrlcm/client/migration_controller"
	"github.com/cars/terraform-provider-vrlcm/client/node_controller"
	"github.com/cars/terraform-provider-vrlcm/client/notification_client_controller"
	"github.com/cars/terraform-provider-vrlcm/client/prevalidation_report_controller"
	"github.com/cars/terraform-provider-vrlcm/client/product_policy_controller"
	"github.com/cars/terraform-provider-vrlcm/client/product_reference_a_p_is"
	"github.com/cars/terraform-provider-vrlcm/client/property_controller"
	"github.com/cars/terraform-provider-vrlcm/client/property_interface_controller_impl"
	"github.com/cars/terraform-provider-vrlcm/client/property_schema_controller"
	"github.com/cars/terraform-provider-vrlcm/client/request_controller"
	"github.com/cars/terraform-provider-vrlcm/client/request_definition_controller"
	"github.com/cars/terraform-provider-vrlcm/client/resource_controller"
	"github.com/cars/terraform-provider-vrlcm/client/resource_group_schema_controller"
	"github.com/cars/terraform-provider-vrlcm/client/resource_schema_controller"
	"github.com/cars/terraform-provider-vrlcm/client/root_node_controller"
	"github.com/cars/terraform-provider-vrlcm/client/settings_controller"
	"github.com/cars/terraform-provider-vrlcm/client/tenant_management_a_p_is"
	"github.com/cars/terraform-provider-vrlcm/client/vm_ware_hello_controller"
	"github.com/cars/terraform-provider-vrlcm/models"
  	"github.com/cars/terraform-provider-vrlcm/client"
)

// NewSaveAllMarketPlaceAppsUsingPOSTParams creates a new SaveAllMarketPlaceAppsUsingPOSTParams object
// with the default values initialized.
func NewSaveAllMarketPlaceAppsUsingPOSTParams() *SaveAllMarketPlaceAppsUsingPOSTParams {
  var (
  
  )
  return &SaveAllMarketPlaceAppsUsingPOSTParams{
  
    timeout: cr.DefaultTimeout,
  }
}

// NewSaveAllMarketPlaceAppsUsingPOSTParamsWithTimeout creates a new SaveAllMarketPlaceAppsUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSaveAllMarketPlaceAppsUsingPOSTParamsWithTimeout(timeout time.Duration) *SaveAllMarketPlaceAppsUsingPOSTParams {
  var (
  
  )
  return &SaveAllMarketPlaceAppsUsingPOSTParams{
  
    timeout: timeout,
  }
}

// NewSaveAllMarketPlaceAppsUsingPOSTParamsWithContext creates a new SaveAllMarketPlaceAppsUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewSaveAllMarketPlaceAppsUsingPOSTParamsWithContext(ctx context.Context) *SaveAllMarketPlaceAppsUsingPOSTParams {
  var (
  
  )
  return &SaveAllMarketPlaceAppsUsingPOSTParams{
  
    Context: ctx,
  }
}

// NewSaveAllMarketPlaceAppsUsingPOSTParamsWithHTTPClient creates a new SaveAllMarketPlaceAppsUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSaveAllMarketPlaceAppsUsingPOSTParamsWithHTTPClient(client *http.Client) *SaveAllMarketPlaceAppsUsingPOSTParams {
  var (
  
  )
  return &SaveAllMarketPlaceAppsUsingPOSTParams{
  HTTPClient: client,
  }
}

/*SaveAllMarketPlaceAppsUsingPOSTParams contains all the parameters to send to the API endpoint
for the save all market place apps using p o s t operation typically these are written to a http.Request
*/
type SaveAllMarketPlaceAppsUsingPOSTParams struct {

  /*AppsDto
  appsDto

  */
  AppsDto []
  

  timeout time.Duration
  Context context.Context
  HTTPClient *http.Client
}

// WithTimeout adds the timeout to the save all market place apps using p o s t params
func (o *SaveAllMarketPlaceAppsUsingPOSTParams) WithTimeout(timeout time.Duration) *SaveAllMarketPlaceAppsUsingPOSTParams {
  o.SetTimeout(timeout)
  return o
}

// SetTimeout adds the timeout to the save all market place apps using p o s t params
func (o *SaveAllMarketPlaceAppsUsingPOSTParams) SetTimeout(timeout time.Duration) {
  o.timeout = timeout
}

// WithContext adds the context to the save all market place apps using p o s t params
func (o *SaveAllMarketPlaceAppsUsingPOSTParams) WithContext(ctx context.Context) *SaveAllMarketPlaceAppsUsingPOSTParams {
  o.SetContext(ctx)
  return o
}

// SetContext adds the context to the save all market place apps using p o s t params
func (o *SaveAllMarketPlaceAppsUsingPOSTParams) SetContext(ctx context.Context) {
  o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the save all market place apps using p o s t params
func (o *SaveAllMarketPlaceAppsUsingPOSTParams) WithHTTPClient(client *http.Client) *SaveAllMarketPlaceAppsUsingPOSTParams {
  o.SetHTTPClient(client)
  return o
}

// SetHTTPClient adds the HTTPClient to the save all market place apps using p o s t params
func (o *SaveAllMarketPlaceAppsUsingPOSTParams) SetHTTPClient(client *http.Client) {
  o.HTTPClient = client
}


// WithAppsDto adds the appsDto to the save all market place apps using p o s t params
func (o *SaveAllMarketPlaceAppsUsingPOSTParams) WithAppsDto(appsDto []) *SaveAllMarketPlaceAppsUsingPOSTParams {
  o.SetAppsDto(appsDto)
  return o
}

// SetAppsDto adds the appsDto to the save all market place apps using p o s t params
func (o *SaveAllMarketPlaceAppsUsingPOSTParams) SetAppsDto(appsDto []) {
  o.AppsDto = appsDto
}


// WriteToRequest writes these params to a swagger request
func (o *SaveAllMarketPlaceAppsUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

  if err := r.SetTimeout(o.timeout); err != nil {
    return err
  }
  var res []error
  

  
  valuesAppsDto := o.AppsDto
  
  joinedAppsDto := swag.JoinByFormat(valuesAppsDto, "multi")
  // query array param appsDto
  if err := r.SetQueryParam("appsDto", joinedAppsDto...); err != nil {
    return err
  }
  

  

  

  
  if len(res) > 0 {
    return errors.CompositeValidationError(res...)
  }
  return nil
}