import { BasePage } from '../framework/base.page';

export class ProjectsPage extends BasePage {
	readonly href = '/projects';

	readonly cardProjects = this.page.getByTestId('projects-card');
	readonly tableProjects = this.page.getByTestId('projects-table');
	readonly tableProjectsHeader = this.page.getByTestId('projects-table-header');
	readonly tableProjectsRow = this.page.getByTestId('projects-table-row');
	readonly buttonCreateProject = this.page.getByTestId('projects-create-button');
	readonly modalCreateProject = this.page.getByTestId('projects-create-dialog');
	readonly inputProjectName = this.page.getByTestId('projects-create-dialog-input-name');
	readonly selectTriggerProjectStatus = this.page.getByTestId(
		'projects-create-dialog-select-trigger-status'
	);
	readonly selectContentProjectStatus = this.page.getByTestId(
		'projects-create-dialog-select-content-status'
	);
	readonly buttonCreateProjectSubmit = this.page.getByTestId(
		'projects-create-dialog-button-submit'
	);
}
