// Code generated by go-swagger; DO NOT EDIT.

package fittings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"fleet-dash-core/eveclient/models"
)

// GetCharactersCharacterIDFittingsReader is a Reader for the GetCharactersCharacterIDFittings structure.
type GetCharactersCharacterIDFittingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDFittingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDFittingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDFittingsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDFittingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDFittingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDFittingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDFittingsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDFittingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDFittingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDFittingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersCharacterIDFittingsOK creates a GetCharactersCharacterIDFittingsOK with default headers values
func NewGetCharactersCharacterIDFittingsOK() *GetCharactersCharacterIDFittingsOK {
	return &GetCharactersCharacterIDFittingsOK{}
}

/* GetCharactersCharacterIDFittingsOK describes a response with status code 200, with default header values.

A list of fittings
*/
type GetCharactersCharacterIDFittingsOK struct {

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

	Payload []*GetCharactersCharacterIDFittingsOKBodyItems0
}

func (o *GetCharactersCharacterIDFittingsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fittings/][%d] getCharactersCharacterIdFittingsOK  %+v", 200, o.Payload)
}
func (o *GetCharactersCharacterIDFittingsOK) GetPayload() []*GetCharactersCharacterIDFittingsOKBodyItems0 {
	return o.Payload
}

func (o *GetCharactersCharacterIDFittingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDFittingsNotModified creates a GetCharactersCharacterIDFittingsNotModified with default headers values
func NewGetCharactersCharacterIDFittingsNotModified() *GetCharactersCharacterIDFittingsNotModified {
	return &GetCharactersCharacterIDFittingsNotModified{}
}

/* GetCharactersCharacterIDFittingsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDFittingsNotModified struct {

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

func (o *GetCharactersCharacterIDFittingsNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fittings/][%d] getCharactersCharacterIdFittingsNotModified ", 304)
}

func (o *GetCharactersCharacterIDFittingsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDFittingsBadRequest creates a GetCharactersCharacterIDFittingsBadRequest with default headers values
func NewGetCharactersCharacterIDFittingsBadRequest() *GetCharactersCharacterIDFittingsBadRequest {
	return &GetCharactersCharacterIDFittingsBadRequest{}
}

/* GetCharactersCharacterIDFittingsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDFittingsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDFittingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fittings/][%d] getCharactersCharacterIdFittingsBadRequest  %+v", 400, o.Payload)
}
func (o *GetCharactersCharacterIDFittingsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDFittingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFittingsUnauthorized creates a GetCharactersCharacterIDFittingsUnauthorized with default headers values
func NewGetCharactersCharacterIDFittingsUnauthorized() *GetCharactersCharacterIDFittingsUnauthorized {
	return &GetCharactersCharacterIDFittingsUnauthorized{}
}

/* GetCharactersCharacterIDFittingsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDFittingsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDFittingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fittings/][%d] getCharactersCharacterIdFittingsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCharactersCharacterIDFittingsUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDFittingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFittingsForbidden creates a GetCharactersCharacterIDFittingsForbidden with default headers values
func NewGetCharactersCharacterIDFittingsForbidden() *GetCharactersCharacterIDFittingsForbidden {
	return &GetCharactersCharacterIDFittingsForbidden{}
}

/* GetCharactersCharacterIDFittingsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDFittingsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDFittingsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fittings/][%d] getCharactersCharacterIdFittingsForbidden  %+v", 403, o.Payload)
}
func (o *GetCharactersCharacterIDFittingsForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDFittingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFittingsEnhanceYourCalm creates a GetCharactersCharacterIDFittingsEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDFittingsEnhanceYourCalm() *GetCharactersCharacterIDFittingsEnhanceYourCalm {
	return &GetCharactersCharacterIDFittingsEnhanceYourCalm{}
}

/* GetCharactersCharacterIDFittingsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDFittingsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDFittingsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fittings/][%d] getCharactersCharacterIdFittingsEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCharactersCharacterIDFittingsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDFittingsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFittingsInternalServerError creates a GetCharactersCharacterIDFittingsInternalServerError with default headers values
func NewGetCharactersCharacterIDFittingsInternalServerError() *GetCharactersCharacterIDFittingsInternalServerError {
	return &GetCharactersCharacterIDFittingsInternalServerError{}
}

/* GetCharactersCharacterIDFittingsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDFittingsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDFittingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fittings/][%d] getCharactersCharacterIdFittingsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCharactersCharacterIDFittingsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDFittingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFittingsServiceUnavailable creates a GetCharactersCharacterIDFittingsServiceUnavailable with default headers values
func NewGetCharactersCharacterIDFittingsServiceUnavailable() *GetCharactersCharacterIDFittingsServiceUnavailable {
	return &GetCharactersCharacterIDFittingsServiceUnavailable{}
}

/* GetCharactersCharacterIDFittingsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDFittingsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDFittingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fittings/][%d] getCharactersCharacterIdFittingsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCharactersCharacterIDFittingsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDFittingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFittingsGatewayTimeout creates a GetCharactersCharacterIDFittingsGatewayTimeout with default headers values
func NewGetCharactersCharacterIDFittingsGatewayTimeout() *GetCharactersCharacterIDFittingsGatewayTimeout {
	return &GetCharactersCharacterIDFittingsGatewayTimeout{}
}

/* GetCharactersCharacterIDFittingsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDFittingsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDFittingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fittings/][%d] getCharactersCharacterIdFittingsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCharactersCharacterIDFittingsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDFittingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDFittingsOKBodyItems0 get_characters_character_id_fittings_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDFittingsOKBodyItems0
*/
type GetCharactersCharacterIDFittingsOKBodyItems0 struct {

	// get_characters_character_id_fittings_description
	//
	// description string
	// Required: true
	Description *string `json:"description"`

	// get_characters_character_id_fittings_fitting_id
	//
	// fitting_id integer
	// Required: true
	FittingID *int32 `json:"fitting_id"`

	// get_characters_character_id_fittings_items
	//
	// items array
	// Required: true
	// Max Items: 255
	Items []*GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0 `json:"items"`

	// get_characters_character_id_fittings_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_characters_character_id_fittings_ship_type_id
	//
	// ship_type_id integer
	// Required: true
	ShipTypeID *int32 `json:"ship_type_id"`
}

// Validate validates this get characters character ID fittings o k body items0
func (o *GetCharactersCharacterIDFittingsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFittingID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateShipTypeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDFittingsOKBodyItems0) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDFittingsOKBodyItems0) validateFittingID(formats strfmt.Registry) error {

	if err := validate.Required("fitting_id", "body", o.FittingID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDFittingsOKBodyItems0) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items", "body", o.Items); err != nil {
		return err
	}

	iItemsSize := int64(len(o.Items))

	if err := validate.MaxItems("items", "body", iItemsSize, 255); err != nil {
		return err
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetCharactersCharacterIDFittingsOKBodyItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDFittingsOKBodyItems0) validateShipTypeID(formats strfmt.Registry) error {

	if err := validate.Required("ship_type_id", "body", o.ShipTypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get characters character ID fittings o k body items0 based on the context it is used
func (o *GetCharactersCharacterIDFittingsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDFittingsOKBodyItems0) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {
			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDFittingsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDFittingsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDFittingsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0 get_characters_character_id_fittings_item
//
// item object
swagger:model GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0
*/
type GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0 struct {

	// get_characters_character_id_fittings_flag
	//
	// flag string
	// Required: true
	// Enum: [Cargo DroneBay FighterBay HiSlot0 HiSlot1 HiSlot2 HiSlot3 HiSlot4 HiSlot5 HiSlot6 HiSlot7 Invalid LoSlot0 LoSlot1 LoSlot2 LoSlot3 LoSlot4 LoSlot5 LoSlot6 LoSlot7 MedSlot0 MedSlot1 MedSlot2 MedSlot3 MedSlot4 MedSlot5 MedSlot6 MedSlot7 RigSlot0 RigSlot1 RigSlot2 ServiceSlot0 ServiceSlot1 ServiceSlot2 ServiceSlot3 ServiceSlot4 ServiceSlot5 ServiceSlot6 ServiceSlot7 SubSystemSlot0 SubSystemSlot1 SubSystemSlot2 SubSystemSlot3]
	Flag *string `json:"flag"`

	// get_characters_character_id_fittings_quantity
	//
	// quantity integer
	// Required: true
	Quantity *int32 `json:"quantity"`

	// get_characters_character_id_fittings_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get characters character ID fittings o k body items0 items items0
func (o *GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTypeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getCharactersCharacterIdFittingsOKBodyItems0ItemsItems0TypeFlagPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Cargo","DroneBay","FighterBay","HiSlot0","HiSlot1","HiSlot2","HiSlot3","HiSlot4","HiSlot5","HiSlot6","HiSlot7","Invalid","LoSlot0","LoSlot1","LoSlot2","LoSlot3","LoSlot4","LoSlot5","LoSlot6","LoSlot7","MedSlot0","MedSlot1","MedSlot2","MedSlot3","MedSlot4","MedSlot5","MedSlot6","MedSlot7","RigSlot0","RigSlot1","RigSlot2","ServiceSlot0","ServiceSlot1","ServiceSlot2","ServiceSlot3","ServiceSlot4","ServiceSlot5","ServiceSlot6","ServiceSlot7","SubSystemSlot0","SubSystemSlot1","SubSystemSlot2","SubSystemSlot3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdFittingsOKBodyItems0ItemsItems0TypeFlagPropEnum = append(getCharactersCharacterIdFittingsOKBodyItems0ItemsItems0TypeFlagPropEnum, v)
	}
}

const (

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagCargo captures enum value "Cargo"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagCargo string = "Cargo"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagDroneBay captures enum value "DroneBay"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagDroneBay string = "DroneBay"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagFighterBay captures enum value "FighterBay"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagFighterBay string = "FighterBay"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagHiSlot0 captures enum value "HiSlot0"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagHiSlot0 string = "HiSlot0"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagHiSlot1 captures enum value "HiSlot1"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagHiSlot1 string = "HiSlot1"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagHiSlot2 captures enum value "HiSlot2"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagHiSlot2 string = "HiSlot2"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagHiSlot3 captures enum value "HiSlot3"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagHiSlot3 string = "HiSlot3"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagHiSlot4 captures enum value "HiSlot4"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagHiSlot4 string = "HiSlot4"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagHiSlot5 captures enum value "HiSlot5"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagHiSlot5 string = "HiSlot5"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagHiSlot6 captures enum value "HiSlot6"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagHiSlot6 string = "HiSlot6"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagHiSlot7 captures enum value "HiSlot7"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagHiSlot7 string = "HiSlot7"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagInvalid captures enum value "Invalid"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagInvalid string = "Invalid"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagLoSlot0 captures enum value "LoSlot0"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagLoSlot0 string = "LoSlot0"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagLoSlot1 captures enum value "LoSlot1"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagLoSlot1 string = "LoSlot1"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagLoSlot2 captures enum value "LoSlot2"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagLoSlot2 string = "LoSlot2"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagLoSlot3 captures enum value "LoSlot3"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagLoSlot3 string = "LoSlot3"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagLoSlot4 captures enum value "LoSlot4"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagLoSlot4 string = "LoSlot4"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagLoSlot5 captures enum value "LoSlot5"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagLoSlot5 string = "LoSlot5"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagLoSlot6 captures enum value "LoSlot6"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagLoSlot6 string = "LoSlot6"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagLoSlot7 captures enum value "LoSlot7"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagLoSlot7 string = "LoSlot7"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagMedSlot0 captures enum value "MedSlot0"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagMedSlot0 string = "MedSlot0"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagMedSlot1 captures enum value "MedSlot1"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagMedSlot1 string = "MedSlot1"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagMedSlot2 captures enum value "MedSlot2"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagMedSlot2 string = "MedSlot2"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagMedSlot3 captures enum value "MedSlot3"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagMedSlot3 string = "MedSlot3"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagMedSlot4 captures enum value "MedSlot4"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagMedSlot4 string = "MedSlot4"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagMedSlot5 captures enum value "MedSlot5"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagMedSlot5 string = "MedSlot5"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagMedSlot6 captures enum value "MedSlot6"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagMedSlot6 string = "MedSlot6"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagMedSlot7 captures enum value "MedSlot7"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagMedSlot7 string = "MedSlot7"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagRigSlot0 captures enum value "RigSlot0"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagRigSlot0 string = "RigSlot0"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagRigSlot1 captures enum value "RigSlot1"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagRigSlot1 string = "RigSlot1"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagRigSlot2 captures enum value "RigSlot2"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagRigSlot2 string = "RigSlot2"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagServiceSlot0 captures enum value "ServiceSlot0"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagServiceSlot0 string = "ServiceSlot0"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagServiceSlot1 captures enum value "ServiceSlot1"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagServiceSlot1 string = "ServiceSlot1"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagServiceSlot2 captures enum value "ServiceSlot2"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagServiceSlot2 string = "ServiceSlot2"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagServiceSlot3 captures enum value "ServiceSlot3"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagServiceSlot3 string = "ServiceSlot3"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagServiceSlot4 captures enum value "ServiceSlot4"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagServiceSlot4 string = "ServiceSlot4"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagServiceSlot5 captures enum value "ServiceSlot5"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagServiceSlot5 string = "ServiceSlot5"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagServiceSlot6 captures enum value "ServiceSlot6"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagServiceSlot6 string = "ServiceSlot6"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagServiceSlot7 captures enum value "ServiceSlot7"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagServiceSlot7 string = "ServiceSlot7"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagSubSystemSlot0 captures enum value "SubSystemSlot0"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagSubSystemSlot0 string = "SubSystemSlot0"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagSubSystemSlot1 captures enum value "SubSystemSlot1"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagSubSystemSlot1 string = "SubSystemSlot1"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagSubSystemSlot2 captures enum value "SubSystemSlot2"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagSubSystemSlot2 string = "SubSystemSlot2"

	// GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagSubSystemSlot3 captures enum value "SubSystemSlot3"
	GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0FlagSubSystemSlot3 string = "SubSystemSlot3"
)

// prop value enum
func (o *GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0) validateFlagEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCharactersCharacterIdFittingsOKBodyItems0ItemsItems0TypeFlagPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0) validateFlag(formats strfmt.Registry) error {

	if err := validate.Required("flag", "body", o.Flag); err != nil {
		return err
	}

	// value enum
	if err := o.validateFlagEnum("flag", "body", *o.Flag); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", o.Quantity); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get characters character ID fittings o k body items0 items items0 based on context it is used
func (o *GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDFittingsOKBodyItems0ItemsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}