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
)

// NewDescribeClusterParams creates a new DescribeClusterParams object
// with the default values initialized.
func NewDescribeClusterParams() *DescribeClusterParams {
	var ()
	return &DescribeClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeClusterParamsWithTimeout creates a new DescribeClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeClusterParamsWithTimeout(timeout time.Duration) *DescribeClusterParams {
	var ()
	return &DescribeClusterParams{

		timeout: timeout,
	}
}

// NewDescribeClusterParamsWithContext creates a new DescribeClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeClusterParamsWithContext(ctx context.Context) *DescribeClusterParams {
	var ()
	return &DescribeClusterParams{

		Context: ctx,
	}
}

// NewDescribeClusterParamsWithHTTPClient creates a new DescribeClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeClusterParamsWithHTTPClient(client *http.Client) *DescribeClusterParams {
	var ()
	return &DescribeClusterParams{
		HTTPClient: client,
	}
}

/*DescribeClusterParams contains all the parameters to send to the API endpoint
for the describe cluster operation typically these are written to a http.Request
*/
type DescribeClusterParams struct {

	/*Cluster
	  The identifier of a cluster

	*/
	Cluster string
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

// WithTimeout adds the timeout to the describe cluster params
func (o *DescribeClusterParams) WithTimeout(timeout time.Duration) *DescribeClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe cluster params
func (o *DescribeClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe cluster params
func (o *DescribeClusterParams) WithContext(ctx context.Context) *DescribeClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe cluster params
func (o *DescribeClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe cluster params
func (o *DescribeClusterParams) WithHTTPClient(client *http.Client) *DescribeClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe cluster params
func (o *DescribeClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCluster adds the cluster to the describe cluster params
func (o *DescribeClusterParams) WithCluster(cluster string) *DescribeClusterParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the describe cluster params
func (o *DescribeClusterParams) SetCluster(cluster string) {
	o.Cluster = cluster
}

// WithOrg adds the org to the describe cluster params
func (o *DescribeClusterParams) WithOrg(org string) *DescribeClusterParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the describe cluster params
func (o *DescribeClusterParams) SetOrg(org string) {
	o.Org = org
}

// WithProject adds the project to the describe cluster params
func (o *DescribeClusterParams) WithProject(project string) *DescribeClusterParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the describe cluster params
func (o *DescribeClusterParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster
	if err := r.SetPathParam("cluster", o.Cluster); err != nil {
		return err
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
