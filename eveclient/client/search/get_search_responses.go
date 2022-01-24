// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"fleet-dash-core/eveclient/models"
)

// GetSearchReader is a Reader for the GetSearch structure.
type GetSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetSearchNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetSearchEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSearchServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetSearchGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSearchOK creates a GetSearchOK with default headers values
func NewGetSearchOK() *GetSearchOK {
	return &GetSearchOK{}
}

/* GetSearchOK describes a response with status code 200, with default header values.

A list of search results
*/
type GetSearchOK struct {

	/* The caching mechanism used
	 */
	CacheControl string

	/* The language used in the response
	 */
	ContentLanguage string

	/* RFC7232 compliant entity tag
	 */
	ETag string

	/* RFC7231 formatted datetime string
	 */
	Expires string

	/* RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *GetSearchOKBody
}

func (o *GetSearchOK) Error() string {
	return fmt.Sprintf("[GET /search/][%d] getSearchOK  %+v", 200, o.Payload)
}
func (o *GetSearchOK) GetPayload() *GetSearchOKBody {
	return o.Payload
}

func (o *GetSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Cache-Control
	hdrCacheControl := response.GetHeader("Cache-Control")

	if hdrCacheControl != "" {
		o.CacheControl = hdrCacheControl
	}

	// hydrates response header Content-Language
	hdrContentLanguage := response.GetHeader("Content-Language")

	if hdrContentLanguage != "" {
		o.ContentLanguage = hdrContentLanguage
	}

	// hydrates response header ETag
	hdrETag := response.GetHeader("ETag")

	if hdrETag != "" {
		o.ETag = hdrETag
	}

	// hydrates response header Expires
	hdrExpires := response.GetHeader("Expires")

	if hdrExpires != "" {
		o.Expires = hdrExpires
	}

	// hydrates response header Last-Modified
	hdrLastModified := response.GetHeader("Last-Modified")

	if hdrLastModified != "" {
		o.LastModified = hdrLastModified
	}

	o.Payload = new(GetSearchOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchNotModified creates a GetSearchNotModified with default headers values
func NewGetSearchNotModified() *GetSearchNotModified {
	return &GetSearchNotModified{}
}

/* GetSearchNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetSearchNotModified struct {

	/* The caching mechanism used
	 */
	CacheControl string

	/* RFC7232 compliant entity tag
	 */
	ETag string

	/* RFC7231 formatted datetime string
	 */
	Expires string

	/* RFC7231 formatted datetime string
	 */
	LastModified string
}

func (o *GetSearchNotModified) Error() string {
	return fmt.Sprintf("[GET /search/][%d] getSearchNotModified ", 304)
}

func (o *GetSearchNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Cache-Control
	hdrCacheControl := response.GetHeader("Cache-Control")

	if hdrCacheControl != "" {
		o.CacheControl = hdrCacheControl
	}

	// hydrates response header ETag
	hdrETag := response.GetHeader("ETag")

	if hdrETag != "" {
		o.ETag = hdrETag
	}

	// hydrates response header Expires
	hdrExpires := response.GetHeader("Expires")

	if hdrExpires != "" {
		o.Expires = hdrExpires
	}

	// hydrates response header Last-Modified
	hdrLastModified := response.GetHeader("Last-Modified")

	if hdrLastModified != "" {
		o.LastModified = hdrLastModified
	}

	return nil
}

// NewGetSearchBadRequest creates a GetSearchBadRequest with default headers values
func NewGetSearchBadRequest() *GetSearchBadRequest {
	return &GetSearchBadRequest{}
}

/* GetSearchBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetSearchBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetSearchBadRequest) Error() string {
	return fmt.Sprintf("[GET /search/][%d] getSearchBadRequest  %+v", 400, o.Payload)
}
func (o *GetSearchBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchEnhanceYourCalm creates a GetSearchEnhanceYourCalm with default headers values
func NewGetSearchEnhanceYourCalm() *GetSearchEnhanceYourCalm {
	return &GetSearchEnhanceYourCalm{}
}

/* GetSearchEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetSearchEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetSearchEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /search/][%d] getSearchEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetSearchEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetSearchEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchInternalServerError creates a GetSearchInternalServerError with default headers values
func NewGetSearchInternalServerError() *GetSearchInternalServerError {
	return &GetSearchInternalServerError{}
}

/* GetSearchInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetSearchInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetSearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /search/][%d] getSearchInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSearchInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchServiceUnavailable creates a GetSearchServiceUnavailable with default headers values
func NewGetSearchServiceUnavailable() *GetSearchServiceUnavailable {
	return &GetSearchServiceUnavailable{}
}

/* GetSearchServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetSearchServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetSearchServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /search/][%d] getSearchServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetSearchServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetSearchServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchGatewayTimeout creates a GetSearchGatewayTimeout with default headers values
func NewGetSearchGatewayTimeout() *GetSearchGatewayTimeout {
	return &GetSearchGatewayTimeout{}
}

/* GetSearchGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetSearchGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetSearchGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /search/][%d] getSearchGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetSearchGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetSearchGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetSearchOKBody get_search_ok
//
// 200 ok object
swagger:model GetSearchOKBody
*/
type GetSearchOKBody struct {

	// get_search_agent
	//
	// agent array
	// Max Items: 500
	Agent []int32 `json:"agent"`

	// get_search_alliance
	//
	// alliance array
	// Max Items: 500
	Alliance []int32 `json:"alliance"`

	// get_search_character
	//
	// character array
	// Max Items: 500
	Character []int32 `json:"character"`

	// get_search_constellation
	//
	// constellation array
	// Max Items: 500
	Constellation []int32 `json:"constellation"`

	// get_search_corporation
	//
	// corporation array
	// Max Items: 500
	Corporation []int32 `json:"corporation"`

	// get_search_faction
	//
	// faction array
	// Max Items: 500
	Faction []int32 `json:"faction"`

	// get_search_inventory_type
	//
	// inventory_type array
	// Max Items: 500
	InventoryType []int32 `json:"inventory_type"`

	// get_search_region
	//
	// region array
	// Max Items: 500
	Region []int32 `json:"region"`

	// get_search_solar_system
	//
	// solar_system array
	// Max Items: 500
	SolarSystem []int32 `json:"solar_system"`

	// get_search_station
	//
	// station array
	// Max Items: 500
	Station []int32 `json:"station"`
}

