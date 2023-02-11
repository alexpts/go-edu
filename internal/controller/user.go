package controller

import (
	"github.com/alexpts/go-next/next/layer"
	"github.com/rs/zerolog"

	"github.com/alexpts/edu-go/internal/repo"
	"github.com/alexpts/edu-go/pkg/convert"
)

type User struct {
	Logger   *zerolog.Logger
	UserRepo *repo.User
}

func (c *User) ActionGet(ctx *layer.HandlerCtx) {
	userId := ctx.UriParams["id"]
	model, _ := c.UserRepo.FindOneById(convert.MustInt(userId), AllRelation)
	sendJsonModel(ctx, model)
}

func (c *User) ActionFind(ctx *layer.HandlerCtx) {
	model, _ := c.UserRepo.FindAll(AllRelation)
	sendJsonModel(ctx, model)
}
