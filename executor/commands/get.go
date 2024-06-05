package commands

import (
	"jasmine-cli/domain"
	"jasmine-cli/executor/collections"
	"jasmine-cli/executor/commands/internal_errors"
	"jasmine-cli/storage"
)

func Get(args []string) ([]interface{}, error) {
	if len(args) != 2 {
		return []interface{}{}, &internal_errors.InvalidArgsCountError{
			Command: "get",
			Want:    2,
			Have:    len(args),
		}
	}

	collection := args[0]
	target := args[1]

	switch collection {
	case collections.SRC:
		{
			return executeGet[domain.Src](target, storage.Src)
		}
	case collections.TX:
		{
			return executeGet[domain.Tx](target, storage.Tx)
		}
	default:
		{
			return []interface{}{}, &internal_errors.InvalidCollectionError{
				Collection: collection,
			}
		}
	}
}

func executeGet[T any](
	target string,
	entityStorage storage.Storage[string, *T],
) ([]interface{}, error) {
	rawStorage := entityStorage.GetStorage()

	if target == "*" {
		entities := make([]interface{}, len(rawStorage))
		i := 0

		for _, v := range rawStorage {
			entities[i] = v

			i++
		}

		return entities, nil
	}

	if entity := entityStorage.Get(target); entity == nil {
		return []interface{}{}, nil
	} else {
		return []interface{}{entity}, nil
	}
}
