// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kelvintaywl/circleci-go-sdk/models"
)

// AddProjectCheckoutKeyReader is a Reader for the AddProjectCheckoutKey structure.
type AddProjectCheckoutKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddProjectCheckoutKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddProjectCheckoutKeyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddProjectCheckoutKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddProjectCheckoutKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddProjectCheckoutKeyCreated creates a AddProjectCheckoutKeyCreated with default headers values
func NewAddProjectCheckoutKeyCreated() *AddProjectCheckoutKeyCreated {
	return &AddProjectCheckoutKeyCreated{}
}

/*
AddProjectCheckoutKeyCreated describes a response with status code 201, with default header values.

Successfully added
*/
type AddProjectCheckoutKeyCreated struct {
	Payload *models.ProjectCheckoutKeyInfo
}

// IsSuccess returns true when this add project checkout key created response has a 2xx status code
func (o *AddProjectCheckoutKeyCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add project checkout key created response has a 3xx status code
func (o *AddProjectCheckoutKeyCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add project checkout key created response has a 4xx status code
func (o *AddProjectCheckoutKeyCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add project checkout key created response has a 5xx status code
func (o *AddProjectCheckoutKeyCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add project checkout key created response a status code equal to that given
func (o *AddProjectCheckoutKeyCreated) IsCode(code int) bool {
	return code == 201
}

func (o *AddProjectCheckoutKeyCreated) Error() string {
	return fmt.Sprintf("[POST /project/{project-slug}/checkout-key][%d] addProjectCheckoutKeyCreated  %+v", 201, o.Payload)
}

func (o *AddProjectCheckoutKeyCreated) String() string {
	return fmt.Sprintf("[POST /project/{project-slug}/checkout-key][%d] addProjectCheckoutKeyCreated  %+v", 201, o.Payload)
}

func (o *AddProjectCheckoutKeyCreated) GetPayload() *models.ProjectCheckoutKeyInfo {
	return o.Payload
}

func (o *AddProjectCheckoutKeyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectCheckoutKeyInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProjectCheckoutKeyBadRequest creates a AddProjectCheckoutKeyBadRequest with default headers values
func NewAddProjectCheckoutKeyBadRequest() *AddProjectCheckoutKeyBadRequest {
	return &AddProjectCheckoutKeyBadRequest{}
}

/*
AddProjectCheckoutKeyBadRequest describes a response with status code 400, with default header values.

Invalid input
*/
type AddProjectCheckoutKeyBadRequest struct {
	Payload *models.Errored
}

// IsSuccess returns true when this add project checkout key bad request response has a 2xx status code
func (o *AddProjectCheckoutKeyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add project checkout key bad request response has a 3xx status code
func (o *AddProjectCheckoutKeyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add project checkout key bad request response has a 4xx status code
func (o *AddProjectCheckoutKeyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add project checkout key bad request response has a 5xx status code
func (o *AddProjectCheckoutKeyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add project checkout key bad request response a status code equal to that given
func (o *AddProjectCheckoutKeyBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AddProjectCheckoutKeyBadRequest) Error() string {
	return fmt.Sprintf("[POST /project/{project-slug}/checkout-key][%d] addProjectCheckoutKeyBadRequest  %+v", 400, o.Payload)
}

func (o *AddProjectCheckoutKeyBadRequest) String() string {
	return fmt.Sprintf("[POST /project/{project-slug}/checkout-key][%d] addProjectCheckoutKeyBadRequest  %+v", 400, o.Payload)
}

func (o *AddProjectCheckoutKeyBadRequest) GetPayload() *models.Errored {
	return o.Payload
}

func (o *AddProjectCheckoutKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errored)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProjectCheckoutKeyNotFound creates a AddProjectCheckoutKeyNotFound with default headers values
func NewAddProjectCheckoutKeyNotFound() *AddProjectCheckoutKeyNotFound {
	return &AddProjectCheckoutKeyNotFound{}
}

/*
AddProjectCheckoutKeyNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AddProjectCheckoutKeyNotFound struct {
	Payload *models.Errored
}

// IsSuccess returns true when this add project checkout key not found response has a 2xx status code
func (o *AddProjectCheckoutKeyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add project checkout key not found response has a 3xx status code
func (o *AddProjectCheckoutKeyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add project checkout key not found response has a 4xx status code
func (o *AddProjectCheckoutKeyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add project checkout key not found response has a 5xx status code
func (o *AddProjectCheckoutKeyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add project checkout key not found response a status code equal to that given
func (o *AddProjectCheckoutKeyNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AddProjectCheckoutKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /project/{project-slug}/checkout-key][%d] addProjectCheckoutKeyNotFound  %+v", 404, o.Payload)
}

func (o *AddProjectCheckoutKeyNotFound) String() string {
	return fmt.Sprintf("[POST /project/{project-slug}/checkout-key][%d] addProjectCheckoutKeyNotFound  %+v", 404, o.Payload)
}

func (o *AddProjectCheckoutKeyNotFound) GetPayload() *models.Errored {
	return o.Payload
}

func (o *AddProjectCheckoutKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errored)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
