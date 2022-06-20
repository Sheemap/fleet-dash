// Code generated by go-swagger; DO NOT EDIT.

package fleets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"fleet-dash-core/eveclient/models"
)

// DeleteFleetsFleetIDMembersMemberIDReader is a Reader for the DeleteFleetsFleetIDMembersMemberID structure.
type DeleteFleetsFleetIDMembersMemberIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFleetsFleetIDMembersMemberIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteFleetsFleetIDMembersMemberIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteFleetsFleetIDMembersMemberIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteFleetsFleetIDMembersMemberIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteFleetsFleetIDMembersMemberIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteFleetsFleetIDMembersMemberIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewDeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteFleetsFleetIDMembersMemberIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteFleetsFleetIDMembersMemberIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteFleetsFleetIDMembersMemberIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteFleetsFleetIDMembersMemberIDNoContent creates a DeleteFleetsFleetIDMembersMemberIDNoContent with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDNoContent() *DeleteFleetsFleetIDMembersMemberIDNoContent {
	return &DeleteFleetsFleetIDMembersMemberIDNoContent{}
}

/* DeleteFleetsFleetIDMembersMemberIDNoContent describes a response with status code 204, with default header values.

Fleet member kicked
*/
type DeleteFleetsFleetIDMembersMemberIDNoContent struct {
}

func (o *DeleteFleetsFleetIDMembersMemberIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdNoContent ", 204)
}

func (o *DeleteFleetsFleetIDMembersMemberIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFleetsFleetIDMembersMemberIDBadRequest creates a DeleteFleetsFleetIDMembersMemberIDBadRequest with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDBadRequest() *DeleteFleetsFleetIDMembersMemberIDBadRequest {
	return &DeleteFleetsFleetIDMembersMemberIDBadRequest{}
}

/* DeleteFleetsFleetIDMembersMemberIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteFleetsFleetIDMembersMemberIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *DeleteFleetsFleetIDMembersMemberIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteFleetsFleetIDMembersMemberIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *DeleteFleetsFleetIDMembersMemberIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDMembersMemberIDUnauthorized creates a DeleteFleetsFleetIDMembersMemberIDUnauthorized with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDUnauthorized() *DeleteFleetsFleetIDMembersMemberIDUnauthorized {
	return &DeleteFleetsFleetIDMembersMemberIDUnauthorized{}
}

/* DeleteFleetsFleetIDMembersMemberIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteFleetsFleetIDMembersMemberIDUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *DeleteFleetsFleetIDMembersMemberIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteFleetsFleetIDMembersMemberIDUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *DeleteFleetsFleetIDMembersMemberIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDMembersMemberIDForbidden creates a DeleteFleetsFleetIDMembersMemberIDForbidden with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDForbidden() *DeleteFleetsFleetIDMembersMemberIDForbidden {
	return &DeleteFleetsFleetIDMembersMemberIDForbidden{}
}

/* DeleteFleetsFleetIDMembersMemberIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteFleetsFleetIDMembersMemberIDForbidden struct {
	Payload *models.Forbidden
}

func (o *DeleteFleetsFleetIDMembersMemberIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdForbidden  %+v", 403, o.Payload)
}
func (o *DeleteFleetsFleetIDMembersMemberIDForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *DeleteFleetsFleetIDMembersMemberIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDMembersMemberIDNotFound creates a DeleteFleetsFleetIDMembersMemberIDNotFound with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDNotFound() *DeleteFleetsFleetIDMembersMemberIDNotFound {
	return &DeleteFleetsFleetIDMembersMemberIDNotFound{}
}

/* DeleteFleetsFleetIDMembersMemberIDNotFound describes a response with status code 404, with default header values.

The fleet does not exist or you don't have access to it
*/
type DeleteFleetsFleetIDMembersMemberIDNotFound struct {
	Payload *DeleteFleetsFleetIDMembersMemberIDNotFoundBody
}

func (o *DeleteFleetsFleetIDMembersMemberIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdNotFound  %+v", 404, o.Payload)
}
func (o *DeleteFleetsFleetIDMembersMemberIDNotFound) GetPayload() *DeleteFleetsFleetIDMembersMemberIDNotFoundBody {
	return o.Payload
}

func (o *DeleteFleetsFleetIDMembersMemberIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteFleetsFleetIDMembersMemberIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm creates a DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm() *DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm {
	return &DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm{}
}

/* DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDMembersMemberIDInternalServerError creates a DeleteFleetsFleetIDMembersMemberIDInternalServerError with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDInternalServerError() *DeleteFleetsFleetIDMembersMemberIDInternalServerError {
	return &DeleteFleetsFleetIDMembersMemberIDInternalServerError{}
}

/* DeleteFleetsFleetIDMembersMemberIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteFleetsFleetIDMembersMemberIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *DeleteFleetsFleetIDMembersMemberIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteFleetsFleetIDMembersMemberIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *DeleteFleetsFleetIDMembersMemberIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDMembersMemberIDServiceUnavailable creates a DeleteFleetsFleetIDMembersMemberIDServiceUnavailable with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDServiceUnavailable() *DeleteFleetsFleetIDMembersMemberIDServiceUnavailable {
	return &DeleteFleetsFleetIDMembersMemberIDServiceUnavailable{}
}

/* DeleteFleetsFleetIDMembersMemberIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type DeleteFleetsFleetIDMembersMemberIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *DeleteFleetsFleetIDMembersMemberIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *DeleteFleetsFleetIDMembersMemberIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *DeleteFleetsFleetIDMembersMemberIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDMembersMemberIDGatewayTimeout creates a DeleteFleetsFleetIDMembersMemberIDGatewayTimeout with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDGatewayTimeout() *DeleteFleetsFleetIDMembersMemberIDGatewayTimeout {
	return &DeleteFleetsFleetIDMembersMemberIDGatewayTimeout{}
}

/* DeleteFleetsFleetIDMembersMemberIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type DeleteFleetsFleetIDMembersMemberIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *DeleteFleetsFleetIDMembersMemberIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdGatewayTimeout  %+v", 504, o.Payload)
}
func (o *DeleteFleetsFleetIDMembersMemberIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *DeleteFleetsFleetIDMembersMemberIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteFleetsFleetIDMembersMemberIDNotFoundBody delete_fleets_fleet_id_members_member_id_not_found
//
// Not found
swagger:model DeleteFleetsFleetIDMembersMemberIDNotFoundBody
*/
type DeleteFleetsFleetIDMembersMemberIDNotFoundBody struct {

	// delete_fleets_fleet_id_members_member_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this delete fleets fleet ID members member ID not found body
func (o *DeleteFleetsFleetIDMembersMemberIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete fleets fleet ID members member ID not found body based on context it is used
func (o *DeleteFleetsFleetIDMembersMemberIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteFleetsFleetIDMembersMemberIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteFleetsFleetIDMembersMemberIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteFleetsFleetIDMembersMemberIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}