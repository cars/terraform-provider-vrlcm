// Code generated by go-swagger; DO NOT EDIT.

package application_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new application controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for application controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateApplicationUsingPOST(params *CreateApplicationUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*CreateApplicationUsingPOSTOK, *CreateApplicationUsingPOSTCreated, error)

	DeleteApplicationUsingDELETE(params *DeleteApplicationUsingDELETEParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteApplicationUsingDELETEOK, *DeleteApplicationUsingDELETENoContent, error)

	ExportApplicationUsingPOST(params *ExportApplicationUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*ExportApplicationUsingPOSTOK, *ExportApplicationUsingPOSTCreated, error)

	GetAllApplicationsUsingGET(params *GetAllApplicationsUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllApplicationsUsingGETOK, error)

	GetPropertyUsingGET(params *GetPropertyUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetPropertyUsingGETOK, error)

	ImportApplicationUsingPOST(params *ImportApplicationUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*ImportApplicationUsingPOSTOK, *ImportApplicationUsingPOSTCreated, error)

	ReplaceApplicationUsingPUT(params *ReplaceApplicationUsingPUTParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceApplicationUsingPUTOK, *ReplaceApplicationUsingPUTCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateApplicationUsingPOST creates application
*/
func (a *Client) CreateApplicationUsingPOST(params *CreateApplicationUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*CreateApplicationUsingPOSTOK, *CreateApplicationUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateApplicationUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createApplicationUsingPOST",
		Method:             "POST",
		PathPattern:        "/lcm/shell/api/apps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateApplicationUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateApplicationUsingPOSTOK:
		return value, nil, nil
	case *CreateApplicationUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for application_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteApplicationUsingDELETE deletes application
*/
func (a *Client) DeleteApplicationUsingDELETE(params *DeleteApplicationUsingDELETEParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteApplicationUsingDELETEOK, *DeleteApplicationUsingDELETENoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteApplicationUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteApplicationUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/lcm/shell/api/apps/{vmid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteApplicationUsingDELETEReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteApplicationUsingDELETEOK:
		return value, nil, nil
	case *DeleteApplicationUsingDELETENoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for application_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExportApplicationUsingPOST exports application
*/
func (a *Client) ExportApplicationUsingPOST(params *ExportApplicationUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*ExportApplicationUsingPOSTOK, *ExportApplicationUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportApplicationUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "exportApplicationUsingPOST",
		Method:             "POST",
		PathPattern:        "/lcm/shell/api/apps/export",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ExportApplicationUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ExportApplicationUsingPOSTOK:
		return value, nil, nil
	case *ExportApplicationUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for application_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllApplicationsUsingGET gets all applications
*/
func (a *Client) GetAllApplicationsUsingGET(params *GetAllApplicationsUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllApplicationsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllApplicationsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllApplicationsUsingGET",
		Method:             "GET",
		PathPattern:        "/lcm/shell/api/apps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllApplicationsUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllApplicationsUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllApplicationsUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPropertyUsingGET gets property
*/
func (a *Client) GetPropertyUsingGET(params *GetPropertyUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetPropertyUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPropertyUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPropertyUsingGET",
		Method:             "GET",
		PathPattern:        "/lcm/shell/api/apps/{vmid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPropertyUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPropertyUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPropertyUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImportApplicationUsingPOST imports application
*/
func (a *Client) ImportApplicationUsingPOST(params *ImportApplicationUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*ImportApplicationUsingPOSTOK, *ImportApplicationUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportApplicationUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "importApplicationUsingPOST",
		Method:             "POST",
		PathPattern:        "/lcm/shell/api/apps/import",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ImportApplicationUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ImportApplicationUsingPOSTOK:
		return value, nil, nil
	case *ImportApplicationUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for application_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReplaceApplicationUsingPUT replaces application
*/
func (a *Client) ReplaceApplicationUsingPUT(params *ReplaceApplicationUsingPUTParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceApplicationUsingPUTOK, *ReplaceApplicationUsingPUTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceApplicationUsingPUTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceApplicationUsingPUT",
		Method:             "PUT",
		PathPattern:        "/lcm/shell/api/apps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceApplicationUsingPUTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceApplicationUsingPUTOK:
		return value, nil, nil
	case *ReplaceApplicationUsingPUTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for application_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
