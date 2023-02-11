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
	home controller.Home,
	user controller.User,
	post controller.Post,
	notFoundController controller.NotFound,
	panicMiddleware middleware.PanicMiddleware,
) []layer.Layer {
	return []layer.Layer{
		{
			Name:     `main-page`,
			Path:     `/`,
			Handlers: m(home.ActionPost),
		},
		{
			Name:     `user-by-id`,
			Path:     `/users/{id:\d+}/`,
			Handlers: m(user.ActionGet),
		},
		{
			Name:     `users`,
			Path:     `/users/`,
			Handlers: m(user.ActionFind),
		},
		{
			Name:     `post-by-id`,
			Path:     `/posts/{id:\d+}/`,
			Handlers: m(post.ActionGet),
		},
		{
			Name:     `posts`,
			Path:     `/posts/`,
			Handlers: m(post.ActionFind),
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
