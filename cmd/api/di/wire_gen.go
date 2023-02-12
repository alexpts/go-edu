// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/alexpts/edu-go/internal/controller"
	"github.com/alexpts/edu-go/internal/middleware"
	panic2 "github.com/alexpts/edu-go/internal/middleware/panic"
	"github.com/alexpts/edu-go/internal/provider"
	"github.com/alexpts/edu-go/internal/repo"
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
	stdMarshaler := provider.ProvideStdEncodingJsonMarshaler()
	restController := controller.RestController{
		Json: stdMarshaler,
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
	gormDB, err := provider.ProvideGormDb(db)
	if err != nil {
		return next.MicroApp{}, err
	}
	user := &repo.User{
		Db: gormDB,
	}
	controllerUser := controller.User{
		RestController: restController,
		Logger:         logger,
		UserRepo:       user,
	}
	post := &repo.Post{
		Db: gormDB,
	}
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
		Payload: []byte(`{"error": "not found"}`),
	}
)

func InjectLogger() *zerolog.Logger {
	logger := provider.ProvideZeroLogger()
	return logger
}

// wire.go:

var repoSet = wire.NewSet(wire.Struct(new(repo.Post), "Db"), wire.Struct(new(repo.User), "Db"))

var controllerSet = wire.NewSet(wire.Value(controller.NotFound{
	Payload: []byte(`{"error": "not found"}`),
}), wire.Struct(new(controller.RestController), "*"), wire.Struct(new(controller.Home), "*"), wire.Struct(new(controller.User), "*"), wire.Struct(new(controller.Post), "*"),
)

var middlewareSet = wire.NewSet(middleware.ProvideMiddlewarePanic)

var marshalerSet = wire.NewSet(provider.ProvideStdEncodingJsonMarshaler, provider.ProvideFastestSonicJsonMarshaler)

func InjectApiLogger() *zerolog.Logger {
	logger := provider.ProvideZeroLogger().
		With().Str("app", "api").Logger()

	return &logger
}
