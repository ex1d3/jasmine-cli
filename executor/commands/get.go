package commands

import (
	"errors"
	"nolono-cli/executor/collections"
	"nolono-cli/executor/commands/internal_errors"
	"nolono-cli/storage"
)

func Get(args []string) ([]string, error) {
	params := GetParams{}

	if len(args) != 1 && len(args) != 2 {
		return []string{}, errors.New(
			internal_errors.InvalidArgsCount("get", "2", len(args)),
		)
	}

	if args[0] == " " {
		return []string{}, errors.New(
			internal_errors.InvalidArgsCount("get", "2", len(args)-1),
		)
	}

	collection := args[0]

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
			return []string{}, errors.New(internal_errors.InvalidCollection(collection))
		}
	}
}

type GetParams struct {
	Target string
}

func getSrc(params GetParams) ([]string, error) {
	srcStorage := storage.Src.GetStorage()

	if params.Target == "" {
		srcs := make([]string, len(srcStorage))
		i := 0

		for k, v := range srcStorage {
			if v {
				srcs[i] = k
			}

			i++
		}

		return srcs, nil
	}

	if !storage.Src.Get(params.Target) {
		return []string{}, nil
	}

	return []string{params.Target}, nil
}

func getTx(params GetParams) ([]string, error) {
	txStorage := storage.Tx

	if params.Target == "" {
		txs := make([]string, len(txStorage.GetStorage()))
		i := 0

		for k, v := range txStorage.GetStorage() {
			txs[i] = v.ToStr(k)
			i++
		}

		return txs, nil
	}

	if tx := txStorage.Get(params.Target); tx == nil {
		return []string{}, nil
	} else {
		return []string{tx.ToStr(params.Target)}, nil
	}
}
