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
	tx := repo.Db.Take(post, id)

	if tx.RowsAffected == 0 {
		return nil
	}

	return post
}
