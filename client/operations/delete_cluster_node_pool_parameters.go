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

// NewDeleteClusterNodePoolParams creates a new DeleteClusterNodePoolParams object
// with the default values initialized.
func NewDeleteClusterNodePoolParams() *DeleteClusterNodePoolParams {
	var ()
	return &DeleteClusterNodePoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClusterNodePoolParamsWithTimeout creates a new DeleteClusterNodePoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClusterNodePoolParamsWithTimeout(timeout time.Duration) *DeleteClusterNodePoolParams {
	var ()
	return &DeleteClusterNodePoolParams{

		timeout: timeout,
	}
}

// NewDeleteClusterNodePoolParamsWithContext creates a new DeleteClusterNodePoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClusterNodePoolParamsWithContext(ctx context.Context) *DeleteClusterNodePoolParams {
	var ()
	return &DeleteClusterNodePoolParams{

		Context: ctx,
	}
}

// NewDeleteClusterNodePoolParamsWithHTTPClient creates a new DeleteClusterNodePoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteClusterNodePoolParamsWithHTTPClient(client *http.Client) *DeleteClusterNodePoolParams {
	var ()
	return &DeleteClusterNodePoolParams{
		HTTPClient: client,
	}
}

/*DeleteClusterNodePoolParams contains all the parameters to send to the API endpoint
for the delete cluster node pool operation typically these are written to a http.Request
*/
type DeleteClusterNodePoolParams struct {

	/*Cluster
	  The identifier of a cluster.

	*/
	Cluster string
	/*NodePool
	  The identifier of a node pool in a cluster

	*/
	NodePool string
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

// WithTimeout adds the timeout to the delete cluster node pool params
func (o *DeleteClusterNodePoolParams) WithTimeout(timeout time.Duration) *DeleteClusterNodePoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cluster node pool params
func (o *DeleteClusterNodePoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cluster node pool params
func (o *DeleteClusterNodePoolParams) WithContext(ctx context.Context) *DeleteClusterNodePoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cluster node pool params
func (o *DeleteClusterNodePoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cluster node pool params
func (o *DeleteClusterNodePoolParams) WithHTTPClient(client *http.Client) *DeleteClusterNodePoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cluster node pool params
func (o *DeleteClusterNodePoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCluster adds the cluster to the delete cluster node pool params
func (o *DeleteClusterNodePoolParams) WithCluster(cluster string) *DeleteClusterNodePoolParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the delete cluster node pool params
func (o *DeleteClusterNodePoolParams) SetCluster(cluster string) {
	o.Cluster = cluster
}

// WithNodePool adds the nodePool to the delete cluster node pool params
func (o *DeleteClusterNodePoolParams) WithNodePool(nodePool string) *DeleteClusterNodePoolParams {
	o.SetNodePool(nodePool)
	return o
}

// SetNodePool adds the nodePool to the delete cluster node pool params
func (o *DeleteClusterNodePoolParams) SetNodePool(nodePool string) {
	o.NodePool = nodePool
}

// WithOrg adds the org to the delete cluster node pool params
func (o *DeleteClusterNodePoolParams) WithOrg(org string) *DeleteClusterNodePoolParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the delete cluster node pool params
func (o *DeleteClusterNodePoolParams) SetOrg(org string) {
	o.Org = org
}

// WithProject adds the project to the delete cluster node pool params
func (o *DeleteClusterNodePoolParams) WithProject(project string) *DeleteClusterNodePoolParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the delete cluster node pool params
func (o *DeleteClusterNodePoolParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClusterNodePoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster
	if err := r.SetPathParam("cluster", o.Cluster); err != nil {
		return err
	}

	// path param node_pool
	if err := r.SetPathParam("node_pool", o.NodePool); err != nil {
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
