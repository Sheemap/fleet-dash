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

// GetUniverseStargatesStargateIDReader is a Reader for the GetUniverseStargatesStargateID structure.
type GetUniverseStargatesStargateIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseStargatesStargateIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUniverseStargatesStargateIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetUniverseStargatesStargateIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetUniverseStargatesStargateIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUniverseStargatesStargateIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetUniverseStargatesStargateIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUniverseStargatesStargateIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUniverseStargatesStargateIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUniverseStargatesStargateIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUniverseStargatesStargateIDOK creates a GetUniverseStargatesStargateIDOK with default headers values
func NewGetUniverseStargatesStargateIDOK() *GetUniverseStargatesStargateIDOK {
	return &GetUniverseStargatesStargateIDOK{}
}

/* GetUniverseStargatesStargateIDOK describes a response with status code 200, with default header values.

Information about a stargate
*/
type GetUniverseStargatesStargateIDOK struct {

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

	Payload *GetUniverseStargatesStargateIDOKBody
}

func (o *GetUniverseStargatesStargateIDOK) Error() string {
	return fmt.Sprintf("[GET /universe/stargates/{stargate_id}/][%d] getUniverseStargatesStargateIdOK  %+v", 200, o.Payload)
}
func (o *GetUniverseStargatesStargateIDOK) GetPayload() *GetUniverseStargatesStargateIDOKBody {
	return o.Payload
}

func (o *GetUniverseStargatesStargateIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetUniverseStargatesStargateIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStargatesStargateIDNotModified creates a GetUniverseStargatesStargateIDNotModified with default headers values
func NewGetUniverseStargatesStargateIDNotModified() *GetUniverseStargatesStargateIDNotModified {
	return &GetUniverseStargatesStargateIDNotModified{}
}

/* GetUniverseStargatesStargateIDNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetUniverseStargatesStargateIDNotModified struct {

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

func (o *GetUniverseStargatesStargateIDNotModified) Error() string {
	return fmt.Sprintf("[GET /universe/stargates/{stargate_id}/][%d] getUniverseStargatesStargateIdNotModified ", 304)
}

func (o *GetUniverseStargatesStargateIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseStargatesStargateIDBadRequest creates a GetUniverseStargatesStargateIDBadRequest with default headers values
func NewGetUniverseStargatesStargateIDBadRequest() *GetUniverseStargatesStargateIDBadRequest {
	return &GetUniverseStargatesStargateIDBadRequest{}
}

/* GetUniverseStargatesStargateIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUniverseStargatesStargateIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetUniverseStargatesStargateIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /universe/stargates/{stargate_id}/][%d] getUniverseStargatesStargateIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetUniverseStargatesStargateIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetUniverseStargatesStargateIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStargatesStargateIDNotFound creates a GetUniverseStargatesStargateIDNotFound with default headers values
func NewGetUniverseStargatesStargateIDNotFound() *GetUniverseStargatesStargateIDNotFound {
	return &GetUniverseStargatesStargateIDNotFound{}
}

/* GetUniverseStargatesStargateIDNotFound describes a response with status code 404, with default header values.

Stargate not found
*/
type GetUniverseStargatesStargateIDNotFound struct {
	Payload *GetUniverseStargatesStargateIDNotFoundBody
}

func (o *GetUniverseStargatesStargateIDNotFound) Error() string {
	return fmt.Sprintf("[GET /universe/stargates/{stargate_id}/][%d] getUniverseStargatesStargateIdNotFound  %+v", 404, o.Payload)
}
func (o *GetUniverseStargatesStargateIDNotFound) GetPayload() *GetUniverseStargatesStargateIDNotFoundBody {
	return o.Payload
}

func (o *GetUniverseStargatesStargateIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetUniverseStargatesStargateIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStargatesStargateIDEnhanceYourCalm creates a GetUniverseStargatesStargateIDEnhanceYourCalm with default headers values
func NewGetUniverseStargatesStargateIDEnhanceYourCalm() *GetUniverseStargatesStargateIDEnhanceYourCalm {
	return &GetUniverseStargatesStargateIDEnhanceYourCalm{}
}

/* GetUniverseStargatesStargateIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetUniverseStargatesStargateIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetUniverseStargatesStargateIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /universe/stargates/{stargate_id}/][%d] getUniverseStargatesStargateIdEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetUniverseStargatesStargateIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetUniverseStargatesStargateIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStargatesStargateIDInternalServerError creates a GetUniverseStargatesStargateIDInternalServerError with default headers values
func NewGetUniverseStargatesStargateIDInternalServerError() *GetUniverseStargatesStargateIDInternalServerError {
	return &GetUniverseStargatesStargateIDInternalServerError{}
}

/* GetUniverseStargatesStargateIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUniverseStargatesStargateIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUniverseStargatesStargateIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/stargates/{stargate_id}/][%d] getUniverseStargatesStargateIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetUniverseStargatesStargateIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUniverseStargatesStargateIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStargatesStargateIDServiceUnavailable creates a GetUniverseStargatesStargateIDServiceUnavailable with default headers values
func NewGetUniverseStargatesStargateIDServiceUnavailable() *GetUniverseStargatesStargateIDServiceUnavailable {
	return &GetUniverseStargatesStargateIDServiceUnavailable{}
}

/* GetUniverseStargatesStargateIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetUniverseStargatesStargateIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetUniverseStargatesStargateIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /universe/stargates/{stargate_id}/][%d] getUniverseStargatesStargateIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetUniverseStargatesStargateIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetUniverseStargatesStargateIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStargatesStargateIDGatewayTimeout creates a GetUniverseStargatesStargateIDGatewayTimeout with default headers values
func NewGetUniverseStargatesStargateIDGatewayTimeout() *GetUniverseStargatesStargateIDGatewayTimeout {
	return &GetUniverseStargatesStargateIDGatewayTimeout{}
}

/* GetUniverseStargatesStargateIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetUniverseStargatesStargateIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetUniverseStargatesStargateIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /universe/stargates/{stargate_id}/][%d] getUniverseStargatesStargateIdGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetUniverseStargatesStargateIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetUniverseStargatesStargateIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetUniverseStargatesStargateIDNotFoundBody get_universe_stargates_stargate_id_not_found
//
// Not found
swagger:model GetUniverseStargatesStargateIDNotFoundBody
*/
type GetUniverseStargatesStargateIDNotFoundBody struct {

	// get_universe_stargates_stargate_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get universe stargates stargate ID not found body
func (o *GetUniverseStargatesStargateIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get universe stargates stargate ID not found body based on context it is used
func (o *GetUniverseStargatesStargateIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseStargatesStargateIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseStargatesStargateIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseStargatesStargateIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUniverseStargatesStargateIDOKBody get_universe_stargates_stargate_id_ok
//
// 200 ok object
swagger:model GetUniverseStargatesStargateIDOKBody
*/
type GetUniverseStargatesStargateIDOKBody struct {

	// destination
	// Required: true
	Destination *GetUniverseStargatesStargateIDOKBodyDestination `json:"destination"`

	// get_universe_stargates_stargate_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// position
	// Required: true
	Position *GetUniverseStargatesStargateIDOKBodyPosition `json:"position"`

	// get_universe_stargates_stargate_id_stargate_id
	//
	// stargate_id integer
	// Required: true
	StargateID *int32 `json:"stargate_id"`

	// get_universe_stargates_stargate_id_system_id
	//
	// The solar system this stargate is in
	// Required: true
	SystemID *int32 `json:"system_id"`

	// get_universe_stargates_stargate_id_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get universe stargates stargate ID o k body
func (o *GetUniverseStargatesStargateIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStargateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSystemID(formats); err != nil {
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

func (o *GetUniverseStargatesStargateIDOKBody) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStargatesStargateIdOK"+"."+"destination", "body", o.Destination); err != nil {
		return err
	}

	if o.Destination != nil {
		if err := o.Destination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUniverseStargatesStargateIdOK" + "." + "destination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getUniverseStargatesStargateIdOK" + "." + "destination")
			}
			return err
		}
	}

	return nil
}

func (o *GetUniverseStargatesStargateIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStargatesStargateIdOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStargatesStargateIDOKBody) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStargatesStargateIdOK"+"."+"position", "body", o.Position); err != nil {
		return err
	}

	if o.Position != nil {
		if err := o.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUniverseStargatesStargateIdOK" + "." + "position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getUniverseStargatesStargateIdOK" + "." + "position")
			}
			return err
		}
	}

	return nil
}

