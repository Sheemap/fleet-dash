// Code generated by go-swagger; DO NOT EDIT.

package killmails

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

// GetCharactersCharacterIDKillmailsRecentReader is a Reader for the GetCharactersCharacterIDKillmailsRecent structure.
type GetCharactersCharacterIDKillmailsRecentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDKillmailsRecentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDKillmailsRecentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDKillmailsRecentNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDKillmailsRecentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDKillmailsRecentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDKillmailsRecentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDKillmailsRecentEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDKillmailsRecentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDKillmailsRecentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDKillmailsRecentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersCharacterIDKillmailsRecentOK creates a GetCharactersCharacterIDKillmailsRecentOK with default headers values
func NewGetCharactersCharacterIDKillmailsRecentOK() *GetCharactersCharacterIDKillmailsRecentOK {
	var (
		// initialize headers with default values
		xPagesDefault = int32(1)
	)

	return &GetCharactersCharacterIDKillmailsRecentOK{

		XPages: xPagesDefault,
	}
}

/* GetCharactersCharacterIDKillmailsRecentOK describes a response with status code 200, with default header values.

A list of killmail IDs and hashes
*/
type GetCharactersCharacterIDKillmailsRecentOK struct {

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

	Payload []*GetCharactersCharacterIDKillmailsRecentOKBodyItems0
}

func (o *GetCharactersCharacterIDKillmailsRecentOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/killmails/recent/][%d] getCharactersCharacterIdKillmailsRecentOK  %+v", 200, o.Payload)
}
func (o *GetCharactersCharacterIDKillmailsRecentOK) GetPayload() []*GetCharactersCharacterIDKillmailsRecentOKBodyItems0 {
	return o.Payload
}

