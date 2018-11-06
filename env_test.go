package assets

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	t.Run("environment variable is not defined", func(t *testing.T) {
		assert := assert.New(t)

		assert.NoError(os.Setenv("key", ""))
		s := Env("key")
		_, err := s.Path()
		assert.Error(err)
	})

	t.Run("environment variable is defined", func(t *testing.T) {
		assert := assert.New(t)

		assert.NoError(os.Setenv("key", "/foo/bar"))
		s := Env("key")
		path, err := s.Path()
		assert.NoError(err)
		assert.Equal("/foo/bar", path)
		assert.NoError(s.Close())
	})
}
