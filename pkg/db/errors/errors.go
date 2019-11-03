package errors

import (
	"fmt"
)

type InvalidTypeIndexError struct {
	TypeName string
	Index    int8
}

func (e *InvalidTypeIndexError) Error() string {
	return fmt.Sprintf("Could not find a %q type associated with the index %d", e.TypeName, e.Index)
}

type InvalidTypeError struct {
	TypeName string
	Value    string
}

func (e *InvalidTypeError) Error() string {
	return fmt.Sprintf("Invalid %q type value: %q", e.TypeName, e.Value)
}

type ValueNotExistError struct {
	Value         string
	ContainerName string
}

func (e *ValueNotExistError) Error() string {
	return fmt.Sprintf("The value %q does not exist in the container: %q", e.Value, e.ContainerName)
}

type TableNotExistError struct {
	NameTable string
}

func (e *TableNotExistError) Error() string {
	return fmt.Sprintf("The table %q does not exist", e.NameTable)
}
