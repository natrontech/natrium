package git

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/natrongmbh/natrium/pkg/env"
	"github.com/natrongmbh/natrium/pkg/models"
	"github.com/natrongmbh/natrium/pkg/server"
)

type Repository struct {
	*models.Repository
}

func (repo *Repository) Validate() error {
	return nil
}

func (repo *Repository) Sync() (string, error) {
	var r *git.Repository
	var repoAuth *models.RepositoryAuth
	var err error
	path := fmt.Sprintf("%s/%s", env.REPO_DIR, repo.GetName())

	if repo.RepositoryAuth != "" {

		repoAuth, err = models.FindRepositoryAuthById(server.App.Dao(), repo.RepositoryAuth)
		if err != nil {
			return "", err
		}

		r, err = git.PlainClone(path, false, &git.CloneOptions{
			URL: repo.Url,
			Auth: &http.BasicAuth{
				Username: repoAuth.Username,
				Password: repoAuth.Password,
			},
		})
	} else {
		r, err = git.PlainClone(path, false, &git.CloneOptions{
			URL: repo.Url,
		})
	}

	if err == git.ErrRepositoryAlreadyExists {
		// pull
		r, err = git.PlainOpen(env.REPO_DIR + "/" + repo.GetName())
		if err != nil {
			return "", err
		}

		w, err := r.Worktree()
		if err != nil {
			return "", err
		}

		if repo.RepositoryAuth != "" {
			err = w.Pull(&git.PullOptions{
				RemoteName: "origin",
				Auth: &http.BasicAuth{
					Username: repoAuth.Username,
					Password: repoAuth.Password,
				},
			})
		} else {
			err = w.Pull(&git.PullOptions{RemoteName: "origin"})
		}
		if err != git.NoErrAlreadyUpToDate {
			return "", err
		}
	}

	ref, err := r.Head()
	if err != nil {
		return "", err
	}

	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return "", err
	}

	return commit.Hash.String(), nil
}

func (repo *Repository) GetName() string {
	// if repoUrl contains .git, remove it
	if repo.Url[len(repo.Url)-4:] == ".git" {
		repo.Url = repo.Url[:len(repo.Url)-4]
	}

	// get last part of url
	repoName := repo.Url[strings.LastIndex(repo.Url, "/")+1:]

	return repoName
}
