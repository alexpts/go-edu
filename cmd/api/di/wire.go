//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
)

type A struct {
	Name string
}

func provideA() (A, func()) {
	return A{"Alex"}, func() {}
}

func InjectA() (A, func(), error) {
	wire.Build(
		provideA,
	)

	return A{}, nil, nil
}
