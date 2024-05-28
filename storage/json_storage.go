package storage

import (
	"encoding/json"
	"io"
	"os"
)

type JSONStorage[K string, V any] struct {
	filename string
	storage  map[K]V
}

func NewJSONStorage[K string, V any](filename string) *JSONStorage[K, V] {
	return &JSONStorage[K, V]{
		filename: filename,
		storage:  make(map[K]V),
	}
}

func (jS JSONStorage[K, V]) GetFilename() string {
	return jS.filename
}

func (jS JSONStorage[K, V]) GetStorage() map[K]V {
	return jS.storage
}

func (jS JSONStorage[K, V]) Get(k K) V {
	return jS.storage[k]
}

func (jS JSONStorage[K, V]) Set(k K, v V) {
	jS.storage[k] = v
}

func (jS JSONStorage[K, V]) Delete(k K) {
	delete(jS.storage, k)
}

func (jS JSONStorage[K, V]) Load(filename string) (int, error) {
	_, err := os.Stat(filename)

	if os.IsExist(err) {
		return 0, nil
	}

	file, err := os.Open(filename)

	if err != nil {
		return -1, err
	}

	defer func(file *os.File) {
		file.Close()
	}(file)

	byteValue, _ := io.ReadAll(file)

	json.Unmarshal(byteValue, &jS.storage)

	return len(jS.storage), nil
}

func (jS JSONStorage[K, V]) Unload() (int, error) {
	oldLength := len(jS.storage)

	for k := range jS.storage {
		jS.Delete(k)
	}

	return oldLength, nil
}

func (jS JSONStorage[K, V]) Save(filename string) (int, error) {
	b, err := json.Marshal(jS.storage)

	if err != nil {
		return -1, err
	}

	file, err := os.Create(filename)

	if err != nil {
		return -1, err
	}

	defer func(file *os.File) {
		file.Close()
	}(file)

	file.Write(b)

	return len(jS.storage), nil
}

func (jS JSONStorage[K, V]) Discard() (int, error) {
	return jS.Load(jS.GetFilename())
}
