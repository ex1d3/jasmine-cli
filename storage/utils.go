package storage

import "fmt"

func NullStoragePointer(value string) string {
	return fmt.Sprintf("%s => null", value)
}
