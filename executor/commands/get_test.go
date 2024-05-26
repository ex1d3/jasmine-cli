package commands

import (
	"nolono-cli/domain"
	"nolono-cli/storage"
	"strings"
	"testing"
)

func TestGetAllTxs(t *testing.T) {
	storage.ClearTx()

	storage.Src["adam"] = true

	storage.Tx["1"] = domain.NewTx("adam", 200)
	storage.Tx["2"] = domain.NewTx("adam", 30)

	txs, err := Get(strings.Split("get tx", " "))

	if err != nil {
		t.Fatal("unexpected error")
	}

	if len(txs) != 2 {
		t.Fatal("unexpected return value length")
	}

	if (txs[0] != storage.Tx["1"].ToStr("1")) ||
		(txs[1] != storage.Tx["2"].ToStr("2")) {
		t.Fatal("unexpected return value")
	}
}

func TestGetExistingTx(t *testing.T) {
	txId := "3"

	storage.Src["ben"] = true
	storage.Tx[txId] = domain.NewTx("ben", 20)

	txs, err := Get(strings.Split("get tx 3", " "))

	if err != nil {
		t.Fatal("unexpected error")
	}

	if len(txs) != 1 {
		t.Fatal("unexpected return value length")
	}

	if txs[0] != storage.Tx[txId].ToStr(txId) {
		t.Fatal("unexpected return value")
	}
}

func TestGetNonExistingTx(t *testing.T) {
	storage.ClearTx()

	txs, err := Get(strings.Split("get tx 1", " "))

	if err != nil {
		t.Fatal("unexpected error")
	}

	if len(txs) != 0 {
		t.Fatal("unexpected return value length")
	}
}

func TestGetAllSrcs(t *testing.T) {
	storage.ClearSrc()

	storage.Src["adam"] = true
	storage.Src["ben"] = true

	srcs, err := Get(strings.Split("get src", " "))

	if err != nil {
		t.Fatal("unexpected error")
	}

	if len(srcs) != 2 {
		t.Fatal("unexpected return value length")
	}

	if (srcs[0] != "adam") || (srcs[1] != "ben") {
		t.Fatal("unexpected return value")
	}
}

func TestGetExistingSrc(t *testing.T) {
	storage.ClearSrc()

	storage.Src["adam"] = true

	src, err := Get(strings.Split("get src adam", " "))

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
	storage.ClearSrc()

	src, err := Get(strings.Split("get src adam", " "))

	if err != nil {
		t.Fatal("unexpected error")
	}

	if len(src) != 0 {
		t.Fatal("unexpected return value length")
	}
}
