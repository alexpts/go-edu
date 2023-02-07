package home

import (
	"encoding/json"
	"github.com/alexpts/edu-go/internal/repo"
	"github.com/alexpts/go-next/next/layer"
	"github.com/rs/zerolog"
)

type ControllerHome struct {
	Logger   *zerolog.Logger
	PostRepo *repo.Post
}

// @todo need Response less syntax

func (c *ControllerHome) ActionIndex(ctx *layer.HandlerCtx) {
	post := c.PostRepo.FindById(2)
	ctx.Response.Header.Add("content-type", "application/json") // @todo move to middleware for API application

	if post == nil {
		ctx.Response.SetStatusCode(404)
		ctx.Response.AppendBodyString(`{"error": "not found""}`)
		return
	}

	respBytes, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}

	ctx.Response.SetStatusCode(200)
	ctx.Response.AppendBody(respBytes)
}
