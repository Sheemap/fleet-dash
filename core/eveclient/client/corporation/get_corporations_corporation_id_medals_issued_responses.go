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

// GetCorporationsCorporationIDMedalsIssuedReader is a Reader for the GetCorporationsCorporationIDMedalsIssued structure.
type GetCorporationsCorporationIDMedalsIssuedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDMedalsIssuedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationsCorporationIDMedalsIssuedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationsCorporationIDMedalsIssuedNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationsCorporationIDMedalsIssuedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationsCorporationIDMedalsIssuedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationsCorporationIDMedalsIssuedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationsCorporationIDMedalsIssuedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationsCorporationIDMedalsIssuedServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationsCorporationIDMedalsIssuedGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDMedalsIssuedOK creates a GetCorporationsCorporationIDMedalsIssuedOK with default headers values
func NewGetCorporationsCorporationIDMedalsIssuedOK() *GetCorporationsCorporationIDMedalsIssuedOK {
	var (
		// initialize headers with default values
		xPagesDefault = int32(1)
	)

	return &GetCorporationsCorporationIDMedalsIssuedOK{

		XPages: xPagesDefault,
	}
}

/* GetCorporationsCorporationIDMedalsIssuedOK describes a response with status code 200, with default header values.

A list of issued medals
*/
type GetCorporationsCorporationIDMedalsIssuedOK struct {

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

	Payload []*GetCorporationsCorporationIDMedalsIssuedOKBodyItems0
}

