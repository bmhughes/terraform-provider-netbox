// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2020 Samuel Mutel <12967891+smutel@users.noreply.github.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
//

package extras

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

	"github.com/smutel/go-netbox/netbox/models"
)

// NewExtrasJournalEntriesBulkPartialUpdateParams creates a new ExtrasJournalEntriesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasJournalEntriesBulkPartialUpdateParams() *ExtrasJournalEntriesBulkPartialUpdateParams {
	return &ExtrasJournalEntriesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasJournalEntriesBulkPartialUpdateParamsWithTimeout creates a new ExtrasJournalEntriesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasJournalEntriesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasJournalEntriesBulkPartialUpdateParams {
	return &ExtrasJournalEntriesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasJournalEntriesBulkPartialUpdateParamsWithContext creates a new ExtrasJournalEntriesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasJournalEntriesBulkPartialUpdateParamsWithContext(ctx context.Context) *ExtrasJournalEntriesBulkPartialUpdateParams {
	return &ExtrasJournalEntriesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasJournalEntriesBulkPartialUpdateParamsWithHTTPClient creates a new ExtrasJournalEntriesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasJournalEntriesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasJournalEntriesBulkPartialUpdateParams {
	return &ExtrasJournalEntriesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasJournalEntriesBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras journal entries bulk partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasJournalEntriesBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableJournalEntry

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras journal entries bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJournalEntriesBulkPartialUpdateParams) WithDefaults() *ExtrasJournalEntriesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras journal entries bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJournalEntriesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras journal entries bulk partial update params
func (o *ExtrasJournalEntriesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasJournalEntriesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras journal entries bulk partial update params
func (o *ExtrasJournalEntriesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras journal entries bulk partial update params
func (o *ExtrasJournalEntriesBulkPartialUpdateParams) WithContext(ctx context.Context) *ExtrasJournalEntriesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras journal entries bulk partial update params
func (o *ExtrasJournalEntriesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras journal entries bulk partial update params
func (o *ExtrasJournalEntriesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasJournalEntriesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras journal entries bulk partial update params
func (o *ExtrasJournalEntriesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras journal entries bulk partial update params
func (o *ExtrasJournalEntriesBulkPartialUpdateParams) WithData(data *models.WritableJournalEntry) *ExtrasJournalEntriesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras journal entries bulk partial update params
func (o *ExtrasJournalEntriesBulkPartialUpdateParams) SetData(data *models.WritableJournalEntry) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasJournalEntriesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
