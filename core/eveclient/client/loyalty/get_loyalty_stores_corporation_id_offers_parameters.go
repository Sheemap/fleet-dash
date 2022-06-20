// Code generated by go-swagger; DO NOT EDIT.

package loyalty

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetLoyaltyStoresCorporationIDOffersParams creates a new GetLoyaltyStoresCorporationIDOffersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLoyaltyStoresCorporationIDOffersParams() *GetLoyaltyStoresCorporationIDOffersParams {
	return &GetLoyaltyStoresCorporationIDOffersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoyaltyStoresCorporationIDOffersParamsWithTimeout creates a new GetLoyaltyStoresCorporationIDOffersParams object
// with the ability to set a timeout on a request.
func NewGetLoyaltyStoresCorporationIDOffersParamsWithTimeout(timeout time.Duration) *GetLoyaltyStoresCorporationIDOffersParams {
	return &GetLoyaltyStoresCorporationIDOffersParams{
		timeout: timeout,
	}
}

// NewGetLoyaltyStoresCorporationIDOffersParamsWithContext creates a new GetLoyaltyStoresCorporationIDOffersParams object
// with the ability to set a context for a request.
func NewGetLoyaltyStoresCorporationIDOffersParamsWithContext(ctx context.Context) *GetLoyaltyStoresCorporationIDOffersParams {
	return &GetLoyaltyStoresCorporationIDOffersParams{
		Context: ctx,
	}
}

// NewGetLoyaltyStoresCorporationIDOffersParamsWithHTTPClient creates a new GetLoyaltyStoresCorporationIDOffersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLoyaltyStoresCorporationIDOffersParamsWithHTTPClient(client *http.Client) *GetLoyaltyStoresCorporationIDOffersParams {
	return &GetLoyaltyStoresCorporationIDOffersParams{
		HTTPClient: client,
	}
}

/* GetLoyaltyStoresCorporationIDOffersParams contains all the parameters to send to the API endpoint
   for the get loyalty stores corporation id offers operation.

   Typically these are written to a http.Request.
*/
type GetLoyaltyStoresCorporationIDOffersParams struct {

	/* IfNoneMatch.

	   ETag from a previous request. A 304 will be returned if this matches the current ETag
	*/
	IfNoneMatch *string

	/* CorporationID.

	   An EVE corporation ID

	   Format: int32
	*/
	CorporationID int32

	/* Datasource.

	   The server name you would like data from

	   Default: "tranquility"
	*/
	Datasource *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get loyalty stores corporation id offers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoyaltyStoresCorporationIDOffersParams) WithDefaults() *GetLoyaltyStoresCorporationIDOffersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get loyalty stores corporation id offers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoyaltyStoresCorporationIDOffersParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := GetLoyaltyStoresCorporationIDOffersParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get loyalty stores corporation id offers params
func (o *GetLoyaltyStoresCorporationIDOffersParams) WithTimeout(timeout time.Duration) *GetLoyaltyStoresCorporationIDOffersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get loyalty stores corporation id offers params
func (o *GetLoyaltyStoresCorporationIDOffersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get loyalty stores corporation id offers params
func (o *GetLoyaltyStoresCorporationIDOffersParams) WithContext(ctx context.Context) *GetLoyaltyStoresCorporationIDOffersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get loyalty stores corporation id offers params
func (o *GetLoyaltyStoresCorporationIDOffersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get loyalty stores corporation id offers params
func (o *GetLoyaltyStoresCorporationIDOffersParams) WithHTTPClient(client *http.Client) *GetLoyaltyStoresCorporationIDOffersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get loyalty stores corporation id offers params
func (o *GetLoyaltyStoresCorporationIDOffersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get loyalty stores corporation id offers params
func (o *GetLoyaltyStoresCorporationIDOffersParams) WithIfNoneMatch(ifNoneMatch *string) *GetLoyaltyStoresCorporationIDOffersParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get loyalty stores corporation id offers params
func (o *GetLoyaltyStoresCorporationIDOffersParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCorporationID adds the corporationID to the get loyalty stores corporation id offers params
func (o *GetLoyaltyStoresCorporationIDOffersParams) WithCorporationID(corporationID int32) *GetLoyaltyStoresCorporationIDOffersParams {
	o.SetCorporationID(corporationID)
	return o
}

// SetCorporationID adds the corporationId to the get loyalty stores corporation id offers params
func (o *GetLoyaltyStoresCorporationIDOffersParams) SetCorporationID(corporationID int32) {
	o.CorporationID = corporationID
}

// WithDatasource adds the datasource to the get loyalty stores corporation id offers params
func (o *GetLoyaltyStoresCorporationIDOffersParams) WithDatasource(datasource *string) *GetLoyaltyStoresCorporationIDOffersParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get loyalty stores corporation id offers params
func (o *GetLoyaltyStoresCorporationIDOffersParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoyaltyStoresCorporationIDOffersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IfNoneMatch != nil {

		// header param If-None-Match
		if err := r.SetHeaderParam("If-None-Match", *o.IfNoneMatch); err != nil {
			return err
		}
	}

	// path param corporation_id
	if err := r.SetPathParam("corporation_id", swag.FormatInt32(o.CorporationID)); err != nil {
		return err
	}

	if o.Datasource != nil {

		// query param datasource
		var qrDatasource string

		if o.Datasource != nil {
			qrDatasource = *o.Datasource
		}
		qDatasource := qrDatasource
		if qDatasource != "" {

			if err := r.SetQueryParam("datasource", qDatasource); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}