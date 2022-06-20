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

// GetCorporationsCorporationIDOrdersHistoryReader is a Reader for the GetCorporationsCorporationIDOrdersHistory structure.
type GetCorporationsCorporationIDOrdersHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDOrdersHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationsCorporationIDOrdersHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationsCorporationIDOrdersHistoryNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationsCorporationIDOrdersHistoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationsCorporationIDOrdersHistoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationsCorporationIDOrdersHistoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationsCorporationIDOrdersHistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationsCorporationIDOrdersHistoryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationsCorporationIDOrdersHistoryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDOrdersHistoryOK creates a GetCorporationsCorporationIDOrdersHistoryOK with default headers values
func NewGetCorporationsCorporationIDOrdersHistoryOK() *GetCorporationsCorporationIDOrdersHistoryOK {
	var (
		// initialize headers with default values
		xPagesDefault = int32(1)
	)

	return &GetCorporationsCorporationIDOrdersHistoryOK{

		XPages: xPagesDefault,
	}
}

/* GetCorporationsCorporationIDOrdersHistoryOK describes a response with status code 200, with default header values.

Expired and cancelled market orders placed on behalf of a corporation
*/
type GetCorporationsCorporationIDOrdersHistoryOK struct {

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

	Payload []*GetCorporationsCorporationIDOrdersHistoryOKBodyItems0
}

func (o *GetCorporationsCorporationIDOrdersHistoryOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/history/][%d] getCorporationsCorporationIdOrdersHistoryOK  %+v", 200, o.Payload)
}
func (o *GetCorporationsCorporationIDOrdersHistoryOK) GetPayload() []*GetCorporationsCorporationIDOrdersHistoryOKBodyItems0 {
	return o.Payload
}

