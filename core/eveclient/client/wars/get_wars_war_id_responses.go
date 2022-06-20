// Code generated by go-swagger; DO NOT EDIT.

package wars

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"fleet-dash-core/eveclient/models"
)

// GetWarsWarIDReader is a Reader for the GetWarsWarID structure.
type GetWarsWarIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWarsWarIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWarsWarIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetWarsWarIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetWarsWarIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetWarsWarIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetWarsWarIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWarsWarIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWarsWarIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWarsWarIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWarsWarIDOK creates a GetWarsWarIDOK with default headers values
func NewGetWarsWarIDOK() *GetWarsWarIDOK {
	return &GetWarsWarIDOK{}
}

/* GetWarsWarIDOK describes a response with status code 200, with default header values.

Details about a war
*/
type GetWarsWarIDOK struct {

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

	Payload *GetWarsWarIDOKBody
}

func (o *GetWarsWarIDOK) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/][%d] getWarsWarIdOK  %+v", 200, o.Payload)
}
func (o *GetWarsWarIDOK) GetPayload() *GetWarsWarIDOKBody {
	return o.Payload
}

func (o *GetWarsWarIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetWarsWarIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWarsWarIDNotModified creates a GetWarsWarIDNotModified with default headers values
func NewGetWarsWarIDNotModified() *GetWarsWarIDNotModified {
	return &GetWarsWarIDNotModified{}
}

/* GetWarsWarIDNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetWarsWarIDNotModified struct {

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

func (o *GetWarsWarIDNotModified) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/][%d] getWarsWarIdNotModified ", 304)
}

func (o *GetWarsWarIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetWarsWarIDBadRequest creates a GetWarsWarIDBadRequest with default headers values
func NewGetWarsWarIDBadRequest() *GetWarsWarIDBadRequest {
	return &GetWarsWarIDBadRequest{}
}

/* GetWarsWarIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetWarsWarIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetWarsWarIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/][%d] getWarsWarIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetWarsWarIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetWarsWarIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWarsWarIDEnhanceYourCalm creates a GetWarsWarIDEnhanceYourCalm with default headers values
func NewGetWarsWarIDEnhanceYourCalm() *GetWarsWarIDEnhanceYourCalm {
	return &GetWarsWarIDEnhanceYourCalm{}
}

/* GetWarsWarIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetWarsWarIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetWarsWarIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/][%d] getWarsWarIdEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetWarsWarIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetWarsWarIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWarsWarIDUnprocessableEntity creates a GetWarsWarIDUnprocessableEntity with default headers values
func NewGetWarsWarIDUnprocessableEntity() *GetWarsWarIDUnprocessableEntity {
	return &GetWarsWarIDUnprocessableEntity{}
}

/* GetWarsWarIDUnprocessableEntity describes a response with status code 422, with default header values.

War not found
*/
type GetWarsWarIDUnprocessableEntity struct {
	Payload *GetWarsWarIDUnprocessableEntityBody
}

func (o *GetWarsWarIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/][%d] getWarsWarIdUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *GetWarsWarIDUnprocessableEntity) GetPayload() *GetWarsWarIDUnprocessableEntityBody {
	return o.Payload
}

