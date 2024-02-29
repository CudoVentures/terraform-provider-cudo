// Code generated by go-swagger; DO NOT EDIT.

package billing

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

// NewGetBillingAccountSpendDetailsParams creates a new GetBillingAccountSpendDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBillingAccountSpendDetailsParams() *GetBillingAccountSpendDetailsParams {
	return &GetBillingAccountSpendDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBillingAccountSpendDetailsParamsWithTimeout creates a new GetBillingAccountSpendDetailsParams object
// with the ability to set a timeout on a request.
func NewGetBillingAccountSpendDetailsParamsWithTimeout(timeout time.Duration) *GetBillingAccountSpendDetailsParams {
	return &GetBillingAccountSpendDetailsParams{
		timeout: timeout,
	}
}

// NewGetBillingAccountSpendDetailsParamsWithContext creates a new GetBillingAccountSpendDetailsParams object
// with the ability to set a context for a request.
func NewGetBillingAccountSpendDetailsParamsWithContext(ctx context.Context) *GetBillingAccountSpendDetailsParams {
	return &GetBillingAccountSpendDetailsParams{
		Context: ctx,
	}
}

// NewGetBillingAccountSpendDetailsParamsWithHTTPClient creates a new GetBillingAccountSpendDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBillingAccountSpendDetailsParamsWithHTTPClient(client *http.Client) *GetBillingAccountSpendDetailsParams {
	return &GetBillingAccountSpendDetailsParams{
		HTTPClient: client,
	}
}

/*
GetBillingAccountSpendDetailsParams contains all the parameters to send to the API endpoint

	for the get billing account spend details operation.

	Typically these are written to a http.Request.
*/
type GetBillingAccountSpendDetailsParams struct {

	// BillingAccountID.
	BillingAccountID string

	// EndTime.
	//
	// Format: date-time
	EndTime strfmt.DateTime

	// StartTime.
	//
	// Format: date-time
	StartTime strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get billing account spend details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBillingAccountSpendDetailsParams) WithDefaults() *GetBillingAccountSpendDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get billing account spend details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBillingAccountSpendDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get billing account spend details params
func (o *GetBillingAccountSpendDetailsParams) WithTimeout(timeout time.Duration) *GetBillingAccountSpendDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get billing account spend details params
func (o *GetBillingAccountSpendDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get billing account spend details params
func (o *GetBillingAccountSpendDetailsParams) WithContext(ctx context.Context) *GetBillingAccountSpendDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get billing account spend details params
func (o *GetBillingAccountSpendDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get billing account spend details params
func (o *GetBillingAccountSpendDetailsParams) WithHTTPClient(client *http.Client) *GetBillingAccountSpendDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get billing account spend details params
func (o *GetBillingAccountSpendDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBillingAccountID adds the billingAccountID to the get billing account spend details params
func (o *GetBillingAccountSpendDetailsParams) WithBillingAccountID(billingAccountID string) *GetBillingAccountSpendDetailsParams {
	o.SetBillingAccountID(billingAccountID)
	return o
}

// SetBillingAccountID adds the billingAccountId to the get billing account spend details params
func (o *GetBillingAccountSpendDetailsParams) SetBillingAccountID(billingAccountID string) {
	o.BillingAccountID = billingAccountID
}

// WithEndTime adds the endTime to the get billing account spend details params
func (o *GetBillingAccountSpendDetailsParams) WithEndTime(endTime strfmt.DateTime) *GetBillingAccountSpendDetailsParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the get billing account spend details params
func (o *GetBillingAccountSpendDetailsParams) SetEndTime(endTime strfmt.DateTime) {
	o.EndTime = endTime
}

// WithStartTime adds the startTime to the get billing account spend details params
func (o *GetBillingAccountSpendDetailsParams) WithStartTime(startTime strfmt.DateTime) *GetBillingAccountSpendDetailsParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the get billing account spend details params
func (o *GetBillingAccountSpendDetailsParams) SetStartTime(startTime strfmt.DateTime) {
	o.StartTime = startTime
}

// WriteToRequest writes these params to a swagger request
func (o *GetBillingAccountSpendDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param billingAccountId
	if err := r.SetPathParam("billingAccountId", o.BillingAccountID); err != nil {
		return err
	}

	// query param endTime
	qrEndTime := o.EndTime
	qEndTime := qrEndTime.String()
	if qEndTime != "" {

		if err := r.SetQueryParam("endTime", qEndTime); err != nil {
			return err
		}
	}

	// query param startTime
	qrStartTime := o.StartTime
	qStartTime := qrStartTime.String()
	if qStartTime != "" {

		if err := r.SetQueryParam("startTime", qStartTime); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