func (o *GetCorporationsCorporationIDOrdersHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDOrdersHistoryNotModified creates a GetCorporationsCorporationIDOrdersHistoryNotModified with default headers values
func NewGetCorporationsCorporationIDOrdersHistoryNotModified() *GetCorporationsCorporationIDOrdersHistoryNotModified {
	return &GetCorporationsCorporationIDOrdersHistoryNotModified{}
}

/* GetCorporationsCorporationIDOrdersHistoryNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationsCorporationIDOrdersHistoryNotModified struct {

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

func (o *GetCorporationsCorporationIDOrdersHistoryNotModified) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/history/][%d] getCorporationsCorporationIdOrdersHistoryNotModified ", 304)
}

func (o *GetCorporationsCorporationIDOrdersHistoryNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDOrdersHistoryBadRequest creates a GetCorporationsCorporationIDOrdersHistoryBadRequest with default headers values
func NewGetCorporationsCorporationIDOrdersHistoryBadRequest() *GetCorporationsCorporationIDOrdersHistoryBadRequest {
	return &GetCorporationsCorporationIDOrdersHistoryBadRequest{}
}

/* GetCorporationsCorporationIDOrdersHistoryBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationsCorporationIDOrdersHistoryBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDOrdersHistoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/history/][%d] getCorporationsCorporationIdOrdersHistoryBadRequest  %+v", 400, o.Payload)
}
func (o *GetCorporationsCorporationIDOrdersHistoryBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationsCorporationIDOrdersHistoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDOrdersHistoryUnauthorized creates a GetCorporationsCorporationIDOrdersHistoryUnauthorized with default headers values
func NewGetCorporationsCorporationIDOrdersHistoryUnauthorized() *GetCorporationsCorporationIDOrdersHistoryUnauthorized {
	return &GetCorporationsCorporationIDOrdersHistoryUnauthorized{}
}

/* GetCorporationsCorporationIDOrdersHistoryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDOrdersHistoryUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDOrdersHistoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/history/][%d] getCorporationsCorporationIdOrdersHistoryUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCorporationsCorporationIDOrdersHistoryUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationsCorporationIDOrdersHistoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDOrdersHistoryForbidden creates a GetCorporationsCorporationIDOrdersHistoryForbidden with default headers values
func NewGetCorporationsCorporationIDOrdersHistoryForbidden() *GetCorporationsCorporationIDOrdersHistoryForbidden {
	return &GetCorporationsCorporationIDOrdersHistoryForbidden{}
}

/* GetCorporationsCorporationIDOrdersHistoryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationsCorporationIDOrdersHistoryForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDOrdersHistoryForbidden) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/history/][%d] getCorporationsCorporationIdOrdersHistoryForbidden  %+v", 403, o.Payload)
}
func (o *GetCorporationsCorporationIDOrdersHistoryForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationsCorporationIDOrdersHistoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm creates a GetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm() *GetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm {
	return &GetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm{}
}

/* GetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/history/][%d] getCorporationsCorporationIdOrdersHistoryEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDOrdersHistoryInternalServerError creates a GetCorporationsCorporationIDOrdersHistoryInternalServerError with default headers values
func NewGetCorporationsCorporationIDOrdersHistoryInternalServerError() *GetCorporationsCorporationIDOrdersHistoryInternalServerError {
	return &GetCorporationsCorporationIDOrdersHistoryInternalServerError{}
}

/* GetCorporationsCorporationIDOrdersHistoryInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationsCorporationIDOrdersHistoryInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDOrdersHistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/history/][%d] getCorporationsCorporationIdOrdersHistoryInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCorporationsCorporationIDOrdersHistoryInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationsCorporationIDOrdersHistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDOrdersHistoryServiceUnavailable creates a GetCorporationsCorporationIDOrdersHistoryServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDOrdersHistoryServiceUnavailable() *GetCorporationsCorporationIDOrdersHistoryServiceUnavailable {
	return &GetCorporationsCorporationIDOrdersHistoryServiceUnavailable{}
}

/* GetCorporationsCorporationIDOrdersHistoryServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDOrdersHistoryServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDOrdersHistoryServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/history/][%d] getCorporationsCorporationIdOrdersHistoryServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCorporationsCorporationIDOrdersHistoryServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationsCorporationIDOrdersHistoryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDOrdersHistoryGatewayTimeout creates a GetCorporationsCorporationIDOrdersHistoryGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDOrdersHistoryGatewayTimeout() *GetCorporationsCorporationIDOrdersHistoryGatewayTimeout {
	return &GetCorporationsCorporationIDOrdersHistoryGatewayTimeout{}
}

/* GetCorporationsCorporationIDOrdersHistoryGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDOrdersHistoryGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDOrdersHistoryGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/history/][%d] getCorporationsCorporationIdOrdersHistoryGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCorporationsCorporationIDOrdersHistoryGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationsCorporationIDOrdersHistoryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDOrdersHistoryOKBodyItems0 get_corporations_corporation_id_orders_history_200_ok
//
// 200 ok object
swagger:model GetCorporationsCorporationIDOrdersHistoryOKBodyItems0
*/
type GetCorporationsCorporationIDOrdersHistoryOKBodyItems0 struct {

	// get_corporations_corporation_id_orders_history_duration
	//
	// Number of days the order was valid for (starting from the issued date). An order expires at time issued + duration
	// Required: true
	Duration *int32 `json:"duration"`

	// get_corporations_corporation_id_orders_history_escrow
	//
	// For buy orders, the amount of ISK in escrow
	Escrow float64 `json:"escrow,omitempty"`

	// get_corporations_corporation_id_orders_history_is_buy_order
	//
	// True if the order is a bid (buy) order
	IsBuyOrder bool `json:"is_buy_order,omitempty"`

	// get_corporations_corporation_id_orders_history_issued
	//
	// Date and time when this order was issued
	// Required: true
	// Format: date-time
	Issued *strfmt.DateTime `json:"issued"`

	// get_corporations_corporation_id_orders_history_issued_by
	//
	// The character who issued this order
	IssuedBy int32 `json:"issued_by,omitempty"`

	// get_corporations_corporation_id_orders_history_location_id
	//
	// ID of the location where order was placed
	// Required: true
	LocationID *int64 `json:"location_id"`

	// get_corporations_corporation_id_orders_history_min_volume
	//
	// For buy orders, the minimum quantity that will be accepted in a matching sell order
	MinVolume int32 `json:"min_volume,omitempty"`

	// get_corporations_corporation_id_orders_history_order_id
	//
	// Unique order ID
	// Required: true
	OrderID *int64 `json:"order_id"`

	// get_corporations_corporation_id_orders_history_price
	//
	// Cost per unit for this order
	// Required: true
	Price *float64 `json:"price"`

	// get_corporations_corporation_id_orders_history_range
	//
	// Valid order range, numbers are ranges in jumps
	// Required: true
	// Enum: [1 10 2 20 3 30 4 40 5 region solarsystem station]
	Range *string `json:"range"`

	// get_corporations_corporation_id_orders_history_region_id
	//
	// ID of the region where order was placed
	// Required: true
	RegionID *int32 `json:"region_id"`

	// get_corporations_corporation_id_orders_history_state
	//
	// Current order state
	// Required: true
	// Enum: [cancelled expired]
	State *string `json:"state"`

	// get_corporations_corporation_id_orders_history_type_id
	//
	// The type ID of the item transacted in this order
	// Required: true
	TypeID *int32 `json:"type_id"`

	// get_corporations_corporation_id_orders_history_volume_remain
	//
	// Quantity of items still required or offered
	// Required: true
	VolumeRemain *int32 `json:"volume_remain"`

	// get_corporations_corporation_id_orders_history_volume_total
	//
	// Quantity of items required or offered at time order was placed
	// Required: true
	VolumeTotal *int32 `json:"volume_total"`

	// get_corporations_corporation_id_orders_history_wallet_division
	//
	// The corporation wallet division used for this order
	// Required: true
	// Maximum: 7
	// Minimum: 1
	WalletDivision *int32 `json:"wallet_division"`
}

// Validate validates this get corporations corporation ID orders history o k body items0
func (o *GetCorporationsCorporationIDOrdersHistoryOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDuration(formats); err != nil {
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

	if err := o.validateWalletDivision(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationsCorporationIDOrdersHistoryOKBodyItems0) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", o.Duration); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDOrdersHistoryOKBodyItems0) validateIssued(formats strfmt.Registry) error {

	if err := validate.Required("issued", "body", o.Issued); err != nil {
		return err
	}

	if err := validate.FormatOf("issued", "body", "date-time", o.Issued.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDOrdersHistoryOKBodyItems0) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", o.LocationID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDOrdersHistoryOKBodyItems0) validateOrderID(formats strfmt.Registry) error {

	if err := validate.Required("order_id", "body", o.OrderID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDOrdersHistoryOKBodyItems0) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", o.Price); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdOrdersHistoryOKBodyItems0TypeRangePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["1","10","2","20","3","30","4","40","5","region","solarsystem","station"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdOrdersHistoryOKBodyItems0TypeRangePropEnum = append(getCorporationsCorporationIdOrdersHistoryOKBodyItems0TypeRangePropEnum, v)
	}
}

const (

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeNr1 captures enum value "1"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeNr1 string = "1"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeNr10 captures enum value "10"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeNr10 string = "10"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeNr2 captures enum value "2"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeNr2 string = "2"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeNr20 captures enum value "20"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeNr20 string = "20"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeNr3 captures enum value "3"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeNr3 string = "3"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeNr30 captures enum value "30"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeNr30 string = "30"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeNr4 captures enum value "4"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeNr4 string = "4"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeNr40 captures enum value "40"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeNr40 string = "40"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeNr5 captures enum value "5"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeNr5 string = "5"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeRegion captures enum value "region"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeRegion string = "region"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeSolarsystem captures enum value "solarsystem"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeSolarsystem string = "solarsystem"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeStation captures enum value "station"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItems0RangeStation string = "station"
)

// prop value enum
func (o *GetCorporationsCorporationIDOrdersHistoryOKBodyItems0) validateRangeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCorporationsCorporationIdOrdersHistoryOKBodyItems0TypeRangePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDOrdersHistoryOKBodyItems0) validateRange(formats strfmt.Registry) error {

	if err := validate.Required("range", "body", o.Range); err != nil {
		return err
	}

	// value enum
	if err := o.validateRangeEnum("range", "body", *o.Range); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDOrdersHistoryOKBodyItems0) validateRegionID(formats strfmt.Registry) error {

	if err := validate.Required("region_id", "body", o.RegionID); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdOrdersHistoryOKBodyItems0TypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cancelled","expired"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdOrdersHistoryOKBodyItems0TypeStatePropEnum = append(getCorporationsCorporationIdOrdersHistoryOKBodyItems0TypeStatePropEnum, v)
	}
}

