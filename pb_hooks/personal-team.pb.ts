onRecordAfterCreateSuccess(({ next, record, app }) => {
	const id = record.id;
	let name = record.getString('name');

	if (!name || name === '') name = 'Anonymous';

	try {
		const teamsCollection = app.findCollectionByNameOrId('teams');

		const personalTeam = new Record(teamsCollection);
		personalTeam.set('name', `${name}'s team`);
		personalTeam.set('user', id);

		app.save(personalTeam);

		next();
	} catch (e) {
		console.error('Teams collection not found', e);
		return next();
	}
}, 'users');
