package internal_errors

type ObjectConstructorNotFoundError struct {
}

func (e *ObjectConstructorNotFoundError) Error() string {
	return "object constructor not found"
}
