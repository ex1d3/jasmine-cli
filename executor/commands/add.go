package commands

import (
	"errors"
	"fmt"
	"nolono-cli/domain"
	"nolono-cli/executor/collections"
	"nolono-cli/executor/commands/internal_errors"
	"nolono-cli/storage"
	"strconv"
	"strings"
	"time"
)

// example calls:
//
// > add tx (adam;20)
// > add src (adam)
func Add(
	// args param length is actually gonna be pre-checked in #ExecuteCall() func, so
	// nevermind about this unspecified slice size
	args []string,
) (string, error) {
	collection := args[1]
	rawValues := args[2]

	openBracketIndex := strings.Index(rawValues, "(")
	closeBracketIndex := strings.Index(rawValues, ")")

	if (openBracketIndex == -1) || (closeBracketIndex == -1) {
		return "", errors.New(
			"object constructor not found",
		)
	}

	values := strings.Split(
		rawValues[openBracketIndex+1:closeBracketIndex],
		";",
	)

	if len(values) == 1 && values[0] == "" {
		return "", errors.New("object constructor is empty")
	}

	switch collection {
	case collections.SRC:
		{
			return addSrc(values)
		}
	case collections.TX:
		{
			return addTx(values)
		}
	default:
		{
			return "", errors.New(
				internal_errors.InvalidCollection(collection),
			)
		}
	}
}

func addSrc(values []string) (string, error) {
	if len(values) != 1 {
		return "", errors.New(
			invalidValuesAmount(collections.SRC, 1, len(values)),
		)
	}

	name := values[0]
	storage.Src[name] = true

	return name, nil
}

func addTx(values []string) (string, error) {
	if len(values) != 2 {
		return "", errors.New(
			invalidValuesAmount(collections.TX, 2, len(values)),
		)
	}

	source := values[0]
	amount, err := strconv.ParseFloat(values[1], 32)

	if err != nil {
		return "", errors.New(invalidElementForTx("amount", amount))
	}

	if !storage.Src[source] {
		return "", errors.New(invalidElementForTx("source", source))
	}

	id := strconv.FormatInt(time.Now().UnixMicro(), 10)
	storage.Tx[id] = domain.NewTx(source, float32(amount))

	return storage.Tx[id].ToStr(id), nil
}

func invalidValuesAmount(collection string, expected int, received int) string {
	return fmt.Sprintf(
		"invalid values amount for '%s' collection (expected: %d received: %d)",
		collection,
		expected,
		received,
	)
}

func invalidElementForTx(elem string, value any) string {
	return fmt.Sprintf("invalid %s for tx (%s)", elem, value)
}
