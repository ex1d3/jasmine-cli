package internal_errors

import "fmt"

func InvalidArgsCount(cmd string, expected int, received int) string {
	return fmt.Sprintf(
		"invalid args count for '%s' command (expected: %d, received: %d)",
		cmd,
		expected,
		received,
	)
}
