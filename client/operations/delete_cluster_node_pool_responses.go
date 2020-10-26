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

// DeleteClusterNodePoolReader is a Reader for the DeleteClusterNodePool structure.
type DeleteClusterNodePoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterNodePoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteClusterNodePoolNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteClusterNodePoolBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteClusterNodePoolForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteClusterNodePoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteClusterNodePoolNoContent creates a DeleteClusterNodePoolNoContent with default headers values
func NewDeleteClusterNodePoolNoContent() *DeleteClusterNodePoolNoContent {
	return &DeleteClusterNodePoolNoContent{}
}

/*DeleteClusterNodePoolNoContent handles this case with default header values.

OK
*/
type DeleteClusterNodePoolNoContent struct {
}

func (o *DeleteClusterNodePoolNoContent) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/projects/{project}/clusters/{cluster}/node-pools/{node_pool}][%d] deleteClusterNodePoolNoContent ", 204)
}

func (o *DeleteClusterNodePoolNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterNodePoolBadRequest creates a DeleteClusterNodePoolBadRequest with default headers values
func NewDeleteClusterNodePoolBadRequest() *DeleteClusterNodePoolBadRequest {
	return &DeleteClusterNodePoolBadRequest{}
}

/*DeleteClusterNodePoolBadRequest handles this case with default header values.

Invalid request parameters
*/
type DeleteClusterNodePoolBadRequest struct {
	Payload string
}

func (o *DeleteClusterNodePoolBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/projects/{project}/clusters/{cluster}/node-pools/{node_pool}][%d] deleteClusterNodePoolBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteClusterNodePoolBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteClusterNodePoolBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterNodePoolForbidden creates a DeleteClusterNodePoolForbidden with default headers values
func NewDeleteClusterNodePoolForbidden() *DeleteClusterNodePoolForbidden {
	return &DeleteClusterNodePoolForbidden{}
}

/*DeleteClusterNodePoolForbidden handles this case with default header values.

Forbidden
*/
type DeleteClusterNodePoolForbidden struct {
	Payload string
}

func (o *DeleteClusterNodePoolForbidden) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/projects/{project}/clusters/{cluster}/node-pools/{node_pool}][%d] deleteClusterNodePoolForbidden  %+v", 403, o.Payload)
}

func (o *DeleteClusterNodePoolForbidden) GetPayload() string {
	return o.Payload
}

func (o *DeleteClusterNodePoolForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterNodePoolNotFound creates a DeleteClusterNodePoolNotFound with default headers values
func NewDeleteClusterNodePoolNotFound() *DeleteClusterNodePoolNotFound {
	return &DeleteClusterNodePoolNotFound{}
}

/*DeleteClusterNodePoolNotFound handles this case with default header values.

Resource not found
*/
type DeleteClusterNodePoolNotFound struct {
	Payload string
}

func (o *DeleteClusterNodePoolNotFound) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/projects/{project}/clusters/{cluster}/node-pools/{node_pool}][%d] deleteClusterNodePoolNotFound  %+v", 404, o.Payload)
}

func (o *DeleteClusterNodePoolNotFound) GetPayload() string {
	return o.Payload
}

func (o *DeleteClusterNodePoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
