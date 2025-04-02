import { expect } from '@playwright/test';

import { login } from './flows/login';
import { test } from './framework/test';

test('As a user, I can login using my email and password', async ({
	page,
	loginPage,
	dashboardPage,
	testUser
}) => {
	await test.step('Login', async () => {
		await login({ testUser, page, loginPage, dashboardPage });
	});

	await test.step('Check if user is logged in', async () => {
		await expect(page).toHaveURL(dashboardPage.href);
	});
});

test('As a user, I can see an error message when I try to login with invalid credentials', async ({
	page,
	loginPage,
	testUser
}) => {
	await test.step('Login', async () => {
		await page.goto(loginPage.href);

		await loginPage.inputEmail.fill(testUser.email);
		await loginPage.inputPassword.fill('wrong-password');
		await loginPage.buttonSignIn.click();
	});

	await test.step('Check if user is not logged in', async () => {
		await expect(page).toHaveURL(loginPage.href);
	});

	await test.step('Check if error message is displayed', async () => {
		await expect(loginPage.textErrorTitle).toBeVisible();
		await expect(loginPage.textErrorTitle).toHaveText('Oops!');
		await expect(loginPage.textErrorDescription).toBeVisible();
		await expect(loginPage.textErrorDescription).toHaveText(
			'Invalid email or password. Please try again.'
		);
	});
});
