package storage

import "jasmine-cli/domain"

var Src = NewJSONStorage[string, bool]("src-storage.json")
var Tx = NewJSONStorage[string, *domain.Tx]("tx-storage.json")
