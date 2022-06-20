// Code generated by go-swagger; DO NOT EDIT.

package character

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

// GetCharactersCharacterIDCorporationhistoryReader is a Reader for the GetCharactersCharacterIDCorporationhistory structure.
type GetCharactersCharacterIDCorporationhistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDCorporationhistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDCorporationhistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDCorporationhistoryNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDCorporationhistoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDCorporationhistoryEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDCorporationhistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDCorporationhistoryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDCorporationhistoryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersCharacterIDCorporationhistoryOK creates a GetCharactersCharacterIDCorporationhistoryOK with default headers values
func NewGetCharactersCharacterIDCorporationhistoryOK() *GetCharactersCharacterIDCorporationhistoryOK {
	return &GetCharactersCharacterIDCorporationhistoryOK{}
}

/* GetCharactersCharacterIDCorporationhistoryOK describes a response with status code 200, with default header values.

Corporation history for the given character
*/
type GetCharactersCharacterIDCorporationhistoryOK struct {

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

	Payload []*GetCharactersCharacterIDCorporationhistoryOKBodyItems0
}

func (o *GetCharactersCharacterIDCorporationhistoryOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/corporationhistory/][%d] getCharactersCharacterIdCorporationhistoryOK  %+v", 200, o.Payload)
}
func (o *GetCharactersCharacterIDCorporationhistoryOK) GetPayload() []*GetCharactersCharacterIDCorporationhistoryOKBodyItems0 {
	return o.Payload
}

func (o *GetCharactersCharacterIDCorporationhistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDCorporationhistoryNotModified creates a GetCharactersCharacterIDCorporationhistoryNotModified with default headers values
func NewGetCharactersCharacterIDCorporationhistoryNotModified() *GetCharactersCharacterIDCorporationhistoryNotModified {
	return &GetCharactersCharacterIDCorporationhistoryNotModified{}
}

/* GetCharactersCharacterIDCorporationhistoryNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDCorporationhistoryNotModified struct {

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

func (o *GetCharactersCharacterIDCorporationhistoryNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/corporationhistory/][%d] getCharactersCharacterIdCorporationhistoryNotModified ", 304)
}

func (o *GetCharactersCharacterIDCorporationhistoryNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDCorporationhistoryBadRequest creates a GetCharactersCharacterIDCorporationhistoryBadRequest with default headers values
func NewGetCharactersCharacterIDCorporationhistoryBadRequest() *GetCharactersCharacterIDCorporationhistoryBadRequest {
	return &GetCharactersCharacterIDCorporationhistoryBadRequest{}
}

/* GetCharactersCharacterIDCorporationhistoryBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDCorporationhistoryBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDCorporationhistoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/corporationhistory/][%d] getCharactersCharacterIdCorporationhistoryBadRequest  %+v", 400, o.Payload)
}
func (o *GetCharactersCharacterIDCorporationhistoryBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDCorporationhistoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDCorporationhistoryEnhanceYourCalm creates a GetCharactersCharacterIDCorporationhistoryEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDCorporationhistoryEnhanceYourCalm() *GetCharactersCharacterIDCorporationhistoryEnhanceYourCalm {
	return &GetCharactersCharacterIDCorporationhistoryEnhanceYourCalm{}
}

/* GetCharactersCharacterIDCorporationhistoryEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDCorporationhistoryEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDCorporationhistoryEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/corporationhistory/][%d] getCharactersCharacterIdCorporationhistoryEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCharactersCharacterIDCorporationhistoryEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDCorporationhistoryEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDCorporationhistoryInternalServerError creates a GetCharactersCharacterIDCorporationhistoryInternalServerError with default headers values
func NewGetCharactersCharacterIDCorporationhistoryInternalServerError() *GetCharactersCharacterIDCorporationhistoryInternalServerError {
	return &GetCharactersCharacterIDCorporationhistoryInternalServerError{}
}

/* GetCharactersCharacterIDCorporationhistoryInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDCorporationhistoryInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDCorporationhistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/corporationhistory/][%d] getCharactersCharacterIdCorporationhistoryInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCharactersCharacterIDCorporationhistoryInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDCorporationhistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDCorporationhistoryServiceUnavailable creates a GetCharactersCharacterIDCorporationhistoryServiceUnavailable with default headers values
func NewGetCharactersCharacterIDCorporationhistoryServiceUnavailable() *GetCharactersCharacterIDCorporationhistoryServiceUnavailable {
	return &GetCharactersCharacterIDCorporationhistoryServiceUnavailable{}
}

/* GetCharactersCharacterIDCorporationhistoryServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDCorporationhistoryServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDCorporationhistoryServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/corporationhistory/][%d] getCharactersCharacterIdCorporationhistoryServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCharactersCharacterIDCorporationhistoryServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDCorporationhistoryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDCorporationhistoryGatewayTimeout creates a GetCharactersCharacterIDCorporationhistoryGatewayTimeout with default headers values
func NewGetCharactersCharacterIDCorporationhistoryGatewayTimeout() *GetCharactersCharacterIDCorporationhistoryGatewayTimeout {
	return &GetCharactersCharacterIDCorporationhistoryGatewayTimeout{}
}

/* GetCharactersCharacterIDCorporationhistoryGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDCorporationhistoryGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDCorporationhistoryGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/corporationhistory/][%d] getCharactersCharacterIdCorporationhistoryGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCharactersCharacterIDCorporationhistoryGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDCorporationhistoryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDCorporationhistoryOKBodyItems0 get_characters_character_id_corporationhistory_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDCorporationhistoryOKBodyItems0
*/
type GetCharactersCharacterIDCorporationhistoryOKBodyItems0 struct {

	// get_characters_character_id_corporationhistory_corporation_id
	//
	// corporation_id integer
	// Required: true
	CorporationID *int32 `json:"corporation_id"`

	// get_characters_character_id_corporationhistory_is_deleted
	//
	// True if the corporation has been deleted
	IsDeleted bool `json:"is_deleted,omitempty"`

	// get_characters_character_id_corporationhistory_record_id
	//
	// An incrementing ID that can be used to canonically establish order of records in cases where dates may be ambiguous
	// Required: true
	RecordID *int32 `json:"record_id"`

	// get_characters_character_id_corporationhistory_start_date
	//
	// start_date string
	// Required: true
	// Format: date-time
	StartDate *strfmt.DateTime `json:"start_date"`
}

// Validate validates this get characters character ID corporationhistory o k body items0
func (o *GetCharactersCharacterIDCorporationhistoryOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCorporationID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRecordID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDCorporationhistoryOKBodyItems0) validateCorporationID(formats strfmt.Registry) error {

	if err := validate.Required("corporation_id", "body", o.CorporationID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDCorporationhistoryOKBodyItems0) validateRecordID(formats strfmt.Registry) error {

	if err := validate.Required("record_id", "body", o.RecordID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDCorporationhistoryOKBodyItems0) validateStartDate(formats strfmt.Registry) error {

	if err := validate.Required("start_date", "body", o.StartDate); err != nil {
		return err
	}

	if err := validate.FormatOf("start_date", "body", "date-time", o.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get characters character ID corporationhistory o k body items0 based on context it is used
func (o *GetCharactersCharacterIDCorporationhistoryOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDCorporationhistoryOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDCorporationhistoryOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDCorporationhistoryOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}