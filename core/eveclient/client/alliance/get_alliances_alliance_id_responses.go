// Code generated by go-swagger; DO NOT EDIT.

package alliance

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

// GetAlliancesAllianceIDReader is a Reader for the GetAlliancesAllianceID structure.
type GetAlliancesAllianceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlliancesAllianceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlliancesAllianceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetAlliancesAllianceIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetAlliancesAllianceIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAlliancesAllianceIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetAlliancesAllianceIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAlliancesAllianceIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAlliancesAllianceIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAlliancesAllianceIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAlliancesAllianceIDOK creates a GetAlliancesAllianceIDOK with default headers values
func NewGetAlliancesAllianceIDOK() *GetAlliancesAllianceIDOK {
	return &GetAlliancesAllianceIDOK{}
}

/* GetAlliancesAllianceIDOK describes a response with status code 200, with default header values.

Public data about an alliance
*/
type GetAlliancesAllianceIDOK struct {

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

	Payload *GetAlliancesAllianceIDOKBody
}

func (o *GetAlliancesAllianceIDOK) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/][%d] getAlliancesAllianceIdOK  %+v", 200, o.Payload)
}
func (o *GetAlliancesAllianceIDOK) GetPayload() *GetAlliancesAllianceIDOKBody {
	return o.Payload
}

func (o *GetAlliancesAllianceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetAlliancesAllianceIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDNotModified creates a GetAlliancesAllianceIDNotModified with default headers values
func NewGetAlliancesAllianceIDNotModified() *GetAlliancesAllianceIDNotModified {
	return &GetAlliancesAllianceIDNotModified{}
}

/* GetAlliancesAllianceIDNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetAlliancesAllianceIDNotModified struct {

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

func (o *GetAlliancesAllianceIDNotModified) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/][%d] getAlliancesAllianceIdNotModified ", 304)
}

func (o *GetAlliancesAllianceIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlliancesAllianceIDBadRequest creates a GetAlliancesAllianceIDBadRequest with default headers values
func NewGetAlliancesAllianceIDBadRequest() *GetAlliancesAllianceIDBadRequest {
	return &GetAlliancesAllianceIDBadRequest{}
}

/* GetAlliancesAllianceIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetAlliancesAllianceIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetAlliancesAllianceIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/][%d] getAlliancesAllianceIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetAlliancesAllianceIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetAlliancesAllianceIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDNotFound creates a GetAlliancesAllianceIDNotFound with default headers values
func NewGetAlliancesAllianceIDNotFound() *GetAlliancesAllianceIDNotFound {
	return &GetAlliancesAllianceIDNotFound{}
}

/* GetAlliancesAllianceIDNotFound describes a response with status code 404, with default header values.

Alliance not found
*/
type GetAlliancesAllianceIDNotFound struct {
	Payload *GetAlliancesAllianceIDNotFoundBody
}

func (o *GetAlliancesAllianceIDNotFound) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/][%d] getAlliancesAllianceIdNotFound  %+v", 404, o.Payload)
}
func (o *GetAlliancesAllianceIDNotFound) GetPayload() *GetAlliancesAllianceIDNotFoundBody {
	return o.Payload
}

