import { expect } from '@playwright/test';

import { login } from './flows/login';
import { TestProject } from './framework/fixtures/project';
import { test } from './framework/test';

test.describe('Projects - List', () => {
	const projects: TestProject[] = [];

	test.beforeEach(async ({ testTeam, testData }) => {
		await Promise.all(
			Array.from({ length: 500 }).map(async () => {
				const project = await testData.project(testTeam);
				projects.push(project);
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
});

test.describe('Projects - Create', () => {
	test('As a user, I can create a project', async ({
		page,
		loginPage,
		dashboardPage,
		projectsPage,
		testUser,
		testTeam,
		testData
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

		await test.step('Click on create project button', async () => {
			await projectsPage.buttonCreateProject.click();
		});

		await test.step('Check if create project modal is visible', async () => {
			await expect(projectsPage.modalCreateProject).toBeVisible();
		});

		const projectName = `Test Project ${Date.now()}`;

		await test.step('Fill in project name', async () => {
			await projectsPage.inputProjectName.fill(projectName);
		});

		await test.step('Select project status', async () => {
			await projectsPage.selectTriggerProjectStatus.click();
			await projectsPage.selectContentProjectStatus.getByRole('option', { name: 'Active' }).click();
		});

		await test.step('Click on create project button', async () => {
			await projectsPage.buttonCreateProjectSubmit.click();
		});

		testData.markProjectAsDirty(testTeam);

		await test.step('Check if project is created', async () => {
			await expect(projectsPage.tableProjectsRow).toHaveCount(1);
		});

		await test.step('Check if project name is visible', async () => {
			await expect(projectsPage.tableProjectsRow.first().getByRole('cell').first()).toHaveText(
				projectName
			);
		});

		await test.step('Check if project status is visible', async () => {
			await expect(projectsPage.tableProjectsRow.first().getByRole('cell').nth(1)).toHaveText(
				'Active'
			);
		});
	});

	['Planned', 'Active', 'Completed'].forEach((status) => {
		test(`As a user, I can create a project with status ${status}`, async ({
			page,
			loginPage,
			dashboardPage,
			projectsPage,
			testUser,
			testTeam,
			testData
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

			await test.step('Click on create project button', async () => {
				await projectsPage.buttonCreateProject.click();
			});

			await test.step('Check if create project modal is visible', async () => {
				await expect(projectsPage.modalCreateProject).toBeVisible();
			});

			const projectName = `Test Project ${Date.now()}`;

			await test.step('Fill in project name', async () => {
				await projectsPage.inputProjectName.fill(projectName);
			});

			await test.step('Select project status', async () => {
				await projectsPage.selectTriggerProjectStatus.click();
				await projectsPage.selectContentProjectStatus
					.getByRole('option', { name: 'Active' })
					.click();
			});

			await test.step('Click on create project button', async () => {
				await projectsPage.buttonCreateProjectSubmit.click();
			});

			testData.markProjectAsDirty(testTeam);

			await test.step('Check if project is created', async () => {
				await expect(projectsPage.tableProjectsRow).toHaveCount(1);
			});

			await test.step('Check if project name is visible', async () => {
				await expect(projectsPage.tableProjectsRow.first().getByRole('cell').first()).toHaveText(
					projectName
				);
			});

			await test.step('Check if project status is visible', async () => {
				await expect(projectsPage.tableProjectsRow.first().getByRole('cell').nth(1)).toHaveText(
					'Active'
				);
			});
		});
	});
});
