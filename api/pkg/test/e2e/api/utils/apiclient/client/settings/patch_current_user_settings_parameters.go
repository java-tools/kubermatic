// Code generated by go-swagger; DO NOT EDIT.

package settings

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

// NewPatchCurrentUserSettingsParams creates a new PatchCurrentUserSettingsParams object
// with the default values initialized.
func NewPatchCurrentUserSettingsParams() *PatchCurrentUserSettingsParams {
	var ()
	return &PatchCurrentUserSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCurrentUserSettingsParamsWithTimeout creates a new PatchCurrentUserSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchCurrentUserSettingsParamsWithTimeout(timeout time.Duration) *PatchCurrentUserSettingsParams {
	var ()
	return &PatchCurrentUserSettingsParams{

		timeout: timeout,
	}
}

// NewPatchCurrentUserSettingsParamsWithContext creates a new PatchCurrentUserSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchCurrentUserSettingsParamsWithContext(ctx context.Context) *PatchCurrentUserSettingsParams {
	var ()
	return &PatchCurrentUserSettingsParams{

		Context: ctx,
	}
}

// NewPatchCurrentUserSettingsParamsWithHTTPClient creates a new PatchCurrentUserSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchCurrentUserSettingsParamsWithHTTPClient(client *http.Client) *PatchCurrentUserSettingsParams {
	var ()
	return &PatchCurrentUserSettingsParams{
		HTTPClient: client,
	}
}

/*PatchCurrentUserSettingsParams contains all the parameters to send to the API endpoint
for the patch current user settings operation typically these are written to a http.Request
*/
type PatchCurrentUserSettingsParams struct {

	/*Patch*/
	Patch interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch current user settings params
func (o *PatchCurrentUserSettingsParams) WithTimeout(timeout time.Duration) *PatchCurrentUserSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch current user settings params
func (o *PatchCurrentUserSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch current user settings params
func (o *PatchCurrentUserSettingsParams) WithContext(ctx context.Context) *PatchCurrentUserSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch current user settings params
func (o *PatchCurrentUserSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch current user settings params
func (o *PatchCurrentUserSettingsParams) WithHTTPClient(client *http.Client) *PatchCurrentUserSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch current user settings params
func (o *PatchCurrentUserSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatch adds the patch to the patch current user settings params
func (o *PatchCurrentUserSettingsParams) WithPatch(patch interface{}) *PatchCurrentUserSettingsParams {
	o.SetPatch(patch)
	return o
}

// SetPatch adds the patch to the patch current user settings params
func (o *PatchCurrentUserSettingsParams) SetPatch(patch interface{}) {
	o.Patch = patch
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCurrentUserSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Patch != nil {
		if err := r.SetBodyParam(o.Patch); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
