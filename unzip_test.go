package assets

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnzip(t *testing.T) {
	t.Run("bundled", func(t *testing.T) {
		assert := assert.New(t)

		e := executable
		defer func() { executable = e }()
		executable = func() (string, error) {
			return "testdata/bin/hello-bundled", nil
		}

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

		e := executable
		defer func() { executable = e }()
		executable = func() (string, error) {
			return "testdata/bin/hello-zipslip", nil
		}

		s := Unzip()
		_, err := s.Path()
		assert.Error(err)
	})

	t.Run("non-zip", func(t *testing.T) {
		assert := assert.New(t)

		e := executable
		defer func() { executable = e }()
		executable = func() (string, error) {
			return "testdata/bin/hello", nil
		}

		s := Unzip()
		_, err := s.Path()
		assert.Error(err)
	})
}
