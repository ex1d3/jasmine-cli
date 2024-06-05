package internal_errors

import "fmt"

type InvalidObjectConstructorValueError struct {
	Index int
	Value any
}

func (e *InvalidObjectConstructorValueError) Error() string {
	return fmt.Sprintf(
		"invalid object constructor value - %v; index - %d",
		e.Value,
		e.Index,
	)
}