// Validate validates this get search o k body
func (o *GetSearchOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAlliance(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCharacter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateConstellation(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCorporation(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFaction(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateInventoryType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSolarSystem(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSearchOKBody) validateAgent(formats strfmt.Registry) error {
	if swag.IsZero(o.Agent) { // not required
		return nil
	}

	iAgentSize := int64(len(o.Agent))

	if err := validate.MaxItems("getSearchOK"+"."+"agent", "body", iAgentSize, 500); err != nil {
		return err
	}

	return nil
}

func (o *GetSearchOKBody) validateAlliance(formats strfmt.Registry) error {
	if swag.IsZero(o.Alliance) { // not required
		return nil
	}

	iAllianceSize := int64(len(o.Alliance))

	if err := validate.MaxItems("getSearchOK"+"."+"alliance", "body", iAllianceSize, 500); err != nil {
		return err
	}

	return nil
}

func (o *GetSearchOKBody) validateCharacter(formats strfmt.Registry) error {
	if swag.IsZero(o.Character) { // not required
		return nil
	}

	iCharacterSize := int64(len(o.Character))

	if err := validate.MaxItems("getSearchOK"+"."+"character", "body", iCharacterSize, 500); err != nil {
		return err
	}

	return nil
}

func (o *GetSearchOKBody) validateConstellation(formats strfmt.Registry) error {
	if swag.IsZero(o.Constellation) { // not required
		return nil
	}

	iConstellationSize := int64(len(o.Constellation))

	if err := validate.MaxItems("getSearchOK"+"."+"constellation", "body", iConstellationSize, 500); err != nil {
		return err
	}

	return nil
}

func (o *GetSearchOKBody) validateCorporation(formats strfmt.Registry) error {
	if swag.IsZero(o.Corporation) { // not required
		return nil
	}

	iCorporationSize := int64(len(o.Corporation))

	if err := validate.MaxItems("getSearchOK"+"."+"corporation", "body", iCorporationSize, 500); err != nil {
		return err
	}

	return nil
}

func (o *GetSearchOKBody) validateFaction(formats strfmt.Registry) error {
	if swag.IsZero(o.Faction) { // not required
		return nil
	}

	iFactionSize := int64(len(o.Faction))

	if err := validate.MaxItems("getSearchOK"+"."+"faction", "body", iFactionSize, 500); err != nil {
		return err
	}

	return nil
}

func (o *GetSearchOKBody) validateInventoryType(formats strfmt.Registry) error {
	if swag.IsZero(o.InventoryType) { // not required
		return nil
	}

	iInventoryTypeSize := int64(len(o.InventoryType))

	if err := validate.MaxItems("getSearchOK"+"."+"inventory_type", "body", iInventoryTypeSize, 500); err != nil {
		return err
	}

	return nil
}

func (o *GetSearchOKBody) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(o.Region) { // not required
		return nil
	}

	iRegionSize := int64(len(o.Region))

	if err := validate.MaxItems("getSearchOK"+"."+"region", "body", iRegionSize, 500); err != nil {
		return err
	}

	return nil
}

func (o *GetSearchOKBody) validateSolarSystem(formats strfmt.Registry) error {
	if swag.IsZero(o.SolarSystem) { // not required
		return nil
	}

	iSolarSystemSize := int64(len(o.SolarSystem))

	if err := validate.MaxItems("getSearchOK"+"."+"solar_system", "body", iSolarSystemSize, 500); err != nil {
		return err
	}

	return nil
}

func (o *GetSearchOKBody) validateStation(formats strfmt.Registry) error {
	if swag.IsZero(o.Station) { // not required
		return nil
	}

	iStationSize := int64(len(o.Station))

	if err := validate.MaxItems("getSearchOK"+"."+"station", "body", iStationSize, 500); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get search o k body based on context it is used
func (o *GetSearchOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSearchOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSearchOKBody) UnmarshalBinary(b []byte) error {
	var res GetSearchOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
