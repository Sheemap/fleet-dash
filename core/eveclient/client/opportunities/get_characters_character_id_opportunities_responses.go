// Code generated by go-swagger; DO NOT EDIT.

package opportunities

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

// GetCharactersCharacterIDOpportunitiesReader is a Reader for the GetCharactersCharacterIDOpportunities structure.
type GetCharactersCharacterIDOpportunitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDOpportunitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDOpportunitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDOpportunitiesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDOpportunitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDOpportunitiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDOpportunitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDOpportunitiesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDOpportunitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDOpportunitiesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDOpportunitiesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersCharacterIDOpportunitiesOK creates a GetCharactersCharacterIDOpportunitiesOK with default headers values
func NewGetCharactersCharacterIDOpportunitiesOK() *GetCharactersCharacterIDOpportunitiesOK {
	return &GetCharactersCharacterIDOpportunitiesOK{}
}

/* GetCharactersCharacterIDOpportunitiesOK describes a response with status code 200, with default header values.

A list of opportunities task ids
*/
type GetCharactersCharacterIDOpportunitiesOK struct {

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

	Payload []*GetCharactersCharacterIDOpportunitiesOKBodyItems0
}

func (o *GetCharactersCharacterIDOpportunitiesOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/opportunities/][%d] getCharactersCharacterIdOpportunitiesOK  %+v", 200, o.Payload)
}
func (o *GetCharactersCharacterIDOpportunitiesOK) GetPayload() []*GetCharactersCharacterIDOpportunitiesOKBodyItems0 {
	return o.Payload
}

func (o *GetCharactersCharacterIDOpportunitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDOpportunitiesNotModified creates a GetCharactersCharacterIDOpportunitiesNotModified with default headers values
func NewGetCharactersCharacterIDOpportunitiesNotModified() *GetCharactersCharacterIDOpportunitiesNotModified {
	return &GetCharactersCharacterIDOpportunitiesNotModified{}
}

/* GetCharactersCharacterIDOpportunitiesNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDOpportunitiesNotModified struct {

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

func (o *GetCharactersCharacterIDOpportunitiesNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/opportunities/][%d] getCharactersCharacterIdOpportunitiesNotModified ", 304)
}

func (o *GetCharactersCharacterIDOpportunitiesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDOpportunitiesBadRequest creates a GetCharactersCharacterIDOpportunitiesBadRequest with default headers values
func NewGetCharactersCharacterIDOpportunitiesBadRequest() *GetCharactersCharacterIDOpportunitiesBadRequest {
	return &GetCharactersCharacterIDOpportunitiesBadRequest{}
}

/* GetCharactersCharacterIDOpportunitiesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDOpportunitiesBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDOpportunitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/opportunities/][%d] getCharactersCharacterIdOpportunitiesBadRequest  %+v", 400, o.Payload)
}
func (o *GetCharactersCharacterIDOpportunitiesBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDOpportunitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOpportunitiesUnauthorized creates a GetCharactersCharacterIDOpportunitiesUnauthorized with default headers values
func NewGetCharactersCharacterIDOpportunitiesUnauthorized() *GetCharactersCharacterIDOpportunitiesUnauthorized {
	return &GetCharactersCharacterIDOpportunitiesUnauthorized{}
}

/* GetCharactersCharacterIDOpportunitiesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDOpportunitiesUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDOpportunitiesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/opportunities/][%d] getCharactersCharacterIdOpportunitiesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCharactersCharacterIDOpportunitiesUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDOpportunitiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOpportunitiesForbidden creates a GetCharactersCharacterIDOpportunitiesForbidden with default headers values
func NewGetCharactersCharacterIDOpportunitiesForbidden() *GetCharactersCharacterIDOpportunitiesForbidden {
	return &GetCharactersCharacterIDOpportunitiesForbidden{}
}

/* GetCharactersCharacterIDOpportunitiesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDOpportunitiesForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDOpportunitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/opportunities/][%d] getCharactersCharacterIdOpportunitiesForbidden  %+v", 403, o.Payload)
}
func (o *GetCharactersCharacterIDOpportunitiesForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDOpportunitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOpportunitiesEnhanceYourCalm creates a GetCharactersCharacterIDOpportunitiesEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDOpportunitiesEnhanceYourCalm() *GetCharactersCharacterIDOpportunitiesEnhanceYourCalm {
	return &GetCharactersCharacterIDOpportunitiesEnhanceYourCalm{}
}

/* GetCharactersCharacterIDOpportunitiesEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDOpportunitiesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDOpportunitiesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/opportunities/][%d] getCharactersCharacterIdOpportunitiesEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCharactersCharacterIDOpportunitiesEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDOpportunitiesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOpportunitiesInternalServerError creates a GetCharactersCharacterIDOpportunitiesInternalServerError with default headers values
func NewGetCharactersCharacterIDOpportunitiesInternalServerError() *GetCharactersCharacterIDOpportunitiesInternalServerError {
	return &GetCharactersCharacterIDOpportunitiesInternalServerError{}
}

/* GetCharactersCharacterIDOpportunitiesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDOpportunitiesInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDOpportunitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/opportunities/][%d] getCharactersCharacterIdOpportunitiesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCharactersCharacterIDOpportunitiesInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDOpportunitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOpportunitiesServiceUnavailable creates a GetCharactersCharacterIDOpportunitiesServiceUnavailable with default headers values
func NewGetCharactersCharacterIDOpportunitiesServiceUnavailable() *GetCharactersCharacterIDOpportunitiesServiceUnavailable {
	return &GetCharactersCharacterIDOpportunitiesServiceUnavailable{}
}

/* GetCharactersCharacterIDOpportunitiesServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDOpportunitiesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDOpportunitiesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/opportunities/][%d] getCharactersCharacterIdOpportunitiesServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCharactersCharacterIDOpportunitiesServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDOpportunitiesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOpportunitiesGatewayTimeout creates a GetCharactersCharacterIDOpportunitiesGatewayTimeout with default headers values
func NewGetCharactersCharacterIDOpportunitiesGatewayTimeout() *GetCharactersCharacterIDOpportunitiesGatewayTimeout {
	return &GetCharactersCharacterIDOpportunitiesGatewayTimeout{}
}

/* GetCharactersCharacterIDOpportunitiesGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDOpportunitiesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDOpportunitiesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/opportunities/][%d] getCharactersCharacterIdOpportunitiesGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCharactersCharacterIDOpportunitiesGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDOpportunitiesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDOpportunitiesOKBodyItems0 get_characters_character_id_opportunities_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDOpportunitiesOKBodyItems0
*/
type GetCharactersCharacterIDOpportunitiesOKBodyItems0 struct {

	// get_characters_character_id_opportunities_completed_at
	//
	// completed_at string
	// Required: true
	// Format: date-time
	CompletedAt *strfmt.DateTime `json:"completed_at"`

	// get_characters_character_id_opportunities_task_id
	//
	// task_id integer
	// Required: true
	TaskID *int32 `json:"task_id"`
}

// Validate validates this get characters character ID opportunities o k body items0
func (o *GetCharactersCharacterIDOpportunitiesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCompletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTaskID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDOpportunitiesOKBodyItems0) validateCompletedAt(formats strfmt.Registry) error {

	if err := validate.Required("completed_at", "body", o.CompletedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("completed_at", "body", "date-time", o.CompletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDOpportunitiesOKBodyItems0) validateTaskID(formats strfmt.Registry) error {

	if err := validate.Required("task_id", "body", o.TaskID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get characters character ID opportunities o k body items0 based on context it is used
func (o *GetCharactersCharacterIDOpportunitiesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDOpportunitiesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDOpportunitiesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDOpportunitiesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
