package commands

import (
	"errors"
	"fmt"
	"nolono-cli/executor/collections"
	"nolono-cli/executor/commands/internal_errors"
	"nolono-cli/storage"
)

func Del(args []string) (string, error) {
	collection := args[1]
	target := args[2]

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

	return nullStoragePointer(target), nil
}

func delTx(target string) (string, error) {
	if _, ok := storage.Tx[target]; !ok {
		return "", errors.New(
			invalidTargetForCollection(target, collections.TX),
		)
	}

	delete(storage.Tx, target)

	return nullStoragePointer(target), nil
}

func nullStoragePointer(source string) string {
	return fmt.Sprintf("%s => null", source)
}

func invalidTargetForCollection(target string, collection string) string {
	return fmt.Sprintf(
		"invalid target for collection '%s' (%s)",
		collection,
		target,
	)
}
