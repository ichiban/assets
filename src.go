package assets

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Src locates resources by path of the caller of assets.New.
// Let's assume assets.New is called in /foo/bar/baz.go and dirName is assets, it first tries /foo/bar/assets, then /foo/assets, then /assets.
// If one of them exists, it returns it.
// Otherwise, it fails.
func Src(dirName string) Strategy {
	return &srcStrategy{dirName: dirName}
}

type srcStrategy struct {
	dirName string
}

func (s *srcStrategy) Path() (string, error) {
	_, file, _, ok := caller(2)
	if !ok {
		return "", errors.New("failed to identify caller")
	}

	for dir := filepath.Dir(file); !strings.HasSuffix(dir, string(filepath.Separator)); dir = filepath.Dir(dir) {
		path := filepath.Join(dir, s.dirName)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("no directory as '%s' found in ancestor directories of '%s'", s.dirName, file)
}

func (s *srcStrategy) Close() error {
	return nil
}

var caller = runtime.Caller
