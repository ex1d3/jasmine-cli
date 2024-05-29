package commands

import (
	"errors"
	"fmt"
	"jasmine-cli/domain"
	"jasmine-cli/executor/collections"
	"jasmine-cli/executor/commands/internal_errors"
	"jasmine-cli/storage"
	"strconv"
	"strings"
)

func Add(
	// example args slices: [tx (adam;20)] or [src (adam)]
	args []string,
) (string, error) {
	if len(args) != 2 {
		return "", errors.New(
			internal_errors.InvalidArgsCount("add", "2", len(args)),
		)
	}

	collection := args[0]
	rawValues := args[1]

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

	storage.Src.Set(name, domain.NewSrc(name))

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

	if storage.Src.Get(source) == nil {
		return "", errors.New(invalidElementForTx("source", source))
	}

	txStorage := storage.Tx

	id := strconv.Itoa(len(txStorage.GetStorage()) + 1)
	txStorage.Set(id, domain.NewTx(source, float32(amount)))

	return txStorage.Get(id).ToStr(id), nil
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
