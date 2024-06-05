package commands

import (
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
		return "", &internal_errors.InvalidArgsCountError{
			Command: "add",
			Want:    2,
			Have:    len(args),
		}
	}

	collection := args[0]
	rawValues := args[1]

	openBracketIndex := strings.Index(rawValues, "(")
	closeBracketIndex := strings.Index(rawValues, ")")

	if (openBracketIndex == -1) || (closeBracketIndex == -1) {
		return "", &internal_errors.ObjectConstructorNotFoundError{}
	}

	values := strings.Split(
		rawValues[openBracketIndex+1:closeBracketIndex],
		";",
	)

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
			return "", &internal_errors.InvalidCollectionError{
				Collection: collection,
			}
		}
	}
}

func addSrc(values []string) (string, error) {
	if len(values) != 1 {
		return "", &internal_errors.InvalidObjectConstructorLengthError{
			Have: len(values),
			Want: 1,
		}
	}

	if values[0] == "" {
		return "", &internal_errors.InvalidObjectConstructorLengthError{
			Have: 0,
			Want: 1,
		}
	}

	name := values[0]

	storage.Src.Set(name, domain.NewSrc(name))

	return name, nil
}

func addTx(values []string) (string, error) {
	if len(values) != 2 {
		return "", &internal_errors.InvalidObjectConstructorLengthError{
			Have: len(values),
			Want: 2,
		}
	}

	source := values[0]

	if storage.Src.Get(source) == nil {
		return "", &internal_errors.InvalidObjectConstructorValueError{
			Index: 1,
			Value: source,
		}
	}

	amount, err := strconv.ParseFloat(values[1], 32)

	if err != nil {
		return "", &internal_errors.InvalidObjectConstructorValueError{
			Index: 1,
			Value: values[1],
		}
	}

	txStorage := storage.Tx

	id := strconv.Itoa(len(txStorage.GetStorage()) + 1)
	txStorage.Set(id, domain.NewTx(id, source, float32(amount)))

	return txStorage.Get(id).ToStr(), nil
}
