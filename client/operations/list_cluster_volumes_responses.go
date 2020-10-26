// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/EricomSoftwareLtd/ridge-go/models"
)

// ListClusterVolumesReader is a Reader for the ListClusterVolumes structure.
type ListClusterVolumesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClusterVolumesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListClusterVolumesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListClusterVolumesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListClusterVolumesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListClusterVolumesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListClusterVolumesOK creates a ListClusterVolumesOK with default headers values
func NewListClusterVolumesOK() *ListClusterVolumesOK {
	return &ListClusterVolumesOK{}
}

/*ListClusterVolumesOK handles this case with default header values.

Success
*/
type ListClusterVolumesOK struct {
	Payload *models.Volumes
}

func (o *ListClusterVolumesOK) Error() string {
	return fmt.Sprintf("[GET /orgs/{org}/projects/{project}/clusters/{cluster}/volumes][%d] listClusterVolumesOK  %+v", 200, o.Payload)
}

func (o *ListClusterVolumesOK) GetPayload() *models.Volumes {
	return o.Payload
}

func (o *ListClusterVolumesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Volumes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClusterVolumesBadRequest creates a ListClusterVolumesBadRequest with default headers values
func NewListClusterVolumesBadRequest() *ListClusterVolumesBadRequest {
	return &ListClusterVolumesBadRequest{}
}

/*ListClusterVolumesBadRequest handles this case with default header values.

Invalid request parameters
*/
type ListClusterVolumesBadRequest struct {
	Payload string
}

func (o *ListClusterVolumesBadRequest) Error() string {
	return fmt.Sprintf("[GET /orgs/{org}/projects/{project}/clusters/{cluster}/volumes][%d] listClusterVolumesBadRequest  %+v", 400, o.Payload)
}

func (o *ListClusterVolumesBadRequest) GetPayload() string {
	return o.Payload
}

func (o *ListClusterVolumesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClusterVolumesForbidden creates a ListClusterVolumesForbidden with default headers values
func NewListClusterVolumesForbidden() *ListClusterVolumesForbidden {
	return &ListClusterVolumesForbidden{}
}

/*ListClusterVolumesForbidden handles this case with default header values.

Forbidden
*/
type ListClusterVolumesForbidden struct {
	Payload string
}

func (o *ListClusterVolumesForbidden) Error() string {
	return fmt.Sprintf("[GET /orgs/{org}/projects/{project}/clusters/{cluster}/volumes][%d] listClusterVolumesForbidden  %+v", 403, o.Payload)
}

func (o *ListClusterVolumesForbidden) GetPayload() string {
	return o.Payload
}

func (o *ListClusterVolumesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClusterVolumesNotFound creates a ListClusterVolumesNotFound with default headers values
func NewListClusterVolumesNotFound() *ListClusterVolumesNotFound {
	return &ListClusterVolumesNotFound{}
}

/*ListClusterVolumesNotFound handles this case with default header values.

Resource not found
*/
type ListClusterVolumesNotFound struct {
	Payload string
}

func (o *ListClusterVolumesNotFound) Error() string {
	return fmt.Sprintf("[GET /orgs/{org}/projects/{project}/clusters/{cluster}/volumes][%d] listClusterVolumesNotFound  %+v", 404, o.Payload)
}

func (o *ListClusterVolumesNotFound) GetPayload() string {
	return o.Payload
}

func (o *ListClusterVolumesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
