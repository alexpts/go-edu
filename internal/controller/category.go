package controller

import (
	"github.com/alexpts/edu-go/internal/model"
	"github.com/alexpts/go-next/next/layer"
	"github.com/rs/zerolog"

	"github.com/alexpts/edu-go/internal/repo"
)

type Category struct {
	RestController
	Logger *zerolog.Logger
	Repo   *repo.Category
}

func (c *Category) ActionGet(ctx *layer.HandlerCtx) {
	catId := ctx.UriParams["id"]
	cat, _ := c.Repo.FindOneById(catId, AllRelation)
	c.sendJsonModel(ctx, cat)
}

func (c *Category) ActionCreate(ctx *layer.HandlerCtx) {
	cat := &model.Category{Title: "IT"}

	cat, _, err := c.Repo.Create(cat)
	if err != nil {
		c.sendError(ctx, err, 400)
		return
	}

	c.sendJsonModel(ctx, cat)
}
