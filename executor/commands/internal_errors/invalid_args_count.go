package internal_errors

import (
	"fmt"
)

type InvalidArgsCountError struct {
	Command string
	Want    int
	Have    int
}

func (e *InvalidArgsCountError) Error() string {
	return fmt.Sprintf(
		"invalid args count for command '%s'; want: %d, have: %d",
		e.Command,
		e.Want,
		e.Have,
	)
}
