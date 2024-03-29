// Code generated by go-swagger; DO NOT EDIT.

package status

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

// GetStatusReader is a Reader for the GetStatus structure.
type GetStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetStatusNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetStatusEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetStatusServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetStatusGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStatusOK creates a GetStatusOK with default headers values
func NewGetStatusOK() *GetStatusOK {
	return &GetStatusOK{}
}

/* GetStatusOK describes a response with status code 200, with default header values.

Server status
*/
type GetStatusOK struct {

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

	Payload *GetStatusOKBody
}

func (o *GetStatusOK) Error() string {
	return fmt.Sprintf("[GET /status/][%d] getStatusOK  %+v", 200, o.Payload)
}
func (o *GetStatusOK) GetPayload() *GetStatusOKBody {
	return o.Payload
}

func (o *GetStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetStatusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatusNotModified creates a GetStatusNotModified with default headers values
func NewGetStatusNotModified() *GetStatusNotModified {
	return &GetStatusNotModified{}
}

/* GetStatusNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetStatusNotModified struct {

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

func (o *GetStatusNotModified) Error() string {
	return fmt.Sprintf("[GET /status/][%d] getStatusNotModified ", 304)
}

func (o *GetStatusNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetStatusBadRequest creates a GetStatusBadRequest with default headers values
func NewGetStatusBadRequest() *GetStatusBadRequest {
	return &GetStatusBadRequest{}
}

/* GetStatusBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetStatusBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /status/][%d] getStatusBadRequest  %+v", 400, o.Payload)
}
func (o *GetStatusBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatusEnhanceYourCalm creates a GetStatusEnhanceYourCalm with default headers values
func NewGetStatusEnhanceYourCalm() *GetStatusEnhanceYourCalm {
	return &GetStatusEnhanceYourCalm{}
}

/* GetStatusEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetStatusEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetStatusEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /status/][%d] getStatusEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetStatusEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetStatusEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatusInternalServerError creates a GetStatusInternalServerError with default headers values
func NewGetStatusInternalServerError() *GetStatusInternalServerError {
	return &GetStatusInternalServerError{}
}

/* GetStatusInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetStatusInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /status/][%d] getStatusInternalServerError  %+v", 500, o.Payload)
}
func (o *GetStatusInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatusServiceUnavailable creates a GetStatusServiceUnavailable with default headers values
func NewGetStatusServiceUnavailable() *GetStatusServiceUnavailable {
	return &GetStatusServiceUnavailable{}
}

/* GetStatusServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetStatusServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetStatusServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /status/][%d] getStatusServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetStatusServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetStatusServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatusGatewayTimeout creates a GetStatusGatewayTimeout with default headers values
func NewGetStatusGatewayTimeout() *GetStatusGatewayTimeout {
	return &GetStatusGatewayTimeout{}
}

/* GetStatusGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetStatusGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetStatusGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /status/][%d] getStatusGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetStatusGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetStatusGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetStatusOKBody get_status_ok
//
// 200 ok object
swagger:model GetStatusOKBody
*/
type GetStatusOKBody struct {

	// get_status_players
	//
	// Current online player count
	// Required: true
	Players *int64 `json:"players"`

	// get_status_server_version
	//
	// Running version as string
	// Required: true
	ServerVersion *string `json:"server_version"`

	// get_status_start_time
	//
	// Server start timestamp
	// Required: true
	// Format: date-time
	StartTime *strfmt.DateTime `json:"start_time"`

	// get_status_vip
	//
	// If the server is in VIP mode
	Vip bool `json:"vip,omitempty"`
}

// Validate validates this get status o k body
func (o *GetStatusOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePlayers(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateServerVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStatusOKBody) validatePlayers(formats strfmt.Registry) error {

	if err := validate.Required("getStatusOK"+"."+"players", "body", o.Players); err != nil {
		return err
	}

	return nil
}

func (o *GetStatusOKBody) validateServerVersion(formats strfmt.Registry) error {

	if err := validate.Required("getStatusOK"+"."+"server_version", "body", o.ServerVersion); err != nil {
		return err
	}

	return nil
}

func (o *GetStatusOKBody) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("getStatusOK"+"."+"start_time", "body", o.StartTime); err != nil {
		return err
	}

	if err := validate.FormatOf("getStatusOK"+"."+"start_time", "body", "date-time", o.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get status o k body based on context it is used
func (o *GetStatusOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetStatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStatusOKBody) UnmarshalBinary(b []byte) error {
	var res GetStatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
