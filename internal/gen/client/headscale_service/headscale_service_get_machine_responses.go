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

// HeadscaleServiceGetMachineReader is a Reader for the HeadscaleServiceGetMachine structure.
type HeadscaleServiceGetMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceGetMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceGetMachineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceGetMachineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceGetMachineOK creates a HeadscaleServiceGetMachineOK with default headers values
func NewHeadscaleServiceGetMachineOK() *HeadscaleServiceGetMachineOK {
	return &HeadscaleServiceGetMachineOK{}
}

/*
HeadscaleServiceGetMachineOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceGetMachineOK struct {
	Payload *models.V1GetMachineResponse
}

// IsSuccess returns true when this headscale service get machine o k response has a 2xx status code
func (o *HeadscaleServiceGetMachineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service get machine o k response has a 3xx status code
func (o *HeadscaleServiceGetMachineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service get machine o k response has a 4xx status code
func (o *HeadscaleServiceGetMachineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service get machine o k response has a 5xx status code
func (o *HeadscaleServiceGetMachineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service get machine o k response a status code equal to that given
func (o *HeadscaleServiceGetMachineOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the headscale service get machine o k response
func (o *HeadscaleServiceGetMachineOK) Code() int {
	return 200
}

func (o *HeadscaleServiceGetMachineOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/machine/{machineId}][%d] headscaleServiceGetMachineOK %s", 200, payload)
}

func (o *HeadscaleServiceGetMachineOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/machine/{machineId}][%d] headscaleServiceGetMachineOK %s", 200, payload)
}

func (o *HeadscaleServiceGetMachineOK) GetPayload() *models.V1GetMachineResponse {
	return o.Payload
}

func (o *HeadscaleServiceGetMachineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetMachineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceGetMachineDefault creates a HeadscaleServiceGetMachineDefault with default headers values
func NewHeadscaleServiceGetMachineDefault(code int) *HeadscaleServiceGetMachineDefault {
	return &HeadscaleServiceGetMachineDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceGetMachineDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceGetMachineDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this headscale service get machine default response has a 2xx status code
func (o *HeadscaleServiceGetMachineDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service get machine default response has a 3xx status code
func (o *HeadscaleServiceGetMachineDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service get machine default response has a 4xx status code
func (o *HeadscaleServiceGetMachineDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service get machine default response has a 5xx status code
func (o *HeadscaleServiceGetMachineDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service get machine default response a status code equal to that given
func (o *HeadscaleServiceGetMachineDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the headscale service get machine default response
func (o *HeadscaleServiceGetMachineDefault) Code() int {
	return o._statusCode
}

func (o *HeadscaleServiceGetMachineDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/machine/{machineId}][%d] HeadscaleService_GetMachine default %s", o._statusCode, payload)
}

func (o *HeadscaleServiceGetMachineDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/machine/{machineId}][%d] HeadscaleService_GetMachine default %s", o._statusCode, payload)
}

func (o *HeadscaleServiceGetMachineDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceGetMachineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
