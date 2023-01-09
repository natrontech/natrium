package bindings

import (
	"errors"

	"github.com/natrongmbh/natrium/pkg/models"
	"github.com/natrongmbh/natrium/pkg/server"
	"github.com/pocketbase/pocketbase/core"
)

var err error

func BindDBHooks() error {
	server.App.OnModelAfterUpdate().Add(func(e *core.ModelEvent) error {
		if e.Model.TableName() == models.Tables.Clusters {
			// cluster := &models.Cluster{}

			if e == nil || e.Model == nil || e.Model.GetId() == "" {
				return errors.New("invalid event or model")
			}

			// record, err := server.App.Dao().FindRecordById(models.Tables.Clusters, e.Model.GetId())
			// if err != nil {
			// 	return err
			// }

			_, err := models.FindClusterById(server.App.Dao(), e.Model.GetId())
			if err != nil {
				return err
			}

			// parse the record into a models.Cluster struct
		}
		return nil
	})
	server.App.OnModelBeforeDelete().Add(func(e *core.ModelEvent) error {
		if e.Model.TableName() == models.Tables.Clusters {

			return err
		}

		return nil
	})

	return nil
}
