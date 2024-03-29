// Code generated by go-swagger; DO NOT EDIT.

package alliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"fleet-dash-core/eveclient/models"
)

// GetAlliancesReader is a Reader for the GetAlliances structure.
type GetAlliancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlliancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlliancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetAlliancesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetAlliancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetAlliancesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAlliancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAlliancesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAlliancesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAlliancesOK creates a GetAlliancesOK with default headers values
func NewGetAlliancesOK() *GetAlliancesOK {
	return &GetAlliancesOK{}
}

/* GetAlliancesOK describes a response with status code 200, with default header values.

List of Alliance IDs
*/
type GetAlliancesOK struct {

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

	Payload []*int32
}

func (o *GetAlliancesOK) Error() string {
	return fmt.Sprintf("[GET /alliances/][%d] getAlliancesOK  %+v", 200, o.Payload)
}
func (o *GetAlliancesOK) GetPayload() []*int32 {
	return o.Payload
}

func (o *GetAlliancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlliancesNotModified creates a GetAlliancesNotModified with default headers values
func NewGetAlliancesNotModified() *GetAlliancesNotModified {
	return &GetAlliancesNotModified{}
}

/* GetAlliancesNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetAlliancesNotModified struct {

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

func (o *GetAlliancesNotModified) Error() string {
	return fmt.Sprintf("[GET /alliances/][%d] getAlliancesNotModified ", 304)
}

func (o *GetAlliancesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlliancesBadRequest creates a GetAlliancesBadRequest with default headers values
func NewGetAlliancesBadRequest() *GetAlliancesBadRequest {
	return &GetAlliancesBadRequest{}
}

/* GetAlliancesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetAlliancesBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetAlliancesBadRequest) Error() string {
	return fmt.Sprintf("[GET /alliances/][%d] getAlliancesBadRequest  %+v", 400, o.Payload)
}
func (o *GetAlliancesBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetAlliancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesEnhanceYourCalm creates a GetAlliancesEnhanceYourCalm with default headers values
func NewGetAlliancesEnhanceYourCalm() *GetAlliancesEnhanceYourCalm {
	return &GetAlliancesEnhanceYourCalm{}
}

/* GetAlliancesEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetAlliancesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetAlliancesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /alliances/][%d] getAlliancesEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetAlliancesEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetAlliancesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesInternalServerError creates a GetAlliancesInternalServerError with default headers values
func NewGetAlliancesInternalServerError() *GetAlliancesInternalServerError {
	return &GetAlliancesInternalServerError{}
}

/* GetAlliancesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetAlliancesInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetAlliancesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /alliances/][%d] getAlliancesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAlliancesInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetAlliancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesServiceUnavailable creates a GetAlliancesServiceUnavailable with default headers values
func NewGetAlliancesServiceUnavailable() *GetAlliancesServiceUnavailable {
	return &GetAlliancesServiceUnavailable{}
}

/* GetAlliancesServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetAlliancesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetAlliancesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /alliances/][%d] getAlliancesServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetAlliancesServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetAlliancesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesGatewayTimeout creates a GetAlliancesGatewayTimeout with default headers values
func NewGetAlliancesGatewayTimeout() *GetAlliancesGatewayTimeout {
	return &GetAlliancesGatewayTimeout{}
}

/* GetAlliancesGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetAlliancesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetAlliancesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /alliances/][%d] getAlliancesGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetAlliancesGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetAlliancesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
