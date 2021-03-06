// Code generated by go-swagger; DO NOT EDIT.

package data_center_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new data center controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for data center controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDataCenterUsingPOST(params *CreateDataCenterUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDataCenterUsingPOSTOK, *CreateDataCenterUsingPOSTCreated, error)

	CreateRegionUsingPOST(params *CreateRegionUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*CreateRegionUsingPOSTOK, *CreateRegionUsingPOSTCreated, error)

	CreateVCenterUsingPOST(params *CreateVCenterUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*CreateVCenterUsingPOSTOK, *CreateVCenterUsingPOSTCreated, error)

	CreateZoneUsingPOST(params *CreateZoneUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*CreateZoneUsingPOSTOK, *CreateZoneUsingPOSTCreated, error)

	DeleteDataCentersUsingDELETE(params *DeleteDataCentersUsingDELETEParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDataCentersUsingDELETEOK, *DeleteDataCentersUsingDELETENoContent, error)

	GetAllDataCentersUsingGET(params *GetAllDataCentersUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllDataCentersUsingGETOK, error)

	GetAllRegionsUsingGET(params *GetAllRegionsUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllRegionsUsingGETOK, error)

	GetAllVCentersUsingGET(params *GetAllVCentersUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllVCentersUsingGETOK, error)

	GetAllZonesUsingGET(params *GetAllZonesUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllZonesUsingGETOK, error)

	GetDataCenterUsingGET(params *GetDataCenterUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetDataCenterUsingGETOK, error)

	GetRegionUsingGET(params *GetRegionUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetRegionUsingGETOK, error)

	GetVCenterUsingGET(params *GetVCenterUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetVCenterUsingGETOK, error)

	GetVCenterUsingGET1(params *GetVCenterUsingGET1Params, authInfo runtime.ClientAuthInfoWriter) (*GetVCenterUsingGET1OK, error)

	GetZoneUsingGET(params *GetZoneUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetZoneUsingGETOK, error)

	ImportVCentersUsingPOST(params *ImportVCentersUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*ImportVCentersUsingPOSTOK, *ImportVCentersUsingPOSTCreated, error)

	TriggerVCenterDataCollectionUsingPUT(params *TriggerVCenterDataCollectionUsingPUTParams, authInfo runtime.ClientAuthInfoWriter) (*TriggerVCenterDataCollectionUsingPUTOK, *TriggerVCenterDataCollectionUsingPUTCreated, error)

	UpdateDataCenterUsingPUT(params *UpdateDataCenterUsingPUTParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateDataCenterUsingPUTOK, *UpdateDataCenterUsingPUTCreated, error)

	ValidateImportVCentersUsingPOST(params *ValidateImportVCentersUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateImportVCentersUsingPOSTOK, *ValidateImportVCentersUsingPOSTCreated, error)

	ValidateVCenterUsingPOST(params *ValidateVCenterUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateVCenterUsingPOSTOK, *ValidateVCenterUsingPOSTCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateDataCenterUsingPOST creates data center
*/
func (a *Client) CreateDataCenterUsingPOST(params *CreateDataCenterUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDataCenterUsingPOSTOK, *CreateDataCenterUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDataCenterUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createDataCenterUsingPOST",
		Method:             "POST",
		PathPattern:        "/lcm/lcops/api/datacenters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateDataCenterUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateDataCenterUsingPOSTOK:
		return value, nil, nil
	case *CreateDataCenterUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for data_center_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateRegionUsingPOST creates region
*/
func (a *Client) CreateRegionUsingPOST(params *CreateRegionUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*CreateRegionUsingPOSTOK, *CreateRegionUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRegionUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createRegionUsingPOST",
		Method:             "POST",
		PathPattern:        "/lcm/lcops/api/datacenters/{dataCenterName}/regions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateRegionUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateRegionUsingPOSTOK:
		return value, nil, nil
	case *CreateRegionUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for data_center_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateVCenterUsingPOST creates v center
*/
func (a *Client) CreateVCenterUsingPOST(params *CreateVCenterUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*CreateVCenterUsingPOSTOK, *CreateVCenterUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVCenterUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createVCenterUsingPOST",
		Method:             "POST",
		PathPattern:        "/lcm/lcops/api/datacenters/{dataCenterName}/regions/{regionName}/zones/{zoneName}/vCenters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateVCenterUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateVCenterUsingPOSTOK:
		return value, nil, nil
	case *CreateVCenterUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for data_center_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateZoneUsingPOST creates zone
*/
func (a *Client) CreateZoneUsingPOST(params *CreateZoneUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*CreateZoneUsingPOSTOK, *CreateZoneUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateZoneUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createZoneUsingPOST",
		Method:             "POST",
		PathPattern:        "/lcm/lcops/api/datacenters/{dataCenterName}/regions/{regionName}/zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateZoneUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateZoneUsingPOSTOK:
		return value, nil, nil
	case *CreateZoneUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for data_center_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteDataCentersUsingDELETE deletes data centers
*/
func (a *Client) DeleteDataCentersUsingDELETE(params *DeleteDataCentersUsingDELETEParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDataCentersUsingDELETEOK, *DeleteDataCentersUsingDELETENoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDataCentersUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDataCentersUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/lcm/lcops/api/datacenters/{dataCenterName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteDataCentersUsingDELETEReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteDataCentersUsingDELETEOK:
		return value, nil, nil
	case *DeleteDataCentersUsingDELETENoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for data_center_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllDataCentersUsingGET gets all data centers
*/
func (a *Client) GetAllDataCentersUsingGET(params *GetAllDataCentersUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllDataCentersUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllDataCentersUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllDataCentersUsingGET",
		Method:             "GET",
		PathPattern:        "/lcm/lcops/api/datacenters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllDataCentersUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllDataCentersUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllDataCentersUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllRegionsUsingGET gets all regions
*/
func (a *Client) GetAllRegionsUsingGET(params *GetAllRegionsUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllRegionsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllRegionsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllRegionsUsingGET",
		Method:             "GET",
		PathPattern:        "/lcm/lcops/api/datacenters/{dataCenterName}/regions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllRegionsUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllRegionsUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllRegionsUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllVCentersUsingGET gets all v centers
*/
func (a *Client) GetAllVCentersUsingGET(params *GetAllVCentersUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllVCentersUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllVCentersUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllVCentersUsingGET",
		Method:             "GET",
		PathPattern:        "/lcm/lcops/api/datacenters/{dataCenterName}/regions/{regionName}/zones/{zoneName}/vCenters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllVCentersUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllVCentersUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllVCentersUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllZonesUsingGET gets all zones
*/
func (a *Client) GetAllZonesUsingGET(params *GetAllZonesUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllZonesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllZonesUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllZonesUsingGET",
		Method:             "GET",
		PathPattern:        "/lcm/lcops/api/datacenters/{dataCenterName}/regions/{regionName}/zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllZonesUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllZonesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllZonesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDataCenterUsingGET gets data center
*/
func (a *Client) GetDataCenterUsingGET(params *GetDataCenterUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetDataCenterUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDataCenterUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDataCenterUsingGET",
		Method:             "GET",
		PathPattern:        "/lcm/lcops/api/datacenters/{dataCenterName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDataCenterUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDataCenterUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDataCenterUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRegionUsingGET gets region
*/
func (a *Client) GetRegionUsingGET(params *GetRegionUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetRegionUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRegionUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRegionUsingGET",
		Method:             "GET",
		PathPattern:        "/lcm/lcops/api/datacenters/{dataCenterName}/regions/{regionName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRegionUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRegionUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRegionUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVCenterUsingGET gets v center
*/
func (a *Client) GetVCenterUsingGET(params *GetVCenterUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetVCenterUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVCenterUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVCenterUsingGET",
		Method:             "GET",
		PathPattern:        "/lcm/lcops/api/datacenters/{dataCenterName}/regions/{regionName}/zones/{zoneName}/vCenters/{vCenterName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVCenterUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVCenterUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVCenterUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVCenterUsingGET1 gets v center
*/
func (a *Client) GetVCenterUsingGET1(params *GetVCenterUsingGET1Params, authInfo runtime.ClientAuthInfoWriter) (*GetVCenterUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVCenterUsingGET1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVCenterUsingGET_1",
		Method:             "GET",
		PathPattern:        "/lcm/lcops/api/datacenters/{dataCenterName}/vCenter/{vCenterName}/query/{queryPattern}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVCenterUsingGET1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVCenterUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVCenterUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetZoneUsingGET gets zone
*/
func (a *Client) GetZoneUsingGET(params *GetZoneUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetZoneUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetZoneUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getZoneUsingGET",
		Method:             "GET",
		PathPattern:        "/lcm/lcops/api/datacenters/{dataCenterName}/regions/{regionName}/zones/{zoneName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetZoneUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetZoneUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getZoneUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImportVCentersUsingPOST imports v centers
*/
func (a *Client) ImportVCentersUsingPOST(params *ImportVCentersUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*ImportVCentersUsingPOSTOK, *ImportVCentersUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportVCentersUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "importVCentersUsingPOST",
		Method:             "POST",
		PathPattern:        "/lcm/lcops/api/datacenters/{dataCenterName}/regions/{regionName}/zones/{zoneName}/importvCenters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ImportVCentersUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ImportVCentersUsingPOSTOK:
		return value, nil, nil
	case *ImportVCentersUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for data_center_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TriggerVCenterDataCollectionUsingPUT triggers v center data collection
*/
func (a *Client) TriggerVCenterDataCollectionUsingPUT(params *TriggerVCenterDataCollectionUsingPUTParams, authInfo runtime.ClientAuthInfoWriter) (*TriggerVCenterDataCollectionUsingPUTOK, *TriggerVCenterDataCollectionUsingPUTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTriggerVCenterDataCollectionUsingPUTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "triggerVCenterDataCollectionUsingPUT",
		Method:             "PUT",
		PathPattern:        "/lcm/lcops/api/datacenters/{dataCenterName}/regions/{regionName}/zones/{zoneName}/vCenters/{vCenterName}/datacollection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TriggerVCenterDataCollectionUsingPUTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *TriggerVCenterDataCollectionUsingPUTOK:
		return value, nil, nil
	case *TriggerVCenterDataCollectionUsingPUTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for data_center_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateDataCenterUsingPUT updates data center
*/
func (a *Client) UpdateDataCenterUsingPUT(params *UpdateDataCenterUsingPUTParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateDataCenterUsingPUTOK, *UpdateDataCenterUsingPUTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDataCenterUsingPUTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateDataCenterUsingPUT",
		Method:             "PUT",
		PathPattern:        "/lcm/lcops/api/datacenters/{dataCenterName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateDataCenterUsingPUTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateDataCenterUsingPUTOK:
		return value, nil, nil
	case *UpdateDataCenterUsingPUTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for data_center_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateImportVCentersUsingPOST validates import v centers
*/
func (a *Client) ValidateImportVCentersUsingPOST(params *ValidateImportVCentersUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateImportVCentersUsingPOSTOK, *ValidateImportVCentersUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateImportVCentersUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateImportVCentersUsingPOST",
		Method:             "POST",
		PathPattern:        "/lcm/lcops/api/datacenters/{dataCenterName}/regions/{regionName}/zones/{zoneName}/validatevCenters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ValidateImportVCentersUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ValidateImportVCentersUsingPOSTOK:
		return value, nil, nil
	case *ValidateImportVCentersUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for data_center_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateVCenterUsingPOST validates v center
*/
func (a *Client) ValidateVCenterUsingPOST(params *ValidateVCenterUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateVCenterUsingPOSTOK, *ValidateVCenterUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateVCenterUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateVCenterUsingPOST",
		Method:             "POST",
		PathPattern:        "/lcm/lcops/api/datacenters/{dataCenterName}/regions/{regionName}/zones/{zoneName}/vCenters/{vCenterName}/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ValidateVCenterUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ValidateVCenterUsingPOSTOK:
		return value, nil, nil
	case *ValidateVCenterUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for data_center_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
