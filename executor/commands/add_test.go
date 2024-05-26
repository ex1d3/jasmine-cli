package commands

import (
	"nolono-cli/storage"
	"strings"
	"testing"
)

func TestAddValidSrc(t *testing.T) {
	result, err := Add(strings.Split("add src (adam)", " "))

	if err != nil {
		t.Fatal("unexpected error")
	}

	if result != "adam" {
		t.Fatal("unexpected return value")
	}

	if !storage.Src["adam"] {
		t.Fatal("source is not inserted into storage")
	}
}

func TestAddValidTx(t *testing.T) {
	storage.Src["adam"] = true

	result, err := Add(strings.Split("add tx (adam;200)", " "))

	if err != nil {
		t.Fatal("unexpected error")
	}

	if result == "" {
		t.Fatal("unexpected return value")
	}

	id := strings.Split(result, " ")[0]

	if _, ok := storage.Tx[id]; !ok {
		t.Fatal("transaction is not inserted into storage")
	}
}

func TestAddTxWithInvalidSrc(t *testing.T) {
	result, err := Add(strings.Split("add tx (invalid_src;200)", " "))

	if err == nil {
		t.Fatal("adding transaction with invalid src is not prevented")
	}

	if result != "" {
		t.Fatal("unexpected return value")
	}
}

func TestAddTxWithInvalidObjectConstructor(t *testing.T) {
	result, err := Add(strings.Split("add tx (adam;200;unexpected_arg)", " "))

	if err == nil {
		t.Fatal(
			("adding transaction with invalid" +
				"object constructor is not prevented"),
		)
	}

	if result != "" {
		t.Fatal("unexpected return value")
	}
}

func TestAddSrcWithInvalidObjectConstructor(t *testing.T) {
	result, err := Add(strings.Split("add src (adam;unexpected_arg)", " "))

	if err == nil {
		t.Fatal(
			("adding src with invalid" +
				"object constructor is not prevented"),
		)
	}

	if result != "" {
		t.Fatal("unexpected return value")
	}
}

func TestAddWithoutObjectConstructor(t *testing.T) {
	result, err := Add(strings.Split("add src adam", " "))

	if err == nil {
		t.Fatal("adding without object constructor is not prevented")
	}

	if result != "" {
		t.Fatal("Unexpected return value")
	}
}

func TestAddWithEmptyObjectConstructor(t *testing.T) {
	result, err := Add(strings.Split("add src ()", " "))

	if err == nil {
		t.Fatal("adding with empty object constructor is not prevented")
	}

	if result != "" {
		t.Fatal("unexpected return value")
	}
}

func TestAddToInvalidCollection(t *testing.T) {
	result, err := Add(strings.Split("add invalid_collection adam", " "))

	if err == nil {
		t.Fatal("adding object to invalid collection is not prevented")
	}

	if result != "" {
		t.Fatal("unexpected return value")
	}
}
