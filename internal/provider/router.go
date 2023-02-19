package provider

import (
	"github.com/alexpts/go-next/next/layer"

	"github.com/alexpts/edu-go/internal/controller"
	"github.com/alexpts/edu-go/internal/middleware"
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
			Path:     `/`,
			Handlers: m(home.ActionPost),
		},
		{
			Name:     `user-by-id`,
			Path:     `/users/{id:\d+}/`,
			Methods:  []string{`GET`},
			Handlers: m(user.ActionGet),
		},
		{
			Name:     `user-by-name`,
			Path:     `/users/{name:[a-z][a-z0-9]*}/`,
			Handlers: m(user.ActionGetByName),
		},
		{
			Name:     `get users`,
			Path:     `/users/`,
			Methods:  []string{`GET`},
			Handlers: m(user.ActionFind),
		},
		{
			Name:     `create user`,
			Path:     `/users/`,
			Methods:  []string{`POST`},
			Handlers: m(user.ActionCreate),
		},
		{
			Name:     `update user`,
			Path:     `/users/{id:\d+}/`,
			Methods:  []string{`PUT`},
			Handlers: m(user.ActionUpdate),
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
			Handlers: m(panicMiddleware.Middleware),
			Priority: 1000,
		},
		{
			Name:     `otherwise`,
			Handlers: m(notFoundController.Action404),
			Priority: -1000,
		},
	}
}
