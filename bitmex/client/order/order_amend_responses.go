// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sumorf/coinex/bitmex/models"
)

// OrderAmendReader is a Reader for the OrderAmend structure.
type OrderAmendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderAmendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOrderAmendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewOrderAmendBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewOrderAmendUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewOrderAmendNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrderAmendOK creates a OrderAmendOK with default headers values
func NewOrderAmendOK() *OrderAmendOK {
	return &OrderAmendOK{}
}

/*OrderAmendOK handles this case with default header values.

Request was successful
*/
type OrderAmendOK struct {
	Payload *models.Order
}

func (o *OrderAmendOK) Error() string {
	return fmt.Sprintf("[PUT /order][%d] orderAmendOK  %+v", 200, o.Payload)
}

func (o *OrderAmendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Order)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderAmendBadRequest creates a OrderAmendBadRequest with default headers values
func NewOrderAmendBadRequest() *OrderAmendBadRequest {
	return &OrderAmendBadRequest{}
}

/*OrderAmendBadRequest handles this case with default header values.

Parameter Error
*/
type OrderAmendBadRequest struct {
	Payload *models.Error
}

func (o *OrderAmendBadRequest) Error() string {
	return fmt.Sprintf("[PUT /order][%d] orderAmendBadRequest  %+v", 400, o.Payload)
}

func (o *OrderAmendBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderAmendUnauthorized creates a OrderAmendUnauthorized with default headers values
func NewOrderAmendUnauthorized() *OrderAmendUnauthorized {
	return &OrderAmendUnauthorized{}
}

/*OrderAmendUnauthorized handles this case with default header values.

Unauthorized
*/
type OrderAmendUnauthorized struct {
	Payload *models.Error
}

func (o *OrderAmendUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /order][%d] orderAmendUnauthorized  %+v", 401, o.Payload)
}

func (o *OrderAmendUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderAmendNotFound creates a OrderAmendNotFound with default headers values
func NewOrderAmendNotFound() *OrderAmendNotFound {
	return &OrderAmendNotFound{}
}

/*OrderAmendNotFound handles this case with default header values.

Not Found
*/
type OrderAmendNotFound struct {
	Payload *models.Error
}

func (o *OrderAmendNotFound) Error() string {
	return fmt.Sprintf("[PUT /order][%d] orderAmendNotFound  %+v", 404, o.Payload)
}

func (o *OrderAmendNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
