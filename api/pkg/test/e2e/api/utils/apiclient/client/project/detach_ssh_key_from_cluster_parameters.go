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

// NewDetachSSHKeyFromClusterParams creates a new DetachSSHKeyFromClusterParams object
// with the default values initialized.
func NewDetachSSHKeyFromClusterParams() *DetachSSHKeyFromClusterParams {
	var ()
	return &DetachSSHKeyFromClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDetachSSHKeyFromClusterParamsWithTimeout creates a new DetachSSHKeyFromClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDetachSSHKeyFromClusterParamsWithTimeout(timeout time.Duration) *DetachSSHKeyFromClusterParams {
	var ()
	return &DetachSSHKeyFromClusterParams{

		timeout: timeout,
	}
}

// NewDetachSSHKeyFromClusterParamsWithContext creates a new DetachSSHKeyFromClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewDetachSSHKeyFromClusterParamsWithContext(ctx context.Context) *DetachSSHKeyFromClusterParams {
	var ()
	return &DetachSSHKeyFromClusterParams{

		Context: ctx,
	}
}

// NewDetachSSHKeyFromClusterParamsWithHTTPClient creates a new DetachSSHKeyFromClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDetachSSHKeyFromClusterParamsWithHTTPClient(client *http.Client) *DetachSSHKeyFromClusterParams {
	var ()
	return &DetachSSHKeyFromClusterParams{
		HTTPClient: client,
	}
}

/*DetachSSHKeyFromClusterParams contains all the parameters to send to the API endpoint
for the detach SSH key from cluster operation typically these are written to a http.Request
*/
type DetachSSHKeyFromClusterParams struct {

	/*ClusterID*/
	ClusterID string
	/*Dc*/
	DC string
	/*KeyID*/
	KeyID string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the detach SSH key from cluster params
func (o *DetachSSHKeyFromClusterParams) WithTimeout(timeout time.Duration) *DetachSSHKeyFromClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the detach SSH key from cluster params
func (o *DetachSSHKeyFromClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the detach SSH key from cluster params
func (o *DetachSSHKeyFromClusterParams) WithContext(ctx context.Context) *DetachSSHKeyFromClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the detach SSH key from cluster params
func (o *DetachSSHKeyFromClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the detach SSH key from cluster params
func (o *DetachSSHKeyFromClusterParams) WithHTTPClient(client *http.Client) *DetachSSHKeyFromClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the detach SSH key from cluster params
func (o *DetachSSHKeyFromClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the detach SSH key from cluster params
func (o *DetachSSHKeyFromClusterParams) WithClusterID(clusterID string) *DetachSSHKeyFromClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the detach SSH key from cluster params
func (o *DetachSSHKeyFromClusterParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the detach SSH key from cluster params
func (o *DetachSSHKeyFromClusterParams) WithDC(dc string) *DetachSSHKeyFromClusterParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the detach SSH key from cluster params
func (o *DetachSSHKeyFromClusterParams) SetDC(dc string) {
	o.DC = dc
}

// WithKeyID adds the keyID to the detach SSH key from cluster params
func (o *DetachSSHKeyFromClusterParams) WithKeyID(keyID string) *DetachSSHKeyFromClusterParams {
	o.SetKeyID(keyID)
	return o
}

// SetKeyID adds the keyId to the detach SSH key from cluster params
func (o *DetachSSHKeyFromClusterParams) SetKeyID(keyID string) {
	o.KeyID = keyID
}

// WithProjectID adds the projectID to the detach SSH key from cluster params
func (o *DetachSSHKeyFromClusterParams) WithProjectID(projectID string) *DetachSSHKeyFromClusterParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the detach SSH key from cluster params
func (o *DetachSSHKeyFromClusterParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DetachSSHKeyFromClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param key_id
	if err := r.SetPathParam("key_id", o.KeyID); err != nil {
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
