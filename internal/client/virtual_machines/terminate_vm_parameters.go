// Code generated by go-swagger; DO NOT EDIT.

package virtual_machines

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

// NewTerminateVMParams creates a new TerminateVMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTerminateVMParams() *TerminateVMParams {
	return &TerminateVMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTerminateVMParamsWithTimeout creates a new TerminateVMParams object
// with the ability to set a timeout on a request.
func NewTerminateVMParamsWithTimeout(timeout time.Duration) *TerminateVMParams {
	return &TerminateVMParams{
		timeout: timeout,
	}
}

// NewTerminateVMParamsWithContext creates a new TerminateVMParams object
// with the ability to set a context for a request.
func NewTerminateVMParamsWithContext(ctx context.Context) *TerminateVMParams {
	return &TerminateVMParams{
		Context: ctx,
	}
}

// NewTerminateVMParamsWithHTTPClient creates a new TerminateVMParams object
// with the ability to set a custom HTTPClient for a request.
func NewTerminateVMParamsWithHTTPClient(client *http.Client) *TerminateVMParams {
	return &TerminateVMParams{
		HTTPClient: client,
	}
}

/*
TerminateVMParams contains all the parameters to send to the API endpoint

	for the terminate VM operation.

	Typically these are written to a http.Request.
*/
type TerminateVMParams struct {

	// ID.
	ID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the terminate VM params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TerminateVMParams) WithDefaults() *TerminateVMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the terminate VM params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TerminateVMParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the terminate VM params
func (o *TerminateVMParams) WithTimeout(timeout time.Duration) *TerminateVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the terminate VM params
func (o *TerminateVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the terminate VM params
func (o *TerminateVMParams) WithContext(ctx context.Context) *TerminateVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the terminate VM params
func (o *TerminateVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the terminate VM params
func (o *TerminateVMParams) WithHTTPClient(client *http.Client) *TerminateVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the terminate VM params
func (o *TerminateVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the terminate VM params
func (o *TerminateVMParams) WithID(id string) *TerminateVMParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the terminate VM params
func (o *TerminateVMParams) SetID(id string) {
	o.ID = id
}

// WithProjectID adds the projectID to the terminate VM params
func (o *TerminateVMParams) WithProjectID(projectID string) *TerminateVMParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the terminate VM params
func (o *TerminateVMParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *TerminateVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param projectId
	if err := r.SetPathParam("projectId", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}