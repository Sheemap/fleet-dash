// Code generated by go-swagger; DO NOT EDIT.

package location

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

// GetCharactersCharacterIDLocationReader is a Reader for the GetCharactersCharacterIDLocation structure.
type GetCharactersCharacterIDLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDLocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDLocationNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDLocationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDLocationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDLocationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDLocationEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDLocationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDLocationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDLocationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersCharacterIDLocationOK creates a GetCharactersCharacterIDLocationOK with default headers values
func NewGetCharactersCharacterIDLocationOK() *GetCharactersCharacterIDLocationOK {
	return &GetCharactersCharacterIDLocationOK{}
}

/* GetCharactersCharacterIDLocationOK describes a response with status code 200, with default header values.

Information about the characters current location. Returns the current solar system id, and also the current station or structure ID if applicable
*/
type GetCharactersCharacterIDLocationOK struct {

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

	Payload *GetCharactersCharacterIDLocationOKBody
}

func (o *GetCharactersCharacterIDLocationOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/location/][%d] getCharactersCharacterIdLocationOK  %+v", 200, o.Payload)
}
func (o *GetCharactersCharacterIDLocationOK) GetPayload() *GetCharactersCharacterIDLocationOKBody {
	return o.Payload
}

func (o *GetCharactersCharacterIDLocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetCharactersCharacterIDLocationOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDLocationNotModified creates a GetCharactersCharacterIDLocationNotModified with default headers values
func NewGetCharactersCharacterIDLocationNotModified() *GetCharactersCharacterIDLocationNotModified {
	return &GetCharactersCharacterIDLocationNotModified{}
}

/* GetCharactersCharacterIDLocationNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDLocationNotModified struct {

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

func (o *GetCharactersCharacterIDLocationNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/location/][%d] getCharactersCharacterIdLocationNotModified ", 304)
}

func (o *GetCharactersCharacterIDLocationNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDLocationBadRequest creates a GetCharactersCharacterIDLocationBadRequest with default headers values
func NewGetCharactersCharacterIDLocationBadRequest() *GetCharactersCharacterIDLocationBadRequest {
	return &GetCharactersCharacterIDLocationBadRequest{}
}

/* GetCharactersCharacterIDLocationBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDLocationBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDLocationBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/location/][%d] getCharactersCharacterIdLocationBadRequest  %+v", 400, o.Payload)
}
func (o *GetCharactersCharacterIDLocationBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDLocationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDLocationUnauthorized creates a GetCharactersCharacterIDLocationUnauthorized with default headers values
func NewGetCharactersCharacterIDLocationUnauthorized() *GetCharactersCharacterIDLocationUnauthorized {
	return &GetCharactersCharacterIDLocationUnauthorized{}
}

/* GetCharactersCharacterIDLocationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDLocationUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDLocationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/location/][%d] getCharactersCharacterIdLocationUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCharactersCharacterIDLocationUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDLocationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDLocationForbidden creates a GetCharactersCharacterIDLocationForbidden with default headers values
func NewGetCharactersCharacterIDLocationForbidden() *GetCharactersCharacterIDLocationForbidden {
	return &GetCharactersCharacterIDLocationForbidden{}
}

/* GetCharactersCharacterIDLocationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDLocationForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDLocationForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/location/][%d] getCharactersCharacterIdLocationForbidden  %+v", 403, o.Payload)
}
func (o *GetCharactersCharacterIDLocationForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDLocationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDLocationEnhanceYourCalm creates a GetCharactersCharacterIDLocationEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDLocationEnhanceYourCalm() *GetCharactersCharacterIDLocationEnhanceYourCalm {
	return &GetCharactersCharacterIDLocationEnhanceYourCalm{}
}

/* GetCharactersCharacterIDLocationEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDLocationEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDLocationEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/location/][%d] getCharactersCharacterIdLocationEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCharactersCharacterIDLocationEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDLocationEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDLocationInternalServerError creates a GetCharactersCharacterIDLocationInternalServerError with default headers values
func NewGetCharactersCharacterIDLocationInternalServerError() *GetCharactersCharacterIDLocationInternalServerError {
	return &GetCharactersCharacterIDLocationInternalServerError{}
}

/* GetCharactersCharacterIDLocationInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDLocationInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDLocationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/location/][%d] getCharactersCharacterIdLocationInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCharactersCharacterIDLocationInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDLocationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDLocationServiceUnavailable creates a GetCharactersCharacterIDLocationServiceUnavailable with default headers values
func NewGetCharactersCharacterIDLocationServiceUnavailable() *GetCharactersCharacterIDLocationServiceUnavailable {
	return &GetCharactersCharacterIDLocationServiceUnavailable{}
}

/* GetCharactersCharacterIDLocationServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDLocationServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDLocationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/location/][%d] getCharactersCharacterIdLocationServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCharactersCharacterIDLocationServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDLocationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDLocationGatewayTimeout creates a GetCharactersCharacterIDLocationGatewayTimeout with default headers values
func NewGetCharactersCharacterIDLocationGatewayTimeout() *GetCharactersCharacterIDLocationGatewayTimeout {
	return &GetCharactersCharacterIDLocationGatewayTimeout{}
}

/* GetCharactersCharacterIDLocationGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDLocationGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDLocationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/location/][%d] getCharactersCharacterIdLocationGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCharactersCharacterIDLocationGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDLocationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDLocationOKBody get_characters_character_id_location_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDLocationOKBody
*/
type GetCharactersCharacterIDLocationOKBody struct {

	// get_characters_character_id_location_solar_system_id
	//
	// solar_system_id integer
	// Required: true
	SolarSystemID *int32 `json:"solar_system_id"`

	// get_characters_character_id_location_station_id
	//
	// station_id integer
	StationID int32 `json:"station_id,omitempty"`

	// get_characters_character_id_location_structure_id
	//
	// structure_id integer
	StructureID int64 `json:"structure_id,omitempty"`
}

// Validate validates this get characters character ID location o k body
func (o *GetCharactersCharacterIDLocationOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSolarSystemID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDLocationOKBody) validateSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdLocationOK"+"."+"solar_system_id", "body", o.SolarSystemID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get characters character ID location o k body based on context it is used
func (o *GetCharactersCharacterIDLocationOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDLocationOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDLocationOKBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDLocationOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
