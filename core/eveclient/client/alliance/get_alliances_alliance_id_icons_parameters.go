// Code generated by go-swagger; DO NOT EDIT.

package alliance

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

// NewGetAlliancesAllianceIDIconsParams creates a new GetAlliancesAllianceIDIconsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAlliancesAllianceIDIconsParams() *GetAlliancesAllianceIDIconsParams {
	return &GetAlliancesAllianceIDIconsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlliancesAllianceIDIconsParamsWithTimeout creates a new GetAlliancesAllianceIDIconsParams object
// with the ability to set a timeout on a request.
func NewGetAlliancesAllianceIDIconsParamsWithTimeout(timeout time.Duration) *GetAlliancesAllianceIDIconsParams {
	return &GetAlliancesAllianceIDIconsParams{
		timeout: timeout,
	}
}

// NewGetAlliancesAllianceIDIconsParamsWithContext creates a new GetAlliancesAllianceIDIconsParams object
// with the ability to set a context for a request.
func NewGetAlliancesAllianceIDIconsParamsWithContext(ctx context.Context) *GetAlliancesAllianceIDIconsParams {
	return &GetAlliancesAllianceIDIconsParams{
		Context: ctx,
	}
}

// NewGetAlliancesAllianceIDIconsParamsWithHTTPClient creates a new GetAlliancesAllianceIDIconsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAlliancesAllianceIDIconsParamsWithHTTPClient(client *http.Client) *GetAlliancesAllianceIDIconsParams {
	return &GetAlliancesAllianceIDIconsParams{
		HTTPClient: client,
	}
}

/* GetAlliancesAllianceIDIconsParams contains all the parameters to send to the API endpoint
   for the get alliances alliance id icons operation.

   Typically these are written to a http.Request.
*/
type GetAlliancesAllianceIDIconsParams struct {

	/* IfNoneMatch.

	   ETag from a previous request. A 304 will be returned if this matches the current ETag
	*/
	IfNoneMatch *string

	/* AllianceID.

	   An EVE alliance ID

	   Format: int32
	*/
	AllianceID int32

	/* Datasource.

	   The server name you would like data from

	   Default: "tranquility"
	*/
	Datasource *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get alliances alliance id icons params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlliancesAllianceIDIconsParams) WithDefaults() *GetAlliancesAllianceIDIconsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get alliances alliance id icons params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlliancesAllianceIDIconsParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := GetAlliancesAllianceIDIconsParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get alliances alliance id icons params
func (o *GetAlliancesAllianceIDIconsParams) WithTimeout(timeout time.Duration) *GetAlliancesAllianceIDIconsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alliances alliance id icons params
func (o *GetAlliancesAllianceIDIconsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alliances alliance id icons params
func (o *GetAlliancesAllianceIDIconsParams) WithContext(ctx context.Context) *GetAlliancesAllianceIDIconsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alliances alliance id icons params
func (o *GetAlliancesAllianceIDIconsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alliances alliance id icons params
func (o *GetAlliancesAllianceIDIconsParams) WithHTTPClient(client *http.Client) *GetAlliancesAllianceIDIconsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alliances alliance id icons params
func (o *GetAlliancesAllianceIDIconsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get alliances alliance id icons params
func (o *GetAlliancesAllianceIDIconsParams) WithIfNoneMatch(ifNoneMatch *string) *GetAlliancesAllianceIDIconsParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get alliances alliance id icons params
func (o *GetAlliancesAllianceIDIconsParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithAllianceID adds the allianceID to the get alliances alliance id icons params
func (o *GetAlliancesAllianceIDIconsParams) WithAllianceID(allianceID int32) *GetAlliancesAllianceIDIconsParams {
	o.SetAllianceID(allianceID)
	return o
}

// SetAllianceID adds the allianceId to the get alliances alliance id icons params
func (o *GetAlliancesAllianceIDIconsParams) SetAllianceID(allianceID int32) {
	o.AllianceID = allianceID
}

// WithDatasource adds the datasource to the get alliances alliance id icons params
func (o *GetAlliancesAllianceIDIconsParams) WithDatasource(datasource *string) *GetAlliancesAllianceIDIconsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get alliances alliance id icons params
func (o *GetAlliancesAllianceIDIconsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlliancesAllianceIDIconsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param alliance_id
	if err := r.SetPathParam("alliance_id", swag.FormatInt32(o.AllianceID)); err != nil {
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