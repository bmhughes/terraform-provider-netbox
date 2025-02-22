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

package circuits

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
	"github.com/go-openapi/swag"
)

// NewCircuitsCircuitTypesReadParams creates a new CircuitsCircuitTypesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsCircuitTypesReadParams() *CircuitsCircuitTypesReadParams {
	return &CircuitsCircuitTypesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTypesReadParamsWithTimeout creates a new CircuitsCircuitTypesReadParams object
// with the ability to set a timeout on a request.
func NewCircuitsCircuitTypesReadParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTypesReadParams {
	return &CircuitsCircuitTypesReadParams{
		timeout: timeout,
	}
}

// NewCircuitsCircuitTypesReadParamsWithContext creates a new CircuitsCircuitTypesReadParams object
// with the ability to set a context for a request.
func NewCircuitsCircuitTypesReadParamsWithContext(ctx context.Context) *CircuitsCircuitTypesReadParams {
	return &CircuitsCircuitTypesReadParams{
		Context: ctx,
	}
}

// NewCircuitsCircuitTypesReadParamsWithHTTPClient creates a new CircuitsCircuitTypesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsCircuitTypesReadParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTypesReadParams {
	return &CircuitsCircuitTypesReadParams{
		HTTPClient: client,
	}
}

/* CircuitsCircuitTypesReadParams contains all the parameters to send to the API endpoint
   for the circuits circuit types read operation.

   Typically these are written to a http.Request.
*/
type CircuitsCircuitTypesReadParams struct {

	/* ID.

	   A unique integer value identifying this circuit type.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits circuit types read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTypesReadParams) WithDefaults() *CircuitsCircuitTypesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits circuit types read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTypesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits circuit types read params
func (o *CircuitsCircuitTypesReadParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTypesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit types read params
func (o *CircuitsCircuitTypesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit types read params
func (o *CircuitsCircuitTypesReadParams) WithContext(ctx context.Context) *CircuitsCircuitTypesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit types read params
func (o *CircuitsCircuitTypesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit types read params
func (o *CircuitsCircuitTypesReadParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTypesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit types read params
func (o *CircuitsCircuitTypesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the circuits circuit types read params
func (o *CircuitsCircuitTypesReadParams) WithID(id int64) *CircuitsCircuitTypesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits circuit types read params
func (o *CircuitsCircuitTypesReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTypesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
