// Code generated by go-swagger; DO NOT EDIT.

package authentication_role_a_p_is

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new authentication role a p is API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for authentication role a p is API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	BulkImportRolesUsingPOST(params *BulkImportRolesUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*BulkImportRolesUsingPOSTOK, *BulkImportRolesUsingPOSTCreated, error)

	CreateRoleUsingPOST(params *CreateRoleUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*CreateRoleUsingPOSTOK, *CreateRoleUsingPOSTCreated, error)

	DeleteAllRolesUsingDELETE(params *DeleteAllRolesUsingDELETEParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAllRolesUsingDELETEOK, *DeleteAllRolesUsingDELETENoContent, error)

	DeleteRoleByIDUsingDELETE(params *DeleteRoleByIDUsingDELETEParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRoleByIDUsingDELETEOK, *DeleteRoleByIDUsingDELETENoContent, error)

	GetAllExternalRolesUsingGET(params *GetAllExternalRolesUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllExternalRolesUsingGETOK, error)

	GetAllInternalRolesUsingGET(params *GetAllInternalRolesUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllInternalRolesUsingGETOK, error)

	GetAllRolesExtendedUsingGET(params *GetAllRolesExtendedUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllRolesExtendedUsingGETOK, error)

	GetAllRolesUsingGET(params *GetAllRolesUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllRolesUsingGETOK, error)

	GetAllUsersForGiveRoleIDUsingGET(params *GetAllUsersForGiveRoleIDUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllUsersForGiveRoleIDUsingGETOK, error)

	GetRoleByIDExtendedUsingGET(params *GetRoleByIDExtendedUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetRoleByIDExtendedUsingGETOK, error)

	GetRoleByIDUsingGET(params *GetRoleByIDUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetRoleByIDUsingGETOK, error)

	UpdateRoleByIDUsingPATCH(params *UpdateRoleByIDUsingPATCHParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateRoleByIDUsingPATCHOK, *UpdateRoleByIDUsingPATCHNoContent, error)

	UpdateRoleByIDUsingPUT(params *UpdateRoleByIDUsingPUTParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateRoleByIDUsingPUTOK, *UpdateRoleByIDUsingPUTCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BulkImportRolesUsingPOST bulks import roles
*/
func (a *Client) BulkImportRolesUsingPOST(params *BulkImportRolesUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*BulkImportRolesUsingPOSTOK, *BulkImportRolesUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBulkImportRolesUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "bulkImportRolesUsingPOST",
		Method:             "POST",
		PathPattern:        "/lcm/authzn/api/roles/bulkimport",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BulkImportRolesUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *BulkImportRolesUsingPOSTOK:
		return value, nil, nil
	case *BulkImportRolesUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for authentication_role_a_p_is: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateRoleUsingPOST creates role
*/
func (a *Client) CreateRoleUsingPOST(params *CreateRoleUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*CreateRoleUsingPOSTOK, *CreateRoleUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRoleUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createRoleUsingPOST",
		Method:             "POST",
		PathPattern:        "/lcm/authzn/api/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateRoleUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateRoleUsingPOSTOK:
		return value, nil, nil
	case *CreateRoleUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for authentication_role_a_p_is: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteAllRolesUsingDELETE deletes all roles
*/
func (a *Client) DeleteAllRolesUsingDELETE(params *DeleteAllRolesUsingDELETEParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAllRolesUsingDELETEOK, *DeleteAllRolesUsingDELETENoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAllRolesUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteAllRolesUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/lcm/authzn/api/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteAllRolesUsingDELETEReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteAllRolesUsingDELETEOK:
		return value, nil, nil
	case *DeleteAllRolesUsingDELETENoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for authentication_role_a_p_is: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteRoleByIDUsingDELETE deletes role by Id
*/
func (a *Client) DeleteRoleByIDUsingDELETE(params *DeleteRoleByIDUsingDELETEParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRoleByIDUsingDELETEOK, *DeleteRoleByIDUsingDELETENoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRoleByIDUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRoleByIdUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/lcm/authzn/api/roles/{vmid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteRoleByIDUsingDELETEReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteRoleByIDUsingDELETEOK:
		return value, nil, nil
	case *DeleteRoleByIDUsingDELETENoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for authentication_role_a_p_is: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllExternalRolesUsingGET gets all external roles
*/
func (a *Client) GetAllExternalRolesUsingGET(params *GetAllExternalRolesUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllExternalRolesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllExternalRolesUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllExternalRolesUsingGET",
		Method:             "GET",
		PathPattern:        "/lcm/authzn/api/roles/external",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllExternalRolesUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllExternalRolesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllExternalRolesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllInternalRolesUsingGET gets all internal roles
*/
func (a *Client) GetAllInternalRolesUsingGET(params *GetAllInternalRolesUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllInternalRolesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllInternalRolesUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllInternalRolesUsingGET",
		Method:             "GET",
		PathPattern:        "/lcm/authzn/api/roles/internal",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllInternalRolesUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllInternalRolesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllInternalRolesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllRolesExtendedUsingGET gets all roles extended
*/
func (a *Client) GetAllRolesExtendedUsingGET(params *GetAllRolesExtendedUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllRolesExtendedUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllRolesExtendedUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllRolesExtendedUsingGET",
		Method:             "GET",
		PathPattern:        "/lcm/authzn/api/roles/extended",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllRolesExtendedUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllRolesExtendedUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllRolesExtendedUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllRolesUsingGET gets all roles
*/
func (a *Client) GetAllRolesUsingGET(params *GetAllRolesUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllRolesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllRolesUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllRolesUsingGET",
		Method:             "GET",
		PathPattern:        "/lcm/authzn/api/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllRolesUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllRolesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllRolesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllUsersForGiveRoleIDUsingGET gets all users for give role Id
*/
func (a *Client) GetAllUsersForGiveRoleIDUsingGET(params *GetAllUsersForGiveRoleIDUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllUsersForGiveRoleIDUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllUsersForGiveRoleIDUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllUsersForGiveRoleIdUsingGET",
		Method:             "GET",
		PathPattern:        "/lcm/authzn/api/roles/{vmid}/mappedusers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllUsersForGiveRoleIDUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllUsersForGiveRoleIDUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllUsersForGiveRoleIdUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRoleByIDExtendedUsingGET gets role by Id extended
*/
func (a *Client) GetRoleByIDExtendedUsingGET(params *GetRoleByIDExtendedUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetRoleByIDExtendedUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRoleByIDExtendedUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRoleByIdExtendedUsingGET",
		Method:             "GET",
		PathPattern:        "/lcm/authzn/api/roles/{vmid}/extended",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRoleByIDExtendedUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRoleByIDExtendedUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRoleByIdExtendedUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRoleByIDUsingGET gets role by Id
*/
func (a *Client) GetRoleByIDUsingGET(params *GetRoleByIDUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetRoleByIDUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRoleByIDUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRoleByIdUsingGET",
		Method:             "GET",
		PathPattern:        "/lcm/authzn/api/roles/{vmid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRoleByIDUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRoleByIDUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRoleByIdUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateRoleByIDUsingPATCH updates role by Id
*/
func (a *Client) UpdateRoleByIDUsingPATCH(params *UpdateRoleByIDUsingPATCHParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateRoleByIDUsingPATCHOK, *UpdateRoleByIDUsingPATCHNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRoleByIDUsingPATCHParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateRoleByIdUsingPATCH",
		Method:             "PATCH",
		PathPattern:        "/lcm/authzn/api/roles/{vmid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateRoleByIDUsingPATCHReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateRoleByIDUsingPATCHOK:
		return value, nil, nil
	case *UpdateRoleByIDUsingPATCHNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for authentication_role_a_p_is: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateRoleByIDUsingPUT updates role by Id
*/
func (a *Client) UpdateRoleByIDUsingPUT(params *UpdateRoleByIDUsingPUTParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateRoleByIDUsingPUTOK, *UpdateRoleByIDUsingPUTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRoleByIDUsingPUTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateRoleByIdUsingPUT",
		Method:             "PUT",
		PathPattern:        "/lcm/authzn/api/roles/{vmid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateRoleByIDUsingPUTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateRoleByIDUsingPUTOK:
		return value, nil, nil
	case *UpdateRoleByIDUsingPUTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for authentication_role_a_p_is: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}