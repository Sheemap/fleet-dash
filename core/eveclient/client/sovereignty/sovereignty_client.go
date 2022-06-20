// Code generated by go-swagger; DO NOT EDIT.

package sovereignty

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new sovereignty API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sovereignty API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetSovereigntyCampaigns(params *GetSovereigntyCampaignsParams, opts ...ClientOption) (*GetSovereigntyCampaignsOK, error)

	GetSovereigntyMap(params *GetSovereigntyMapParams, opts ...ClientOption) (*GetSovereigntyMapOK, error)

	GetSovereigntyStructures(params *GetSovereigntyStructuresParams, opts ...ClientOption) (*GetSovereigntyStructuresOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetSovereigntyCampaigns lists sovereignty campaigns

  Shows sovereignty data for campaigns.

---
Alternate route: `/dev/sovereignty/campaigns/`

Alternate route: `/legacy/sovereignty/campaigns/`

Alternate route: `/v1/sovereignty/campaigns/`

---
This route is cached for up to 5 seconds
*/
func (a *Client) GetSovereigntyCampaigns(params *GetSovereigntyCampaignsParams, opts ...ClientOption) (*GetSovereigntyCampaignsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSovereigntyCampaignsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_sovereignty_campaigns",
		Method:             "GET",
		PathPattern:        "/sovereignty/campaigns/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSovereigntyCampaignsReader{formats: a.formats},
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
	success, ok := result.(*GetSovereigntyCampaignsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_sovereignty_campaigns: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSovereigntyMap lists sovereignty of systems

  Shows sovereignty information for solar systems

---
Alternate route: `/dev/sovereignty/map/`

Alternate route: `/legacy/sovereignty/map/`

Alternate route: `/v1/sovereignty/map/`

---
This route is cached for up to 3600 seconds
*/
func (a *Client) GetSovereigntyMap(params *GetSovereigntyMapParams, opts ...ClientOption) (*GetSovereigntyMapOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSovereigntyMapParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_sovereignty_map",
		Method:             "GET",
		PathPattern:        "/sovereignty/map/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSovereigntyMapReader{formats: a.formats},
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
	success, ok := result.(*GetSovereigntyMapOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_sovereignty_map: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSovereigntyStructures lists sovereignty structures

  Shows sovereignty data for structures.

---
Alternate route: `/dev/sovereignty/structures/`

Alternate route: `/legacy/sovereignty/structures/`

Alternate route: `/v1/sovereignty/structures/`

---
This route is cached for up to 120 seconds
*/
func (a *Client) GetSovereigntyStructures(params *GetSovereigntyStructuresParams, opts ...ClientOption) (*GetSovereigntyStructuresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSovereigntyStructuresParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_sovereignty_structures",
		Method:             "GET",
		PathPattern:        "/sovereignty/structures/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSovereigntyStructuresReader{formats: a.formats},
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
	success, ok := result.(*GetSovereigntyStructuresOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_sovereignty_structures: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}