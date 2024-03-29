// Code generated by go-swagger; DO NOT EDIT.

package assets

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

// PostCorporationsCorporationIDAssetsNamesReader is a Reader for the PostCorporationsCorporationIDAssetsNames structure.
type PostCorporationsCorporationIDAssetsNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCorporationsCorporationIDAssetsNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCorporationsCorporationIDAssetsNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostCorporationsCorporationIDAssetsNamesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostCorporationsCorporationIDAssetsNamesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostCorporationsCorporationIDAssetsNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostCorporationsCorporationIDAssetsNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewPostCorporationsCorporationIDAssetsNamesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostCorporationsCorporationIDAssetsNamesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostCorporationsCorporationIDAssetsNamesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostCorporationsCorporationIDAssetsNamesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostCorporationsCorporationIDAssetsNamesOK creates a PostCorporationsCorporationIDAssetsNamesOK with default headers values
func NewPostCorporationsCorporationIDAssetsNamesOK() *PostCorporationsCorporationIDAssetsNamesOK {
	return &PostCorporationsCorporationIDAssetsNamesOK{}
}

/* PostCorporationsCorporationIDAssetsNamesOK describes a response with status code 200, with default header values.

List of asset names
*/
type PostCorporationsCorporationIDAssetsNamesOK struct {
	Payload []*PostCorporationsCorporationIDAssetsNamesOKBodyItems0
}

func (o *PostCorporationsCorporationIDAssetsNamesOK) Error() string {
	return fmt.Sprintf("[POST /corporations/{corporation_id}/assets/names/][%d] postCorporationsCorporationIdAssetsNamesOK  %+v", 200, o.Payload)
}
func (o *PostCorporationsCorporationIDAssetsNamesOK) GetPayload() []*PostCorporationsCorporationIDAssetsNamesOKBodyItems0 {
	return o.Payload
}

