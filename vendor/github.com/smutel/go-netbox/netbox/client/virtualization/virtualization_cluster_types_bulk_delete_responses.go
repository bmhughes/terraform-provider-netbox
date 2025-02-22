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

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VirtualizationClusterTypesBulkDeleteReader is a Reader for the VirtualizationClusterTypesBulkDelete structure.
type VirtualizationClusterTypesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterTypesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationClusterTypesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClusterTypesBulkDeleteNoContent creates a VirtualizationClusterTypesBulkDeleteNoContent with default headers values
func NewVirtualizationClusterTypesBulkDeleteNoContent() *VirtualizationClusterTypesBulkDeleteNoContent {
	return &VirtualizationClusterTypesBulkDeleteNoContent{}
}

/* VirtualizationClusterTypesBulkDeleteNoContent describes a response with status code 204, with default header values.

VirtualizationClusterTypesBulkDeleteNoContent virtualization cluster types bulk delete no content
*/
type VirtualizationClusterTypesBulkDeleteNoContent struct {
}

func (o *VirtualizationClusterTypesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/cluster-types/][%d] virtualizationClusterTypesBulkDeleteNoContent ", 204)
}

func (o *VirtualizationClusterTypesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
