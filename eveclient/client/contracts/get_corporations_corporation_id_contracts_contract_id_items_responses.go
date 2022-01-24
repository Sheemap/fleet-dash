// Code generated by go-swagger; DO NOT EDIT.

package contracts

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

// GetCorporationsCorporationIDContractsContractIDItemsReader is a Reader for the GetCorporationsCorporationIDContractsContractIDItems structure.
type GetCorporationsCorporationIDContractsContractIDItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDContractsContractIDItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationsCorporationIDContractsContractIDItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationsCorporationIDContractsContractIDItemsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationsCorporationIDContractsContractIDItemsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationsCorporationIDContractsContractIDItemsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationsCorporationIDContractsContractIDItemsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCorporationsCorporationIDContractsContractIDItemsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationsCorporationIDContractsContractIDItemsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationsCorporationIDContractsContractIDItemsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationsCorporationIDContractsContractIDItemsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationsCorporationIDContractsContractIDItemsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 520:
		result := NewGetCorporationsCorporationIDContractsContractIDItemsStatus520()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDContractsContractIDItemsOK creates a GetCorporationsCorporationIDContractsContractIDItemsOK with default headers values
func NewGetCorporationsCorporationIDContractsContractIDItemsOK() *GetCorporationsCorporationIDContractsContractIDItemsOK {
	return &GetCorporationsCorporationIDContractsContractIDItemsOK{}
}

