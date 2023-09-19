package repo

import (
	sq "github.com/Masterminds/squirrel"

	"github.com/alexpts/edu-go/internal/model"
)

//go:generate mockery --name=IUserRepo --structname=IUserRepoMock
type IUserRepo interface {
	IRepo[model.User]

	FindOneUserByName(name string) (*model.User, error)
	FindByNameRawSQL(name string) (*model.User, error)
}

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

// FindByNameRawSQL - create sql via squirrel and custom scan to multi model
func (repo *User) FindByNameRawSQL(name string) (*model.User, error) {
	m := &model.User{}

	query, args, err := sq.Select("u.*, p.id, p.title, p.status, p.user_id").
		From("users as u").
		InnerJoin("posts as p ON u.id = p.user_id").
		Where(sq.Eq{
			"name": name,
		}).
		ToSql()

	if err != nil {
		return nil, err
	}

	// ### Single model scan via GORM.Scan
	//tx := repo.Db.Raw(query, args).Scan(m)
	//return repo.resultOne(tx, m)

	// Native Rows via GORM
	rows, err := repo.Db.Model(m).Raw(query, args).Rows() // (*sql.Rows, error)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := model.Post{}

		// Multi model scan
		err := rows.Scan(&m.Id, &m.Name, &post.Id, &post.Title, &post.Status, &post.UserId)
		if err != nil {
			return nil, err
		}

		m.Posts = append(m.Posts, post)
	}

	return m, err
}
