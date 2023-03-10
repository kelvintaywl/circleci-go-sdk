// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kelvintaywl/circleci-go-sdk/models"
)

// DeleteWebhookReader is a Reader for the DeleteWebhook structure.
type DeleteWebhookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWebhookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteWebhookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteWebhookBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteWebhookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteWebhookOK creates a DeleteWebhookOK with default headers values
func NewDeleteWebhookOK() *DeleteWebhookOK {
	return &DeleteWebhookOK{}
}

/*
DeleteWebhookOK describes a response with status code 200, with default header values.

DeleteWebhookOK delete webhook o k
*/
type DeleteWebhookOK struct {
	Payload *models.Deleted
}

// IsSuccess returns true when this delete webhook o k response has a 2xx status code
func (o *DeleteWebhookOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete webhook o k response has a 3xx status code
func (o *DeleteWebhookOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete webhook o k response has a 4xx status code
func (o *DeleteWebhookOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete webhook o k response has a 5xx status code
func (o *DeleteWebhookOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete webhook o k response a status code equal to that given
func (o *DeleteWebhookOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteWebhookOK) Error() string {
	return fmt.Sprintf("[DELETE /webhook/{id}][%d] deleteWebhookOK  %+v", 200, o.Payload)
}

func (o *DeleteWebhookOK) String() string {
	return fmt.Sprintf("[DELETE /webhook/{id}][%d] deleteWebhookOK  %+v", 200, o.Payload)
}

func (o *DeleteWebhookOK) GetPayload() *models.Deleted {
	return o.Payload
}

func (o *DeleteWebhookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Deleted)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebhookBadRequest creates a DeleteWebhookBadRequest with default headers values
func NewDeleteWebhookBadRequest() *DeleteWebhookBadRequest {
	return &DeleteWebhookBadRequest{}
}

/*
DeleteWebhookBadRequest describes a response with status code 400, with default header values.

Invalid input
*/
type DeleteWebhookBadRequest struct {
	Payload *models.Errored
}

// IsSuccess returns true when this delete webhook bad request response has a 2xx status code
func (o *DeleteWebhookBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete webhook bad request response has a 3xx status code
func (o *DeleteWebhookBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete webhook bad request response has a 4xx status code
func (o *DeleteWebhookBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete webhook bad request response has a 5xx status code
func (o *DeleteWebhookBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete webhook bad request response a status code equal to that given
func (o *DeleteWebhookBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteWebhookBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /webhook/{id}][%d] deleteWebhookBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWebhookBadRequest) String() string {
	return fmt.Sprintf("[DELETE /webhook/{id}][%d] deleteWebhookBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWebhookBadRequest) GetPayload() *models.Errored {
	return o.Payload
}

func (o *DeleteWebhookBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errored)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebhookNotFound creates a DeleteWebhookNotFound with default headers values
func NewDeleteWebhookNotFound() *DeleteWebhookNotFound {
	return &DeleteWebhookNotFound{}
}

/*
DeleteWebhookNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteWebhookNotFound struct {
	Payload *models.Errored
}

// IsSuccess returns true when this delete webhook not found response has a 2xx status code
func (o *DeleteWebhookNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete webhook not found response has a 3xx status code
func (o *DeleteWebhookNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete webhook not found response has a 4xx status code
func (o *DeleteWebhookNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete webhook not found response has a 5xx status code
func (o *DeleteWebhookNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete webhook not found response a status code equal to that given
func (o *DeleteWebhookNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteWebhookNotFound) Error() string {
	return fmt.Sprintf("[DELETE /webhook/{id}][%d] deleteWebhookNotFound  %+v", 404, o.Payload)
}

func (o *DeleteWebhookNotFound) String() string {
	return fmt.Sprintf("[DELETE /webhook/{id}][%d] deleteWebhookNotFound  %+v", 404, o.Payload)
}

func (o *DeleteWebhookNotFound) GetPayload() *models.Errored {
	return o.Payload
}

func (o *DeleteWebhookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errored)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
