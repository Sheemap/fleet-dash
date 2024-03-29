// Code generated by go-swagger; DO NOT EDIT.

package fleets

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

// GetCharactersCharacterIDFleetReader is a Reader for the GetCharactersCharacterIDFleet structure.
type GetCharactersCharacterIDFleetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDFleetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDFleetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDFleetNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDFleetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDFleetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDFleetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCharactersCharacterIDFleetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDFleetEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDFleetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDFleetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDFleetGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersCharacterIDFleetOK creates a GetCharactersCharacterIDFleetOK with default headers values
func NewGetCharactersCharacterIDFleetOK() *GetCharactersCharacterIDFleetOK {
	return &GetCharactersCharacterIDFleetOK{}
}

/* GetCharactersCharacterIDFleetOK describes a response with status code 200, with default header values.

Details about the character's fleet
*/
type GetCharactersCharacterIDFleetOK struct {

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

	Payload *GetCharactersCharacterIDFleetOKBody
}

func (o *GetCharactersCharacterIDFleetOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetOK  %+v", 200, o.Payload)
}
func (o *GetCharactersCharacterIDFleetOK) GetPayload() *GetCharactersCharacterIDFleetOKBody {
	return o.Payload
}

func (o *GetCharactersCharacterIDFleetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetCharactersCharacterIDFleetOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFleetNotModified creates a GetCharactersCharacterIDFleetNotModified with default headers values
func NewGetCharactersCharacterIDFleetNotModified() *GetCharactersCharacterIDFleetNotModified {
	return &GetCharactersCharacterIDFleetNotModified{}
}

/* GetCharactersCharacterIDFleetNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDFleetNotModified struct {

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

func (o *GetCharactersCharacterIDFleetNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetNotModified ", 304)
}

func (o *GetCharactersCharacterIDFleetNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDFleetBadRequest creates a GetCharactersCharacterIDFleetBadRequest with default headers values
func NewGetCharactersCharacterIDFleetBadRequest() *GetCharactersCharacterIDFleetBadRequest {
	return &GetCharactersCharacterIDFleetBadRequest{}
}

/* GetCharactersCharacterIDFleetBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDFleetBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDFleetBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetBadRequest  %+v", 400, o.Payload)
}
func (o *GetCharactersCharacterIDFleetBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDFleetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFleetUnauthorized creates a GetCharactersCharacterIDFleetUnauthorized with default headers values
func NewGetCharactersCharacterIDFleetUnauthorized() *GetCharactersCharacterIDFleetUnauthorized {
	return &GetCharactersCharacterIDFleetUnauthorized{}
}

/* GetCharactersCharacterIDFleetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDFleetUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDFleetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCharactersCharacterIDFleetUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDFleetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFleetForbidden creates a GetCharactersCharacterIDFleetForbidden with default headers values
func NewGetCharactersCharacterIDFleetForbidden() *GetCharactersCharacterIDFleetForbidden {
	return &GetCharactersCharacterIDFleetForbidden{}
}

/* GetCharactersCharacterIDFleetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDFleetForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDFleetForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetForbidden  %+v", 403, o.Payload)
}
func (o *GetCharactersCharacterIDFleetForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDFleetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFleetNotFound creates a GetCharactersCharacterIDFleetNotFound with default headers values
func NewGetCharactersCharacterIDFleetNotFound() *GetCharactersCharacterIDFleetNotFound {
	return &GetCharactersCharacterIDFleetNotFound{}
}

/* GetCharactersCharacterIDFleetNotFound describes a response with status code 404, with default header values.

The character is not in a fleet
*/
type GetCharactersCharacterIDFleetNotFound struct {
	Payload *GetCharactersCharacterIDFleetNotFoundBody
}

func (o *GetCharactersCharacterIDFleetNotFound) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetNotFound  %+v", 404, o.Payload)
}
func (o *GetCharactersCharacterIDFleetNotFound) GetPayload() *GetCharactersCharacterIDFleetNotFoundBody {
	return o.Payload
}

func (o *GetCharactersCharacterIDFleetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCharactersCharacterIDFleetNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFleetEnhanceYourCalm creates a GetCharactersCharacterIDFleetEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDFleetEnhanceYourCalm() *GetCharactersCharacterIDFleetEnhanceYourCalm {
	return &GetCharactersCharacterIDFleetEnhanceYourCalm{}
}

/* GetCharactersCharacterIDFleetEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDFleetEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDFleetEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCharactersCharacterIDFleetEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDFleetEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFleetInternalServerError creates a GetCharactersCharacterIDFleetInternalServerError with default headers values
func NewGetCharactersCharacterIDFleetInternalServerError() *GetCharactersCharacterIDFleetInternalServerError {
	return &GetCharactersCharacterIDFleetInternalServerError{}
}

/* GetCharactersCharacterIDFleetInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDFleetInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDFleetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCharactersCharacterIDFleetInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDFleetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFleetServiceUnavailable creates a GetCharactersCharacterIDFleetServiceUnavailable with default headers values
func NewGetCharactersCharacterIDFleetServiceUnavailable() *GetCharactersCharacterIDFleetServiceUnavailable {
	return &GetCharactersCharacterIDFleetServiceUnavailable{}
}

/* GetCharactersCharacterIDFleetServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDFleetServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDFleetServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCharactersCharacterIDFleetServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDFleetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFleetGatewayTimeout creates a GetCharactersCharacterIDFleetGatewayTimeout with default headers values
func NewGetCharactersCharacterIDFleetGatewayTimeout() *GetCharactersCharacterIDFleetGatewayTimeout {
	return &GetCharactersCharacterIDFleetGatewayTimeout{}
}

/* GetCharactersCharacterIDFleetGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDFleetGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDFleetGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCharactersCharacterIDFleetGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDFleetGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDFleetNotFoundBody get_characters_character_id_fleet_not_found
//
// Not found
swagger:model GetCharactersCharacterIDFleetNotFoundBody
*/
type GetCharactersCharacterIDFleetNotFoundBody struct {

	// get_characters_character_id_fleet_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get characters character ID fleet not found body
func (o *GetCharactersCharacterIDFleetNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get characters character ID fleet not found body based on context it is used
func (o *GetCharactersCharacterIDFleetNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDFleetNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDFleetNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDFleetNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetCharactersCharacterIDFleetOKBody get_characters_character_id_fleet_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDFleetOKBody
*/
type GetCharactersCharacterIDFleetOKBody struct {

	// get_characters_character_id_fleet_fleet_id
	//
	// The character's current fleet ID
	// Required: true
	FleetID *int64 `json:"fleet_id"`

	// get_characters_character_id_fleet_role
	//
	// Member’s role in fleet
	// Required: true
	// Enum: [fleet_commander squad_commander squad_member wing_commander]
	Role *string `json:"role"`

	// get_characters_character_id_fleet_squad_id
	//
	// ID of the squad the member is in. If not applicable, will be set to -1
	// Required: true
	SquadID *int64 `json:"squad_id"`

	// get_characters_character_id_fleet_wing_id
	//
	// ID of the wing the member is in. If not applicable, will be set to -1
	// Required: true
	WingID *int64 `json:"wing_id"`
}

// Validate validates this get characters character ID fleet o k body
func (o *GetCharactersCharacterIDFleetOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFleetID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSquadID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWingID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDFleetOKBody) validateFleetID(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdFleetOK"+"."+"fleet_id", "body", o.FleetID); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdFleetOKBodyTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fleet_commander","squad_commander","squad_member","wing_commander"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdFleetOKBodyTypeRolePropEnum = append(getCharactersCharacterIdFleetOKBodyTypeRolePropEnum, v)
	}
}

const (

	// GetCharactersCharacterIDFleetOKBodyRoleFleetCommander captures enum value "fleet_commander"
	GetCharactersCharacterIDFleetOKBodyRoleFleetCommander string = "fleet_commander"

	// GetCharactersCharacterIDFleetOKBodyRoleSquadCommander captures enum value "squad_commander"
	GetCharactersCharacterIDFleetOKBodyRoleSquadCommander string = "squad_commander"

	// GetCharactersCharacterIDFleetOKBodyRoleSquadMember captures enum value "squad_member"
	GetCharactersCharacterIDFleetOKBodyRoleSquadMember string = "squad_member"

	// GetCharactersCharacterIDFleetOKBodyRoleWingCommander captures enum value "wing_commander"
	GetCharactersCharacterIDFleetOKBodyRoleWingCommander string = "wing_commander"
)

// prop value enum
func (o *GetCharactersCharacterIDFleetOKBody) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCharactersCharacterIdFleetOKBodyTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDFleetOKBody) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdFleetOK"+"."+"role", "body", o.Role); err != nil {
		return err
	}

	// value enum
	if err := o.validateRoleEnum("getCharactersCharacterIdFleetOK"+"."+"role", "body", *o.Role); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDFleetOKBody) validateSquadID(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdFleetOK"+"."+"squad_id", "body", o.SquadID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDFleetOKBody) validateWingID(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdFleetOK"+"."+"wing_id", "body", o.WingID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get characters character ID fleet o k body based on context it is used
func (o *GetCharactersCharacterIDFleetOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDFleetOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDFleetOKBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDFleetOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
