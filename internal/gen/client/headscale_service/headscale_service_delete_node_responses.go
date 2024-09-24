// Code generated by go-swagger; DO NOT EDIT.

package headscale_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/awlsring/terraform-provider-headscale/internal/gen/models"
)

// HeadscaleServiceDeleteNodeReader is a Reader for the HeadscaleServiceDeleteNode structure.
type HeadscaleServiceDeleteNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceDeleteNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceDeleteNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceDeleteNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceDeleteNodeOK creates a HeadscaleServiceDeleteNodeOK with default headers values
func NewHeadscaleServiceDeleteNodeOK() *HeadscaleServiceDeleteNodeOK {
	return &HeadscaleServiceDeleteNodeOK{}
}

/*
HeadscaleServiceDeleteNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceDeleteNodeOK struct {
	Payload models.V1DeleteNodeResponse
}

// IsSuccess returns true when this headscale service delete node o k response has a 2xx status code
func (o *HeadscaleServiceDeleteNodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service delete node o k response has a 3xx status code
func (o *HeadscaleServiceDeleteNodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service delete node o k response has a 4xx status code
func (o *HeadscaleServiceDeleteNodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service delete node o k response has a 5xx status code
func (o *HeadscaleServiceDeleteNodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service delete node o k response a status code equal to that given
func (o *HeadscaleServiceDeleteNodeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the headscale service delete node o k response
func (o *HeadscaleServiceDeleteNodeOK) Code() int {
	return 200
}

func (o *HeadscaleServiceDeleteNodeOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1/node/{nodeId}][%d] headscaleServiceDeleteNodeOK %s", 200, payload)
}

func (o *HeadscaleServiceDeleteNodeOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1/node/{nodeId}][%d] headscaleServiceDeleteNodeOK %s", 200, payload)
}

func (o *HeadscaleServiceDeleteNodeOK) GetPayload() models.V1DeleteNodeResponse {
	return o.Payload
}

func (o *HeadscaleServiceDeleteNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceDeleteNodeDefault creates a HeadscaleServiceDeleteNodeDefault with default headers values
func NewHeadscaleServiceDeleteNodeDefault(code int) *HeadscaleServiceDeleteNodeDefault {
	return &HeadscaleServiceDeleteNodeDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceDeleteNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceDeleteNodeDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this headscale service delete node default response has a 2xx status code
func (o *HeadscaleServiceDeleteNodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service delete node default response has a 3xx status code
func (o *HeadscaleServiceDeleteNodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service delete node default response has a 4xx status code
func (o *HeadscaleServiceDeleteNodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service delete node default response has a 5xx status code
func (o *HeadscaleServiceDeleteNodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service delete node default response a status code equal to that given
func (o *HeadscaleServiceDeleteNodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the headscale service delete node default response
func (o *HeadscaleServiceDeleteNodeDefault) Code() int {
	return o._statusCode
}

func (o *HeadscaleServiceDeleteNodeDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1/node/{nodeId}][%d] HeadscaleService_DeleteNode default %s", o._statusCode, payload)
}

func (o *HeadscaleServiceDeleteNodeDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1/node/{nodeId}][%d] HeadscaleService_DeleteNode default %s", o._statusCode, payload)
}

func (o *HeadscaleServiceDeleteNodeDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceDeleteNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}