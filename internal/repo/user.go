package repo

import "github.com/alexpts/edu-go/internal/model"

// type User = Repo[model.Profile] - cant` extend
// type User Repo[model.Profile] - cant` extend

// User - extend general repo via embed Repo and custom methods
type User struct {
	Repo[model.User]
}

func (repo *User) FindOneUserByName(name string) (*model.User, error) {
	m := &model.User{}

	tx := repo.Db.
		Where(model.User{Name: name}).
		Take(m)

	return repo.resultOne(tx, m)
}
