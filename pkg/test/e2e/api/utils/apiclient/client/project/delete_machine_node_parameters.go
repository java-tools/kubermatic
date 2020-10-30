// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewDeleteMachineNodeParams creates a new DeleteMachineNodeParams object
// with the default values initialized.
func NewDeleteMachineNodeParams() *DeleteMachineNodeParams {
	var ()
	return &DeleteMachineNodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMachineNodeParamsWithTimeout creates a new DeleteMachineNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMachineNodeParamsWithTimeout(timeout time.Duration) *DeleteMachineNodeParams {
	var ()
	return &DeleteMachineNodeParams{

		timeout: timeout,
	}
}

// NewDeleteMachineNodeParamsWithContext creates a new DeleteMachineNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMachineNodeParamsWithContext(ctx context.Context) *DeleteMachineNodeParams {
	var ()
	return &DeleteMachineNodeParams{

		Context: ctx,
	}
}

// NewDeleteMachineNodeParamsWithHTTPClient creates a new DeleteMachineNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMachineNodeParamsWithHTTPClient(client *http.Client) *DeleteMachineNodeParams {
	var ()
	return &DeleteMachineNodeParams{
		HTTPClient: client,
	}
}

/*DeleteMachineNodeParams contains all the parameters to send to the API endpoint
for the delete machine node operation typically these are written to a http.Request
*/
type DeleteMachineNodeParams struct {

	/*ClusterID*/
	ClusterID string
	/*NodeID*/
	NodeID string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete machine node params
func (o *DeleteMachineNodeParams) WithTimeout(timeout time.Duration) *DeleteMachineNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete machine node params
func (o *DeleteMachineNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete machine node params
func (o *DeleteMachineNodeParams) WithContext(ctx context.Context) *DeleteMachineNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete machine node params
func (o *DeleteMachineNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete machine node params
func (o *DeleteMachineNodeParams) WithHTTPClient(client *http.Client) *DeleteMachineNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete machine node params
func (o *DeleteMachineNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the delete machine node params
func (o *DeleteMachineNodeParams) WithClusterID(clusterID string) *DeleteMachineNodeParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete machine node params
func (o *DeleteMachineNodeParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithNodeID adds the nodeID to the delete machine node params
func (o *DeleteMachineNodeParams) WithNodeID(nodeID string) *DeleteMachineNodeParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the delete machine node params
func (o *DeleteMachineNodeParams) SetNodeID(nodeID string) {
	o.NodeID = nodeID
}

// WithProjectID adds the projectID to the delete machine node params
func (o *DeleteMachineNodeParams) WithProjectID(projectID string) *DeleteMachineNodeParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete machine node params
func (o *DeleteMachineNodeParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMachineNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param node_id
	if err := r.SetPathParam("node_id", o.NodeID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}