package assets

import (
	"fmt"
	"os"
)

func Env(key string) Strategy {
	return &envStrategy{key: key}
}

type envStrategy struct {
	key string
}

func (s *envStrategy) Path() (string, error) {
	path := os.Getenv(s.key)
	if path == "" {
		return "", fmt.Errorf("env var '%s' is not set", s.key)
	}
	return path, nil
}

func (s *envStrategy) Close() error {
	return nil
}
