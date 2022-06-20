// Code generated by go-swagger; DO NOT EDIT.

package corporation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"fleet-dash-core/eveclient/models"
)

// GetCorporationsCorporationIDIconsReader is a Reader for the GetCorporationsCorporationIDIcons structure.
type GetCorporationsCorporationIDIconsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDIconsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationsCorporationIDIconsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationsCorporationIDIconsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationsCorporationIDIconsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCorporationsCorporationIDIconsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationsCorporationIDIconsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationsCorporationIDIconsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationsCorporationIDIconsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationsCorporationIDIconsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDIconsOK creates a GetCorporationsCorporationIDIconsOK with default headers values
func NewGetCorporationsCorporationIDIconsOK() *GetCorporationsCorporationIDIconsOK {
	return &GetCorporationsCorporationIDIconsOK{}
}

/* GetCorporationsCorporationIDIconsOK describes a response with status code 200, with default header values.

Urls for icons for the given corporation id and server
*/
type GetCorporationsCorporationIDIconsOK struct {

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

	Payload *GetCorporationsCorporationIDIconsOKBody
}

func (o *GetCorporationsCorporationIDIconsOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/icons/][%d] getCorporationsCorporationIdIconsOK  %+v", 200, o.Payload)
}
func (o *GetCorporationsCorporationIDIconsOK) GetPayload() *GetCorporationsCorporationIDIconsOKBody {
	return o.Payload
}

func (o *GetCorporationsCorporationIDIconsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetCorporationsCorporationIDIconsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDIconsNotModified creates a GetCorporationsCorporationIDIconsNotModified with default headers values
func NewGetCorporationsCorporationIDIconsNotModified() *GetCorporationsCorporationIDIconsNotModified {
	return &GetCorporationsCorporationIDIconsNotModified{}
}

/* GetCorporationsCorporationIDIconsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationsCorporationIDIconsNotModified struct {

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

func (o *GetCorporationsCorporationIDIconsNotModified) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/icons/][%d] getCorporationsCorporationIdIconsNotModified ", 304)
}

func (o *GetCorporationsCorporationIDIconsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDIconsBadRequest creates a GetCorporationsCorporationIDIconsBadRequest with default headers values
func NewGetCorporationsCorporationIDIconsBadRequest() *GetCorporationsCorporationIDIconsBadRequest {
	return &GetCorporationsCorporationIDIconsBadRequest{}
}

/* GetCorporationsCorporationIDIconsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationsCorporationIDIconsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDIconsBadRequest) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/icons/][%d] getCorporationsCorporationIdIconsBadRequest  %+v", 400, o.Payload)
}
func (o *GetCorporationsCorporationIDIconsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationsCorporationIDIconsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDIconsNotFound creates a GetCorporationsCorporationIDIconsNotFound with default headers values
func NewGetCorporationsCorporationIDIconsNotFound() *GetCorporationsCorporationIDIconsNotFound {
	return &GetCorporationsCorporationIDIconsNotFound{}
}

/* GetCorporationsCorporationIDIconsNotFound describes a response with status code 404, with default header values.

No image server for this datasource
*/
type GetCorporationsCorporationIDIconsNotFound struct {
	Payload *GetCorporationsCorporationIDIconsNotFoundBody
}

func (o *GetCorporationsCorporationIDIconsNotFound) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/icons/][%d] getCorporationsCorporationIdIconsNotFound  %+v", 404, o.Payload)
}
func (o *GetCorporationsCorporationIDIconsNotFound) GetPayload() *GetCorporationsCorporationIDIconsNotFoundBody {
	return o.Payload
}

func (o *GetCorporationsCorporationIDIconsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCorporationsCorporationIDIconsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDIconsEnhanceYourCalm creates a GetCorporationsCorporationIDIconsEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDIconsEnhanceYourCalm() *GetCorporationsCorporationIDIconsEnhanceYourCalm {
	return &GetCorporationsCorporationIDIconsEnhanceYourCalm{}
}

/* GetCorporationsCorporationIDIconsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationsCorporationIDIconsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDIconsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/icons/][%d] getCorporationsCorporationIdIconsEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCorporationsCorporationIDIconsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationsCorporationIDIconsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDIconsInternalServerError creates a GetCorporationsCorporationIDIconsInternalServerError with default headers values
func NewGetCorporationsCorporationIDIconsInternalServerError() *GetCorporationsCorporationIDIconsInternalServerError {
	return &GetCorporationsCorporationIDIconsInternalServerError{}
}

/* GetCorporationsCorporationIDIconsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationsCorporationIDIconsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDIconsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/icons/][%d] getCorporationsCorporationIdIconsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCorporationsCorporationIDIconsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationsCorporationIDIconsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDIconsServiceUnavailable creates a GetCorporationsCorporationIDIconsServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDIconsServiceUnavailable() *GetCorporationsCorporationIDIconsServiceUnavailable {
	return &GetCorporationsCorporationIDIconsServiceUnavailable{}
}

/* GetCorporationsCorporationIDIconsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDIconsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDIconsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/icons/][%d] getCorporationsCorporationIdIconsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCorporationsCorporationIDIconsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationsCorporationIDIconsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDIconsGatewayTimeout creates a GetCorporationsCorporationIDIconsGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDIconsGatewayTimeout() *GetCorporationsCorporationIDIconsGatewayTimeout {
	return &GetCorporationsCorporationIDIconsGatewayTimeout{}
}

/* GetCorporationsCorporationIDIconsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDIconsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDIconsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/icons/][%d] getCorporationsCorporationIdIconsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCorporationsCorporationIDIconsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationsCorporationIDIconsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDIconsNotFoundBody get_corporations_corporation_id_icons_not_found
//
// No image server for this datasource
swagger:model GetCorporationsCorporationIDIconsNotFoundBody
*/
type GetCorporationsCorporationIDIconsNotFoundBody struct {

	// get_corporations_corporation_id_icons_error
	//
	// error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get corporations corporation ID icons not found body
func (o *GetCorporationsCorporationIDIconsNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get corporations corporation ID icons not found body based on context it is used
func (o *GetCorporationsCorporationIDIconsNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDIconsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDIconsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDIconsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetCorporationsCorporationIDIconsOKBody get_corporations_corporation_id_icons_ok
//
// 200 ok object
swagger:model GetCorporationsCorporationIDIconsOKBody
*/
type GetCorporationsCorporationIDIconsOKBody struct {

	// get_corporations_corporation_id_icons_px128x128
	//
	// px128x128 string
	Px128x128 string `json:"px128x128,omitempty"`

	// get_corporations_corporation_id_icons_px256x256
	//
	// px256x256 string
	Px256x256 string `json:"px256x256,omitempty"`

	// get_corporations_corporation_id_icons_px64x64
	//
	// px64x64 string
	Px64x64 string `json:"px64x64,omitempty"`
}

// Validate validates this get corporations corporation ID icons o k body
func (o *GetCorporationsCorporationIDIconsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get corporations corporation ID icons o k body based on context it is used
func (o *GetCorporationsCorporationIDIconsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDIconsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDIconsOKBody) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDIconsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}