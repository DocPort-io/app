/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_3446931122")

  // update collection data
  unmarshal({
    "createRule": "@request.auth.id ?= versions.project.team.user.id",
    "deleteRule": "@request.auth.id ?= versions.project.team.user.id",
    "listRule": "@request.auth.id ?= versions.project.team.user.id",
    "updateRule": "@request.auth.id ?= versions.project.team.user.id",
    "viewRule": "@request.auth.id ?= versions.project.team.user.id"
  }, collection)

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_3446931122")

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
