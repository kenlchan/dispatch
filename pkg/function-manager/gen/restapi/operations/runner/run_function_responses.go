///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package runner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.eng.vmware.com/serverless/serverless/pkg/function-manager/gen/models"
)

// RunFunctionOKCode is the HTTP code returned for type RunFunctionOK
const RunFunctionOKCode int = 200

/*RunFunctionOK Successful execution (blocking call)

swagger:response runFunctionOK
*/
type RunFunctionOK struct {

	/*
	  In: Body
	*/
	Payload *models.Run `json:"body,omitempty"`
}

// NewRunFunctionOK creates RunFunctionOK with default headers values
func NewRunFunctionOK() *RunFunctionOK {
	return &RunFunctionOK{}
}

// WithPayload adds the payload to the run function o k response
func (o *RunFunctionOK) WithPayload(payload *models.Run) *RunFunctionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the run function o k response
func (o *RunFunctionOK) SetPayload(payload *models.Run) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RunFunctionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RunFunctionAcceptedCode is the HTTP code returned for type RunFunctionAccepted
const RunFunctionAcceptedCode int = 202

/*RunFunctionAccepted Execution started (non-blocking call)

swagger:response runFunctionAccepted
*/
type RunFunctionAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.Run `json:"body,omitempty"`
}

// NewRunFunctionAccepted creates RunFunctionAccepted with default headers values
func NewRunFunctionAccepted() *RunFunctionAccepted {
	return &RunFunctionAccepted{}
}

// WithPayload adds the payload to the run function accepted response
func (o *RunFunctionAccepted) WithPayload(payload *models.Run) *RunFunctionAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the run function accepted response
func (o *RunFunctionAccepted) SetPayload(payload *models.Run) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RunFunctionAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RunFunctionBadRequestCode is the HTTP code returned for type RunFunctionBadRequest
const RunFunctionBadRequestCode int = 400

/*RunFunctionBadRequest User error

swagger:response runFunctionBadRequest
*/
type RunFunctionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRunFunctionBadRequest creates RunFunctionBadRequest with default headers values
func NewRunFunctionBadRequest() *RunFunctionBadRequest {
	return &RunFunctionBadRequest{}
}

// WithPayload adds the payload to the run function bad request response
func (o *RunFunctionBadRequest) WithPayload(payload *models.Error) *RunFunctionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the run function bad request response
func (o *RunFunctionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RunFunctionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RunFunctionNotFoundCode is the HTTP code returned for type RunFunctionNotFound
const RunFunctionNotFoundCode int = 404

/*RunFunctionNotFound Function not found

swagger:response runFunctionNotFound
*/
type RunFunctionNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRunFunctionNotFound creates RunFunctionNotFound with default headers values
func NewRunFunctionNotFound() *RunFunctionNotFound {
	return &RunFunctionNotFound{}
}

// WithPayload adds the payload to the run function not found response
func (o *RunFunctionNotFound) WithPayload(payload *models.Error) *RunFunctionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the run function not found response
func (o *RunFunctionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RunFunctionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RunFunctionUnprocessableEntityCode is the HTTP code returned for type RunFunctionUnprocessableEntity
const RunFunctionUnprocessableEntityCode int = 422

/*RunFunctionUnprocessableEntity Input object validation failed

swagger:response runFunctionUnprocessableEntity
*/
type RunFunctionUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRunFunctionUnprocessableEntity creates RunFunctionUnprocessableEntity with default headers values
func NewRunFunctionUnprocessableEntity() *RunFunctionUnprocessableEntity {
	return &RunFunctionUnprocessableEntity{}
}

// WithPayload adds the payload to the run function unprocessable entity response
func (o *RunFunctionUnprocessableEntity) WithPayload(payload *models.Error) *RunFunctionUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the run function unprocessable entity response
func (o *RunFunctionUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RunFunctionUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RunFunctionInternalServerErrorCode is the HTTP code returned for type RunFunctionInternalServerError
const RunFunctionInternalServerErrorCode int = 500

/*RunFunctionInternalServerError Internal error

swagger:response runFunctionInternalServerError
*/
type RunFunctionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRunFunctionInternalServerError creates RunFunctionInternalServerError with default headers values
func NewRunFunctionInternalServerError() *RunFunctionInternalServerError {
	return &RunFunctionInternalServerError{}
}

// WithPayload adds the payload to the run function internal server error response
func (o *RunFunctionInternalServerError) WithPayload(payload *models.Error) *RunFunctionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the run function internal server error response
func (o *RunFunctionInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RunFunctionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RunFunctionBadGatewayCode is the HTTP code returned for type RunFunctionBadGateway
const RunFunctionBadGatewayCode int = 502

/*RunFunctionBadGateway Function error occurred (blocking call)

swagger:response runFunctionBadGateway
*/
type RunFunctionBadGateway struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRunFunctionBadGateway creates RunFunctionBadGateway with default headers values
func NewRunFunctionBadGateway() *RunFunctionBadGateway {
	return &RunFunctionBadGateway{}
}

// WithPayload adds the payload to the run function bad gateway response
func (o *RunFunctionBadGateway) WithPayload(payload *models.Error) *RunFunctionBadGateway {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the run function bad gateway response
func (o *RunFunctionBadGateway) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RunFunctionBadGateway) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(502)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
