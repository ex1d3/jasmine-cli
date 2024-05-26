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
			if !storage.Src[target] {
				return "", errors.New(
					invalidTargetForCollection(target, collection),
				)
			}

			delete(storage.Src, target)

			break
		}
	case collections.TX:
		{
			if _, ok := storage.Tx[target]; !ok {
				return "", errors.New(
					invalidTargetForCollection(target, collection),
				)
			}

			delete(storage.Tx, target)
		}
	default:
		{
			return "", errors.New(
				internal_errors.InvalidCollection(collection),
			)
		}
	}

	return nullStoragePointer(target), nil
}

func nullStoragePointer(source string) string {
	return fmt.Sprintf("%s => null", source)
}

func invalidTargetForCollection(target string, collection string) string {
	return fmt.Sprintf(
		"Invalid target for collection '%s' (%s)",
		collection,
		target,
	)
}
