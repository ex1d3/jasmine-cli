package domain

import (
	"fmt"
)

type Tx struct {
	Id     string  `json:"id"`
	Source string  `json:"source"`
	Amount float32 `json:"amount"`
}

func NewTx(id string, source string, amount float32) *Tx {
	return &Tx{
		Id:     id,
		Source: source,
		Amount: amount,
	}
}

func (t *Tx) ToStr() string {
	return fmt.Sprintf("%s => %s, %f", t.Id, t.Source, t.Amount)
}
