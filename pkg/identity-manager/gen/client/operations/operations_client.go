///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Home ans placehold home page will be redirected to if successfully logged in
*/
func (a *Client) Home(params *HomeParams, authInfo runtime.ClientAuthInfoWriter) (*HomeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHomeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "home",
		Method:             "GET",
		PathPattern:        "/v1/iam/home",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &HomeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*HomeOK), nil

}

/*
Redirect redirects to localhost for vs cli login testing
*/
func (a *Client) Redirect(params *RedirectParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRedirectParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "redirect",
		Method:             "GET",
		PathPattern:        "/v1/iam/redirect",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RedirectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
