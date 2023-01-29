package panic

import (
	"github.com/alexpts/go-next/next"
	"github.com/rs/zerolog"
)

type MiddlewarePanic struct {
	logger *zerolog.Logger
}

func ProvideMiddlewarePanic(logger *zerolog.Logger) MiddlewarePanic {
	return MiddlewarePanic{logger}
}

func (m *MiddlewarePanic) convertPanicToResponse(recovery any, ctx *next.HandlerCxt) {
	switch recovery := recovery.(type) {
	case nil:
		return
	case next.PanicMessage:
		m.logger.Error().Err(recovery.Error).Msg("convert panic to response")
	default:
		m.logger.Panic().Any("panic", recovery)
	}

	ctx.Response.SetStatusCode(500)
	ctx.Response.AppendBodyString(`{"error": "server error"}`)
}

func (m *MiddlewarePanic) Middleware(ctx *next.HandlerCxt) {
	defer func() {
		m.convertPanicToResponse(recover(), ctx)
	}()
	ctx.Next()
}
