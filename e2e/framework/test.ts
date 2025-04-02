import { test as baseTest } from '@playwright/test';
import PocketBase from 'pocketbase';

import { LoginPage } from '../pages/auth/login.page';
import { DashboardPage } from '../pages/dashboard.page';
import { createPocketBase } from './fixtures/pocketbase';
import { createUser, deleteUser, TestUser } from './fixtures/user';

type TestFixtures = {
	// Pages
	loginPage: LoginPage;
	dashboardPage: DashboardPage;

	// Fixtures
	pocketBase: PocketBase;
	testUser: TestUser;
};

export const test = baseTest.extend<TestFixtures>({
	loginPage: async ({ page }, use) => {
		await use(new LoginPage(page));
	},
	dashboardPage: async ({ page }, use) => {
		await use(new DashboardPage(page));
	},
	pocketBase: async ({}, use) => {
		const pocketBase = await createPocketBase();
		await use(pocketBase);
	},
	testUser: async ({ pocketBase }, use) => {
		const user = await createUser(pocketBase);
		await use(user);
		await deleteUser(pocketBase, user);
	}
});
