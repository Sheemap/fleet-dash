// Code generated by go-swagger; DO NOT EDIT.

package universe

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

// GetUniverseAsteroidBeltsAsteroidBeltIDReader is a Reader for the GetUniverseAsteroidBeltsAsteroidBeltID structure.
type GetUniverseAsteroidBeltsAsteroidBeltIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUniverseAsteroidBeltsAsteroidBeltIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetUniverseAsteroidBeltsAsteroidBeltIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetUniverseAsteroidBeltsAsteroidBeltIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUniverseAsteroidBeltsAsteroidBeltIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUniverseAsteroidBeltsAsteroidBeltIDOK creates a GetUniverseAsteroidBeltsAsteroidBeltIDOK with default headers values
func NewGetUniverseAsteroidBeltsAsteroidBeltIDOK() *GetUniverseAsteroidBeltsAsteroidBeltIDOK {
	return &GetUniverseAsteroidBeltsAsteroidBeltIDOK{}
}

/* GetUniverseAsteroidBeltsAsteroidBeltIDOK describes a response with status code 200, with default header values.

Information about an asteroid belt
*/
type GetUniverseAsteroidBeltsAsteroidBeltIDOK struct {

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

	Payload *GetUniverseAsteroidBeltsAsteroidBeltIDOKBody
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDOK) Error() string {
	return fmt.Sprintf("[GET /universe/asteroid_belts/{asteroid_belt_id}/][%d] getUniverseAsteroidBeltsAsteroidBeltIdOK  %+v", 200, o.Payload)
}
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDOK) GetPayload() *GetUniverseAsteroidBeltsAsteroidBeltIDOKBody {
	return o.Payload
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetUniverseAsteroidBeltsAsteroidBeltIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseAsteroidBeltsAsteroidBeltIDNotModified creates a GetUniverseAsteroidBeltsAsteroidBeltIDNotModified with default headers values
func NewGetUniverseAsteroidBeltsAsteroidBeltIDNotModified() *GetUniverseAsteroidBeltsAsteroidBeltIDNotModified {
	return &GetUniverseAsteroidBeltsAsteroidBeltIDNotModified{}
}

/* GetUniverseAsteroidBeltsAsteroidBeltIDNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetUniverseAsteroidBeltsAsteroidBeltIDNotModified struct {

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

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDNotModified) Error() string {
	return fmt.Sprintf("[GET /universe/asteroid_belts/{asteroid_belt_id}/][%d] getUniverseAsteroidBeltsAsteroidBeltIdNotModified ", 304)
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseAsteroidBeltsAsteroidBeltIDBadRequest creates a GetUniverseAsteroidBeltsAsteroidBeltIDBadRequest with default headers values
func NewGetUniverseAsteroidBeltsAsteroidBeltIDBadRequest() *GetUniverseAsteroidBeltsAsteroidBeltIDBadRequest {
	return &GetUniverseAsteroidBeltsAsteroidBeltIDBadRequest{}
}

/* GetUniverseAsteroidBeltsAsteroidBeltIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUniverseAsteroidBeltsAsteroidBeltIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /universe/asteroid_belts/{asteroid_belt_id}/][%d] getUniverseAsteroidBeltsAsteroidBeltIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseAsteroidBeltsAsteroidBeltIDNotFound creates a GetUniverseAsteroidBeltsAsteroidBeltIDNotFound with default headers values
func NewGetUniverseAsteroidBeltsAsteroidBeltIDNotFound() *GetUniverseAsteroidBeltsAsteroidBeltIDNotFound {
	return &GetUniverseAsteroidBeltsAsteroidBeltIDNotFound{}
}

/* GetUniverseAsteroidBeltsAsteroidBeltIDNotFound describes a response with status code 404, with default header values.

Asteroid belt not found
*/
type GetUniverseAsteroidBeltsAsteroidBeltIDNotFound struct {
	Payload *GetUniverseAsteroidBeltsAsteroidBeltIDNotFoundBody
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDNotFound) Error() string {
	return fmt.Sprintf("[GET /universe/asteroid_belts/{asteroid_belt_id}/][%d] getUniverseAsteroidBeltsAsteroidBeltIdNotFound  %+v", 404, o.Payload)
}
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDNotFound) GetPayload() *GetUniverseAsteroidBeltsAsteroidBeltIDNotFoundBody {
	return o.Payload
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetUniverseAsteroidBeltsAsteroidBeltIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm creates a GetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm with default headers values
func NewGetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm() *GetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm {
	return &GetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm{}
}

/* GetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /universe/asteroid_belts/{asteroid_belt_id}/][%d] getUniverseAsteroidBeltsAsteroidBeltIdEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError creates a GetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError with default headers values
func NewGetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError() *GetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError {
	return &GetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError{}
}

/* GetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/asteroid_belts/{asteroid_belt_id}/][%d] getUniverseAsteroidBeltsAsteroidBeltIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable creates a GetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable with default headers values
func NewGetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable() *GetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable {
	return &GetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable{}
}

/* GetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /universe/asteroid_belts/{asteroid_belt_id}/][%d] getUniverseAsteroidBeltsAsteroidBeltIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout creates a GetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout with default headers values
func NewGetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout() *GetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout {
	return &GetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout{}
}

/* GetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /universe/asteroid_belts/{asteroid_belt_id}/][%d] getUniverseAsteroidBeltsAsteroidBeltIdGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetUniverseAsteroidBeltsAsteroidBeltIDNotFoundBody get_universe_asteroid_belts_asteroid_belt_id_not_found
//
// Not found
swagger:model GetUniverseAsteroidBeltsAsteroidBeltIDNotFoundBody
*/
type GetUniverseAsteroidBeltsAsteroidBeltIDNotFoundBody struct {

	// get_universe_asteroid_belts_asteroid_belt_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get universe asteroid belts asteroid belt ID not found body
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get universe asteroid belts asteroid belt ID not found body based on context it is used
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseAsteroidBeltsAsteroidBeltIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUniverseAsteroidBeltsAsteroidBeltIDOKBody get_universe_asteroid_belts_asteroid_belt_id_ok
//
// 200 ok object
swagger:model GetUniverseAsteroidBeltsAsteroidBeltIDOKBody
*/
type GetUniverseAsteroidBeltsAsteroidBeltIDOKBody struct {

	// get_universe_asteroid_belts_asteroid_belt_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// position
	// Required: true
	Position *GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition `json:"position"`

	// get_universe_asteroid_belts_asteroid_belt_id_system_id
	//
	// The solar system this asteroid belt is in
	// Required: true
	SystemID *int32 `json:"system_id"`
}

// Validate validates this get universe asteroid belts asteroid belt ID o k body
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSystemID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseAsteroidBeltsAsteroidBeltIdOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDOKBody) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseAsteroidBeltsAsteroidBeltIdOK"+"."+"position", "body", o.Position); err != nil {
		return err
	}

	if o.Position != nil {
		if err := o.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUniverseAsteroidBeltsAsteroidBeltIdOK" + "." + "position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getUniverseAsteroidBeltsAsteroidBeltIdOK" + "." + "position")
			}
			return err
		}
	}

	return nil
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDOKBody) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseAsteroidBeltsAsteroidBeltIdOK"+"."+"system_id", "body", o.SystemID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get universe asteroid belts asteroid belt ID o k body based on the context it is used
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDOKBody) contextValidatePosition(ctx context.Context, formats strfmt.Registry) error {

	if o.Position != nil {
		if err := o.Position.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUniverseAsteroidBeltsAsteroidBeltIdOK" + "." + "position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getUniverseAsteroidBeltsAsteroidBeltIdOK" + "." + "position")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseAsteroidBeltsAsteroidBeltIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition get_universe_asteroid_belts_asteroid_belt_id_position
//
// position object
swagger:model GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition
*/
type GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition struct {

	// get_universe_asteroid_belts_asteroid_belt_id_x
	//
	// x number
	// Required: true
	X *float64 `json:"x"`

	// get_universe_asteroid_belts_asteroid_belt_id_y
	//
	// y number
	// Required: true
	Y *float64 `json:"y"`

	// get_universe_asteroid_belts_asteroid_belt_id_z
	//
	// z number
	// Required: true
	Z *float64 `json:"z"`
}

// Validate validates this get universe asteroid belts asteroid belt ID o k body position
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateX(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateY(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateZ(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition) validateX(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseAsteroidBeltsAsteroidBeltIdOK"+"."+"position"+"."+"x", "body", o.X); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition) validateY(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseAsteroidBeltsAsteroidBeltIdOK"+"."+"position"+"."+"y", "body", o.Y); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition) validateZ(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseAsteroidBeltsAsteroidBeltIdOK"+"."+"position"+"."+"z", "body", o.Z); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get universe asteroid belts asteroid belt ID o k body position based on context it is used
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition) UnmarshalBinary(b []byte) error {
	var res GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
