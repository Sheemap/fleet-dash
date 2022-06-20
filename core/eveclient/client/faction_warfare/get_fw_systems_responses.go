// Code generated by go-swagger; DO NOT EDIT.

package faction_warfare

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

// GetFwSystemsReader is a Reader for the GetFwSystems structure.
type GetFwSystemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFwSystemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFwSystemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetFwSystemsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetFwSystemsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetFwSystemsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFwSystemsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFwSystemsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFwSystemsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFwSystemsOK creates a GetFwSystemsOK with default headers values
func NewGetFwSystemsOK() *GetFwSystemsOK {
	return &GetFwSystemsOK{}
}

/* GetFwSystemsOK describes a response with status code 200, with default header values.

All faction warfare solar systems
*/
type GetFwSystemsOK struct {

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

	Payload []*GetFwSystemsOKBodyItems0
}

func (o *GetFwSystemsOK) Error() string {
	return fmt.Sprintf("[GET /fw/systems/][%d] getFwSystemsOK  %+v", 200, o.Payload)
}
func (o *GetFwSystemsOK) GetPayload() []*GetFwSystemsOKBodyItems0 {
	return o.Payload
}

func (o *GetFwSystemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFwSystemsNotModified creates a GetFwSystemsNotModified with default headers values
func NewGetFwSystemsNotModified() *GetFwSystemsNotModified {
	return &GetFwSystemsNotModified{}
}

/* GetFwSystemsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetFwSystemsNotModified struct {

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

func (o *GetFwSystemsNotModified) Error() string {
	return fmt.Sprintf("[GET /fw/systems/][%d] getFwSystemsNotModified ", 304)
}

func (o *GetFwSystemsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFwSystemsBadRequest creates a GetFwSystemsBadRequest with default headers values
func NewGetFwSystemsBadRequest() *GetFwSystemsBadRequest {
	return &GetFwSystemsBadRequest{}
}

/* GetFwSystemsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetFwSystemsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetFwSystemsBadRequest) Error() string {
	return fmt.Sprintf("[GET /fw/systems/][%d] getFwSystemsBadRequest  %+v", 400, o.Payload)
}
func (o *GetFwSystemsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetFwSystemsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFwSystemsEnhanceYourCalm creates a GetFwSystemsEnhanceYourCalm with default headers values
func NewGetFwSystemsEnhanceYourCalm() *GetFwSystemsEnhanceYourCalm {
	return &GetFwSystemsEnhanceYourCalm{}
}

/* GetFwSystemsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetFwSystemsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetFwSystemsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /fw/systems/][%d] getFwSystemsEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetFwSystemsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetFwSystemsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFwSystemsInternalServerError creates a GetFwSystemsInternalServerError with default headers values
func NewGetFwSystemsInternalServerError() *GetFwSystemsInternalServerError {
	return &GetFwSystemsInternalServerError{}
}

/* GetFwSystemsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetFwSystemsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetFwSystemsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fw/systems/][%d] getFwSystemsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetFwSystemsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetFwSystemsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFwSystemsServiceUnavailable creates a GetFwSystemsServiceUnavailable with default headers values
func NewGetFwSystemsServiceUnavailable() *GetFwSystemsServiceUnavailable {
	return &GetFwSystemsServiceUnavailable{}
}

/* GetFwSystemsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetFwSystemsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetFwSystemsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /fw/systems/][%d] getFwSystemsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetFwSystemsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetFwSystemsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFwSystemsGatewayTimeout creates a GetFwSystemsGatewayTimeout with default headers values
func NewGetFwSystemsGatewayTimeout() *GetFwSystemsGatewayTimeout {
	return &GetFwSystemsGatewayTimeout{}
}

/* GetFwSystemsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetFwSystemsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetFwSystemsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /fw/systems/][%d] getFwSystemsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetFwSystemsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetFwSystemsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetFwSystemsOKBodyItems0 get_fw_systems_200_ok
//
// 200 ok object
swagger:model GetFwSystemsOKBodyItems0
*/
type GetFwSystemsOKBodyItems0 struct {

	// get_fw_systems_contested
	//
	// contested string
	// Required: true
	// Enum: [captured contested uncontested vulnerable]
	Contested *string `json:"contested"`

	// get_fw_systems_occupier_faction_id
	//
	// occupier_faction_id integer
	// Required: true
	OccupierFactionID *int32 `json:"occupier_faction_id"`

	// get_fw_systems_owner_faction_id
	//
	// owner_faction_id integer
	// Required: true
	OwnerFactionID *int32 `json:"owner_faction_id"`

	// get_fw_systems_solar_system_id
	//
	// solar_system_id integer
	// Required: true
	SolarSystemID *int32 `json:"solar_system_id"`

	// get_fw_systems_victory_points
	//
	// victory_points integer
	// Required: true
	VictoryPoints *int32 `json:"victory_points"`

	// get_fw_systems_victory_points_threshold
	//
	// victory_points_threshold integer
	// Required: true
	VictoryPointsThreshold *int32 `json:"victory_points_threshold"`
}

// Validate validates this get fw systems o k body items0
func (o *GetFwSystemsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateContested(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOccupierFactionID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOwnerFactionID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSolarSystemID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVictoryPoints(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVictoryPointsThreshold(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getFwSystemsOKBodyItems0TypeContestedPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["captured","contested","uncontested","vulnerable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getFwSystemsOKBodyItems0TypeContestedPropEnum = append(getFwSystemsOKBodyItems0TypeContestedPropEnum, v)
	}
}

const (

	// GetFwSystemsOKBodyItems0ContestedCaptured captures enum value "captured"
	GetFwSystemsOKBodyItems0ContestedCaptured string = "captured"

	// GetFwSystemsOKBodyItems0ContestedContested captures enum value "contested"
	GetFwSystemsOKBodyItems0ContestedContested string = "contested"

	// GetFwSystemsOKBodyItems0ContestedUncontested captures enum value "uncontested"
	GetFwSystemsOKBodyItems0ContestedUncontested string = "uncontested"

	// GetFwSystemsOKBodyItems0ContestedVulnerable captures enum value "vulnerable"
	GetFwSystemsOKBodyItems0ContestedVulnerable string = "vulnerable"
)

// prop value enum
func (o *GetFwSystemsOKBodyItems0) validateContestedEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getFwSystemsOKBodyItems0TypeContestedPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetFwSystemsOKBodyItems0) validateContested(formats strfmt.Registry) error {

	if err := validate.Required("contested", "body", o.Contested); err != nil {
		return err
	}

	// value enum
	if err := o.validateContestedEnum("contested", "body", *o.Contested); err != nil {
		return err
	}

	return nil
}

func (o *GetFwSystemsOKBodyItems0) validateOccupierFactionID(formats strfmt.Registry) error {

	if err := validate.Required("occupier_faction_id", "body", o.OccupierFactionID); err != nil {
		return err
	}

	return nil
}

func (o *GetFwSystemsOKBodyItems0) validateOwnerFactionID(formats strfmt.Registry) error {

	if err := validate.Required("owner_faction_id", "body", o.OwnerFactionID); err != nil {
		return err
	}

	return nil
}

func (o *GetFwSystemsOKBodyItems0) validateSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("solar_system_id", "body", o.SolarSystemID); err != nil {
		return err
	}

	return nil
}

func (o *GetFwSystemsOKBodyItems0) validateVictoryPoints(formats strfmt.Registry) error {

	if err := validate.Required("victory_points", "body", o.VictoryPoints); err != nil {
		return err
	}

	return nil
}

func (o *GetFwSystemsOKBodyItems0) validateVictoryPointsThreshold(formats strfmt.Registry) error {

	if err := validate.Required("victory_points_threshold", "body", o.VictoryPointsThreshold); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get fw systems o k body items0 based on context it is used
func (o *GetFwSystemsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetFwSystemsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFwSystemsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetFwSystemsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}