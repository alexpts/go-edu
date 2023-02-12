package controller

import (
	"github.com/alexpts/go-next/next/layer"
	"github.com/rs/zerolog"
)

type Home struct {
	RestController
	Logger *zerolog.Logger
}

func (c *Home) ActionPost(ctx *layer.HandlerCtx) {
	c.sendJsonModel(ctx, map[string]any{
		"message": "main page",
	})
}
