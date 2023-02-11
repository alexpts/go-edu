package repo

import (
	"gorm.io/gorm"

	"github.com/alexpts/edu-go/internal/model"
)

type Post struct {
	Db *gorm.DB
}

func (repo *Post) FindOneById(id int, relations ...string) (*model.PostRel, error) {
	post := &model.PostRel{}
	tx := repo.withRelations(relations).Take(post, id)

	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, nil
	}

	return post, nil
}

func (repo *Post) FindAll(relations ...string) ([]model.Post, error) {
	var models []model.Post
	tx := repo.withRelations(relations).Find(&models)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return models, nil
}

func (repo *Post) withRelations(relations []string) *gorm.DB {
	tx := repo.Db

	for _, relName := range relations {
		tx = repo.Db.Preload(relName)
	}

	return tx
}

// @todo generic for any type and move to base repository
func (repo *Post) prepareOneResult(tx *gorm.DB, model *any) (*any, error) {
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, nil
	}

	return model, nil
}
