// Code generated by go-swagger; DO NOT EDIT.

package corporation

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

// GetCorporationsCorporationIDMembertrackingReader is a Reader for the GetCorporationsCorporationIDMembertracking structure.
type GetCorporationsCorporationIDMembertrackingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDMembertrackingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationsCorporationIDMembertrackingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationsCorporationIDMembertrackingNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationsCorporationIDMembertrackingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationsCorporationIDMembertrackingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationsCorporationIDMembertrackingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationsCorporationIDMembertrackingEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationsCorporationIDMembertrackingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationsCorporationIDMembertrackingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationsCorporationIDMembertrackingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDMembertrackingOK creates a GetCorporationsCorporationIDMembertrackingOK with default headers values
func NewGetCorporationsCorporationIDMembertrackingOK() *GetCorporationsCorporationIDMembertrackingOK {
	return &GetCorporationsCorporationIDMembertrackingOK{}
}

/* GetCorporationsCorporationIDMembertrackingOK describes a response with status code 200, with default header values.

List of member character IDs
*/
type GetCorporationsCorporationIDMembertrackingOK struct {

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

	Payload []*GetCorporationsCorporationIDMembertrackingOKBodyItems0
}

func (o *GetCorporationsCorporationIDMembertrackingOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/membertracking/][%d] getCorporationsCorporationIdMembertrackingOK  %+v", 200, o.Payload)
}
func (o *GetCorporationsCorporationIDMembertrackingOK) GetPayload() []*GetCorporationsCorporationIDMembertrackingOKBodyItems0 {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembertrackingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDMembertrackingNotModified creates a GetCorporationsCorporationIDMembertrackingNotModified with default headers values
func NewGetCorporationsCorporationIDMembertrackingNotModified() *GetCorporationsCorporationIDMembertrackingNotModified {
	return &GetCorporationsCorporationIDMembertrackingNotModified{}
}

/* GetCorporationsCorporationIDMembertrackingNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationsCorporationIDMembertrackingNotModified struct {

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

func (o *GetCorporationsCorporationIDMembertrackingNotModified) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/membertracking/][%d] getCorporationsCorporationIdMembertrackingNotModified ", 304)
}

func (o *GetCorporationsCorporationIDMembertrackingNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDMembertrackingBadRequest creates a GetCorporationsCorporationIDMembertrackingBadRequest with default headers values
func NewGetCorporationsCorporationIDMembertrackingBadRequest() *GetCorporationsCorporationIDMembertrackingBadRequest {
	return &GetCorporationsCorporationIDMembertrackingBadRequest{}
}

/* GetCorporationsCorporationIDMembertrackingBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationsCorporationIDMembertrackingBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDMembertrackingBadRequest) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/membertracking/][%d] getCorporationsCorporationIdMembertrackingBadRequest  %+v", 400, o.Payload)
}
func (o *GetCorporationsCorporationIDMembertrackingBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembertrackingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembertrackingUnauthorized creates a GetCorporationsCorporationIDMembertrackingUnauthorized with default headers values
func NewGetCorporationsCorporationIDMembertrackingUnauthorized() *GetCorporationsCorporationIDMembertrackingUnauthorized {
	return &GetCorporationsCorporationIDMembertrackingUnauthorized{}
}

/* GetCorporationsCorporationIDMembertrackingUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDMembertrackingUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDMembertrackingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/membertracking/][%d] getCorporationsCorporationIdMembertrackingUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCorporationsCorporationIDMembertrackingUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembertrackingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembertrackingForbidden creates a GetCorporationsCorporationIDMembertrackingForbidden with default headers values
func NewGetCorporationsCorporationIDMembertrackingForbidden() *GetCorporationsCorporationIDMembertrackingForbidden {
	return &GetCorporationsCorporationIDMembertrackingForbidden{}
}

/* GetCorporationsCorporationIDMembertrackingForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationsCorporationIDMembertrackingForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDMembertrackingForbidden) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/membertracking/][%d] getCorporationsCorporationIdMembertrackingForbidden  %+v", 403, o.Payload)
}
func (o *GetCorporationsCorporationIDMembertrackingForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembertrackingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembertrackingEnhanceYourCalm creates a GetCorporationsCorporationIDMembertrackingEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDMembertrackingEnhanceYourCalm() *GetCorporationsCorporationIDMembertrackingEnhanceYourCalm {
	return &GetCorporationsCorporationIDMembertrackingEnhanceYourCalm{}
}

/* GetCorporationsCorporationIDMembertrackingEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationsCorporationIDMembertrackingEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDMembertrackingEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/membertracking/][%d] getCorporationsCorporationIdMembertrackingEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCorporationsCorporationIDMembertrackingEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembertrackingEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembertrackingInternalServerError creates a GetCorporationsCorporationIDMembertrackingInternalServerError with default headers values
func NewGetCorporationsCorporationIDMembertrackingInternalServerError() *GetCorporationsCorporationIDMembertrackingInternalServerError {
	return &GetCorporationsCorporationIDMembertrackingInternalServerError{}
}

/* GetCorporationsCorporationIDMembertrackingInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationsCorporationIDMembertrackingInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDMembertrackingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/membertracking/][%d] getCorporationsCorporationIdMembertrackingInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCorporationsCorporationIDMembertrackingInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembertrackingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembertrackingServiceUnavailable creates a GetCorporationsCorporationIDMembertrackingServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDMembertrackingServiceUnavailable() *GetCorporationsCorporationIDMembertrackingServiceUnavailable {
	return &GetCorporationsCorporationIDMembertrackingServiceUnavailable{}
}

/* GetCorporationsCorporationIDMembertrackingServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDMembertrackingServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDMembertrackingServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/membertracking/][%d] getCorporationsCorporationIdMembertrackingServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCorporationsCorporationIDMembertrackingServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembertrackingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembertrackingGatewayTimeout creates a GetCorporationsCorporationIDMembertrackingGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDMembertrackingGatewayTimeout() *GetCorporationsCorporationIDMembertrackingGatewayTimeout {
	return &GetCorporationsCorporationIDMembertrackingGatewayTimeout{}
}

/* GetCorporationsCorporationIDMembertrackingGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDMembertrackingGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDMembertrackingGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/membertracking/][%d] getCorporationsCorporationIdMembertrackingGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCorporationsCorporationIDMembertrackingGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembertrackingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDMembertrackingOKBodyItems0 get_corporations_corporation_id_membertracking_200_ok
//
// 200 ok object
swagger:model GetCorporationsCorporationIDMembertrackingOKBodyItems0
*/
type GetCorporationsCorporationIDMembertrackingOKBodyItems0 struct {

	// get_corporations_corporation_id_membertracking_base_id
	//
	// base_id integer
	BaseID int32 `json:"base_id,omitempty"`

	// get_corporations_corporation_id_membertracking_character_id
	//
	// character_id integer
	// Required: true
	CharacterID *int32 `json:"character_id"`

	// get_corporations_corporation_id_membertracking_location_id
	//
	// location_id integer
	LocationID int64 `json:"location_id,omitempty"`

	// get_corporations_corporation_id_membertracking_logoff_date
	//
	// logoff_date string
	// Format: date-time
	LogoffDate strfmt.DateTime `json:"logoff_date,omitempty"`

	// get_corporations_corporation_id_membertracking_logon_date
	//
	// logon_date string
	// Format: date-time
	LogonDate strfmt.DateTime `json:"logon_date,omitempty"`

	// get_corporations_corporation_id_membertracking_ship_type_id
	//
	// ship_type_id integer
	ShipTypeID int32 `json:"ship_type_id,omitempty"`

	// get_corporations_corporation_id_membertracking_start_date
	//
	// start_date string
	// Format: date-time
	StartDate strfmt.DateTime `json:"start_date,omitempty"`
}

// Validate validates this get corporations corporation ID membertracking o k body items0
func (o *GetCorporationsCorporationIDMembertrackingOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCharacterID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLogoffDate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLogonDate(formats); err != nil {
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

func (o *GetCorporationsCorporationIDMembertrackingOKBodyItems0) validateCharacterID(formats strfmt.Registry) error {

	if err := validate.Required("character_id", "body", o.CharacterID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDMembertrackingOKBodyItems0) validateLogoffDate(formats strfmt.Registry) error {
	if swag.IsZero(o.LogoffDate) { // not required
		return nil
	}

	if err := validate.FormatOf("logoff_date", "body", "date-time", o.LogoffDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDMembertrackingOKBodyItems0) validateLogonDate(formats strfmt.Registry) error {
	if swag.IsZero(o.LogonDate) { // not required
		return nil
	}

	if err := validate.FormatOf("logon_date", "body", "date-time", o.LogonDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDMembertrackingOKBodyItems0) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(o.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("start_date", "body", "date-time", o.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get corporations corporation ID membertracking o k body items0 based on context it is used
func (o *GetCorporationsCorporationIDMembertrackingOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDMembertrackingOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDMembertrackingOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDMembertrackingOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
