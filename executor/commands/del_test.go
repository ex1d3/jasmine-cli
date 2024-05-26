package commands

import (
	"strings"
	"testing"
)

func TestDelNonExistingSrc(t *testing.T) {
	result, err := Del(strings.Split("del src invalid", " "))

	if err == nil {
		t.Fatal("deletion of non existing src not pervented")
	}

	if result != "" {
		t.Fatal("unexpected return value")
	}
}

func TestDelNonExistingTx(t *testing.T) {
	result, err := Del(strings.Split("del tx invalid", " "))

	if err == nil {
		t.Fatal("deletion of non existing src not pervented")
	}

	if result != "" {
		t.Fatal("unexpected return value")
	}
}

func TestDelElementFromNonexistingCollection(t *testing.T) {
	result, err := Del(strings.Split("del unknown_collection id", " "))

	if err == nil {
		t.Fatal(
			"calling 'del' command on " +
				"non existing collection is not pervented",
		)
	}

	if result != "" {
		t.Fatal("unexpected return value")
	}
}

func TestDelWithUnexpectedArgs(t *testing.T) {
	result, err := Del(strings.Split("del tx id invalid_arg another_arg", " "))

	if err == nil {
		t.Fatal("calling 'del' command with unexpecting args is not pervented")
	}

	if result != "" {
		t.Fatal("unexpected return value")
	}
}
