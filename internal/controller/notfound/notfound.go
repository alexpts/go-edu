package notfound

import (
	"github.com/alexpts/go-next/next/layer"
)

type ControllerNotFound struct {
	Payload []byte
}

func (c *ControllerNotFound) Action404(ctx *layer.HandlerCtx) {
	ctx.Response.SetStatusCode(404)
	ctx.SetContentType("application/json")
	ctx.Response.AppendBody(c.Payload)
}