/* GetCorporationsCorporationIDContractsContractIDItemsOK describes a response with status code 200, with default header values.

A list of items in this contract
*/
type GetCorporationsCorporationIDContractsContractIDItemsOK struct {

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

	Payload []*GetCorporationsCorporationIDContractsContractIDItemsOKBodyItems0
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/contracts/{contract_id}/items/][%d] getCorporationsCorporationIdContractsContractIdItemsOK  %+v", 200, o.Payload)
}
func (o *GetCorporationsCorporationIDContractsContractIDItemsOK) GetPayload() []*GetCorporationsCorporationIDContractsContractIDItemsOKBodyItems0 {
	return o.Payload
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDContractsContractIDItemsNotModified creates a GetCorporationsCorporationIDContractsContractIDItemsNotModified with default headers values
func NewGetCorporationsCorporationIDContractsContractIDItemsNotModified() *GetCorporationsCorporationIDContractsContractIDItemsNotModified {
	return &GetCorporationsCorporationIDContractsContractIDItemsNotModified{}
}

/* GetCorporationsCorporationIDContractsContractIDItemsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationsCorporationIDContractsContractIDItemsNotModified struct {

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

func (o *GetCorporationsCorporationIDContractsContractIDItemsNotModified) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/contracts/{contract_id}/items/][%d] getCorporationsCorporationIdContractsContractIdItemsNotModified ", 304)
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDContractsContractIDItemsBadRequest creates a GetCorporationsCorporationIDContractsContractIDItemsBadRequest with default headers values
func NewGetCorporationsCorporationIDContractsContractIDItemsBadRequest() *GetCorporationsCorporationIDContractsContractIDItemsBadRequest {
	return &GetCorporationsCorporationIDContractsContractIDItemsBadRequest{}
}

/* GetCorporationsCorporationIDContractsContractIDItemsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationsCorporationIDContractsContractIDItemsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsBadRequest) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/contracts/{contract_id}/items/][%d] getCorporationsCorporationIdContractsContractIdItemsBadRequest  %+v", 400, o.Payload)
}
func (o *GetCorporationsCorporationIDContractsContractIDItemsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDContractsContractIDItemsUnauthorized creates a GetCorporationsCorporationIDContractsContractIDItemsUnauthorized with default headers values
func NewGetCorporationsCorporationIDContractsContractIDItemsUnauthorized() *GetCorporationsCorporationIDContractsContractIDItemsUnauthorized {
	return &GetCorporationsCorporationIDContractsContractIDItemsUnauthorized{}
}

/* GetCorporationsCorporationIDContractsContractIDItemsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDContractsContractIDItemsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/contracts/{contract_id}/items/][%d] getCorporationsCorporationIdContractsContractIdItemsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCorporationsCorporationIDContractsContractIDItemsUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDContractsContractIDItemsForbidden creates a GetCorporationsCorporationIDContractsContractIDItemsForbidden with default headers values
func NewGetCorporationsCorporationIDContractsContractIDItemsForbidden() *GetCorporationsCorporationIDContractsContractIDItemsForbidden {
	return &GetCorporationsCorporationIDContractsContractIDItemsForbidden{}
}

/* GetCorporationsCorporationIDContractsContractIDItemsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationsCorporationIDContractsContractIDItemsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsForbidden) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/contracts/{contract_id}/items/][%d] getCorporationsCorporationIdContractsContractIdItemsForbidden  %+v", 403, o.Payload)
}
func (o *GetCorporationsCorporationIDContractsContractIDItemsForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDContractsContractIDItemsNotFound creates a GetCorporationsCorporationIDContractsContractIDItemsNotFound with default headers values
func NewGetCorporationsCorporationIDContractsContractIDItemsNotFound() *GetCorporationsCorporationIDContractsContractIDItemsNotFound {
	return &GetCorporationsCorporationIDContractsContractIDItemsNotFound{}
}

/* GetCorporationsCorporationIDContractsContractIDItemsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetCorporationsCorporationIDContractsContractIDItemsNotFound struct {
	Payload *GetCorporationsCorporationIDContractsContractIDItemsNotFoundBody
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsNotFound) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/contracts/{contract_id}/items/][%d] getCorporationsCorporationIdContractsContractIdItemsNotFound  %+v", 404, o.Payload)
}
func (o *GetCorporationsCorporationIDContractsContractIDItemsNotFound) GetPayload() *GetCorporationsCorporationIDContractsContractIDItemsNotFoundBody {
	return o.Payload
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCorporationsCorporationIDContractsContractIDItemsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDContractsContractIDItemsEnhanceYourCalm creates a GetCorporationsCorporationIDContractsContractIDItemsEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDContractsContractIDItemsEnhanceYourCalm() *GetCorporationsCorporationIDContractsContractIDItemsEnhanceYourCalm {
	return &GetCorporationsCorporationIDContractsContractIDItemsEnhanceYourCalm{}
}

/* GetCorporationsCorporationIDContractsContractIDItemsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationsCorporationIDContractsContractIDItemsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/contracts/{contract_id}/items/][%d] getCorporationsCorporationIdContractsContractIdItemsEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCorporationsCorporationIDContractsContractIDItemsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDContractsContractIDItemsInternalServerError creates a GetCorporationsCorporationIDContractsContractIDItemsInternalServerError with default headers values
func NewGetCorporationsCorporationIDContractsContractIDItemsInternalServerError() *GetCorporationsCorporationIDContractsContractIDItemsInternalServerError {
	return &GetCorporationsCorporationIDContractsContractIDItemsInternalServerError{}
}

/* GetCorporationsCorporationIDContractsContractIDItemsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationsCorporationIDContractsContractIDItemsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/contracts/{contract_id}/items/][%d] getCorporationsCorporationIdContractsContractIdItemsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCorporationsCorporationIDContractsContractIDItemsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDContractsContractIDItemsServiceUnavailable creates a GetCorporationsCorporationIDContractsContractIDItemsServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDContractsContractIDItemsServiceUnavailable() *GetCorporationsCorporationIDContractsContractIDItemsServiceUnavailable {
	return &GetCorporationsCorporationIDContractsContractIDItemsServiceUnavailable{}
}

/* GetCorporationsCorporationIDContractsContractIDItemsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDContractsContractIDItemsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/contracts/{contract_id}/items/][%d] getCorporationsCorporationIdContractsContractIdItemsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCorporationsCorporationIDContractsContractIDItemsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDContractsContractIDItemsGatewayTimeout creates a GetCorporationsCorporationIDContractsContractIDItemsGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDContractsContractIDItemsGatewayTimeout() *GetCorporationsCorporationIDContractsContractIDItemsGatewayTimeout {
	return &GetCorporationsCorporationIDContractsContractIDItemsGatewayTimeout{}
}

/* GetCorporationsCorporationIDContractsContractIDItemsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDContractsContractIDItemsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/contracts/{contract_id}/items/][%d] getCorporationsCorporationIdContractsContractIdItemsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCorporationsCorporationIDContractsContractIDItemsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDContractsContractIDItemsStatus520 creates a GetCorporationsCorporationIDContractsContractIDItemsStatus520 with default headers values
func NewGetCorporationsCorporationIDContractsContractIDItemsStatus520() *GetCorporationsCorporationIDContractsContractIDItemsStatus520 {
	return &GetCorporationsCorporationIDContractsContractIDItemsStatus520{}
}

/* GetCorporationsCorporationIDContractsContractIDItemsStatus520 describes a response with status code 520, with default header values.

Internal error thrown from the EVE server. Most of the time this means you have hit an EVE server rate limit
*/
type GetCorporationsCorporationIDContractsContractIDItemsStatus520 struct {
	Payload *GetCorporationsCorporationIDContractsContractIDItemsStatus520Body
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsStatus520) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/contracts/{contract_id}/items/][%d] getCorporationsCorporationIdContractsContractIdItemsStatus520  %+v", 520, o.Payload)
}
func (o *GetCorporationsCorporationIDContractsContractIDItemsStatus520) GetPayload() *GetCorporationsCorporationIDContractsContractIDItemsStatus520Body {
	return o.Payload
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsStatus520) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCorporationsCorporationIDContractsContractIDItemsStatus520Body)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDContractsContractIDItemsNotFoundBody get_corporations_corporation_id_contracts_contract_id_items_not_found
//
// Not found
swagger:model GetCorporationsCorporationIDContractsContractIDItemsNotFoundBody
*/
type GetCorporationsCorporationIDContractsContractIDItemsNotFoundBody struct {

	// get_corporations_corporation_id_contracts_contract_id_items_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get corporations corporation ID contracts contract ID items not found body
func (o *GetCorporationsCorporationIDContractsContractIDItemsNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get corporations corporation ID contracts contract ID items not found body based on context it is used
func (o *GetCorporationsCorporationIDContractsContractIDItemsNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDContractsContractIDItemsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDContractsContractIDItemsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDContractsContractIDItemsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetCorporationsCorporationIDContractsContractIDItemsOKBodyItems0 get_corporations_corporation_id_contracts_contract_id_items_200_ok
//
// 200 ok object
swagger:model GetCorporationsCorporationIDContractsContractIDItemsOKBodyItems0
*/
type GetCorporationsCorporationIDContractsContractIDItemsOKBodyItems0 struct {

	// get_corporations_corporation_id_contracts_contract_id_items_is_included
	//
	// true if the contract issuer has submitted this item with the contract, false if the isser is asking for this item in the contract
	// Required: true
	IsIncluded *bool `json:"is_included"`

	// get_corporations_corporation_id_contracts_contract_id_items_is_singleton
	//
	// is_singleton boolean
	// Required: true
	IsSingleton *bool `json:"is_singleton"`

	// get_corporations_corporation_id_contracts_contract_id_items_quantity
	//
	// Number of items in the stack
	// Required: true
	Quantity *int32 `json:"quantity"`

	// get_corporations_corporation_id_contracts_contract_id_items_raw_quantity
	//
	// -1 indicates that the item is a singleton (non-stackable). If the item happens to be a Blueprint, -1 is an Original and -2 is a Blueprint Copy
	RawQuantity int32 `json:"raw_quantity,omitempty"`

	// get_corporations_corporation_id_contracts_contract_id_items_record_id
	//
	// Unique ID for the item
	// Required: true
	RecordID *int64 `json:"record_id"`

	// get_corporations_corporation_id_contracts_contract_id_items_type_id
	//
	// Type ID for item
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get corporations corporation ID contracts contract ID items o k body items0
func (o *GetCorporationsCorporationIDContractsContractIDItemsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIsIncluded(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIsSingleton(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRecordID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTypeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsOKBodyItems0) validateIsIncluded(formats strfmt.Registry) error {

	if err := validate.Required("is_included", "body", o.IsIncluded); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsOKBodyItems0) validateIsSingleton(formats strfmt.Registry) error {

	if err := validate.Required("is_singleton", "body", o.IsSingleton); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsOKBodyItems0) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", o.Quantity); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsOKBodyItems0) validateRecordID(formats strfmt.Registry) error {

	if err := validate.Required("record_id", "body", o.RecordID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDContractsContractIDItemsOKBodyItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get corporations corporation ID contracts contract ID items o k body items0 based on context it is used
func (o *GetCorporationsCorporationIDContractsContractIDItemsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDContractsContractIDItemsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDContractsContractIDItemsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDContractsContractIDItemsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetCorporationsCorporationIDContractsContractIDItemsStatus520Body get_corporations_corporation_id_contracts_contract_id_items_error_520
//
// Error 520
swagger:model GetCorporationsCorporationIDContractsContractIDItemsStatus520Body
*/
type GetCorporationsCorporationIDContractsContractIDItemsStatus520Body struct {

	// get_corporations_corporation_id_contracts_contract_id_items_520_error_520
	//
	// Error 520 message
	Error string `json:"error,omitempty"`
}

// Validate validates this get corporations corporation ID contracts contract ID items status520 body
func (o *GetCorporationsCorporationIDContractsContractIDItemsStatus520Body) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get corporations corporation ID contracts contract ID items status520 body based on context it is used
func (o *GetCorporationsCorporationIDContractsContractIDItemsStatus520Body) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDContractsContractIDItemsStatus520Body) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDContractsContractIDItemsStatus520Body) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDContractsContractIDItemsStatus520Body
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
