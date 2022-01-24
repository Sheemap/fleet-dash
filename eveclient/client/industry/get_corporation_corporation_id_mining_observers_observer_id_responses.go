// Code generated by go-swagger; DO NOT EDIT.

package industry

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

// GetCorporationCorporationIDMiningObserversObserverIDReader is a Reader for the GetCorporationCorporationIDMiningObserversObserverID structure.
type GetCorporationCorporationIDMiningObserversObserverIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationCorporationIDMiningObserversObserverIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationCorporationIDMiningObserversObserverIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationCorporationIDMiningObserversObserverIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationCorporationIDMiningObserversObserverIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationCorporationIDMiningObserversObserverIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationCorporationIDMiningObserversObserverIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationCorporationIDMiningObserversObserverIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationCorporationIDMiningObserversObserverIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationCorporationIDMiningObserversObserverIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationCorporationIDMiningObserversObserverIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCorporationCorporationIDMiningObserversObserverIDOK creates a GetCorporationCorporationIDMiningObserversObserverIDOK with default headers values
func NewGetCorporationCorporationIDMiningObserversObserverIDOK() *GetCorporationCorporationIDMiningObserversObserverIDOK {
	var (
		// initialize headers with default values
		xPagesDefault = int32(1)
	)

	return &GetCorporationCorporationIDMiningObserversObserverIDOK{

		XPages: xPagesDefault,
	}
}

