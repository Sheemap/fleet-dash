// Code generated by go-swagger; DO NOT EDIT.

package incursions

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

// GetIncursionsReader is a Reader for the GetIncursions structure.
type GetIncursionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIncursionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIncursionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetIncursionsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetIncursionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetIncursionsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIncursionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIncursionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIncursionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIncursionsOK creates a GetIncursionsOK with default headers values
func NewGetIncursionsOK() *GetIncursionsOK {
	return &GetIncursionsOK{}
}

/* GetIncursionsOK describes a response with status code 200, with default header values.

A list of incursions
*/
type GetIncursionsOK struct {

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

	Payload []*GetIncursionsOKBodyItems0
}

func (o *GetIncursionsOK) Error() string {
	return fmt.Sprintf("[GET /incursions/][%d] getIncursionsOK  %+v", 200, o.Payload)
}
func (o *GetIncursionsOK) GetPayload() []*GetIncursionsOKBodyItems0 {
	return o.Payload
}

func (o *GetIncursionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIncursionsNotModified creates a GetIncursionsNotModified with default headers values
func NewGetIncursionsNotModified() *GetIncursionsNotModified {
	return &GetIncursionsNotModified{}
}

/* GetIncursionsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetIncursionsNotModified struct {

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

func (o *GetIncursionsNotModified) Error() string {
	return fmt.Sprintf("[GET /incursions/][%d] getIncursionsNotModified ", 304)
}

func (o *GetIncursionsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIncursionsBadRequest creates a GetIncursionsBadRequest with default headers values
func NewGetIncursionsBadRequest() *GetIncursionsBadRequest {
	return &GetIncursionsBadRequest{}
}

/* GetIncursionsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetIncursionsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetIncursionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /incursions/][%d] getIncursionsBadRequest  %+v", 400, o.Payload)
}
func (o *GetIncursionsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetIncursionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIncursionsEnhanceYourCalm creates a GetIncursionsEnhanceYourCalm with default headers values
func NewGetIncursionsEnhanceYourCalm() *GetIncursionsEnhanceYourCalm {
	return &GetIncursionsEnhanceYourCalm{}
}

/* GetIncursionsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetIncursionsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetIncursionsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /incursions/][%d] getIncursionsEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetIncursionsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetIncursionsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIncursionsInternalServerError creates a GetIncursionsInternalServerError with default headers values
func NewGetIncursionsInternalServerError() *GetIncursionsInternalServerError {
	return &GetIncursionsInternalServerError{}
}

/* GetIncursionsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetIncursionsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetIncursionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /incursions/][%d] getIncursionsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetIncursionsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetIncursionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIncursionsServiceUnavailable creates a GetIncursionsServiceUnavailable with default headers values
func NewGetIncursionsServiceUnavailable() *GetIncursionsServiceUnavailable {
	return &GetIncursionsServiceUnavailable{}
}

/* GetIncursionsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetIncursionsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetIncursionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /incursions/][%d] getIncursionsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetIncursionsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetIncursionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIncursionsGatewayTimeout creates a GetIncursionsGatewayTimeout with default headers values
func NewGetIncursionsGatewayTimeout() *GetIncursionsGatewayTimeout {
	return &GetIncursionsGatewayTimeout{}
}

/* GetIncursionsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetIncursionsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetIncursionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /incursions/][%d] getIncursionsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetIncursionsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetIncursionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetIncursionsOKBodyItems0 get_incursions_200_ok
//
// 200 ok object
swagger:model GetIncursionsOKBodyItems0
*/
type GetIncursionsOKBodyItems0 struct {

	// get_incursions_constellation_id
	//
	// The constellation id in which this incursion takes place
	// Required: true
	ConstellationID *int32 `json:"constellation_id"`

	// get_incursions_faction_id
	//
	// The attacking faction's id
	// Required: true
	FactionID *int32 `json:"faction_id"`

	// get_incursions_has_boss
	//
	// Whether the final encounter has boss or not
	// Required: true
	HasBoss *bool `json:"has_boss"`

	// get_incursions_infested_solar_systems
	//
	// A list of infested solar system ids that are a part of this incursion
	// Required: true
	// Max Items: 100
	InfestedSolarSystems []int32 `json:"infested_solar_systems"`

	// get_incursions_influence
	//
	// Influence of this incursion as a float from 0 to 1
	// Required: true
	Influence *float32 `json:"influence"`

	// get_incursions_staging_solar_system_id
	//
	// Staging solar system for this incursion
	// Required: true
	StagingSolarSystemID *int32 `json:"staging_solar_system_id"`

	// get_incursions_state
	//
	// The state of this incursion
	// Required: true
	// Enum: [withdrawing mobilizing established]
	State *string `json:"state"`

	// get_incursions_type
	//
	// The type of this incursion
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this get incursions o k body items0
func (o *GetIncursionsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConstellationID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFactionID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateHasBoss(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateInfestedSolarSystems(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateInfluence(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStagingSolarSystemID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetIncursionsOKBodyItems0) validateConstellationID(formats strfmt.Registry) error {

	if err := validate.Required("constellation_id", "body", o.ConstellationID); err != nil {
		return err
	}

	return nil
}

func (o *GetIncursionsOKBodyItems0) validateFactionID(formats strfmt.Registry) error {

	if err := validate.Required("faction_id", "body", o.FactionID); err != nil {
		return err
	}

	return nil
}

func (o *GetIncursionsOKBodyItems0) validateHasBoss(formats strfmt.Registry) error {

	if err := validate.Required("has_boss", "body", o.HasBoss); err != nil {
		return err
	}

	return nil
}

func (o *GetIncursionsOKBodyItems0) validateInfestedSolarSystems(formats strfmt.Registry) error {

	if err := validate.Required("infested_solar_systems", "body", o.InfestedSolarSystems); err != nil {
		return err
	}

	iInfestedSolarSystemsSize := int64(len(o.InfestedSolarSystems))

	if err := validate.MaxItems("infested_solar_systems", "body", iInfestedSolarSystemsSize, 100); err != nil {
		return err
	}

	return nil
}

func (o *GetIncursionsOKBodyItems0) validateInfluence(formats strfmt.Registry) error {

	if err := validate.Required("influence", "body", o.Influence); err != nil {
		return err
	}

	return nil
}

func (o *GetIncursionsOKBodyItems0) validateStagingSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("staging_solar_system_id", "body", o.StagingSolarSystemID); err != nil {
		return err
	}

	return nil
}

var getIncursionsOKBodyItems0TypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["withdrawing","mobilizing","established"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getIncursionsOKBodyItems0TypeStatePropEnum = append(getIncursionsOKBodyItems0TypeStatePropEnum, v)
	}
}

const (

	// GetIncursionsOKBodyItems0StateWithdrawing captures enum value "withdrawing"
	GetIncursionsOKBodyItems0StateWithdrawing string = "withdrawing"

	// GetIncursionsOKBodyItems0StateMobilizing captures enum value "mobilizing"
	GetIncursionsOKBodyItems0StateMobilizing string = "mobilizing"

	// GetIncursionsOKBodyItems0StateEstablished captures enum value "established"
	GetIncursionsOKBodyItems0StateEstablished string = "established"
)

// prop value enum
func (o *GetIncursionsOKBodyItems0) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getIncursionsOKBodyItems0TypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetIncursionsOKBodyItems0) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", o.State); err != nil {
		return err
	}

	// value enum
	if err := o.validateStateEnum("state", "body", *o.State); err != nil {
		return err
	}

	return nil
}

func (o *GetIncursionsOKBodyItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get incursions o k body items0 based on context it is used
func (o *GetIncursionsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetIncursionsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIncursionsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetIncursionsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
