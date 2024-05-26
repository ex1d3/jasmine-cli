package storage

import "nolono-cli/domain"

var Tx = make(map[string]*domain.Tx)

func ClearTx() {
	for k := range Tx {
		delete(Tx, k)
	}
}
