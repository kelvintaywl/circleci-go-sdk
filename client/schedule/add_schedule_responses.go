// Code generated by go-swagger; DO NOT EDIT.

package schedule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kelvintaywl/circleci-go-sdk/models"
)

// AddScheduleReader is a Reader for the AddSchedule structure.
type AddScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddScheduleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddScheduleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddScheduleCreated creates a AddScheduleCreated with default headers values
func NewAddScheduleCreated() *AddScheduleCreated {
	return &AddScheduleCreated{}
}

/*
AddScheduleCreated describes a response with status code 201, with default header values.

Successfully added
*/
type AddScheduleCreated struct {
	Payload *models.ScheduleInfo
}

// IsSuccess returns true when this add schedule created response has a 2xx status code
func (o *AddScheduleCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add schedule created response has a 3xx status code
func (o *AddScheduleCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add schedule created response has a 4xx status code
func (o *AddScheduleCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add schedule created response has a 5xx status code
func (o *AddScheduleCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add schedule created response a status code equal to that given
func (o *AddScheduleCreated) IsCode(code int) bool {
	return code == 201
}

func (o *AddScheduleCreated) Error() string {
	return fmt.Sprintf("[POST /project/{project-slug}/schedule][%d] addScheduleCreated  %+v", 201, o.Payload)
}

func (o *AddScheduleCreated) String() string {
	return fmt.Sprintf("[POST /project/{project-slug}/schedule][%d] addScheduleCreated  %+v", 201, o.Payload)
}

func (o *AddScheduleCreated) GetPayload() *models.ScheduleInfo {
	return o.Payload
}

func (o *AddScheduleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScheduleInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddScheduleBadRequest creates a AddScheduleBadRequest with default headers values
func NewAddScheduleBadRequest() *AddScheduleBadRequest {
	return &AddScheduleBadRequest{}
}

/*
AddScheduleBadRequest describes a response with status code 400, with default header values.

Invalid input
*/
type AddScheduleBadRequest struct {
	Payload *models.Errored
}

// IsSuccess returns true when this add schedule bad request response has a 2xx status code
func (o *AddScheduleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add schedule bad request response has a 3xx status code
func (o *AddScheduleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add schedule bad request response has a 4xx status code
func (o *AddScheduleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add schedule bad request response has a 5xx status code
func (o *AddScheduleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add schedule bad request response a status code equal to that given
func (o *AddScheduleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AddScheduleBadRequest) Error() string {
	return fmt.Sprintf("[POST /project/{project-slug}/schedule][%d] addScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *AddScheduleBadRequest) String() string {
	return fmt.Sprintf("[POST /project/{project-slug}/schedule][%d] addScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *AddScheduleBadRequest) GetPayload() *models.Errored {
	return o.Payload
}

func (o *AddScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errored)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddScheduleNotFound creates a AddScheduleNotFound with default headers values
func NewAddScheduleNotFound() *AddScheduleNotFound {
	return &AddScheduleNotFound{}
}

/*
AddScheduleNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AddScheduleNotFound struct {
	Payload *models.Errored
}

// IsSuccess returns true when this add schedule not found response has a 2xx status code
func (o *AddScheduleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add schedule not found response has a 3xx status code
func (o *AddScheduleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add schedule not found response has a 4xx status code
func (o *AddScheduleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add schedule not found response has a 5xx status code
func (o *AddScheduleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add schedule not found response a status code equal to that given
func (o *AddScheduleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AddScheduleNotFound) Error() string {
	return fmt.Sprintf("[POST /project/{project-slug}/schedule][%d] addScheduleNotFound  %+v", 404, o.Payload)
}

func (o *AddScheduleNotFound) String() string {
	return fmt.Sprintf("[POST /project/{project-slug}/schedule][%d] addScheduleNotFound  %+v", 404, o.Payload)
}

func (o *AddScheduleNotFound) GetPayload() *models.Errored {
	return o.Payload
}

func (o *AddScheduleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errored)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
