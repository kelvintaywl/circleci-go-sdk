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

// DeleteContextEnvVarReader is a Reader for the DeleteContextEnvVar structure.
type DeleteContextEnvVarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteContextEnvVarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteContextEnvVarOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteContextEnvVarBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteContextEnvVarOK creates a DeleteContextEnvVarOK with default headers values
func NewDeleteContextEnvVarOK() *DeleteContextEnvVarOK {
	return &DeleteContextEnvVarOK{}
}

/*
DeleteContextEnvVarOK describes a response with status code 200, with default header values.

DeleteContextEnvVarOK delete context env var o k
*/
type DeleteContextEnvVarOK struct {
	Payload *models.Deleted
}

// IsSuccess returns true when this delete context env var o k response has a 2xx status code
func (o *DeleteContextEnvVarOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete context env var o k response has a 3xx status code
func (o *DeleteContextEnvVarOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete context env var o k response has a 4xx status code
func (o *DeleteContextEnvVarOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete context env var o k response has a 5xx status code
func (o *DeleteContextEnvVarOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete context env var o k response a status code equal to that given
func (o *DeleteContextEnvVarOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteContextEnvVarOK) Error() string {
	return fmt.Sprintf("[DELETE /context/{id}/environment-variable/{name}][%d] deleteContextEnvVarOK  %+v", 200, o.Payload)
}

func (o *DeleteContextEnvVarOK) String() string {
	return fmt.Sprintf("[DELETE /context/{id}/environment-variable/{name}][%d] deleteContextEnvVarOK  %+v", 200, o.Payload)
}

func (o *DeleteContextEnvVarOK) GetPayload() *models.Deleted {
	return o.Payload
}

func (o *DeleteContextEnvVarOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Deleted)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContextEnvVarBadRequest creates a DeleteContextEnvVarBadRequest with default headers values
func NewDeleteContextEnvVarBadRequest() *DeleteContextEnvVarBadRequest {
	return &DeleteContextEnvVarBadRequest{}
}

/*
DeleteContextEnvVarBadRequest describes a response with status code 400, with default header values.

Invalid input
*/
type DeleteContextEnvVarBadRequest struct {
	Payload *models.Errored
}

// IsSuccess returns true when this delete context env var bad request response has a 2xx status code
func (o *DeleteContextEnvVarBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete context env var bad request response has a 3xx status code
func (o *DeleteContextEnvVarBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete context env var bad request response has a 4xx status code
func (o *DeleteContextEnvVarBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete context env var bad request response has a 5xx status code
func (o *DeleteContextEnvVarBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete context env var bad request response a status code equal to that given
func (o *DeleteContextEnvVarBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteContextEnvVarBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /context/{id}/environment-variable/{name}][%d] deleteContextEnvVarBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteContextEnvVarBadRequest) String() string {
	return fmt.Sprintf("[DELETE /context/{id}/environment-variable/{name}][%d] deleteContextEnvVarBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteContextEnvVarBadRequest) GetPayload() *models.Errored {
	return o.Payload
}

func (o *DeleteContextEnvVarBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errored)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
