package commands

import (
	"nolono-cli/storage"
	"strings"
	"testing"
)

func TestAddSrc(t *testing.T) {
	result := Add([]string{"add", "src", "(adam)"})

	if (len(result) != 1) || (result[0] != "adam") {
		t.Fatal("Invalid return value")
	}

	if !storage.Src["adam"] {
		t.Fatal("Source does not inserted into storage")
	}
}

func TestAddTx(t *testing.T) {
	storage.Src["adam"] = true

	result := Add([]string{"add", "tx", "(adam;200)"})

	if len(result) != 1 || result[0] == "" {
		t.Fatal("Invalid return value")
	}

	id := strings.Split(result[0], " ")[0]

	if _, ok := storage.Tx[id]; !ok {
		t.Fatal("Transaction does not inserted into storage")
	}
}
