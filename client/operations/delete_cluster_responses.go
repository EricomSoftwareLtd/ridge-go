// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteClusterReader is a Reader for the DeleteCluster structure.
type DeleteClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteClusterNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteClusterNoContent creates a DeleteClusterNoContent with default headers values
func NewDeleteClusterNoContent() *DeleteClusterNoContent {
	return &DeleteClusterNoContent{}
}

/*DeleteClusterNoContent handles this case with default header values.

OK
*/
type DeleteClusterNoContent struct {
}

func (o *DeleteClusterNoContent) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/projects/{project}/clusters/{cluster}][%d] deleteClusterNoContent ", 204)
}

func (o *DeleteClusterNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterBadRequest creates a DeleteClusterBadRequest with default headers values
func NewDeleteClusterBadRequest() *DeleteClusterBadRequest {
	return &DeleteClusterBadRequest{}
}

/*DeleteClusterBadRequest handles this case with default header values.

Invalid request parameters
*/
type DeleteClusterBadRequest struct {
	Payload string
}

func (o *DeleteClusterBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/projects/{project}/clusters/{cluster}][%d] deleteClusterBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteClusterBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterForbidden creates a DeleteClusterForbidden with default headers values
func NewDeleteClusterForbidden() *DeleteClusterForbidden {
	return &DeleteClusterForbidden{}
}

/*DeleteClusterForbidden handles this case with default header values.

Forbidden
*/
type DeleteClusterForbidden struct {
	Payload string
}

func (o *DeleteClusterForbidden) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/projects/{project}/clusters/{cluster}][%d] deleteClusterForbidden  %+v", 403, o.Payload)
}

func (o *DeleteClusterForbidden) GetPayload() string {
	return o.Payload
}

func (o *DeleteClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterNotFound creates a DeleteClusterNotFound with default headers values
func NewDeleteClusterNotFound() *DeleteClusterNotFound {
	return &DeleteClusterNotFound{}
}

/*DeleteClusterNotFound handles this case with default header values.

Resource not found
*/
type DeleteClusterNotFound struct {
	Payload string
}

func (o *DeleteClusterNotFound) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/projects/{project}/clusters/{cluster}][%d] deleteClusterNotFound  %+v", 404, o.Payload)
}

func (o *DeleteClusterNotFound) GetPayload() string {
	return o.Payload
}

func (o *DeleteClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