func (o *GetUniverseStargatesStargateIDOKBody) validateStargateID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStargatesStargateIdOK"+"."+"stargate_id", "body", o.StargateID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStargatesStargateIDOKBody) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStargatesStargateIdOK"+"."+"system_id", "body", o.SystemID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStargatesStargateIDOKBody) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStargatesStargateIdOK"+"."+"type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get universe stargates stargate ID o k body based on the context it is used
func (o *GetUniverseStargatesStargateIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDestination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseStargatesStargateIDOKBody) contextValidateDestination(ctx context.Context, formats strfmt.Registry) error {

	if o.Destination != nil {
		if err := o.Destination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUniverseStargatesStargateIdOK" + "." + "destination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getUniverseStargatesStargateIdOK" + "." + "destination")
			}
			return err
		}
	}

	return nil
}

func (o *GetUniverseStargatesStargateIDOKBody) contextValidatePosition(ctx context.Context, formats strfmt.Registry) error {

	if o.Position != nil {
		if err := o.Position.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUniverseStargatesStargateIdOK" + "." + "position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getUniverseStargatesStargateIdOK" + "." + "position")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseStargatesStargateIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseStargatesStargateIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseStargatesStargateIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUniverseStargatesStargateIDOKBodyDestination get_universe_stargates_stargate_id_destination
//
// destination object
swagger:model GetUniverseStargatesStargateIDOKBodyDestination
*/
type GetUniverseStargatesStargateIDOKBodyDestination struct {

	// get_universe_stargates_stargate_id_destination_stargate_id
	//
	// The stargate this stargate connects to
	// Required: true
	StargateID *int32 `json:"stargate_id"`

	// get_universe_stargates_stargate_id_destination_system_id
	//
	// The solar system this stargate connects to
	// Required: true
	SystemID *int32 `json:"system_id"`
}

// Validate validates this get universe stargates stargate ID o k body destination
func (o *GetUniverseStargatesStargateIDOKBodyDestination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStargateID(formats); err != nil {
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

func (o *GetUniverseStargatesStargateIDOKBodyDestination) validateStargateID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStargatesStargateIdOK"+"."+"destination"+"."+"stargate_id", "body", o.StargateID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStargatesStargateIDOKBodyDestination) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStargatesStargateIdOK"+"."+"destination"+"."+"system_id", "body", o.SystemID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get universe stargates stargate ID o k body destination based on context it is used
func (o *GetUniverseStargatesStargateIDOKBodyDestination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseStargatesStargateIDOKBodyDestination) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseStargatesStargateIDOKBodyDestination) UnmarshalBinary(b []byte) error {
	var res GetUniverseStargatesStargateIDOKBodyDestination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUniverseStargatesStargateIDOKBodyPosition get_universe_stargates_stargate_id_position
//
// position object
swagger:model GetUniverseStargatesStargateIDOKBodyPosition
*/
type GetUniverseStargatesStargateIDOKBodyPosition struct {

	// get_universe_stargates_stargate_id_x
	//
	// x number
	// Required: true
	X *float64 `json:"x"`

	// get_universe_stargates_stargate_id_y
	//
	// y number
	// Required: true
	Y *float64 `json:"y"`

	// get_universe_stargates_stargate_id_z
	//
	// z number
	// Required: true
	Z *float64 `json:"z"`
}

// Validate validates this get universe stargates stargate ID o k body position
func (o *GetUniverseStargatesStargateIDOKBodyPosition) Validate(formats strfmt.Registry) error {
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

func (o *GetUniverseStargatesStargateIDOKBodyPosition) validateX(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStargatesStargateIdOK"+"."+"position"+"."+"x", "body", o.X); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStargatesStargateIDOKBodyPosition) validateY(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStargatesStargateIdOK"+"."+"position"+"."+"y", "body", o.Y); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStargatesStargateIDOKBodyPosition) validateZ(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStargatesStargateIdOK"+"."+"position"+"."+"z", "body", o.Z); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get universe stargates stargate ID o k body position based on context it is used
func (o *GetUniverseStargatesStargateIDOKBodyPosition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseStargatesStargateIDOKBodyPosition) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseStargatesStargateIDOKBodyPosition) UnmarshalBinary(b []byte) error {
	var res GetUniverseStargatesStargateIDOKBodyPosition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
