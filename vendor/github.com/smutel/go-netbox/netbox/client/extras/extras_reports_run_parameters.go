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
)

// NewExtrasReportsRunParams creates a new ExtrasReportsRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasReportsRunParams() *ExtrasReportsRunParams {
	return &ExtrasReportsRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasReportsRunParamsWithTimeout creates a new ExtrasReportsRunParams object
// with the ability to set a timeout on a request.
func NewExtrasReportsRunParamsWithTimeout(timeout time.Duration) *ExtrasReportsRunParams {
	return &ExtrasReportsRunParams{
		timeout: timeout,
	}
}

// NewExtrasReportsRunParamsWithContext creates a new ExtrasReportsRunParams object
// with the ability to set a context for a request.
func NewExtrasReportsRunParamsWithContext(ctx context.Context) *ExtrasReportsRunParams {
	return &ExtrasReportsRunParams{
		Context: ctx,
	}
}

// NewExtrasReportsRunParamsWithHTTPClient creates a new ExtrasReportsRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasReportsRunParamsWithHTTPClient(client *http.Client) *ExtrasReportsRunParams {
	return &ExtrasReportsRunParams{
		HTTPClient: client,
	}
}

/* ExtrasReportsRunParams contains all the parameters to send to the API endpoint
   for the extras reports run operation.

   Typically these are written to a http.Request.
*/
type ExtrasReportsRunParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras reports run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasReportsRunParams) WithDefaults() *ExtrasReportsRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras reports run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasReportsRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras reports run params
func (o *ExtrasReportsRunParams) WithTimeout(timeout time.Duration) *ExtrasReportsRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras reports run params
func (o *ExtrasReportsRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras reports run params
func (o *ExtrasReportsRunParams) WithContext(ctx context.Context) *ExtrasReportsRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras reports run params
func (o *ExtrasReportsRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras reports run params
func (o *ExtrasReportsRunParams) WithHTTPClient(client *http.Client) *ExtrasReportsRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras reports run params
func (o *ExtrasReportsRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras reports run params
func (o *ExtrasReportsRunParams) WithID(id string) *ExtrasReportsRunParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras reports run params
func (o *ExtrasReportsRunParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasReportsRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
