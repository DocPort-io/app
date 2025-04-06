const tryCatch = <T>(
	fn: () => T
): { success: true; data: T } | { success: false; error: Error } => {
	try {
		return { success: true, data: fn() };
	} catch (error) {
		return { success: false, error: error as Error };
	}
};

onRecordAfterCreateSuccess(({ next, record, app }) => {
	const id = record.id;
	let name = record.getString('name');

	if (!name || name === '') name = 'Anonymous';

	const teamsCollectionResult = tryCatch(() => app.findCollectionByNameOrId('teams'));
	if (!teamsCollectionResult.success) return next();
	const { data: teamsCollection } = teamsCollectionResult;

	const personalTeam = new Record(teamsCollection);
	personalTeam.set('name', `${name}'s team`);
	personalTeam.set('user', id);

	app.save(personalTeam);

	next();
}, 'users');
