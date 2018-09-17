///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package runner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// GetRunReader is a Reader for the GetRun structure.
type GetRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRunBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetRunUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRunOK creates a GetRunOK with default headers values
func NewGetRunOK() *GetRunOK {
	return &GetRunOK{}
}

/*GetRunOK handles this case with default header values.

Function Run
*/
type GetRunOK struct {
	Payload *v1.Run
}

func (o *GetRunOK) Error() string {
	return fmt.Sprintf("[GET /runs/{runName}][%d] getRunOK  %+v", 200, o.Payload)
}

func (o *GetRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Run)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunBadRequest creates a GetRunBadRequest with default headers values
func NewGetRunBadRequest() *GetRunBadRequest {
	return &GetRunBadRequest{}
}

/*GetRunBadRequest handles this case with default header values.

Bad Request
*/
type GetRunBadRequest struct {
	Payload *v1.Error
}

func (o *GetRunBadRequest) Error() string {
	return fmt.Sprintf("[GET /runs/{runName}][%d] getRunBadRequest  %+v", 400, o.Payload)
}

func (o *GetRunBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunUnauthorized creates a GetRunUnauthorized with default headers values
func NewGetRunUnauthorized() *GetRunUnauthorized {
	return &GetRunUnauthorized{}
}

/*GetRunUnauthorized handles this case with default header values.

Unauthorized Request
*/
type GetRunUnauthorized struct {
	Payload *v1.Error
}

func (o *GetRunUnauthorized) Error() string {
	return fmt.Sprintf("[GET /runs/{runName}][%d] getRunUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRunUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunForbidden creates a GetRunForbidden with default headers values
func NewGetRunForbidden() *GetRunForbidden {
	return &GetRunForbidden{}
}

/*GetRunForbidden handles this case with default header values.

access to this resource is forbidden
*/
type GetRunForbidden struct {
	Payload *v1.Error
}

func (o *GetRunForbidden) Error() string {
	return fmt.Sprintf("[GET /runs/{runName}][%d] getRunForbidden  %+v", 403, o.Payload)
}

func (o *GetRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunNotFound creates a GetRunNotFound with default headers values
func NewGetRunNotFound() *GetRunNotFound {
	return &GetRunNotFound{}
}

/*GetRunNotFound handles this case with default header values.

Function or Run not found
*/
type GetRunNotFound struct {
	Payload *v1.Error
}

func (o *GetRunNotFound) Error() string {
	return fmt.Sprintf("[GET /runs/{runName}][%d] getRunNotFound  %+v", 404, o.Payload)
}

func (o *GetRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunDefault creates a GetRunDefault with default headers values
func NewGetRunDefault(code int) *GetRunDefault {
	return &GetRunDefault{
		_statusCode: code,
	}
}

/*GetRunDefault handles this case with default header values.

Unknown error
*/
type GetRunDefault struct {
	_statusCode int

	Payload *v1.Error
}

// Code gets the status code for the get run default response
func (o *GetRunDefault) Code() int {
	return o._statusCode
}

func (o *GetRunDefault) Error() string {
	return fmt.Sprintf("[GET /runs/{runName}][%d] getRun default  %+v", o._statusCode, o.Payload)
}

func (o *GetRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}