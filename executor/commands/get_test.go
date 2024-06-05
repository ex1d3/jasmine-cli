package commands

import (
	"fmt"
	"jasmine-cli/domain"
	"jasmine-cli/storage"
	"strings"
	"testing"
)

func TestGetAllTxs(t *testing.T) {
	txStorage := storage.Tx
	txStorage.Unload()

	storage.Src.Set("adam", domain.NewSrc("adam"))

	want := [2]*domain.Tx{
		domain.NewTx("1", "adam", 200),
		domain.NewTx("2", "adam", 30),
	}

	txStorage.Set("1", want[0])
	txStorage.Set("2", want[1])

	command := strings.Split("tx *", " ")
	txs, err := Get(command)

	errorMessage := fmt.Sprintf(
		"Get(%q) = %v, %v; want %v, %v",
		command,
		txs,
		err,
		want,
		nil,
	)

	if err != nil {
		t.Fatalf(errorMessage)
	}

	if len(txs) != 2 {
		t.Fatalf(errorMessage)
	}

	for _, tx := range txs {
		if tx != want[0] && tx != want[1] {
			t.Fatalf(errorMessage)
		}
	}
}

func TestGetTx(t *testing.T) {
	txId := "1"
	tx := domain.NewTx(txId, "ben", 20)

	storage.Src.Set("ben", domain.NewSrc("ben"))
	storage.Tx.Set(txId, tx)

	command := strings.Split(fmt.Sprintf("tx %s", txId), " ")
	txs, err := Get(command)

	errorMessage := fmt.Sprintf(
		"Get(%q) = %v, %v; want %v, %v",
		command,
		txs,
		err,
		[]interface{}{tx},
		nil,
	)

	if err != nil {
		t.Fatalf(errorMessage)
	}

	if len(txs) != 1 {
		t.Fatalf(errorMessage)
	}

	if txs[0] != tx {
		t.Fatalf(errorMessage)
	}
}

func TestGetInvalidTx(t *testing.T) {
	storage.Tx.Unload()

	command := strings.Split("tx 1", " ")
	txs, err := Get(command)

	errorMessage := fmt.Sprintf(
		"Get(%q) = %v, %v; want %v, %v",
		command,
		txs,
		err,
		[]interface{}{},
		nil,
	)

	if err != nil {
		t.Fatal(errorMessage)
	}

	if len(txs) != 0 {
		t.Fatal(errorMessage)
	}
}

func TestGetAllSrcs(t *testing.T) {
	srcStorage := storage.Src

	srcStorage.Unload()

	want := [2]*domain.Src{
		domain.NewSrc("adam"),
		domain.NewSrc("ben"),
	}

	adam := want[0]
	ben := want[1]

	srcStorage.Set("adam", adam)
	srcStorage.Set("ben", ben)

	command := strings.Split("src *", " ")
	srcs, err := Get(command)

	errorMessage := fmt.Sprintf(
		"Get(%q) = %v, %v; want %v, %v",
		command,
		srcs,
		err,
		want,
		nil,
	)

	if err != nil {
		t.Fatal(errorMessage)
	}

	if len(srcs) != 2 {
		t.Fatal(errorMessage)
	}

	for _, src := range srcs {
		if src != adam && src != ben {
			t.Fatal(errorMessage)
		}
	}
}

func TestGetSrc(t *testing.T) {
	srcStorage := storage.Src

	srcStorage.Unload()

	adam := domain.NewSrc("adam")
	srcStorage.Set("adam", adam)

	command := strings.Split("src adam", " ")
	srcs, err := Get(command)

	errorMessage := fmt.Sprintf(
		"Get(%q) = %v, %v; want %v, %v",
		command,
		srcs,
		err,
		[]interface{}{adam},
		nil,
	)

	if err != nil {
		t.Fatal(errorMessage)
	}

	if len(srcs) != 1 {
		t.Fatal(errorMessage)
	}

	if srcs[0] != adam {
		t.Fatal(errorMessage)
	}
}

func TestGetInvalidSrc(t *testing.T) {
	storage.Src.Unload()

	command := strings.Split("src adam", " ")
	src, err := Get(command)

	errorMessage := fmt.Sprintf(
		"Get(%q) = %v, %v; want %v, %v",
		command,
		src,
		err,
		[]interface{}{},
		nil,
	)

	if err != nil {
		t.Fatal(errorMessage)
	}

	if len(src) != 0 {
		t.Fatal(errorMessage)
	}
}

func TestGetWithUnexpectedArgs(t *testing.T) {
	command := strings.Split("abc abc abc", " ")
	result, err := Get(command)

	errorMessage := fmt.Sprintf(
		"Get(%q) = %v, %v; want %v, %v",
		command,
		result,
		err,
		[]interface{}{},
		"error",
	)

	if err == nil {
		t.Fatal(errorMessage)
	}

	if len(result) != 0 {
		t.Fatal(errorMessage)
	}
}

func TestGetWithInvalidCollection(t *testing.T) {
	command := strings.Split("abc abc", " ")
	result, err := Get(command)

	errorMessage := fmt.Sprintf(
		"Get(%q) = %v, %v; want %v, %v",
		command,
		result,
		err,
		[]interface{}{},
		"error",
	)

	if err == nil {
		t.Fatal(errorMessage)
	}

	if len(result) != 0 {
		t.Fatal(errorMessage)
	}
}
