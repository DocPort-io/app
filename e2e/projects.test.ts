import { expect } from '@playwright/test';

import { login } from './flows/login';
import { createProject, deleteProject, TestProject } from './framework/fixtures/project';
import { test } from './framework/test';

const projects: TestProject[] = [];

test.beforeEach(async ({ testTeam, pocketBase }) => {
	await Promise.all(
		Array.from({ length: 500 }).map(async () => {
			const project = await createProject(pocketBase, testTeam);
			projects.push(project);
		})
	);
});

test.afterEach(async ({ pocketBase }) => {
	await Promise.all(
		projects.map(async (project) => {
			await deleteProject(pocketBase, project);
		})
	);
});

test('As a user, I can see a list of projects', async ({
	page,
	loginPage,
	dashboardPage,
	projectsPage,
	testUser,
	testTeam
}) => {
	await test.step('Login', async () => {
		await login({ testUser, page, loginPage, dashboardPage });
	});

	await test.step('Check if user is logged in', async () => {
		await expect(page).toHaveURL(dashboardPage.href);
	});

	await test.step('Check if team is selected', async () => {
		await expect(page.getByTestId('team-switcher-button-text')).toHaveText(testTeam.name);
	});

	await test.step('Go to projects page', async () => {
		await page.goto(projectsPage.href);
	});

	await test.step('Check if projects page is loaded', async () => {
		await expect(page).toHaveURL(projectsPage.href);
	});

	await test.step('Check if projects card is visible', async () => {
		await expect(projectsPage.cardProjects).toBeVisible();
	});

	await test.step('Check if projects table is visible', async () => {
		await expect(projectsPage.tableProjects).toBeVisible();
	});

	await test.step('Check if projects table header is visible', async () => {
		await expect(projectsPage.tableProjectsHeader).toBeVisible();
	});

	await test.step('Check if projects table has rows', async () => {
		const rows = await projectsPage.tableProjectsRow.count();
		expect(rows).toBeGreaterThan(0);
	});
});
