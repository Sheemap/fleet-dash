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

// GetMarketsStructuresStructureIDReader is a Reader for the GetMarketsStructuresStructureID structure.
type GetMarketsStructuresStructureIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMarketsStructuresStructureIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMarketsStructuresStructureIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetMarketsStructuresStructureIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetMarketsStructuresStructureIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetMarketsStructuresStructureIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMarketsStructuresStructureIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetMarketsStructuresStructureIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMarketsStructuresStructureIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetMarketsStructuresStructureIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetMarketsStructuresStructureIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMarketsStructuresStructureIDOK creates a GetMarketsStructuresStructureIDOK with default headers values
func NewGetMarketsStructuresStructureIDOK() *GetMarketsStructuresStructureIDOK {
	var (
		// initialize headers with default values
		xPagesDefault = int32(1)
	)

	return &GetMarketsStructuresStructureIDOK{

		XPages: xPagesDefault,
	}
}

/* GetMarketsStructuresStructureIDOK describes a response with status code 200, with default header values.

A list of orders
*/
type GetMarketsStructuresStructureIDOK struct {

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

	Payload []*GetMarketsStructuresStructureIDOKBodyItems0
}

func (o *GetMarketsStructuresStructureIDOK) Error() string {
	return fmt.Sprintf("[GET /markets/structures/{structure_id}/][%d] getMarketsStructuresStructureIdOK  %+v", 200, o.Payload)
}
func (o *GetMarketsStructuresStructureIDOK) GetPayload() []*GetMarketsStructuresStructureIDOKBodyItems0 {
	return o.Payload
}

func (o *GetMarketsStructuresStructureIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMarketsStructuresStructureIDNotModified creates a GetMarketsStructuresStructureIDNotModified with default headers values
func NewGetMarketsStructuresStructureIDNotModified() *GetMarketsStructuresStructureIDNotModified {
	return &GetMarketsStructuresStructureIDNotModified{}
}

/* GetMarketsStructuresStructureIDNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetMarketsStructuresStructureIDNotModified struct {

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

func (o *GetMarketsStructuresStructureIDNotModified) Error() string {
	return fmt.Sprintf("[GET /markets/structures/{structure_id}/][%d] getMarketsStructuresStructureIdNotModified ", 304)
}

func (o *GetMarketsStructuresStructureIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMarketsStructuresStructureIDBadRequest creates a GetMarketsStructuresStructureIDBadRequest with default headers values
func NewGetMarketsStructuresStructureIDBadRequest() *GetMarketsStructuresStructureIDBadRequest {
	return &GetMarketsStructuresStructureIDBadRequest{}
}

/* GetMarketsStructuresStructureIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetMarketsStructuresStructureIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetMarketsStructuresStructureIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /markets/structures/{structure_id}/][%d] getMarketsStructuresStructureIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetMarketsStructuresStructureIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetMarketsStructuresStructureIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsStructuresStructureIDUnauthorized creates a GetMarketsStructuresStructureIDUnauthorized with default headers values
func NewGetMarketsStructuresStructureIDUnauthorized() *GetMarketsStructuresStructureIDUnauthorized {
	return &GetMarketsStructuresStructureIDUnauthorized{}
}

/* GetMarketsStructuresStructureIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetMarketsStructuresStructureIDUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetMarketsStructuresStructureIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /markets/structures/{structure_id}/][%d] getMarketsStructuresStructureIdUnauthorized  %+v", 401, o.Payload)
}
func (o *GetMarketsStructuresStructureIDUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetMarketsStructuresStructureIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsStructuresStructureIDForbidden creates a GetMarketsStructuresStructureIDForbidden with default headers values
func NewGetMarketsStructuresStructureIDForbidden() *GetMarketsStructuresStructureIDForbidden {
	return &GetMarketsStructuresStructureIDForbidden{}
}

/* GetMarketsStructuresStructureIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetMarketsStructuresStructureIDForbidden struct {
	Payload *models.Forbidden
}

func (o *GetMarketsStructuresStructureIDForbidden) Error() string {
	return fmt.Sprintf("[GET /markets/structures/{structure_id}/][%d] getMarketsStructuresStructureIdForbidden  %+v", 403, o.Payload)
}
func (o *GetMarketsStructuresStructureIDForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetMarketsStructuresStructureIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsStructuresStructureIDEnhanceYourCalm creates a GetMarketsStructuresStructureIDEnhanceYourCalm with default headers values
func NewGetMarketsStructuresStructureIDEnhanceYourCalm() *GetMarketsStructuresStructureIDEnhanceYourCalm {
	return &GetMarketsStructuresStructureIDEnhanceYourCalm{}
}

/* GetMarketsStructuresStructureIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetMarketsStructuresStructureIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetMarketsStructuresStructureIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /markets/structures/{structure_id}/][%d] getMarketsStructuresStructureIdEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetMarketsStructuresStructureIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetMarketsStructuresStructureIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsStructuresStructureIDInternalServerError creates a GetMarketsStructuresStructureIDInternalServerError with default headers values
func NewGetMarketsStructuresStructureIDInternalServerError() *GetMarketsStructuresStructureIDInternalServerError {
	return &GetMarketsStructuresStructureIDInternalServerError{}
}

/* GetMarketsStructuresStructureIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetMarketsStructuresStructureIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetMarketsStructuresStructureIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /markets/structures/{structure_id}/][%d] getMarketsStructuresStructureIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetMarketsStructuresStructureIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetMarketsStructuresStructureIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsStructuresStructureIDServiceUnavailable creates a GetMarketsStructuresStructureIDServiceUnavailable with default headers values
func NewGetMarketsStructuresStructureIDServiceUnavailable() *GetMarketsStructuresStructureIDServiceUnavailable {
	return &GetMarketsStructuresStructureIDServiceUnavailable{}
}

/* GetMarketsStructuresStructureIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetMarketsStructuresStructureIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetMarketsStructuresStructureIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /markets/structures/{structure_id}/][%d] getMarketsStructuresStructureIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetMarketsStructuresStructureIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetMarketsStructuresStructureIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsStructuresStructureIDGatewayTimeout creates a GetMarketsStructuresStructureIDGatewayTimeout with default headers values
func NewGetMarketsStructuresStructureIDGatewayTimeout() *GetMarketsStructuresStructureIDGatewayTimeout {
	return &GetMarketsStructuresStructureIDGatewayTimeout{}
}

/* GetMarketsStructuresStructureIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetMarketsStructuresStructureIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetMarketsStructuresStructureIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /markets/structures/{structure_id}/][%d] getMarketsStructuresStructureIdGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetMarketsStructuresStructureIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetMarketsStructuresStructureIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetMarketsStructuresStructureIDOKBodyItems0 get_markets_structures_structure_id_200_ok
//
// 200 ok object
swagger:model GetMarketsStructuresStructureIDOKBodyItems0
*/
type GetMarketsStructuresStructureIDOKBodyItems0 struct {

	// get_markets_structures_structure_id_duration
	//
	// duration integer
	// Required: true
	Duration *int32 `json:"duration"`

	// get_markets_structures_structure_id_is_buy_order
	//
	// is_buy_order boolean
	// Required: true
	IsBuyOrder *bool `json:"is_buy_order"`

	// get_markets_structures_structure_id_issued
	//
	// issued string
	// Required: true
	// Format: date-time
	Issued *strfmt.DateTime `json:"issued"`

	// get_markets_structures_structure_id_location_id
	//
	// location_id integer
	// Required: true
	LocationID *int64 `json:"location_id"`

	// get_markets_structures_structure_id_min_volume
	//
	// min_volume integer
	// Required: true
	MinVolume *int32 `json:"min_volume"`

	// get_markets_structures_structure_id_order_id
	//
	// order_id integer
	// Required: true
	OrderID *int64 `json:"order_id"`

	// get_markets_structures_structure_id_price
	//
	// price number
	// Required: true
	Price *float64 `json:"price"`

	// get_markets_structures_structure_id_range
	//
	// range string
	// Required: true
	// Enum: [station region solarsystem 1 2 3 4 5 10 20 30 40]
	Range *string `json:"range"`

	// get_markets_structures_structure_id_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`

	// get_markets_structures_structure_id_volume_remain
	//
	// volume_remain integer
	// Required: true
	VolumeRemain *int32 `json:"volume_remain"`

	// get_markets_structures_structure_id_volume_total
	//
	// volume_total integer
	// Required: true
	VolumeTotal *int32 `json:"volume_total"`
}

// Validate validates this get markets structures structure ID o k body items0
func (o *GetMarketsStructuresStructureIDOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIsBuyOrder(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIssued(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMinVolume(formats); err != nil {
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

func (o *GetMarketsStructuresStructureIDOKBodyItems0) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", o.Duration); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsStructuresStructureIDOKBodyItems0) validateIsBuyOrder(formats strfmt.Registry) error {

	if err := validate.Required("is_buy_order", "body", o.IsBuyOrder); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsStructuresStructureIDOKBodyItems0) validateIssued(formats strfmt.Registry) error {

	if err := validate.Required("issued", "body", o.Issued); err != nil {
		return err
	}

	if err := validate.FormatOf("issued", "body", "date-time", o.Issued.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsStructuresStructureIDOKBodyItems0) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", o.LocationID); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsStructuresStructureIDOKBodyItems0) validateMinVolume(formats strfmt.Registry) error {

	if err := validate.Required("min_volume", "body", o.MinVolume); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsStructuresStructureIDOKBodyItems0) validateOrderID(formats strfmt.Registry) error {

	if err := validate.Required("order_id", "body", o.OrderID); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsStructuresStructureIDOKBodyItems0) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", o.Price); err != nil {
		return err
	}

	return nil
}

var getMarketsStructuresStructureIdOKBodyItems0TypeRangePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["station","region","solarsystem","1","2","3","4","5","10","20","30","40"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getMarketsStructuresStructureIdOKBodyItems0TypeRangePropEnum = append(getMarketsStructuresStructureIdOKBodyItems0TypeRangePropEnum, v)
	}
}

const (

	// GetMarketsStructuresStructureIDOKBodyItems0RangeStation captures enum value "station"
	GetMarketsStructuresStructureIDOKBodyItems0RangeStation string = "station"

	// GetMarketsStructuresStructureIDOKBodyItems0RangeRegion captures enum value "region"
	GetMarketsStructuresStructureIDOKBodyItems0RangeRegion string = "region"

	// GetMarketsStructuresStructureIDOKBodyItems0RangeSolarsystem captures enum value "solarsystem"
	GetMarketsStructuresStructureIDOKBodyItems0RangeSolarsystem string = "solarsystem"

	// GetMarketsStructuresStructureIDOKBodyItems0RangeNr1 captures enum value "1"
	GetMarketsStructuresStructureIDOKBodyItems0RangeNr1 string = "1"

	// GetMarketsStructuresStructureIDOKBodyItems0RangeNr2 captures enum value "2"
	GetMarketsStructuresStructureIDOKBodyItems0RangeNr2 string = "2"

	// GetMarketsStructuresStructureIDOKBodyItems0RangeNr3 captures enum value "3"
	GetMarketsStructuresStructureIDOKBodyItems0RangeNr3 string = "3"

	// GetMarketsStructuresStructureIDOKBodyItems0RangeNr4 captures enum value "4"
	GetMarketsStructuresStructureIDOKBodyItems0RangeNr4 string = "4"

	// GetMarketsStructuresStructureIDOKBodyItems0RangeNr5 captures enum value "5"
	GetMarketsStructuresStructureIDOKBodyItems0RangeNr5 string = "5"

	// GetMarketsStructuresStructureIDOKBodyItems0RangeNr10 captures enum value "10"
	GetMarketsStructuresStructureIDOKBodyItems0RangeNr10 string = "10"

	// GetMarketsStructuresStructureIDOKBodyItems0RangeNr20 captures enum value "20"
	GetMarketsStructuresStructureIDOKBodyItems0RangeNr20 string = "20"

	// GetMarketsStructuresStructureIDOKBodyItems0RangeNr30 captures enum value "30"
	GetMarketsStructuresStructureIDOKBodyItems0RangeNr30 string = "30"

	// GetMarketsStructuresStructureIDOKBodyItems0RangeNr40 captures enum value "40"
	GetMarketsStructuresStructureIDOKBodyItems0RangeNr40 string = "40"
)

// prop value enum
func (o *GetMarketsStructuresStructureIDOKBodyItems0) validateRangeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getMarketsStructuresStructureIdOKBodyItems0TypeRangePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetMarketsStructuresStructureIDOKBodyItems0) validateRange(formats strfmt.Registry) error {

	if err := validate.Required("range", "body", o.Range); err != nil {
		return err
	}

	// value enum
	if err := o.validateRangeEnum("range", "body", *o.Range); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsStructuresStructureIDOKBodyItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsStructuresStructureIDOKBodyItems0) validateVolumeRemain(formats strfmt.Registry) error {

	if err := validate.Required("volume_remain", "body", o.VolumeRemain); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsStructuresStructureIDOKBodyItems0) validateVolumeTotal(formats strfmt.Registry) error {

	if err := validate.Required("volume_total", "body", o.VolumeTotal); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get markets structures structure ID o k body items0 based on context it is used
func (o *GetMarketsStructuresStructureIDOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetMarketsStructuresStructureIDOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMarketsStructuresStructureIDOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetMarketsStructuresStructureIDOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
