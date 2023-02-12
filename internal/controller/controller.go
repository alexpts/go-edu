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

	respBytes, err := c.Json.Marshal(model)
	if err != nil {
		panic(err)
	}

	ctx.Response.SetStatusCode(200)
	ctx.Response.AppendBody(respBytes)
}
