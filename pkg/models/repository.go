package models

import (
	"strings"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var RepositoryStatus = newRepositoryStatusRegistry()

func newRepositoryStatusRegistry() *repositoryStatusRegistry {
	return &repositoryStatusRegistry{
		PENDING: "PENDING",
		READY:   "READY",
		SYNCING: "SYNCING",
		ERROR:   "ERROR",
	}
}

type repositoryStatusRegistry struct {
	PENDING string
	READY   string
	SYNCING string
	ERROR   string
}

var _ models.Model = (*Repository)(nil)

type Repository struct {
	models.BaseModel

	Url            string `json:"url"`
	Branch         string `json:"branch"`
	RepositoryAuth string `json:"repository_auth"`
	Status         string `json:"status"`
	Commit         string `json:"head"`
}

func (r *Repository) TableName() string {
	return "repositories"
}

func RepositoryQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Repository{})
}

func FindRepositoryById(dao *daos.Dao, id string) (*Repository, error) {
	repository := &Repository{}

	err := RepositoryQuery(dao).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(repository)

	if err != nil {
		return nil, err
	}

	return repository, nil
}

func FindRepositoryByName(dao *daos.Dao, name string) (*Repository, error) {
	repository := &Repository{}

	err := RepositoryQuery(dao).
		// case insensitive match
		AndWhere(dbx.NewExp("LOWER(name)={:name}", dbx.Params{
			"name": strings.ToLower(name),
		})).
		Limit(1).
		One(repository)

	if err != nil {
		return nil, err
	}

	return repository, nil
}

func FindLast10Repositorys(dao *daos.Dao, name string) ([]*Repository, error) {
	repositorys := []*Repository{}

	err := RepositoryQuery(dao).
		Limit(10).
		All(&repositorys)

	if err != nil {
		return nil, err
	}

	return repositorys, nil
}
