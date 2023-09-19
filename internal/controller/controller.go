package controller

import (
	"net/http"
	"reflect"

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

	if model == nil || reflect.TypeOf(model) == nil || reflect.ValueOf(model).IsNil() {
		ctx.Response.SetStatusCode(http.StatusNotFound)
		ctx.Response.AppendBodyString(`{"error": "not found""}`)
		return
	}

	respBytes := c.mustJsonMap(map[string]any{
		"data": model,
	})

	ctx.Response.SetStatusCode(http.StatusOK)
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
