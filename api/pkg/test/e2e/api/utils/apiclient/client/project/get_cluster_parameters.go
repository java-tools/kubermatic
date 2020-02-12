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

// NewGetClusterParams creates a new GetClusterParams object
// with the default values initialized.
func NewGetClusterParams() *GetClusterParams {
	var ()
	return &GetClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterParamsWithTimeout creates a new GetClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterParamsWithTimeout(timeout time.Duration) *GetClusterParams {
	var ()
	return &GetClusterParams{

		timeout: timeout,
	}
}

// NewGetClusterParamsWithContext creates a new GetClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterParamsWithContext(ctx context.Context) *GetClusterParams {
	var ()
	return &GetClusterParams{

		Context: ctx,
	}
}

// NewGetClusterParamsWithHTTPClient creates a new GetClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterParamsWithHTTPClient(client *http.Client) *GetClusterParams {
	var ()
	return &GetClusterParams{
		HTTPClient: client,
	}
}

/*GetClusterParams contains all the parameters to send to the API endpoint
for the get cluster operation typically these are written to a http.Request
*/
type GetClusterParams struct {

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

// WithTimeout adds the timeout to the get cluster params
func (o *GetClusterParams) WithTimeout(timeout time.Duration) *GetClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster params
func (o *GetClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster params
func (o *GetClusterParams) WithContext(ctx context.Context) *GetClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster params
func (o *GetClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster params
func (o *GetClusterParams) WithHTTPClient(client *http.Client) *GetClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster params
func (o *GetClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get cluster params
func (o *GetClusterParams) WithClusterID(clusterID string) *GetClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster params
func (o *GetClusterParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the get cluster params
func (o *GetClusterParams) WithDC(dc string) *GetClusterParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the get cluster params
func (o *GetClusterParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the get cluster params
func (o *GetClusterParams) WithProjectID(projectID string) *GetClusterParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get cluster params
func (o *GetClusterParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
