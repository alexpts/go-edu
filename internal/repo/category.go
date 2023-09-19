package repo

import (
	"context"

	"github.com/alexpts/edu-go/internal/model"
)

type PostRepo = Repo[model.Post]

//go:generate mockery --name=ICategoryRepo --structname=ICategoryRepoMock
type ICategoryRepo interface {
	// IRepo[model.Category] // перекрыть метод интерфейса хочется, но интерфейсы не дают и приходится дублировать
	Create(ctx context.Context, model *model.Category) (*model.Category, int64, error)
	Update(ctx context.Context, model *model.Category) (*model.Category, int64, error)
	Persist(ctx context.Context, model *model.Category) (*model.Category, int64, error)
	FindAll(ctx context.Context, relations ...string) ([]model.Category, error)
	// FindOneById(ctx context.Context, id int, relations ...string) (*model.Category, error)

	// FindOneById overlapping method
	FindOneById(ctx context.Context, uuid string, relations ...string) (*model.Category, error)
}

type Category struct {
	Repo[model.Category]
	//PostRepo
}

func (repo *Category) FindOneById(ctx context.Context, uuid string, relations ...string) (*model.Category, error) {
	m := new(model.Category)
	tx := repo.withRelations(relations).Debug().Take(m, "id = ?", uuid)

	return repo.resultOne(tx, m)
	//return Repo[model.Category].resultOne(repo.Repo, tx, m)
}
