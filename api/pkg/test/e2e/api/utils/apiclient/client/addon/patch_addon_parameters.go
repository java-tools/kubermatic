// Code generated by go-swagger; DO NOT EDIT.

package addon

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

	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// NewPatchAddonParams creates a new PatchAddonParams object
// with the default values initialized.
func NewPatchAddonParams() *PatchAddonParams {
	var ()
	return &PatchAddonParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAddonParamsWithTimeout creates a new PatchAddonParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAddonParamsWithTimeout(timeout time.Duration) *PatchAddonParams {
	var ()
	return &PatchAddonParams{

		timeout: timeout,
	}
}

// NewPatchAddonParamsWithContext creates a new PatchAddonParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAddonParamsWithContext(ctx context.Context) *PatchAddonParams {
	var ()
	return &PatchAddonParams{

		Context: ctx,
	}
}

// NewPatchAddonParamsWithHTTPClient creates a new PatchAddonParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAddonParamsWithHTTPClient(client *http.Client) *PatchAddonParams {
	var ()
	return &PatchAddonParams{
		HTTPClient: client,
	}
}

/*PatchAddonParams contains all the parameters to send to the API endpoint
for the patch addon operation typically these are written to a http.Request
*/
type PatchAddonParams struct {

	/*Body*/
	Body *models.Addon
	/*AddonID*/
	AddonID string
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

// WithTimeout adds the timeout to the patch addon params
func (o *PatchAddonParams) WithTimeout(timeout time.Duration) *PatchAddonParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch addon params
func (o *PatchAddonParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch addon params
func (o *PatchAddonParams) WithContext(ctx context.Context) *PatchAddonParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch addon params
func (o *PatchAddonParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch addon params
func (o *PatchAddonParams) WithHTTPClient(client *http.Client) *PatchAddonParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch addon params
func (o *PatchAddonParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch addon params
func (o *PatchAddonParams) WithBody(body *models.Addon) *PatchAddonParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch addon params
func (o *PatchAddonParams) SetBody(body *models.Addon) {
	o.Body = body
}

// WithAddonID adds the addonID to the patch addon params
func (o *PatchAddonParams) WithAddonID(addonID string) *PatchAddonParams {
	o.SetAddonID(addonID)
	return o
}

// SetAddonID adds the addonId to the patch addon params
func (o *PatchAddonParams) SetAddonID(addonID string) {
	o.AddonID = addonID
}

// WithClusterID adds the clusterID to the patch addon params
func (o *PatchAddonParams) WithClusterID(clusterID string) *PatchAddonParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the patch addon params
func (o *PatchAddonParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the patch addon params
func (o *PatchAddonParams) WithDC(dc string) *PatchAddonParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the patch addon params
func (o *PatchAddonParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the patch addon params
func (o *PatchAddonParams) WithProjectID(projectID string) *PatchAddonParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the patch addon params
func (o *PatchAddonParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAddonParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param addon_id
	if err := r.SetPathParam("addon_id", o.AddonID); err != nil {
		return err
	}

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
