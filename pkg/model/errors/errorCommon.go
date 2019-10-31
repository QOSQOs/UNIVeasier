package errors

import (
	"fmt"
)

type NullValueNotAllowed struct {
	Attribute string
}

func (e *NullValueNotAllowed) Error() string {
	return fmt.Sprintf("The %q attribute can not be NULL", e.Attribute)
}
