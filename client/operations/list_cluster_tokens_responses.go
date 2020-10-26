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

// ListClusterTokensReader is a Reader for the ListClusterTokens structure.
type ListClusterTokensReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClusterTokensReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListClusterTokensOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListClusterTokensBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListClusterTokensForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListClusterTokensNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListClusterTokensOK creates a ListClusterTokensOK with default headers values
func NewListClusterTokensOK() *ListClusterTokensOK {
	return &ListClusterTokensOK{}
}

/*ListClusterTokensOK handles this case with default header values.

Success
*/
type ListClusterTokensOK struct {
	Payload *models.Tokens
}

func (o *ListClusterTokensOK) Error() string {
	return fmt.Sprintf("[GET /orgs/{org}/projects/{project}/clusters/{cluster}/tokens][%d] listClusterTokensOK  %+v", 200, o.Payload)
}

func (o *ListClusterTokensOK) GetPayload() *models.Tokens {
	return o.Payload
}

func (o *ListClusterTokensOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tokens)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClusterTokensBadRequest creates a ListClusterTokensBadRequest with default headers values
func NewListClusterTokensBadRequest() *ListClusterTokensBadRequest {
	return &ListClusterTokensBadRequest{}
}

/*ListClusterTokensBadRequest handles this case with default header values.

Invalid request parameters
*/
type ListClusterTokensBadRequest struct {
	Payload string
}

func (o *ListClusterTokensBadRequest) Error() string {
	return fmt.Sprintf("[GET /orgs/{org}/projects/{project}/clusters/{cluster}/tokens][%d] listClusterTokensBadRequest  %+v", 400, o.Payload)
}

func (o *ListClusterTokensBadRequest) GetPayload() string {
	return o.Payload
}

func (o *ListClusterTokensBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClusterTokensForbidden creates a ListClusterTokensForbidden with default headers values
func NewListClusterTokensForbidden() *ListClusterTokensForbidden {
	return &ListClusterTokensForbidden{}
}

/*ListClusterTokensForbidden handles this case with default header values.

Forbidden
*/
type ListClusterTokensForbidden struct {
	Payload string
}

func (o *ListClusterTokensForbidden) Error() string {
	return fmt.Sprintf("[GET /orgs/{org}/projects/{project}/clusters/{cluster}/tokens][%d] listClusterTokensForbidden  %+v", 403, o.Payload)
}

func (o *ListClusterTokensForbidden) GetPayload() string {
	return o.Payload
}

func (o *ListClusterTokensForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClusterTokensNotFound creates a ListClusterTokensNotFound with default headers values
func NewListClusterTokensNotFound() *ListClusterTokensNotFound {
	return &ListClusterTokensNotFound{}
}

/*ListClusterTokensNotFound handles this case with default header values.

Resource not found
*/
type ListClusterTokensNotFound struct {
	Payload string
}

func (o *ListClusterTokensNotFound) Error() string {
	return fmt.Sprintf("[GET /orgs/{org}/projects/{project}/clusters/{cluster}/tokens][%d] listClusterTokensNotFound  %+v", 404, o.Payload)
}

func (o *ListClusterTokensNotFound) GetPayload() string {
	return o.Payload
}

func (o *ListClusterTokensNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
