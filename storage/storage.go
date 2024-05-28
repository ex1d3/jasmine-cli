package storage

type Storage[K string, V any] interface {
	GetFilename() string
	GetStorage() map[K]V

	Get(k K) V
	Set(k K, v V)
	Delete(k K)

	Load(filename string) (int, error)
	Unload() (int, error)
	Save(filename string) (int, error)
	Discard() (int, error)
}
