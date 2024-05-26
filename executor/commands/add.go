package commands

import (
	"fmt"
	"nolono-cli/domain"
	"nolono-cli/executor/collections"
	"nolono-cli/executor/commands/errors"
	"nolono-cli/executor/utils"
	"nolono-cli/storage"
	"strconv"
	"strings"
	"time"
)

// example calls:
//
// > add tx (adam;20)
// > add src (adam)
func Add(args [3]string) [1]string {
	collection := args[1]
	values := strings.Split(
		args[2][strings.Index(args[2], "(")+1:strings.Index(args[2], ")")],
		";",
	)

	switch collection {
	case collections.SRC:
		{
			if len(values) != 1 {
				return utils.FResult(
					invalidValuesAmount(collection, 1, len(values)),
				)
			}

			storage.Src[values[0]] = true

			return utils.FResult(values[0])
		}
	case collections.TX:
		{
			if len(values) != 2 {
				return utils.FResult(
					invalidValuesAmount(collection, 2, len(values)),
				)
			}

			source := values[0]
			amount, err := strconv.ParseFloat(values[1], 32)

			if err != nil {
				return utils.FResult(invalidElementForTx("amount", amount))
			}

			if !storage.Src[source] {
				return utils.FResult(invalidElementForTx("source", source))
			}

			id := strconv.FormatInt(time.Now().UnixMicro(), 10)
			storage.Tx[id] = domain.NewTx(source, float32(amount))

			return utils.FResult(storage.Tx[id].ToStr(id))
		}
	default:
		{
			return utils.FResult(errors.InvalidCollection(collection))
		}
	}
}

func invalidValuesAmount(collection string, expected int, received int) string {
	return fmt.Sprintf(
		"Invalid values amount for '%s' collection (expected: %d received: %d)",
		collection,
		expected,
		received,
	)
}

func invalidElementForTx(elem string, value any) string {
	return fmt.Sprintf("Invalid %s for tx (%s)", elem, value)
}