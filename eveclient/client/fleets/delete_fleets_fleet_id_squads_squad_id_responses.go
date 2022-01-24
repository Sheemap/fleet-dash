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

// DeleteFleetsFleetIDSquadsSquadIDReader is a Reader for the DeleteFleetsFleetIDSquadsSquadID structure.
type DeleteFleetsFleetIDSquadsSquadIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFleetsFleetIDSquadsSquadIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteFleetsFleetIDSquadsSquadIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteFleetsFleetIDSquadsSquadIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteFleetsFleetIDSquadsSquadIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteFleetsFleetIDSquadsSquadIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteFleetsFleetIDSquadsSquadIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewDeleteFleetsFleetIDSquadsSquadIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteFleetsFleetIDSquadsSquadIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteFleetsFleetIDSquadsSquadIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteFleetsFleetIDSquadsSquadIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteFleetsFleetIDSquadsSquadIDNoContent creates a DeleteFleetsFleetIDSquadsSquadIDNoContent with default headers values
func NewDeleteFleetsFleetIDSquadsSquadIDNoContent() *DeleteFleetsFleetIDSquadsSquadIDNoContent {
	return &DeleteFleetsFleetIDSquadsSquadIDNoContent{}
}

/* DeleteFleetsFleetIDSquadsSquadIDNoContent describes a response with status code 204, with default header values.

Squad deleted
*/
type DeleteFleetsFleetIDSquadsSquadIDNoContent struct {
}

func (o *DeleteFleetsFleetIDSquadsSquadIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/squads/{squad_id}/][%d] deleteFleetsFleetIdSquadsSquadIdNoContent ", 204)
}

func (o *DeleteFleetsFleetIDSquadsSquadIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFleetsFleetIDSquadsSquadIDBadRequest creates a DeleteFleetsFleetIDSquadsSquadIDBadRequest with default headers values
func NewDeleteFleetsFleetIDSquadsSquadIDBadRequest() *DeleteFleetsFleetIDSquadsSquadIDBadRequest {
	return &DeleteFleetsFleetIDSquadsSquadIDBadRequest{}
}

/* DeleteFleetsFleetIDSquadsSquadIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteFleetsFleetIDSquadsSquadIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *DeleteFleetsFleetIDSquadsSquadIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/squads/{squad_id}/][%d] deleteFleetsFleetIdSquadsSquadIdBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteFleetsFleetIDSquadsSquadIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *DeleteFleetsFleetIDSquadsSquadIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDSquadsSquadIDUnauthorized creates a DeleteFleetsFleetIDSquadsSquadIDUnauthorized with default headers values
func NewDeleteFleetsFleetIDSquadsSquadIDUnauthorized() *DeleteFleetsFleetIDSquadsSquadIDUnauthorized {
	return &DeleteFleetsFleetIDSquadsSquadIDUnauthorized{}
}

/* DeleteFleetsFleetIDSquadsSquadIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteFleetsFleetIDSquadsSquadIDUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *DeleteFleetsFleetIDSquadsSquadIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/squads/{squad_id}/][%d] deleteFleetsFleetIdSquadsSquadIdUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteFleetsFleetIDSquadsSquadIDUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *DeleteFleetsFleetIDSquadsSquadIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDSquadsSquadIDForbidden creates a DeleteFleetsFleetIDSquadsSquadIDForbidden with default headers values
func NewDeleteFleetsFleetIDSquadsSquadIDForbidden() *DeleteFleetsFleetIDSquadsSquadIDForbidden {
	return &DeleteFleetsFleetIDSquadsSquadIDForbidden{}
}

/* DeleteFleetsFleetIDSquadsSquadIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteFleetsFleetIDSquadsSquadIDForbidden struct {
	Payload *models.Forbidden
}

func (o *DeleteFleetsFleetIDSquadsSquadIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/squads/{squad_id}/][%d] deleteFleetsFleetIdSquadsSquadIdForbidden  %+v", 403, o.Payload)
}
func (o *DeleteFleetsFleetIDSquadsSquadIDForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *DeleteFleetsFleetIDSquadsSquadIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDSquadsSquadIDNotFound creates a DeleteFleetsFleetIDSquadsSquadIDNotFound with default headers values
func NewDeleteFleetsFleetIDSquadsSquadIDNotFound() *DeleteFleetsFleetIDSquadsSquadIDNotFound {
	return &DeleteFleetsFleetIDSquadsSquadIDNotFound{}
}

/* DeleteFleetsFleetIDSquadsSquadIDNotFound describes a response with status code 404, with default header values.

The fleet does not exist or you don't have access to it
*/
type DeleteFleetsFleetIDSquadsSquadIDNotFound struct {
	Payload *DeleteFleetsFleetIDSquadsSquadIDNotFoundBody
}

func (o *DeleteFleetsFleetIDSquadsSquadIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/squads/{squad_id}/][%d] deleteFleetsFleetIdSquadsSquadIdNotFound  %+v", 404, o.Payload)
}
func (o *DeleteFleetsFleetIDSquadsSquadIDNotFound) GetPayload() *DeleteFleetsFleetIDSquadsSquadIDNotFoundBody {
	return o.Payload
}

