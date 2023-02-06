package provider

import (
	"github.com/alexpts/edu-go/internal/controller"
	"github.com/alexpts/edu-go/internal/middleware"
	"github.com/alexpts/go-next/next/layer"
)

// m - multi handler wrapper for decomposition
func m(handlers ...layer.Handler) []layer.Handler {
	return handlers
}

func ProvideNextLayers(
	homeController controller.Home,
	notFoundController controller.NotFound,
	panicMiddleware middleware.PanicMiddleware,
) []layer.Layer {
	return []layer.Layer{
		{
			Name:     `main-page`,
			Path:     `/`,
			Handlers: m(homeController.ActionIndex),
		},
		{
			Name:     `panic-to-response`,
			Handlers: m(panicMiddleware.Middleware),
			Priority: 1000,
		},
		{
			Name:     `otherwise`,
			Handlers: m(notFoundController.Action404),
			Priority: -100,
		},
	}
}
