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

// NewIpamIPRangesDeleteParams creates a new IpamIPRangesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamIPRangesDeleteParams() *IpamIPRangesDeleteParams {
	return &IpamIPRangesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamIPRangesDeleteParamsWithTimeout creates a new IpamIPRangesDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamIPRangesDeleteParamsWithTimeout(timeout time.Duration) *IpamIPRangesDeleteParams {
	return &IpamIPRangesDeleteParams{
		timeout: timeout,
	}
}

// NewIpamIPRangesDeleteParamsWithContext creates a new IpamIPRangesDeleteParams object
// with the ability to set a context for a request.
func NewIpamIPRangesDeleteParamsWithContext(ctx context.Context) *IpamIPRangesDeleteParams {
	return &IpamIPRangesDeleteParams{
		Context: ctx,
	}
}

// NewIpamIPRangesDeleteParamsWithHTTPClient creates a new IpamIPRangesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamIPRangesDeleteParamsWithHTTPClient(client *http.Client) *IpamIPRangesDeleteParams {
	return &IpamIPRangesDeleteParams{
		HTTPClient: client,
	}
}

/* IpamIPRangesDeleteParams contains all the parameters to send to the API endpoint
   for the ipam ip ranges delete operation.

   Typically these are written to a http.Request.
*/
type IpamIPRangesDeleteParams struct {

	/* ID.

	   A unique integer value identifying this IP range.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam ip ranges delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPRangesDeleteParams) WithDefaults() *IpamIPRangesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam ip ranges delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPRangesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam ip ranges delete params
func (o *IpamIPRangesDeleteParams) WithTimeout(timeout time.Duration) *IpamIPRangesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam ip ranges delete params
func (o *IpamIPRangesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam ip ranges delete params
func (o *IpamIPRangesDeleteParams) WithContext(ctx context.Context) *IpamIPRangesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam ip ranges delete params
func (o *IpamIPRangesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam ip ranges delete params
func (o *IpamIPRangesDeleteParams) WithHTTPClient(client *http.Client) *IpamIPRangesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam ip ranges delete params
func (o *IpamIPRangesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam ip ranges delete params
func (o *IpamIPRangesDeleteParams) WithID(id int64) *IpamIPRangesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam ip ranges delete params
func (o *IpamIPRangesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamIPRangesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
