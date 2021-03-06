// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/EricomSoftwareLtd/ridge-go/models"
)

// NewCreateClusterParams creates a new CreateClusterParams object
// with the default values initialized.
func NewCreateClusterParams() *CreateClusterParams {
	var ()
	return &CreateClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClusterParamsWithTimeout creates a new CreateClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateClusterParamsWithTimeout(timeout time.Duration) *CreateClusterParams {
	var ()
	return &CreateClusterParams{

		timeout: timeout,
	}
}

// NewCreateClusterParamsWithContext creates a new CreateClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateClusterParamsWithContext(ctx context.Context) *CreateClusterParams {
	var ()
	return &CreateClusterParams{

		Context: ctx,
	}
}

// NewCreateClusterParamsWithHTTPClient creates a new CreateClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateClusterParamsWithHTTPClient(client *http.Client) *CreateClusterParams {
	var ()
	return &CreateClusterParams{
		HTTPClient: client,
	}
}

/*CreateClusterParams contains all the parameters to send to the API endpoint
for the create cluster operation typically these are written to a http.Request
*/
type CreateClusterParams struct {

	/*Body*/
	Body *models.ClusterCreate
	/*Org
	  The organization identifier

	*/
	Org string
	/*Project
	  The project identifier

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create cluster params
func (o *CreateClusterParams) WithTimeout(timeout time.Duration) *CreateClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cluster params
func (o *CreateClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cluster params
func (o *CreateClusterParams) WithContext(ctx context.Context) *CreateClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cluster params
func (o *CreateClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cluster params
func (o *CreateClusterParams) WithHTTPClient(client *http.Client) *CreateClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cluster params
func (o *CreateClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create cluster params
func (o *CreateClusterParams) WithBody(body *models.ClusterCreate) *CreateClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create cluster params
func (o *CreateClusterParams) SetBody(body *models.ClusterCreate) {
	o.Body = body
}

// WithOrg adds the org to the create cluster params
func (o *CreateClusterParams) WithOrg(org string) *CreateClusterParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the create cluster params
func (o *CreateClusterParams) SetOrg(org string) {
	o.Org = org
}

// WithProject adds the project to the create cluster params
func (o *CreateClusterParams) WithProject(project string) *CreateClusterParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the create cluster params
func (o *CreateClusterParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param org
	if err := r.SetPathParam("org", o.Org); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
