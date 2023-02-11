package controller

import (
	"github.com/alexpts/edu-go/internal/repo"
	"github.com/alexpts/edu-go/pkg/convert"
	"github.com/alexpts/go-next/next/layer"
	"github.com/rs/zerolog"
)

type User struct {
	Logger   *zerolog.Logger
	UserRepo *repo.User
}

func (c *User) ActionGet(ctx *layer.HandlerCtx) {
	userId := ctx.UriParams["id"]
	model, _ := c.UserRepo.FindOneById(convert.MustInt(userId), "Posts")
	sendJsonModel(ctx, model)
}

func (c *User) ActionFind(ctx *layer.HandlerCtx) {
	model, _ := c.UserRepo.FindAll("Posts")
	sendJsonModel(ctx, model)
}
