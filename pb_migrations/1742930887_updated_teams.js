/// <reference path="../pb_data/types.d.ts" />
migrate(
	(app) => {
		const collection = app.findCollectionByNameOrId('pbc_1568971955');

		// add field
		collection.fields.addAt(
			1,
			new Field({
				autogeneratePattern: '',
				hidden: false,
				id: 'text1579384326',
				max: 0,
				min: 0,
				name: 'name',
				pattern: '',
				presentable: false,
				primaryKey: false,
				required: true,
				system: false,
				type: 'text'
			})
		);

		// add field
		collection.fields.addAt(
			2,
			new Field({
				cascadeDelete: false,
				collectionId: '_pb_users_auth_',
				hidden: false,
				id: 'relation2375276105',
				maxSelect: 10000,
				minSelect: 0,
				name: 'user',
				presentable: false,
				required: false,
				system: false,
				type: 'relation'
			})
		);

		return app.save(collection);
	},
	(app) => {
		const collection = app.findCollectionByNameOrId('pbc_1568971955');

		// remove field
		collection.fields.removeById('text1579384326');

		// remove field
		collection.fields.removeById('relation2375276105');

		return app.save(collection);
	}
);
