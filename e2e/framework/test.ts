import { test as baseTest } from '@playwright/test';
import PocketBase from 'pocketbase';

import { LoginPage } from '../pages/auth/login.page';
import { DashboardPage } from '../pages/dashboard.page';
import { ProjectsPage } from '../pages/projects.page';
import { createPocketBase } from './fixtures/pocketbase';
import { createTeam, deleteTeam, TestTeam } from './fixtures/team';
import { createUser, deleteUser, TestUser } from './fixtures/user';

type TestFixtures = {
	// Pages
	loginPage: LoginPage;
	dashboardPage: DashboardPage;
	projectsPage: ProjectsPage;

	// Fixtures
	pocketBase: PocketBase;
	testUser: TestUser;
	testTeam: TestTeam;
};

export const test = baseTest.extend<TestFixtures>({
	// Pages
	loginPage: async ({ page }, use) => {
		await use(new LoginPage(page));
	},
	dashboardPage: async ({ page }, use) => {
		await use(new DashboardPage(page));
	},
	projectsPage: async ({ page }, use) => {
		await use(new ProjectsPage(page));
	},

	// Fixtures
	pocketBase: async ({}, use) => {
		const pocketBase = await createPocketBase();
		await use(pocketBase);
	},
	testUser: async ({ pocketBase }, use) => {
		const user = await createUser(pocketBase);
		await use(user);
		await deleteUser(pocketBase, user);
	},
	testTeam: async ({ pocketBase, testUser }, use) => {
		const team = await createTeam(pocketBase, testUser);
		await use(team);
		await deleteTeam(pocketBase, team);
	}
});
