package post

import (
	"github.com/alexpts/edu-go/internal/model"
	"gorm.io/gorm"
)

type Repo struct {
	Db *gorm.DB
}

func (repo *Repo) FindById(id int) *model.Post {
	post := new(model.Post)
	tx := repo.Db.InnerJoins("Cat").Select("*").Take(post, id)

	if tx.RowsAffected == 0 {
		return nil
	}

	return post
}

func (repo *Repo) FindShortPostById(id int) *model.ShortPost {
	post := new(model.ShortPost)
	tx := repo.Db.InnerJoins("Cat").Take(post, id)

	if tx.RowsAffected == 0 {
		return nil
	}

	return post
}
