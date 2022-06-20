// Code generated by go-swagger; DO NOT EDIT.

package industry

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

// NewGetIndustrySystemsParams creates a new GetIndustrySystemsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIndustrySystemsParams() *GetIndustrySystemsParams {
	return &GetIndustrySystemsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIndustrySystemsParamsWithTimeout creates a new GetIndustrySystemsParams object
// with the ability to set a timeout on a request.
func NewGetIndustrySystemsParamsWithTimeout(timeout time.Duration) *GetIndustrySystemsParams {
	return &GetIndustrySystemsParams{
		timeout: timeout,
	}
}

// NewGetIndustrySystemsParamsWithContext creates a new GetIndustrySystemsParams object
// with the ability to set a context for a request.
func NewGetIndustrySystemsParamsWithContext(ctx context.Context) *GetIndustrySystemsParams {
	return &GetIndustrySystemsParams{
		Context: ctx,
	}
}

// NewGetIndustrySystemsParamsWithHTTPClient creates a new GetIndustrySystemsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIndustrySystemsParamsWithHTTPClient(client *http.Client) *GetIndustrySystemsParams {
	return &GetIndustrySystemsParams{
		HTTPClient: client,
	}
}

/* GetIndustrySystemsParams contains all the parameters to send to the API endpoint
   for the get industry systems operation.

   Typically these are written to a http.Request.
*/
type GetIndustrySystemsParams struct {

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

// WithDefaults hydrates default values in the get industry systems params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIndustrySystemsParams) WithDefaults() *GetIndustrySystemsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get industry systems params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIndustrySystemsParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := GetIndustrySystemsParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get industry systems params
func (o *GetIndustrySystemsParams) WithTimeout(timeout time.Duration) *GetIndustrySystemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get industry systems params
func (o *GetIndustrySystemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get industry systems params
func (o *GetIndustrySystemsParams) WithContext(ctx context.Context) *GetIndustrySystemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get industry systems params
func (o *GetIndustrySystemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get industry systems params
func (o *GetIndustrySystemsParams) WithHTTPClient(client *http.Client) *GetIndustrySystemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get industry systems params
func (o *GetIndustrySystemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get industry systems params
func (o *GetIndustrySystemsParams) WithIfNoneMatch(ifNoneMatch *string) *GetIndustrySystemsParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get industry systems params
func (o *GetIndustrySystemsParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get industry systems params
func (o *GetIndustrySystemsParams) WithDatasource(datasource *string) *GetIndustrySystemsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get industry systems params
func (o *GetIndustrySystemsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetIndustrySystemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
