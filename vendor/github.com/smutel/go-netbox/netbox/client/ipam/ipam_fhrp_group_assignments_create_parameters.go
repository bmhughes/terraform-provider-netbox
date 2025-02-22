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

	"github.com/smutel/go-netbox/netbox/models"
)

// NewIpamFhrpGroupAssignmentsCreateParams creates a new IpamFhrpGroupAssignmentsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamFhrpGroupAssignmentsCreateParams() *IpamFhrpGroupAssignmentsCreateParams {
	return &IpamFhrpGroupAssignmentsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamFhrpGroupAssignmentsCreateParamsWithTimeout creates a new IpamFhrpGroupAssignmentsCreateParams object
// with the ability to set a timeout on a request.
func NewIpamFhrpGroupAssignmentsCreateParamsWithTimeout(timeout time.Duration) *IpamFhrpGroupAssignmentsCreateParams {
	return &IpamFhrpGroupAssignmentsCreateParams{
		timeout: timeout,
	}
}

// NewIpamFhrpGroupAssignmentsCreateParamsWithContext creates a new IpamFhrpGroupAssignmentsCreateParams object
// with the ability to set a context for a request.
func NewIpamFhrpGroupAssignmentsCreateParamsWithContext(ctx context.Context) *IpamFhrpGroupAssignmentsCreateParams {
	return &IpamFhrpGroupAssignmentsCreateParams{
		Context: ctx,
	}
}

// NewIpamFhrpGroupAssignmentsCreateParamsWithHTTPClient creates a new IpamFhrpGroupAssignmentsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamFhrpGroupAssignmentsCreateParamsWithHTTPClient(client *http.Client) *IpamFhrpGroupAssignmentsCreateParams {
	return &IpamFhrpGroupAssignmentsCreateParams{
		HTTPClient: client,
	}
}

/* IpamFhrpGroupAssignmentsCreateParams contains all the parameters to send to the API endpoint
   for the ipam fhrp group assignments create operation.

   Typically these are written to a http.Request.
*/
type IpamFhrpGroupAssignmentsCreateParams struct {

	// Data.
	Data *models.WritableFHRPGroupAssignment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam fhrp group assignments create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamFhrpGroupAssignmentsCreateParams) WithDefaults() *IpamFhrpGroupAssignmentsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam fhrp group assignments create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamFhrpGroupAssignmentsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam fhrp group assignments create params
func (o *IpamFhrpGroupAssignmentsCreateParams) WithTimeout(timeout time.Duration) *IpamFhrpGroupAssignmentsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam fhrp group assignments create params
func (o *IpamFhrpGroupAssignmentsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam fhrp group assignments create params
func (o *IpamFhrpGroupAssignmentsCreateParams) WithContext(ctx context.Context) *IpamFhrpGroupAssignmentsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam fhrp group assignments create params
func (o *IpamFhrpGroupAssignmentsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam fhrp group assignments create params
func (o *IpamFhrpGroupAssignmentsCreateParams) WithHTTPClient(client *http.Client) *IpamFhrpGroupAssignmentsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam fhrp group assignments create params
func (o *IpamFhrpGroupAssignmentsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam fhrp group assignments create params
func (o *IpamFhrpGroupAssignmentsCreateParams) WithData(data *models.WritableFHRPGroupAssignment) *IpamFhrpGroupAssignmentsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam fhrp group assignments create params
func (o *IpamFhrpGroupAssignmentsCreateParams) SetData(data *models.WritableFHRPGroupAssignment) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamFhrpGroupAssignmentsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
