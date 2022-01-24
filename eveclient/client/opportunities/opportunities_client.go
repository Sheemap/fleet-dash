// Code generated by go-swagger; DO NOT EDIT.

package opportunities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new opportunities API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for opportunities API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCharactersCharacterIDOpportunities(params *GetCharactersCharacterIDOpportunitiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDOpportunitiesOK, error)

	GetOpportunitiesGroups(params *GetOpportunitiesGroupsParams, opts ...ClientOption) (*GetOpportunitiesGroupsOK, error)

	GetOpportunitiesGroupsGroupID(params *GetOpportunitiesGroupsGroupIDParams, opts ...ClientOption) (*GetOpportunitiesGroupsGroupIDOK, error)

	GetOpportunitiesTasks(params *GetOpportunitiesTasksParams, opts ...ClientOption) (*GetOpportunitiesTasksOK, error)

	GetOpportunitiesTasksTaskID(params *GetOpportunitiesTasksTaskIDParams, opts ...ClientOption) (*GetOpportunitiesTasksTaskIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCharactersCharacterIDOpportunities gets a character s completed tasks

  Return a list of tasks finished by a character

---
Alternate route: `/dev/characters/{character_id}/opportunities/`

Alternate route: `/legacy/characters/{character_id}/opportunities/`

Alternate route: `/v1/characters/{character_id}/opportunities/`

---
This route is cached for up to 3600 seconds
*/
func (a *Client) GetCharactersCharacterIDOpportunities(params *GetCharactersCharacterIDOpportunitiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDOpportunitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDOpportunitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_opportunities",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/opportunities/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDOpportunitiesReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDOpportunitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_opportunities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOpportunitiesGroups gets opportunities groups

  Return a list of opportunities groups

---
Alternate route: `/dev/opportunities/groups/`

Alternate route: `/legacy/opportunities/groups/`

Alternate route: `/v1/opportunities/groups/`

---
This route expires daily at 11:05
*/
func (a *Client) GetOpportunitiesGroups(params *GetOpportunitiesGroupsParams, opts ...ClientOption) (*GetOpportunitiesGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOpportunitiesGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_opportunities_groups",
		Method:             "GET",
		PathPattern:        "/opportunities/groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOpportunitiesGroupsReader{formats: a.formats},
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
	success, ok := result.(*GetOpportunitiesGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_opportunities_groups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOpportunitiesGroupsGroupID gets opportunities group

  Return information of an opportunities group

---
Alternate route: `/dev/opportunities/groups/{group_id}/`

Alternate route: `/legacy/opportunities/groups/{group_id}/`

Alternate route: `/v1/opportunities/groups/{group_id}/`

---
This route expires daily at 11:05
*/
func (a *Client) GetOpportunitiesGroupsGroupID(params *GetOpportunitiesGroupsGroupIDParams, opts ...ClientOption) (*GetOpportunitiesGroupsGroupIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOpportunitiesGroupsGroupIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_opportunities_groups_group_id",
		Method:             "GET",
		PathPattern:        "/opportunities/groups/{group_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOpportunitiesGroupsGroupIDReader{formats: a.formats},
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
	success, ok := result.(*GetOpportunitiesGroupsGroupIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_opportunities_groups_group_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOpportunitiesTasks gets opportunities tasks

  Return a list of opportunities tasks

---
Alternate route: `/dev/opportunities/tasks/`

Alternate route: `/legacy/opportunities/tasks/`

Alternate route: `/v1/opportunities/tasks/`

---
This route expires daily at 11:05
*/
func (a *Client) GetOpportunitiesTasks(params *GetOpportunitiesTasksParams, opts ...ClientOption) (*GetOpportunitiesTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOpportunitiesTasksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_opportunities_tasks",
		Method:             "GET",
		PathPattern:        "/opportunities/tasks/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOpportunitiesTasksReader{formats: a.formats},
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
	success, ok := result.(*GetOpportunitiesTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_opportunities_tasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOpportunitiesTasksTaskID gets opportunities task

  Return information of an opportunities task

---
Alternate route: `/dev/opportunities/tasks/{task_id}/`

Alternate route: `/legacy/opportunities/tasks/{task_id}/`

Alternate route: `/v1/opportunities/tasks/{task_id}/`

---
This route expires daily at 11:05
*/
func (a *Client) GetOpportunitiesTasksTaskID(params *GetOpportunitiesTasksTaskIDParams, opts ...ClientOption) (*GetOpportunitiesTasksTaskIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOpportunitiesTasksTaskIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_opportunities_tasks_task_id",
		Method:             "GET",
		PathPattern:        "/opportunities/tasks/{task_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOpportunitiesTasksTaskIDReader{formats: a.formats},
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
	success, ok := result.(*GetOpportunitiesTasksTaskIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_opportunities_tasks_task_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
