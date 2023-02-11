package controller

import (
	"encoding/json"
	"github.com/alexpts/go-next/next/layer"
)

func sendJsonModel(ctx *layer.HandlerCtx, model any) {
	ctx.Response.Header.Add("content-type", "application/json")

	if model == nil {
		ctx.Response.SetStatusCode(404)
		ctx.Response.AppendBodyString(`{"error": "not found""}`)
		return
	}

	respBytes, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}

	ctx.Response.SetStatusCode(200)
	ctx.Response.AppendBody(respBytes)
}
