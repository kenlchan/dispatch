///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// DeleteImageByNameReader is a Reader for the DeleteImageByName structure.
type DeleteImageByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteImageByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteImageByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteImageByNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteImageByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteImageByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteImageByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteImageByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteImageByNameOK creates a DeleteImageByNameOK with default headers values
func NewDeleteImageByNameOK() *DeleteImageByNameOK {
	return &DeleteImageByNameOK{}
}

/*DeleteImageByNameOK handles this case with default header values.

successful operation
*/
type DeleteImageByNameOK struct {
	Payload *v1.Image
}

func (o *DeleteImageByNameOK) Error() string {
	return fmt.Sprintf("[DELETE /image/{imageName}][%d] deleteImageByNameOK  %+v", 200, o.Payload)
}

func (o *DeleteImageByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Image)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImageByNameBadRequest creates a DeleteImageByNameBadRequest with default headers values
func NewDeleteImageByNameBadRequest() *DeleteImageByNameBadRequest {
	return &DeleteImageByNameBadRequest{}
}

/*DeleteImageByNameBadRequest handles this case with default header values.

Invalid ID supplied
*/
type DeleteImageByNameBadRequest struct {
	Payload *v1.Error
}

func (o *DeleteImageByNameBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /image/{imageName}][%d] deleteImageByNameBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteImageByNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImageByNameUnauthorized creates a DeleteImageByNameUnauthorized with default headers values
func NewDeleteImageByNameUnauthorized() *DeleteImageByNameUnauthorized {
	return &DeleteImageByNameUnauthorized{}
}

/*DeleteImageByNameUnauthorized handles this case with default header values.

Unauthorized Request
*/
type DeleteImageByNameUnauthorized struct {
	Payload *v1.Error
}

func (o *DeleteImageByNameUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /image/{imageName}][%d] deleteImageByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteImageByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImageByNameForbidden creates a DeleteImageByNameForbidden with default headers values
func NewDeleteImageByNameForbidden() *DeleteImageByNameForbidden {
	return &DeleteImageByNameForbidden{}
}

/*DeleteImageByNameForbidden handles this case with default header values.

access to this resource is forbidden
*/
type DeleteImageByNameForbidden struct {
	Payload *v1.Error
}

func (o *DeleteImageByNameForbidden) Error() string {
	return fmt.Sprintf("[DELETE /image/{imageName}][%d] deleteImageByNameForbidden  %+v", 403, o.Payload)
}

func (o *DeleteImageByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImageByNameNotFound creates a DeleteImageByNameNotFound with default headers values
func NewDeleteImageByNameNotFound() *DeleteImageByNameNotFound {
	return &DeleteImageByNameNotFound{}
}

/*DeleteImageByNameNotFound handles this case with default header values.

Image not found
*/
type DeleteImageByNameNotFound struct {
	Payload *v1.Error
}

func (o *DeleteImageByNameNotFound) Error() string {
	return fmt.Sprintf("[DELETE /image/{imageName}][%d] deleteImageByNameNotFound  %+v", 404, o.Payload)
}

func (o *DeleteImageByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImageByNameDefault creates a DeleteImageByNameDefault with default headers values
func NewDeleteImageByNameDefault(code int) *DeleteImageByNameDefault {
	return &DeleteImageByNameDefault{
		_statusCode: code,
	}
}

/*DeleteImageByNameDefault handles this case with default header values.

Generic error response
*/
type DeleteImageByNameDefault struct {
	_statusCode int

	Payload *v1.Error
}

// Code gets the status code for the delete image by name default response
func (o *DeleteImageByNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteImageByNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /image/{imageName}][%d] deleteImageByName default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteImageByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}