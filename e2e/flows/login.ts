import { Page } from '@playwright/test';

import { TestUser } from '../framework/fixtures/user';
import { LoginPage } from '../pages/auth/login.page';
import { DashboardPage } from '../pages/dashboard.page';

export type LoginOptions = {
	testUser: TestUser;
	page: Page;
	loginPage: LoginPage;
	dashboardPage: DashboardPage;
};

export const login = async ({ testUser, page, loginPage, dashboardPage }: LoginOptions) => {
	await page.goto(loginPage.href);

	await loginPage.inputEmail.fill(testUser.email);
	await loginPage.inputPassword.fill(testUser.password);
	await loginPage.buttonSignIn.click();

	await page.waitForURL(dashboardPage.href);
};
