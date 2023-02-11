package controller

import (
	"github.com/alexpts/go-next/next/layer"
)

type NotFound struct {
	Payload []byte
}

func (c *NotFound) Action404(ctx *layer.HandlerCtx) {
	ctx.Response.SetStatusCode(404)
	ctx.SetContentType("application/json")
	ctx.Response.AppendBody(c.Payload)
}
