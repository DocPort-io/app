import PocketBase from 'pocketbase';

import { createProject, deleteProject, TestProject } from './project';
import { createTeam, deleteTeam, TestTeam } from './team';
import { createUser, deleteUser, TestUser } from './user';

export type TestDataCleanupFn = () => Promise<void>;
export type TestUserFn = () => Promise<TestUser>;
export type TestTeamFn = (user: TestUser) => Promise<TestTeam>;
export type TestProjectFn = (team: TestTeam) => Promise<TestProject>;
export type TestProjectDirtyFn = (team: TestTeam) => void;

export type TestData = {
	cleanup: TestDataCleanupFn;
	user: TestUserFn;
	team: TestTeamFn;
	project: TestProjectFn;
	markProjectAsDirty: TestProjectDirtyFn;
};

export const createTestData = (pocketBase: PocketBase): TestData => {
	const resources: TestDataCleanupFn[] = [];

	const cleanup = async () => {
		for await (const resource of [...resources].reverse()) {
			await resource();
		}
	};

	const user = async () => {
		const user = await createUser(pocketBase);
		resources.push(() => deleteUser(pocketBase, user));
		return user;
	};

	const team = async (user: TestUser) => {
		const team = await createTeam(pocketBase, user);
		resources.push(() => deleteTeam(pocketBase, team));
		return team;
	};

	const project = async (team: TestTeam) => {
		const project = await createProject(pocketBase, team);
		resources.push(() => deleteProject(pocketBase, project));
		return project;
	};

	const markProjectAsDirty = (team: TestTeam) => {
		resources.push(async () => {
			const projects = await pocketBase.collection('projects').getFullList({
				filter: `team = '${team.id}'`,
				sort: '-created'
			});
			await Promise.all(
				projects.map(async (project) => {
					await pocketBase.collection('projects').delete(project.id);
				})
			);
		});
	};

	return {
		cleanup,
		user,
		team,
		project,
		markProjectAsDirty
	};
};
