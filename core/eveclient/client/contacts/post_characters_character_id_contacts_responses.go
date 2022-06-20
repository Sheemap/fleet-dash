// Code generated by go-swagger; DO NOT EDIT.

package contacts

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

// PostCharactersCharacterIDContactsReader is a Reader for the PostCharactersCharacterIDContacts structure.
type PostCharactersCharacterIDContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCharactersCharacterIDContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostCharactersCharacterIDContactsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostCharactersCharacterIDContactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostCharactersCharacterIDContactsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostCharactersCharacterIDContactsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewPostCharactersCharacterIDContactsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostCharactersCharacterIDContactsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostCharactersCharacterIDContactsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostCharactersCharacterIDContactsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 520:
		result := NewPostCharactersCharacterIDContactsStatus520()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostCharactersCharacterIDContactsCreated creates a PostCharactersCharacterIDContactsCreated with default headers values
func NewPostCharactersCharacterIDContactsCreated() *PostCharactersCharacterIDContactsCreated {
	return &PostCharactersCharacterIDContactsCreated{}
}

/* PostCharactersCharacterIDContactsCreated describes a response with status code 201, with default header values.

A list of contact ids that successfully created
*/
type PostCharactersCharacterIDContactsCreated struct {
	Payload []int32
}

func (o *PostCharactersCharacterIDContactsCreated) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/contacts/][%d] postCharactersCharacterIdContactsCreated  %+v", 201, o.Payload)
}
func (o *PostCharactersCharacterIDContactsCreated) GetPayload() []int32 {
	return o.Payload
}

func (o *PostCharactersCharacterIDContactsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDContactsBadRequest creates a PostCharactersCharacterIDContactsBadRequest with default headers values
func NewPostCharactersCharacterIDContactsBadRequest() *PostCharactersCharacterIDContactsBadRequest {
	return &PostCharactersCharacterIDContactsBadRequest{}
}

/* PostCharactersCharacterIDContactsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostCharactersCharacterIDContactsBadRequest struct {
	Payload *models.BadRequest
}

func (o *PostCharactersCharacterIDContactsBadRequest) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/contacts/][%d] postCharactersCharacterIdContactsBadRequest  %+v", 400, o.Payload)
}
func (o *PostCharactersCharacterIDContactsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *PostCharactersCharacterIDContactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDContactsUnauthorized creates a PostCharactersCharacterIDContactsUnauthorized with default headers values
func NewPostCharactersCharacterIDContactsUnauthorized() *PostCharactersCharacterIDContactsUnauthorized {
	return &PostCharactersCharacterIDContactsUnauthorized{}
}

/* PostCharactersCharacterIDContactsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostCharactersCharacterIDContactsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *PostCharactersCharacterIDContactsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/contacts/][%d] postCharactersCharacterIdContactsUnauthorized  %+v", 401, o.Payload)
}
func (o *PostCharactersCharacterIDContactsUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *PostCharactersCharacterIDContactsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDContactsForbidden creates a PostCharactersCharacterIDContactsForbidden with default headers values
func NewPostCharactersCharacterIDContactsForbidden() *PostCharactersCharacterIDContactsForbidden {
	return &PostCharactersCharacterIDContactsForbidden{}
}

/* PostCharactersCharacterIDContactsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostCharactersCharacterIDContactsForbidden struct {
	Payload *models.Forbidden
}

func (o *PostCharactersCharacterIDContactsForbidden) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/contacts/][%d] postCharactersCharacterIdContactsForbidden  %+v", 403, o.Payload)
}
func (o *PostCharactersCharacterIDContactsForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *PostCharactersCharacterIDContactsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDContactsEnhanceYourCalm creates a PostCharactersCharacterIDContactsEnhanceYourCalm with default headers values
func NewPostCharactersCharacterIDContactsEnhanceYourCalm() *PostCharactersCharacterIDContactsEnhanceYourCalm {
	return &PostCharactersCharacterIDContactsEnhanceYourCalm{}
}

/* PostCharactersCharacterIDContactsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type PostCharactersCharacterIDContactsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *PostCharactersCharacterIDContactsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/contacts/][%d] postCharactersCharacterIdContactsEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *PostCharactersCharacterIDContactsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *PostCharactersCharacterIDContactsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDContactsInternalServerError creates a PostCharactersCharacterIDContactsInternalServerError with default headers values
func NewPostCharactersCharacterIDContactsInternalServerError() *PostCharactersCharacterIDContactsInternalServerError {
	return &PostCharactersCharacterIDContactsInternalServerError{}
}

/* PostCharactersCharacterIDContactsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostCharactersCharacterIDContactsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *PostCharactersCharacterIDContactsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/contacts/][%d] postCharactersCharacterIdContactsInternalServerError  %+v", 500, o.Payload)
}
func (o *PostCharactersCharacterIDContactsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *PostCharactersCharacterIDContactsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDContactsServiceUnavailable creates a PostCharactersCharacterIDContactsServiceUnavailable with default headers values
func NewPostCharactersCharacterIDContactsServiceUnavailable() *PostCharactersCharacterIDContactsServiceUnavailable {
	return &PostCharactersCharacterIDContactsServiceUnavailable{}
}

/* PostCharactersCharacterIDContactsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type PostCharactersCharacterIDContactsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *PostCharactersCharacterIDContactsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/contacts/][%d] postCharactersCharacterIdContactsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PostCharactersCharacterIDContactsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *PostCharactersCharacterIDContactsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDContactsGatewayTimeout creates a PostCharactersCharacterIDContactsGatewayTimeout with default headers values
func NewPostCharactersCharacterIDContactsGatewayTimeout() *PostCharactersCharacterIDContactsGatewayTimeout {
	return &PostCharactersCharacterIDContactsGatewayTimeout{}
}

/* PostCharactersCharacterIDContactsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type PostCharactersCharacterIDContactsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *PostCharactersCharacterIDContactsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/contacts/][%d] postCharactersCharacterIdContactsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *PostCharactersCharacterIDContactsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *PostCharactersCharacterIDContactsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDContactsStatus520 creates a PostCharactersCharacterIDContactsStatus520 with default headers values
func NewPostCharactersCharacterIDContactsStatus520() *PostCharactersCharacterIDContactsStatus520 {
	return &PostCharactersCharacterIDContactsStatus520{}
}

/* PostCharactersCharacterIDContactsStatus520 describes a response with status code 520, with default header values.

Internal error thrown from the EVE server
*/
type PostCharactersCharacterIDContactsStatus520 struct {
	Payload *PostCharactersCharacterIDContactsStatus520Body
}

func (o *PostCharactersCharacterIDContactsStatus520) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/contacts/][%d] postCharactersCharacterIdContactsStatus520  %+v", 520, o.Payload)
}
func (o *PostCharactersCharacterIDContactsStatus520) GetPayload() *PostCharactersCharacterIDContactsStatus520Body {
	return o.Payload
}

func (o *PostCharactersCharacterIDContactsStatus520) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostCharactersCharacterIDContactsStatus520Body)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostCharactersCharacterIDContactsStatus520Body post_characters_character_id_contacts_error_520
//
// Error 520
swagger:model PostCharactersCharacterIDContactsStatus520Body
*/
type PostCharactersCharacterIDContactsStatus520Body struct {

	// post_characters_character_id_contacts_520_error_520
	//
	// Error 520 message
	Error string `json:"error,omitempty"`
}

// Validate validates this post characters character ID contacts status520 body
func (o *PostCharactersCharacterIDContactsStatus520Body) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post characters character ID contacts status520 body based on context it is used
func (o *PostCharactersCharacterIDContactsStatus520Body) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCharactersCharacterIDContactsStatus520Body) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCharactersCharacterIDContactsStatus520Body) UnmarshalBinary(b []byte) error {
	var res PostCharactersCharacterIDContactsStatus520Body
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}