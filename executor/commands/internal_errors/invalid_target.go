package internal_errors

import "fmt"

type InvalidTargetError struct {
	Target string
}

func (e *InvalidTargetError) Error() string {
	return fmt.Sprintf("invalid taget '%s'", e.Target)
}
