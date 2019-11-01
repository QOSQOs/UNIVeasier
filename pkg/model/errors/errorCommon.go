package errors

import (
	"fmt"
)

type NullValueNotAllowedError struct {
	Attribute string
}

func (e *NullValueNotAllowedError) Error() string {
	return fmt.Sprintf("The %q attribute can not be NULL", e.Attribute)
}
