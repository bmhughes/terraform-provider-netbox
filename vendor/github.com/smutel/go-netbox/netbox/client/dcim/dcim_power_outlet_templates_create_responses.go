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

// DcimPowerOutletTemplatesCreateReader is a Reader for the DcimPowerOutletTemplatesCreate structure.
type DcimPowerOutletTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimPowerOutletTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerOutletTemplatesCreateCreated creates a DcimPowerOutletTemplatesCreateCreated with default headers values
func NewDcimPowerOutletTemplatesCreateCreated() *DcimPowerOutletTemplatesCreateCreated {
	return &DcimPowerOutletTemplatesCreateCreated{}
}

/* DcimPowerOutletTemplatesCreateCreated describes a response with status code 201, with default header values.

DcimPowerOutletTemplatesCreateCreated dcim power outlet templates create created
*/
type DcimPowerOutletTemplatesCreateCreated struct {
	Payload *models.PowerOutletTemplate
}

func (o *DcimPowerOutletTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/power-outlet-templates/][%d] dcimPowerOutletTemplatesCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimPowerOutletTemplatesCreateCreated) GetPayload() *models.PowerOutletTemplate {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutletTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
