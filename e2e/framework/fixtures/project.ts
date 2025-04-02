import PocketBase from 'pocketbase';

import { TestTeam } from './team';

export type TestProject = {
	id: string;
	name: string;
	status: string;
};

export const createProject = async (
	pocketBase: PocketBase,
	team: TestTeam
): Promise<TestProject> => {
	const name = `Test Project ${Date.now()}`;
	const status = 'active';

	const { id } = await pocketBase.collection('projects').create({
		name,
		status,
		team: team.id
	});

	return {
		id,
		name,
		status
	};
};

export const deleteProject = async (pocketBase: PocketBase, project: TestProject) => {
	await pocketBase.collection('projects').delete(project.id);
};
