// Code generated by go-swagger; DO NOT EDIT.

package serviceaccounts

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

// NewDeleteServiceAccountParams creates a new DeleteServiceAccountParams object
// with the default values initialized.
func NewDeleteServiceAccountParams() *DeleteServiceAccountParams {
	var ()
	return &DeleteServiceAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteServiceAccountParamsWithTimeout creates a new DeleteServiceAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteServiceAccountParamsWithTimeout(timeout time.Duration) *DeleteServiceAccountParams {
	var ()
	return &DeleteServiceAccountParams{

		timeout: timeout,
	}
}

// NewDeleteServiceAccountParamsWithContext creates a new DeleteServiceAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteServiceAccountParamsWithContext(ctx context.Context) *DeleteServiceAccountParams {
	var ()
	return &DeleteServiceAccountParams{

		Context: ctx,
	}
}

// NewDeleteServiceAccountParamsWithHTTPClient creates a new DeleteServiceAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteServiceAccountParamsWithHTTPClient(client *http.Client) *DeleteServiceAccountParams {
	var ()
	return &DeleteServiceAccountParams{
		HTTPClient: client,
	}
}

/*DeleteServiceAccountParams contains all the parameters to send to the API endpoint
for the delete service account operation typically these are written to a http.Request
*/
type DeleteServiceAccountParams struct {

	/*ProjectID*/
	ProjectID string
	/*ServiceaccountID*/
	ServiceAccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete service account params
func (o *DeleteServiceAccountParams) WithTimeout(timeout time.Duration) *DeleteServiceAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete service account params
func (o *DeleteServiceAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete service account params
func (o *DeleteServiceAccountParams) WithContext(ctx context.Context) *DeleteServiceAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete service account params
func (o *DeleteServiceAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete service account params
func (o *DeleteServiceAccountParams) WithHTTPClient(client *http.Client) *DeleteServiceAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete service account params
func (o *DeleteServiceAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the delete service account params
func (o *DeleteServiceAccountParams) WithProjectID(projectID string) *DeleteServiceAccountParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete service account params
func (o *DeleteServiceAccountParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithServiceAccountID adds the serviceaccountID to the delete service account params
func (o *DeleteServiceAccountParams) WithServiceAccountID(serviceaccountID string) *DeleteServiceAccountParams {
	o.SetServiceAccountID(serviceaccountID)
	return o
}

// SetServiceAccountID adds the serviceaccountId to the delete service account params
func (o *DeleteServiceAccountParams) SetServiceAccountID(serviceaccountID string) {
	o.ServiceAccountID = serviceaccountID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteServiceAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	// path param serviceaccount_id
	if err := r.SetPathParam("serviceaccount_id", o.ServiceAccountID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
