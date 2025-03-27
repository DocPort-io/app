import type {
	ProjectCreateSchema,
	ProjectDeleteSchema,
	ProjectSchema,
	ProjectUpdateSchema
} from '$lib/schemas/project.schema';
import type { IProjectService } from '$lib/services/interfaces/project-service.interface';

import { ProjectService } from '$lib/services/project.service';
import { getContext, setContext } from 'svelte';

export class Projects {
	projects = $state<ProjectSchema[]>([]);
	loading = $state(false);
	error = $state<string | null>(null);

	constructor(protected readonly service: IProjectService = new ProjectService()) {}

	load() {
		this.#fetch();
	}

	async #fetch() {
		this.loading = true;
		this.error = null;

		return this.service
			.getProjects()
			.then((projects) => {
				this.projects = projects;
				this.loading = false;
			})
			.catch((err) => {
				this.error = err instanceof Error ? err.message : err;
				this.loading = false;
			});
	}

	async add(project: ProjectCreateSchema) {
		await this.service.createProject(project);
		await this.#fetch();
	}

	async edit(id: string, project: ProjectUpdateSchema) {
		await this.service.updateProject(id, project);
		await this.#fetch();
	}

	async remove(project: ProjectDeleteSchema) {
		await this.service.deleteProject(project.id);
		await this.#fetch();
	}
}

const PROJECTS_KEY = Symbol('PROJECTS');

export const setProjects = () => {
	return setContext(PROJECTS_KEY, new Projects());
};

export const getProjects = () => {
	return getContext<ReturnType<typeof setProjects>>(PROJECTS_KEY);
};
