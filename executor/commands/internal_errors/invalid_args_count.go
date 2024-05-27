package internal_errors

import "fmt"

func InvalidArgsCount(
	cmd string,
	// for support "from x to y" statements in errors
	expected string,
	received int,
) string {
	return fmt.Sprintf(
		"invalid args count for '%s' command (expected: %s, received: %d)",
		cmd,
		expected,
		received,
	)
}
