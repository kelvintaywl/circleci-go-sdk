// Code generated by go-swagger; DO NOT EDIT.

package contexts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kelvintaywl/circleci-go-sdk/models"
)

// AddContextReader is a Reader for the AddContext structure.
type AddContextReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddContextReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddContextOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddContextBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddContextOK creates a AddContextOK with default headers values
func NewAddContextOK() *AddContextOK {
	return &AddContextOK{}
}

/*
AddContextOK describes a response with status code 200, with default header values.

Successfully added
*/
type AddContextOK struct {
	Payload *models.ContextInfo
}

// IsSuccess returns true when this add context o k response has a 2xx status code
func (o *AddContextOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add context o k response has a 3xx status code
func (o *AddContextOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add context o k response has a 4xx status code
func (o *AddContextOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add context o k response has a 5xx status code
func (o *AddContextOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add context o k response a status code equal to that given
func (o *AddContextOK) IsCode(code int) bool {
	return code == 200
}

func (o *AddContextOK) Error() string {
	return fmt.Sprintf("[POST /context][%d] addContextOK  %+v", 200, o.Payload)
}

func (o *AddContextOK) String() string {
	return fmt.Sprintf("[POST /context][%d] addContextOK  %+v", 200, o.Payload)
}

func (o *AddContextOK) GetPayload() *models.ContextInfo {
	return o.Payload
}

func (o *AddContextOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContextInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddContextBadRequest creates a AddContextBadRequest with default headers values
func NewAddContextBadRequest() *AddContextBadRequest {
	return &AddContextBadRequest{}
}

/*
AddContextBadRequest describes a response with status code 400, with default header values.

Invalid input
*/
type AddContextBadRequest struct {
	Payload *models.Errored
}

// IsSuccess returns true when this add context bad request response has a 2xx status code
func (o *AddContextBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add context bad request response has a 3xx status code
func (o *AddContextBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add context bad request response has a 4xx status code
func (o *AddContextBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add context bad request response has a 5xx status code
func (o *AddContextBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add context bad request response a status code equal to that given
func (o *AddContextBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AddContextBadRequest) Error() string {
	return fmt.Sprintf("[POST /context][%d] addContextBadRequest  %+v", 400, o.Payload)
}

func (o *AddContextBadRequest) String() string {
	return fmt.Sprintf("[POST /context][%d] addContextBadRequest  %+v", 400, o.Payload)
}

func (o *AddContextBadRequest) GetPayload() *models.Errored {
	return o.Payload
}

func (o *AddContextBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errored)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
