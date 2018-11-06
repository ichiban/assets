// Package assets provides an environment-agnostic means of locating resources such as stylesheets, scripts, images, and templates.
package assets

import (
	"io"
)

// Locator holds the path to the root directory of resources.
// Caller must close the locator when it's not needed anymore.
// After it had closed, accessing resources under the path is undefined.
type Locator struct {
	io.Closer
	Path string
}

// New creates an instance of locator.
// It tries the given strategies from left to right until one strategy succeeds at locating resources.
// When no strategies couldn't locate resources, it returns non-nil error.
// If no strategies are given, it uses DefaultStrategies.
func New(strategies ...Strategy) (*Locator, error) {
	if len(strategies) == 0 {
		strategies = DefaultStrategies
	}

	var path string
	var err error
	for _, s := range strategies {
		path, err = s.Path()
		if err == nil {
			return &Locator{
				Closer: s,
				Path:   path,
			}, nil
		}
	}

	return nil, err
}

// Strategy represents a means to locate resources.
type Strategy interface {
	io.Closer
	Path() (string, error)
}

// DefaultStrategies are a recommended sequence of strategies.
var DefaultStrategies = []Strategy{Env("ASSETS"), Unzip(), Src("assets")}
