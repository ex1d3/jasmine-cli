package commands

import (
	"errors"
	"fmt"
	"jasmine-cli/domain"
	"jasmine-cli/executor/collections"
	"jasmine-cli/executor/commands/internal_errors"
	"jasmine-cli/storage"
)

// Del example args slices: [tx 1] or [src adam]
func Del(args []string) ([]interface{}, error) {
	if len(args) != 2 {
		return []interface{}{}, errors.New(
			internal_errors.InvalidArgsCount("del", "2", len(args)),
		)
	}

	collection := args[0]
	target := args[1]

	switch collection {
	case collections.SRC:
		{
			return executeDel[domain.Src](target, storage.Src)
		}
	case collections.TX:
		{
			return executeDel[domain.Tx](target, storage.Tx)
		}
	default:
		{
			return []interface{}{}, errors.New(
				internal_errors.InvalidCollection(collection),
			)
		}
	}
}

func executeDel[T any](
	target string,
	entityStorage storage.Storage[string, *T],
) ([]interface{}, error) {
	if target == "*" {
		entityStorage.Unload()

		return []interface{}{}, nil
	}

	if entity := entityStorage.Get(target); entity == nil {
		return []interface{}{}, errors.New(
			invalidTargetForCollection(target, collections.SRC),
		)
	} else {
		entityStorage.Delete(target)
		return []interface{}{entity}, nil
	}

}

func invalidTargetForCollection(target string, collection string) string {
	return fmt.Sprintf(
		"invalid target for collection '%s' (%s)",
		collection,
		target,
	)
}
