package domain

import "fmt"

type Tx struct {
	Source string  `json:"source"`
	Amount float32 `json:"amount"`
}

func NewTx(source string, amount float32) *Tx {
	return &Tx{Source: source, Amount: amount}
}

func (t *Tx) ToStr(id string) string {
	return fmt.Sprintf("%s => %s, %f", id, t.Source, t.Amount)
}
