/// <reference path="../pb_data/types.d.ts" />
migrate(
	(app) => {
		const collection = app.findCollectionByNameOrId('pbc_484305853');

		// update field
		collection.fields.addAt(
			2,
			new Field({
				hidden: false,
				id: 'select2063623452',
				maxSelect: 1,
				name: 'status',
				presentable: false,
				required: true,
				system: false,
				type: 'select',
				values: ['planned', 'active', 'completed']
			})
		);

		return app.save(collection);
	},
	(app) => {
		const collection = app.findCollectionByNameOrId('pbc_484305853');

		// update field
		collection.fields.addAt(
			2,
			new Field({
				hidden: false,
				id: 'select2063623452',
				maxSelect: 1,
				name: 'status',
				presentable: false,
				required: true,
				system: false,
				type: 'select',
				values: ['active', 'completed']
			})
		);

		return app.save(collection);
	}
);
