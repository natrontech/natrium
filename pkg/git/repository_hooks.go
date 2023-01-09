package git

import (
	"errors"
	"os"

	"github.com/natrongmbh/natrium/pkg/env"
	"github.com/natrongmbh/natrium/pkg/models"
	"github.com/natrongmbh/natrium/pkg/server"
	"github.com/natrongmbh/natrium/pkg/util"
)

func (repo *Repository) UpdateStatus() error {
	var err error

	if repo.Status != models.RepositoryStatus.SYNCING {
		return nil
	}

	commit, err := repo.Sync()
	if err != nil {
		util.Logger.Error(err)
		repo.Status = models.RepositoryStatus.ERROR
	} else {
		repo.Status = models.RepositoryStatus.READY
	}

	record, err := server.App.Dao().FindRecordsByIds(models.Tables.Repositories, []string{repo.Id})
	if err != nil {
		return err
	}

	if len(record) == 0 {
		return errors.New("repository record not found")
	}

	record[0].Set("status", repo.Status)

	if commit != "" {
		record[0].Set("commit", commit)
	}

	err = server.App.Dao().SaveRecord(record[0])
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) DeleteRepository() error {
	// delete directory
	err := os.RemoveAll(env.REPO_DIR + "/" + repo.GetName())
	if err != nil {
		util.Logger.Error(err)
	}

	return nil
}
