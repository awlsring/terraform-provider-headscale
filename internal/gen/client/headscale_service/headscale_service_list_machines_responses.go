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

// HeadscaleServiceListMachinesReader is a Reader for the HeadscaleServiceListMachines structure.
type HeadscaleServiceListMachinesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceListMachinesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceListMachinesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceListMachinesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceListMachinesOK creates a HeadscaleServiceListMachinesOK with default headers values
func NewHeadscaleServiceListMachinesOK() *HeadscaleServiceListMachinesOK {
	return &HeadscaleServiceListMachinesOK{}
}

/*
HeadscaleServiceListMachinesOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceListMachinesOK struct {
	Payload *models.V1ListMachinesResponse
}

// IsSuccess returns true when this headscale service list machines o k response has a 2xx status code
func (o *HeadscaleServiceListMachinesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service list machines o k response has a 3xx status code
func (o *HeadscaleServiceListMachinesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service list machines o k response has a 4xx status code
func (o *HeadscaleServiceListMachinesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service list machines o k response has a 5xx status code
func (o *HeadscaleServiceListMachinesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service list machines o k response a status code equal to that given
func (o *HeadscaleServiceListMachinesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the headscale service list machines o k response
func (o *HeadscaleServiceListMachinesOK) Code() int {
	return 200
}

func (o *HeadscaleServiceListMachinesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/machine][%d] headscaleServiceListMachinesOK %s", 200, payload)
}

func (o *HeadscaleServiceListMachinesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/machine][%d] headscaleServiceListMachinesOK %s", 200, payload)
}

func (o *HeadscaleServiceListMachinesOK) GetPayload() *models.V1ListMachinesResponse {
	return o.Payload
}

func (o *HeadscaleServiceListMachinesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ListMachinesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceListMachinesDefault creates a HeadscaleServiceListMachinesDefault with default headers values
func NewHeadscaleServiceListMachinesDefault(code int) *HeadscaleServiceListMachinesDefault {
	return &HeadscaleServiceListMachinesDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceListMachinesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceListMachinesDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this headscale service list machines default response has a 2xx status code
func (o *HeadscaleServiceListMachinesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service list machines default response has a 3xx status code
func (o *HeadscaleServiceListMachinesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service list machines default response has a 4xx status code
func (o *HeadscaleServiceListMachinesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service list machines default response has a 5xx status code
func (o *HeadscaleServiceListMachinesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service list machines default response a status code equal to that given
func (o *HeadscaleServiceListMachinesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the headscale service list machines default response
func (o *HeadscaleServiceListMachinesDefault) Code() int {
	return o._statusCode
}

func (o *HeadscaleServiceListMachinesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/machine][%d] HeadscaleService_ListMachines default %s", o._statusCode, payload)
}

func (o *HeadscaleServiceListMachinesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/machine][%d] HeadscaleService_ListMachines default %s", o._statusCode, payload)
}

func (o *HeadscaleServiceListMachinesDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceListMachinesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
