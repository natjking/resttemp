package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// Response response
// swagger:model Response
type Response struct {

	// celsius
	Celsius float32 `json:"celsius,omitempty"`

	// fahrenheit
	Fahrenheit float32 `json:"fahrenheit,omitempty"`
}

// Validate validates this response
func (m *Response) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
