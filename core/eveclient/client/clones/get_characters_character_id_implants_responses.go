// Code generated by go-swagger; DO NOT EDIT.

package clones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"fleet-dash-core/eveclient/models"
)

// GetCharactersCharacterIDImplantsReader is a Reader for the GetCharactersCharacterIDImplants structure.
type GetCharactersCharacterIDImplantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDImplantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDImplantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDImplantsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDImplantsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDImplantsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDImplantsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDImplantsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDImplantsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDImplantsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDImplantsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersCharacterIDImplantsOK creates a GetCharactersCharacterIDImplantsOK with default headers values
func NewGetCharactersCharacterIDImplantsOK() *GetCharactersCharacterIDImplantsOK {
	return &GetCharactersCharacterIDImplantsOK{}
}

/* GetCharactersCharacterIDImplantsOK describes a response with status code 200, with default header values.

A list of implant type ids
*/
type GetCharactersCharacterIDImplantsOK struct {

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

	Payload []int32
}

func (o *GetCharactersCharacterIDImplantsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/implants/][%d] getCharactersCharacterIdImplantsOK  %+v", 200, o.Payload)
}
func (o *GetCharactersCharacterIDImplantsOK) GetPayload() []int32 {
	return o.Payload
}

func (o *GetCharactersCharacterIDImplantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDImplantsNotModified creates a GetCharactersCharacterIDImplantsNotModified with default headers values
func NewGetCharactersCharacterIDImplantsNotModified() *GetCharactersCharacterIDImplantsNotModified {
	return &GetCharactersCharacterIDImplantsNotModified{}
}

/* GetCharactersCharacterIDImplantsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDImplantsNotModified struct {

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

func (o *GetCharactersCharacterIDImplantsNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/implants/][%d] getCharactersCharacterIdImplantsNotModified ", 304)
}

func (o *GetCharactersCharacterIDImplantsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDImplantsBadRequest creates a GetCharactersCharacterIDImplantsBadRequest with default headers values
func NewGetCharactersCharacterIDImplantsBadRequest() *GetCharactersCharacterIDImplantsBadRequest {
	return &GetCharactersCharacterIDImplantsBadRequest{}
}

/* GetCharactersCharacterIDImplantsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDImplantsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDImplantsBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/implants/][%d] getCharactersCharacterIdImplantsBadRequest  %+v", 400, o.Payload)
}
func (o *GetCharactersCharacterIDImplantsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDImplantsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDImplantsUnauthorized creates a GetCharactersCharacterIDImplantsUnauthorized with default headers values
func NewGetCharactersCharacterIDImplantsUnauthorized() *GetCharactersCharacterIDImplantsUnauthorized {
	return &GetCharactersCharacterIDImplantsUnauthorized{}
}

/* GetCharactersCharacterIDImplantsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDImplantsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDImplantsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/implants/][%d] getCharactersCharacterIdImplantsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCharactersCharacterIDImplantsUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDImplantsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDImplantsForbidden creates a GetCharactersCharacterIDImplantsForbidden with default headers values
func NewGetCharactersCharacterIDImplantsForbidden() *GetCharactersCharacterIDImplantsForbidden {
	return &GetCharactersCharacterIDImplantsForbidden{}
}

/* GetCharactersCharacterIDImplantsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDImplantsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDImplantsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/implants/][%d] getCharactersCharacterIdImplantsForbidden  %+v", 403, o.Payload)
}
func (o *GetCharactersCharacterIDImplantsForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDImplantsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDImplantsEnhanceYourCalm creates a GetCharactersCharacterIDImplantsEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDImplantsEnhanceYourCalm() *GetCharactersCharacterIDImplantsEnhanceYourCalm {
	return &GetCharactersCharacterIDImplantsEnhanceYourCalm{}
}

/* GetCharactersCharacterIDImplantsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDImplantsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDImplantsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/implants/][%d] getCharactersCharacterIdImplantsEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCharactersCharacterIDImplantsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDImplantsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDImplantsInternalServerError creates a GetCharactersCharacterIDImplantsInternalServerError with default headers values
func NewGetCharactersCharacterIDImplantsInternalServerError() *GetCharactersCharacterIDImplantsInternalServerError {
	return &GetCharactersCharacterIDImplantsInternalServerError{}
}

/* GetCharactersCharacterIDImplantsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDImplantsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDImplantsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/implants/][%d] getCharactersCharacterIdImplantsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCharactersCharacterIDImplantsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDImplantsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDImplantsServiceUnavailable creates a GetCharactersCharacterIDImplantsServiceUnavailable with default headers values
func NewGetCharactersCharacterIDImplantsServiceUnavailable() *GetCharactersCharacterIDImplantsServiceUnavailable {
	return &GetCharactersCharacterIDImplantsServiceUnavailable{}
}

/* GetCharactersCharacterIDImplantsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDImplantsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDImplantsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/implants/][%d] getCharactersCharacterIdImplantsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCharactersCharacterIDImplantsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDImplantsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDImplantsGatewayTimeout creates a GetCharactersCharacterIDImplantsGatewayTimeout with default headers values
func NewGetCharactersCharacterIDImplantsGatewayTimeout() *GetCharactersCharacterIDImplantsGatewayTimeout {
	return &GetCharactersCharacterIDImplantsGatewayTimeout{}
}

/* GetCharactersCharacterIDImplantsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDImplantsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDImplantsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/implants/][%d] getCharactersCharacterIdImplantsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCharactersCharacterIDImplantsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDImplantsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
