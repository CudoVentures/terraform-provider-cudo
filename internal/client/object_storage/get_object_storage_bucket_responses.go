// Code generated by go-swagger; DO NOT EDIT.

package object_storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// GetObjectStorageBucketReader is a Reader for the GetObjectStorageBucket structure.
type GetObjectStorageBucketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetObjectStorageBucketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetObjectStorageBucketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetObjectStorageBucketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetObjectStorageBucketOK creates a GetObjectStorageBucketOK with default headers values
func NewGetObjectStorageBucketOK() *GetObjectStorageBucketOK {
	return &GetObjectStorageBucketOK{}
}

/*
GetObjectStorageBucketOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetObjectStorageBucketOK struct {
	Payload *models.ObjectStorageBucket
}

// IsSuccess returns true when this get object storage bucket o k response has a 2xx status code
func (o *GetObjectStorageBucketOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get object storage bucket o k response has a 3xx status code
func (o *GetObjectStorageBucketOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get object storage bucket o k response has a 4xx status code
func (o *GetObjectStorageBucketOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get object storage bucket o k response has a 5xx status code
func (o *GetObjectStorageBucketOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get object storage bucket o k response a status code equal to that given
func (o *GetObjectStorageBucketOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get object storage bucket o k response
func (o *GetObjectStorageBucketOK) Code() int {
	return 200
}

func (o *GetObjectStorageBucketOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/object-storage/buckets/{id}][%d] getObjectStorageBucketOK  %+v", 200, o.Payload)
}

func (o *GetObjectStorageBucketOK) String() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/object-storage/buckets/{id}][%d] getObjectStorageBucketOK  %+v", 200, o.Payload)
}

func (o *GetObjectStorageBucketOK) GetPayload() *models.ObjectStorageBucket {
	return o.Payload
}

func (o *GetObjectStorageBucketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectStorageBucket)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetObjectStorageBucketDefault creates a GetObjectStorageBucketDefault with default headers values
func NewGetObjectStorageBucketDefault(code int) *GetObjectStorageBucketDefault {
	return &GetObjectStorageBucketDefault{
		_statusCode: code,
	}
}

/*
GetObjectStorageBucketDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetObjectStorageBucketDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this get object storage bucket default response has a 2xx status code
func (o *GetObjectStorageBucketDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get object storage bucket default response has a 3xx status code
func (o *GetObjectStorageBucketDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get object storage bucket default response has a 4xx status code
func (o *GetObjectStorageBucketDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get object storage bucket default response has a 5xx status code
func (o *GetObjectStorageBucketDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get object storage bucket default response a status code equal to that given
func (o *GetObjectStorageBucketDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get object storage bucket default response
func (o *GetObjectStorageBucketDefault) Code() int {
	return o._statusCode
}

func (o *GetObjectStorageBucketDefault) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/object-storage/buckets/{id}][%d] GetObjectStorageBucket default  %+v", o._statusCode, o.Payload)
}

func (o *GetObjectStorageBucketDefault) String() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/object-storage/buckets/{id}][%d] GetObjectStorageBucket default  %+v", o._statusCode, o.Payload)
}

func (o *GetObjectStorageBucketDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *GetObjectStorageBucketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
