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
	"github.com/go-openapi/swag"

	"github.com/smutel/go-netbox/netbox/models"
)

// NewDcimSiteGroupsPartialUpdateParams creates a new DcimSiteGroupsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimSiteGroupsPartialUpdateParams() *DcimSiteGroupsPartialUpdateParams {
	return &DcimSiteGroupsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimSiteGroupsPartialUpdateParamsWithTimeout creates a new DcimSiteGroupsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimSiteGroupsPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimSiteGroupsPartialUpdateParams {
	return &DcimSiteGroupsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimSiteGroupsPartialUpdateParamsWithContext creates a new DcimSiteGroupsPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimSiteGroupsPartialUpdateParamsWithContext(ctx context.Context) *DcimSiteGroupsPartialUpdateParams {
	return &DcimSiteGroupsPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimSiteGroupsPartialUpdateParamsWithHTTPClient creates a new DcimSiteGroupsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimSiteGroupsPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimSiteGroupsPartialUpdateParams {
	return &DcimSiteGroupsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimSiteGroupsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim site groups partial update operation.

   Typically these are written to a http.Request.
*/
type DcimSiteGroupsPartialUpdateParams struct {

	// Data.
	Data *models.WritableSiteGroup

	/* ID.

	   A unique integer value identifying this site group.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim site groups partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSiteGroupsPartialUpdateParams) WithDefaults() *DcimSiteGroupsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim site groups partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSiteGroupsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim site groups partial update params
func (o *DcimSiteGroupsPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimSiteGroupsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim site groups partial update params
func (o *DcimSiteGroupsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim site groups partial update params
func (o *DcimSiteGroupsPartialUpdateParams) WithContext(ctx context.Context) *DcimSiteGroupsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim site groups partial update params
func (o *DcimSiteGroupsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim site groups partial update params
func (o *DcimSiteGroupsPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimSiteGroupsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim site groups partial update params
func (o *DcimSiteGroupsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim site groups partial update params
func (o *DcimSiteGroupsPartialUpdateParams) WithData(data *models.WritableSiteGroup) *DcimSiteGroupsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim site groups partial update params
func (o *DcimSiteGroupsPartialUpdateParams) SetData(data *models.WritableSiteGroup) {
	o.Data = data
}

// WithID adds the id to the dcim site groups partial update params
func (o *DcimSiteGroupsPartialUpdateParams) WithID(id int64) *DcimSiteGroupsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim site groups partial update params
func (o *DcimSiteGroupsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimSiteGroupsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
