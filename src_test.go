package assets

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSrc(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		assert := assert.New(t)

		c := caller
		defer func() { caller = c }()
		caller = func(skip int) (pc uintptr, file string, line int, ok bool) {
			path, err := filepath.Abs("testdata/main.go")
			assert.NoError(err)
			return uintptr(0), path, 0, true
		}

		s := Src("assets")
		path, err := s.Path()
		assert.NoError(err)

		assetsPath, err := filepath.Abs("testdata/assets")
		assert.NoError(err)
		assert.Equal(assetsPath, path)

		assert.NoError(s.Close())
	})

	t.Run("not found", func(t *testing.T) {
		assert := assert.New(t)

		c := caller
		defer func() { caller = c }()
		caller = func(skip int) (pc uintptr, file string, line int, ok bool) {
			path, err := filepath.Abs("testdata/main.go")
			assert.NoError(err)
			return uintptr(0), path, 0, true
		}

		s := Src("I hope you don't have this weirdly named directory")
		_, err := s.Path()
		assert.Error(err)
	})
}
