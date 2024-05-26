package commands

import (
	"nolono-cli/storage"
	"strings"
	"testing"
)

func TestInvalidSrcDelTarget(t *testing.T) {
	t.Log(storage.Src)

	_, err := Del(strings.Split("del src invalid", " "))

	if err == nil {
		t.Fatal("Invalid target 'del' call execution not pervented")
	}
}

func TestInvalidTxDelTarget(t *testing.T) {
	_, err := Del(strings.Split("del tx invalid", " "))

	if err == nil {
		t.Fatal("Invalid target 'del' call execution not pervented")
	}
}

func TestInvalidDelCollection(t *testing.T) {
	_, err := Del(strings.Split("del unknown_collection id", " "))

	if err == nil {
		t.Fatal("Invalid collection 'del' call execution not pervented")
	}
}
