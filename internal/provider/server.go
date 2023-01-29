package provider

import (
	"github.com/alexpts/go-next/next"
	"github.com/valyala/fasthttp"
)

func ProvideHttpServer(handler fasthttp.RequestHandler) fasthttp.Server {
	return fasthttp.Server{
		Handler:               handler,
		NoDefaultDate:         true,
		NoDefaultContentType:  true,
		NoDefaultServerHeader: true,
		TCPKeepalive:          true,
		//GetOnly:                       true,
		//DisableHeaderNamesNormalizing: true,
	}
}

func ProvideNextApp(layers []next.Layer) next.App {
	app := next.NewApp()

	for i := range layers {
		app.AddLayer(&layers[i])
	}

	return app
}
