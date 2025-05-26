/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_3446931122")

  // add field
  collection.fields.addAt(5, new Field({
    "cascadeDelete": false,
    "collectionId": "pbc_1502746827",
    "hidden": false,
    "id": "relation433864914",
    "maxSelect": 10000,
    "minSelect": 0,
    "name": "versions",
    "presentable": false,
    "required": true,
    "system": false,
    "type": "relation"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_3446931122")

  // remove field
  collection.fields.removeById("relation433864914")

  return app.save(collection)
})
