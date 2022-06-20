// Code generated by go-swagger; DO NOT EDIT.

package market

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"fleet-dash-core/eveclient/models"
)

// GetCharactersCharacterIDOrdersHistoryReader is a Reader for the GetCharactersCharacterIDOrdersHistory structure.
type GetCharactersCharacterIDOrdersHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDOrdersHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDOrdersHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDOrdersHistoryNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDOrdersHistoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDOrdersHistoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDOrdersHistoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDOrdersHistoryEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDOrdersHistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDOrdersHistoryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDOrdersHistoryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersCharacterIDOrdersHistoryOK creates a GetCharactersCharacterIDOrdersHistoryOK with default headers values
func NewGetCharactersCharacterIDOrdersHistoryOK() *GetCharactersCharacterIDOrdersHistoryOK {
	var (
		// initialize headers with default values
		xPagesDefault = int32(1)
	)

	return &GetCharactersCharacterIDOrdersHistoryOK{

		XPages: xPagesDefault,
	}
}

/* GetCharactersCharacterIDOrdersHistoryOK describes a response with status code 200, with default header values.

Expired and cancelled market orders placed by a character
*/
type GetCharactersCharacterIDOrdersHistoryOK struct {

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

	/* Maximum page number

	   Format: int32
	   Default: 1
	*/
	XPages int32

	Payload []*GetCharactersCharacterIDOrdersHistoryOKBodyItems0
}

func (o *GetCharactersCharacterIDOrdersHistoryOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/history/][%d] getCharactersCharacterIdOrdersHistoryOK  %+v", 200, o.Payload)
}
func (o *GetCharactersCharacterIDOrdersHistoryOK) GetPayload() []*GetCharactersCharacterIDOrdersHistoryOKBodyItems0 {
	return o.Payload
}

func (o *GetCharactersCharacterIDOrdersHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	// hydrates response header X-Pages
	hdrXPages := response.GetHeader("X-Pages")

	if hdrXPages != "" {
		valxPages, err := swag.ConvertInt32(hdrXPages)
		if err != nil {
			return errors.InvalidType("X-Pages", "header", "int32", hdrXPages)
		}
		o.XPages = valxPages
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersHistoryNotModified creates a GetCharactersCharacterIDOrdersHistoryNotModified with default headers values
func NewGetCharactersCharacterIDOrdersHistoryNotModified() *GetCharactersCharacterIDOrdersHistoryNotModified {
	return &GetCharactersCharacterIDOrdersHistoryNotModified{}
}

/* GetCharactersCharacterIDOrdersHistoryNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDOrdersHistoryNotModified struct {

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

func (o *GetCharactersCharacterIDOrdersHistoryNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/history/][%d] getCharactersCharacterIdOrdersHistoryNotModified ", 304)
}

func (o *GetCharactersCharacterIDOrdersHistoryNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDOrdersHistoryBadRequest creates a GetCharactersCharacterIDOrdersHistoryBadRequest with default headers values
func NewGetCharactersCharacterIDOrdersHistoryBadRequest() *GetCharactersCharacterIDOrdersHistoryBadRequest {
	return &GetCharactersCharacterIDOrdersHistoryBadRequest{}
}

/* GetCharactersCharacterIDOrdersHistoryBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDOrdersHistoryBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDOrdersHistoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/history/][%d] getCharactersCharacterIdOrdersHistoryBadRequest  %+v", 400, o.Payload)
}
func (o *GetCharactersCharacterIDOrdersHistoryBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDOrdersHistoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersHistoryUnauthorized creates a GetCharactersCharacterIDOrdersHistoryUnauthorized with default headers values
func NewGetCharactersCharacterIDOrdersHistoryUnauthorized() *GetCharactersCharacterIDOrdersHistoryUnauthorized {
	return &GetCharactersCharacterIDOrdersHistoryUnauthorized{}
}

/* GetCharactersCharacterIDOrdersHistoryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDOrdersHistoryUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDOrdersHistoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/history/][%d] getCharactersCharacterIdOrdersHistoryUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCharactersCharacterIDOrdersHistoryUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDOrdersHistoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersHistoryForbidden creates a GetCharactersCharacterIDOrdersHistoryForbidden with default headers values
func NewGetCharactersCharacterIDOrdersHistoryForbidden() *GetCharactersCharacterIDOrdersHistoryForbidden {
	return &GetCharactersCharacterIDOrdersHistoryForbidden{}
}

/* GetCharactersCharacterIDOrdersHistoryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDOrdersHistoryForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDOrdersHistoryForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/history/][%d] getCharactersCharacterIdOrdersHistoryForbidden  %+v", 403, o.Payload)
}
func (o *GetCharactersCharacterIDOrdersHistoryForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDOrdersHistoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersHistoryEnhanceYourCalm creates a GetCharactersCharacterIDOrdersHistoryEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDOrdersHistoryEnhanceYourCalm() *GetCharactersCharacterIDOrdersHistoryEnhanceYourCalm {
	return &GetCharactersCharacterIDOrdersHistoryEnhanceYourCalm{}
}

/* GetCharactersCharacterIDOrdersHistoryEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDOrdersHistoryEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDOrdersHistoryEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/history/][%d] getCharactersCharacterIdOrdersHistoryEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCharactersCharacterIDOrdersHistoryEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDOrdersHistoryEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersHistoryInternalServerError creates a GetCharactersCharacterIDOrdersHistoryInternalServerError with default headers values
func NewGetCharactersCharacterIDOrdersHistoryInternalServerError() *GetCharactersCharacterIDOrdersHistoryInternalServerError {
	return &GetCharactersCharacterIDOrdersHistoryInternalServerError{}
}

/* GetCharactersCharacterIDOrdersHistoryInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDOrdersHistoryInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDOrdersHistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/history/][%d] getCharactersCharacterIdOrdersHistoryInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCharactersCharacterIDOrdersHistoryInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDOrdersHistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersHistoryServiceUnavailable creates a GetCharactersCharacterIDOrdersHistoryServiceUnavailable with default headers values
func NewGetCharactersCharacterIDOrdersHistoryServiceUnavailable() *GetCharactersCharacterIDOrdersHistoryServiceUnavailable {
	return &GetCharactersCharacterIDOrdersHistoryServiceUnavailable{}
}

/* GetCharactersCharacterIDOrdersHistoryServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDOrdersHistoryServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDOrdersHistoryServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/history/][%d] getCharactersCharacterIdOrdersHistoryServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCharactersCharacterIDOrdersHistoryServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDOrdersHistoryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersHistoryGatewayTimeout creates a GetCharactersCharacterIDOrdersHistoryGatewayTimeout with default headers values
func NewGetCharactersCharacterIDOrdersHistoryGatewayTimeout() *GetCharactersCharacterIDOrdersHistoryGatewayTimeout {
	return &GetCharactersCharacterIDOrdersHistoryGatewayTimeout{}
}

/* GetCharactersCharacterIDOrdersHistoryGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDOrdersHistoryGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDOrdersHistoryGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/history/][%d] getCharactersCharacterIdOrdersHistoryGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCharactersCharacterIDOrdersHistoryGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDOrdersHistoryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDOrdersHistoryOKBodyItems0 get_characters_character_id_orders_history_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDOrdersHistoryOKBodyItems0
*/
type GetCharactersCharacterIDOrdersHistoryOKBodyItems0 struct {

	// get_characters_character_id_orders_history_duration
	//
	// Number of days the order was valid for (starting from the issued date). An order expires at time issued + duration
	// Required: true
	Duration *int32 `json:"duration"`

	// get_characters_character_id_orders_history_escrow
	//
	// For buy orders, the amount of ISK in escrow
	Escrow float64 `json:"escrow,omitempty"`

	// get_characters_character_id_orders_history_is_buy_order
	//
	// True if the order is a bid (buy) order
	IsBuyOrder bool `json:"is_buy_order,omitempty"`

	// get_characters_character_id_orders_history_is_corporation
	//
	// Signifies whether the buy/sell order was placed on behalf of a corporation.
	// Required: true
	IsCorporation *bool `json:"is_corporation"`

	// get_characters_character_id_orders_history_issued
	//
	// Date and time when this order was issued
	// Required: true
	// Format: date-time
	Issued *strfmt.DateTime `json:"issued"`

	// get_characters_character_id_orders_history_location_id
	//
	// ID of the location where order was placed
	// Required: true
	LocationID *int64 `json:"location_id"`

	// get_characters_character_id_orders_history_min_volume
	//
	// For buy orders, the minimum quantity that will be accepted in a matching sell order
	MinVolume int32 `json:"min_volume,omitempty"`

	// get_characters_character_id_orders_history_order_id
	//
	// Unique order ID
	// Required: true
	OrderID *int64 `json:"order_id"`

	// get_characters_character_id_orders_history_price
	//
	// Cost per unit for this order
	// Required: true
	Price *float64 `json:"price"`

	// get_characters_character_id_orders_history_range
	//
	// Valid order range, numbers are ranges in jumps
	// Required: true
	// Enum: [1 10 2 20 3 30 4 40 5 region solarsystem station]
	Range *string `json:"range"`

	// get_characters_character_id_orders_history_region_id
	//
	// ID of the region where order was placed
	// Required: true
	RegionID *int32 `json:"region_id"`

	// get_characters_character_id_orders_history_state
	//
	// Current order state
	// Required: true
	// Enum: [cancelled expired]
	State *string `json:"state"`

	// get_characters_character_id_orders_history_type_id
	//
	// The type ID of the item transacted in this order
	// Required: true
	TypeID *int32 `json:"type_id"`

	// get_characters_character_id_orders_history_volume_remain
	//
	// Quantity of items still required or offered
	// Required: true
	VolumeRemain *int32 `json:"volume_remain"`

	// get_characters_character_id_orders_history_volume_total
	//
	// Quantity of items required or offered at time order was placed
	// Required: true
	VolumeTotal *int32 `json:"volume_total"`
}

// Validate validates this get characters character ID orders history o k body items0
func (o *GetCharactersCharacterIDOrdersHistoryOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIsCorporation(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIssued(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRange(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRegionID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTypeID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVolumeRemain(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVolumeTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDOrdersHistoryOKBodyItems0) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", o.Duration); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDOrdersHistoryOKBodyItems0) validateIsCorporation(formats strfmt.Registry) error {

	if err := validate.Required("is_corporation", "body", o.IsCorporation); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDOrdersHistoryOKBodyItems0) validateIssued(formats strfmt.Registry) error {

	if err := validate.Required("issued", "body", o.Issued); err != nil {
		return err
	}

	if err := validate.FormatOf("issued", "body", "date-time", o.Issued.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDOrdersHistoryOKBodyItems0) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", o.LocationID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDOrdersHistoryOKBodyItems0) validateOrderID(formats strfmt.Registry) error {

	if err := validate.Required("order_id", "body", o.OrderID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDOrdersHistoryOKBodyItems0) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", o.Price); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdOrdersHistoryOKBodyItems0TypeRangePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["1","10","2","20","3","30","4","40","5","region","solarsystem","station"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdOrdersHistoryOKBodyItems0TypeRangePropEnum = append(getCharactersCharacterIdOrdersHistoryOKBodyItems0TypeRangePropEnum, v)
	}
}

const (

	// GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeNr1 captures enum value "1"
	GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeNr1 string = "1"

	// GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeNr10 captures enum value "10"
	GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeNr10 string = "10"

	// GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeNr2 captures enum value "2"
	GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeNr2 string = "2"

	// GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeNr20 captures enum value "20"
	GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeNr20 string = "20"

	// GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeNr3 captures enum value "3"
	GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeNr3 string = "3"

	// GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeNr30 captures enum value "30"
	GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeNr30 string = "30"

	// GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeNr4 captures enum value "4"
	GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeNr4 string = "4"

	// GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeNr40 captures enum value "40"
	GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeNr40 string = "40"

	// GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeNr5 captures enum value "5"
	GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeNr5 string = "5"

	// GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeRegion captures enum value "region"
	GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeRegion string = "region"

	// GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeSolarsystem captures enum value "solarsystem"
	GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeSolarsystem string = "solarsystem"

	// GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeStation captures enum value "station"
	GetCharactersCharacterIDOrdersHistoryOKBodyItems0RangeStation string = "station"
)

// prop value enum
func (o *GetCharactersCharacterIDOrdersHistoryOKBodyItems0) validateRangeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCharactersCharacterIdOrdersHistoryOKBodyItems0TypeRangePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDOrdersHistoryOKBodyItems0) validateRange(formats strfmt.Registry) error {

	if err := validate.Required("range", "body", o.Range); err != nil {
		return err
	}

	// value enum
	if err := o.validateRangeEnum("range", "body", *o.Range); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDOrdersHistoryOKBodyItems0) validateRegionID(formats strfmt.Registry) error {

	if err := validate.Required("region_id", "body", o.RegionID); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdOrdersHistoryOKBodyItems0TypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cancelled","expired"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdOrdersHistoryOKBodyItems0TypeStatePropEnum = append(getCharactersCharacterIdOrdersHistoryOKBodyItems0TypeStatePropEnum, v)
	}
}

const (

	// GetCharactersCharacterIDOrdersHistoryOKBodyItems0StateCancelled captures enum value "cancelled"
	GetCharactersCharacterIDOrdersHistoryOKBodyItems0StateCancelled string = "cancelled"

	// GetCharactersCharacterIDOrdersHistoryOKBodyItems0StateExpired captures enum value "expired"
	GetCharactersCharacterIDOrdersHistoryOKBodyItems0StateExpired string = "expired"
)

// prop value enum
func (o *GetCharactersCharacterIDOrdersHistoryOKBodyItems0) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCharactersCharacterIdOrdersHistoryOKBodyItems0TypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDOrdersHistoryOKBodyItems0) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", o.State); err != nil {
		return err
	}

	// value enum
	if err := o.validateStateEnum("state", "body", *o.State); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDOrdersHistoryOKBodyItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDOrdersHistoryOKBodyItems0) validateVolumeRemain(formats strfmt.Registry) error {

	if err := validate.Required("volume_remain", "body", o.VolumeRemain); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDOrdersHistoryOKBodyItems0) validateVolumeTotal(formats strfmt.Registry) error {

	if err := validate.Required("volume_total", "body", o.VolumeTotal); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get characters character ID orders history o k body items0 based on context it is used
func (o *GetCharactersCharacterIDOrdersHistoryOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDOrdersHistoryOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDOrdersHistoryOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDOrdersHistoryOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
