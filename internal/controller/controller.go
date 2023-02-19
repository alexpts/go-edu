package controller

import (
	"github.com/alexpts/go-next/next/layer"
	"gorm.io/gorm/clause"

	"github.com/alexpts/edu-go/pkg/convert"
)

type RestController struct {
	Json convert.IJsonMarshaler
}

const AllRelation = clause.Associations

func (c *RestController) sendJsonModel(ctx *layer.HandlerCtx, model any) {
	ctx.Response.Header.Add("content-type", "application/json")

	if model == nil {
		ctx.Response.SetStatusCode(404)
		ctx.Response.AppendBodyString(`{"error": "not found""}`)
		return
	}

	respBytes := c.mustJsonMap(map[string]any{
		"data": model,
	})

	ctx.Response.SetStatusCode(200)
	ctx.Response.AppendBody(respBytes)
}

func (c *RestController) sendError(ctx *layer.HandlerCtx, err error, statusCode int) {
	bytes := c.mustJsonMap(map[string]any{
		"error": err.Error(),
	})

	ctx.Response.SetStatusCode(statusCode)
	ctx.Response.AppendBody(bytes)
}

func (c *RestController) mustJsonMap(data map[string]any) []byte {
	bytes, err := c.Json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return bytes
}
