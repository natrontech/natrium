package models

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var ClusterStatus = newClusterStatus()

func newClusterStatus() *clusterStatusRegistry {
	return &clusterStatusRegistry{
		PENDING: "PENDING",
		RUNNING: "RUNNING",
		SYNCING: "SYNCING",
		ERROR:   "ERROR",
	}
}

type clusterStatusRegistry struct {
	PENDING string
	RUNNING string
	SYNCING string
	ERROR   string
}

var _ models.Model = (*Cluster)(nil)

type Cluster struct {
	models.BaseModel

	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

func (r *Cluster) TableName() string {
	return "clusters"
}

func ClusterQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Cluster{})
}

func FindClusterById(dao *daos.Dao, id string) (*models.Record, error) {

	collection, err := dao.FindCollectionByNameOrId(Tables.Clusters)
	if err != nil {
		return nil, err
	}

	query := dao.RecordQuery(collection).
		AndWhere(dbx.HashExp{"id": id}).
		OrderBy("created DESC").
		Limit(1)

	record := dbx.NullStringMap{}
	if err := query.One(&record); err != nil {
		return nil, err
	}

	return models.NewRecordFromNullStringMap(collection, record), nil
}

func FindClusterByName(dao *daos.Dao, name string) (*models.Record, error) {
	collection, err := dao.FindCollectionByNameOrId(Tables.Clusters)
	if err != nil {
		return nil, err
	}

	query := dao.RecordQuery(collection).
		AndWhere(dbx.HashExp{"name": name}).
		OrderBy("created DESC").
		Limit(1)

	record := dbx.NullStringMap{}
	if err := query.One(&record); err != nil {
		return nil, err
	}

	return models.NewRecordFromNullStringMap(collection, record), nil
}

func FindLast10Clusters(dao *daos.Dao) ([]*models.Record, error) {
	collection, err := dao.FindCollectionByNameOrId(Tables.Clusters)
	if err != nil {
		return nil, err
	}

	query := dao.RecordQuery(collection).
		OrderBy("created DESC").
		Limit(10)

	rows := []dbx.NullStringMap{}
	if err := query.All(&rows); err != nil {
		return nil, err
	}

	return models.NewRecordsFromNullStringMaps(collection, rows), nil
}
