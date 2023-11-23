// Code generated by go-swagger; DO NOT EDIT.

package object_storage

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

// NewDeleteObjectStorageUserParams creates a new DeleteObjectStorageUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteObjectStorageUserParams() *DeleteObjectStorageUserParams {
	return &DeleteObjectStorageUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteObjectStorageUserParamsWithTimeout creates a new DeleteObjectStorageUserParams object
// with the ability to set a timeout on a request.
func NewDeleteObjectStorageUserParamsWithTimeout(timeout time.Duration) *DeleteObjectStorageUserParams {
	return &DeleteObjectStorageUserParams{
		timeout: timeout,
	}
}

// NewDeleteObjectStorageUserParamsWithContext creates a new DeleteObjectStorageUserParams object
// with the ability to set a context for a request.
func NewDeleteObjectStorageUserParamsWithContext(ctx context.Context) *DeleteObjectStorageUserParams {
	return &DeleteObjectStorageUserParams{
		Context: ctx,
	}
}

// NewDeleteObjectStorageUserParamsWithHTTPClient creates a new DeleteObjectStorageUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteObjectStorageUserParamsWithHTTPClient(client *http.Client) *DeleteObjectStorageUserParams {
	return &DeleteObjectStorageUserParams{
		HTTPClient: client,
	}
}

/*
DeleteObjectStorageUserParams contains all the parameters to send to the API endpoint

	for the delete object storage user operation.

	Typically these are written to a http.Request.
*/
type DeleteObjectStorageUserParams struct {

	// ID.
	ID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete object storage user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteObjectStorageUserParams) WithDefaults() *DeleteObjectStorageUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete object storage user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteObjectStorageUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete object storage user params
func (o *DeleteObjectStorageUserParams) WithTimeout(timeout time.Duration) *DeleteObjectStorageUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete object storage user params
func (o *DeleteObjectStorageUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete object storage user params
func (o *DeleteObjectStorageUserParams) WithContext(ctx context.Context) *DeleteObjectStorageUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete object storage user params
func (o *DeleteObjectStorageUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete object storage user params
func (o *DeleteObjectStorageUserParams) WithHTTPClient(client *http.Client) *DeleteObjectStorageUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete object storage user params
func (o *DeleteObjectStorageUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete object storage user params
func (o *DeleteObjectStorageUserParams) WithID(id string) *DeleteObjectStorageUserParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete object storage user params
func (o *DeleteObjectStorageUserParams) SetID(id string) {
	o.ID = id
}

// WithProjectID adds the projectID to the delete object storage user params
func (o *DeleteObjectStorageUserParams) WithProjectID(projectID string) *DeleteObjectStorageUserParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete object storage user params
func (o *DeleteObjectStorageUserParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteObjectStorageUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
