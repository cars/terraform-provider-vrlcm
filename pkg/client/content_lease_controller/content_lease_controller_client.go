// Code generated by go-swagger; DO NOT EDIT.

package content_lease_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new content lease controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for content lease controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CleanContentLeasesUsingDELETE(params *CleanContentLeasesUsingDELETEParams, authInfo runtime.ClientAuthInfoWriter) (*CleanContentLeasesUsingDELETEOK, *CleanContentLeasesUsingDELETENoContent, error)

	CreateContentLeaseUsingPOST(params *CreateContentLeaseUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*CreateContentLeaseUsingPOSTOK, *CreateContentLeaseUsingPOSTCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CleanContentLeasesUsingDELETE cleans content leases
*/
func (a *Client) CleanContentLeasesUsingDELETE(params *CleanContentLeasesUsingDELETEParams, authInfo runtime.ClientAuthInfoWriter) (*CleanContentLeasesUsingDELETEOK, *CleanContentLeasesUsingDELETENoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCleanContentLeasesUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "cleanContentLeasesUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/lcm/crepo/api/leases/clean",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CleanContentLeasesUsingDELETEReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CleanContentLeasesUsingDELETEOK:
		return value, nil, nil
	case *CleanContentLeasesUsingDELETENoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for content_lease_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateContentLeaseUsingPOST creates content lease
*/
func (a *Client) CreateContentLeaseUsingPOST(params *CreateContentLeaseUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*CreateContentLeaseUsingPOSTOK, *CreateContentLeaseUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateContentLeaseUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createContentLeaseUsingPOST",
		Method:             "POST",
		PathPattern:        "/lcm/crepo/api/leases",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateContentLeaseUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateContentLeaseUsingPOSTOK:
		return value, nil, nil
	case *CreateContentLeaseUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for content_lease_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
