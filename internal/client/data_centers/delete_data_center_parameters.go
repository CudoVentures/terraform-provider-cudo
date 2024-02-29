// Code generated by go-swagger; DO NOT EDIT.

package data_centers

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

// NewDeleteDataCenterParams creates a new DeleteDataCenterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDataCenterParams() *DeleteDataCenterParams {
	return &DeleteDataCenterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDataCenterParamsWithTimeout creates a new DeleteDataCenterParams object
// with the ability to set a timeout on a request.
func NewDeleteDataCenterParamsWithTimeout(timeout time.Duration) *DeleteDataCenterParams {
	return &DeleteDataCenterParams{
		timeout: timeout,
	}
}

// NewDeleteDataCenterParamsWithContext creates a new DeleteDataCenterParams object
// with the ability to set a context for a request.
func NewDeleteDataCenterParamsWithContext(ctx context.Context) *DeleteDataCenterParams {
	return &DeleteDataCenterParams{
		Context: ctx,
	}
}

// NewDeleteDataCenterParamsWithHTTPClient creates a new DeleteDataCenterParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDataCenterParamsWithHTTPClient(client *http.Client) *DeleteDataCenterParams {
	return &DeleteDataCenterParams{
		HTTPClient: client,
	}
}

/*
DeleteDataCenterParams contains all the parameters to send to the API endpoint

	for the delete data center operation.

	Typically these are written to a http.Request.
*/
type DeleteDataCenterParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete data center params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDataCenterParams) WithDefaults() *DeleteDataCenterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete data center params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDataCenterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete data center params
func (o *DeleteDataCenterParams) WithTimeout(timeout time.Duration) *DeleteDataCenterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete data center params
func (o *DeleteDataCenterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete data center params
func (o *DeleteDataCenterParams) WithContext(ctx context.Context) *DeleteDataCenterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete data center params
func (o *DeleteDataCenterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete data center params
func (o *DeleteDataCenterParams) WithHTTPClient(client *http.Client) *DeleteDataCenterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete data center params
func (o *DeleteDataCenterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete data center params
func (o *DeleteDataCenterParams) WithID(id string) *DeleteDataCenterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete data center params
func (o *DeleteDataCenterParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDataCenterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
