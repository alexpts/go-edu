package repo

import (
	"github.com/alexpts/edu-go/internal/model"
)

//go:generate mockery --name=IPostRepo --structname=IPostRepoMock
type IPostRepo interface {
	IRepo[model.Post]
}

type Post struct {
	Repo[model.Post]
}
