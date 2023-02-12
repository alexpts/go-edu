package repo

import (
	"errors"

	"gorm.io/gorm"
)

// Repo - @todo add constraint T ~struct
type Repo[T any] struct {
	Db *gorm.DB
}

func (repo *Repo[T]) FindAll(relations ...string) ([]T, error) {
	var models []T
	tx := repo.withRelations(relations).Find(&models)
	return repo.resultMany(models, tx.Error)
}

func (repo *Repo[T]) FindOneById(id int, relations ...string) (*T, error) {
	model := new(T)
	tx := repo.withRelations(relations).Take(model, id)
	return repo.resultOne(tx, model)
}

func (repo *Repo[T]) withRelations(relations []string) *gorm.DB {
	tx := repo.Db

	for _, relName := range relations {
		tx = repo.Db.Preload(relName)
	}

	return tx
}

func (repo *Repo[T]) resultMany(models []T, err error) ([]T, error) {
	if err != nil {
		return nil, err
	}

	return models, nil
}

func (repo *Repo[T]) resultOne(tx *gorm.DB, model *T) (*T, error) {
	if tx.Error != nil {
		if errors.As(tx.Error, &gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, nil
	}

	return model, nil
}
