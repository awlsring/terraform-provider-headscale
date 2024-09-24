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

// HeadscaleServiceExpireNodeReader is a Reader for the HeadscaleServiceExpireNode structure.
type HeadscaleServiceExpireNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceExpireNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceExpireNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceExpireNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceExpireNodeOK creates a HeadscaleServiceExpireNodeOK with default headers values
func NewHeadscaleServiceExpireNodeOK() *HeadscaleServiceExpireNodeOK {
	return &HeadscaleServiceExpireNodeOK{}
}

/*
HeadscaleServiceExpireNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceExpireNodeOK struct {
	Payload *models.V1ExpireNodeResponse
}

// IsSuccess returns true when this headscale service expire node o k response has a 2xx status code
func (o *HeadscaleServiceExpireNodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service expire node o k response has a 3xx status code
func (o *HeadscaleServiceExpireNodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service expire node o k response has a 4xx status code
func (o *HeadscaleServiceExpireNodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service expire node o k response has a 5xx status code
func (o *HeadscaleServiceExpireNodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service expire node o k response a status code equal to that given
func (o *HeadscaleServiceExpireNodeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the headscale service expire node o k response
func (o *HeadscaleServiceExpireNodeOK) Code() int {
	return 200
}

func (o *HeadscaleServiceExpireNodeOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/node/{nodeId}/expire][%d] headscaleServiceExpireNodeOK %s", 200, payload)
}

func (o *HeadscaleServiceExpireNodeOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/node/{nodeId}/expire][%d] headscaleServiceExpireNodeOK %s", 200, payload)
}

func (o *HeadscaleServiceExpireNodeOK) GetPayload() *models.V1ExpireNodeResponse {
	return o.Payload
}

func (o *HeadscaleServiceExpireNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ExpireNodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceExpireNodeDefault creates a HeadscaleServiceExpireNodeDefault with default headers values
func NewHeadscaleServiceExpireNodeDefault(code int) *HeadscaleServiceExpireNodeDefault {
	return &HeadscaleServiceExpireNodeDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceExpireNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceExpireNodeDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this headscale service expire node default response has a 2xx status code
func (o *HeadscaleServiceExpireNodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service expire node default response has a 3xx status code
func (o *HeadscaleServiceExpireNodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service expire node default response has a 4xx status code
func (o *HeadscaleServiceExpireNodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service expire node default response has a 5xx status code
func (o *HeadscaleServiceExpireNodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service expire node default response a status code equal to that given
func (o *HeadscaleServiceExpireNodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the headscale service expire node default response
func (o *HeadscaleServiceExpireNodeDefault) Code() int {
	return o._statusCode
}

func (o *HeadscaleServiceExpireNodeDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/node/{nodeId}/expire][%d] HeadscaleService_ExpireNode default %s", o._statusCode, payload)
}

func (o *HeadscaleServiceExpireNodeDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/node/{nodeId}/expire][%d] HeadscaleService_ExpireNode default %s", o._statusCode, payload)
}

func (o *HeadscaleServiceExpireNodeDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceExpireNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}