func (o *GetCharactersCharacterIDKillmailsRecentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDKillmailsRecentNotModified creates a GetCharactersCharacterIDKillmailsRecentNotModified with default headers values
func NewGetCharactersCharacterIDKillmailsRecentNotModified() *GetCharactersCharacterIDKillmailsRecentNotModified {
	return &GetCharactersCharacterIDKillmailsRecentNotModified{}
}

/* GetCharactersCharacterIDKillmailsRecentNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDKillmailsRecentNotModified struct {

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

func (o *GetCharactersCharacterIDKillmailsRecentNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/killmails/recent/][%d] getCharactersCharacterIdKillmailsRecentNotModified ", 304)
}

func (o *GetCharactersCharacterIDKillmailsRecentNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDKillmailsRecentBadRequest creates a GetCharactersCharacterIDKillmailsRecentBadRequest with default headers values
func NewGetCharactersCharacterIDKillmailsRecentBadRequest() *GetCharactersCharacterIDKillmailsRecentBadRequest {
	return &GetCharactersCharacterIDKillmailsRecentBadRequest{}
}

/* GetCharactersCharacterIDKillmailsRecentBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDKillmailsRecentBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDKillmailsRecentBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/killmails/recent/][%d] getCharactersCharacterIdKillmailsRecentBadRequest  %+v", 400, o.Payload)
}
func (o *GetCharactersCharacterIDKillmailsRecentBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDKillmailsRecentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDKillmailsRecentUnauthorized creates a GetCharactersCharacterIDKillmailsRecentUnauthorized with default headers values
func NewGetCharactersCharacterIDKillmailsRecentUnauthorized() *GetCharactersCharacterIDKillmailsRecentUnauthorized {
	return &GetCharactersCharacterIDKillmailsRecentUnauthorized{}
}

/* GetCharactersCharacterIDKillmailsRecentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDKillmailsRecentUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDKillmailsRecentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/killmails/recent/][%d] getCharactersCharacterIdKillmailsRecentUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCharactersCharacterIDKillmailsRecentUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDKillmailsRecentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDKillmailsRecentForbidden creates a GetCharactersCharacterIDKillmailsRecentForbidden with default headers values
func NewGetCharactersCharacterIDKillmailsRecentForbidden() *GetCharactersCharacterIDKillmailsRecentForbidden {
	return &GetCharactersCharacterIDKillmailsRecentForbidden{}
}

/* GetCharactersCharacterIDKillmailsRecentForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDKillmailsRecentForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDKillmailsRecentForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/killmails/recent/][%d] getCharactersCharacterIdKillmailsRecentForbidden  %+v", 403, o.Payload)
}
func (o *GetCharactersCharacterIDKillmailsRecentForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDKillmailsRecentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDKillmailsRecentEnhanceYourCalm creates a GetCharactersCharacterIDKillmailsRecentEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDKillmailsRecentEnhanceYourCalm() *GetCharactersCharacterIDKillmailsRecentEnhanceYourCalm {
	return &GetCharactersCharacterIDKillmailsRecentEnhanceYourCalm{}
}

/* GetCharactersCharacterIDKillmailsRecentEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDKillmailsRecentEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDKillmailsRecentEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/killmails/recent/][%d] getCharactersCharacterIdKillmailsRecentEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCharactersCharacterIDKillmailsRecentEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDKillmailsRecentEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDKillmailsRecentInternalServerError creates a GetCharactersCharacterIDKillmailsRecentInternalServerError with default headers values
func NewGetCharactersCharacterIDKillmailsRecentInternalServerError() *GetCharactersCharacterIDKillmailsRecentInternalServerError {
	return &GetCharactersCharacterIDKillmailsRecentInternalServerError{}
}

/* GetCharactersCharacterIDKillmailsRecentInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDKillmailsRecentInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDKillmailsRecentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/killmails/recent/][%d] getCharactersCharacterIdKillmailsRecentInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCharactersCharacterIDKillmailsRecentInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDKillmailsRecentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDKillmailsRecentServiceUnavailable creates a GetCharactersCharacterIDKillmailsRecentServiceUnavailable with default headers values
func NewGetCharactersCharacterIDKillmailsRecentServiceUnavailable() *GetCharactersCharacterIDKillmailsRecentServiceUnavailable {
	return &GetCharactersCharacterIDKillmailsRecentServiceUnavailable{}
}

/* GetCharactersCharacterIDKillmailsRecentServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDKillmailsRecentServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDKillmailsRecentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/killmails/recent/][%d] getCharactersCharacterIdKillmailsRecentServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCharactersCharacterIDKillmailsRecentServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDKillmailsRecentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDKillmailsRecentGatewayTimeout creates a GetCharactersCharacterIDKillmailsRecentGatewayTimeout with default headers values
func NewGetCharactersCharacterIDKillmailsRecentGatewayTimeout() *GetCharactersCharacterIDKillmailsRecentGatewayTimeout {
	return &GetCharactersCharacterIDKillmailsRecentGatewayTimeout{}
}

/* GetCharactersCharacterIDKillmailsRecentGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDKillmailsRecentGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDKillmailsRecentGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/killmails/recent/][%d] getCharactersCharacterIdKillmailsRecentGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCharactersCharacterIDKillmailsRecentGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDKillmailsRecentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDKillmailsRecentOKBodyItems0 get_characters_character_id_killmails_recent_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDKillmailsRecentOKBodyItems0
*/
type GetCharactersCharacterIDKillmailsRecentOKBodyItems0 struct {

	// get_characters_character_id_killmails_recent_killmail_hash
	//
	// A hash of this killmail
	// Required: true
	KillmailHash *string `json:"killmail_hash"`

	// get_characters_character_id_killmails_recent_killmail_id
	//
	// ID of this killmail
	// Required: true
	KillmailID *int32 `json:"killmail_id"`
}

// Validate validates this get characters character ID killmails recent o k body items0
func (o *GetCharactersCharacterIDKillmailsRecentOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateKillmailHash(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateKillmailID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDKillmailsRecentOKBodyItems0) validateKillmailHash(formats strfmt.Registry) error {

	if err := validate.Required("killmail_hash", "body", o.KillmailHash); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDKillmailsRecentOKBodyItems0) validateKillmailID(formats strfmt.Registry) error {

	if err := validate.Required("killmail_id", "body", o.KillmailID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get characters character ID killmails recent o k body items0 based on context it is used
func (o *GetCharactersCharacterIDKillmailsRecentOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDKillmailsRecentOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDKillmailsRecentOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDKillmailsRecentOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
