// Code generated by go-swagger; DO NOT EDIT.

package corporation

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

// GetCorporationsCorporationIDShareholdersReader is a Reader for the GetCorporationsCorporationIDShareholders structure.
type GetCorporationsCorporationIDShareholdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDShareholdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationsCorporationIDShareholdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationsCorporationIDShareholdersNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationsCorporationIDShareholdersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationsCorporationIDShareholdersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationsCorporationIDShareholdersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationsCorporationIDShareholdersEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationsCorporationIDShareholdersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationsCorporationIDShareholdersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationsCorporationIDShareholdersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDShareholdersOK creates a GetCorporationsCorporationIDShareholdersOK with default headers values
func NewGetCorporationsCorporationIDShareholdersOK() *GetCorporationsCorporationIDShareholdersOK {
	var (
		// initialize headers with default values
		xPagesDefault = int32(1)
	)

	return &GetCorporationsCorporationIDShareholdersOK{

		XPages: xPagesDefault,
	}
}

/* GetCorporationsCorporationIDShareholdersOK describes a response with status code 200, with default header values.

List of shareholders
*/
type GetCorporationsCorporationIDShareholdersOK struct {

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

	Payload []*GetCorporationsCorporationIDShareholdersOKBodyItems0
}

func (o *GetCorporationsCorporationIDShareholdersOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/shareholders/][%d] getCorporationsCorporationIdShareholdersOK  %+v", 200, o.Payload)
}
func (o *GetCorporationsCorporationIDShareholdersOK) GetPayload() []*GetCorporationsCorporationIDShareholdersOKBodyItems0 {
	return o.Payload
}

func (o *GetCorporationsCorporationIDShareholdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDShareholdersNotModified creates a GetCorporationsCorporationIDShareholdersNotModified with default headers values
func NewGetCorporationsCorporationIDShareholdersNotModified() *GetCorporationsCorporationIDShareholdersNotModified {
	return &GetCorporationsCorporationIDShareholdersNotModified{}
}

/* GetCorporationsCorporationIDShareholdersNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationsCorporationIDShareholdersNotModified struct {

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

func (o *GetCorporationsCorporationIDShareholdersNotModified) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/shareholders/][%d] getCorporationsCorporationIdShareholdersNotModified ", 304)
}

func (o *GetCorporationsCorporationIDShareholdersNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDShareholdersBadRequest creates a GetCorporationsCorporationIDShareholdersBadRequest with default headers values
func NewGetCorporationsCorporationIDShareholdersBadRequest() *GetCorporationsCorporationIDShareholdersBadRequest {
	return &GetCorporationsCorporationIDShareholdersBadRequest{}
}

/* GetCorporationsCorporationIDShareholdersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationsCorporationIDShareholdersBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDShareholdersBadRequest) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/shareholders/][%d] getCorporationsCorporationIdShareholdersBadRequest  %+v", 400, o.Payload)
}
func (o *GetCorporationsCorporationIDShareholdersBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationsCorporationIDShareholdersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDShareholdersUnauthorized creates a GetCorporationsCorporationIDShareholdersUnauthorized with default headers values
func NewGetCorporationsCorporationIDShareholdersUnauthorized() *GetCorporationsCorporationIDShareholdersUnauthorized {
	return &GetCorporationsCorporationIDShareholdersUnauthorized{}
}

/* GetCorporationsCorporationIDShareholdersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDShareholdersUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDShareholdersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/shareholders/][%d] getCorporationsCorporationIdShareholdersUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCorporationsCorporationIDShareholdersUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationsCorporationIDShareholdersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDShareholdersForbidden creates a GetCorporationsCorporationIDShareholdersForbidden with default headers values
func NewGetCorporationsCorporationIDShareholdersForbidden() *GetCorporationsCorporationIDShareholdersForbidden {
	return &GetCorporationsCorporationIDShareholdersForbidden{}
}

/* GetCorporationsCorporationIDShareholdersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationsCorporationIDShareholdersForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDShareholdersForbidden) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/shareholders/][%d] getCorporationsCorporationIdShareholdersForbidden  %+v", 403, o.Payload)
}
func (o *GetCorporationsCorporationIDShareholdersForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationsCorporationIDShareholdersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDShareholdersEnhanceYourCalm creates a GetCorporationsCorporationIDShareholdersEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDShareholdersEnhanceYourCalm() *GetCorporationsCorporationIDShareholdersEnhanceYourCalm {
	return &GetCorporationsCorporationIDShareholdersEnhanceYourCalm{}
}

/* GetCorporationsCorporationIDShareholdersEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationsCorporationIDShareholdersEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDShareholdersEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/shareholders/][%d] getCorporationsCorporationIdShareholdersEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCorporationsCorporationIDShareholdersEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationsCorporationIDShareholdersEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDShareholdersInternalServerError creates a GetCorporationsCorporationIDShareholdersInternalServerError with default headers values
func NewGetCorporationsCorporationIDShareholdersInternalServerError() *GetCorporationsCorporationIDShareholdersInternalServerError {
	return &GetCorporationsCorporationIDShareholdersInternalServerError{}
}

/* GetCorporationsCorporationIDShareholdersInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationsCorporationIDShareholdersInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDShareholdersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/shareholders/][%d] getCorporationsCorporationIdShareholdersInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCorporationsCorporationIDShareholdersInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationsCorporationIDShareholdersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDShareholdersServiceUnavailable creates a GetCorporationsCorporationIDShareholdersServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDShareholdersServiceUnavailable() *GetCorporationsCorporationIDShareholdersServiceUnavailable {
	return &GetCorporationsCorporationIDShareholdersServiceUnavailable{}
}

/* GetCorporationsCorporationIDShareholdersServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDShareholdersServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDShareholdersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/shareholders/][%d] getCorporationsCorporationIdShareholdersServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCorporationsCorporationIDShareholdersServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationsCorporationIDShareholdersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDShareholdersGatewayTimeout creates a GetCorporationsCorporationIDShareholdersGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDShareholdersGatewayTimeout() *GetCorporationsCorporationIDShareholdersGatewayTimeout {
	return &GetCorporationsCorporationIDShareholdersGatewayTimeout{}
}

/* GetCorporationsCorporationIDShareholdersGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDShareholdersGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDShareholdersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/shareholders/][%d] getCorporationsCorporationIdShareholdersGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCorporationsCorporationIDShareholdersGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationsCorporationIDShareholdersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDShareholdersOKBodyItems0 get_corporations_corporation_id_shareholders_200_ok
//
// 200 ok object
swagger:model GetCorporationsCorporationIDShareholdersOKBodyItems0
*/
type GetCorporationsCorporationIDShareholdersOKBodyItems0 struct {

	// get_corporations_corporation_id_shareholders_share_count
	//
	// share_count integer
	// Required: true
	ShareCount *int64 `json:"share_count"`

	// get_corporations_corporation_id_shareholders_shareholder_id
	//
	// shareholder_id integer
	// Required: true
	ShareholderID *int32 `json:"shareholder_id"`

	// get_corporations_corporation_id_shareholders_shareholder_type
	//
	// shareholder_type string
	// Required: true
	// Enum: [character corporation]
	ShareholderType *string `json:"shareholder_type"`
}

// Validate validates this get corporations corporation ID shareholders o k body items0
func (o *GetCorporationsCorporationIDShareholdersOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateShareCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateShareholderID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateShareholderType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationsCorporationIDShareholdersOKBodyItems0) validateShareCount(formats strfmt.Registry) error {

	if err := validate.Required("share_count", "body", o.ShareCount); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDShareholdersOKBodyItems0) validateShareholderID(formats strfmt.Registry) error {

	if err := validate.Required("shareholder_id", "body", o.ShareholderID); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdShareholdersOKBodyItems0TypeShareholderTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["character","corporation"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdShareholdersOKBodyItems0TypeShareholderTypePropEnum = append(getCorporationsCorporationIdShareholdersOKBodyItems0TypeShareholderTypePropEnum, v)
	}
}

const (

	// GetCorporationsCorporationIDShareholdersOKBodyItems0ShareholderTypeCharacter captures enum value "character"
	GetCorporationsCorporationIDShareholdersOKBodyItems0ShareholderTypeCharacter string = "character"

	// GetCorporationsCorporationIDShareholdersOKBodyItems0ShareholderTypeCorporation captures enum value "corporation"
	GetCorporationsCorporationIDShareholdersOKBodyItems0ShareholderTypeCorporation string = "corporation"
)

// prop value enum
func (o *GetCorporationsCorporationIDShareholdersOKBodyItems0) validateShareholderTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCorporationsCorporationIdShareholdersOKBodyItems0TypeShareholderTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDShareholdersOKBodyItems0) validateShareholderType(formats strfmt.Registry) error {

	if err := validate.Required("shareholder_type", "body", o.ShareholderType); err != nil {
		return err
	}

	// value enum
	if err := o.validateShareholderTypeEnum("shareholder_type", "body", *o.ShareholderType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get corporations corporation ID shareholders o k body items0 based on context it is used
func (o *GetCorporationsCorporationIDShareholdersOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDShareholdersOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDShareholdersOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDShareholdersOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
