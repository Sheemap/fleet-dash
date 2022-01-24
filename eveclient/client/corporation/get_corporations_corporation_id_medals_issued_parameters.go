// Code generated by go-swagger; DO NOT EDIT.

package corporation

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

// NewGetCorporationsCorporationIDMedalsIssuedParams creates a new GetCorporationsCorporationIDMedalsIssuedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCorporationsCorporationIDMedalsIssuedParams() *GetCorporationsCorporationIDMedalsIssuedParams {
	return &GetCorporationsCorporationIDMedalsIssuedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCorporationsCorporationIDMedalsIssuedParamsWithTimeout creates a new GetCorporationsCorporationIDMedalsIssuedParams object
// with the ability to set a timeout on a request.
func NewGetCorporationsCorporationIDMedalsIssuedParamsWithTimeout(timeout time.Duration) *GetCorporationsCorporationIDMedalsIssuedParams {
	return &GetCorporationsCorporationIDMedalsIssuedParams{
		timeout: timeout,
	}
}

// NewGetCorporationsCorporationIDMedalsIssuedParamsWithContext creates a new GetCorporationsCorporationIDMedalsIssuedParams object
// with the ability to set a context for a request.
func NewGetCorporationsCorporationIDMedalsIssuedParamsWithContext(ctx context.Context) *GetCorporationsCorporationIDMedalsIssuedParams {
	return &GetCorporationsCorporationIDMedalsIssuedParams{
		Context: ctx,
	}
}

// NewGetCorporationsCorporationIDMedalsIssuedParamsWithHTTPClient creates a new GetCorporationsCorporationIDMedalsIssuedParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCorporationsCorporationIDMedalsIssuedParamsWithHTTPClient(client *http.Client) *GetCorporationsCorporationIDMedalsIssuedParams {
	return &GetCorporationsCorporationIDMedalsIssuedParams{
		HTTPClient: client,
	}
}

/* GetCorporationsCorporationIDMedalsIssuedParams contains all the parameters to send to the API endpoint
   for the get corporations corporation id medals issued operation.

   Typically these are written to a http.Request.
*/
type GetCorporationsCorporationIDMedalsIssuedParams struct {

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

	/* Page.

	   Which page of results to return

	   Format: int32
	   Default: 1
	*/
	Page *int32

	/* Token.

	   Access token to use if unable to set a header
	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get corporations corporation id medals issued params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCorporationsCorporationIDMedalsIssuedParams) WithDefaults() *GetCorporationsCorporationIDMedalsIssuedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get corporations corporation id medals issued params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCorporationsCorporationIDMedalsIssuedParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")

		pageDefault = int32(1)
	)

	val := GetCorporationsCorporationIDMedalsIssuedParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get corporations corporation id medals issued params
func (o *GetCorporationsCorporationIDMedalsIssuedParams) WithTimeout(timeout time.Duration) *GetCorporationsCorporationIDMedalsIssuedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get corporations corporation id medals issued params
func (o *GetCorporationsCorporationIDMedalsIssuedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get corporations corporation id medals issued params
func (o *GetCorporationsCorporationIDMedalsIssuedParams) WithContext(ctx context.Context) *GetCorporationsCorporationIDMedalsIssuedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get corporations corporation id medals issued params
func (o *GetCorporationsCorporationIDMedalsIssuedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get corporations corporation id medals issued params
func (o *GetCorporationsCorporationIDMedalsIssuedParams) WithHTTPClient(client *http.Client) *GetCorporationsCorporationIDMedalsIssuedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get corporations corporation id medals issued params
func (o *GetCorporationsCorporationIDMedalsIssuedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get corporations corporation id medals issued params
func (o *GetCorporationsCorporationIDMedalsIssuedParams) WithIfNoneMatch(ifNoneMatch *string) *GetCorporationsCorporationIDMedalsIssuedParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get corporations corporation id medals issued params
func (o *GetCorporationsCorporationIDMedalsIssuedParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCorporationID adds the corporationID to the get corporations corporation id medals issued params
func (o *GetCorporationsCorporationIDMedalsIssuedParams) WithCorporationID(corporationID int32) *GetCorporationsCorporationIDMedalsIssuedParams {
	o.SetCorporationID(corporationID)
	return o
}

// SetCorporationID adds the corporationId to the get corporations corporation id medals issued params
func (o *GetCorporationsCorporationIDMedalsIssuedParams) SetCorporationID(corporationID int32) {
	o.CorporationID = corporationID
}

// WithDatasource adds the datasource to the get corporations corporation id medals issued params
func (o *GetCorporationsCorporationIDMedalsIssuedParams) WithDatasource(datasource *string) *GetCorporationsCorporationIDMedalsIssuedParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get corporations corporation id medals issued params
func (o *GetCorporationsCorporationIDMedalsIssuedParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithPage adds the page to the get corporations corporation id medals issued params
func (o *GetCorporationsCorporationIDMedalsIssuedParams) WithPage(page *int32) *GetCorporationsCorporationIDMedalsIssuedParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get corporations corporation id medals issued params
func (o *GetCorporationsCorporationIDMedalsIssuedParams) SetPage(page *int32) {
	o.Page = page
}

// WithToken adds the token to the get corporations corporation id medals issued params
func (o *GetCorporationsCorporationIDMedalsIssuedParams) WithToken(token *string) *GetCorporationsCorporationIDMedalsIssuedParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get corporations corporation id medals issued params
func (o *GetCorporationsCorporationIDMedalsIssuedParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCorporationsCorporationIDMedalsIssuedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
