// Code generated by go-swagger; DO NOT EDIT.

package assets

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

// NewPostCorporationsCorporationIDAssetsNamesParams creates a new PostCorporationsCorporationIDAssetsNamesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostCorporationsCorporationIDAssetsNamesParams() *PostCorporationsCorporationIDAssetsNamesParams {
	return &PostCorporationsCorporationIDAssetsNamesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostCorporationsCorporationIDAssetsNamesParamsWithTimeout creates a new PostCorporationsCorporationIDAssetsNamesParams object
// with the ability to set a timeout on a request.
func NewPostCorporationsCorporationIDAssetsNamesParamsWithTimeout(timeout time.Duration) *PostCorporationsCorporationIDAssetsNamesParams {
	return &PostCorporationsCorporationIDAssetsNamesParams{
		timeout: timeout,
	}
}

// NewPostCorporationsCorporationIDAssetsNamesParamsWithContext creates a new PostCorporationsCorporationIDAssetsNamesParams object
// with the ability to set a context for a request.
func NewPostCorporationsCorporationIDAssetsNamesParamsWithContext(ctx context.Context) *PostCorporationsCorporationIDAssetsNamesParams {
	return &PostCorporationsCorporationIDAssetsNamesParams{
		Context: ctx,
	}
}

// NewPostCorporationsCorporationIDAssetsNamesParamsWithHTTPClient creates a new PostCorporationsCorporationIDAssetsNamesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostCorporationsCorporationIDAssetsNamesParamsWithHTTPClient(client *http.Client) *PostCorporationsCorporationIDAssetsNamesParams {
	return &PostCorporationsCorporationIDAssetsNamesParams{
		HTTPClient: client,
	}
}

/* PostCorporationsCorporationIDAssetsNamesParams contains all the parameters to send to the API endpoint
   for the post corporations corporation id assets names operation.

   Typically these are written to a http.Request.
*/
type PostCorporationsCorporationIDAssetsNamesParams struct {

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

	/* ItemIds.

	   A list of item ids
	*/
	ItemIds []int64

	/* Token.

	   Access token to use if unable to set a header
	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post corporations corporation id assets names params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCorporationsCorporationIDAssetsNamesParams) WithDefaults() *PostCorporationsCorporationIDAssetsNamesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post corporations corporation id assets names params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCorporationsCorporationIDAssetsNamesParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := PostCorporationsCorporationIDAssetsNamesParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the post corporations corporation id assets names params
func (o *PostCorporationsCorporationIDAssetsNamesParams) WithTimeout(timeout time.Duration) *PostCorporationsCorporationIDAssetsNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post corporations corporation id assets names params
func (o *PostCorporationsCorporationIDAssetsNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post corporations corporation id assets names params
func (o *PostCorporationsCorporationIDAssetsNamesParams) WithContext(ctx context.Context) *PostCorporationsCorporationIDAssetsNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post corporations corporation id assets names params
func (o *PostCorporationsCorporationIDAssetsNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post corporations corporation id assets names params
func (o *PostCorporationsCorporationIDAssetsNamesParams) WithHTTPClient(client *http.Client) *PostCorporationsCorporationIDAssetsNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post corporations corporation id assets names params
func (o *PostCorporationsCorporationIDAssetsNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCorporationID adds the corporationID to the post corporations corporation id assets names params
func (o *PostCorporationsCorporationIDAssetsNamesParams) WithCorporationID(corporationID int32) *PostCorporationsCorporationIDAssetsNamesParams {
	o.SetCorporationID(corporationID)
	return o
}

// SetCorporationID adds the corporationId to the post corporations corporation id assets names params
func (o *PostCorporationsCorporationIDAssetsNamesParams) SetCorporationID(corporationID int32) {
	o.CorporationID = corporationID
}

// WithDatasource adds the datasource to the post corporations corporation id assets names params
func (o *PostCorporationsCorporationIDAssetsNamesParams) WithDatasource(datasource *string) *PostCorporationsCorporationIDAssetsNamesParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the post corporations corporation id assets names params
func (o *PostCorporationsCorporationIDAssetsNamesParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithItemIds adds the itemIds to the post corporations corporation id assets names params
func (o *PostCorporationsCorporationIDAssetsNamesParams) WithItemIds(itemIds []int64) *PostCorporationsCorporationIDAssetsNamesParams {
	o.SetItemIds(itemIds)
	return o
}

// SetItemIds adds the itemIds to the post corporations corporation id assets names params
func (o *PostCorporationsCorporationIDAssetsNamesParams) SetItemIds(itemIds []int64) {
	o.ItemIds = itemIds
}

// WithToken adds the token to the post corporations corporation id assets names params
func (o *PostCorporationsCorporationIDAssetsNamesParams) WithToken(token *string) *PostCorporationsCorporationIDAssetsNamesParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the post corporations corporation id assets names params
func (o *PostCorporationsCorporationIDAssetsNamesParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *PostCorporationsCorporationIDAssetsNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
	if o.ItemIds != nil {
		if err := r.SetBodyParam(o.ItemIds); err != nil {
			return err
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