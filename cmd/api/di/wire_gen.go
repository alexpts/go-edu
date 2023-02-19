// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/alexpts/edu-go/internal/controller"
	"github.com/alexpts/edu-go/internal/middleware"
	panic2 "github.com/alexpts/edu-go/internal/middleware/panic"
	"github.com/alexpts/edu-go/internal/provider"
	"github.com/alexpts/go-next/next"
	"github.com/google/wire"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
)

// Injectors from wire.go:

func InjectHttpServer(handler fasthttp.RequestHandler) fasthttp.Server {
	server := provider.ProvideHttpServer(handler)
	return server
}

func InjectApp() (next.MicroApp, error) {
	api := provider.ProvideFastestSonicJsonMarshaler()
	restController := controller.RestController{
		Json: api,
	}
	logger := provider.ProvideZeroLogger()
	home := controller.Home{
		RestController: restController,
		Logger:         logger,
	}
	viper := provider.ProvideConfig()
	db, err := provider.ProvideDbConnect(viper)
	if err != nil {
		return next.MicroApp{}, err
	}
	gormDB, err := provider.ProvideGormDb(db, logger)
	if err != nil {
		return next.MicroApp{}, err
	}
	user := provider.ProvideUserRepo(gormDB)
	controllerUser := controller.User{
		RestController: restController,
		Logger:         logger,
		UserRepo:       user,
	}
	post := provider.ProvidePostRepo(gormDB)
	controllerPost := controller.Post{
		RestController: restController,
		Logger:         logger,
		Repo:           post,
	}
	notFound := _wireNotFoundValue
	middlewarePanic := panic2.ProvideMiddlewarePanic(logger)
	v := provider.ProvideNextLayers(home, controllerUser, controllerPost, notFound, middlewarePanic)
	microApp := provider.ProvideNextApp(v)
	return microApp, nil
}

var (
	_wireNotFoundValue = controller.NotFound{
		Payload: []byte(`{"error": "not found handler"}`),
	}
)

func InjectLogger() *zerolog.Logger {
	logger := provider.ProvideZeroLogger()
	return logger
}

// wire.go:

var repoSet = wire.NewSet(provider.ProvideUserRepo, provider.ProvidePostRepo)

var controllerSet = wire.NewSet(wire.Value(controller.NotFound{
	Payload: []byte(`{"error": "not found handler"}`),
}), wire.Struct(new(controller.RestController), "*"), wire.Struct(new(controller.Home), "*"), wire.Struct(new(controller.User), "*"), wire.Struct(new(controller.Post), "*"),
)

var middlewareSet = wire.NewSet(middleware.ProvideMiddlewarePanic)

var marshalerSet = wire.NewSet(provider.ProvideStdEncodingJsonMarshaler, provider.ProvideFastestSonicJsonMarshaler)

func InjectApiLogger() *zerolog.Logger {
	logger := provider.ProvideZeroLogger().
		With().Str("app", "api").Logger()

	return &logger
}