const (

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItems0StateCancelled captures enum value "cancelled"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItems0StateCancelled string = "cancelled"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItems0StateExpired captures enum value "expired"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItems0StateExpired string = "expired"
)

// prop value enum
func (o *GetCorporationsCorporationIDOrdersHistoryOKBodyItems0) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCorporationsCorporationIdOrdersHistoryOKBodyItems0TypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDOrdersHistoryOKBodyItems0) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", o.State); err != nil {
		return err
	}

	// value enum
	if err := o.validateStateEnum("state", "body", *o.State); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDOrdersHistoryOKBodyItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDOrdersHistoryOKBodyItems0) validateVolumeRemain(formats strfmt.Registry) error {

	if err := validate.Required("volume_remain", "body", o.VolumeRemain); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDOrdersHistoryOKBodyItems0) validateVolumeTotal(formats strfmt.Registry) error {

	if err := validate.Required("volume_total", "body", o.VolumeTotal); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDOrdersHistoryOKBodyItems0) validateWalletDivision(formats strfmt.Registry) error {

	if err := validate.Required("wallet_division", "body", o.WalletDivision); err != nil {
		return err
	}

	if err := validate.MinimumInt("wallet_division", "body", int64(*o.WalletDivision), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("wallet_division", "body", int64(*o.WalletDivision), 7, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get corporations corporation ID orders history o k body items0 based on context it is used
func (o *GetCorporationsCorporationIDOrdersHistoryOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDOrdersHistoryOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDOrdersHistoryOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDOrdersHistoryOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
