// Code generated by go-swagger; DO NOT EDIT.

package wars

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

// GetWarsWarIDKillmailsReader is a Reader for the GetWarsWarIDKillmails structure.
type GetWarsWarIDKillmailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWarsWarIDKillmailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWarsWarIDKillmailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetWarsWarIDKillmailsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetWarsWarIDKillmailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetWarsWarIDKillmailsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetWarsWarIDKillmailsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWarsWarIDKillmailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWarsWarIDKillmailsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWarsWarIDKillmailsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWarsWarIDKillmailsOK creates a GetWarsWarIDKillmailsOK with default headers values
func NewGetWarsWarIDKillmailsOK() *GetWarsWarIDKillmailsOK {
	var (
		// initialize headers with default values
		xPagesDefault = int32(1)
	)

	return &GetWarsWarIDKillmailsOK{

		XPages: xPagesDefault,
	}
}

/* GetWarsWarIDKillmailsOK describes a response with status code 200, with default header values.

A list of killmail IDs and hashes
*/
type GetWarsWarIDKillmailsOK struct {

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

	Payload []*GetWarsWarIDKillmailsOKBodyItems0
}

func (o *GetWarsWarIDKillmailsOK) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/killmails/][%d] getWarsWarIdKillmailsOK  %+v", 200, o.Payload)
}
func (o *GetWarsWarIDKillmailsOK) GetPayload() []*GetWarsWarIDKillmailsOKBodyItems0 {
	return o.Payload
}

func (o *GetWarsWarIDKillmailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetWarsWarIDKillmailsNotModified creates a GetWarsWarIDKillmailsNotModified with default headers values
func NewGetWarsWarIDKillmailsNotModified() *GetWarsWarIDKillmailsNotModified {
	return &GetWarsWarIDKillmailsNotModified{}
}

/* GetWarsWarIDKillmailsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetWarsWarIDKillmailsNotModified struct {

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

func (o *GetWarsWarIDKillmailsNotModified) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/killmails/][%d] getWarsWarIdKillmailsNotModified ", 304)
}

func (o *GetWarsWarIDKillmailsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetWarsWarIDKillmailsBadRequest creates a GetWarsWarIDKillmailsBadRequest with default headers values
func NewGetWarsWarIDKillmailsBadRequest() *GetWarsWarIDKillmailsBadRequest {
	return &GetWarsWarIDKillmailsBadRequest{}
}

/* GetWarsWarIDKillmailsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetWarsWarIDKillmailsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetWarsWarIDKillmailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/killmails/][%d] getWarsWarIdKillmailsBadRequest  %+v", 400, o.Payload)
}
func (o *GetWarsWarIDKillmailsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetWarsWarIDKillmailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWarsWarIDKillmailsEnhanceYourCalm creates a GetWarsWarIDKillmailsEnhanceYourCalm with default headers values
func NewGetWarsWarIDKillmailsEnhanceYourCalm() *GetWarsWarIDKillmailsEnhanceYourCalm {
	return &GetWarsWarIDKillmailsEnhanceYourCalm{}
}

/* GetWarsWarIDKillmailsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetWarsWarIDKillmailsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetWarsWarIDKillmailsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/killmails/][%d] getWarsWarIdKillmailsEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetWarsWarIDKillmailsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetWarsWarIDKillmailsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWarsWarIDKillmailsUnprocessableEntity creates a GetWarsWarIDKillmailsUnprocessableEntity with default headers values
func NewGetWarsWarIDKillmailsUnprocessableEntity() *GetWarsWarIDKillmailsUnprocessableEntity {
	return &GetWarsWarIDKillmailsUnprocessableEntity{}
}

/* GetWarsWarIDKillmailsUnprocessableEntity describes a response with status code 422, with default header values.

War not found
*/
type GetWarsWarIDKillmailsUnprocessableEntity struct {
	Payload *GetWarsWarIDKillmailsUnprocessableEntityBody
}

func (o *GetWarsWarIDKillmailsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/killmails/][%d] getWarsWarIdKillmailsUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *GetWarsWarIDKillmailsUnprocessableEntity) GetPayload() *GetWarsWarIDKillmailsUnprocessableEntityBody {
	return o.Payload
}

func (o *GetWarsWarIDKillmailsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWarsWarIDKillmailsUnprocessableEntityBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWarsWarIDKillmailsInternalServerError creates a GetWarsWarIDKillmailsInternalServerError with default headers values
func NewGetWarsWarIDKillmailsInternalServerError() *GetWarsWarIDKillmailsInternalServerError {
	return &GetWarsWarIDKillmailsInternalServerError{}
}

/* GetWarsWarIDKillmailsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetWarsWarIDKillmailsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetWarsWarIDKillmailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/killmails/][%d] getWarsWarIdKillmailsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetWarsWarIDKillmailsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetWarsWarIDKillmailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWarsWarIDKillmailsServiceUnavailable creates a GetWarsWarIDKillmailsServiceUnavailable with default headers values
func NewGetWarsWarIDKillmailsServiceUnavailable() *GetWarsWarIDKillmailsServiceUnavailable {
	return &GetWarsWarIDKillmailsServiceUnavailable{}
}

/* GetWarsWarIDKillmailsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetWarsWarIDKillmailsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetWarsWarIDKillmailsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/killmails/][%d] getWarsWarIdKillmailsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetWarsWarIDKillmailsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetWarsWarIDKillmailsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWarsWarIDKillmailsGatewayTimeout creates a GetWarsWarIDKillmailsGatewayTimeout with default headers values
func NewGetWarsWarIDKillmailsGatewayTimeout() *GetWarsWarIDKillmailsGatewayTimeout {
	return &GetWarsWarIDKillmailsGatewayTimeout{}
}

/* GetWarsWarIDKillmailsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetWarsWarIDKillmailsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetWarsWarIDKillmailsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/killmails/][%d] getWarsWarIdKillmailsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetWarsWarIDKillmailsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetWarsWarIDKillmailsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetWarsWarIDKillmailsOKBodyItems0 get_wars_war_id_killmails_200_ok
//
// 200 ok object
swagger:model GetWarsWarIDKillmailsOKBodyItems0
*/
type GetWarsWarIDKillmailsOKBodyItems0 struct {

	// get_wars_war_id_killmails_killmail_hash
	//
	// A hash of this killmail
	// Required: true
	KillmailHash *string `json:"killmail_hash"`

	// get_wars_war_id_killmails_killmail_id
	//
	// ID of this killmail
	// Required: true
	KillmailID *int32 `json:"killmail_id"`
}

// Validate validates this get wars war ID killmails o k body items0
func (o *GetWarsWarIDKillmailsOKBodyItems0) Validate(formats strfmt.Registry) error {
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

func (o *GetWarsWarIDKillmailsOKBodyItems0) validateKillmailHash(formats strfmt.Registry) error {

	if err := validate.Required("killmail_hash", "body", o.KillmailHash); err != nil {
		return err
	}

	return nil
}

func (o *GetWarsWarIDKillmailsOKBodyItems0) validateKillmailID(formats strfmt.Registry) error {

	if err := validate.Required("killmail_id", "body", o.KillmailID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get wars war ID killmails o k body items0 based on context it is used
func (o *GetWarsWarIDKillmailsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWarsWarIDKillmailsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWarsWarIDKillmailsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetWarsWarIDKillmailsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWarsWarIDKillmailsUnprocessableEntityBody get_wars_war_id_killmails_unprocessable_entity
//
// Unprocessable entity
swagger:model GetWarsWarIDKillmailsUnprocessableEntityBody
*/
type GetWarsWarIDKillmailsUnprocessableEntityBody struct {

	// get_wars_war_id_killmails_422_unprocessable_entity
	//
	// Unprocessable entity message
	Error string `json:"error,omitempty"`
}

// Validate validates this get wars war ID killmails unprocessable entity body
func (o *GetWarsWarIDKillmailsUnprocessableEntityBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get wars war ID killmails unprocessable entity body based on context it is used
func (o *GetWarsWarIDKillmailsUnprocessableEntityBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWarsWarIDKillmailsUnprocessableEntityBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWarsWarIDKillmailsUnprocessableEntityBody) UnmarshalBinary(b []byte) error {
	var res GetWarsWarIDKillmailsUnprocessableEntityBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
