// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/pydio-sdk-go/models"
)

// UpdateRoleReader is a Reader for the UpdateRole structure.
type UpdateRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateRoleOK creates a UpdateRoleOK with default headers values
func NewUpdateRoleOK() *UpdateRoleOK {
	return &UpdateRoleOK{}
}

/*UpdateRoleOK handles this case with default header values.

Successful response
*/
type UpdateRoleOK struct {
	Payload *models.Role
}

func (o *UpdateRoleOK) Error() string {
	return fmt.Sprintf("[PATCH /admin/roles/{roleId}][%d] updateRoleOK  %+v", 200, o.Payload)
}

func (o *UpdateRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
