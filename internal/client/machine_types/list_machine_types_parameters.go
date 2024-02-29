// Code generated by go-swagger; DO NOT EDIT.

package machine_types

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

// NewListMachineTypesParams creates a new ListMachineTypesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListMachineTypesParams() *ListMachineTypesParams {
	return &ListMachineTypesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListMachineTypesParamsWithTimeout creates a new ListMachineTypesParams object
// with the ability to set a timeout on a request.
func NewListMachineTypesParamsWithTimeout(timeout time.Duration) *ListMachineTypesParams {
	return &ListMachineTypesParams{
		timeout: timeout,
	}
}

// NewListMachineTypesParamsWithContext creates a new ListMachineTypesParams object
// with the ability to set a context for a request.
func NewListMachineTypesParamsWithContext(ctx context.Context) *ListMachineTypesParams {
	return &ListMachineTypesParams{
		Context: ctx,
	}
}

// NewListMachineTypesParamsWithHTTPClient creates a new ListMachineTypesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListMachineTypesParamsWithHTTPClient(client *http.Client) *ListMachineTypesParams {
	return &ListMachineTypesParams{
		HTTPClient: client,
	}
}

/*
ListMachineTypesParams contains all the parameters to send to the API endpoint

	for the list machine types operation.

	Typically these are written to a http.Request.
*/
type ListMachineTypesParams struct {

	// DataCenterID.
	DataCenterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list machine types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMachineTypesParams) WithDefaults() *ListMachineTypesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list machine types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMachineTypesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list machine types params
func (o *ListMachineTypesParams) WithTimeout(timeout time.Duration) *ListMachineTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list machine types params
func (o *ListMachineTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list machine types params
func (o *ListMachineTypesParams) WithContext(ctx context.Context) *ListMachineTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list machine types params
func (o *ListMachineTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list machine types params
func (o *ListMachineTypesParams) WithHTTPClient(client *http.Client) *ListMachineTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list machine types params
func (o *ListMachineTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDataCenterID adds the dataCenterID to the list machine types params
func (o *ListMachineTypesParams) WithDataCenterID(dataCenterID string) *ListMachineTypesParams {
	o.SetDataCenterID(dataCenterID)
	return o
}

// SetDataCenterID adds the dataCenterId to the list machine types params
func (o *ListMachineTypesParams) SetDataCenterID(dataCenterID string) {
	o.DataCenterID = dataCenterID
}

// WriteToRequest writes these params to a swagger request
func (o *ListMachineTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dataCenterId
	if err := r.SetPathParam("dataCenterId", o.DataCenterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
