package common

import (
	"bookTrade-backend/conf"
	"io/ioutil"
	"path/filepath"
)

type StorageService interface {
	Store(string, []byte) error
	Fetch(string) ([]byte, error)
}

type defaultStorageService struct {
	path string
}

func NewDefaultStorageService(conf *conf.AppConfig) *defaultStorageService {
	return &defaultStorageService{
		path: conf.StorageConfig.Path,
	}
}

func (s *defaultStorageService) Store(key string, data []byte) error {
	storePath := filepath.Join(s.path, key)
	return ioutil.WriteFile(storePath, data, 0644)
}

func (s *defaultStorageService) Fetch(key string) ([]byte, error) {
	storePath := filepath.Join(s.path, key)
	return ioutil.ReadFile(storePath)
}
