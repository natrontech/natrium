package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

// Auto generated migration with the most recent collections configuration.
func InitCollections() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
					"id": "_pb_users_auth_",
					"name": "users",
					"type": "auth",
					"system": false,
					"schema": [
							{
									"id": "users_name",
									"name": "name",
									"type": "text",
									"system": false,
									"required": false,
									"unique": false,
									"options": {
											"min": null,
											"max": null,
											"pattern": ""
									}
							},
							{
									"id": "users_avatar",
									"name": "avatar",
									"type": "file",
									"system": false,
									"required": false,
									"unique": false,
									"options": {
											"maxSelect": 1,
											"maxSize": 5242880,
											"mimeTypes": [
													"image/jpg",
													"image/jpeg",
													"image/png",
													"image/svg+xml",
													"image/gif"
											],
											"thumbs": null
									}
							}
					],
					"listRule": "id = @request.auth.id",
					"viewRule": "id = @request.auth.id",
					"createRule": "",
					"updateRule": "id = @request.auth.id",
					"deleteRule": "id = @request.auth.id",
					"options": {
							"allowEmailAuth": true,
							"allowOAuth2Auth": true,
							"allowUsernameAuth": true,
							"exceptEmailDomains": null,
							"manageRule": null,
							"minPasswordLength": 8,
							"onlyEmailDomains": null,
							"requireEmail": false
					}
			},
			{
					"id": "7xstv9iyb4uxnaj",
					"name": "clusters",
					"type": "base",
					"system": false,
					"schema": [
							{
									"id": "6mtd0gvl",
									"name": "user",
									"type": "relation",
									"system": false,
									"required": false,
									"unique": false,
									"options": {
											"maxSelect": 1,
											"collectionId": "_pb_users_auth_",
											"cascadeDelete": false
									}
							},
							{
									"id": "khuj1f6l",
									"name": "name",
									"type": "text",
									"system": false,
									"required": true,
									"unique": false,
									"options": {
											"min": null,
											"max": null,
											"pattern": ""
									}
							},
							{
									"id": "pyndqytd",
									"name": "type",
									"type": "select",
									"system": false,
									"required": true,
									"unique": false,
									"options": {
											"maxSelect": 1,
											"values": [
													"NORMAL",
													"DEDICATED",
													"GPU"
											]
									}
							},
							{
									"id": "nx4eva63",
									"name": "cpu",
									"type": "number",
									"system": false,
									"required": true,
									"unique": false,
									"options": {
											"min": null,
											"max": null
									}
							},
							{
									"id": "zzqpe1jh",
									"name": "memory",
									"type": "number",
									"system": false,
									"required": true,
									"unique": false,
									"options": {
											"min": null,
											"max": null
									}
							},
							{
									"id": "rfprhk9g",
									"name": "storage",
									"type": "text",
									"system": false,
									"required": true,
									"unique": false,
									"options": {
											"min": null,
											"max": null,
											"pattern": ""
									}
							},
							{
									"id": "aipvrkv1",
									"name": "status",
									"type": "select",
									"system": false,
									"required": true,
									"unique": false,
									"options": {
											"maxSelect": 1,
											"values": [
													"PENDING",
													"RUNNING",
													"SYNCING",
													"ERROR"
											]
									}
							}
					],
					"listRule": "@request.auth.id != \"\" ",
					"viewRule": "@request.auth.id != \"\"",
					"createRule": "@request.auth.id != \"\"",
					"updateRule": "@request.auth.id != \"\"",
					"deleteRule": "@request.auth.id != \"\" ",
					"options": {}
			}
	]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		// no revert since the configuration on the environment, on which
		// the migration was executed, could have changed via the UI/API
		return nil
	})
}
