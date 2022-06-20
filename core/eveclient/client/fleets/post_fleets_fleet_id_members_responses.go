// Code generated by go-swagger; DO NOT EDIT.

package fleets

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

// PostFleetsFleetIDMembersReader is a Reader for the PostFleetsFleetIDMembers structure.
type PostFleetsFleetIDMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFleetsFleetIDMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostFleetsFleetIDMembersNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostFleetsFleetIDMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostFleetsFleetIDMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostFleetsFleetIDMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostFleetsFleetIDMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewPostFleetsFleetIDMembersEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostFleetsFleetIDMembersUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostFleetsFleetIDMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostFleetsFleetIDMembersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostFleetsFleetIDMembersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostFleetsFleetIDMembersNoContent creates a PostFleetsFleetIDMembersNoContent with default headers values
func NewPostFleetsFleetIDMembersNoContent() *PostFleetsFleetIDMembersNoContent {
	return &PostFleetsFleetIDMembersNoContent{}
}

/* PostFleetsFleetIDMembersNoContent describes a response with status code 204, with default header values.

Fleet invitation sent
*/
type PostFleetsFleetIDMembersNoContent struct {
}

func (o *PostFleetsFleetIDMembersNoContent) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/members/][%d] postFleetsFleetIdMembersNoContent ", 204)
}

func (o *PostFleetsFleetIDMembersNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostFleetsFleetIDMembersBadRequest creates a PostFleetsFleetIDMembersBadRequest with default headers values
func NewPostFleetsFleetIDMembersBadRequest() *PostFleetsFleetIDMembersBadRequest {
	return &PostFleetsFleetIDMembersBadRequest{}
}

/* PostFleetsFleetIDMembersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostFleetsFleetIDMembersBadRequest struct {
	Payload *models.BadRequest
}

func (o *PostFleetsFleetIDMembersBadRequest) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/members/][%d] postFleetsFleetIdMembersBadRequest  %+v", 400, o.Payload)
}
func (o *PostFleetsFleetIDMembersBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *PostFleetsFleetIDMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDMembersUnauthorized creates a PostFleetsFleetIDMembersUnauthorized with default headers values
func NewPostFleetsFleetIDMembersUnauthorized() *PostFleetsFleetIDMembersUnauthorized {
	return &PostFleetsFleetIDMembersUnauthorized{}
}

/* PostFleetsFleetIDMembersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostFleetsFleetIDMembersUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *PostFleetsFleetIDMembersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/members/][%d] postFleetsFleetIdMembersUnauthorized  %+v", 401, o.Payload)
}
func (o *PostFleetsFleetIDMembersUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *PostFleetsFleetIDMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDMembersForbidden creates a PostFleetsFleetIDMembersForbidden with default headers values
func NewPostFleetsFleetIDMembersForbidden() *PostFleetsFleetIDMembersForbidden {
	return &PostFleetsFleetIDMembersForbidden{}
}

/* PostFleetsFleetIDMembersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostFleetsFleetIDMembersForbidden struct {
	Payload *models.Forbidden
}

func (o *PostFleetsFleetIDMembersForbidden) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/members/][%d] postFleetsFleetIdMembersForbidden  %+v", 403, o.Payload)
}
func (o *PostFleetsFleetIDMembersForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *PostFleetsFleetIDMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDMembersNotFound creates a PostFleetsFleetIDMembersNotFound with default headers values
func NewPostFleetsFleetIDMembersNotFound() *PostFleetsFleetIDMembersNotFound {
	return &PostFleetsFleetIDMembersNotFound{}
}

/* PostFleetsFleetIDMembersNotFound describes a response with status code 404, with default header values.

The fleet does not exist or you don't have access to it
*/
type PostFleetsFleetIDMembersNotFound struct {
	Payload *PostFleetsFleetIDMembersNotFoundBody
}

func (o *PostFleetsFleetIDMembersNotFound) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/members/][%d] postFleetsFleetIdMembersNotFound  %+v", 404, o.Payload)
}
func (o *PostFleetsFleetIDMembersNotFound) GetPayload() *PostFleetsFleetIDMembersNotFoundBody {
	return o.Payload
}

