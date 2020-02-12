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

// NewDeleteClusterRoleParams creates a new DeleteClusterRoleParams object
// with the default values initialized.
func NewDeleteClusterRoleParams() *DeleteClusterRoleParams {
	var ()
	return &DeleteClusterRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClusterRoleParamsWithTimeout creates a new DeleteClusterRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClusterRoleParamsWithTimeout(timeout time.Duration) *DeleteClusterRoleParams {
	var ()
	return &DeleteClusterRoleParams{

		timeout: timeout,
	}
}

// NewDeleteClusterRoleParamsWithContext creates a new DeleteClusterRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClusterRoleParamsWithContext(ctx context.Context) *DeleteClusterRoleParams {
	var ()
	return &DeleteClusterRoleParams{

		Context: ctx,
	}
}

// NewDeleteClusterRoleParamsWithHTTPClient creates a new DeleteClusterRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteClusterRoleParamsWithHTTPClient(client *http.Client) *DeleteClusterRoleParams {
	var ()
	return &DeleteClusterRoleParams{
		HTTPClient: client,
	}
}

/*DeleteClusterRoleParams contains all the parameters to send to the API endpoint
for the delete cluster role operation typically these are written to a http.Request
*/
type DeleteClusterRoleParams struct {

	/*ClusterID*/
	ClusterID string
	/*Dc*/
	DC string
	/*ProjectID*/
	ProjectID string
	/*RoleID*/
	RoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cluster role params
func (o *DeleteClusterRoleParams) WithTimeout(timeout time.Duration) *DeleteClusterRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cluster role params
func (o *DeleteClusterRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cluster role params
func (o *DeleteClusterRoleParams) WithContext(ctx context.Context) *DeleteClusterRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cluster role params
func (o *DeleteClusterRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cluster role params
func (o *DeleteClusterRoleParams) WithHTTPClient(client *http.Client) *DeleteClusterRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cluster role params
func (o *DeleteClusterRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the delete cluster role params
func (o *DeleteClusterRoleParams) WithClusterID(clusterID string) *DeleteClusterRoleParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete cluster role params
func (o *DeleteClusterRoleParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the delete cluster role params
func (o *DeleteClusterRoleParams) WithDC(dc string) *DeleteClusterRoleParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the delete cluster role params
func (o *DeleteClusterRoleParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the delete cluster role params
func (o *DeleteClusterRoleParams) WithProjectID(projectID string) *DeleteClusterRoleParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete cluster role params
func (o *DeleteClusterRoleParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithRoleID adds the roleID to the delete cluster role params
func (o *DeleteClusterRoleParams) WithRoleID(roleID string) *DeleteClusterRoleParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the delete cluster role params
func (o *DeleteClusterRoleParams) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClusterRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	// path param role_id
	if err := r.SetPathParam("role_id", o.RoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
