package errors

import "fmt"

func InvalidCollection(collection string) string {
	return fmt.Sprintf("Invalid collection (%s)", collection)
}
