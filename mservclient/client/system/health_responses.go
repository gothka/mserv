// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/TykTechnologies/mserv/mservclient/models"
)

// HealthReader is a Reader for the Health structure.
type HealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewHealthInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHealthOK creates a HealthOK with default headers values
func NewHealthOK() *HealthOK {
	return &HealthOK{}
}

/*HealthOK handles this case with default header values.

Health status response
*/
type HealthOK struct {
	Payload *HealthOKBody
}

func (o *HealthOK) Error() string {
	return fmt.Sprintf("[GET /health][%d] healthOK  %+v", 200, o.Payload)
}

func (o *HealthOK) GetPayload() *HealthOKBody {
	return o.Payload
}

func (o *HealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(HealthOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHealthInternalServerError creates a HealthInternalServerError with default headers values
func NewHealthInternalServerError() *HealthInternalServerError {
	return &HealthInternalServerError{}
}

/*HealthInternalServerError handles this case with default header values.

Health status response
*/
type HealthInternalServerError struct {
	Payload *HealthInternalServerErrorBody
}

func (o *HealthInternalServerError) Error() string {
	return fmt.Sprintf("[GET /health][%d] healthInternalServerError  %+v", 500, o.Payload)
}

func (o *HealthInternalServerError) GetPayload() *HealthInternalServerErrorBody {
	return o.Payload
}

func (o *HealthInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(HealthInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*HealthInternalServerErrorBody health internal server error body
swagger:model HealthInternalServerErrorBody
*/
type HealthInternalServerErrorBody struct {

	// error
	Error string `json:"Error,omitempty"`

	// payload
	Payload *models.HReport `json:"Payload,omitempty"`

	// status
	Status string `json:"Status,omitempty"`
}

// Validate validates this health internal server error body
func (o *HealthInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *HealthInternalServerErrorBody) validatePayload(formats strfmt.Registry) error {

	if swag.IsZero(o.Payload) { // not required
		return nil
	}

	if o.Payload != nil {
		if err := o.Payload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("healthInternalServerError" + "." + "Payload")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *HealthInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *HealthInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res HealthInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*HealthOKBody health o k body
swagger:model HealthOKBody
*/
type HealthOKBody struct {

	// error
	Error string `json:"Error,omitempty"`

	// payload
	Payload *models.HReport `json:"Payload,omitempty"`

	// status
	Status string `json:"Status,omitempty"`
}

// Validate validates this health o k body
func (o *HealthOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *HealthOKBody) validatePayload(formats strfmt.Registry) error {

	if swag.IsZero(o.Payload) { // not required
		return nil
	}

	if o.Payload != nil {
		if err := o.Payload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("healthOK" + "." + "Payload")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *HealthOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *HealthOKBody) UnmarshalBinary(b []byte) error {
	var res HealthOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