func (o *GetCorporationsCorporationIDMedalsIssuedOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/issued/][%d] getCorporationsCorporationIdMedalsIssuedOK  %+v", 200, o.Payload)
}
func (o *GetCorporationsCorporationIDMedalsIssuedOK) GetPayload() []*GetCorporationsCorporationIDMedalsIssuedOKBodyItems0 {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMedalsIssuedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDMedalsIssuedNotModified creates a GetCorporationsCorporationIDMedalsIssuedNotModified with default headers values
func NewGetCorporationsCorporationIDMedalsIssuedNotModified() *GetCorporationsCorporationIDMedalsIssuedNotModified {
	return &GetCorporationsCorporationIDMedalsIssuedNotModified{}
}

/* GetCorporationsCorporationIDMedalsIssuedNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationsCorporationIDMedalsIssuedNotModified struct {

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

func (o *GetCorporationsCorporationIDMedalsIssuedNotModified) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/issued/][%d] getCorporationsCorporationIdMedalsIssuedNotModified ", 304)
}

func (o *GetCorporationsCorporationIDMedalsIssuedNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDMedalsIssuedBadRequest creates a GetCorporationsCorporationIDMedalsIssuedBadRequest with default headers values
func NewGetCorporationsCorporationIDMedalsIssuedBadRequest() *GetCorporationsCorporationIDMedalsIssuedBadRequest {
	return &GetCorporationsCorporationIDMedalsIssuedBadRequest{}
}

/* GetCorporationsCorporationIDMedalsIssuedBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationsCorporationIDMedalsIssuedBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDMedalsIssuedBadRequest) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/issued/][%d] getCorporationsCorporationIdMedalsIssuedBadRequest  %+v", 400, o.Payload)
}
func (o *GetCorporationsCorporationIDMedalsIssuedBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMedalsIssuedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsIssuedUnauthorized creates a GetCorporationsCorporationIDMedalsIssuedUnauthorized with default headers values
func NewGetCorporationsCorporationIDMedalsIssuedUnauthorized() *GetCorporationsCorporationIDMedalsIssuedUnauthorized {
	return &GetCorporationsCorporationIDMedalsIssuedUnauthorized{}
}

/* GetCorporationsCorporationIDMedalsIssuedUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDMedalsIssuedUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDMedalsIssuedUnauthorized) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/issued/][%d] getCorporationsCorporationIdMedalsIssuedUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCorporationsCorporationIDMedalsIssuedUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMedalsIssuedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsIssuedForbidden creates a GetCorporationsCorporationIDMedalsIssuedForbidden with default headers values
func NewGetCorporationsCorporationIDMedalsIssuedForbidden() *GetCorporationsCorporationIDMedalsIssuedForbidden {
	return &GetCorporationsCorporationIDMedalsIssuedForbidden{}
}

/* GetCorporationsCorporationIDMedalsIssuedForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationsCorporationIDMedalsIssuedForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDMedalsIssuedForbidden) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/issued/][%d] getCorporationsCorporationIdMedalsIssuedForbidden  %+v", 403, o.Payload)
}
func (o *GetCorporationsCorporationIDMedalsIssuedForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMedalsIssuedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm creates a GetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm() *GetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm {
	return &GetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm{}
}

/* GetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/issued/][%d] getCorporationsCorporationIdMedalsIssuedEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsIssuedInternalServerError creates a GetCorporationsCorporationIDMedalsIssuedInternalServerError with default headers values
func NewGetCorporationsCorporationIDMedalsIssuedInternalServerError() *GetCorporationsCorporationIDMedalsIssuedInternalServerError {
	return &GetCorporationsCorporationIDMedalsIssuedInternalServerError{}
}

/* GetCorporationsCorporationIDMedalsIssuedInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationsCorporationIDMedalsIssuedInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDMedalsIssuedInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/issued/][%d] getCorporationsCorporationIdMedalsIssuedInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCorporationsCorporationIDMedalsIssuedInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMedalsIssuedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsIssuedServiceUnavailable creates a GetCorporationsCorporationIDMedalsIssuedServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDMedalsIssuedServiceUnavailable() *GetCorporationsCorporationIDMedalsIssuedServiceUnavailable {
	return &GetCorporationsCorporationIDMedalsIssuedServiceUnavailable{}
}

/* GetCorporationsCorporationIDMedalsIssuedServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDMedalsIssuedServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDMedalsIssuedServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/issued/][%d] getCorporationsCorporationIdMedalsIssuedServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCorporationsCorporationIDMedalsIssuedServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMedalsIssuedServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsIssuedGatewayTimeout creates a GetCorporationsCorporationIDMedalsIssuedGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDMedalsIssuedGatewayTimeout() *GetCorporationsCorporationIDMedalsIssuedGatewayTimeout {
	return &GetCorporationsCorporationIDMedalsIssuedGatewayTimeout{}
}

/* GetCorporationsCorporationIDMedalsIssuedGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDMedalsIssuedGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDMedalsIssuedGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/issued/][%d] getCorporationsCorporationIdMedalsIssuedGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCorporationsCorporationIDMedalsIssuedGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMedalsIssuedGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDMedalsIssuedOKBodyItems0 get_corporations_corporation_id_medals_issued_200_ok
//
// 200 ok object
swagger:model GetCorporationsCorporationIDMedalsIssuedOKBodyItems0
*/
type GetCorporationsCorporationIDMedalsIssuedOKBodyItems0 struct {

	// get_corporations_corporation_id_medals_issued_character_id
	//
	// ID of the character who was rewarded this medal
	// Required: true
	CharacterID *int32 `json:"character_id"`

	// get_corporations_corporation_id_medals_issued_issued_at
	//
	// issued_at string
	// Required: true
	// Format: date-time
	IssuedAt *strfmt.DateTime `json:"issued_at"`

	// get_corporations_corporation_id_medals_issued_issuer_id
	//
	// ID of the character who issued the medal
	// Required: true
	IssuerID *int32 `json:"issuer_id"`

	// get_corporations_corporation_id_medals_issued_medal_id
	//
	// medal_id integer
	// Required: true
	MedalID *int32 `json:"medal_id"`

	// get_corporations_corporation_id_medals_issued_reason
	//
	// reason string
	// Required: true
	// Max Length: 1000
	Reason *string `json:"reason"`

	// get_corporations_corporation_id_medals_issued_status
	//
	// status string
	// Required: true
	// Enum: [private public]
	Status *string `json:"status"`
}

// Validate validates this get corporations corporation ID medals issued o k body items0
func (o *GetCorporationsCorporationIDMedalsIssuedOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCharacterID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIssuedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIssuerID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMedalID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationsCorporationIDMedalsIssuedOKBodyItems0) validateCharacterID(formats strfmt.Registry) error {

	if err := validate.Required("character_id", "body", o.CharacterID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDMedalsIssuedOKBodyItems0) validateIssuedAt(formats strfmt.Registry) error {

	if err := validate.Required("issued_at", "body", o.IssuedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("issued_at", "body", "date-time", o.IssuedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDMedalsIssuedOKBodyItems0) validateIssuerID(formats strfmt.Registry) error {

	if err := validate.Required("issuer_id", "body", o.IssuerID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDMedalsIssuedOKBodyItems0) validateMedalID(formats strfmt.Registry) error {

	if err := validate.Required("medal_id", "body", o.MedalID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDMedalsIssuedOKBodyItems0) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("reason", "body", o.Reason); err != nil {
		return err
	}

	if err := validate.MaxLength("reason", "body", *o.Reason, 1000); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdMedalsIssuedOKBodyItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["private","public"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdMedalsIssuedOKBodyItems0TypeStatusPropEnum = append(getCorporationsCorporationIdMedalsIssuedOKBodyItems0TypeStatusPropEnum, v)
	}
}

const (

	// GetCorporationsCorporationIDMedalsIssuedOKBodyItems0StatusPrivate captures enum value "private"
	GetCorporationsCorporationIDMedalsIssuedOKBodyItems0StatusPrivate string = "private"

	// GetCorporationsCorporationIDMedalsIssuedOKBodyItems0StatusPublic captures enum value "public"
	GetCorporationsCorporationIDMedalsIssuedOKBodyItems0StatusPublic string = "public"
)

// prop value enum
func (o *GetCorporationsCorporationIDMedalsIssuedOKBodyItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCorporationsCorporationIdMedalsIssuedOKBodyItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDMedalsIssuedOKBodyItems0) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", o.Status); err != nil {
		return err
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get corporations corporation ID medals issued o k body items0 based on context it is used
func (o *GetCorporationsCorporationIDMedalsIssuedOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDMedalsIssuedOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDMedalsIssuedOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDMedalsIssuedOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
