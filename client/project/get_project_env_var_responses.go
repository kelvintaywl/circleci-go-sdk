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

// GetProjectEnvVarReader is a Reader for the GetProjectEnvVar structure.
type GetProjectEnvVarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectEnvVarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectEnvVarOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectEnvVarBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectEnvVarOK creates a GetProjectEnvVarOK with default headers values
func NewGetProjectEnvVarOK() *GetProjectEnvVarOK {
	return &GetProjectEnvVarOK{}
}

/*
GetProjectEnvVarOK describes a response with status code 200, with default header values.

A project environment variable.
*/
type GetProjectEnvVarOK struct {
	Payload *models.ProjectEnvVarInfo
}

// IsSuccess returns true when this get project env var o k response has a 2xx status code
func (o *GetProjectEnvVarOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get project env var o k response has a 3xx status code
func (o *GetProjectEnvVarOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project env var o k response has a 4xx status code
func (o *GetProjectEnvVarOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project env var o k response has a 5xx status code
func (o *GetProjectEnvVarOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get project env var o k response a status code equal to that given
func (o *GetProjectEnvVarOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetProjectEnvVarOK) Error() string {
	return fmt.Sprintf("[GET /project/{project-slug}/envvar/{name}][%d] getProjectEnvVarOK  %+v", 200, o.Payload)
}

func (o *GetProjectEnvVarOK) String() string {
	return fmt.Sprintf("[GET /project/{project-slug}/envvar/{name}][%d] getProjectEnvVarOK  %+v", 200, o.Payload)
}

func (o *GetProjectEnvVarOK) GetPayload() *models.ProjectEnvVarInfo {
	return o.Payload
}

func (o *GetProjectEnvVarOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectEnvVarInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectEnvVarBadRequest creates a GetProjectEnvVarBadRequest with default headers values
func NewGetProjectEnvVarBadRequest() *GetProjectEnvVarBadRequest {
	return &GetProjectEnvVarBadRequest{}
}

/*
GetProjectEnvVarBadRequest describes a response with status code 400, with default header values.

Invalid input
*/
type GetProjectEnvVarBadRequest struct {
	Payload *models.Errored
}

// IsSuccess returns true when this get project env var bad request response has a 2xx status code
func (o *GetProjectEnvVarBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project env var bad request response has a 3xx status code
func (o *GetProjectEnvVarBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project env var bad request response has a 4xx status code
func (o *GetProjectEnvVarBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project env var bad request response has a 5xx status code
func (o *GetProjectEnvVarBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get project env var bad request response a status code equal to that given
func (o *GetProjectEnvVarBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetProjectEnvVarBadRequest) Error() string {
	return fmt.Sprintf("[GET /project/{project-slug}/envvar/{name}][%d] getProjectEnvVarBadRequest  %+v", 400, o.Payload)
}

func (o *GetProjectEnvVarBadRequest) String() string {
	return fmt.Sprintf("[GET /project/{project-slug}/envvar/{name}][%d] getProjectEnvVarBadRequest  %+v", 400, o.Payload)
}

func (o *GetProjectEnvVarBadRequest) GetPayload() *models.Errored {
	return o.Payload
}

func (o *GetProjectEnvVarBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errored)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
