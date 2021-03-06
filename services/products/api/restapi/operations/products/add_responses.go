// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"eventsourcing/services/products/api/models"
)

// AddCreatedCode is the HTTP code returned for type AddCreated
const AddCreatedCode int = 201

/*AddCreated Created

swagger:response addCreated
*/
type AddCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Product `json:"body,omitempty"`
}

// NewAddCreated creates AddCreated with default headers values
func NewAddCreated() *AddCreated {

	return &AddCreated{}
}

// WithPayload adds the payload to the add created response
func (o *AddCreated) WithPayload(payload *models.Product) *AddCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add created response
func (o *AddCreated) SetPayload(payload *models.Product) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddBadRequestCode is the HTTP code returned for type AddBadRequest
const AddBadRequestCode int = 400

/*AddBadRequest BadRequest

swagger:response addBadRequest
*/
type AddBadRequest struct {
}

// NewAddBadRequest creates AddBadRequest with default headers values
func NewAddBadRequest() *AddBadRequest {

	return &AddBadRequest{}
}

// WriteResponse to the client
func (o *AddBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// AddInternalServerErrorCode is the HTTP code returned for type AddInternalServerError
const AddInternalServerErrorCode int = 500

/*AddInternalServerError InternalServerError

swagger:response addInternalServerError
*/
type AddInternalServerError struct {
}

// NewAddInternalServerError creates AddInternalServerError with default headers values
func NewAddInternalServerError() *AddInternalServerError {

	return &AddInternalServerError{}
}

// WriteResponse to the client
func (o *AddInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
