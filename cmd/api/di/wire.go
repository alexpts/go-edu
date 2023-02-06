//go:build wireinject
// +build wireinject

package di

import (
	"github.com/alexpts/edu-go/internal/controller"
	"github.com/alexpts/edu-go/internal/middleware"
	"github.com/alexpts/edu-go/internal/provider"
	"github.com/alexpts/go-next/next"
	"github.com/google/wire"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
)

//var miscSet = wire.NewSet(
//	provider.ProviderZeroLogger,
//)

func InjectHttpServer(handler fasthttp.RequestHandler) fasthttp.Server {
	wire.Build(provider.ProvideHttpServer)
	return fasthttp.Server{}
}

func InjectApp() next.MicroApp {
	wire.Build(
		// controllers
		wire.Struct(new(controller.Home), "Logger"),
		wire.Value(controller.NotFound{
			Payload: []byte(`{"error": "not found"}`),
		}),
		middleware.ProvideMiddlewarePanic,

		provider.ProvideZeroLogger,
		provider.ProvideNextLayers,
		provider.ProvideNextApp,
	)

	return next.MicroApp{}
}

func InjectLogger() *zerolog.Logger {
	panic(wire.Build(provider.ProvideZeroLogger))
}

func InjectApiLogger() *zerolog.Logger {
	logger := provider.ProvideZeroLogger().
		With().Str("app", "api").Logger()

	return &logger
}
