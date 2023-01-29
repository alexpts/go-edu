package notfound

import (
	"github.com/alexpts/go-next/next"
)

type ControllerNotFound struct {
	Payload []byte
}

func (c *ControllerNotFound) Action404(ctx *next.HandlerCxt) {
	ctx.Response.SetStatusCode(404)
	ctx.Response.AppendBody(c.Payload)
}
