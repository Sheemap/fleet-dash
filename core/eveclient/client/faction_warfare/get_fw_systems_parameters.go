// Code generated by go-swagger; DO NOT EDIT.

package faction_warfare

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
)

// NewGetFwSystemsParams creates a new GetFwSystemsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFwSystemsParams() *GetFwSystemsParams {
	return &GetFwSystemsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFwSystemsParamsWithTimeout creates a new GetFwSystemsParams object
// with the ability to set a timeout on a request.
func NewGetFwSystemsParamsWithTimeout(timeout time.Duration) *GetFwSystemsParams {
	return &GetFwSystemsParams{
		timeout: timeout,
	}
}

// NewGetFwSystemsParamsWithContext creates a new GetFwSystemsParams object
// with the ability to set a context for a request.
func NewGetFwSystemsParamsWithContext(ctx context.Context) *GetFwSystemsParams {
	return &GetFwSystemsParams{
		Context: ctx,
	}
}

// NewGetFwSystemsParamsWithHTTPClient creates a new GetFwSystemsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFwSystemsParamsWithHTTPClient(client *http.Client) *GetFwSystemsParams {
	return &GetFwSystemsParams{
		HTTPClient: client,
	}
}

/* GetFwSystemsParams contains all the parameters to send to the API endpoint
   for the get fw systems operation.

   Typically these are written to a http.Request.
*/
type GetFwSystemsParams struct {

	/* IfNoneMatch.

	   ETag from a previous request. A 304 will be returned if this matches the current ETag
	*/
	IfNoneMatch *string

	/* Datasource.

	   The server name you would like data from

	   Default: "tranquility"
	*/
	Datasource *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get fw systems params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFwSystemsParams) WithDefaults() *GetFwSystemsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get fw systems params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFwSystemsParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := GetFwSystemsParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get fw systems params
func (o *GetFwSystemsParams) WithTimeout(timeout time.Duration) *GetFwSystemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fw systems params
func (o *GetFwSystemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fw systems params
func (o *GetFwSystemsParams) WithContext(ctx context.Context) *GetFwSystemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fw systems params
func (o *GetFwSystemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fw systems params
func (o *GetFwSystemsParams) WithHTTPClient(client *http.Client) *GetFwSystemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fw systems params
func (o *GetFwSystemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get fw systems params
func (o *GetFwSystemsParams) WithIfNoneMatch(ifNoneMatch *string) *GetFwSystemsParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get fw systems params
func (o *GetFwSystemsParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get fw systems params
func (o *GetFwSystemsParams) WithDatasource(datasource *string) *GetFwSystemsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get fw systems params
func (o *GetFwSystemsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetFwSystemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}