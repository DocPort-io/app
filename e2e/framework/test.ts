import { test as baseTest } from '@playwright/test';
import PocketBase from 'pocketbase';

import { LoginPage } from '../pages/auth/login.page';
import { DashboardPage } from '../pages/dashboard.page';
import { ProjectsPage } from '../pages/projects.page';
import { createPocketBase } from './fixtures/pocketbase';
import { TestTeam } from './fixtures/team';
import { createTestData, TestData } from './fixtures/test-data';
import { TestUser } from './fixtures/user';

type TestFixtures = {
	// Pages
	loginPage: LoginPage;
	dashboardPage: DashboardPage;
	projectsPage: ProjectsPage;

	// Fixtures
	pocketBase: PocketBase;
	testUser: TestUser;
	testTeam: TestTeam;
	testData: TestData;
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
	testUser: async ({ testData }, use) => {
		const user = await testData.user();
		await use(user);
	},
	testTeam: async ({ testUser, testData }, use) => {
		const team = await testData.team(testUser);
		await use(team);
	},
	testData: async ({ pocketBase }, use) => {
		const testData = createTestData(pocketBase);
		await use(testData);

		const { status, expectedStatus } = test.info();
		if (status !== expectedStatus) return;

		await testData.cleanup();
	}
});
