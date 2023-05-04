// Code generated by go-swagger; DO NOT EDIT.

package ssh_keys

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

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// NewCreateSSHKeyParams creates a new CreateSSHKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSSHKeyParams() *CreateSSHKeyParams {
	return &CreateSSHKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSSHKeyParamsWithTimeout creates a new CreateSSHKeyParams object
// with the ability to set a timeout on a request.
func NewCreateSSHKeyParamsWithTimeout(timeout time.Duration) *CreateSSHKeyParams {
	return &CreateSSHKeyParams{
		timeout: timeout,
	}
}

// NewCreateSSHKeyParamsWithContext creates a new CreateSSHKeyParams object
// with the ability to set a context for a request.
func NewCreateSSHKeyParamsWithContext(ctx context.Context) *CreateSSHKeyParams {
	return &CreateSSHKeyParams{
		Context: ctx,
	}
}

// NewCreateSSHKeyParamsWithHTTPClient creates a new CreateSSHKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSSHKeyParamsWithHTTPClient(client *http.Client) *CreateSSHKeyParams {
	return &CreateSSHKeyParams{
		HTTPClient: client,
	}
}

/*
CreateSSHKeyParams contains all the parameters to send to the API endpoint

	for the create Ssh key operation.

	Typically these are written to a http.Request.
*/
type CreateSSHKeyParams struct {

	// SSHKey.
	SSHKey *models.SSHKey

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create Ssh key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSSHKeyParams) WithDefaults() *CreateSSHKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create Ssh key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSSHKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create Ssh key params
func (o *CreateSSHKeyParams) WithTimeout(timeout time.Duration) *CreateSSHKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create Ssh key params
func (o *CreateSSHKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create Ssh key params
func (o *CreateSSHKeyParams) WithContext(ctx context.Context) *CreateSSHKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create Ssh key params
func (o *CreateSSHKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create Ssh key params
func (o *CreateSSHKeyParams) WithHTTPClient(client *http.Client) *CreateSSHKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create Ssh key params
func (o *CreateSSHKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSSHKey adds the sSHKey to the create Ssh key params
func (o *CreateSSHKeyParams) WithSSHKey(sSHKey *models.SSHKey) *CreateSSHKeyParams {
	o.SetSSHKey(sSHKey)
	return o
}

// SetSSHKey adds the sshKey to the create Ssh key params
func (o *CreateSSHKeyParams) SetSSHKey(sSHKey *models.SSHKey) {
	o.SSHKey = sSHKey
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSSHKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.SSHKey != nil {
		if err := r.SetBodyParam(o.SSHKey); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
