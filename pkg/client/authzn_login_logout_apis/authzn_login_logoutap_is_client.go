// Code generated by go-swagger; DO NOT EDIT.

package authzn_login_logout_a_p_is

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new authzn login logout a p is API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for authzn login logout a p is API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	LoginBasicUsingPOST(params *LoginBasicUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*LoginBasicUsingPOSTOK, *LoginBasicUsingPOSTCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  LoginBasicUsingPOST logins basic
*/
func (a *Client) LoginBasicUsingPOST(params *LoginBasicUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*LoginBasicUsingPOSTOK, *LoginBasicUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoginBasicUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "loginBasicUsingPOST",
		Method:             "POST",
		PathPattern:        "/lcm/authzn/api/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LoginBasicUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *LoginBasicUsingPOSTOK:
		return value, nil, nil
	case *LoginBasicUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for authzn_login_logout_a_p_is: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}