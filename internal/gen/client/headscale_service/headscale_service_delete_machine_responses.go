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

// HeadscaleServiceDeleteMachineReader is a Reader for the HeadscaleServiceDeleteMachine structure.
type HeadscaleServiceDeleteMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceDeleteMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceDeleteMachineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceDeleteMachineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceDeleteMachineOK creates a HeadscaleServiceDeleteMachineOK with default headers values
func NewHeadscaleServiceDeleteMachineOK() *HeadscaleServiceDeleteMachineOK {
	return &HeadscaleServiceDeleteMachineOK{}
}

/*
HeadscaleServiceDeleteMachineOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceDeleteMachineOK struct {
	Payload models.V1DeleteMachineResponse
}

// IsSuccess returns true when this headscale service delete machine o k response has a 2xx status code
func (o *HeadscaleServiceDeleteMachineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service delete machine o k response has a 3xx status code
func (o *HeadscaleServiceDeleteMachineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service delete machine o k response has a 4xx status code
func (o *HeadscaleServiceDeleteMachineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service delete machine o k response has a 5xx status code
func (o *HeadscaleServiceDeleteMachineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service delete machine o k response a status code equal to that given
func (o *HeadscaleServiceDeleteMachineOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the headscale service delete machine o k response
func (o *HeadscaleServiceDeleteMachineOK) Code() int {
	return 200
}

func (o *HeadscaleServiceDeleteMachineOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1/machine/{machineId}][%d] headscaleServiceDeleteMachineOK %s", 200, payload)
}

func (o *HeadscaleServiceDeleteMachineOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1/machine/{machineId}][%d] headscaleServiceDeleteMachineOK %s", 200, payload)
}

func (o *HeadscaleServiceDeleteMachineOK) GetPayload() models.V1DeleteMachineResponse {
	return o.Payload
}

func (o *HeadscaleServiceDeleteMachineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceDeleteMachineDefault creates a HeadscaleServiceDeleteMachineDefault with default headers values
func NewHeadscaleServiceDeleteMachineDefault(code int) *HeadscaleServiceDeleteMachineDefault {
	return &HeadscaleServiceDeleteMachineDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceDeleteMachineDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceDeleteMachineDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this headscale service delete machine default response has a 2xx status code
func (o *HeadscaleServiceDeleteMachineDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service delete machine default response has a 3xx status code
func (o *HeadscaleServiceDeleteMachineDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service delete machine default response has a 4xx status code
func (o *HeadscaleServiceDeleteMachineDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service delete machine default response has a 5xx status code
func (o *HeadscaleServiceDeleteMachineDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service delete machine default response a status code equal to that given
func (o *HeadscaleServiceDeleteMachineDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the headscale service delete machine default response
func (o *HeadscaleServiceDeleteMachineDefault) Code() int {
	return o._statusCode
}

func (o *HeadscaleServiceDeleteMachineDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1/machine/{machineId}][%d] HeadscaleService_DeleteMachine default %s", o._statusCode, payload)
}

func (o *HeadscaleServiceDeleteMachineDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v1/machine/{machineId}][%d] HeadscaleService_DeleteMachine default %s", o._statusCode, payload)
}

func (o *HeadscaleServiceDeleteMachineDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceDeleteMachineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