/* GetCorporationCorporationIDMiningObserversObserverIDOK describes a response with status code 200, with default header values.

Mining ledger of an observer
*/
type GetCorporationCorporationIDMiningObserversObserverIDOK struct {

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

	Payload []*GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems0
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDOK) Error() string {
	return fmt.Sprintf("[GET /corporation/{corporation_id}/mining/observers/{observer_id}/][%d] getCorporationCorporationIdMiningObserversObserverIdOK  %+v", 200, o.Payload)
}
func (o *GetCorporationCorporationIDMiningObserversObserverIDOK) GetPayload() []*GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems0 {
	return o.Payload
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationCorporationIDMiningObserversObserverIDNotModified creates a GetCorporationCorporationIDMiningObserversObserverIDNotModified with default headers values
func NewGetCorporationCorporationIDMiningObserversObserverIDNotModified() *GetCorporationCorporationIDMiningObserversObserverIDNotModified {
	return &GetCorporationCorporationIDMiningObserversObserverIDNotModified{}
}

/* GetCorporationCorporationIDMiningObserversObserverIDNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationCorporationIDMiningObserversObserverIDNotModified struct {

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

func (o *GetCorporationCorporationIDMiningObserversObserverIDNotModified) Error() string {
	return fmt.Sprintf("[GET /corporation/{corporation_id}/mining/observers/{observer_id}/][%d] getCorporationCorporationIdMiningObserversObserverIdNotModified ", 304)
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationCorporationIDMiningObserversObserverIDBadRequest creates a GetCorporationCorporationIDMiningObserversObserverIDBadRequest with default headers values
func NewGetCorporationCorporationIDMiningObserversObserverIDBadRequest() *GetCorporationCorporationIDMiningObserversObserverIDBadRequest {
	return &GetCorporationCorporationIDMiningObserversObserverIDBadRequest{}
}

/* GetCorporationCorporationIDMiningObserversObserverIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationCorporationIDMiningObserversObserverIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /corporation/{corporation_id}/mining/observers/{observer_id}/][%d] getCorporationCorporationIdMiningObserversObserverIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetCorporationCorporationIDMiningObserversObserverIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationCorporationIDMiningObserversObserverIDUnauthorized creates a GetCorporationCorporationIDMiningObserversObserverIDUnauthorized with default headers values
func NewGetCorporationCorporationIDMiningObserversObserverIDUnauthorized() *GetCorporationCorporationIDMiningObserversObserverIDUnauthorized {
	return &GetCorporationCorporationIDMiningObserversObserverIDUnauthorized{}
}

/* GetCorporationCorporationIDMiningObserversObserverIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationCorporationIDMiningObserversObserverIDUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /corporation/{corporation_id}/mining/observers/{observer_id}/][%d] getCorporationCorporationIdMiningObserversObserverIdUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCorporationCorporationIDMiningObserversObserverIDUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationCorporationIDMiningObserversObserverIDForbidden creates a GetCorporationCorporationIDMiningObserversObserverIDForbidden with default headers values
func NewGetCorporationCorporationIDMiningObserversObserverIDForbidden() *GetCorporationCorporationIDMiningObserversObserverIDForbidden {
	return &GetCorporationCorporationIDMiningObserversObserverIDForbidden{}
}

/* GetCorporationCorporationIDMiningObserversObserverIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationCorporationIDMiningObserversObserverIDForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDForbidden) Error() string {
	return fmt.Sprintf("[GET /corporation/{corporation_id}/mining/observers/{observer_id}/][%d] getCorporationCorporationIdMiningObserversObserverIdForbidden  %+v", 403, o.Payload)
}
func (o *GetCorporationCorporationIDMiningObserversObserverIDForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationCorporationIDMiningObserversObserverIDEnhanceYourCalm creates a GetCorporationCorporationIDMiningObserversObserverIDEnhanceYourCalm with default headers values
func NewGetCorporationCorporationIDMiningObserversObserverIDEnhanceYourCalm() *GetCorporationCorporationIDMiningObserversObserverIDEnhanceYourCalm {
	return &GetCorporationCorporationIDMiningObserversObserverIDEnhanceYourCalm{}
}

/* GetCorporationCorporationIDMiningObserversObserverIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationCorporationIDMiningObserversObserverIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /corporation/{corporation_id}/mining/observers/{observer_id}/][%d] getCorporationCorporationIdMiningObserversObserverIdEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCorporationCorporationIDMiningObserversObserverIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationCorporationIDMiningObserversObserverIDInternalServerError creates a GetCorporationCorporationIDMiningObserversObserverIDInternalServerError with default headers values
func NewGetCorporationCorporationIDMiningObserversObserverIDInternalServerError() *GetCorporationCorporationIDMiningObserversObserverIDInternalServerError {
	return &GetCorporationCorporationIDMiningObserversObserverIDInternalServerError{}
}

/* GetCorporationCorporationIDMiningObserversObserverIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationCorporationIDMiningObserversObserverIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporation/{corporation_id}/mining/observers/{observer_id}/][%d] getCorporationCorporationIdMiningObserversObserverIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCorporationCorporationIDMiningObserversObserverIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationCorporationIDMiningObserversObserverIDServiceUnavailable creates a GetCorporationCorporationIDMiningObserversObserverIDServiceUnavailable with default headers values
func NewGetCorporationCorporationIDMiningObserversObserverIDServiceUnavailable() *GetCorporationCorporationIDMiningObserversObserverIDServiceUnavailable {
	return &GetCorporationCorporationIDMiningObserversObserverIDServiceUnavailable{}
}

/* GetCorporationCorporationIDMiningObserversObserverIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationCorporationIDMiningObserversObserverIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /corporation/{corporation_id}/mining/observers/{observer_id}/][%d] getCorporationCorporationIdMiningObserversObserverIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCorporationCorporationIDMiningObserversObserverIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationCorporationIDMiningObserversObserverIDGatewayTimeout creates a GetCorporationCorporationIDMiningObserversObserverIDGatewayTimeout with default headers values
func NewGetCorporationCorporationIDMiningObserversObserverIDGatewayTimeout() *GetCorporationCorporationIDMiningObserversObserverIDGatewayTimeout {
	return &GetCorporationCorporationIDMiningObserversObserverIDGatewayTimeout{}
}

/* GetCorporationCorporationIDMiningObserversObserverIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationCorporationIDMiningObserversObserverIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /corporation/{corporation_id}/mining/observers/{observer_id}/][%d] getCorporationCorporationIdMiningObserversObserverIdGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCorporationCorporationIDMiningObserversObserverIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems0 get_corporation_corporation_id_mining_observers_observer_id_200_ok
//
// 200 ok object
swagger:model GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems0
*/
type GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems0 struct {

	// get_corporation_corporation_id_mining_observers_observer_id_character_id
	//
	// The character that did the mining
	//
	// Required: true
	CharacterID *int32 `json:"character_id"`

	// get_corporation_corporation_id_mining_observers_observer_id_last_updated
	//
	// last_updated string
	// Required: true
	// Format: date
	LastUpdated *strfmt.Date `json:"last_updated"`

	// get_corporation_corporation_id_mining_observers_observer_id_quantity
	//
	// quantity integer
	// Required: true
	Quantity *int64 `json:"quantity"`

	// get_corporation_corporation_id_mining_observers_observer_id_recorded_corporation_id
	//
	// The corporation id of the character at the time data was recorded.
	//
	// Required: true
	RecordedCorporationID *int32 `json:"recorded_corporation_id"`

	// get_corporation_corporation_id_mining_observers_observer_id_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get corporation corporation ID mining observers observer ID o k body items0
func (o *GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCharacterID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRecordedCorporationID(formats); err != nil {
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

func (o *GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems0) validateCharacterID(formats strfmt.Registry) error {

	if err := validate.Required("character_id", "body", o.CharacterID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems0) validateLastUpdated(formats strfmt.Registry) error {

	if err := validate.Required("last_updated", "body", o.LastUpdated); err != nil {
		return err
	}

	if err := validate.FormatOf("last_updated", "body", "date", o.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems0) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", o.Quantity); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems0) validateRecordedCorporationID(formats strfmt.Registry) error {

	if err := validate.Required("recorded_corporation_id", "body", o.RecordedCorporationID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get corporation corporation ID mining observers observer ID o k body items0 based on context it is used
func (o *GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
