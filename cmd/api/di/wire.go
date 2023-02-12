//go:build wireinject
// +build wireinject

package di

import (
	"github.com/alexpts/go-next/next"
	"github.com/bytedance/sonic"
	"github.com/google/wire"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"

	"github.com/alexpts/edu-go/internal/controller"
	"github.com/alexpts/edu-go/internal/middleware"
	"github.com/alexpts/edu-go/internal/provider"
	"github.com/alexpts/edu-go/internal/repo"
	"github.com/alexpts/edu-go/pkg/convert"
)

var repoSet = wire.NewSet(
	wire.Struct(new(repo.Post), "Db"),
	wire.Struct(new(repo.User), "Db"),
)

var controllerSet = wire.NewSet(
	wire.Value(controller.NotFound{
		Payload: []byte(`{"error": "not found"}`),
	}),
	wire.Struct(new(controller.RestController), "*"),
	wire.Struct(new(controller.Home), "*"),
	wire.Struct(new(controller.User), "*"),
	wire.Struct(new(controller.Post), "*"),
)

var middlewareSet = wire.NewSet(
	middleware.ProvideMiddlewarePanic,
)

var marshalerSet = wire.NewSet(
	provider.ProvideStdEncodingJsonMarshaler,

	// select sonic instance with any configs
	provider.ProvideFastestSonicJsonMarshaler,
	//provider.ProvideStdSonicJsonMarshaler,
	//provider.ProvideDefaultSonicJsonMarshaler,
)

func InjectHttpServer(handler fasthttp.RequestHandler) fasthttp.Server {
	wire.Build(provider.ProvideHttpServer)
	return fasthttp.Server{}
}

func InjectApp() (next.MicroApp, error) {
	wire.Build(
		repoSet,
		controllerSet,
		middlewareSet,
		marshalerSet,

		// bind interface (left) to implementation (right)
		//wire.Bind(new(convert.IJsonMarshaler), new(*convert.StdMarshaler)),
		wire.Bind(new(convert.IJsonMarshaler), new(sonic.API)), // without ref because sonic.API is interface type

		// other
		provider.ProvideConfig,
		provider.ProvideZeroLogger,
		provider.ProvideNextLayers,
		provider.ProvideNextApp,
		provider.ProvideDbConnect,
		provider.ProvideGormDb,
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
