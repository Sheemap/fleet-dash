// Code generated by go-swagger; DO NOT EDIT.

package fleets

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

// GetFleetsFleetIDReader is a Reader for the GetFleetsFleetID structure.
type GetFleetsFleetIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFleetsFleetIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFleetsFleetIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetFleetsFleetIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetFleetsFleetIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFleetsFleetIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFleetsFleetIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFleetsFleetIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetFleetsFleetIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFleetsFleetIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFleetsFleetIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFleetsFleetIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFleetsFleetIDOK creates a GetFleetsFleetIDOK with default headers values
func NewGetFleetsFleetIDOK() *GetFleetsFleetIDOK {
	return &GetFleetsFleetIDOK{}
}

/* GetFleetsFleetIDOK describes a response with status code 200, with default header values.

Details about a fleet
*/
type GetFleetsFleetIDOK struct {

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

	Payload *GetFleetsFleetIDOKBody
}

func (o *GetFleetsFleetIDOK) Error() string {
	return fmt.Sprintf("[GET /fleets/{fleet_id}/][%d] getFleetsFleetIdOK  %+v", 200, o.Payload)
}
func (o *GetFleetsFleetIDOK) GetPayload() *GetFleetsFleetIDOKBody {
	return o.Payload
}

func (o *GetFleetsFleetIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetFleetsFleetIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDNotModified creates a GetFleetsFleetIDNotModified with default headers values
func NewGetFleetsFleetIDNotModified() *GetFleetsFleetIDNotModified {
	return &GetFleetsFleetIDNotModified{}
}

/* GetFleetsFleetIDNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetFleetsFleetIDNotModified struct {

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

func (o *GetFleetsFleetIDNotModified) Error() string {
	return fmt.Sprintf("[GET /fleets/{fleet_id}/][%d] getFleetsFleetIdNotModified ", 304)
}

func (o *GetFleetsFleetIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFleetsFleetIDBadRequest creates a GetFleetsFleetIDBadRequest with default headers values
func NewGetFleetsFleetIDBadRequest() *GetFleetsFleetIDBadRequest {
	return &GetFleetsFleetIDBadRequest{}
}

/* GetFleetsFleetIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetFleetsFleetIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetFleetsFleetIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /fleets/{fleet_id}/][%d] getFleetsFleetIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetFleetsFleetIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetFleetsFleetIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDUnauthorized creates a GetFleetsFleetIDUnauthorized with default headers values
func NewGetFleetsFleetIDUnauthorized() *GetFleetsFleetIDUnauthorized {
	return &GetFleetsFleetIDUnauthorized{}
}

/* GetFleetsFleetIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetFleetsFleetIDUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetFleetsFleetIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /fleets/{fleet_id}/][%d] getFleetsFleetIdUnauthorized  %+v", 401, o.Payload)
}
func (o *GetFleetsFleetIDUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetFleetsFleetIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDForbidden creates a GetFleetsFleetIDForbidden with default headers values
func NewGetFleetsFleetIDForbidden() *GetFleetsFleetIDForbidden {
	return &GetFleetsFleetIDForbidden{}
}

/* GetFleetsFleetIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetFleetsFleetIDForbidden struct {
	Payload *models.Forbidden
}

func (o *GetFleetsFleetIDForbidden) Error() string {
	return fmt.Sprintf("[GET /fleets/{fleet_id}/][%d] getFleetsFleetIdForbidden  %+v", 403, o.Payload)
}
func (o *GetFleetsFleetIDForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetFleetsFleetIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDNotFound creates a GetFleetsFleetIDNotFound with default headers values
func NewGetFleetsFleetIDNotFound() *GetFleetsFleetIDNotFound {
	return &GetFleetsFleetIDNotFound{}
}

/* GetFleetsFleetIDNotFound describes a response with status code 404, with default header values.

The fleet does not exist or you don't have access to it
*/
type GetFleetsFleetIDNotFound struct {
	Payload *GetFleetsFleetIDNotFoundBody
}

func (o *GetFleetsFleetIDNotFound) Error() string {
	return fmt.Sprintf("[GET /fleets/{fleet_id}/][%d] getFleetsFleetIdNotFound  %+v", 404, o.Payload)
}
func (o *GetFleetsFleetIDNotFound) GetPayload() *GetFleetsFleetIDNotFoundBody {
	return o.Payload
}

func (o *GetFleetsFleetIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetFleetsFleetIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDEnhanceYourCalm creates a GetFleetsFleetIDEnhanceYourCalm with default headers values
func NewGetFleetsFleetIDEnhanceYourCalm() *GetFleetsFleetIDEnhanceYourCalm {
	return &GetFleetsFleetIDEnhanceYourCalm{}
}

/* GetFleetsFleetIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetFleetsFleetIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetFleetsFleetIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /fleets/{fleet_id}/][%d] getFleetsFleetIdEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetFleetsFleetIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetFleetsFleetIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDInternalServerError creates a GetFleetsFleetIDInternalServerError with default headers values
func NewGetFleetsFleetIDInternalServerError() *GetFleetsFleetIDInternalServerError {
	return &GetFleetsFleetIDInternalServerError{}
}

/* GetFleetsFleetIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetFleetsFleetIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetFleetsFleetIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fleets/{fleet_id}/][%d] getFleetsFleetIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetFleetsFleetIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetFleetsFleetIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDServiceUnavailable creates a GetFleetsFleetIDServiceUnavailable with default headers values
func NewGetFleetsFleetIDServiceUnavailable() *GetFleetsFleetIDServiceUnavailable {
	return &GetFleetsFleetIDServiceUnavailable{}
}

/* GetFleetsFleetIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetFleetsFleetIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetFleetsFleetIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /fleets/{fleet_id}/][%d] getFleetsFleetIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetFleetsFleetIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetFleetsFleetIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDGatewayTimeout creates a GetFleetsFleetIDGatewayTimeout with default headers values
func NewGetFleetsFleetIDGatewayTimeout() *GetFleetsFleetIDGatewayTimeout {
	return &GetFleetsFleetIDGatewayTimeout{}
}

/* GetFleetsFleetIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetFleetsFleetIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetFleetsFleetIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /fleets/{fleet_id}/][%d] getFleetsFleetIdGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetFleetsFleetIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetFleetsFleetIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetFleetsFleetIDNotFoundBody get_fleets_fleet_id_not_found
//
// Not found
swagger:model GetFleetsFleetIDNotFoundBody
*/
type GetFleetsFleetIDNotFoundBody struct {

	// get_fleets_fleet_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get fleets fleet ID not found body
func (o *GetFleetsFleetIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get fleets fleet ID not found body based on context it is used
func (o *GetFleetsFleetIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetFleetsFleetIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFleetsFleetIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetFleetsFleetIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetFleetsFleetIDOKBody get_fleets_fleet_id_ok
//
// 200 ok object
swagger:model GetFleetsFleetIDOKBody
*/
type GetFleetsFleetIDOKBody struct {

	// get_fleets_fleet_id_is_free_move
	//
	// Is free-move enabled
	// Required: true
	IsFreeMove *bool `json:"is_free_move"`

	// get_fleets_fleet_id_is_registered
	//
	// Does the fleet have an active fleet advertisement
	// Required: true
	IsRegistered *bool `json:"is_registered"`

	// get_fleets_fleet_id_is_voice_enabled
	//
	// Is EVE Voice enabled
	// Required: true
	IsVoiceEnabled *bool `json:"is_voice_enabled"`

	// get_fleets_fleet_id_motd
	//
	// Fleet MOTD in CCP flavoured HTML
	// Required: true
	Motd *string `json:"motd"`
}

// Validate validates this get fleets fleet ID o k body
func (o *GetFleetsFleetIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIsFreeMove(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIsRegistered(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIsVoiceEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMotd(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFleetsFleetIDOKBody) validateIsFreeMove(formats strfmt.Registry) error {

	if err := validate.Required("getFleetsFleetIdOK"+"."+"is_free_move", "body", o.IsFreeMove); err != nil {
		return err
	}

	return nil
}

func (o *GetFleetsFleetIDOKBody) validateIsRegistered(formats strfmt.Registry) error {

	if err := validate.Required("getFleetsFleetIdOK"+"."+"is_registered", "body", o.IsRegistered); err != nil {
		return err
	}

	return nil
}

func (o *GetFleetsFleetIDOKBody) validateIsVoiceEnabled(formats strfmt.Registry) error {

	if err := validate.Required("getFleetsFleetIdOK"+"."+"is_voice_enabled", "body", o.IsVoiceEnabled); err != nil {
		return err
	}

	return nil
}

func (o *GetFleetsFleetIDOKBody) validateMotd(formats strfmt.Registry) error {

	if err := validate.Required("getFleetsFleetIdOK"+"."+"motd", "body", o.Motd); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get fleets fleet ID o k body based on context it is used
func (o *GetFleetsFleetIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetFleetsFleetIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFleetsFleetIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetFleetsFleetIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}