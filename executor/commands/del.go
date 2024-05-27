package commands

import (
	"errors"
	"fmt"
	"nolono-cli/executor/collections"
	"nolono-cli/executor/commands/internal_errors"
	"nolono-cli/storage"
)

// example args silces: [tx 1] or [src adam]
func Del(args []string) (string, error) {
	if len(args) != 2 {
		return "", errors.New(
			internal_errors.InvalidArgsCount("del", "2", len(args)),
		)
	}

	collection := args[0]
	target := args[1]

	switch collection {
	case collections.SRC:
		{
			return delSrc(target)
		}
	case collections.TX:
		{
			return delTx(target)
		}
	default:
		{
			return "", errors.New(
				internal_errors.InvalidCollection(collection),
			)
		}
	}
}

func delSrc(target string) (string, error) {
	if !storage.Src[target] {
		return "", errors.New(
			invalidTargetForCollection(target, collections.SRC),
		)
	}

	delete(storage.Src, target)

	return storage.NullStoragePointer(target), nil
}

func delTx(target string) (string, error) {
	if _, ok := storage.Tx[target]; !ok {
		return "", errors.New(
			invalidTargetForCollection(target, collections.TX),
		)
	}

	delete(storage.Tx, target)

	return storage.NullStoragePointer(target), nil
}

func invalidTargetForCollection(target string, collection string) string {
	return fmt.Sprintf(
		"invalid target for collection '%s' (%s)",
		collection,
		target,
	)
}