func (o *PostCorporationsCorporationIDAssetsNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCorporationsCorporationIDAssetsNamesBadRequest creates a PostCorporationsCorporationIDAssetsNamesBadRequest with default headers values
func NewPostCorporationsCorporationIDAssetsNamesBadRequest() *PostCorporationsCorporationIDAssetsNamesBadRequest {
	return &PostCorporationsCorporationIDAssetsNamesBadRequest{}
}

/* PostCorporationsCorporationIDAssetsNamesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostCorporationsCorporationIDAssetsNamesBadRequest struct {
	Payload *models.BadRequest
}

func (o *PostCorporationsCorporationIDAssetsNamesBadRequest) Error() string {
	return fmt.Sprintf("[POST /corporations/{corporation_id}/assets/names/][%d] postCorporationsCorporationIdAssetsNamesBadRequest  %+v", 400, o.Payload)
}
func (o *PostCorporationsCorporationIDAssetsNamesBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *PostCorporationsCorporationIDAssetsNamesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCorporationsCorporationIDAssetsNamesUnauthorized creates a PostCorporationsCorporationIDAssetsNamesUnauthorized with default headers values
func NewPostCorporationsCorporationIDAssetsNamesUnauthorized() *PostCorporationsCorporationIDAssetsNamesUnauthorized {
	return &PostCorporationsCorporationIDAssetsNamesUnauthorized{}
}

/* PostCorporationsCorporationIDAssetsNamesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostCorporationsCorporationIDAssetsNamesUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *PostCorporationsCorporationIDAssetsNamesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /corporations/{corporation_id}/assets/names/][%d] postCorporationsCorporationIdAssetsNamesUnauthorized  %+v", 401, o.Payload)
}
func (o *PostCorporationsCorporationIDAssetsNamesUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *PostCorporationsCorporationIDAssetsNamesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCorporationsCorporationIDAssetsNamesForbidden creates a PostCorporationsCorporationIDAssetsNamesForbidden with default headers values
func NewPostCorporationsCorporationIDAssetsNamesForbidden() *PostCorporationsCorporationIDAssetsNamesForbidden {
	return &PostCorporationsCorporationIDAssetsNamesForbidden{}
}

/* PostCorporationsCorporationIDAssetsNamesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostCorporationsCorporationIDAssetsNamesForbidden struct {
	Payload *models.Forbidden
}

func (o *PostCorporationsCorporationIDAssetsNamesForbidden) Error() string {
	return fmt.Sprintf("[POST /corporations/{corporation_id}/assets/names/][%d] postCorporationsCorporationIdAssetsNamesForbidden  %+v", 403, o.Payload)
}
func (o *PostCorporationsCorporationIDAssetsNamesForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *PostCorporationsCorporationIDAssetsNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCorporationsCorporationIDAssetsNamesNotFound creates a PostCorporationsCorporationIDAssetsNamesNotFound with default headers values
func NewPostCorporationsCorporationIDAssetsNamesNotFound() *PostCorporationsCorporationIDAssetsNamesNotFound {
	return &PostCorporationsCorporationIDAssetsNamesNotFound{}
}

/* PostCorporationsCorporationIDAssetsNamesNotFound describes a response with status code 404, with default header values.

Invalid IDs
*/
type PostCorporationsCorporationIDAssetsNamesNotFound struct {
	Payload *PostCorporationsCorporationIDAssetsNamesNotFoundBody
}

func (o *PostCorporationsCorporationIDAssetsNamesNotFound) Error() string {
	return fmt.Sprintf("[POST /corporations/{corporation_id}/assets/names/][%d] postCorporationsCorporationIdAssetsNamesNotFound  %+v", 404, o.Payload)
}
func (o *PostCorporationsCorporationIDAssetsNamesNotFound) GetPayload() *PostCorporationsCorporationIDAssetsNamesNotFoundBody {
	return o.Payload
}

func (o *PostCorporationsCorporationIDAssetsNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostCorporationsCorporationIDAssetsNamesNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCorporationsCorporationIDAssetsNamesEnhanceYourCalm creates a PostCorporationsCorporationIDAssetsNamesEnhanceYourCalm with default headers values
func NewPostCorporationsCorporationIDAssetsNamesEnhanceYourCalm() *PostCorporationsCorporationIDAssetsNamesEnhanceYourCalm {
	return &PostCorporationsCorporationIDAssetsNamesEnhanceYourCalm{}
}

/* PostCorporationsCorporationIDAssetsNamesEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type PostCorporationsCorporationIDAssetsNamesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *PostCorporationsCorporationIDAssetsNamesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /corporations/{corporation_id}/assets/names/][%d] postCorporationsCorporationIdAssetsNamesEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *PostCorporationsCorporationIDAssetsNamesEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *PostCorporationsCorporationIDAssetsNamesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCorporationsCorporationIDAssetsNamesInternalServerError creates a PostCorporationsCorporationIDAssetsNamesInternalServerError with default headers values
func NewPostCorporationsCorporationIDAssetsNamesInternalServerError() *PostCorporationsCorporationIDAssetsNamesInternalServerError {
	return &PostCorporationsCorporationIDAssetsNamesInternalServerError{}
}

/* PostCorporationsCorporationIDAssetsNamesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostCorporationsCorporationIDAssetsNamesInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *PostCorporationsCorporationIDAssetsNamesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /corporations/{corporation_id}/assets/names/][%d] postCorporationsCorporationIdAssetsNamesInternalServerError  %+v", 500, o.Payload)
}
func (o *PostCorporationsCorporationIDAssetsNamesInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *PostCorporationsCorporationIDAssetsNamesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCorporationsCorporationIDAssetsNamesServiceUnavailable creates a PostCorporationsCorporationIDAssetsNamesServiceUnavailable with default headers values
func NewPostCorporationsCorporationIDAssetsNamesServiceUnavailable() *PostCorporationsCorporationIDAssetsNamesServiceUnavailable {
	return &PostCorporationsCorporationIDAssetsNamesServiceUnavailable{}
}

/* PostCorporationsCorporationIDAssetsNamesServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type PostCorporationsCorporationIDAssetsNamesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *PostCorporationsCorporationIDAssetsNamesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /corporations/{corporation_id}/assets/names/][%d] postCorporationsCorporationIdAssetsNamesServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PostCorporationsCorporationIDAssetsNamesServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *PostCorporationsCorporationIDAssetsNamesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCorporationsCorporationIDAssetsNamesGatewayTimeout creates a PostCorporationsCorporationIDAssetsNamesGatewayTimeout with default headers values
func NewPostCorporationsCorporationIDAssetsNamesGatewayTimeout() *PostCorporationsCorporationIDAssetsNamesGatewayTimeout {
	return &PostCorporationsCorporationIDAssetsNamesGatewayTimeout{}
}

/* PostCorporationsCorporationIDAssetsNamesGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type PostCorporationsCorporationIDAssetsNamesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *PostCorporationsCorporationIDAssetsNamesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /corporations/{corporation_id}/assets/names/][%d] postCorporationsCorporationIdAssetsNamesGatewayTimeout  %+v", 504, o.Payload)
}
func (o *PostCorporationsCorporationIDAssetsNamesGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *PostCorporationsCorporationIDAssetsNamesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostCorporationsCorporationIDAssetsNamesNotFoundBody post_corporations_corporation_id_assets_names_not_found
//
// Not found
swagger:model PostCorporationsCorporationIDAssetsNamesNotFoundBody
*/
type PostCorporationsCorporationIDAssetsNamesNotFoundBody struct {

	// post_corporations_corporation_id_assets_names_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this post corporations corporation ID assets names not found body
func (o *PostCorporationsCorporationIDAssetsNamesNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post corporations corporation ID assets names not found body based on context it is used
func (o *PostCorporationsCorporationIDAssetsNamesNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCorporationsCorporationIDAssetsNamesNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCorporationsCorporationIDAssetsNamesNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PostCorporationsCorporationIDAssetsNamesNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostCorporationsCorporationIDAssetsNamesOKBodyItems0 post_corporations_corporation_id_assets_names_200_ok
//
// 200 ok object
swagger:model PostCorporationsCorporationIDAssetsNamesOKBodyItems0
*/
type PostCorporationsCorporationIDAssetsNamesOKBodyItems0 struct {

	// post_corporations_corporation_id_assets_names_item_id
	//
	// item_id integer
	// Required: true
	ItemID *int64 `json:"item_id"`

	// post_corporations_corporation_id_assets_names_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this post corporations corporation ID assets names o k body items0
func (o *PostCorporationsCorporationIDAssetsNamesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCorporationsCorporationIDAssetsNamesOKBodyItems0) validateItemID(formats strfmt.Registry) error {

	if err := validate.Required("item_id", "body", o.ItemID); err != nil {
		return err
	}

	return nil
}

func (o *PostCorporationsCorporationIDAssetsNamesOKBodyItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post corporations corporation ID assets names o k body items0 based on context it is used
func (o *PostCorporationsCorporationIDAssetsNamesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCorporationsCorporationIDAssetsNamesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCorporationsCorporationIDAssetsNamesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res PostCorporationsCorporationIDAssetsNamesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
