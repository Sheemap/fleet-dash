// Code generated by go-swagger; DO NOT EDIT.

package bookmarks

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

// GetCorporationsCorporationIDBookmarksFoldersReader is a Reader for the GetCorporationsCorporationIDBookmarksFolders structure.
type GetCorporationsCorporationIDBookmarksFoldersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDBookmarksFoldersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationsCorporationIDBookmarksFoldersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationsCorporationIDBookmarksFoldersNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationsCorporationIDBookmarksFoldersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationsCorporationIDBookmarksFoldersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationsCorporationIDBookmarksFoldersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationsCorporationIDBookmarksFoldersEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationsCorporationIDBookmarksFoldersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationsCorporationIDBookmarksFoldersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationsCorporationIDBookmarksFoldersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDBookmarksFoldersOK creates a GetCorporationsCorporationIDBookmarksFoldersOK with default headers values
func NewGetCorporationsCorporationIDBookmarksFoldersOK() *GetCorporationsCorporationIDBookmarksFoldersOK {
	var (
		// initialize headers with default values
		xPagesDefault = int32(1)
	)

	return &GetCorporationsCorporationIDBookmarksFoldersOK{

		XPages: xPagesDefault,
	}
}

/* GetCorporationsCorporationIDBookmarksFoldersOK describes a response with status code 200, with default header values.

List of corporation owned bookmark folders
*/
type GetCorporationsCorporationIDBookmarksFoldersOK struct {

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

	/* Maximum page number

	   Format: int32
	   Default: 1
	*/
	XPages int32

	Payload []*GetCorporationsCorporationIDBookmarksFoldersOKBodyItems0
}

func (o *GetCorporationsCorporationIDBookmarksFoldersOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/bookmarks/folders/][%d] getCorporationsCorporationIdBookmarksFoldersOK  %+v", 200, o.Payload)
}
func (o *GetCorporationsCorporationIDBookmarksFoldersOK) GetPayload() []*GetCorporationsCorporationIDBookmarksFoldersOKBodyItems0 {
	return o.Payload
}

