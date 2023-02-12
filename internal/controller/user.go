package controller

import (
	"github.com/alexpts/go-next/next/layer"
	"github.com/rs/zerolog"

	"github.com/alexpts/edu-go/internal/repo"
	"github.com/alexpts/edu-go/pkg/convert"
)

type User struct {
	RestController
	Logger   *zerolog.Logger
	UserRepo *repo.User
}

func (c *User) ActionGet(ctx *layer.HandlerCtx) {
	userId := convert.MustInt(ctx.UriParams["id"])
	model, _ := c.UserRepo.FindOneById(userId, AllRelation)
	c.sendJsonModel(ctx, model)
}

func (c *User) ActionFind(ctx *layer.HandlerCtx) {
	model, _ := c.UserRepo.FindAll(AllRelation)
	c.sendJsonModel(ctx, model)
}

func (c *User) ActionGetByName(ctx *layer.HandlerCtx) {
	model, _ := c.UserRepo.FindOneUserByName(ctx.UriParams["name"])
	c.sendJsonModel(ctx, model)
}
