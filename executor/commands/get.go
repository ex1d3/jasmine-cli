package commands

import (
	"errors"
	"nolono-cli/executor/collections"
	"nolono-cli/executor/commands/internal_errors"
	"nolono-cli/storage"
)

func Get(args []string) ([]string, error) {
	collection := args[0]
	params := GetParams{}

	argsLen := len(args)

	if argsLen < 1 || argsLen > 2 {
		return []string{}, errors.New(
			internal_errors.InvalidArgsCount("add", "2", len(args)),
		)
	}

	if len(args) == 2 {
		params.Target = args[1]
	}

	switch collection {
	case collections.SRC:
		{
			return getSrc(params)
		}
	case collections.TX:
		{
			return getTx(params)
		}
	default:
		{
			internal_errors.InvalidCollection(collection)
		}
	}

	return []string{}, nil
}

type GetParams struct {
	Target string
}

func getSrc(params GetParams) ([]string, error) {
	if params.Target == "" {
		srcs := make([]string, len(storage.Src))
		i := 0

		for k := range storage.Src {
			srcs[i] = k
			i++
		}

		return srcs, nil
	}

	if !storage.Src[params.Target] {
		return []string{}, nil
	}

	return []string{params.Target}, nil
}

func getTx(params GetParams) ([]string, error) {
	if params.Target == "" {
		txs := make([]string, len(storage.Tx))
		i := 0

		for k, v := range storage.Tx {
			txs[i] = v.ToStr(k)
			i++
		}

		return txs, nil
	}

	if tx, ok := storage.Tx[params.Target]; !ok {
		return []string{}, nil
	} else {
		return []string{tx.ToStr(params.Target)}, nil
	}
}
