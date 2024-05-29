package commands

import (
	"jasmine-cli/domain"
	"jasmine-cli/storage"
	"strings"
	"testing"
)

func TestGetAllTxs(t *testing.T) {
	txStorage := storage.Tx
	txStorage.Unload()

	storage.Src.Set("adam", domain.NewSrc("adam"))

	firstTx := domain.NewTx("1", "adam", 200)
	secondTx := domain.NewTx("2", "adam", 30)

	txStorage.Set("1", firstTx)
	txStorage.Set("2", secondTx)

	txs, err := Get(strings.Split("tx *", " "))

	if err != nil {
		t.Fatal("unexpected error")
	}

	if len(txs) != 2 {
		t.Fatal("unexpected return value length")
	}

	for _, tx := range txs {
		if tx != firstTx && tx != secondTx {
			t.Fatal("unexpected return value")
		}
	}
}

func TestGetExistingTx(t *testing.T) {
	txId := "3"
	tx := domain.NewTx("1", "ben", 20)

	storage.Src.Set("ben", domain.NewSrc("ben"))
	storage.Tx.Set(txId, tx)

	txs, err := Get(strings.Split("tx 3", " "))

	if err != nil {
		t.Fatal("unexpected error")
	}

	if len(txs) != 1 {
		t.Fatal("unexpected return value length")
	}

	if txs[0] != storage.Tx.Get(txId) {
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

	adam := domain.NewSrc("adam")
	ben := domain.NewSrc("ben")

	srcStorage.Set("adam", adam)
	srcStorage.Set("ben", ben)

	srcs, err := Get(strings.Split("src *", " "))

	if err != nil {
		t.Fatal("unexpected error")
	}

	if len(srcs) != 2 {
		t.Fatal("unexpected return value length")
	}

	for _, src := range srcs {
		if src != adam && src != ben {
			t.Fatal("unexpected return value")
		}
	}
}

func TestGetExistingSrc(t *testing.T) {
	srcStorage := storage.Src

	srcStorage.Unload()

	adam := domain.NewSrc("adam")

	srcStorage.Set("adam", adam)
	srcs, err := Get(strings.Split("src adam", " "))

	if err != nil {
		t.Fatal("unexpected error")
	}

	if len(srcs) != 1 {
		t.Fatal("unexpected return value length")
	}

	if srcs[0] != adam {
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

func TestGetWithUnexpectedArgsCount(t *testing.T) {
	result, err := Get(strings.Split("abc abc abc", " "))

	if err == nil {
		t.Fatal("call with unexpected args count not pervented")
	}

	if len(result) != 0 {
		t.Fatal("unexpected return value len")
	}
}

func TestGetFromNonexistingCollection(t *testing.T) {
	result, err := Get(strings.Split("abc abc", " "))

	if err == nil {
		t.Fatal("call to non existing collection not pervented")
	}

	if len(result) != 0 {
		t.Fatal("unexpected return value len")
	}
}
