// Code generated by go-swagger; DO NOT EDIT.

package corporation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
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

// GetCorporationsCorporationIDRolesReader is a Reader for the GetCorporationsCorporationIDRoles structure.
type GetCorporationsCorporationIDRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationsCorporationIDRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationsCorporationIDRolesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationsCorporationIDRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationsCorporationIDRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationsCorporationIDRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationsCorporationIDRolesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationsCorporationIDRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationsCorporationIDRolesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationsCorporationIDRolesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDRolesOK creates a GetCorporationsCorporationIDRolesOK with default headers values
func NewGetCorporationsCorporationIDRolesOK() *GetCorporationsCorporationIDRolesOK {
	return &GetCorporationsCorporationIDRolesOK{}
}

/* GetCorporationsCorporationIDRolesOK describes a response with status code 200, with default header values.

List of member character ID's and roles
*/
type GetCorporationsCorporationIDRolesOK struct {

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

	Payload []*GetCorporationsCorporationIDRolesOKBodyItems0
}

func (o *GetCorporationsCorporationIDRolesOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/roles/][%d] getCorporationsCorporationIdRolesOK  %+v", 200, o.Payload)
}
func (o *GetCorporationsCorporationIDRolesOK) GetPayload() []*GetCorporationsCorporationIDRolesOKBodyItems0 {
	return o.Payload
}

func (o *GetCorporationsCorporationIDRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDRolesNotModified creates a GetCorporationsCorporationIDRolesNotModified with default headers values
func NewGetCorporationsCorporationIDRolesNotModified() *GetCorporationsCorporationIDRolesNotModified {
	return &GetCorporationsCorporationIDRolesNotModified{}
}

/* GetCorporationsCorporationIDRolesNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationsCorporationIDRolesNotModified struct {

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

func (o *GetCorporationsCorporationIDRolesNotModified) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/roles/][%d] getCorporationsCorporationIdRolesNotModified ", 304)
}

func (o *GetCorporationsCorporationIDRolesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDRolesBadRequest creates a GetCorporationsCorporationIDRolesBadRequest with default headers values
func NewGetCorporationsCorporationIDRolesBadRequest() *GetCorporationsCorporationIDRolesBadRequest {
	return &GetCorporationsCorporationIDRolesBadRequest{}
}

/* GetCorporationsCorporationIDRolesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationsCorporationIDRolesBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/roles/][%d] getCorporationsCorporationIdRolesBadRequest  %+v", 400, o.Payload)
}
func (o *GetCorporationsCorporationIDRolesBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationsCorporationIDRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDRolesUnauthorized creates a GetCorporationsCorporationIDRolesUnauthorized with default headers values
func NewGetCorporationsCorporationIDRolesUnauthorized() *GetCorporationsCorporationIDRolesUnauthorized {
	return &GetCorporationsCorporationIDRolesUnauthorized{}
}

/* GetCorporationsCorporationIDRolesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDRolesUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDRolesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/roles/][%d] getCorporationsCorporationIdRolesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCorporationsCorporationIDRolesUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationsCorporationIDRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDRolesForbidden creates a GetCorporationsCorporationIDRolesForbidden with default headers values
func NewGetCorporationsCorporationIDRolesForbidden() *GetCorporationsCorporationIDRolesForbidden {
	return &GetCorporationsCorporationIDRolesForbidden{}
}

/* GetCorporationsCorporationIDRolesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationsCorporationIDRolesForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/roles/][%d] getCorporationsCorporationIdRolesForbidden  %+v", 403, o.Payload)
}
func (o *GetCorporationsCorporationIDRolesForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationsCorporationIDRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDRolesEnhanceYourCalm creates a GetCorporationsCorporationIDRolesEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDRolesEnhanceYourCalm() *GetCorporationsCorporationIDRolesEnhanceYourCalm {
	return &GetCorporationsCorporationIDRolesEnhanceYourCalm{}
}

/* GetCorporationsCorporationIDRolesEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationsCorporationIDRolesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDRolesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/roles/][%d] getCorporationsCorporationIdRolesEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCorporationsCorporationIDRolesEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationsCorporationIDRolesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDRolesInternalServerError creates a GetCorporationsCorporationIDRolesInternalServerError with default headers values
func NewGetCorporationsCorporationIDRolesInternalServerError() *GetCorporationsCorporationIDRolesInternalServerError {
	return &GetCorporationsCorporationIDRolesInternalServerError{}
}

/* GetCorporationsCorporationIDRolesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationsCorporationIDRolesInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/roles/][%d] getCorporationsCorporationIdRolesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCorporationsCorporationIDRolesInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationsCorporationIDRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDRolesServiceUnavailable creates a GetCorporationsCorporationIDRolesServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDRolesServiceUnavailable() *GetCorporationsCorporationIDRolesServiceUnavailable {
	return &GetCorporationsCorporationIDRolesServiceUnavailable{}
}

/* GetCorporationsCorporationIDRolesServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDRolesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDRolesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/roles/][%d] getCorporationsCorporationIdRolesServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCorporationsCorporationIDRolesServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationsCorporationIDRolesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDRolesGatewayTimeout creates a GetCorporationsCorporationIDRolesGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDRolesGatewayTimeout() *GetCorporationsCorporationIDRolesGatewayTimeout {
	return &GetCorporationsCorporationIDRolesGatewayTimeout{}
}

/* GetCorporationsCorporationIDRolesGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDRolesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDRolesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/roles/][%d] getCorporationsCorporationIdRolesGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCorporationsCorporationIDRolesGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationsCorporationIDRolesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDRolesOKBodyItems0 get_corporations_corporation_id_roles_200_ok
//
// 200 ok object
swagger:model GetCorporationsCorporationIDRolesOKBodyItems0
*/
type GetCorporationsCorporationIDRolesOKBodyItems0 struct {

	// get_corporations_corporation_id_roles_character_id
	//
	// character_id integer
	// Required: true
	CharacterID *int32 `json:"character_id"`

	// get_corporations_corporation_id_roles_grantable_roles
	//
	// grantable_roles array
	// Max Items: 50
	GrantableRoles []string `json:"grantable_roles"`

	// get_corporations_corporation_id_roles_grantable_roles_at_base
	//
	// grantable_roles_at_base array
	// Max Items: 50
	GrantableRolesAtBase []string `json:"grantable_roles_at_base"`

	// get_corporations_corporation_id_roles_grantable_roles_at_hq
	//
	// grantable_roles_at_hq array
	// Max Items: 50
	GrantableRolesAtHq []string `json:"grantable_roles_at_hq"`

	// get_corporations_corporation_id_roles_grantable_roles_at_other
	//
	// grantable_roles_at_other array
	// Max Items: 50
	GrantableRolesAtOther []string `json:"grantable_roles_at_other"`

	// get_corporations_corporation_id_roles_roles
	//
	// roles array
	// Max Items: 50
	Roles []string `json:"roles"`

	// get_corporations_corporation_id_roles_roles_at_base
	//
	// roles_at_base array
	// Max Items: 50
	RolesAtBase []string `json:"roles_at_base"`

	// get_corporations_corporation_id_roles_roles_at_hq
	//
	// roles_at_hq array
	// Max Items: 50
	RolesAtHq []string `json:"roles_at_hq"`

	// get_corporations_corporation_id_roles_roles_at_other
	//
	// roles_at_other array
	// Max Items: 50
	RolesAtOther []string `json:"roles_at_other"`
}

// Validate validates this get corporations corporation ID roles o k body items0
func (o *GetCorporationsCorporationIDRolesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCharacterID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGrantableRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGrantableRolesAtBase(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGrantableRolesAtHq(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGrantableRolesAtOther(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRolesAtBase(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRolesAtHq(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRolesAtOther(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateCharacterID(formats strfmt.Registry) error {

	if err := validate.Required("character_id", "body", o.CharacterID); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Accountant","Auditor","Communications_Officer","Config_Equipment","Config_Starbase_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Contract_Manager","Diplomat","Director","Factory_Manager","Fitting_Manager","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Junior_Accountant","Personnel_Manager","Rent_Factory_Facility","Rent_Office","Rent_Research_Facility","Security_Officer","Starbase_Defense_Operator","Starbase_Fuel_Technician","Station_Manager","Trader"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesItemsEnum = append(getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesItemsEnum, v)
	}
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateGrantableRolesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateGrantableRoles(formats strfmt.Registry) error {
	if swag.IsZero(o.GrantableRoles) { // not required
		return nil
	}

	iGrantableRolesSize := int64(len(o.GrantableRoles))

	if err := validate.MaxItems("grantable_roles", "body", iGrantableRolesSize, 50); err != nil {
		return err
	}

	for i := 0; i < len(o.GrantableRoles); i++ {

		// value enum
		if err := o.validateGrantableRolesItemsEnum("grantable_roles"+"."+strconv.Itoa(i), "body", o.GrantableRoles[i]); err != nil {
			return err
		}

	}

	return nil
}

var getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtBaseItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Accountant","Auditor","Communications_Officer","Config_Equipment","Config_Starbase_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Contract_Manager","Diplomat","Director","Factory_Manager","Fitting_Manager","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Junior_Accountant","Personnel_Manager","Rent_Factory_Facility","Rent_Office","Rent_Research_Facility","Security_Officer","Starbase_Defense_Operator","Starbase_Fuel_Technician","Station_Manager","Trader"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtBaseItemsEnum = append(getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtBaseItemsEnum, v)
	}
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateGrantableRolesAtBaseItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtBaseItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateGrantableRolesAtBase(formats strfmt.Registry) error {
	if swag.IsZero(o.GrantableRolesAtBase) { // not required
		return nil
	}

	iGrantableRolesAtBaseSize := int64(len(o.GrantableRolesAtBase))

	if err := validate.MaxItems("grantable_roles_at_base", "body", iGrantableRolesAtBaseSize, 50); err != nil {
		return err
	}

	for i := 0; i < len(o.GrantableRolesAtBase); i++ {

		// value enum
		if err := o.validateGrantableRolesAtBaseItemsEnum("grantable_roles_at_base"+"."+strconv.Itoa(i), "body", o.GrantableRolesAtBase[i]); err != nil {
			return err
		}

	}

	return nil
}

var getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtHqItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Accountant","Auditor","Communications_Officer","Config_Equipment","Config_Starbase_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Contract_Manager","Diplomat","Director","Factory_Manager","Fitting_Manager","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Junior_Accountant","Personnel_Manager","Rent_Factory_Facility","Rent_Office","Rent_Research_Facility","Security_Officer","Starbase_Defense_Operator","Starbase_Fuel_Technician","Station_Manager","Trader"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtHqItemsEnum = append(getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtHqItemsEnum, v)
	}
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateGrantableRolesAtHqItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtHqItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateGrantableRolesAtHq(formats strfmt.Registry) error {
	if swag.IsZero(o.GrantableRolesAtHq) { // not required
		return nil
	}

	iGrantableRolesAtHqSize := int64(len(o.GrantableRolesAtHq))

	if err := validate.MaxItems("grantable_roles_at_hq", "body", iGrantableRolesAtHqSize, 50); err != nil {
		return err
	}

	for i := 0; i < len(o.GrantableRolesAtHq); i++ {

		// value enum
		if err := o.validateGrantableRolesAtHqItemsEnum("grantable_roles_at_hq"+"."+strconv.Itoa(i), "body", o.GrantableRolesAtHq[i]); err != nil {
			return err
		}

	}

	return nil
}

var getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtOtherItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Accountant","Auditor","Communications_Officer","Config_Equipment","Config_Starbase_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Contract_Manager","Diplomat","Director","Factory_Manager","Fitting_Manager","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Junior_Accountant","Personnel_Manager","Rent_Factory_Facility","Rent_Office","Rent_Research_Facility","Security_Officer","Starbase_Defense_Operator","Starbase_Fuel_Technician","Station_Manager","Trader"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtOtherItemsEnum = append(getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtOtherItemsEnum, v)
	}
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateGrantableRolesAtOtherItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCorporationsCorporationIdRolesOKBodyItems0GrantableRolesAtOtherItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateGrantableRolesAtOther(formats strfmt.Registry) error {
	if swag.IsZero(o.GrantableRolesAtOther) { // not required
		return nil
	}

	iGrantableRolesAtOtherSize := int64(len(o.GrantableRolesAtOther))

	if err := validate.MaxItems("grantable_roles_at_other", "body", iGrantableRolesAtOtherSize, 50); err != nil {
		return err
	}

	for i := 0; i < len(o.GrantableRolesAtOther); i++ {

		// value enum
		if err := o.validateGrantableRolesAtOtherItemsEnum("grantable_roles_at_other"+"."+strconv.Itoa(i), "body", o.GrantableRolesAtOther[i]); err != nil {
			return err
		}

	}

	return nil
}

var getCorporationsCorporationIdRolesOKBodyItems0RolesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Accountant","Auditor","Communications_Officer","Config_Equipment","Config_Starbase_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Contract_Manager","Diplomat","Director","Factory_Manager","Fitting_Manager","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Junior_Accountant","Personnel_Manager","Rent_Factory_Facility","Rent_Office","Rent_Research_Facility","Security_Officer","Starbase_Defense_Operator","Starbase_Fuel_Technician","Station_Manager","Trader"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdRolesOKBodyItems0RolesItemsEnum = append(getCorporationsCorporationIdRolesOKBodyItems0RolesItemsEnum, v)
	}
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateRolesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCorporationsCorporationIdRolesOKBodyItems0RolesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateRoles(formats strfmt.Registry) error {
	if swag.IsZero(o.Roles) { // not required
		return nil
	}

	iRolesSize := int64(len(o.Roles))

	if err := validate.MaxItems("roles", "body", iRolesSize, 50); err != nil {
		return err
	}

	for i := 0; i < len(o.Roles); i++ {

		// value enum
		if err := o.validateRolesItemsEnum("roles"+"."+strconv.Itoa(i), "body", o.Roles[i]); err != nil {
			return err
		}

	}

	return nil
}

var getCorporationsCorporationIdRolesOKBodyItems0RolesAtBaseItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Accountant","Auditor","Communications_Officer","Config_Equipment","Config_Starbase_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Contract_Manager","Diplomat","Director","Factory_Manager","Fitting_Manager","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Junior_Accountant","Personnel_Manager","Rent_Factory_Facility","Rent_Office","Rent_Research_Facility","Security_Officer","Starbase_Defense_Operator","Starbase_Fuel_Technician","Station_Manager","Trader"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdRolesOKBodyItems0RolesAtBaseItemsEnum = append(getCorporationsCorporationIdRolesOKBodyItems0RolesAtBaseItemsEnum, v)
	}
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateRolesAtBaseItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCorporationsCorporationIdRolesOKBodyItems0RolesAtBaseItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateRolesAtBase(formats strfmt.Registry) error {
	if swag.IsZero(o.RolesAtBase) { // not required
		return nil
	}

	iRolesAtBaseSize := int64(len(o.RolesAtBase))

	if err := validate.MaxItems("roles_at_base", "body", iRolesAtBaseSize, 50); err != nil {
		return err
	}

	for i := 0; i < len(o.RolesAtBase); i++ {

		// value enum
		if err := o.validateRolesAtBaseItemsEnum("roles_at_base"+"."+strconv.Itoa(i), "body", o.RolesAtBase[i]); err != nil {
			return err
		}

	}

	return nil
}

var getCorporationsCorporationIdRolesOKBodyItems0RolesAtHqItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Accountant","Auditor","Communications_Officer","Config_Equipment","Config_Starbase_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Contract_Manager","Diplomat","Director","Factory_Manager","Fitting_Manager","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Junior_Accountant","Personnel_Manager","Rent_Factory_Facility","Rent_Office","Rent_Research_Facility","Security_Officer","Starbase_Defense_Operator","Starbase_Fuel_Technician","Station_Manager","Trader"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdRolesOKBodyItems0RolesAtHqItemsEnum = append(getCorporationsCorporationIdRolesOKBodyItems0RolesAtHqItemsEnum, v)
	}
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateRolesAtHqItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCorporationsCorporationIdRolesOKBodyItems0RolesAtHqItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateRolesAtHq(formats strfmt.Registry) error {
	if swag.IsZero(o.RolesAtHq) { // not required
		return nil
	}

	iRolesAtHqSize := int64(len(o.RolesAtHq))

	if err := validate.MaxItems("roles_at_hq", "body", iRolesAtHqSize, 50); err != nil {
		return err
	}

	for i := 0; i < len(o.RolesAtHq); i++ {

		// value enum
		if err := o.validateRolesAtHqItemsEnum("roles_at_hq"+"."+strconv.Itoa(i), "body", o.RolesAtHq[i]); err != nil {
			return err
		}

	}

	return nil
}

var getCorporationsCorporationIdRolesOKBodyItems0RolesAtOtherItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Accountant","Auditor","Communications_Officer","Config_Equipment","Config_Starbase_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Contract_Manager","Diplomat","Director","Factory_Manager","Fitting_Manager","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Junior_Accountant","Personnel_Manager","Rent_Factory_Facility","Rent_Office","Rent_Research_Facility","Security_Officer","Starbase_Defense_Operator","Starbase_Fuel_Technician","Station_Manager","Trader"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdRolesOKBodyItems0RolesAtOtherItemsEnum = append(getCorporationsCorporationIdRolesOKBodyItems0RolesAtOtherItemsEnum, v)
	}
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateRolesAtOtherItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCorporationsCorporationIdRolesOKBodyItems0RolesAtOtherItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesOKBodyItems0) validateRolesAtOther(formats strfmt.Registry) error {
	if swag.IsZero(o.RolesAtOther) { // not required
		return nil
	}

	iRolesAtOtherSize := int64(len(o.RolesAtOther))

	if err := validate.MaxItems("roles_at_other", "body", iRolesAtOtherSize, 50); err != nil {
		return err
	}

	for i := 0; i < len(o.RolesAtOther); i++ {

		// value enum
		if err := o.validateRolesAtOtherItemsEnum("roles_at_other"+"."+strconv.Itoa(i), "body", o.RolesAtOther[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this get corporations corporation ID roles o k body items0 based on context it is used
func (o *GetCorporationsCorporationIDRolesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDRolesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDRolesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDRolesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