func (o *PostFleetsFleetIDMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostFleetsFleetIDMembersNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDMembersEnhanceYourCalm creates a PostFleetsFleetIDMembersEnhanceYourCalm with default headers values
func NewPostFleetsFleetIDMembersEnhanceYourCalm() *PostFleetsFleetIDMembersEnhanceYourCalm {
	return &PostFleetsFleetIDMembersEnhanceYourCalm{}
}

/* PostFleetsFleetIDMembersEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type PostFleetsFleetIDMembersEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *PostFleetsFleetIDMembersEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/members/][%d] postFleetsFleetIdMembersEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *PostFleetsFleetIDMembersEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *PostFleetsFleetIDMembersEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDMembersUnprocessableEntity creates a PostFleetsFleetIDMembersUnprocessableEntity with default headers values
func NewPostFleetsFleetIDMembersUnprocessableEntity() *PostFleetsFleetIDMembersUnprocessableEntity {
	return &PostFleetsFleetIDMembersUnprocessableEntity{}
}

/* PostFleetsFleetIDMembersUnprocessableEntity describes a response with status code 422, with default header values.

Errors in invitation
*/
type PostFleetsFleetIDMembersUnprocessableEntity struct {
	Payload *PostFleetsFleetIDMembersUnprocessableEntityBody
}

func (o *PostFleetsFleetIDMembersUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/members/][%d] postFleetsFleetIdMembersUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PostFleetsFleetIDMembersUnprocessableEntity) GetPayload() *PostFleetsFleetIDMembersUnprocessableEntityBody {
	return o.Payload
}

func (o *PostFleetsFleetIDMembersUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostFleetsFleetIDMembersUnprocessableEntityBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDMembersInternalServerError creates a PostFleetsFleetIDMembersInternalServerError with default headers values
func NewPostFleetsFleetIDMembersInternalServerError() *PostFleetsFleetIDMembersInternalServerError {
	return &PostFleetsFleetIDMembersInternalServerError{}
}

/* PostFleetsFleetIDMembersInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostFleetsFleetIDMembersInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *PostFleetsFleetIDMembersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/members/][%d] postFleetsFleetIdMembersInternalServerError  %+v", 500, o.Payload)
}
func (o *PostFleetsFleetIDMembersInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *PostFleetsFleetIDMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDMembersServiceUnavailable creates a PostFleetsFleetIDMembersServiceUnavailable with default headers values
func NewPostFleetsFleetIDMembersServiceUnavailable() *PostFleetsFleetIDMembersServiceUnavailable {
	return &PostFleetsFleetIDMembersServiceUnavailable{}
}

/* PostFleetsFleetIDMembersServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type PostFleetsFleetIDMembersServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *PostFleetsFleetIDMembersServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/members/][%d] postFleetsFleetIdMembersServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PostFleetsFleetIDMembersServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *PostFleetsFleetIDMembersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDMembersGatewayTimeout creates a PostFleetsFleetIDMembersGatewayTimeout with default headers values
func NewPostFleetsFleetIDMembersGatewayTimeout() *PostFleetsFleetIDMembersGatewayTimeout {
	return &PostFleetsFleetIDMembersGatewayTimeout{}
}

/* PostFleetsFleetIDMembersGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type PostFleetsFleetIDMembersGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *PostFleetsFleetIDMembersGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/members/][%d] postFleetsFleetIdMembersGatewayTimeout  %+v", 504, o.Payload)
}
func (o *PostFleetsFleetIDMembersGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *PostFleetsFleetIDMembersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostFleetsFleetIDMembersBody post_fleets_fleet_id_members_invitation
//
// invitation object
swagger:model PostFleetsFleetIDMembersBody
*/
type PostFleetsFleetIDMembersBody struct {

	// post_fleets_fleet_id_members_character_id
	//
	// The character you want to invite
	// Required: true
	CharacterID *int32 `json:"character_id"`

	// post_fleets_fleet_id_members_role
	//
	// If a character is invited with the `fleet_commander` role, neither `wing_id` or `squad_id` should be specified. If a character is invited with the `wing_commander` role, only `wing_id` should be specified. If a character is invited with the `squad_commander` role, both `wing_id` and `squad_id` should be specified. If a character is invited with the `squad_member` role, `wing_id` and `squad_id` should either both be specified or not specified at all. If they aren’t specified, the invited character will join any squad with available positions.
	// Required: true
	// Enum: [fleet_commander wing_commander squad_commander squad_member]
	Role *string `json:"role"`

	// post_fleets_fleet_id_members_squad_id
	//
	// squad_id integer
	// Minimum: 0
	SquadID *int64 `json:"squad_id,omitempty"`

	// post_fleets_fleet_id_members_wing_id
	//
	// wing_id integer
	// Minimum: 0
	WingID *int64 `json:"wing_id,omitempty"`
}

// Validate validates this post fleets fleet ID members body
func (o *PostFleetsFleetIDMembersBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCharacterID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSquadID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWingID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostFleetsFleetIDMembersBody) validateCharacterID(formats strfmt.Registry) error {

	if err := validate.Required("invitation"+"."+"character_id", "body", o.CharacterID); err != nil {
		return err
	}

	return nil
}

var postFleetsFleetIdMembersBodyTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fleet_commander","wing_commander","squad_commander","squad_member"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postFleetsFleetIdMembersBodyTypeRolePropEnum = append(postFleetsFleetIdMembersBodyTypeRolePropEnum, v)
	}
}