func (o *GetWarsWarIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWarsWarIDUnprocessableEntityBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWarsWarIDInternalServerError creates a GetWarsWarIDInternalServerError with default headers values
func NewGetWarsWarIDInternalServerError() *GetWarsWarIDInternalServerError {
	return &GetWarsWarIDInternalServerError{}
}

/* GetWarsWarIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetWarsWarIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetWarsWarIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/][%d] getWarsWarIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetWarsWarIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetWarsWarIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWarsWarIDServiceUnavailable creates a GetWarsWarIDServiceUnavailable with default headers values
func NewGetWarsWarIDServiceUnavailable() *GetWarsWarIDServiceUnavailable {
	return &GetWarsWarIDServiceUnavailable{}
}

/* GetWarsWarIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetWarsWarIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetWarsWarIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/][%d] getWarsWarIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetWarsWarIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetWarsWarIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWarsWarIDGatewayTimeout creates a GetWarsWarIDGatewayTimeout with default headers values
func NewGetWarsWarIDGatewayTimeout() *GetWarsWarIDGatewayTimeout {
	return &GetWarsWarIDGatewayTimeout{}
}

/* GetWarsWarIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetWarsWarIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetWarsWarIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/][%d] getWarsWarIdGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetWarsWarIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetWarsWarIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetWarsWarIDOKBody get_wars_war_id_ok
//
// 200 ok object
swagger:model GetWarsWarIDOKBody
*/
type GetWarsWarIDOKBody struct {

	// aggressor
	// Required: true
	Aggressor *GetWarsWarIDOKBodyAggressor `json:"aggressor"`

	// get_wars_war_id_allies
	//
	// allied corporations or alliances, each object contains either corporation_id or alliance_id
	// Max Items: 10000
	Allies []*GetWarsWarIDOKBodyAlliesItems0 `json:"allies"`

	// get_wars_war_id_declared
	//
	// Time that the war was declared
	// Required: true
	// Format: date-time
	Declared *strfmt.DateTime `json:"declared"`

	// defender
	// Required: true
	Defender *GetWarsWarIDOKBodyDefender `json:"defender"`

	// get_wars_war_id_finished
	//
	// Time the war ended and shooting was no longer allowed
	// Format: date-time
	Finished strfmt.DateTime `json:"finished,omitempty"`

	// get_wars_war_id_id
	//
	// ID of the specified war
	// Required: true
	ID *int32 `json:"id"`

	// get_wars_war_id_mutual
	//
	// Was the war declared mutual by both parties
	// Required: true
	Mutual *bool `json:"mutual"`

	// get_wars_war_id_open_for_allies
	//
	// Is the war currently open for allies or not
	// Required: true
	OpenForAllies *bool `json:"open_for_allies"`

	// get_wars_war_id_retracted
	//
	// Time the war was retracted but both sides could still shoot each other
	// Format: date-time
	Retracted strfmt.DateTime `json:"retracted,omitempty"`

	// get_wars_war_id_started
	//
	// Time when the war started and both sides could shoot each other
	// Format: date-time
	Started strfmt.DateTime `json:"started,omitempty"`
}

// Validate validates this get wars war ID o k body
func (o *GetWarsWarIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAggressor(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAllies(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDeclared(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDefender(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFinished(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMutual(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOpenForAllies(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRetracted(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStarted(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWarsWarIDOKBody) validateAggressor(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"aggressor", "body", o.Aggressor); err != nil {
		return err
	}

	if o.Aggressor != nil {
		if err := o.Aggressor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWarsWarIdOK" + "." + "aggressor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getWarsWarIdOK" + "." + "aggressor")
			}
			return err
		}
	}

	return nil
}

func (o *GetWarsWarIDOKBody) validateAllies(formats strfmt.Registry) error {
	if swag.IsZero(o.Allies) { // not required
		return nil
	}

	iAlliesSize := int64(len(o.Allies))

	if err := validate.MaxItems("getWarsWarIdOK"+"."+"allies", "body", iAlliesSize, 10000); err != nil {
		return err
	}

	for i := 0; i < len(o.Allies); i++ {
		if swag.IsZero(o.Allies[i]) { // not required
			continue
		}

		if o.Allies[i] != nil {
			if err := o.Allies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getWarsWarIdOK" + "." + "allies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getWarsWarIdOK" + "." + "allies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetWarsWarIDOKBody) validateDeclared(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"declared", "body", o.Declared); err != nil {
		return err
	}

	if err := validate.FormatOf("getWarsWarIdOK"+"."+"declared", "body", "date-time", o.Declared.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetWarsWarIDOKBody) validateDefender(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"defender", "body", o.Defender); err != nil {
		return err
	}

	if o.Defender != nil {
		if err := o.Defender.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWarsWarIdOK" + "." + "defender")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getWarsWarIdOK" + "." + "defender")
			}
			return err
		}
	}

	return nil
}

func (o *GetWarsWarIDOKBody) validateFinished(formats strfmt.Registry) error {
	if swag.IsZero(o.Finished) { // not required
		return nil
	}

	if err := validate.FormatOf("getWarsWarIdOK"+"."+"finished", "body", "date-time", o.Finished.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetWarsWarIDOKBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *GetWarsWarIDOKBody) validateMutual(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"mutual", "body", o.Mutual); err != nil {
		return err
	}

	return nil
}

func (o *GetWarsWarIDOKBody) validateOpenForAllies(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"open_for_allies", "body", o.OpenForAllies); err != nil {
		return err
	}

	return nil
}

func (o *GetWarsWarIDOKBody) validateRetracted(formats strfmt.Registry) error {
	if swag.IsZero(o.Retracted) { // not required
		return nil
	}

	if err := validate.FormatOf("getWarsWarIdOK"+"."+"retracted", "body", "date-time", o.Retracted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetWarsWarIDOKBody) validateStarted(formats strfmt.Registry) error {
	if swag.IsZero(o.Started) { // not required
		return nil
	}

	if err := validate.FormatOf("getWarsWarIdOK"+"."+"started", "body", "date-time", o.Started.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get wars war ID o k body based on the context it is used
func (o *GetWarsWarIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAggressor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateAllies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDefender(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWarsWarIDOKBody) contextValidateAggressor(ctx context.Context, formats strfmt.Registry) error {

	if o.Aggressor != nil {
		if err := o.Aggressor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWarsWarIdOK" + "." + "aggressor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getWarsWarIdOK" + "." + "aggressor")
			}
			return err
		}
	}

	return nil
}

func (o *GetWarsWarIDOKBody) contextValidateAllies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Allies); i++ {

		if o.Allies[i] != nil {
			if err := o.Allies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getWarsWarIdOK" + "." + "allies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getWarsWarIdOK" + "." + "allies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetWarsWarIDOKBody) contextValidateDefender(ctx context.Context, formats strfmt.Registry) error {

	if o.Defender != nil {
		if err := o.Defender.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWarsWarIdOK" + "." + "defender")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getWarsWarIdOK" + "." + "defender")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWarsWarIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWarsWarIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetWarsWarIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWarsWarIDOKBodyAggressor get_wars_war_id_aggressor
//
// The aggressor corporation or alliance that declared this war, only contains either corporation_id or alliance_id
swagger:model GetWarsWarIDOKBodyAggressor
*/
type GetWarsWarIDOKBodyAggressor struct {

	// get_wars_war_id_alliance_id
	//
	// Alliance ID if and only if the aggressor is an alliance
	AllianceID int32 `json:"alliance_id,omitempty"`

	// get_wars_war_id_corporation_id
	//
	// Corporation ID if and only if the aggressor is a corporation
	CorporationID int32 `json:"corporation_id,omitempty"`

	// get_wars_war_id_isk_destroyed
	//
	// ISK value of ships the aggressor has destroyed
	// Required: true
	IskDestroyed *float32 `json:"isk_destroyed"`

	// get_wars_war_id_ships_killed
	//
	// The number of ships the aggressor has killed
	// Required: true
	ShipsKilled *int32 `json:"ships_killed"`
}

// Validate validates this get wars war ID o k body aggressor
func (o *GetWarsWarIDOKBodyAggressor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIskDestroyed(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateShipsKilled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWarsWarIDOKBodyAggressor) validateIskDestroyed(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"aggressor"+"."+"isk_destroyed", "body", o.IskDestroyed); err != nil {
		return err
	}

	return nil
}

func (o *GetWarsWarIDOKBodyAggressor) validateShipsKilled(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"aggressor"+"."+"ships_killed", "body", o.ShipsKilled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get wars war ID o k body aggressor based on context it is used
func (o *GetWarsWarIDOKBodyAggressor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWarsWarIDOKBodyAggressor) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWarsWarIDOKBodyAggressor) UnmarshalBinary(b []byte) error {
	var res GetWarsWarIDOKBodyAggressor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWarsWarIDOKBodyAlliesItems0 get_wars_war_id_ally
//
// ally object
swagger:model GetWarsWarIDOKBodyAlliesItems0
*/
type GetWarsWarIDOKBodyAlliesItems0 struct {

	// get_wars_war_id_ally_alliance_id
	//
	// Alliance ID if and only if this ally is an alliance
	AllianceID int32 `json:"alliance_id,omitempty"`

	// get_wars_war_id_ally_corporation_id
	//
	// Corporation ID if and only if this ally is a corporation
	CorporationID int32 `json:"corporation_id,omitempty"`
}

// Validate validates this get wars war ID o k body allies items0
func (o *GetWarsWarIDOKBodyAlliesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get wars war ID o k body allies items0 based on context it is used
func (o *GetWarsWarIDOKBodyAlliesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWarsWarIDOKBodyAlliesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWarsWarIDOKBodyAlliesItems0) UnmarshalBinary(b []byte) error {
	var res GetWarsWarIDOKBodyAlliesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWarsWarIDOKBodyDefender get_wars_war_id_defender
//
// The defending corporation or alliance that declared this war, only contains either corporation_id or alliance_id
swagger:model GetWarsWarIDOKBodyDefender
*/
type GetWarsWarIDOKBodyDefender struct {

	// get_wars_war_id_defender_alliance_id
	//
	// Alliance ID if and only if the defender is an alliance
	AllianceID int32 `json:"alliance_id,omitempty"`

	// get_wars_war_id_defender_corporation_id
	//
	// Corporation ID if and only if the defender is a corporation
	CorporationID int32 `json:"corporation_id,omitempty"`

	// get_wars_war_id_defender_isk_destroyed
	//
	// ISK value of ships the defender has killed
	// Required: true
	IskDestroyed *float32 `json:"isk_destroyed"`

	// get_wars_war_id_defender_ships_killed
	//
	// The number of ships the defender has killed
	// Required: true
	ShipsKilled *int32 `json:"ships_killed"`
}

// Validate validates this get wars war ID o k body defender
func (o *GetWarsWarIDOKBodyDefender) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIskDestroyed(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateShipsKilled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWarsWarIDOKBodyDefender) validateIskDestroyed(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"defender"+"."+"isk_destroyed", "body", o.IskDestroyed); err != nil {
		return err
	}

	return nil
}

func (o *GetWarsWarIDOKBodyDefender) validateShipsKilled(formats strfmt.Registry) error {

	if err := validate.Required("getWarsWarIdOK"+"."+"defender"+"."+"ships_killed", "body", o.ShipsKilled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get wars war ID o k body defender based on context it is used
func (o *GetWarsWarIDOKBodyDefender) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWarsWarIDOKBodyDefender) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWarsWarIDOKBodyDefender) UnmarshalBinary(b []byte) error {
	var res GetWarsWarIDOKBodyDefender
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWarsWarIDUnprocessableEntityBody get_wars_war_id_unprocessable_entity
//
// Unprocessable entity
swagger:model GetWarsWarIDUnprocessableEntityBody
*/
type GetWarsWarIDUnprocessableEntityBody struct {

	// get_wars_war_id_422_unprocessable_entity
	//
	// Unprocessable entity message
	Error string `json:"error,omitempty"`
}

// Validate validates this get wars war ID unprocessable entity body
func (o *GetWarsWarIDUnprocessableEntityBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get wars war ID unprocessable entity body based on context it is used
func (o *GetWarsWarIDUnprocessableEntityBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWarsWarIDUnprocessableEntityBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWarsWarIDUnprocessableEntityBody) UnmarshalBinary(b []byte) error {
	var res GetWarsWarIDUnprocessableEntityBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}