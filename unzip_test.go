package assets

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestUnzip(t *testing.T) {
	t.Run("bundled", func(t *testing.T) {
		assert := assert.New(t)

		var g *monkey.PatchGuard
		g = monkey.Patch(ioutil.ReadFile, func(string) ([]byte, error) {
			g.Unpatch()
			defer g.Restore()

			return ioutil.ReadFile("testdata/bin/hello-bundled")
		})
		defer g.Unpatch()

		s := Unzip()
		p, err := s.Path()
		assert.NoError(err)
		assert.True(strings.HasPrefix(p, os.TempDir()))

		_, err = os.Stat(p)
		assert.NoError(err)

		assert.NoError(s.Close())

		_, err = os.Stat(p)
		assert.Error(err)
	})

	t.Run("zipslip", func(t *testing.T) {
		assert := assert.New(t)

		var g *monkey.PatchGuard
		g = monkey.Patch(ioutil.ReadFile, func(string) ([]byte, error) {
			g.Unpatch()
			defer g.Restore()

			return ioutil.ReadFile("testdata/bin/hello-zipslip")
		})
		defer g.Unpatch()

		s := Unzip()
		_, err := s.Path()
		assert.Error(err)
	})

	t.Run("non-zip", func(t *testing.T) {
		assert := assert.New(t)

		var g *monkey.PatchGuard
		g = monkey.Patch(ioutil.ReadFile, func(string) ([]byte, error) {
			g.Unpatch()
			defer g.Restore()

			return ioutil.ReadFile("testdata/bin/hello")
		})
		defer g.Unpatch()

		s := Unzip()
		_, err := s.Path()
		assert.Error(err)
	})
}
