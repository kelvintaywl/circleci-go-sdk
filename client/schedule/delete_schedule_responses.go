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

// DeleteScheduleReader is a Reader for the DeleteSchedule structure.
type DeleteScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteScheduleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteScheduleOK creates a DeleteScheduleOK with default headers values
func NewDeleteScheduleOK() *DeleteScheduleOK {
	return &DeleteScheduleOK{}
}

/*
DeleteScheduleOK describes a response with status code 200, with default header values.

DeleteScheduleOK delete schedule o k
*/
type DeleteScheduleOK struct {
	Payload *models.Deleted
}

// IsSuccess returns true when this delete schedule o k response has a 2xx status code
func (o *DeleteScheduleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete schedule o k response has a 3xx status code
func (o *DeleteScheduleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete schedule o k response has a 4xx status code
func (o *DeleteScheduleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete schedule o k response has a 5xx status code
func (o *DeleteScheduleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete schedule o k response a status code equal to that given
func (o *DeleteScheduleOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteScheduleOK) Error() string {
	return fmt.Sprintf("[DELETE /schedule/{id}][%d] deleteScheduleOK  %+v", 200, o.Payload)
}

func (o *DeleteScheduleOK) String() string {
	return fmt.Sprintf("[DELETE /schedule/{id}][%d] deleteScheduleOK  %+v", 200, o.Payload)
}

func (o *DeleteScheduleOK) GetPayload() *models.Deleted {
	return o.Payload
}

func (o *DeleteScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Deleted)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScheduleBadRequest creates a DeleteScheduleBadRequest with default headers values
func NewDeleteScheduleBadRequest() *DeleteScheduleBadRequest {
	return &DeleteScheduleBadRequest{}
}

/*
DeleteScheduleBadRequest describes a response with status code 400, with default header values.

Invalid input
*/
type DeleteScheduleBadRequest struct {
	Payload *models.Errored
}

// IsSuccess returns true when this delete schedule bad request response has a 2xx status code
func (o *DeleteScheduleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete schedule bad request response has a 3xx status code
func (o *DeleteScheduleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete schedule bad request response has a 4xx status code
func (o *DeleteScheduleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete schedule bad request response has a 5xx status code
func (o *DeleteScheduleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete schedule bad request response a status code equal to that given
func (o *DeleteScheduleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteScheduleBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /schedule/{id}][%d] deleteScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteScheduleBadRequest) String() string {
	return fmt.Sprintf("[DELETE /schedule/{id}][%d] deleteScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteScheduleBadRequest) GetPayload() *models.Errored {
	return o.Payload
}

func (o *DeleteScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errored)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScheduleNotFound creates a DeleteScheduleNotFound with default headers values
func NewDeleteScheduleNotFound() *DeleteScheduleNotFound {
	return &DeleteScheduleNotFound{}
}

/*
DeleteScheduleNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteScheduleNotFound struct {
	Payload *models.Errored
}

// IsSuccess returns true when this delete schedule not found response has a 2xx status code
func (o *DeleteScheduleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete schedule not found response has a 3xx status code
func (o *DeleteScheduleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete schedule not found response has a 4xx status code
func (o *DeleteScheduleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete schedule not found response has a 5xx status code
func (o *DeleteScheduleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete schedule not found response a status code equal to that given
func (o *DeleteScheduleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteScheduleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /schedule/{id}][%d] deleteScheduleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteScheduleNotFound) String() string {
	return fmt.Sprintf("[DELETE /schedule/{id}][%d] deleteScheduleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteScheduleNotFound) GetPayload() *models.Errored {
	return o.Payload
}

func (o *DeleteScheduleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errored)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
