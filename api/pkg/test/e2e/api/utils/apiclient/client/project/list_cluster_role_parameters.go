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

// NewListClusterRoleParams creates a new ListClusterRoleParams object
// with the default values initialized.
func NewListClusterRoleParams() *ListClusterRoleParams {
	var ()
	return &ListClusterRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListClusterRoleParamsWithTimeout creates a new ListClusterRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListClusterRoleParamsWithTimeout(timeout time.Duration) *ListClusterRoleParams {
	var ()
	return &ListClusterRoleParams{

		timeout: timeout,
	}
}

// NewListClusterRoleParamsWithContext creates a new ListClusterRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewListClusterRoleParamsWithContext(ctx context.Context) *ListClusterRoleParams {
	var ()
	return &ListClusterRoleParams{

		Context: ctx,
	}
}

// NewListClusterRoleParamsWithHTTPClient creates a new ListClusterRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListClusterRoleParamsWithHTTPClient(client *http.Client) *ListClusterRoleParams {
	var ()
	return &ListClusterRoleParams{
		HTTPClient: client,
	}
}

/*ListClusterRoleParams contains all the parameters to send to the API endpoint
for the list cluster role operation typically these are written to a http.Request
*/
type ListClusterRoleParams struct {

	/*ClusterID*/
	ClusterID string
	/*Dc*/
	DC string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list cluster role params
func (o *ListClusterRoleParams) WithTimeout(timeout time.Duration) *ListClusterRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list cluster role params
func (o *ListClusterRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list cluster role params
func (o *ListClusterRoleParams) WithContext(ctx context.Context) *ListClusterRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list cluster role params
func (o *ListClusterRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list cluster role params
func (o *ListClusterRoleParams) WithHTTPClient(client *http.Client) *ListClusterRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list cluster role params
func (o *ListClusterRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list cluster role params
func (o *ListClusterRoleParams) WithClusterID(clusterID string) *ListClusterRoleParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list cluster role params
func (o *ListClusterRoleParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the list cluster role params
func (o *ListClusterRoleParams) WithDC(dc string) *ListClusterRoleParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the list cluster role params
func (o *ListClusterRoleParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the list cluster role params
func (o *ListClusterRoleParams) WithProjectID(projectID string) *ListClusterRoleParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list cluster role params
func (o *ListClusterRoleParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListClusterRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
