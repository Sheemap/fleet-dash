// Code generated by go-swagger; DO NOT EDIT.

package insurance

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

// NewGetInsurancePricesParams creates a new GetInsurancePricesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInsurancePricesParams() *GetInsurancePricesParams {
	return &GetInsurancePricesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInsurancePricesParamsWithTimeout creates a new GetInsurancePricesParams object
// with the ability to set a timeout on a request.
func NewGetInsurancePricesParamsWithTimeout(timeout time.Duration) *GetInsurancePricesParams {
	return &GetInsurancePricesParams{
		timeout: timeout,
	}
}

// NewGetInsurancePricesParamsWithContext creates a new GetInsurancePricesParams object
// with the ability to set a context for a request.
func NewGetInsurancePricesParamsWithContext(ctx context.Context) *GetInsurancePricesParams {
	return &GetInsurancePricesParams{
		Context: ctx,
	}
}

// NewGetInsurancePricesParamsWithHTTPClient creates a new GetInsurancePricesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInsurancePricesParamsWithHTTPClient(client *http.Client) *GetInsurancePricesParams {
	return &GetInsurancePricesParams{
		HTTPClient: client,
	}
}

/* GetInsurancePricesParams contains all the parameters to send to the API endpoint
   for the get insurance prices operation.

   Typically these are written to a http.Request.
*/
type GetInsurancePricesParams struct {

	/* AcceptLanguage.

	   Language to use in the response

	   Default: "en"
	*/
	AcceptLanguage *string

	/* IfNoneMatch.

	   ETag from a previous request. A 304 will be returned if this matches the current ETag
	*/
	IfNoneMatch *string

	/* Datasource.

	   The server name you would like data from

	   Default: "tranquility"
	*/
	Datasource *string

	/* Language.

	   Language to use in the response, takes precedence over Accept-Language

	   Default: "en"
	*/
	Language *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get insurance prices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInsurancePricesParams) WithDefaults() *GetInsurancePricesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get insurance prices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInsurancePricesParams) SetDefaults() {
	var (
		acceptLanguageDefault = string("en")

		datasourceDefault = string("tranquility")

		languageDefault = string("en")
	)

	val := GetInsurancePricesParams{
		AcceptLanguage: &acceptLanguageDefault,
		Datasource:     &datasourceDefault,
		Language:       &languageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get insurance prices params
func (o *GetInsurancePricesParams) WithTimeout(timeout time.Duration) *GetInsurancePricesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get insurance prices params
func (o *GetInsurancePricesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get insurance prices params
func (o *GetInsurancePricesParams) WithContext(ctx context.Context) *GetInsurancePricesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get insurance prices params
func (o *GetInsurancePricesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get insurance prices params
func (o *GetInsurancePricesParams) WithHTTPClient(client *http.Client) *GetInsurancePricesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get insurance prices params
func (o *GetInsurancePricesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get insurance prices params
func (o *GetInsurancePricesParams) WithAcceptLanguage(acceptLanguage *string) *GetInsurancePricesParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get insurance prices params
func (o *GetInsurancePricesParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithIfNoneMatch adds the ifNoneMatch to the get insurance prices params
func (o *GetInsurancePricesParams) WithIfNoneMatch(ifNoneMatch *string) *GetInsurancePricesParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get insurance prices params
func (o *GetInsurancePricesParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get insurance prices params
func (o *GetInsurancePricesParams) WithDatasource(datasource *string) *GetInsurancePricesParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get insurance prices params
func (o *GetInsurancePricesParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithLanguage adds the language to the get insurance prices params
func (o *GetInsurancePricesParams) WithLanguage(language *string) *GetInsurancePricesParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the get insurance prices params
func (o *GetInsurancePricesParams) SetLanguage(language *string) {
	o.Language = language
}

// WriteToRequest writes these params to a swagger request
func (o *GetInsurancePricesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AcceptLanguage != nil {

		// header param Accept-Language
		if err := r.SetHeaderParam("Accept-Language", *o.AcceptLanguage); err != nil {
			return err
		}
	}

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

	if o.Language != nil {

		// query param language
		var qrLanguage string

		if o.Language != nil {
			qrLanguage = *o.Language
		}
		qLanguage := qrLanguage
		if qLanguage != "" {

			if err := r.SetQueryParam("language", qLanguage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}