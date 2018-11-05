package assets

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNew(t *testing.T) {
	t.Run("if the 1st strategy succeeds, it employs the 1st strategy", func(t *testing.T) {
		assert := assert.New(t)

		var s1 MockStrategy
		s1.On("Path").Return("/foo/bar", nil).Once()
		s1.On("Close").Return(nil).Once()
		var s2 MockStrategy

		l, err := New(&s1, &s2)
		assert.NoError(err)
		assert.Equal("/foo/bar", l.Path)
		assert.NoError(l.Close())

		s1.AssertExpectations(t)
		s2.AssertExpectations(t)
	})

	t.Run("if the 2nd strategy succeeds, it employs the 2nd strategy", func(t *testing.T) {
		assert := assert.New(t)

		var s1 MockStrategy
		s1.On("Path").Return("", errors.New("failed")).Once()
		var s2 MockStrategy
		s2.On("Path").Return("/foo/baz", nil).Once()
		s2.On("Close").Return(nil).Once()

		l, err := New(&s1, &s2)
		assert.NoError(err)
		assert.Equal("/foo/baz", l.Path)
		assert.NoError(l.Close())

		s1.AssertExpectations(t)
		s2.AssertExpectations(t)
	})

	t.Run("if no strategies succeed, it returns error", func(t *testing.T) {
		assert := assert.New(t)

		var s1 MockStrategy
		s1.On("Path").Return("", errors.New("failed")).Once()
		var s2 MockStrategy
		s2.On("Path").Return("", errors.New("failed")).Once()

		_, err := New(&s1, &s2)
		assert.Error(err)

		s1.AssertExpectations(t)
		s2.AssertExpectations(t)
	})

	t.Run("without arguments, it employs default strategies", func(t *testing.T) {
		assert := assert.New(t)

		defaults := DefaultStrategies
		defer func() {
			DefaultStrategies = defaults
		}()

		var s1 MockStrategy
		s1.On("Path").Return("/foo/bar", nil).Once()
		s1.On("Close").Return(nil).Once()
		var s2 MockStrategy

		DefaultStrategies = []Strategy{&s1, &s2}

		l, err := New()
		assert.NoError(err)
		assert.Equal("/foo/bar", l.Path)
		assert.NoError(l.Close())

		s1.AssertExpectations(t)
		s2.AssertExpectations(t)
	})
}

type MockStrategy struct {
	mock.Mock
}

func (s *MockStrategy) Path() (string, error) {
	args := s.Called()
	return args.String(0), args.Error(1)
}

func (s *MockStrategy) Close() error {
	args := s.Called()
	return args.Error(0)
}
