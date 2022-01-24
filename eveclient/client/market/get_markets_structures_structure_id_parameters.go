// Code generated by go-swagger; DO NOT EDIT.

package market

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

// NewGetMarketsStructuresStructureIDParams creates a new GetMarketsStructuresStructureIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMarketsStructuresStructureIDParams() *GetMarketsStructuresStructureIDParams {
	return &GetMarketsStructuresStructureIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMarketsStructuresStructureIDParamsWithTimeout creates a new GetMarketsStructuresStructureIDParams object
// with the ability to set a timeout on a request.
func NewGetMarketsStructuresStructureIDParamsWithTimeout(timeout time.Duration) *GetMarketsStructuresStructureIDParams {
	return &GetMarketsStructuresStructureIDParams{
		timeout: timeout,
	}
}

// NewGetMarketsStructuresStructureIDParamsWithContext creates a new GetMarketsStructuresStructureIDParams object
// with the ability to set a context for a request.
func NewGetMarketsStructuresStructureIDParamsWithContext(ctx context.Context) *GetMarketsStructuresStructureIDParams {
	return &GetMarketsStructuresStructureIDParams{
		Context: ctx,
	}
}

// NewGetMarketsStructuresStructureIDParamsWithHTTPClient creates a new GetMarketsStructuresStructureIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMarketsStructuresStructureIDParamsWithHTTPClient(client *http.Client) *GetMarketsStructuresStructureIDParams {
	return &GetMarketsStructuresStructureIDParams{
		HTTPClient: client,
	}
}

/* GetMarketsStructuresStructureIDParams contains all the parameters to send to the API endpoint
   for the get markets structures structure id operation.

   Typically these are written to a http.Request.
*/
type GetMarketsStructuresStructureIDParams struct {

	/* IfNoneMatch.

	   ETag from a previous request. A 304 will be returned if this matches the current ETag
	*/
	IfNoneMatch *string

	/* Datasource.

	   The server name you would like data from

	   Default: "tranquility"
	*/
	Datasource *string

	/* Page.

	   Which page of results to return

	   Format: int32
	   Default: 1
	*/
	Page *int32

	/* StructureID.

	   Return orders in this structure

	   Format: int64
	*/
	StructureID int64

	/* Token.

	   Access token to use if unable to set a header
	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get markets structures structure id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMarketsStructuresStructureIDParams) WithDefaults() *GetMarketsStructuresStructureIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get markets structures structure id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMarketsStructuresStructureIDParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")

		pageDefault = int32(1)
	)

	val := GetMarketsStructuresStructureIDParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get markets structures structure id params
func (o *GetMarketsStructuresStructureIDParams) WithTimeout(timeout time.Duration) *GetMarketsStructuresStructureIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get markets structures structure id params
func (o *GetMarketsStructuresStructureIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get markets structures structure id params
func (o *GetMarketsStructuresStructureIDParams) WithContext(ctx context.Context) *GetMarketsStructuresStructureIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get markets structures structure id params
func (o *GetMarketsStructuresStructureIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get markets structures structure id params
func (o *GetMarketsStructuresStructureIDParams) WithHTTPClient(client *http.Client) *GetMarketsStructuresStructureIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get markets structures structure id params
func (o *GetMarketsStructuresStructureIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get markets structures structure id params
func (o *GetMarketsStructuresStructureIDParams) WithIfNoneMatch(ifNoneMatch *string) *GetMarketsStructuresStructureIDParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get markets structures structure id params
func (o *GetMarketsStructuresStructureIDParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get markets structures structure id params
func (o *GetMarketsStructuresStructureIDParams) WithDatasource(datasource *string) *GetMarketsStructuresStructureIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get markets structures structure id params
func (o *GetMarketsStructuresStructureIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithPage adds the page to the get markets structures structure id params
func (o *GetMarketsStructuresStructureIDParams) WithPage(page *int32) *GetMarketsStructuresStructureIDParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get markets structures structure id params
func (o *GetMarketsStructuresStructureIDParams) SetPage(page *int32) {
	o.Page = page
}

// WithStructureID adds the structureID to the get markets structures structure id params
func (o *GetMarketsStructuresStructureIDParams) WithStructureID(structureID int64) *GetMarketsStructuresStructureIDParams {
	o.SetStructureID(structureID)
	return o
}

// SetStructureID adds the structureId to the get markets structures structure id params
func (o *GetMarketsStructuresStructureIDParams) SetStructureID(structureID int64) {
	o.StructureID = structureID
}

// WithToken adds the token to the get markets structures structure id params
func (o *GetMarketsStructuresStructureIDParams) WithToken(token *string) *GetMarketsStructuresStructureIDParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get markets structures structure id params
func (o *GetMarketsStructuresStructureIDParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetMarketsStructuresStructureIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	// path param structure_id
	if err := r.SetPathParam("structure_id", swag.FormatInt64(o.StructureID)); err != nil {
		return err
	}

	if o.Token != nil {

		// query param token
		var qrToken string

		if o.Token != nil {
			qrToken = *o.Token
		}
		qToken := qrToken
		if qToken != "" {

			if err := r.SetQueryParam("token", qToken); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
