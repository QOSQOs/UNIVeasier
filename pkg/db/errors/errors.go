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
	Value          string
	CollectionName string
}

func (e *ValueNotExistError) Error() string {
	return fmt.Sprintf("The value %q does not exist in the collection: %q", e.Value, e.CollectionName)
}

type TableNotExistError struct {
	NameTable string
}

func (e *TableNotExistError) Error() string {
	return fmt.Sprintf("The table %q does not exist", e.NameTable)
}

type NotEqualsSizeError struct {
	FirstCollection   string
	SecondtCollection string
}

func (e *NotEqualsSizeError) Error() string {
	return fmt.Sprintf("%q and %q does not have the same size", e.FirstCollection, e.SecondtCollection)
}

type OperationNotSupportedError struct {
	MethodName string
	Operation  string
}

func (e *OperationNotSupportedError) Error() string {
	return fmt.Sprintf("The %q method is not supported by the %q operation", e.MethodName, e.Operation)
}

type InvalidFirstFilterError struct{}

func (e *InvalidFirstFilterError) Error() string {
	return fmt.Sprint("From the second filter the logical operator is mandatory")
}

type InvalidFilterError struct{}

func (e *InvalidFilterError) Error() string {
	return fmt.Sprint("The first filter can not contain a logical operator")
}

type EmptyHeaderExpressionError struct{}

func (e *EmptyHeaderExpressionError) Error() string {
	return fmt.Sprint("The UPDATE operation must contain at least one column")
}

type InvalidHeaderExpressionError struct{}

func (e *InvalidHeaderExpressionError) Error() string {
	return fmt.Sprint("The DELETE operation must not contain any columns")
}
