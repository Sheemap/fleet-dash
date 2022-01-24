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

// GetCharactersCharacterIDOrdersReader is a Reader for the GetCharactersCharacterIDOrders structure.
type GetCharactersCharacterIDOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDOrdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDOrdersNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDOrdersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDOrdersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDOrdersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDOrdersEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDOrdersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDOrdersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDOrdersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersCharacterIDOrdersOK creates a GetCharactersCharacterIDOrdersOK with default headers values
func NewGetCharactersCharacterIDOrdersOK() *GetCharactersCharacterIDOrdersOK {
	return &GetCharactersCharacterIDOrdersOK{}
}

/* GetCharactersCharacterIDOrdersOK describes a response with status code 200, with default header values.

Open market orders placed by a character
*/
type GetCharactersCharacterIDOrdersOK struct {

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

	Payload []*GetCharactersCharacterIDOrdersOKBodyItems0
}

func (o *GetCharactersCharacterIDOrdersOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/][%d] getCharactersCharacterIdOrdersOK  %+v", 200, o.Payload)
}
func (o *GetCharactersCharacterIDOrdersOK) GetPayload() []*GetCharactersCharacterIDOrdersOKBodyItems0 {
	return o.Payload
}

func (o *GetCharactersCharacterIDOrdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersNotModified creates a GetCharactersCharacterIDOrdersNotModified with default headers values
func NewGetCharactersCharacterIDOrdersNotModified() *GetCharactersCharacterIDOrdersNotModified {
	return &GetCharactersCharacterIDOrdersNotModified{}
}

/* GetCharactersCharacterIDOrdersNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDOrdersNotModified struct {

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

func (o *GetCharactersCharacterIDOrdersNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/][%d] getCharactersCharacterIdOrdersNotModified ", 304)
}

func (o *GetCharactersCharacterIDOrdersNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDOrdersBadRequest creates a GetCharactersCharacterIDOrdersBadRequest with default headers values
func NewGetCharactersCharacterIDOrdersBadRequest() *GetCharactersCharacterIDOrdersBadRequest {
	return &GetCharactersCharacterIDOrdersBadRequest{}
}

/* GetCharactersCharacterIDOrdersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDOrdersBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDOrdersBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/][%d] getCharactersCharacterIdOrdersBadRequest  %+v", 400, o.Payload)
}
func (o *GetCharactersCharacterIDOrdersBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDOrdersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersUnauthorized creates a GetCharactersCharacterIDOrdersUnauthorized with default headers values
func NewGetCharactersCharacterIDOrdersUnauthorized() *GetCharactersCharacterIDOrdersUnauthorized {
	return &GetCharactersCharacterIDOrdersUnauthorized{}
}

/* GetCharactersCharacterIDOrdersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDOrdersUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDOrdersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/][%d] getCharactersCharacterIdOrdersUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCharactersCharacterIDOrdersUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDOrdersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersForbidden creates a GetCharactersCharacterIDOrdersForbidden with default headers values
func NewGetCharactersCharacterIDOrdersForbidden() *GetCharactersCharacterIDOrdersForbidden {
	return &GetCharactersCharacterIDOrdersForbidden{}
}

/* GetCharactersCharacterIDOrdersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDOrdersForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDOrdersForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/][%d] getCharactersCharacterIdOrdersForbidden  %+v", 403, o.Payload)
}
func (o *GetCharactersCharacterIDOrdersForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDOrdersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersEnhanceYourCalm creates a GetCharactersCharacterIDOrdersEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDOrdersEnhanceYourCalm() *GetCharactersCharacterIDOrdersEnhanceYourCalm {
	return &GetCharactersCharacterIDOrdersEnhanceYourCalm{}
}

/* GetCharactersCharacterIDOrdersEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDOrdersEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDOrdersEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/][%d] getCharactersCharacterIdOrdersEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCharactersCharacterIDOrdersEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDOrdersEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersInternalServerError creates a GetCharactersCharacterIDOrdersInternalServerError with default headers values
func NewGetCharactersCharacterIDOrdersInternalServerError() *GetCharactersCharacterIDOrdersInternalServerError {
	return &GetCharactersCharacterIDOrdersInternalServerError{}
}

/* GetCharactersCharacterIDOrdersInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDOrdersInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDOrdersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/][%d] getCharactersCharacterIdOrdersInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCharactersCharacterIDOrdersInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDOrdersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersServiceUnavailable creates a GetCharactersCharacterIDOrdersServiceUnavailable with default headers values
func NewGetCharactersCharacterIDOrdersServiceUnavailable() *GetCharactersCharacterIDOrdersServiceUnavailable {
	return &GetCharactersCharacterIDOrdersServiceUnavailable{}
}

/* GetCharactersCharacterIDOrdersServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDOrdersServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDOrdersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/][%d] getCharactersCharacterIdOrdersServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCharactersCharacterIDOrdersServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDOrdersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersGatewayTimeout creates a GetCharactersCharacterIDOrdersGatewayTimeout with default headers values
func NewGetCharactersCharacterIDOrdersGatewayTimeout() *GetCharactersCharacterIDOrdersGatewayTimeout {
	return &GetCharactersCharacterIDOrdersGatewayTimeout{}
}

/* GetCharactersCharacterIDOrdersGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDOrdersGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDOrdersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/][%d] getCharactersCharacterIdOrdersGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCharactersCharacterIDOrdersGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDOrdersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDOrdersOKBodyItems0 get_characters_character_id_orders_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDOrdersOKBodyItems0
*/
type GetCharactersCharacterIDOrdersOKBodyItems0 struct {

	// get_characters_character_id_orders_duration
	//
	// Number of days for which order is valid (starting from the issued date). An order expires at time issued + duration
	// Required: true
	Duration *int32 `json:"duration"`

	// get_characters_character_id_orders_escrow
	//
	// For buy orders, the amount of ISK in escrow
	Escrow float64 `json:"escrow,omitempty"`

	// get_characters_character_id_orders_is_buy_order
	//
	// True if the order is a bid (buy) order
	IsBuyOrder bool `json:"is_buy_order,omitempty"`

	// get_characters_character_id_orders_is_corporation
	//
	// Signifies whether the buy/sell order was placed on behalf of a corporation.
	// Required: true
	IsCorporation *bool `json:"is_corporation"`

	// get_characters_character_id_orders_issued
	//
	// Date and time when this order was issued
	// Required: true
	// Format: date-time
	Issued *strfmt.DateTime `json:"issued"`

	// get_characters_character_id_orders_location_id
	//
	// ID of the location where order was placed
	// Required: true
	LocationID *int64 `json:"location_id"`

	// get_characters_character_id_orders_min_volume
	//
	// For buy orders, the minimum quantity that will be accepted in a matching sell order
	MinVolume int32 `json:"min_volume,omitempty"`

	// get_characters_character_id_orders_order_id
	//
	// Unique order ID
	// Required: true
	OrderID *int64 `json:"order_id"`

	// get_characters_character_id_orders_price
	//
	// Cost per unit for this order
	// Required: true
	Price *float64 `json:"price"`

	// get_characters_character_id_orders_range
	//
	// Valid order range, numbers are ranges in jumps
	// Required: true
	// Enum: [1 10 2 20 3 30 4 40 5 region solarsystem station]
	Range *string `json:"range"`

	// get_characters_character_id_orders_region_id
	//
	// ID of the region where order was placed
	// Required: true
	RegionID *int32 `json:"region_id"`

	// get_characters_character_id_orders_type_id
	//
	// The type ID of the item transacted in this order
	// Required: true
	TypeID *int32 `json:"type_id"`

	// get_characters_character_id_orders_volume_remain
	//
	// Quantity of items still required or offered
	// Required: true
	VolumeRemain *int32 `json:"volume_remain"`

	// get_characters_character_id_orders_volume_total
	//
	// Quantity of items required or offered at time order was placed
	// Required: true
	VolumeTotal *int32 `json:"volume_total"`
}

// Validate validates this get characters character ID orders o k body items0
func (o *GetCharactersCharacterIDOrdersOKBodyItems0) Validate(formats strfmt.Registry) error {
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

func (o *GetCharactersCharacterIDOrdersOKBodyItems0) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", o.Duration); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDOrdersOKBodyItems0) validateIsCorporation(formats strfmt.Registry) error {

	if err := validate.Required("is_corporation", "body", o.IsCorporation); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDOrdersOKBodyItems0) validateIssued(formats strfmt.Registry) error {

	if err := validate.Required("issued", "body", o.Issued); err != nil {
		return err
	}

	if err := validate.FormatOf("issued", "body", "date-time", o.Issued.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDOrdersOKBodyItems0) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", o.LocationID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDOrdersOKBodyItems0) validateOrderID(formats strfmt.Registry) error {

	if err := validate.Required("order_id", "body", o.OrderID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDOrdersOKBodyItems0) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", o.Price); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdOrdersOKBodyItems0TypeRangePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["1","10","2","20","3","30","4","40","5","region","solarsystem","station"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdOrdersOKBodyItems0TypeRangePropEnum = append(getCharactersCharacterIdOrdersOKBodyItems0TypeRangePropEnum, v)
	}
}

