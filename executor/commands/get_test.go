package commands

import (
	"nolono-cli/domain"
	"nolono-cli/storage"
	"strings"
	"testing"
)

func TestGetAllTxs(t *testing.T) {
	txStorage := storage.Tx
	txStorage.Unload()

	storage.Src.Set("adam", true)

	txStorage.Set("1", domain.NewTx("adam", 200))
	txStorage.Set("2", domain.NewTx("adam", 30))

	txs, err := Get(strings.Split("tx", " "))

	if err != nil {
		t.Fatal("unexpected error")
	}

	if len(txs) != 2 {
		t.Fatal("unexpected return value length")
	}

	firstTx := txStorage.Get("1").ToStr("1")
	secondTx := txStorage.Get("2").ToStr("2")

	for _, tx := range txs {
		if tx != firstTx &&
			tx != secondTx {
			t.Fatal("unexpected return value")
		}
	}
}

func TestGetExistingTx(t *testing.T) {
	txId := "3"

	storage.Src.Set("ben", true)
	storage.Tx.Set(txId, domain.NewTx("ben", 20))

	txs, err := Get(strings.Split("tx 3", " "))

	if err != nil {
		t.Fatal("unexpected error")
	}

	if len(txs) != 1 {
		t.Fatal("unexpected return value length")
	}

	if txs[0] != storage.Tx.Get(txId).ToStr(txId) {
		t.Fatal("unexpected return value")
	}
}

func TestGetNonExistingTx(t *testing.T) {
	storage.Tx.Unload()

	txs, err := Get(strings.Split("tx 1", " "))

	if err != nil {
		t.Fatal("unexpected error")
	}

	if len(txs) != 0 {
		t.Fatal("unexpected return value length")
	}
}

func TestGetAllSrcs(t *testing.T) {
	srcStorage := storage.Src

	srcStorage.Unload()
	srcStorage.Set("adam", true)
	srcStorage.Set("ben", true)

	srcs, err := Get(strings.Split("src", " "))

	if err != nil {
		t.Fatal("unexpected error")
	}

	if len(srcs) != 2 {
		t.Fatal("unexpected return value length")
	}

	for _, src := range srcs {
		if src != "adam" && src != "ben" {
			t.Fatal("unexpected return value")
		}
	}
}

func TestGetExistingSrc(t *testing.T) {
	srcStorage := storage.Src

	srcStorage.Unload()
	srcStorage.Set("adam", true)

	src, err := Get(strings.Split("src adam", " "))

	if err != nil {
		t.Fatal("unexpected error")
	}

	if len(src) != 1 {
		t.Fatal("unexpected return value length")
	}

	if src[0] != "adam" {
		t.Fatal("unexpected return value")
	}
}

func TestGetNonExistingSrc(t *testing.T) {
	storage.Src.Unload()

	src, err := Get(strings.Split("src adam", " "))

	if err != nil {
		t.Fatal("unexpected error")
	}

	if len(src) != 0 {
		t.Fatal("unexpected return value length")
	}
}
