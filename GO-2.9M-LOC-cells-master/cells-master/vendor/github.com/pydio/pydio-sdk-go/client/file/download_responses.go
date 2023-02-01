// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DownloadReader is a Reader for the Download structure.
type DownloadReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *DownloadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDownloadOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDownloadOK creates a DownloadOK with default headers values
func NewDownloadOK(writer io.Writer) *DownloadOK {
	return &DownloadOK{
		Payload: writer,
	}
}

/*DownloadOK handles this case with default header values.

Successful Response
*/
type DownloadOK struct {
	Payload io.Writer
}

func (o *DownloadOK) Error() string {
	return fmt.Sprintf("[GET /io/{path}][%d] downloadOK  %+v", 200, o.Payload)
}

func (o *DownloadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
