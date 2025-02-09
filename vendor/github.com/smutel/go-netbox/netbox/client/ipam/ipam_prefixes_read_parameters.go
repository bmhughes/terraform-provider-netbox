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

package ipam

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

// NewIpamPrefixesReadParams creates a new IpamPrefixesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamPrefixesReadParams() *IpamPrefixesReadParams {
	return &IpamPrefixesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamPrefixesReadParamsWithTimeout creates a new IpamPrefixesReadParams object
// with the ability to set a timeout on a request.
func NewIpamPrefixesReadParamsWithTimeout(timeout time.Duration) *IpamPrefixesReadParams {
	return &IpamPrefixesReadParams{
		timeout: timeout,
	}
}

// NewIpamPrefixesReadParamsWithContext creates a new IpamPrefixesReadParams object
// with the ability to set a context for a request.
func NewIpamPrefixesReadParamsWithContext(ctx context.Context) *IpamPrefixesReadParams {
	return &IpamPrefixesReadParams{
		Context: ctx,
	}
}

// NewIpamPrefixesReadParamsWithHTTPClient creates a new IpamPrefixesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamPrefixesReadParamsWithHTTPClient(client *http.Client) *IpamPrefixesReadParams {
	return &IpamPrefixesReadParams{
		HTTPClient: client,
	}
}

/* IpamPrefixesReadParams contains all the parameters to send to the API endpoint
   for the ipam prefixes read operation.

   Typically these are written to a http.Request.
*/
type IpamPrefixesReadParams struct {

	/* ID.

	   A unique integer value identifying this prefix.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam prefixes read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesReadParams) WithDefaults() *IpamPrefixesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam prefixes read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam prefixes read params
func (o *IpamPrefixesReadParams) WithTimeout(timeout time.Duration) *IpamPrefixesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes read params
func (o *IpamPrefixesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes read params
func (o *IpamPrefixesReadParams) WithContext(ctx context.Context) *IpamPrefixesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes read params
func (o *IpamPrefixesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes read params
func (o *IpamPrefixesReadParams) WithHTTPClient(client *http.Client) *IpamPrefixesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes read params
func (o *IpamPrefixesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam prefixes read params
func (o *IpamPrefixesReadParams) WithID(id int64) *IpamPrefixesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam prefixes read params
func (o *IpamPrefixesReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamPrefixesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
