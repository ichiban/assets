package assets

import (
	"fmt"
	"os"
)

// Env locates resources by environment variable.
// If the environment variable identified by the key has non-empty string value, it uses the value as the root path to resources.
// Otherwise, it fails.
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
