import { BasePage } from '../framework/base.page';

export class ProjectsPage extends BasePage {
	readonly href = '/projects';

	readonly cardProjects = this.page.getByTestId('projects-card');
	readonly tableProjects = this.page.getByTestId('projects-table');
	readonly tableProjectsHeader = this.page.getByTestId('projects-table-header');
	readonly tableProjectsRow = this.page.getByTestId('projects-table-row');
}
