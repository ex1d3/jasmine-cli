package commands

import (
	"fmt"
	"nolono-cli/executor/collections"
	"nolono-cli/executor/utils"
	"nolono-cli/storage"
)

func Del(args []string) [1]string {
	collection := args[1]
	target := args[2]

	switch collection {
	case collections.SRC:
		{
			if !storage.Src[target] {
				return utils.FResult(
					invalidTargetForCollection(target, collection),
				)
			}

			delete(storage.Src, target)

			break
		}
	case collections.TX:
		{
			if _, ok := storage.Tx[target]; !ok {
				return utils.FResult(
					invalidTargetForCollection(target, collection),
				)
			}

			delete(storage.Tx, target)
		}
	}

	return utils.FResult(nullStoragePointer(target))
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
