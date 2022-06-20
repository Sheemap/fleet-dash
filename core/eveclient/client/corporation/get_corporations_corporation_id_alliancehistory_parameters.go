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

// NewGetCorporationsCorporationIDAlliancehistoryParams creates a new GetCorporationsCorporationIDAlliancehistoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCorporationsCorporationIDAlliancehistoryParams() *GetCorporationsCorporationIDAlliancehistoryParams {
	return &GetCorporationsCorporationIDAlliancehistoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCorporationsCorporationIDAlliancehistoryParamsWithTimeout creates a new GetCorporationsCorporationIDAlliancehistoryParams object
// with the ability to set a timeout on a request.
func NewGetCorporationsCorporationIDAlliancehistoryParamsWithTimeout(timeout time.Duration) *GetCorporationsCorporationIDAlliancehistoryParams {
	return &GetCorporationsCorporationIDAlliancehistoryParams{
		timeout: timeout,
	}
}

// NewGetCorporationsCorporationIDAlliancehistoryParamsWithContext creates a new GetCorporationsCorporationIDAlliancehistoryParams object
// with the ability to set a context for a request.
func NewGetCorporationsCorporationIDAlliancehistoryParamsWithContext(ctx context.Context) *GetCorporationsCorporationIDAlliancehistoryParams {
	return &GetCorporationsCorporationIDAlliancehistoryParams{
		Context: ctx,
	}
}

// NewGetCorporationsCorporationIDAlliancehistoryParamsWithHTTPClient creates a new GetCorporationsCorporationIDAlliancehistoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCorporationsCorporationIDAlliancehistoryParamsWithHTTPClient(client *http.Client) *GetCorporationsCorporationIDAlliancehistoryParams {
	return &GetCorporationsCorporationIDAlliancehistoryParams{
		HTTPClient: client,
	}
}

/* GetCorporationsCorporationIDAlliancehistoryParams contains all the parameters to send to the API endpoint
   for the get corporations corporation id alliancehistory operation.

   Typically these are written to a http.Request.
*/
type GetCorporationsCorporationIDAlliancehistoryParams struct {

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

// WithDefaults hydrates default values in the get corporations corporation id alliancehistory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCorporationsCorporationIDAlliancehistoryParams) WithDefaults() *GetCorporationsCorporationIDAlliancehistoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get corporations corporation id alliancehistory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCorporationsCorporationIDAlliancehistoryParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := GetCorporationsCorporationIDAlliancehistoryParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get corporations corporation id alliancehistory params
func (o *GetCorporationsCorporationIDAlliancehistoryParams) WithTimeout(timeout time.Duration) *GetCorporationsCorporationIDAlliancehistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get corporations corporation id alliancehistory params
func (o *GetCorporationsCorporationIDAlliancehistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get corporations corporation id alliancehistory params
func (o *GetCorporationsCorporationIDAlliancehistoryParams) WithContext(ctx context.Context) *GetCorporationsCorporationIDAlliancehistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get corporations corporation id alliancehistory params
func (o *GetCorporationsCorporationIDAlliancehistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get corporations corporation id alliancehistory params
func (o *GetCorporationsCorporationIDAlliancehistoryParams) WithHTTPClient(client *http.Client) *GetCorporationsCorporationIDAlliancehistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get corporations corporation id alliancehistory params
func (o *GetCorporationsCorporationIDAlliancehistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get corporations corporation id alliancehistory params
func (o *GetCorporationsCorporationIDAlliancehistoryParams) WithIfNoneMatch(ifNoneMatch *string) *GetCorporationsCorporationIDAlliancehistoryParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get corporations corporation id alliancehistory params
func (o *GetCorporationsCorporationIDAlliancehistoryParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCorporationID adds the corporationID to the get corporations corporation id alliancehistory params
func (o *GetCorporationsCorporationIDAlliancehistoryParams) WithCorporationID(corporationID int32) *GetCorporationsCorporationIDAlliancehistoryParams {
	o.SetCorporationID(corporationID)
	return o
}

// SetCorporationID adds the corporationId to the get corporations corporation id alliancehistory params
func (o *GetCorporationsCorporationIDAlliancehistoryParams) SetCorporationID(corporationID int32) {
	o.CorporationID = corporationID
}

// WithDatasource adds the datasource to the get corporations corporation id alliancehistory params
func (o *GetCorporationsCorporationIDAlliancehistoryParams) WithDatasource(datasource *string) *GetCorporationsCorporationIDAlliancehistoryParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get corporations corporation id alliancehistory params
func (o *GetCorporationsCorporationIDAlliancehistoryParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetCorporationsCorporationIDAlliancehistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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