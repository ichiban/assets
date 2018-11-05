package assets

import (
	"path/filepath"
	"runtime"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestSrc(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		assert := assert.New(t)

		g := monkey.Patch(runtime.Caller, func(skip int) (pc uintptr, file string, line int, ok bool) {
			path, err := filepath.Abs("testdata/main.go")
			assert.NoError(err)
			return uintptr(0), path, 0, true
		})
		defer g.Unpatch()

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

		g := monkey.Patch(runtime.Caller, func(skip int) (pc uintptr, file string, line int, ok bool) {
			path, err := filepath.Abs("testdata/main.go")
			assert.NoError(err)
			return uintptr(0), path, 0, true
		})
		defer g.Unpatch()

		s := Src("I hope you don't have this weirdly named directory")
		_, err := s.Path()
		assert.Error(err)
	})
}
