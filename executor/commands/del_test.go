package commands

import (
	"strings"
	"testing"
)

func TestDelNonExistingSrc(t *testing.T) {
	result, err := Del(strings.Split("src invalid", " "))

	if err == nil {
		t.Fatal("deletion of non existing src not pervented")
	}

	if len(result) != 0 {
		t.Fatal("unexpected return value len")
	}
}

func TestDelNonExistingTx(t *testing.T) {
	result, err := Del(strings.Split("tx invalid", " "))

	if err == nil {
		t.Fatal("deletion of non existing src not pervented")
	}

	if len(result) != 0 {
		t.Fatal("unexpected return value len")
	}
}

func TestDelElementFromNonexistingCollection(t *testing.T) {
	result, err := Del(strings.Split("unknown_collection id", " "))

	if err == nil {
		t.Fatal(
			"calling 'del' command on " +
				"non existing collection is not pervented",
		)
	}

	if len(result) != 0 {
		t.Fatal("unexpected return value len")
	}
}

func TestDelWithUnexpectedArgs(t *testing.T) {
	result, err := Del(strings.Split("tx id invalid_arg 123 another_arg", " "))

	if err == nil {
		t.Fatal("calling 'del' command with unexpecting args is not pervented")
	}

	if len(result) != 0 {
		t.Fatal("unexpected return value len")
	}
}
