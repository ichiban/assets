package assets

import (
	"io"
)

type Locator struct {
	io.Closer
	Path string
}

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

type Strategy interface {
	Path() (string, error)
	io.Closer
}

var DefaultStrategies = []Strategy{Env("ASSETS"), Unzip(), Src("assets")}
