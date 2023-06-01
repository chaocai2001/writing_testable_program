package untestable

import (
	"errors"
)

type MemoryStorage struct {
	storage map[string]string
}

func NewMemoryStorage() *MemoryStorage {
	storage := make(map[string]string)
	return &MemoryStorage{storage}
}

func (mem *MemoryStorage) RetiveData(token string) (string, error) {
	val, isOK := mem.storage[token]
	if !isOK {
		return "", errors.New("invalid token")
	}
	return val, nil
}

func (mem *MemoryStorage) StoreData(token string, data string) error {
	mem.storage[token] = data
	return nil
}
