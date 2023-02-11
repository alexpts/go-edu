package repo

import (
	"github.com/alexpts/edu-go/internal/model"
	"gorm.io/gorm"
)

type User struct {
	Db *gorm.DB
}

func (repo *User) FindOneById(id int, relations ...string) (*model.User, error) {
	m := &model.User{}
	tx := repo.withRelations(relations).Take(m, id)

	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, nil
	}

	return m, nil
}

func (repo *User) FindAll(relations ...string) ([]model.User, error) {
	var models []model.User
	tx := repo.withRelations(relations).Find(&models)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return models, nil
}

func (repo *User) withRelations(relations []string) *gorm.DB {
	tx := repo.Db

	for _, relName := range relations {
		tx = repo.Db.Preload(relName)
	}

	return tx
}

// @todo generic for any type and move to base repository
func (repo *User) prepareOneResult(tx *gorm.DB, model *any) (*any, error) {
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, nil
	}

	return model, nil
}