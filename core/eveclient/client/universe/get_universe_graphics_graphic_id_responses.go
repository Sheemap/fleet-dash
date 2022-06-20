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

// GetUniverseGraphicsGraphicIDReader is a Reader for the GetUniverseGraphicsGraphicID structure.
type GetUniverseGraphicsGraphicIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseGraphicsGraphicIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUniverseGraphicsGraphicIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetUniverseGraphicsGraphicIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetUniverseGraphicsGraphicIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUniverseGraphicsGraphicIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetUniverseGraphicsGraphicIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUniverseGraphicsGraphicIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUniverseGraphicsGraphicIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUniverseGraphicsGraphicIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUniverseGraphicsGraphicIDOK creates a GetUniverseGraphicsGraphicIDOK with default headers values
func NewGetUniverseGraphicsGraphicIDOK() *GetUniverseGraphicsGraphicIDOK {
	return &GetUniverseGraphicsGraphicIDOK{}
}

/* GetUniverseGraphicsGraphicIDOK describes a response with status code 200, with default header values.

Information about a graphic
*/
type GetUniverseGraphicsGraphicIDOK struct {

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

	Payload *GetUniverseGraphicsGraphicIDOKBody
}

func (o *GetUniverseGraphicsGraphicIDOK) Error() string {
	return fmt.Sprintf("[GET /universe/graphics/{graphic_id}/][%d] getUniverseGraphicsGraphicIdOK  %+v", 200, o.Payload)
}
func (o *GetUniverseGraphicsGraphicIDOK) GetPayload() *GetUniverseGraphicsGraphicIDOKBody {
	return o.Payload
}

func (o *GetUniverseGraphicsGraphicIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetUniverseGraphicsGraphicIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseGraphicsGraphicIDNotModified creates a GetUniverseGraphicsGraphicIDNotModified with default headers values
func NewGetUniverseGraphicsGraphicIDNotModified() *GetUniverseGraphicsGraphicIDNotModified {
	return &GetUniverseGraphicsGraphicIDNotModified{}
}

/* GetUniverseGraphicsGraphicIDNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetUniverseGraphicsGraphicIDNotModified struct {

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

func (o *GetUniverseGraphicsGraphicIDNotModified) Error() string {
	return fmt.Sprintf("[GET /universe/graphics/{graphic_id}/][%d] getUniverseGraphicsGraphicIdNotModified ", 304)
}

func (o *GetUniverseGraphicsGraphicIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseGraphicsGraphicIDBadRequest creates a GetUniverseGraphicsGraphicIDBadRequest with default headers values
func NewGetUniverseGraphicsGraphicIDBadRequest() *GetUniverseGraphicsGraphicIDBadRequest {
	return &GetUniverseGraphicsGraphicIDBadRequest{}
}

/* GetUniverseGraphicsGraphicIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUniverseGraphicsGraphicIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetUniverseGraphicsGraphicIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /universe/graphics/{graphic_id}/][%d] getUniverseGraphicsGraphicIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetUniverseGraphicsGraphicIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetUniverseGraphicsGraphicIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseGraphicsGraphicIDNotFound creates a GetUniverseGraphicsGraphicIDNotFound with default headers values
func NewGetUniverseGraphicsGraphicIDNotFound() *GetUniverseGraphicsGraphicIDNotFound {
	return &GetUniverseGraphicsGraphicIDNotFound{}
}

/* GetUniverseGraphicsGraphicIDNotFound describes a response with status code 404, with default header values.

Graphic not found
*/
type GetUniverseGraphicsGraphicIDNotFound struct {
	Payload *GetUniverseGraphicsGraphicIDNotFoundBody
}

func (o *GetUniverseGraphicsGraphicIDNotFound) Error() string {
	return fmt.Sprintf("[GET /universe/graphics/{graphic_id}/][%d] getUniverseGraphicsGraphicIdNotFound  %+v", 404, o.Payload)
}
func (o *GetUniverseGraphicsGraphicIDNotFound) GetPayload() *GetUniverseGraphicsGraphicIDNotFoundBody {
	return o.Payload
}

func (o *GetUniverseGraphicsGraphicIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetUniverseGraphicsGraphicIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseGraphicsGraphicIDEnhanceYourCalm creates a GetUniverseGraphicsGraphicIDEnhanceYourCalm with default headers values
func NewGetUniverseGraphicsGraphicIDEnhanceYourCalm() *GetUniverseGraphicsGraphicIDEnhanceYourCalm {
	return &GetUniverseGraphicsGraphicIDEnhanceYourCalm{}
}

/* GetUniverseGraphicsGraphicIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetUniverseGraphicsGraphicIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetUniverseGraphicsGraphicIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /universe/graphics/{graphic_id}/][%d] getUniverseGraphicsGraphicIdEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetUniverseGraphicsGraphicIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetUniverseGraphicsGraphicIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseGraphicsGraphicIDInternalServerError creates a GetUniverseGraphicsGraphicIDInternalServerError with default headers values
func NewGetUniverseGraphicsGraphicIDInternalServerError() *GetUniverseGraphicsGraphicIDInternalServerError {
	return &GetUniverseGraphicsGraphicIDInternalServerError{}
}

/* GetUniverseGraphicsGraphicIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUniverseGraphicsGraphicIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUniverseGraphicsGraphicIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/graphics/{graphic_id}/][%d] getUniverseGraphicsGraphicIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetUniverseGraphicsGraphicIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUniverseGraphicsGraphicIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseGraphicsGraphicIDServiceUnavailable creates a GetUniverseGraphicsGraphicIDServiceUnavailable with default headers values
func NewGetUniverseGraphicsGraphicIDServiceUnavailable() *GetUniverseGraphicsGraphicIDServiceUnavailable {
	return &GetUniverseGraphicsGraphicIDServiceUnavailable{}
}

/* GetUniverseGraphicsGraphicIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetUniverseGraphicsGraphicIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetUniverseGraphicsGraphicIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /universe/graphics/{graphic_id}/][%d] getUniverseGraphicsGraphicIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetUniverseGraphicsGraphicIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetUniverseGraphicsGraphicIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseGraphicsGraphicIDGatewayTimeout creates a GetUniverseGraphicsGraphicIDGatewayTimeout with default headers values
func NewGetUniverseGraphicsGraphicIDGatewayTimeout() *GetUniverseGraphicsGraphicIDGatewayTimeout {
	return &GetUniverseGraphicsGraphicIDGatewayTimeout{}
}

/* GetUniverseGraphicsGraphicIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetUniverseGraphicsGraphicIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetUniverseGraphicsGraphicIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /universe/graphics/{graphic_id}/][%d] getUniverseGraphicsGraphicIdGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetUniverseGraphicsGraphicIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetUniverseGraphicsGraphicIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetUniverseGraphicsGraphicIDNotFoundBody get_universe_graphics_graphic_id_not_found
//
// Not found
swagger:model GetUniverseGraphicsGraphicIDNotFoundBody
*/
type GetUniverseGraphicsGraphicIDNotFoundBody struct {

	// get_universe_graphics_graphic_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get universe graphics graphic ID not found body
func (o *GetUniverseGraphicsGraphicIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get universe graphics graphic ID not found body based on context it is used
func (o *GetUniverseGraphicsGraphicIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseGraphicsGraphicIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseGraphicsGraphicIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseGraphicsGraphicIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUniverseGraphicsGraphicIDOKBody get_universe_graphics_graphic_id_ok
//
// 200 ok object
swagger:model GetUniverseGraphicsGraphicIDOKBody
*/
type GetUniverseGraphicsGraphicIDOKBody struct {

	// get_universe_graphics_graphic_id_collision_file
	//
	// collision_file string
	CollisionFile string `json:"collision_file,omitempty"`

	// get_universe_graphics_graphic_id_graphic_file
	//
	// graphic_file string
	GraphicFile string `json:"graphic_file,omitempty"`

	// get_universe_graphics_graphic_id_graphic_id
	//
	// graphic_id integer
	// Required: true
	GraphicID *int32 `json:"graphic_id"`

	// get_universe_graphics_graphic_id_icon_folder
	//
	// icon_folder string
	IconFolder string `json:"icon_folder,omitempty"`

	// get_universe_graphics_graphic_id_sof_dna
	//
	// sof_dna string
	SofDna string `json:"sof_dna,omitempty"`

	// get_universe_graphics_graphic_id_sof_fation_name
	//
	// sof_fation_name string
	SofFationName string `json:"sof_fation_name,omitempty"`

	// get_universe_graphics_graphic_id_sof_hull_name
	//
	// sof_hull_name string
	SofHullName string `json:"sof_hull_name,omitempty"`

	// get_universe_graphics_graphic_id_sof_race_name
	//
	// sof_race_name string
	SofRaceName string `json:"sof_race_name,omitempty"`
}

// Validate validates this get universe graphics graphic ID o k body
func (o *GetUniverseGraphicsGraphicIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGraphicID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseGraphicsGraphicIDOKBody) validateGraphicID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseGraphicsGraphicIdOK"+"."+"graphic_id", "body", o.GraphicID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get universe graphics graphic ID o k body based on context it is used
func (o *GetUniverseGraphicsGraphicIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseGraphicsGraphicIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseGraphicsGraphicIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseGraphicsGraphicIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}