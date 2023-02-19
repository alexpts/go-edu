package repo

import (
	"errors"
	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

// Repo - @todo add constraint T ~struct
type Repo[T any] struct {
	Db *gorm.DB
}

func (repo *Repo[T]) Create(model *T) (*T, int64, error) {
	result := repo.Db.Omit(clause.Associations).Create(model)
	return model, result.RowsAffected, result.Error
}

func (repo *Repo[T]) Update(model *T) (*T, int64, error) {
	result := repo.Db.Omit(clause.Associations).Select("*").Updates(*model) // Select(*) update all fields, without zero-value skipped
	return model, result.RowsAffected, result.Error
}

func (repo *Repo[T]) Persist(model *T) (*T, int64, error) {
	tx := repo.Db.Debug().Omit(clause.Associations).Save(model)
	return model, tx.RowsAffected, tx.Error
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
