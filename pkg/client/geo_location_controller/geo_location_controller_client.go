// Code generated by go-swagger; DO NOT EDIT.

package geo_location_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new geo location controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for geo location controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	SearchGeoLocationsUsingGET(params *SearchGeoLocationsUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*SearchGeoLocationsUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SearchGeoLocationsUsingGET searches geo locations
*/
func (a *Client) SearchGeoLocationsUsingGET(params *SearchGeoLocationsUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*SearchGeoLocationsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchGeoLocationsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "searchGeoLocationsUsingGET",
		Method:             "GET",
		PathPattern:        "/lcm/lcops/api/geolocation/search/{pattern}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SearchGeoLocationsUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchGeoLocationsUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchGeoLocationsUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