const (

	// GetCharactersCharacterIDOrdersOKBodyItems0RangeNr1 captures enum value "1"
	GetCharactersCharacterIDOrdersOKBodyItems0RangeNr1 string = "1"

	// GetCharactersCharacterIDOrdersOKBodyItems0RangeNr10 captures enum value "10"
	GetCharactersCharacterIDOrdersOKBodyItems0RangeNr10 string = "10"

	// GetCharactersCharacterIDOrdersOKBodyItems0RangeNr2 captures enum value "2"
	GetCharactersCharacterIDOrdersOKBodyItems0RangeNr2 string = "2"

	// GetCharactersCharacterIDOrdersOKBodyItems0RangeNr20 captures enum value "20"
	GetCharactersCharacterIDOrdersOKBodyItems0RangeNr20 string = "20"

	// GetCharactersCharacterIDOrdersOKBodyItems0RangeNr3 captures enum value "3"
	GetCharactersCharacterIDOrdersOKBodyItems0RangeNr3 string = "3"

	// GetCharactersCharacterIDOrdersOKBodyItems0RangeNr30 captures enum value "30"
	GetCharactersCharacterIDOrdersOKBodyItems0RangeNr30 string = "30"

	// GetCharactersCharacterIDOrdersOKBodyItems0RangeNr4 captures enum value "4"
	GetCharactersCharacterIDOrdersOKBodyItems0RangeNr4 string = "4"

	// GetCharactersCharacterIDOrdersOKBodyItems0RangeNr40 captures enum value "40"
	GetCharactersCharacterIDOrdersOKBodyItems0RangeNr40 string = "40"

	// GetCharactersCharacterIDOrdersOKBodyItems0RangeNr5 captures enum value "5"
	GetCharactersCharacterIDOrdersOKBodyItems0RangeNr5 string = "5"

	// GetCharactersCharacterIDOrdersOKBodyItems0RangeRegion captures enum value "region"
	GetCharactersCharacterIDOrdersOKBodyItems0RangeRegion string = "region"

	// GetCharactersCharacterIDOrdersOKBodyItems0RangeSolarsystem captures enum value "solarsystem"
	GetCharactersCharacterIDOrdersOKBodyItems0RangeSolarsystem string = "solarsystem"

	// GetCharactersCharacterIDOrdersOKBodyItems0RangeStation captures enum value "station"
	GetCharactersCharacterIDOrdersOKBodyItems0RangeStation string = "station"
)

// prop value enum
func (o *GetCharactersCharacterIDOrdersOKBodyItems0) validateRangeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCharactersCharacterIdOrdersOKBodyItems0TypeRangePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDOrdersOKBodyItems0) validateRange(formats strfmt.Registry) error {

	if err := validate.Required("range", "body", o.Range); err != nil {
		return err
	}

	// value enum
	if err := o.validateRangeEnum("range", "body", *o.Range); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDOrdersOKBodyItems0) validateRegionID(formats strfmt.Registry) error {

	if err := validate.Required("region_id", "body", o.RegionID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDOrdersOKBodyItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDOrdersOKBodyItems0) validateVolumeRemain(formats strfmt.Registry) error {

	if err := validate.Required("volume_remain", "body", o.VolumeRemain); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDOrdersOKBodyItems0) validateVolumeTotal(formats strfmt.Registry) error {

	if err := validate.Required("volume_total", "body", o.VolumeTotal); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get characters character ID orders o k body items0 based on context it is used
func (o *GetCharactersCharacterIDOrdersOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDOrdersOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDOrdersOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDOrdersOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
