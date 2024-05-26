package storage

var Src = make(map[string]bool)

func ClearSrc() {
	for k := range Src {
		delete(Src, k)
	}
}
