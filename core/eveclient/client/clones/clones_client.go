// Code generated by go-swagger; DO NOT EDIT.

package clones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new clones API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for clones API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCharactersCharacterIDClones(params *GetCharactersCharacterIDClonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDClonesOK, error)

	GetCharactersCharacterIDImplants(params *GetCharactersCharacterIDImplantsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDImplantsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCharactersCharacterIDClones gets clones

  A list of the character's clones

---
Alternate route: `/dev/characters/{character_id}/clones/`

Alternate route: `/v3/characters/{character_id}/clones/`

Alternate route: `/v4/characters/{character_id}/clones/`

---
This route is cached for up to 120 seconds
*/
func (a *Client) GetCharactersCharacterIDClones(params *GetCharactersCharacterIDClonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDClonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDClonesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_clones",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/clones/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDClonesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCharactersCharacterIDClonesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_clones: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCharactersCharacterIDImplants gets active implants

  Return implants on the active clone of a character

---
Alternate route: `/dev/characters/{character_id}/implants/`

Alternate route: `/legacy/characters/{character_id}/implants/`

Alternate route: `/v1/characters/{character_id}/implants/`

Alternate route: `/v2/characters/{character_id}/implants/`

---
This route is cached for up to 120 seconds
*/
func (a *Client) GetCharactersCharacterIDImplants(params *GetCharactersCharacterIDImplantsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDImplantsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDImplantsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_implants",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/implants/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDImplantsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCharactersCharacterIDImplantsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_implants: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
