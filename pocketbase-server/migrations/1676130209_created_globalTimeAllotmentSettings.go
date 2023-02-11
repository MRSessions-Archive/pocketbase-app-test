package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "b3kjwovnfq4zrrj",
			"created": "2023-02-11 15:43:29.481Z",
			"updated": "2023-02-11 15:43:29.481Z",
			"name": "globalTimeAllotmentSettings",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "um9a5dx2",
					"name": "dayOfWeek",
					"type": "text",
					"required": true,
					"unique": true,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "k2nsdtyb",
					"name": "hour",
					"type": "number",
					"required": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": 24
					}
				},
				{
					"system": false,
					"id": "hjyzgbje",
					"name": "minute",
					"type": "number",
					"required": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": 59
					}
				}
			],
			"listRule": "",
			"viewRule": "",
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("b3kjwovnfq4zrrj")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
