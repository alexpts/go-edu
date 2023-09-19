package controller

import (
	"net/http"

	"github.com/alexpts/go-next/next/layer"
	"github.com/rs/zerolog"

	m "github.com/alexpts/edu-go/internal/model"
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
	model, _ := c.UserRepo.FindOneById(ctx, userId, AllRelation)
	c.sendJsonModel(ctx, model)
}

func (c *User) ActionFind(ctx *layer.HandlerCtx) {
	models, _ := c.UserRepo.FindAll(ctx, AllRelation)
	c.sendJsonModel(ctx, models)
}

// ActionCreate
// @todo check `content-type` middleware (extra rule for resolver)
// @todo json-schema validate
// @todo mapper payload to model
// @todo expose error and mapping errors levels
// @todo error wrap/unwrap for logging
// @todo return error for any handler (to Next-app)
func (c *User) ActionCreate(ctx *layer.HandlerCtx) {
	user := &m.User{}
	err := c.Json.Unmarshal(ctx.Request.Body(), user)
	if err != nil {
		c.sendError(ctx, err, http.StatusBadRequest)
		return
	}

	// model, _, err := c.UserRepo.Create(user)
	model, _, err := c.UserRepo.Persist(ctx, user)
	if err != nil {
		c.sendError(ctx, err, http.StatusBadRequest)
		return
	}

	c.sendJsonModel(ctx, model)
}

func (c *User) ActionUpdate(ctx *layer.HandlerCtx) {
	userId := convert.MustInt(ctx.UriParams["id"])
	model, _ := c.UserRepo.FindOneById(ctx, userId)

	err := c.Json.Unmarshal(ctx.Request.Body(), model)
	if err != nil {
		c.sendError(ctx, err, http.StatusBadRequest)
		return
	}

	// model, _, err = c.UserRepo.Persist(model) // Create new if not found by model (id + ver)
	model, _, err = c.UserRepo.Update(ctx, model)
	if err != nil {
		c.sendError(ctx, err, http.StatusBadRequest)
		return
	}

	c.sendJsonModel(ctx, model)
}

func (c *User) ActionGetByName(ctx *layer.HandlerCtx) {
	// model, _ := c.UserRepo.FindOneUserByName(ctx.UriParams["name"])
	model, _ := c.UserRepo.FindByNameRawSQL(ctx.UriParams["name"])
	c.sendJsonModel(ctx, model)
}
