package commands

import (
	"errors"
	"jasmine-cli/executor/collections"
	"jasmine-cli/executor/commands/internal_errors"
	"jasmine-cli/storage"
)

func Get(args []string) ([]interface{}, error) {
	if len(args) != 2 {
		return []interface{}{}, errors.New(
			internal_errors.InvalidArgsCount("get", "2", len(args)),
		)
	}

	collection := args[0]
	target := args[1]

	switch collection {
	case collections.SRC:
		{
			return executeGet(target, storage.Src)
		}
	case collections.TX:
		{
			return executeGet(target, storage.Tx)
		}
	default:
		{
			return []interface{}{}, errors.New(
				internal_errors.InvalidCollection(collection),
			)
		}
	}
}

func executeGet[T any](
	target string,
	entityStorage storage.Storage[string, *T],
) ([]interface{}, error) {
	rawStorage := entityStorage.GetStorage()

	if target == "*" {
		srcs := make([]interface{}, len(rawStorage))
		i := 0

		for _, v := range rawStorage {
			srcs[i] = v

			i++
		}

		return srcs, nil
	}

	if entity := entityStorage.Get(target); entity == nil {
		return []interface{}{}, nil
	} else {
		return []interface{}{entity}, nil
	}
}