func (o *GetAlliancesAllianceIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAlliancesAllianceIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDEnhanceYourCalm creates a GetAlliancesAllianceIDEnhanceYourCalm with default headers values
func NewGetAlliancesAllianceIDEnhanceYourCalm() *GetAlliancesAllianceIDEnhanceYourCalm {
	return &GetAlliancesAllianceIDEnhanceYourCalm{}
}

/* GetAlliancesAllianceIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetAlliancesAllianceIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetAlliancesAllianceIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/][%d] getAlliancesAllianceIdEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetAlliancesAllianceIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetAlliancesAllianceIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDInternalServerError creates a GetAlliancesAllianceIDInternalServerError with default headers values
func NewGetAlliancesAllianceIDInternalServerError() *GetAlliancesAllianceIDInternalServerError {
	return &GetAlliancesAllianceIDInternalServerError{}
}

/* GetAlliancesAllianceIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetAlliancesAllianceIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetAlliancesAllianceIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/][%d] getAlliancesAllianceIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAlliancesAllianceIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetAlliancesAllianceIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDServiceUnavailable creates a GetAlliancesAllianceIDServiceUnavailable with default headers values
func NewGetAlliancesAllianceIDServiceUnavailable() *GetAlliancesAllianceIDServiceUnavailable {
	return &GetAlliancesAllianceIDServiceUnavailable{}
}

/* GetAlliancesAllianceIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetAlliancesAllianceIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetAlliancesAllianceIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/][%d] getAlliancesAllianceIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetAlliancesAllianceIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetAlliancesAllianceIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDGatewayTimeout creates a GetAlliancesAllianceIDGatewayTimeout with default headers values
func NewGetAlliancesAllianceIDGatewayTimeout() *GetAlliancesAllianceIDGatewayTimeout {
	return &GetAlliancesAllianceIDGatewayTimeout{}
}

/* GetAlliancesAllianceIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetAlliancesAllianceIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetAlliancesAllianceIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/][%d] getAlliancesAllianceIdGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetAlliancesAllianceIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetAlliancesAllianceIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAlliancesAllianceIDNotFoundBody get_alliances_alliance_id_not_found
//
// Not found
swagger:model GetAlliancesAllianceIDNotFoundBody
*/
type GetAlliancesAllianceIDNotFoundBody struct {

	// get_alliances_alliance_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get alliances alliance ID not found body
func (o *GetAlliancesAllianceIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get alliances alliance ID not found body based on context it is used
func (o *GetAlliancesAllianceIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAlliancesAllianceIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAlliancesAllianceIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetAlliancesAllianceIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAlliancesAllianceIDOKBody get_alliances_alliance_id_ok
//
// 200 ok object
swagger:model GetAlliancesAllianceIDOKBody
*/
type GetAlliancesAllianceIDOKBody struct {

	// get_alliances_alliance_id_creator_corporation_id
	//
	// ID of the corporation that created the alliance
	// Required: true
	CreatorCorporationID *int32 `json:"creator_corporation_id"`

	// get_alliances_alliance_id_creator_id
	//
	// ID of the character that created the alliance
	// Required: true
	CreatorID *int32 `json:"creator_id"`

	// get_alliances_alliance_id_date_founded
	//
	// date_founded string
	// Required: true
	// Format: date-time
	DateFounded *strfmt.DateTime `json:"date_founded"`

	// get_alliances_alliance_id_executor_corporation_id
	//
	// the executor corporation ID, if this alliance is not closed
	ExecutorCorporationID int32 `json:"executor_corporation_id,omitempty"`

	// get_alliances_alliance_id_faction_id
	//
	// Faction ID this alliance is fighting for, if this alliance is enlisted in factional warfare
	FactionID int32 `json:"faction_id,omitempty"`

	// get_alliances_alliance_id_name
	//
	// the full name of the alliance
	// Required: true
	Name *string `json:"name"`

	// get_alliances_alliance_id_ticker
	//
	// the short name of the alliance
	// Required: true
	Ticker *string `json:"ticker"`
}

// Validate validates this get alliances alliance ID o k body
func (o *GetAlliancesAllianceIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreatorCorporationID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreatorID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDateFounded(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTicker(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAlliancesAllianceIDOKBody) validateCreatorCorporationID(formats strfmt.Registry) error {

	if err := validate.Required("getAlliancesAllianceIdOK"+"."+"creator_corporation_id", "body", o.CreatorCorporationID); err != nil {
		return err
	}

	return nil
}

func (o *GetAlliancesAllianceIDOKBody) validateCreatorID(formats strfmt.Registry) error {

	if err := validate.Required("getAlliancesAllianceIdOK"+"."+"creator_id", "body", o.CreatorID); err != nil {
		return err
	}

	return nil
}

func (o *GetAlliancesAllianceIDOKBody) validateDateFounded(formats strfmt.Registry) error {

	if err := validate.Required("getAlliancesAllianceIdOK"+"."+"date_founded", "body", o.DateFounded); err != nil {
		return err
	}

	if err := validate.FormatOf("getAlliancesAllianceIdOK"+"."+"date_founded", "body", "date-time", o.DateFounded.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetAlliancesAllianceIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getAlliancesAllianceIdOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetAlliancesAllianceIDOKBody) validateTicker(formats strfmt.Registry) error {

	if err := validate.Required("getAlliancesAllianceIdOK"+"."+"ticker", "body", o.Ticker); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get alliances alliance ID o k body based on context it is used
func (o *GetAlliancesAllianceIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAlliancesAllianceIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAlliancesAllianceIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetAlliancesAllianceIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}