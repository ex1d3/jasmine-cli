package storage

import "nolono-cli/domain"

var Src = NewJSONStorage[string, bool]("src-storage.json")
var Tx = NewJSONStorage[string, *domain.Tx]("tx-storage.json")
