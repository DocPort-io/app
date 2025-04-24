/// <reference path="../pb_data/types.d.ts" />
migrate(
	(app) => {
		const collection = app.findCollectionByNameOrId('pbc_1568971955');

		// update collection data
		unmarshal(
			{
				createRule: '@request.auth.id ?= user.id',
				deleteRule: '@request.auth.id ?= user.id',
				updateRule: '@request.auth.id ?= user.id',
				viewRule: '@request.auth.id ?= user.id'
			},
			collection
		);

		return app.save(collection);
	},
	(app) => {
		const collection = app.findCollectionByNameOrId('pbc_1568971955');

		// update collection data
		unmarshal(
			{
				createRule: null,
				deleteRule: null,
				updateRule: null,
				viewRule: null
			},
			collection
		);

		return app.save(collection);
	}
);
