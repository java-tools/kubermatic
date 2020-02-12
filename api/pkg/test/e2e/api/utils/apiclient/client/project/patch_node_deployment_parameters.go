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

	strfmt "github.com/go-openapi/strfmt"
)

// NewPatchNodeDeploymentParams creates a new PatchNodeDeploymentParams object
// with the default values initialized.
func NewPatchNodeDeploymentParams() *PatchNodeDeploymentParams {
	var ()
	return &PatchNodeDeploymentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchNodeDeploymentParamsWithTimeout creates a new PatchNodeDeploymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchNodeDeploymentParamsWithTimeout(timeout time.Duration) *PatchNodeDeploymentParams {
	var ()
	return &PatchNodeDeploymentParams{

		timeout: timeout,
	}
}

// NewPatchNodeDeploymentParamsWithContext creates a new PatchNodeDeploymentParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchNodeDeploymentParamsWithContext(ctx context.Context) *PatchNodeDeploymentParams {
	var ()
	return &PatchNodeDeploymentParams{

		Context: ctx,
	}
}

// NewPatchNodeDeploymentParamsWithHTTPClient creates a new PatchNodeDeploymentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchNodeDeploymentParamsWithHTTPClient(client *http.Client) *PatchNodeDeploymentParams {
	var ()
	return &PatchNodeDeploymentParams{
		HTTPClient: client,
	}
}

/*PatchNodeDeploymentParams contains all the parameters to send to the API endpoint
for the patch node deployment operation typically these are written to a http.Request
*/
type PatchNodeDeploymentParams struct {

	/*Patch*/
	Patch interface{}
	/*ClusterID*/
	ClusterID string
	/*Dc*/
	DC string
	/*NodedeploymentID*/
	NodeDeploymentID string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch node deployment params
func (o *PatchNodeDeploymentParams) WithTimeout(timeout time.Duration) *PatchNodeDeploymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch node deployment params
func (o *PatchNodeDeploymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch node deployment params
func (o *PatchNodeDeploymentParams) WithContext(ctx context.Context) *PatchNodeDeploymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch node deployment params
func (o *PatchNodeDeploymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch node deployment params
func (o *PatchNodeDeploymentParams) WithHTTPClient(client *http.Client) *PatchNodeDeploymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch node deployment params
func (o *PatchNodeDeploymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatch adds the patch to the patch node deployment params
func (o *PatchNodeDeploymentParams) WithPatch(patch interface{}) *PatchNodeDeploymentParams {
	o.SetPatch(patch)
	return o
}

// SetPatch adds the patch to the patch node deployment params
func (o *PatchNodeDeploymentParams) SetPatch(patch interface{}) {
	o.Patch = patch
}

// WithClusterID adds the clusterID to the patch node deployment params
func (o *PatchNodeDeploymentParams) WithClusterID(clusterID string) *PatchNodeDeploymentParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the patch node deployment params
func (o *PatchNodeDeploymentParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the patch node deployment params
func (o *PatchNodeDeploymentParams) WithDC(dc string) *PatchNodeDeploymentParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the patch node deployment params
func (o *PatchNodeDeploymentParams) SetDC(dc string) {
	o.DC = dc
}

// WithNodeDeploymentID adds the nodedeploymentID to the patch node deployment params
func (o *PatchNodeDeploymentParams) WithNodeDeploymentID(nodedeploymentID string) *PatchNodeDeploymentParams {
	o.SetNodeDeploymentID(nodedeploymentID)
	return o
}

// SetNodeDeploymentID adds the nodedeploymentId to the patch node deployment params
func (o *PatchNodeDeploymentParams) SetNodeDeploymentID(nodedeploymentID string) {
	o.NodeDeploymentID = nodedeploymentID
}

// WithProjectID adds the projectID to the patch node deployment params
func (o *PatchNodeDeploymentParams) WithProjectID(projectID string) *PatchNodeDeploymentParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the patch node deployment params
func (o *PatchNodeDeploymentParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchNodeDeploymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Patch != nil {
		if err := r.SetBodyParam(o.Patch); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
		return err
	}

	// path param nodedeployment_id
	if err := r.SetPathParam("nodedeployment_id", o.NodeDeploymentID); err != nil {
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
