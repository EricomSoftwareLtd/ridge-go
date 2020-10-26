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

// NewUpdateClusterNodePoolParams creates a new UpdateClusterNodePoolParams object
// with the default values initialized.
func NewUpdateClusterNodePoolParams() *UpdateClusterNodePoolParams {
	var ()
	return &UpdateClusterNodePoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateClusterNodePoolParamsWithTimeout creates a new UpdateClusterNodePoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateClusterNodePoolParamsWithTimeout(timeout time.Duration) *UpdateClusterNodePoolParams {
	var ()
	return &UpdateClusterNodePoolParams{

		timeout: timeout,
	}
}

// NewUpdateClusterNodePoolParamsWithContext creates a new UpdateClusterNodePoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateClusterNodePoolParamsWithContext(ctx context.Context) *UpdateClusterNodePoolParams {
	var ()
	return &UpdateClusterNodePoolParams{

		Context: ctx,
	}
}

// NewUpdateClusterNodePoolParamsWithHTTPClient creates a new UpdateClusterNodePoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateClusterNodePoolParamsWithHTTPClient(client *http.Client) *UpdateClusterNodePoolParams {
	var ()
	return &UpdateClusterNodePoolParams{
		HTTPClient: client,
	}
}

/*UpdateClusterNodePoolParams contains all the parameters to send to the API endpoint
for the update cluster node pool operation typically these are written to a http.Request
*/
type UpdateClusterNodePoolParams struct {

	/*Body*/
	Body *models.NodePoolUpdate
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

// WithTimeout adds the timeout to the update cluster node pool params
func (o *UpdateClusterNodePoolParams) WithTimeout(timeout time.Duration) *UpdateClusterNodePoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update cluster node pool params
func (o *UpdateClusterNodePoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update cluster node pool params
func (o *UpdateClusterNodePoolParams) WithContext(ctx context.Context) *UpdateClusterNodePoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update cluster node pool params
func (o *UpdateClusterNodePoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update cluster node pool params
func (o *UpdateClusterNodePoolParams) WithHTTPClient(client *http.Client) *UpdateClusterNodePoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update cluster node pool params
func (o *UpdateClusterNodePoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update cluster node pool params
func (o *UpdateClusterNodePoolParams) WithBody(body *models.NodePoolUpdate) *UpdateClusterNodePoolParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update cluster node pool params
func (o *UpdateClusterNodePoolParams) SetBody(body *models.NodePoolUpdate) {
	o.Body = body
}

// WithCluster adds the cluster to the update cluster node pool params
func (o *UpdateClusterNodePoolParams) WithCluster(cluster string) *UpdateClusterNodePoolParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the update cluster node pool params
func (o *UpdateClusterNodePoolParams) SetCluster(cluster string) {
	o.Cluster = cluster
}

// WithNodePool adds the nodePool to the update cluster node pool params
func (o *UpdateClusterNodePoolParams) WithNodePool(nodePool string) *UpdateClusterNodePoolParams {
	o.SetNodePool(nodePool)
	return o
}

// SetNodePool adds the nodePool to the update cluster node pool params
func (o *UpdateClusterNodePoolParams) SetNodePool(nodePool string) {
	o.NodePool = nodePool
}

// WithOrg adds the org to the update cluster node pool params
func (o *UpdateClusterNodePoolParams) WithOrg(org string) *UpdateClusterNodePoolParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the update cluster node pool params
func (o *UpdateClusterNodePoolParams) SetOrg(org string) {
	o.Org = org
}

// WithProject adds the project to the update cluster node pool params
func (o *UpdateClusterNodePoolParams) WithProject(project string) *UpdateClusterNodePoolParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the update cluster node pool params
func (o *UpdateClusterNodePoolParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateClusterNodePoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
