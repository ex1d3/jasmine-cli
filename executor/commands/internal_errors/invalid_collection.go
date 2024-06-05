package internal_errors

import "fmt"

type InvalidCollectionError struct {
	Collection string
}

func (e *InvalidCollectionError) Error() string {
	return fmt.Sprintf("invalid collection '%s'", e.Collection)
}
