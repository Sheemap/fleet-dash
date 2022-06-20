// Code generated by go-swagger; DO NOT EDIT.

package contracts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new contracts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for contracts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCharactersCharacterIDContracts(params *GetCharactersCharacterIDContractsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDContractsOK, error)

	GetCharactersCharacterIDContractsContractIDBids(params *GetCharactersCharacterIDContractsContractIDBidsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDContractsContractIDBidsOK, error)

	GetCharactersCharacterIDContractsContractIDItems(params *GetCharactersCharacterIDContractsContractIDItemsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDContractsContractIDItemsOK, error)

	GetContractsPublicBidsContractID(params *GetContractsPublicBidsContractIDParams, opts ...ClientOption) (*GetContractsPublicBidsContractIDOK, *GetContractsPublicBidsContractIDNoContent, error)

	GetContractsPublicItemsContractID(params *GetContractsPublicItemsContractIDParams, opts ...ClientOption) (*GetContractsPublicItemsContractIDOK, *GetContractsPublicItemsContractIDNoContent, error)

	GetContractsPublicRegionID(params *GetContractsPublicRegionIDParams, opts ...ClientOption) (*GetContractsPublicRegionIDOK, error)

	GetCorporationsCorporationIDContracts(params *GetCorporationsCorporationIDContractsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDContractsOK, error)

	GetCorporationsCorporationIDContractsContractIDBids(params *GetCorporationsCorporationIDContractsContractIDBidsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDContractsContractIDBidsOK, error)

	GetCorporationsCorporationIDContractsContractIDItems(params *GetCorporationsCorporationIDContractsContractIDItemsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDContractsContractIDItemsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCharactersCharacterIDContracts gets contracts

  Returns contracts available to a character, only if the character is issuer, acceptor or assignee. Only returns contracts no older than 30 days, or if the status is "in_progress".

---
Alternate route: `/dev/characters/{character_id}/contracts/`

Alternate route: `/legacy/characters/{character_id}/contracts/`

Alternate route: `/v1/characters/{character_id}/contracts/`

---
This route is cached for up to 300 seconds
*/
func (a *Client) GetCharactersCharacterIDContracts(params *GetCharactersCharacterIDContractsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDContractsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDContractsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_contracts",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/contracts/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDContractsReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDContractsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_contracts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCharactersCharacterIDContractsContractIDBids gets contract bids

  Lists bids on a particular auction contract

---
Alternate route: `/dev/characters/{character_id}/contracts/{contract_id}/bids/`

Alternate route: `/legacy/characters/{character_id}/contracts/{contract_id}/bids/`

Alternate route: `/v1/characters/{character_id}/contracts/{contract_id}/bids/`

---
This route is cached for up to 300 seconds
*/
func (a *Client) GetCharactersCharacterIDContractsContractIDBids(params *GetCharactersCharacterIDContractsContractIDBidsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDContractsContractIDBidsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDContractsContractIDBidsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_contracts_contract_id_bids",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/contracts/{contract_id}/bids/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDContractsContractIDBidsReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDContractsContractIDBidsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_contracts_contract_id_bids: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCharactersCharacterIDContractsContractIDItems gets contract items

  Lists items of a particular contract

---
Alternate route: `/dev/characters/{character_id}/contracts/{contract_id}/items/`

Alternate route: `/legacy/characters/{character_id}/contracts/{contract_id}/items/`

Alternate route: `/v1/characters/{character_id}/contracts/{contract_id}/items/`

---
This route is cached for up to 3600 seconds
*/
func (a *Client) GetCharactersCharacterIDContractsContractIDItems(params *GetCharactersCharacterIDContractsContractIDItemsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDContractsContractIDItemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDContractsContractIDItemsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_contracts_contract_id_items",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/contracts/{contract_id}/items/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDContractsContractIDItemsReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDContractsContractIDItemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_contracts_contract_id_items: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetContractsPublicBidsContractID gets public contract bids

  Lists bids on a public auction contract

---
Alternate route: `/dev/contracts/public/bids/{contract_id}/`

Alternate route: `/legacy/contracts/public/bids/{contract_id}/`

Alternate route: `/v1/contracts/public/bids/{contract_id}/`

---
This route is cached for up to 300 seconds
*/
func (a *Client) GetContractsPublicBidsContractID(params *GetContractsPublicBidsContractIDParams, opts ...ClientOption) (*GetContractsPublicBidsContractIDOK, *GetContractsPublicBidsContractIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetContractsPublicBidsContractIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_contracts_public_bids_contract_id",
		Method:             "GET",
		PathPattern:        "/contracts/public/bids/{contract_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetContractsPublicBidsContractIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetContractsPublicBidsContractIDOK:
		return value, nil, nil
	case *GetContractsPublicBidsContractIDNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for contracts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetContractsPublicItemsContractID gets public contract items

  Lists items of a public contract

---
Alternate route: `/dev/contracts/public/items/{contract_id}/`

Alternate route: `/legacy/contracts/public/items/{contract_id}/`

Alternate route: `/v1/contracts/public/items/{contract_id}/`

---
This route is cached for up to 3600 seconds
*/
func (a *Client) GetContractsPublicItemsContractID(params *GetContractsPublicItemsContractIDParams, opts ...ClientOption) (*GetContractsPublicItemsContractIDOK, *GetContractsPublicItemsContractIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetContractsPublicItemsContractIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_contracts_public_items_contract_id",
		Method:             "GET",
		PathPattern:        "/contracts/public/items/{contract_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetContractsPublicItemsContractIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetContractsPublicItemsContractIDOK:
		return value, nil, nil
	case *GetContractsPublicItemsContractIDNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for contracts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetContractsPublicRegionID gets public contracts

  Returns a paginated list of all public contracts in the given region

---
Alternate route: `/dev/contracts/public/{region_id}/`

Alternate route: `/legacy/contracts/public/{region_id}/`

Alternate route: `/v1/contracts/public/{region_id}/`

---
This route is cached for up to 1800 seconds
*/
func (a *Client) GetContractsPublicRegionID(params *GetContractsPublicRegionIDParams, opts ...ClientOption) (*GetContractsPublicRegionIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetContractsPublicRegionIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_contracts_public_region_id",
		Method:             "GET",
		PathPattern:        "/contracts/public/{region_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetContractsPublicRegionIDReader{formats: a.formats},
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
	success, ok := result.(*GetContractsPublicRegionIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_contracts_public_region_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCorporationsCorporationIDContracts gets corporation contracts

  Returns contracts available to a corporation, only if the corporation is issuer, acceptor or assignee. Only returns contracts no older than 30 days, or if the status is "in_progress".

---
Alternate route: `/dev/corporations/{corporation_id}/contracts/`

Alternate route: `/legacy/corporations/{corporation_id}/contracts/`

Alternate route: `/v1/corporations/{corporation_id}/contracts/`

---
This route is cached for up to 300 seconds
*/
func (a *Client) GetCorporationsCorporationIDContracts(params *GetCorporationsCorporationIDContractsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDContractsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDContractsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_contracts",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/contracts/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDContractsReader{formats: a.formats},
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
	success, ok := result.(*GetCorporationsCorporationIDContractsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_corporations_corporation_id_contracts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCorporationsCorporationIDContractsContractIDBids gets corporation contract bids

  Lists bids on a particular auction contract

---
Alternate route: `/dev/corporations/{corporation_id}/contracts/{contract_id}/bids/`

Alternate route: `/legacy/corporations/{corporation_id}/contracts/{contract_id}/bids/`

Alternate route: `/v1/corporations/{corporation_id}/contracts/{contract_id}/bids/`

---
This route is cached for up to 3600 seconds
*/
func (a *Client) GetCorporationsCorporationIDContractsContractIDBids(params *GetCorporationsCorporationIDContractsContractIDBidsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDContractsContractIDBidsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDContractsContractIDBidsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_contracts_contract_id_bids",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/contracts/{contract_id}/bids/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDContractsContractIDBidsReader{formats: a.formats},
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
	success, ok := result.(*GetCorporationsCorporationIDContractsContractIDBidsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_corporations_corporation_id_contracts_contract_id_bids: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCorporationsCorporationIDContractsContractIDItems gets corporation contract items

  Lists items of a particular contract

---
Alternate route: `/dev/corporations/{corporation_id}/contracts/{contract_id}/items/`

Alternate route: `/legacy/corporations/{corporation_id}/contracts/{contract_id}/items/`

Alternate route: `/v1/corporations/{corporation_id}/contracts/{contract_id}/items/`

---
This route is cached for up to 3600 seconds
*/
func (a *Client) GetCorporationsCorporationIDContractsContractIDItems(params *GetCorporationsCorporationIDContractsContractIDItemsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDContractsContractIDItemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDContractsContractIDItemsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_contracts_contract_id_items",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/contracts/{contract_id}/items/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDContractsContractIDItemsReader{formats: a.formats},
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
	success, ok := result.(*GetCorporationsCorporationIDContractsContractIDItemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_corporations_corporation_id_contracts_contract_id_items: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}