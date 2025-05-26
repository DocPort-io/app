/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_1502746827")

  // remove field
  collection.fields.removeById("file104153177")

  // add field
  collection.fields.addAt(4, new Field({
    "cascadeDelete": false,
    "collectionId": "pbc_3446931122",
    "hidden": false,
    "id": "relation104153177",
    "maxSelect": 1000,
    "minSelect": 0,
    "name": "files",
    "presentable": false,
    "required": false,
    "system": false,
    "type": "relation"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_1502746827")

  // add field
  collection.fields.addAt(4, new Field({
    "hidden": false,
    "id": "file104153177",
    "maxSelect": 500,
    "maxSize": 512000000,
    "mimeTypes": [],
    "name": "files",
    "presentable": false,
    "protected": true,
    "required": false,
    "system": false,
    "thumbs": [],
    "type": "file"
  }))

  // remove field
  collection.fields.removeById("relation104153177")

  return app.save(collection)
})
