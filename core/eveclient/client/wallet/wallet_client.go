// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new wallet API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for wallet API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCharactersCharacterIDWallet(params *GetCharactersCharacterIDWalletParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDWalletOK, error)

	GetCharactersCharacterIDWalletJournal(params *GetCharactersCharacterIDWalletJournalParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDWalletJournalOK, error)

	GetCharactersCharacterIDWalletTransactions(params *GetCharactersCharacterIDWalletTransactionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDWalletTransactionsOK, error)

	GetCorporationsCorporationIDWallets(params *GetCorporationsCorporationIDWalletsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDWalletsOK, error)

	GetCorporationsCorporationIDWalletsDivisionJournal(params *GetCorporationsCorporationIDWalletsDivisionJournalParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDWalletsDivisionJournalOK, error)

	GetCorporationsCorporationIDWalletsDivisionTransactions(params *GetCorporationsCorporationIDWalletsDivisionTransactionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDWalletsDivisionTransactionsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCharactersCharacterIDWallet gets a character s wallet balance

  Returns a character's wallet balance

---
Alternate route: `/legacy/characters/{character_id}/wallet/`

Alternate route: `/v1/characters/{character_id}/wallet/`

---
This route is cached for up to 120 seconds

---
[Diff of the upcoming changes](https://esi.evetech.net/diff/latest/dev/#GET-/characters/{character_id}/wallet/)
*/
func (a *Client) GetCharactersCharacterIDWallet(params *GetCharactersCharacterIDWalletParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDWalletOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDWalletParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_wallet",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/wallet/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDWalletReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDWalletOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_wallet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCharactersCharacterIDWalletJournal gets character wallet journal

  Retrieve the given character's wallet journal going 30 days back

---
Alternate route: `/dev/characters/{character_id}/wallet/journal/`

Alternate route: `/v6/characters/{character_id}/wallet/journal/`

---
This route is cached for up to 3600 seconds
*/
func (a *Client) GetCharactersCharacterIDWalletJournal(params *GetCharactersCharacterIDWalletJournalParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDWalletJournalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDWalletJournalParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_wallet_journal",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/wallet/journal/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDWalletJournalReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDWalletJournalOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_wallet_journal: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCharactersCharacterIDWalletTransactions gets wallet transactions

  Get wallet transactions of a character

---
Alternate route: `/dev/characters/{character_id}/wallet/transactions/`

Alternate route: `/legacy/characters/{character_id}/wallet/transactions/`

Alternate route: `/v1/characters/{character_id}/wallet/transactions/`

---
This route is cached for up to 3600 seconds
*/
func (a *Client) GetCharactersCharacterIDWalletTransactions(params *GetCharactersCharacterIDWalletTransactionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDWalletTransactionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDWalletTransactionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_wallet_transactions",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/wallet/transactions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDWalletTransactionsReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDWalletTransactionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_wallet_transactions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCorporationsCorporationIDWallets returns a corporation s wallet balance

  Get a corporation's wallets

---
Alternate route: `/dev/corporations/{corporation_id}/wallets/`

Alternate route: `/legacy/corporations/{corporation_id}/wallets/`

Alternate route: `/v1/corporations/{corporation_id}/wallets/`

---
This route is cached for up to 300 seconds

---
Requires one of the following EVE corporation role(s): Accountant, Junior_Accountant

*/
func (a *Client) GetCorporationsCorporationIDWallets(params *GetCorporationsCorporationIDWalletsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDWalletsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDWalletsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_wallets",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/wallets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDWalletsReader{formats: a.formats},
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
	success, ok := result.(*GetCorporationsCorporationIDWalletsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_corporations_corporation_id_wallets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCorporationsCorporationIDWalletsDivisionJournal gets corporation wallet journal

  Retrieve the given corporation's wallet journal for the given division going 30 days back

---
Alternate route: `/dev/corporations/{corporation_id}/wallets/{division}/journal/`

Alternate route: `/v4/corporations/{corporation_id}/wallets/{division}/journal/`

---
This route is cached for up to 3600 seconds

---
Requires one of the following EVE corporation role(s): Accountant, Junior_Accountant

*/
func (a *Client) GetCorporationsCorporationIDWalletsDivisionJournal(params *GetCorporationsCorporationIDWalletsDivisionJournalParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDWalletsDivisionJournalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDWalletsDivisionJournalParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_wallets_division_journal",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/wallets/{division}/journal/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDWalletsDivisionJournalReader{formats: a.formats},
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
	success, ok := result.(*GetCorporationsCorporationIDWalletsDivisionJournalOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_corporations_corporation_id_wallets_division_journal: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCorporationsCorporationIDWalletsDivisionTransactions gets corporation wallet transactions

  Get wallet transactions of a corporation

---
Alternate route: `/dev/corporations/{corporation_id}/wallets/{division}/transactions/`

Alternate route: `/legacy/corporations/{corporation_id}/wallets/{division}/transactions/`

Alternate route: `/v1/corporations/{corporation_id}/wallets/{division}/transactions/`

---
This route is cached for up to 3600 seconds

---
Requires one of the following EVE corporation role(s): Accountant, Junior_Accountant

*/
func (a *Client) GetCorporationsCorporationIDWalletsDivisionTransactions(params *GetCorporationsCorporationIDWalletsDivisionTransactionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDWalletsDivisionTransactionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDWalletsDivisionTransactionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_wallets_division_transactions",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/wallets/{division}/transactions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDWalletsDivisionTransactionsReader{formats: a.formats},
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
	success, ok := result.(*GetCorporationsCorporationIDWalletsDivisionTransactionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_corporations_corporation_id_wallets_division_transactions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}