func (o *DeleteFleetsFleetIDSquadsSquadIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteFleetsFleetIDSquadsSquadIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDSquadsSquadIDEnhanceYourCalm creates a DeleteFleetsFleetIDSquadsSquadIDEnhanceYourCalm with default headers values
func NewDeleteFleetsFleetIDSquadsSquadIDEnhanceYourCalm() *DeleteFleetsFleetIDSquadsSquadIDEnhanceYourCalm {
	return &DeleteFleetsFleetIDSquadsSquadIDEnhanceYourCalm{}
}

/* DeleteFleetsFleetIDSquadsSquadIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type DeleteFleetsFleetIDSquadsSquadIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *DeleteFleetsFleetIDSquadsSquadIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/squads/{squad_id}/][%d] deleteFleetsFleetIdSquadsSquadIdEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *DeleteFleetsFleetIDSquadsSquadIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *DeleteFleetsFleetIDSquadsSquadIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDSquadsSquadIDInternalServerError creates a DeleteFleetsFleetIDSquadsSquadIDInternalServerError with default headers values
func NewDeleteFleetsFleetIDSquadsSquadIDInternalServerError() *DeleteFleetsFleetIDSquadsSquadIDInternalServerError {
	return &DeleteFleetsFleetIDSquadsSquadIDInternalServerError{}
}

/* DeleteFleetsFleetIDSquadsSquadIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteFleetsFleetIDSquadsSquadIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *DeleteFleetsFleetIDSquadsSquadIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/squads/{squad_id}/][%d] deleteFleetsFleetIdSquadsSquadIdInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteFleetsFleetIDSquadsSquadIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *DeleteFleetsFleetIDSquadsSquadIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDSquadsSquadIDServiceUnavailable creates a DeleteFleetsFleetIDSquadsSquadIDServiceUnavailable with default headers values
func NewDeleteFleetsFleetIDSquadsSquadIDServiceUnavailable() *DeleteFleetsFleetIDSquadsSquadIDServiceUnavailable {
	return &DeleteFleetsFleetIDSquadsSquadIDServiceUnavailable{}
}

/* DeleteFleetsFleetIDSquadsSquadIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type DeleteFleetsFleetIDSquadsSquadIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *DeleteFleetsFleetIDSquadsSquadIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/squads/{squad_id}/][%d] deleteFleetsFleetIdSquadsSquadIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *DeleteFleetsFleetIDSquadsSquadIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *DeleteFleetsFleetIDSquadsSquadIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDSquadsSquadIDGatewayTimeout creates a DeleteFleetsFleetIDSquadsSquadIDGatewayTimeout with default headers values
func NewDeleteFleetsFleetIDSquadsSquadIDGatewayTimeout() *DeleteFleetsFleetIDSquadsSquadIDGatewayTimeout {
	return &DeleteFleetsFleetIDSquadsSquadIDGatewayTimeout{}
}

/* DeleteFleetsFleetIDSquadsSquadIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type DeleteFleetsFleetIDSquadsSquadIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *DeleteFleetsFleetIDSquadsSquadIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/squads/{squad_id}/][%d] deleteFleetsFleetIdSquadsSquadIdGatewayTimeout  %+v", 504, o.Payload)
}
func (o *DeleteFleetsFleetIDSquadsSquadIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *DeleteFleetsFleetIDSquadsSquadIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteFleetsFleetIDSquadsSquadIDNotFoundBody delete_fleets_fleet_id_squads_squad_id_not_found
//
// Not found
swagger:model DeleteFleetsFleetIDSquadsSquadIDNotFoundBody
*/
type DeleteFleetsFleetIDSquadsSquadIDNotFoundBody struct {

	// delete_fleets_fleet_id_squads_squad_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this delete fleets fleet ID squads squad ID not found body
func (o *DeleteFleetsFleetIDSquadsSquadIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete fleets fleet ID squads squad ID not found body based on context it is used
func (o *DeleteFleetsFleetIDSquadsSquadIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteFleetsFleetIDSquadsSquadIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteFleetsFleetIDSquadsSquadIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteFleetsFleetIDSquadsSquadIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
