package home

import (
	"encoding/json"
	"github.com/alexpts/go-next/next"
	"github.com/rs/zerolog"
)

type ControllerHome struct {
	Logger *zerolog.Logger
}

type User struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

// @todo need Response less syntax

func (c *ControllerHome) ActionIndex(ctx *next.HandlerCxt) {

	user := User{Name: "alex", Lastname: "some"}

	ctx.Response.Header.Add("content-type", "application/json")
	respBytes, err := json.Marshal(user)

	if err != nil {
		ctx.Panic(err, map[string]interface{}{})
		return
	}

	ctx.Response.SetStatusCode(200)
	ctx.Response.AppendBody(respBytes)
}
