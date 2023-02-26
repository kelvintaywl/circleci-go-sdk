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

// GetContextReader is a Reader for the GetContext structure.
type GetContextReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContextReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContextOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetContextBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetContextOK creates a GetContextOK with default headers values
func NewGetContextOK() *GetContextOK {
	return &GetContextOK{}
}

/*
GetContextOK describes a response with status code 200, with default header values.

A context object
*/
type GetContextOK struct {
	Payload *models.ContextInfo
}

// IsSuccess returns true when this get context o k response has a 2xx status code
func (o *GetContextOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get context o k response has a 3xx status code
func (o *GetContextOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get context o k response has a 4xx status code
func (o *GetContextOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get context o k response has a 5xx status code
func (o *GetContextOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get context o k response a status code equal to that given
func (o *GetContextOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetContextOK) Error() string {
	return fmt.Sprintf("[GET /context/{id}][%d] getContextOK  %+v", 200, o.Payload)
}

func (o *GetContextOK) String() string {
	return fmt.Sprintf("[GET /context/{id}][%d] getContextOK  %+v", 200, o.Payload)
}

func (o *GetContextOK) GetPayload() *models.ContextInfo {
	return o.Payload
}

func (o *GetContextOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContextInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContextBadRequest creates a GetContextBadRequest with default headers values
func NewGetContextBadRequest() *GetContextBadRequest {
	return &GetContextBadRequest{}
}

/*
GetContextBadRequest describes a response with status code 400, with default header values.

Invalid input
*/
type GetContextBadRequest struct {
	Payload *models.Errored
}

// IsSuccess returns true when this get context bad request response has a 2xx status code
func (o *GetContextBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get context bad request response has a 3xx status code
func (o *GetContextBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get context bad request response has a 4xx status code
func (o *GetContextBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get context bad request response has a 5xx status code
func (o *GetContextBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get context bad request response a status code equal to that given
func (o *GetContextBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetContextBadRequest) Error() string {
	return fmt.Sprintf("[GET /context/{id}][%d] getContextBadRequest  %+v", 400, o.Payload)
}

func (o *GetContextBadRequest) String() string {
	return fmt.Sprintf("[GET /context/{id}][%d] getContextBadRequest  %+v", 400, o.Payload)
}

func (o *GetContextBadRequest) GetPayload() *models.Errored {
	return o.Payload
}

func (o *GetContextBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errored)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}