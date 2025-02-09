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

package tenancy

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

// NewTenancyTenantGroupsBulkDeleteParams creates a new TenancyTenantGroupsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyTenantGroupsBulkDeleteParams() *TenancyTenantGroupsBulkDeleteParams {
	return &TenancyTenantGroupsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantGroupsBulkDeleteParamsWithTimeout creates a new TenancyTenantGroupsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewTenancyTenantGroupsBulkDeleteParamsWithTimeout(timeout time.Duration) *TenancyTenantGroupsBulkDeleteParams {
	return &TenancyTenantGroupsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewTenancyTenantGroupsBulkDeleteParamsWithContext creates a new TenancyTenantGroupsBulkDeleteParams object
// with the ability to set a context for a request.
func NewTenancyTenantGroupsBulkDeleteParamsWithContext(ctx context.Context) *TenancyTenantGroupsBulkDeleteParams {
	return &TenancyTenantGroupsBulkDeleteParams{
		Context: ctx,
	}
}

// NewTenancyTenantGroupsBulkDeleteParamsWithHTTPClient creates a new TenancyTenantGroupsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyTenantGroupsBulkDeleteParamsWithHTTPClient(client *http.Client) *TenancyTenantGroupsBulkDeleteParams {
	return &TenancyTenantGroupsBulkDeleteParams{
		HTTPClient: client,
	}
}

/* TenancyTenantGroupsBulkDeleteParams contains all the parameters to send to the API endpoint
   for the tenancy tenant groups bulk delete operation.

   Typically these are written to a http.Request.
*/
type TenancyTenantGroupsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy tenant groups bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantGroupsBulkDeleteParams) WithDefaults() *TenancyTenantGroupsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy tenant groups bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantGroupsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy tenant groups bulk delete params
func (o *TenancyTenantGroupsBulkDeleteParams) WithTimeout(timeout time.Duration) *TenancyTenantGroupsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenant groups bulk delete params
func (o *TenancyTenantGroupsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenant groups bulk delete params
func (o *TenancyTenantGroupsBulkDeleteParams) WithContext(ctx context.Context) *TenancyTenantGroupsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenant groups bulk delete params
func (o *TenancyTenantGroupsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenant groups bulk delete params
func (o *TenancyTenantGroupsBulkDeleteParams) WithHTTPClient(client *http.Client) *TenancyTenantGroupsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenant groups bulk delete params
func (o *TenancyTenantGroupsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantGroupsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
