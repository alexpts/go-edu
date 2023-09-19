package controller

import (
	"github.com/alexpts/go-next/next/layer"
	"github.com/rs/zerolog"

	"github.com/alexpts/edu-go/internal/repo"
	"github.com/alexpts/edu-go/pkg/convert"
)

type Post struct {
	RestController
	Logger *zerolog.Logger
	Repo   *repo.Post
}

func (c *Post) ActionGet(ctx *layer.HandlerCtx) {
	userId := ctx.UriParams["id"]
	model, _ := c.Repo.FindOneById(ctx, convert.MustInt(userId), AllRelation)
	c.sendJsonModel(ctx, model)
}

func (c *Post) ActionFind(ctx *layer.HandlerCtx) {
	models, _ := c.Repo.FindAll(ctx, AllRelation)
	c.sendJsonModel(ctx, models)
}
