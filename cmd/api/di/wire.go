//go:build wireinject
// +build wireinject

package di

import (
	"github.com/alexpts/edu-go/internal/controller"
	"github.com/alexpts/edu-go/internal/middleware"
	"github.com/alexpts/edu-go/internal/provider"
	"github.com/alexpts/edu-go/internal/repo"
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

func InjectApp() (next.MicroApp, error) {
	wire.Build(
		// controllers
		wire.Struct(new(controller.Home), "Logger", "PostRepo"),
		wire.Value(controller.NotFound{
			Payload: []byte(`{"error": "not found"}`),
		}),
		middleware.ProvideMiddlewarePanic,

		provider.ProvideConfig,
		provider.ProvideZeroLogger,
		provider.ProvideNextLayers,
		provider.ProvideNextApp,
		provider.ProvideDbConnect,
		provider.ProvideGormDb,

		wire.Struct(new(repo.Post), "Db"),
	)

	return next.MicroApp{}, nil
}

func InjectLogger() *zerolog.Logger {
	panic(wire.Build(provider.ProvideZeroLogger))
}

func InjectApiLogger() *zerolog.Logger {
	logger := provider.ProvideZeroLogger().
		With().Str("app", "api").Logger()

	return &logger
}
