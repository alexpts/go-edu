//go:build wireinject
// +build wireinject

package di

import (
	"github.com/alexpts/edu-go/internal/provider"
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

func InjectLogger() zerolog.Logger {
	panic(wire.Build(provider.ProvideZeroLogger))
}

func InjectApiLogger() zerolog.Logger {
	logger := provider.ProvideZeroLogger()
	return logger.With().Str("app", "api").Logger()
}
