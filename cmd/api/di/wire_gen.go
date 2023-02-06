// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/alexpts/edu-go/internal/controller"
	panic2 "github.com/alexpts/edu-go/internal/middleware/panic"
	"github.com/alexpts/edu-go/internal/provider"
	"github.com/alexpts/go-next/next"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
)

// Injectors from wire.go:

func InjectHttpServer(handler fasthttp.RequestHandler) fasthttp.Server {
	server := provider.ProvideHttpServer(handler)
	return server
}

func InjectApp() next.MicroApp {
	logger := provider.ProvideZeroLogger()
	controllerHome := controller.Home{
		Logger: logger,
	}
	controllerNotFound := _wireControllerNotFoundValue
	middlewarePanic := panic2.ProvideMiddlewarePanic(logger)
	v := provider.ProvideNextLayers(controllerHome, controllerNotFound, middlewarePanic)
	microApp := provider.ProvideNextApp(v)
	return microApp
}

var (
	_wireControllerNotFoundValue = controller.NotFound{
		Payload: []byte(`{"error": "not found"}`),
	}
)

func InjectLogger() *zerolog.Logger {
	logger := provider.ProvideZeroLogger()
	return logger
}

// wire.go:

func InjectApiLogger() *zerolog.Logger {
	logger := provider.ProvideZeroLogger().
		With().Str("app", "api").Logger()

	return &logger
}
