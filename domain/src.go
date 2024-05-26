package domain

type Src struct {
	Name string `json:"name"`
}

func NewSrc(name string) *Src {
	return &Src{Name: name}
}