const (

	// PostFleetsFleetIDMembersBodyRoleFleetCommander captures enum value "fleet_commander"
	PostFleetsFleetIDMembersBodyRoleFleetCommander string = "fleet_commander"

	// PostFleetsFleetIDMembersBodyRoleWingCommander captures enum value "wing_commander"
	PostFleetsFleetIDMembersBodyRoleWingCommander string = "wing_commander"

	// PostFleetsFleetIDMembersBodyRoleSquadCommander captures enum value "squad_commander"
	PostFleetsFleetIDMembersBodyRoleSquadCommander string = "squad_commander"

	// PostFleetsFleetIDMembersBodyRoleSquadMember captures enum value "squad_member"
	PostFleetsFleetIDMembersBodyRoleSquadMember string = "squad_member"
)

// prop value enum
func (o *PostFleetsFleetIDMembersBody) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postFleetsFleetIdMembersBodyTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostFleetsFleetIDMembersBody) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("invitation"+"."+"role", "body", o.Role); err != nil {
		return err
	}

	// value enum
	if err := o.validateRoleEnum("invitation"+"."+"role", "body", *o.Role); err != nil {
		return err
	}

	return nil
}

func (o *PostFleetsFleetIDMembersBody) validateSquadID(formats strfmt.Registry) error {
	if swag.IsZero(o.SquadID) { // not required
		return nil
	}

	if err := validate.MinimumInt("invitation"+"."+"squad_id", "body", *o.SquadID, 0, false); err != nil {
		return err
	}

	return nil
}

func (o *PostFleetsFleetIDMembersBody) validateWingID(formats strfmt.Registry) error {
	if swag.IsZero(o.WingID) { // not required
		return nil
	}

	if err := validate.MinimumInt("invitation"+"."+"wing_id", "body", *o.WingID, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post fleets fleet ID members body based on context it is used
func (o *PostFleetsFleetIDMembersBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostFleetsFleetIDMembersBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostFleetsFleetIDMembersBody) UnmarshalBinary(b []byte) error {
	var res PostFleetsFleetIDMembersBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostFleetsFleetIDMembersNotFoundBody post_fleets_fleet_id_members_not_found
//
// Not found
swagger:model PostFleetsFleetIDMembersNotFoundBody
*/
type PostFleetsFleetIDMembersNotFoundBody struct {

	// post_fleets_fleet_id_members_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this post fleets fleet ID members not found body
func (o *PostFleetsFleetIDMembersNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post fleets fleet ID members not found body based on context it is used
func (o *PostFleetsFleetIDMembersNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostFleetsFleetIDMembersNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostFleetsFleetIDMembersNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PostFleetsFleetIDMembersNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostFleetsFleetIDMembersUnprocessableEntityBody post_fleets_fleet_id_members_unprocessable_entity
//
// 422 unprocessable entity object
swagger:model PostFleetsFleetIDMembersUnprocessableEntityBody
*/
type PostFleetsFleetIDMembersUnprocessableEntityBody struct {

	// post_fleets_fleet_id_members_error
	//
	// error message
	Error string `json:"error,omitempty"`
}

// Validate validates this post fleets fleet ID members unprocessable entity body
func (o *PostFleetsFleetIDMembersUnprocessableEntityBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post fleets fleet ID members unprocessable entity body based on context it is used
func (o *PostFleetsFleetIDMembersUnprocessableEntityBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostFleetsFleetIDMembersUnprocessableEntityBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostFleetsFleetIDMembersUnprocessableEntityBody) UnmarshalBinary(b []byte) error {
	var res PostFleetsFleetIDMembersUnprocessableEntityBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}