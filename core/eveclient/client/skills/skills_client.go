// Code generated by go-swagger; DO NOT EDIT.

package skills

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new skills API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for skills API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCharactersCharacterIDAttributes(params *GetCharactersCharacterIDAttributesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDAttributesOK, error)

	GetCharactersCharacterIDSkillqueue(params *GetCharactersCharacterIDSkillqueueParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDSkillqueueOK, error)

	GetCharactersCharacterIDSkills(params *GetCharactersCharacterIDSkillsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDSkillsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCharactersCharacterIDAttributes gets character attributes

  Return attributes of a character

---
Alternate route: `/dev/characters/{character_id}/attributes/`

Alternate route: `/legacy/characters/{character_id}/attributes/`

Alternate route: `/v1/characters/{character_id}/attributes/`

---
This route is cached for up to 120 seconds
*/
func (a *Client) GetCharactersCharacterIDAttributes(params *GetCharactersCharacterIDAttributesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDAttributesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDAttributesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_attributes",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/attributes/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDAttributesReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDAttributesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_attributes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCharactersCharacterIDSkillqueue gets character s skill queue

  List the configured skill queue for the given character

---
Alternate route: `/dev/characters/{character_id}/skillqueue/`

Alternate route: `/legacy/characters/{character_id}/skillqueue/`

Alternate route: `/v2/characters/{character_id}/skillqueue/`

---
This route is cached for up to 120 seconds
*/
func (a *Client) GetCharactersCharacterIDSkillqueue(params *GetCharactersCharacterIDSkillqueueParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDSkillqueueOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDSkillqueueParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_skillqueue",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/skillqueue/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDSkillqueueReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDSkillqueueOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_skillqueue: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCharactersCharacterIDSkills gets character skills

  List all trained skills for the given character

---
Alternate route: `/dev/characters/{character_id}/skills/`

Alternate route: `/v4/characters/{character_id}/skills/`

---
This route is cached for up to 120 seconds
*/
func (a *Client) GetCharactersCharacterIDSkills(params *GetCharactersCharacterIDSkillsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDSkillsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDSkillsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_skills",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/skills/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDSkillsReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDSkillsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_skills: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}