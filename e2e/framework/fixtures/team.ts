import PocketBase from 'pocketbase';

import { TestUser } from './user';

export type TestTeam = {
	id: string;
	name: string;
};

export const createTeam = async (pocketBase: PocketBase, user: TestUser): Promise<TestTeam> => {
	const name = `Test Team ${Date.now()}`;

	const { id } = await pocketBase.collection('teams').create({
		name,
		user: [user.id]
	});

	return {
		id,
		name
	};
};

export const deleteTeam = async (pocketBase: PocketBase, team: TestTeam) => {
	await pocketBase.collection('teams').delete(team.id);
};
