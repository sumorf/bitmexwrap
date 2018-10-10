// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sumorf/coinex/bitmex/models"
)

// UserConfirmWithdrawalReader is a Reader for the UserConfirmWithdrawal structure.
type UserConfirmWithdrawalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserConfirmWithdrawalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserConfirmWithdrawalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserConfirmWithdrawalOK creates a UserConfirmWithdrawalOK with default headers values
func NewUserConfirmWithdrawalOK() *UserConfirmWithdrawalOK {
	return &UserConfirmWithdrawalOK{}
}

/*UserConfirmWithdrawalOK handles this case with default header values.

Request was successful
*/
type UserConfirmWithdrawalOK struct {
	Payload *models.Transaction
}

func (o *UserConfirmWithdrawalOK) Error() string {
	return fmt.Sprintf("[POST /user/confirmWithdrawal][%d] userConfirmWithdrawalOK  %+v", 200, o.Payload)
}

func (o *UserConfirmWithdrawalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Transaction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
