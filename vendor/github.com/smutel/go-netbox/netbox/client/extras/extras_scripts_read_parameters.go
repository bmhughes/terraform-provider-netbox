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

// NewExtrasScriptsReadParams creates a new ExtrasScriptsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasScriptsReadParams() *ExtrasScriptsReadParams {
	return &ExtrasScriptsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasScriptsReadParamsWithTimeout creates a new ExtrasScriptsReadParams object
// with the ability to set a timeout on a request.
func NewExtrasScriptsReadParamsWithTimeout(timeout time.Duration) *ExtrasScriptsReadParams {
	return &ExtrasScriptsReadParams{
		timeout: timeout,
	}
}

// NewExtrasScriptsReadParamsWithContext creates a new ExtrasScriptsReadParams object
// with the ability to set a context for a request.
func NewExtrasScriptsReadParamsWithContext(ctx context.Context) *ExtrasScriptsReadParams {
	return &ExtrasScriptsReadParams{
		Context: ctx,
	}
}

// NewExtrasScriptsReadParamsWithHTTPClient creates a new ExtrasScriptsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasScriptsReadParamsWithHTTPClient(client *http.Client) *ExtrasScriptsReadParams {
	return &ExtrasScriptsReadParams{
		HTTPClient: client,
	}
}

/* ExtrasScriptsReadParams contains all the parameters to send to the API endpoint
   for the extras scripts read operation.

   Typically these are written to a http.Request.
*/
type ExtrasScriptsReadParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras scripts read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasScriptsReadParams) WithDefaults() *ExtrasScriptsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras scripts read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasScriptsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras scripts read params
func (o *ExtrasScriptsReadParams) WithTimeout(timeout time.Duration) *ExtrasScriptsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras scripts read params
func (o *ExtrasScriptsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras scripts read params
func (o *ExtrasScriptsReadParams) WithContext(ctx context.Context) *ExtrasScriptsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras scripts read params
func (o *ExtrasScriptsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras scripts read params
func (o *ExtrasScriptsReadParams) WithHTTPClient(client *http.Client) *ExtrasScriptsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras scripts read params
func (o *ExtrasScriptsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras scripts read params
func (o *ExtrasScriptsReadParams) WithID(id string) *ExtrasScriptsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras scripts read params
func (o *ExtrasScriptsReadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasScriptsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
