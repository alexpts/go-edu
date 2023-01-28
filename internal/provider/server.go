package provider

import "github.com/valyala/fasthttp"

func ProvideHttpServer(handler fasthttp.RequestHandler) fasthttp.Server {
	return fasthttp.Server{
		Handler:                       handler,
		NoDefaultDate:                 true,
		NoDefaultContentType:          true,
		NoDefaultServerHeader:         true,
		TCPKeepalive:                  true,
		GetOnly:                       true,
		DisableHeaderNamesNormalizing: true,
	}
}
