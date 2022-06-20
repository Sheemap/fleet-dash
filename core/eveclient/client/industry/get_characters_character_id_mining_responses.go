// Code generated by go-swagger; DO NOT EDIT.

package industry

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

// GetCharactersCharacterIDMiningReader is a Reader for the GetCharactersCharacterIDMining structure.
type GetCharactersCharacterIDMiningReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDMiningReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDMiningOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDMiningNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDMiningBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDMiningUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDMiningForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDMiningEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDMiningInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDMiningServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDMiningGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersCharacterIDMiningOK creates a GetCharactersCharacterIDMiningOK with default headers values
func NewGetCharactersCharacterIDMiningOK() *GetCharactersCharacterIDMiningOK {
	var (
		// initialize headers with default values
		xPagesDefault = int32(1)
	)

	return &GetCharactersCharacterIDMiningOK{

		XPages: xPagesDefault,
	}
}

/* GetCharactersCharacterIDMiningOK describes a response with status code 200, with default header values.

Mining ledger of a character
*/
type GetCharactersCharacterIDMiningOK struct {

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

	Payload []*GetCharactersCharacterIDMiningOKBodyItems0
}

func (o *GetCharactersCharacterIDMiningOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mining/][%d] getCharactersCharacterIdMiningOK  %+v", 200, o.Payload)
}
func (o *GetCharactersCharacterIDMiningOK) GetPayload() []*GetCharactersCharacterIDMiningOKBodyItems0 {
	return o.Payload
}

func (o *GetCharactersCharacterIDMiningOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDMiningNotModified creates a GetCharactersCharacterIDMiningNotModified with default headers values
func NewGetCharactersCharacterIDMiningNotModified() *GetCharactersCharacterIDMiningNotModified {
	return &GetCharactersCharacterIDMiningNotModified{}
}

/* GetCharactersCharacterIDMiningNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDMiningNotModified struct {

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

func (o *GetCharactersCharacterIDMiningNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mining/][%d] getCharactersCharacterIdMiningNotModified ", 304)
}

func (o *GetCharactersCharacterIDMiningNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDMiningBadRequest creates a GetCharactersCharacterIDMiningBadRequest with default headers values
func NewGetCharactersCharacterIDMiningBadRequest() *GetCharactersCharacterIDMiningBadRequest {
	return &GetCharactersCharacterIDMiningBadRequest{}
}

/* GetCharactersCharacterIDMiningBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDMiningBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDMiningBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mining/][%d] getCharactersCharacterIdMiningBadRequest  %+v", 400, o.Payload)
}
func (o *GetCharactersCharacterIDMiningBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDMiningBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMiningUnauthorized creates a GetCharactersCharacterIDMiningUnauthorized with default headers values
func NewGetCharactersCharacterIDMiningUnauthorized() *GetCharactersCharacterIDMiningUnauthorized {
	return &GetCharactersCharacterIDMiningUnauthorized{}
}

/* GetCharactersCharacterIDMiningUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDMiningUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDMiningUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mining/][%d] getCharactersCharacterIdMiningUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCharactersCharacterIDMiningUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDMiningUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMiningForbidden creates a GetCharactersCharacterIDMiningForbidden with default headers values
func NewGetCharactersCharacterIDMiningForbidden() *GetCharactersCharacterIDMiningForbidden {
	return &GetCharactersCharacterIDMiningForbidden{}
}

/* GetCharactersCharacterIDMiningForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDMiningForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDMiningForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mining/][%d] getCharactersCharacterIdMiningForbidden  %+v", 403, o.Payload)
}
func (o *GetCharactersCharacterIDMiningForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDMiningForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMiningEnhanceYourCalm creates a GetCharactersCharacterIDMiningEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDMiningEnhanceYourCalm() *GetCharactersCharacterIDMiningEnhanceYourCalm {
	return &GetCharactersCharacterIDMiningEnhanceYourCalm{}
}

/* GetCharactersCharacterIDMiningEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDMiningEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDMiningEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mining/][%d] getCharactersCharacterIdMiningEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCharactersCharacterIDMiningEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDMiningEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMiningInternalServerError creates a GetCharactersCharacterIDMiningInternalServerError with default headers values
func NewGetCharactersCharacterIDMiningInternalServerError() *GetCharactersCharacterIDMiningInternalServerError {
	return &GetCharactersCharacterIDMiningInternalServerError{}
}

/* GetCharactersCharacterIDMiningInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDMiningInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDMiningInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mining/][%d] getCharactersCharacterIdMiningInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCharactersCharacterIDMiningInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDMiningInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMiningServiceUnavailable creates a GetCharactersCharacterIDMiningServiceUnavailable with default headers values
func NewGetCharactersCharacterIDMiningServiceUnavailable() *GetCharactersCharacterIDMiningServiceUnavailable {
	return &GetCharactersCharacterIDMiningServiceUnavailable{}
}

/* GetCharactersCharacterIDMiningServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDMiningServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDMiningServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mining/][%d] getCharactersCharacterIdMiningServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCharactersCharacterIDMiningServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDMiningServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMiningGatewayTimeout creates a GetCharactersCharacterIDMiningGatewayTimeout with default headers values
func NewGetCharactersCharacterIDMiningGatewayTimeout() *GetCharactersCharacterIDMiningGatewayTimeout {
	return &GetCharactersCharacterIDMiningGatewayTimeout{}
}

/* GetCharactersCharacterIDMiningGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDMiningGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDMiningGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mining/][%d] getCharactersCharacterIdMiningGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCharactersCharacterIDMiningGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDMiningGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDMiningOKBodyItems0 get_characters_character_id_mining_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDMiningOKBodyItems0
*/
type GetCharactersCharacterIDMiningOKBodyItems0 struct {

	// get_characters_character_id_mining_date
	//
	// date string
	// Required: true
	// Format: date
	Date *strfmt.Date `json:"date"`

	// get_characters_character_id_mining_quantity
	//
	// quantity integer
	// Required: true
	Quantity *int64 `json:"quantity"`

	// get_characters_character_id_mining_solar_system_id
	//
	// solar_system_id integer
	// Required: true
	SolarSystemID *int32 `json:"solar_system_id"`

	// get_characters_character_id_mining_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get characters character ID mining o k body items0
func (o *GetCharactersCharacterIDMiningOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSolarSystemID(formats); err != nil {
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

func (o *GetCharactersCharacterIDMiningOKBodyItems0) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", o.Date); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date", o.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDMiningOKBodyItems0) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", o.Quantity); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDMiningOKBodyItems0) validateSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("solar_system_id", "body", o.SolarSystemID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDMiningOKBodyItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get characters character ID mining o k body items0 based on context it is used
func (o *GetCharactersCharacterIDMiningOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDMiningOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDMiningOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDMiningOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}