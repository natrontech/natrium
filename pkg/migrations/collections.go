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
				"id": "u161bbj7zmbk372",
				"name": "applications",
				"type": "base",
				"system": false,
				"schema": [
					{
						"id": "3ycrwbmk",
						"name": "name",
						"type": "text",
						"system": false,
						"required": true,
						"unique": false,
						"options": {
							"min": 3,
							"max": null,
							"pattern": ""
						}
					},
					{
						"id": "yenaegnm",
						"name": "description",
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
						"id": "mvojeij1",
						"name": "helm_chart_repository_url",
						"type": "url",
						"system": false,
						"required": true,
						"unique": false,
						"options": {
							"exceptDomains": null,
							"onlyDomains": null
						}
					},
					{
						"id": "bnqnxyqs",
						"name": "helm_chart_name",
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
						"id": "hjp5eywb",
						"name": "helm_chart_version",
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
						"id": "ndvjzeq4",
						"name": "helm_chart_default_values",
						"type": "text",
						"system": false,
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\"",
				"updateRule": "@request.auth.id != \"\"",
				"deleteRule": "@request.auth.id != \"\"",
				"options": {}
			},
			{
				"id": "bneldi0njw3lkod",
				"name": "repositories",
				"type": "base",
				"system": false,
				"schema": [
					{
						"id": "yvnbs5ns",
						"name": "url",
						"type": "url",
						"system": false,
						"required": true,
						"unique": false,
						"options": {
							"exceptDomains": null,
							"onlyDomains": null
						}
					},
					{
						"id": "ebpf4rgl",
						"name": "branch",
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
						"id": "jal6o9dj",
						"name": "repository_auth",
						"type": "relation",
						"system": false,
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "vv9qfinbyos05xv",
							"cascadeDelete": false
						}
					},
					{
						"id": "oojnemvs",
						"name": "status",
						"type": "select",
						"system": false,
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"PENDING",
								"READY",
								"SYNCING",
								"ERROR"
							]
						}
					},
					{
						"id": "juuahhsx",
						"name": "commit",
						"type": "text",
						"system": false,
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\"",
				"updateRule": "@request.auth.id != \"\"",
				"deleteRule": "@request.auth.id != \"\"",
				"options": {}
			},
			{
				"id": "vv9qfinbyos05xv",
				"name": "repository_auth",
				"type": "base",
				"system": false,
				"schema": [
					{
						"id": "hw2wffgb",
						"name": "username",
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
						"id": "j4k9fiv8",
						"name": "password",
						"type": "text",
						"system": false,
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\"",
				"updateRule": "@request.auth.id != \"\"",
				"deleteRule": "@request.auth.id != \"\"",
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
