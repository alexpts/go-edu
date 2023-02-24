package repo

import "github.com/alexpts/edu-go/internal/model"

type PostRepo = Repo[model.Post]
type Category struct {
	Repo[model.Category]
	//PostRepo
}

func (repo *Category) FindOneById(uuid string, relations ...string) (*model.Category, error) {
	m := new(model.Category)
	tx := repo.withRelations(relations).Debug().Take(m, "id = ?", uuid)

	return repo.resultOne(tx, m)
	//return Repo[model.Category].resultOne(repo.Repo, tx, m)
}
