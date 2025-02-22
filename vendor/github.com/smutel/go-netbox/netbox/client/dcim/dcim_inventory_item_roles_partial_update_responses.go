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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smutel/go-netbox/netbox/models"
)

// DcimInventoryItemRolesPartialUpdateReader is a Reader for the DcimInventoryItemRolesPartialUpdate structure.
type DcimInventoryItemRolesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemRolesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemRolesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInventoryItemRolesPartialUpdateOK creates a DcimInventoryItemRolesPartialUpdateOK with default headers values
func NewDcimInventoryItemRolesPartialUpdateOK() *DcimInventoryItemRolesPartialUpdateOK {
	return &DcimInventoryItemRolesPartialUpdateOK{}
}

/* DcimInventoryItemRolesPartialUpdateOK describes a response with status code 200, with default header values.

DcimInventoryItemRolesPartialUpdateOK dcim inventory item roles partial update o k
*/
type DcimInventoryItemRolesPartialUpdateOK struct {
	Payload *models.InventoryItemRole
}

func (o *DcimInventoryItemRolesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-roles/{id}/][%d] dcimInventoryItemRolesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimInventoryItemRolesPartialUpdateOK) GetPayload() *models.InventoryItemRole {
	return o.Payload
}

func (o *DcimInventoryItemRolesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItemRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
