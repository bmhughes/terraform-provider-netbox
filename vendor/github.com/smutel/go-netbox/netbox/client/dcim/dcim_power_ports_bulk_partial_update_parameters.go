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

package dcim

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

// NewDcimPowerPortsBulkPartialUpdateParams creates a new DcimPowerPortsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerPortsBulkPartialUpdateParams() *DcimPowerPortsBulkPartialUpdateParams {
	return &DcimPowerPortsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortsBulkPartialUpdateParamsWithTimeout creates a new DcimPowerPortsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimPowerPortsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimPowerPortsBulkPartialUpdateParams {
	return &DcimPowerPortsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimPowerPortsBulkPartialUpdateParamsWithContext creates a new DcimPowerPortsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimPowerPortsBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimPowerPortsBulkPartialUpdateParams {
	return &DcimPowerPortsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimPowerPortsBulkPartialUpdateParamsWithHTTPClient creates a new DcimPowerPortsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerPortsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimPowerPortsBulkPartialUpdateParams {
	return &DcimPowerPortsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimPowerPortsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim power ports bulk partial update operation.

   Typically these are written to a http.Request.
*/
type DcimPowerPortsBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritablePowerPort

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power ports bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortsBulkPartialUpdateParams) WithDefaults() *DcimPowerPortsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power ports bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power ports bulk partial update params
func (o *DcimPowerPortsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimPowerPortsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power ports bulk partial update params
func (o *DcimPowerPortsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power ports bulk partial update params
func (o *DcimPowerPortsBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimPowerPortsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power ports bulk partial update params
func (o *DcimPowerPortsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power ports bulk partial update params
func (o *DcimPowerPortsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimPowerPortsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power ports bulk partial update params
func (o *DcimPowerPortsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power ports bulk partial update params
func (o *DcimPowerPortsBulkPartialUpdateParams) WithData(data *models.WritablePowerPort) *DcimPowerPortsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power ports bulk partial update params
func (o *DcimPowerPortsBulkPartialUpdateParams) SetData(data *models.WritablePowerPort) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
