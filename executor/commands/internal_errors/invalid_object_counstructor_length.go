package internal_errors

import "fmt"

type InvalidObjectConstructorLengthError struct {
	Want int
	Have int
}

func (e *InvalidObjectConstructorLengthError) Error() string {
	return fmt.Sprintf(
		"invalid object constructor length; want - %d, have - %d",
		e.Want,
		e.Have,
	)
}
