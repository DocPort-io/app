/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_484305853")

  // update collection data
  unmarshal({
    "createRule": "@request.auth.id ?= team.user.id",
    "deleteRule": "@request.auth.id ?= team.user.id",
    "listRule": "@request.auth.id ?= team.user.id",
    "updateRule": "@request.auth.id ?= team.user.id",
    "viewRule": "@request.auth.id ?= team.user.id"
  }, collection)

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_484305853")

  // update collection data
  unmarshal({
    "createRule": null,
    "deleteRule": null,
    "listRule": null,
    "updateRule": null,
    "viewRule": null
  }, collection)

  return app.save(collection)
})
