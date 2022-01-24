// Code generated by go-swagger; DO NOT EDIT.

package universe

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

// GetUniverseTypesTypeIDReader is a Reader for the GetUniverseTypesTypeID structure.
type GetUniverseTypesTypeIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseTypesTypeIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUniverseTypesTypeIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetUniverseTypesTypeIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetUniverseTypesTypeIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUniverseTypesTypeIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetUniverseTypesTypeIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUniverseTypesTypeIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUniverseTypesTypeIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUniverseTypesTypeIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUniverseTypesTypeIDOK creates a GetUniverseTypesTypeIDOK with default headers values
func NewGetUniverseTypesTypeIDOK() *GetUniverseTypesTypeIDOK {
	return &GetUniverseTypesTypeIDOK{}
}

/* GetUniverseTypesTypeIDOK describes a response with status code 200, with default header values.

Information about a type
*/
type GetUniverseTypesTypeIDOK struct {

	/* The caching mechanism used
	 */
	CacheControl string

	/* The language used in the response
	 */
	ContentLanguage string

	/* RFC7232 compliant entity tag
	 */
	ETag string

	/* RFC7231 formatted datetime string
	 */
	Expires string

	/* RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *GetUniverseTypesTypeIDOKBody
}

func (o *GetUniverseTypesTypeIDOK) Error() string {
	return fmt.Sprintf("[GET /universe/types/{type_id}/][%d] getUniverseTypesTypeIdOK  %+v", 200, o.Payload)
}
func (o *GetUniverseTypesTypeIDOK) GetPayload() *GetUniverseTypesTypeIDOKBody {
	return o.Payload
}

func (o *GetUniverseTypesTypeIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Cache-Control
	hdrCacheControl := response.GetHeader("Cache-Control")

	if hdrCacheControl != "" {
		o.CacheControl = hdrCacheControl
	}

	// hydrates response header Content-Language
	hdrContentLanguage := response.GetHeader("Content-Language")

	if hdrContentLanguage != "" {
		o.ContentLanguage = hdrContentLanguage
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

	o.Payload = new(GetUniverseTypesTypeIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseTypesTypeIDNotModified creates a GetUniverseTypesTypeIDNotModified with default headers values
func NewGetUniverseTypesTypeIDNotModified() *GetUniverseTypesTypeIDNotModified {
	return &GetUniverseTypesTypeIDNotModified{}
}

/* GetUniverseTypesTypeIDNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetUniverseTypesTypeIDNotModified struct {

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

func (o *GetUniverseTypesTypeIDNotModified) Error() string {
	return fmt.Sprintf("[GET /universe/types/{type_id}/][%d] getUniverseTypesTypeIdNotModified ", 304)
}

func (o *GetUniverseTypesTypeIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseTypesTypeIDBadRequest creates a GetUniverseTypesTypeIDBadRequest with default headers values
func NewGetUniverseTypesTypeIDBadRequest() *GetUniverseTypesTypeIDBadRequest {
	return &GetUniverseTypesTypeIDBadRequest{}
}

/* GetUniverseTypesTypeIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUniverseTypesTypeIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetUniverseTypesTypeIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /universe/types/{type_id}/][%d] getUniverseTypesTypeIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetUniverseTypesTypeIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetUniverseTypesTypeIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseTypesTypeIDNotFound creates a GetUniverseTypesTypeIDNotFound with default headers values
func NewGetUniverseTypesTypeIDNotFound() *GetUniverseTypesTypeIDNotFound {
	return &GetUniverseTypesTypeIDNotFound{}
}

/* GetUniverseTypesTypeIDNotFound describes a response with status code 404, with default header values.

Type not found
*/
type GetUniverseTypesTypeIDNotFound struct {
	Payload *GetUniverseTypesTypeIDNotFoundBody
}

func (o *GetUniverseTypesTypeIDNotFound) Error() string {
	return fmt.Sprintf("[GET /universe/types/{type_id}/][%d] getUniverseTypesTypeIdNotFound  %+v", 404, o.Payload)
}
func (o *GetUniverseTypesTypeIDNotFound) GetPayload() *GetUniverseTypesTypeIDNotFoundBody {
	return o.Payload
}

func (o *GetUniverseTypesTypeIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetUniverseTypesTypeIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseTypesTypeIDEnhanceYourCalm creates a GetUniverseTypesTypeIDEnhanceYourCalm with default headers values
func NewGetUniverseTypesTypeIDEnhanceYourCalm() *GetUniverseTypesTypeIDEnhanceYourCalm {
	return &GetUniverseTypesTypeIDEnhanceYourCalm{}
}

/* GetUniverseTypesTypeIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetUniverseTypesTypeIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetUniverseTypesTypeIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /universe/types/{type_id}/][%d] getUniverseTypesTypeIdEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetUniverseTypesTypeIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetUniverseTypesTypeIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseTypesTypeIDInternalServerError creates a GetUniverseTypesTypeIDInternalServerError with default headers values
func NewGetUniverseTypesTypeIDInternalServerError() *GetUniverseTypesTypeIDInternalServerError {
	return &GetUniverseTypesTypeIDInternalServerError{}
}

/* GetUniverseTypesTypeIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUniverseTypesTypeIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUniverseTypesTypeIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/types/{type_id}/][%d] getUniverseTypesTypeIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetUniverseTypesTypeIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUniverseTypesTypeIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseTypesTypeIDServiceUnavailable creates a GetUniverseTypesTypeIDServiceUnavailable with default headers values
func NewGetUniverseTypesTypeIDServiceUnavailable() *GetUniverseTypesTypeIDServiceUnavailable {
	return &GetUniverseTypesTypeIDServiceUnavailable{}
}

/* GetUniverseTypesTypeIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetUniverseTypesTypeIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetUniverseTypesTypeIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /universe/types/{type_id}/][%d] getUniverseTypesTypeIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetUniverseTypesTypeIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetUniverseTypesTypeIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseTypesTypeIDGatewayTimeout creates a GetUniverseTypesTypeIDGatewayTimeout with default headers values
func NewGetUniverseTypesTypeIDGatewayTimeout() *GetUniverseTypesTypeIDGatewayTimeout {
	return &GetUniverseTypesTypeIDGatewayTimeout{}
}

/* GetUniverseTypesTypeIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetUniverseTypesTypeIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetUniverseTypesTypeIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /universe/types/{type_id}/][%d] getUniverseTypesTypeIdGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetUniverseTypesTypeIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetUniverseTypesTypeIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetUniverseTypesTypeIDNotFoundBody get_universe_types_type_id_not_found
//
// Not found
swagger:model GetUniverseTypesTypeIDNotFoundBody
*/
type GetUniverseTypesTypeIDNotFoundBody struct {

	// get_universe_types_type_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get universe types type ID not found body
func (o *GetUniverseTypesTypeIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get universe types type ID not found body based on context it is used
func (o *GetUniverseTypesTypeIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseTypesTypeIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseTypesTypeIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseTypesTypeIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUniverseTypesTypeIDOKBody get_universe_types_type_id_ok
//
// 200 ok object
swagger:model GetUniverseTypesTypeIDOKBody
*/
type GetUniverseTypesTypeIDOKBody struct {

	// get_universe_types_type_id_capacity
	//
	// capacity number
	Capacity float32 `json:"capacity,omitempty"`

	// get_universe_types_type_id_description
	//
	// description string
	// Required: true
	Description *string `json:"description"`

	// get_universe_types_type_id_dogma_attributes
	//
	// dogma_attributes array
	// Max Items: 1000
	DogmaAttributes []*GetUniverseTypesTypeIDOKBodyDogmaAttributesItems0 `json:"dogma_attributes"`

	// get_universe_types_type_id_dogma_effects
	//
	// dogma_effects array
	// Max Items: 1000
	DogmaEffects []*GetUniverseTypesTypeIDOKBodyDogmaEffectsItems0 `json:"dogma_effects"`

	// get_universe_types_type_id_graphic_id
	//
	// graphic_id integer
	GraphicID int32 `json:"graphic_id,omitempty"`

	// get_universe_types_type_id_group_id
	//
	// group_id integer
	// Required: true
	GroupID *int32 `json:"group_id"`

	// get_universe_types_type_id_icon_id
	//
	// icon_id integer
	IconID int32 `json:"icon_id,omitempty"`

	// get_universe_types_type_id_market_group_id
	//
	// This only exists for types that can be put on the market
	MarketGroupID int32 `json:"market_group_id,omitempty"`

	// get_universe_types_type_id_mass
	//
	// mass number
	Mass float32 `json:"mass,omitempty"`

	// get_universe_types_type_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_universe_types_type_id_packaged_volume
	//
	// packaged_volume number
	PackagedVolume float32 `json:"packaged_volume,omitempty"`

	// get_universe_types_type_id_portion_size
	//
	// portion_size integer
	PortionSize int32 `json:"portion_size,omitempty"`

	// get_universe_types_type_id_published
	//
	// published boolean
	// Required: true
	Published *bool `json:"published"`

	// get_universe_types_type_id_radius
	//
	// radius number
	Radius float32 `json:"radius,omitempty"`

	// get_universe_types_type_id_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`

	// get_universe_types_type_id_volume
	//
	// volume number
	Volume float32 `json:"volume,omitempty"`
}

// Validate validates this get universe types type ID o k body
func (o *GetUniverseTypesTypeIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDogmaAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDogmaEffects(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePublished(formats); err != nil {
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

func (o *GetUniverseTypesTypeIDOKBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseTypesTypeIdOK"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) validateDogmaAttributes(formats strfmt.Registry) error {
	if swag.IsZero(o.DogmaAttributes) { // not required
		return nil
	}

	iDogmaAttributesSize := int64(len(o.DogmaAttributes))

	if err := validate.MaxItems("getUniverseTypesTypeIdOK"+"."+"dogma_attributes", "body", iDogmaAttributesSize, 1000); err != nil {
		return err
	}

	for i := 0; i < len(o.DogmaAttributes); i++ {
		if swag.IsZero(o.DogmaAttributes[i]) { // not required
			continue
		}

		if o.DogmaAttributes[i] != nil {
			if err := o.DogmaAttributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getUniverseTypesTypeIdOK" + "." + "dogma_attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getUniverseTypesTypeIdOK" + "." + "dogma_attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) validateDogmaEffects(formats strfmt.Registry) error {
	if swag.IsZero(o.DogmaEffects) { // not required
		return nil
	}

	iDogmaEffectsSize := int64(len(o.DogmaEffects))

	if err := validate.MaxItems("getUniverseTypesTypeIdOK"+"."+"dogma_effects", "body", iDogmaEffectsSize, 1000); err != nil {
		return err
	}

	for i := 0; i < len(o.DogmaEffects); i++ {
		if swag.IsZero(o.DogmaEffects[i]) { // not required
			continue
		}

		if o.DogmaEffects[i] != nil {
			if err := o.DogmaEffects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getUniverseTypesTypeIdOK" + "." + "dogma_effects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getUniverseTypesTypeIdOK" + "." + "dogma_effects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) validateGroupID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseTypesTypeIdOK"+"."+"group_id", "body", o.GroupID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseTypesTypeIdOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) validatePublished(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseTypesTypeIdOK"+"."+"published", "body", o.Published); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseTypesTypeIdOK"+"."+"type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get universe types type ID o k body based on the context it is used
func (o *GetUniverseTypesTypeIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDogmaAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDogmaEffects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) contextValidateDogmaAttributes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.DogmaAttributes); i++ {

		if o.DogmaAttributes[i] != nil {
			if err := o.DogmaAttributes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getUniverseTypesTypeIdOK" + "." + "dogma_attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getUniverseTypesTypeIdOK" + "." + "dogma_attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBody) contextValidateDogmaEffects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.DogmaEffects); i++ {

		if o.DogmaEffects[i] != nil {
			if err := o.DogmaEffects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getUniverseTypesTypeIdOK" + "." + "dogma_effects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getUniverseTypesTypeIdOK" + "." + "dogma_effects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseTypesTypeIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseTypesTypeIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseTypesTypeIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUniverseTypesTypeIDOKBodyDogmaAttributesItems0 get_universe_types_type_id_dogma_attribute
//
// dogma_attribute object
swagger:model GetUniverseTypesTypeIDOKBodyDogmaAttributesItems0
*/
type GetUniverseTypesTypeIDOKBodyDogmaAttributesItems0 struct {

	// get_universe_types_type_id_attribute_id
	//
	// attribute_id integer
	// Required: true
	AttributeID *int32 `json:"attribute_id"`

	// get_universe_types_type_id_value
	//
	// value number
	// Required: true
	Value *float32 `json:"value"`
}

// Validate validates this get universe types type ID o k body dogma attributes items0
func (o *GetUniverseTypesTypeIDOKBodyDogmaAttributesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAttributeID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseTypesTypeIDOKBodyDogmaAttributesItems0) validateAttributeID(formats strfmt.Registry) error {

	if err := validate.Required("attribute_id", "body", o.AttributeID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBodyDogmaAttributesItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get universe types type ID o k body dogma attributes items0 based on context it is used
func (o *GetUniverseTypesTypeIDOKBodyDogmaAttributesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseTypesTypeIDOKBodyDogmaAttributesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseTypesTypeIDOKBodyDogmaAttributesItems0) UnmarshalBinary(b []byte) error {
	var res GetUniverseTypesTypeIDOKBodyDogmaAttributesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUniverseTypesTypeIDOKBodyDogmaEffectsItems0 get_universe_types_type_id_dogma_effect
//
// dogma_effect object
swagger:model GetUniverseTypesTypeIDOKBodyDogmaEffectsItems0
*/
type GetUniverseTypesTypeIDOKBodyDogmaEffectsItems0 struct {

	// get_universe_types_type_id_effect_id
	//
	// effect_id integer
	// Required: true
	EffectID *int32 `json:"effect_id"`

	// get_universe_types_type_id_is_default
	//
	// is_default boolean
	// Required: true
	IsDefault *bool `json:"is_default"`
}

// Validate validates this get universe types type ID o k body dogma effects items0
func (o *GetUniverseTypesTypeIDOKBodyDogmaEffectsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEffectID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIsDefault(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseTypesTypeIDOKBodyDogmaEffectsItems0) validateEffectID(formats strfmt.Registry) error {

	if err := validate.Required("effect_id", "body", o.EffectID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseTypesTypeIDOKBodyDogmaEffectsItems0) validateIsDefault(formats strfmt.Registry) error {

	if err := validate.Required("is_default", "body", o.IsDefault); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get universe types type ID o k body dogma effects items0 based on context it is used
func (o *GetUniverseTypesTypeIDOKBodyDogmaEffectsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseTypesTypeIDOKBodyDogmaEffectsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseTypesTypeIDOKBodyDogmaEffectsItems0) UnmarshalBinary(b []byte) error {
	var res GetUniverseTypesTypeIDOKBodyDogmaEffectsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
