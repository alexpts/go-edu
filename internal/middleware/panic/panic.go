package panic

import (
	"fmt"
	"github.com/alexpts/go-next/next/layer"
	"github.com/rs/zerolog"
)

type MiddlewarePanic struct {
	logger *zerolog.Logger
}

func ProvideMiddlewarePanic(logger *zerolog.Logger) MiddlewarePanic {
	return MiddlewarePanic{logger}
}

func (m *MiddlewarePanic) convertPanicToResponse(throw any, ctx *layer.HandlerCtx) {
	m.logger.WithLevel(zerolog.PanicLevel).Str("panic", fmt.Sprintf("%v", throw)).Send()

	ctx.Response.SetStatusCode(500)
	ctx.SetContentType("application/json")
	ctx.Response.AppendBodyString(`{"error": "server error"}`)
}

func (m *MiddlewarePanic) Middleware(ctx *layer.HandlerCtx) {
	defer func() {
		throw := recover()
		if throw != nil {
			m.convertPanicToResponse(throw, ctx)
		}
	}()

	ctx.Next()
}
