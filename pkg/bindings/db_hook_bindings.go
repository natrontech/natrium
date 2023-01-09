package bindings

import (
	"errors"

	"github.com/natrongmbh/natrium/pkg/git"
	"github.com/natrongmbh/natrium/pkg/models"
	"github.com/natrongmbh/natrium/pkg/server"
	"github.com/pocketbase/pocketbase/core"
)

var err error

func BindDBHooks() error {
	server.App.OnModelAfterUpdate().Add(func(e *core.ModelEvent) error {
		if e.Model.TableName() == models.Tables.Repositories {
			gitRepo := &git.Repository{
				Repository: &models.Repository{},
			}
			if e == nil || e.Model == nil || e.Model.GetId() == "" {
				return errors.New("invalid event or model")
			}

			gitRepo.Repository, err = models.FindRepositoryById(server.App.Dao(), e.Model.GetId())
			if err != nil {
				return err
			}

			go func() {
				err = gitRepo.UpdateStatus()
			}()
			return err
		}
		return nil
	})
	server.App.OnModelBeforeDelete().Add(func(e *core.ModelEvent) error {
		if e.Model.TableName() == models.Tables.Repositories {
			gitRepo := &git.Repository{
				Repository: &models.Repository{},
			}
			if e == nil || e.Model == nil || e.Model.GetId() == "" {
				return errors.New("invalid event or model")
			}

			gitRepo.Repository, err = models.FindRepositoryById(server.App.Dao(), e.Model.GetId())
			if err != nil {
				return err
			}

			go func() {
				err = gitRepo.DeleteRepository()
			}()
			return err
		}

		return nil
	})

	return nil
}
