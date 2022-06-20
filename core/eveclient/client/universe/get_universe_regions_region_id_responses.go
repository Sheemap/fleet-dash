// Code generated by go-swagger; DO NOT EDIT.

package universe

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

// GetUniverseRegionsRegionIDReader is a Reader for the GetUniverseRegionsRegionID structure.
type GetUniverseRegionsRegionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseRegionsRegionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUniverseRegionsRegionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetUniverseRegionsRegionIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetUniverseRegionsRegionIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUniverseRegionsRegionIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetUniverseRegionsRegionIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUniverseRegionsRegionIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUniverseRegionsRegionIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUniverseRegionsRegionIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUniverseRegionsRegionIDOK creates a GetUniverseRegionsRegionIDOK with default headers values
func NewGetUniverseRegionsRegionIDOK() *GetUniverseRegionsRegionIDOK {
	return &GetUniverseRegionsRegionIDOK{}
}

/* GetUniverseRegionsRegionIDOK describes a response with status code 200, with default header values.

Information about a region
*/
type GetUniverseRegionsRegionIDOK struct {

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

	Payload *GetUniverseRegionsRegionIDOKBody
}

func (o *GetUniverseRegionsRegionIDOK) Error() string {
	return fmt.Sprintf("[GET /universe/regions/{region_id}/][%d] getUniverseRegionsRegionIdOK  %+v", 200, o.Payload)
}
func (o *GetUniverseRegionsRegionIDOK) GetPayload() *GetUniverseRegionsRegionIDOKBody {
	return o.Payload
}

func (o *GetUniverseRegionsRegionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetUniverseRegionsRegionIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseRegionsRegionIDNotModified creates a GetUniverseRegionsRegionIDNotModified with default headers values
func NewGetUniverseRegionsRegionIDNotModified() *GetUniverseRegionsRegionIDNotModified {
	return &GetUniverseRegionsRegionIDNotModified{}
}

/* GetUniverseRegionsRegionIDNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetUniverseRegionsRegionIDNotModified struct {

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

func (o *GetUniverseRegionsRegionIDNotModified) Error() string {
	return fmt.Sprintf("[GET /universe/regions/{region_id}/][%d] getUniverseRegionsRegionIdNotModified ", 304)
}

func (o *GetUniverseRegionsRegionIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseRegionsRegionIDBadRequest creates a GetUniverseRegionsRegionIDBadRequest with default headers values
func NewGetUniverseRegionsRegionIDBadRequest() *GetUniverseRegionsRegionIDBadRequest {
	return &GetUniverseRegionsRegionIDBadRequest{}
}

/* GetUniverseRegionsRegionIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUniverseRegionsRegionIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetUniverseRegionsRegionIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /universe/regions/{region_id}/][%d] getUniverseRegionsRegionIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetUniverseRegionsRegionIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetUniverseRegionsRegionIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseRegionsRegionIDNotFound creates a GetUniverseRegionsRegionIDNotFound with default headers values
func NewGetUniverseRegionsRegionIDNotFound() *GetUniverseRegionsRegionIDNotFound {
	return &GetUniverseRegionsRegionIDNotFound{}
}

/* GetUniverseRegionsRegionIDNotFound describes a response with status code 404, with default header values.

Region not found
*/
type GetUniverseRegionsRegionIDNotFound struct {
	Payload *GetUniverseRegionsRegionIDNotFoundBody
}

func (o *GetUniverseRegionsRegionIDNotFound) Error() string {
	return fmt.Sprintf("[GET /universe/regions/{region_id}/][%d] getUniverseRegionsRegionIdNotFound  %+v", 404, o.Payload)
}
func (o *GetUniverseRegionsRegionIDNotFound) GetPayload() *GetUniverseRegionsRegionIDNotFoundBody {
	return o.Payload
}

func (o *GetUniverseRegionsRegionIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetUniverseRegionsRegionIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseRegionsRegionIDEnhanceYourCalm creates a GetUniverseRegionsRegionIDEnhanceYourCalm with default headers values
func NewGetUniverseRegionsRegionIDEnhanceYourCalm() *GetUniverseRegionsRegionIDEnhanceYourCalm {
	return &GetUniverseRegionsRegionIDEnhanceYourCalm{}
}

/* GetUniverseRegionsRegionIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetUniverseRegionsRegionIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetUniverseRegionsRegionIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /universe/regions/{region_id}/][%d] getUniverseRegionsRegionIdEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetUniverseRegionsRegionIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetUniverseRegionsRegionIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseRegionsRegionIDInternalServerError creates a GetUniverseRegionsRegionIDInternalServerError with default headers values
func NewGetUniverseRegionsRegionIDInternalServerError() *GetUniverseRegionsRegionIDInternalServerError {
	return &GetUniverseRegionsRegionIDInternalServerError{}
}

/* GetUniverseRegionsRegionIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUniverseRegionsRegionIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUniverseRegionsRegionIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/regions/{region_id}/][%d] getUniverseRegionsRegionIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetUniverseRegionsRegionIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUniverseRegionsRegionIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseRegionsRegionIDServiceUnavailable creates a GetUniverseRegionsRegionIDServiceUnavailable with default headers values
func NewGetUniverseRegionsRegionIDServiceUnavailable() *GetUniverseRegionsRegionIDServiceUnavailable {
	return &GetUniverseRegionsRegionIDServiceUnavailable{}
}

/* GetUniverseRegionsRegionIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetUniverseRegionsRegionIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetUniverseRegionsRegionIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /universe/regions/{region_id}/][%d] getUniverseRegionsRegionIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetUniverseRegionsRegionIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetUniverseRegionsRegionIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseRegionsRegionIDGatewayTimeout creates a GetUniverseRegionsRegionIDGatewayTimeout with default headers values
func NewGetUniverseRegionsRegionIDGatewayTimeout() *GetUniverseRegionsRegionIDGatewayTimeout {
	return &GetUniverseRegionsRegionIDGatewayTimeout{}
}

/* GetUniverseRegionsRegionIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetUniverseRegionsRegionIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetUniverseRegionsRegionIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /universe/regions/{region_id}/][%d] getUniverseRegionsRegionIdGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetUniverseRegionsRegionIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetUniverseRegionsRegionIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetUniverseRegionsRegionIDNotFoundBody get_universe_regions_region_id_not_found
//
// Not found
swagger:model GetUniverseRegionsRegionIDNotFoundBody
*/
type GetUniverseRegionsRegionIDNotFoundBody struct {

	// get_universe_regions_region_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get universe regions region ID not found body
func (o *GetUniverseRegionsRegionIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get universe regions region ID not found body based on context it is used
func (o *GetUniverseRegionsRegionIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseRegionsRegionIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseRegionsRegionIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseRegionsRegionIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUniverseRegionsRegionIDOKBody get_universe_regions_region_id_ok
//
// 200 ok object
swagger:model GetUniverseRegionsRegionIDOKBody
*/
type GetUniverseRegionsRegionIDOKBody struct {

	// get_universe_regions_region_id_constellations
	//
	// constellations array
	// Required: true
	// Max Items: 1000
	Constellations []int32 `json:"constellations"`

	// get_universe_regions_region_id_description
	//
	// description string
	Description string `json:"description,omitempty"`

	// get_universe_regions_region_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_universe_regions_region_id_region_id
	//
	// region_id integer
	// Required: true
	RegionID *int32 `json:"region_id"`
}

// Validate validates this get universe regions region ID o k body
func (o *GetUniverseRegionsRegionIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConstellations(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRegionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseRegionsRegionIDOKBody) validateConstellations(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseRegionsRegionIdOK"+"."+"constellations", "body", o.Constellations); err != nil {
		return err
	}

	iConstellationsSize := int64(len(o.Constellations))

	if err := validate.MaxItems("getUniverseRegionsRegionIdOK"+"."+"constellations", "body", iConstellationsSize, 1000); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseRegionsRegionIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseRegionsRegionIdOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseRegionsRegionIDOKBody) validateRegionID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseRegionsRegionIdOK"+"."+"region_id", "body", o.RegionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get universe regions region ID o k body based on context it is used
func (o *GetUniverseRegionsRegionIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseRegionsRegionIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseRegionsRegionIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseRegionsRegionIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}