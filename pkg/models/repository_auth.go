package models

import (
	"strings"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*RepositoryAuth)(nil)

// TODO: add hashing for password
type RepositoryAuth struct {
	models.BaseModel

	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *RepositoryAuth) TableName() string {
	return "repository_auth"
}

func RepositoryAuthQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&RepositoryAuth{})
}

func FindRepositoryAuthById(dao *daos.Dao, id string) (*RepositoryAuth, error) {
	repositoryAuth := &RepositoryAuth{}

	err := RepositoryAuthQuery(dao).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(repositoryAuth)

	if err != nil {
		return nil, err
	}

	return repositoryAuth, nil
}

func FindRepositoryAuthByUsername(dao *daos.Dao, username string) (*RepositoryAuth, error) {
	repositoryAuth := &RepositoryAuth{}

	err := RepositoryAuthQuery(dao).
		// case insensitive match
		AndWhere(dbx.NewExp("LOWER(username)={:username}", dbx.Params{
			"username": strings.ToLower(username),
		})).
		Limit(1).
		One(repositoryAuth)

	if err != nil {
		return nil, err
	}

	return repositoryAuth, nil
}

func FindLast10RepositoryAuths(dao *daos.Dao, name string) ([]*RepositoryAuth, error) {
	repositoryAuths := []*RepositoryAuth{}

	err := RepositoryAuthQuery(dao).
		Limit(10).
		All(&repositoryAuths)

	if err != nil {
		return nil, err
	}

	return repositoryAuths, nil
}