func (o *GetCorporationsCorporationIDBookmarksFoldersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	// hydrates response header X-Pages
	hdrXPages := response.GetHeader("X-Pages")

	if hdrXPages != "" {
		valxPages, err := swag.ConvertInt32(hdrXPages)
		if err != nil {
			return errors.InvalidType("X-Pages", "header", "int32", hdrXPages)
		}
		o.XPages = valxPages
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDBookmarksFoldersNotModified creates a GetCorporationsCorporationIDBookmarksFoldersNotModified with default headers values
func NewGetCorporationsCorporationIDBookmarksFoldersNotModified() *GetCorporationsCorporationIDBookmarksFoldersNotModified {
	return &GetCorporationsCorporationIDBookmarksFoldersNotModified{}
}

/* GetCorporationsCorporationIDBookmarksFoldersNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationsCorporationIDBookmarksFoldersNotModified struct {

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

func (o *GetCorporationsCorporationIDBookmarksFoldersNotModified) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/bookmarks/folders/][%d] getCorporationsCorporationIdBookmarksFoldersNotModified ", 304)
}

func (o *GetCorporationsCorporationIDBookmarksFoldersNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDBookmarksFoldersBadRequest creates a GetCorporationsCorporationIDBookmarksFoldersBadRequest with default headers values
func NewGetCorporationsCorporationIDBookmarksFoldersBadRequest() *GetCorporationsCorporationIDBookmarksFoldersBadRequest {
	return &GetCorporationsCorporationIDBookmarksFoldersBadRequest{}
}

/* GetCorporationsCorporationIDBookmarksFoldersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationsCorporationIDBookmarksFoldersBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDBookmarksFoldersBadRequest) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/bookmarks/folders/][%d] getCorporationsCorporationIdBookmarksFoldersBadRequest  %+v", 400, o.Payload)
}
func (o *GetCorporationsCorporationIDBookmarksFoldersBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationsCorporationIDBookmarksFoldersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDBookmarksFoldersUnauthorized creates a GetCorporationsCorporationIDBookmarksFoldersUnauthorized with default headers values
func NewGetCorporationsCorporationIDBookmarksFoldersUnauthorized() *GetCorporationsCorporationIDBookmarksFoldersUnauthorized {
	return &GetCorporationsCorporationIDBookmarksFoldersUnauthorized{}
}

/* GetCorporationsCorporationIDBookmarksFoldersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDBookmarksFoldersUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDBookmarksFoldersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/bookmarks/folders/][%d] getCorporationsCorporationIdBookmarksFoldersUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCorporationsCorporationIDBookmarksFoldersUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationsCorporationIDBookmarksFoldersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDBookmarksFoldersForbidden creates a GetCorporationsCorporationIDBookmarksFoldersForbidden with default headers values
func NewGetCorporationsCorporationIDBookmarksFoldersForbidden() *GetCorporationsCorporationIDBookmarksFoldersForbidden {
	return &GetCorporationsCorporationIDBookmarksFoldersForbidden{}
}

/* GetCorporationsCorporationIDBookmarksFoldersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationsCorporationIDBookmarksFoldersForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDBookmarksFoldersForbidden) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/bookmarks/folders/][%d] getCorporationsCorporationIdBookmarksFoldersForbidden  %+v", 403, o.Payload)
}
func (o *GetCorporationsCorporationIDBookmarksFoldersForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationsCorporationIDBookmarksFoldersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDBookmarksFoldersEnhanceYourCalm creates a GetCorporationsCorporationIDBookmarksFoldersEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDBookmarksFoldersEnhanceYourCalm() *GetCorporationsCorporationIDBookmarksFoldersEnhanceYourCalm {
	return &GetCorporationsCorporationIDBookmarksFoldersEnhanceYourCalm{}
}

/* GetCorporationsCorporationIDBookmarksFoldersEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationsCorporationIDBookmarksFoldersEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDBookmarksFoldersEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/bookmarks/folders/][%d] getCorporationsCorporationIdBookmarksFoldersEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCorporationsCorporationIDBookmarksFoldersEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationsCorporationIDBookmarksFoldersEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDBookmarksFoldersInternalServerError creates a GetCorporationsCorporationIDBookmarksFoldersInternalServerError with default headers values
func NewGetCorporationsCorporationIDBookmarksFoldersInternalServerError() *GetCorporationsCorporationIDBookmarksFoldersInternalServerError {
	return &GetCorporationsCorporationIDBookmarksFoldersInternalServerError{}
}

/* GetCorporationsCorporationIDBookmarksFoldersInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationsCorporationIDBookmarksFoldersInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDBookmarksFoldersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/bookmarks/folders/][%d] getCorporationsCorporationIdBookmarksFoldersInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCorporationsCorporationIDBookmarksFoldersInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationsCorporationIDBookmarksFoldersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDBookmarksFoldersServiceUnavailable creates a GetCorporationsCorporationIDBookmarksFoldersServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDBookmarksFoldersServiceUnavailable() *GetCorporationsCorporationIDBookmarksFoldersServiceUnavailable {
	return &GetCorporationsCorporationIDBookmarksFoldersServiceUnavailable{}
}

/* GetCorporationsCorporationIDBookmarksFoldersServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDBookmarksFoldersServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDBookmarksFoldersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/bookmarks/folders/][%d] getCorporationsCorporationIdBookmarksFoldersServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCorporationsCorporationIDBookmarksFoldersServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationsCorporationIDBookmarksFoldersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDBookmarksFoldersGatewayTimeout creates a GetCorporationsCorporationIDBookmarksFoldersGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDBookmarksFoldersGatewayTimeout() *GetCorporationsCorporationIDBookmarksFoldersGatewayTimeout {
	return &GetCorporationsCorporationIDBookmarksFoldersGatewayTimeout{}
}

/* GetCorporationsCorporationIDBookmarksFoldersGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDBookmarksFoldersGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDBookmarksFoldersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/bookmarks/folders/][%d] getCorporationsCorporationIdBookmarksFoldersGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCorporationsCorporationIDBookmarksFoldersGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationsCorporationIDBookmarksFoldersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDBookmarksFoldersOKBodyItems0 get_corporations_corporation_id_bookmarks_folders_200_ok
//
// 200 ok object
swagger:model GetCorporationsCorporationIDBookmarksFoldersOKBodyItems0
*/
type GetCorporationsCorporationIDBookmarksFoldersOKBodyItems0 struct {

	// get_corporations_corporation_id_bookmarks_folders_creator_id
	//
	// creator_id integer
	CreatorID int32 `json:"creator_id,omitempty"`

	// get_corporations_corporation_id_bookmarks_folders_folder_id
	//
	// folder_id integer
	// Required: true
	FolderID *int32 `json:"folder_id"`

	// get_corporations_corporation_id_bookmarks_folders_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this get corporations corporation ID bookmarks folders o k body items0
func (o *GetCorporationsCorporationIDBookmarksFoldersOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFolderID(formats); err != nil {
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

func (o *GetCorporationsCorporationIDBookmarksFoldersOKBodyItems0) validateFolderID(formats strfmt.Registry) error {

	if err := validate.Required("folder_id", "body", o.FolderID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDBookmarksFoldersOKBodyItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get corporations corporation ID bookmarks folders o k body items0 based on context it is used
func (o *GetCorporationsCorporationIDBookmarksFoldersOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDBookmarksFoldersOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDBookmarksFoldersOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDBookmarksFoldersOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
