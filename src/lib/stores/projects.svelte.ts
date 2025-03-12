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
	loadingPromise = $state<Promise<void> | undefined>();

	constructor(protected readonly service: IProjectService = new ProjectService()) {}

	load() {
		this.loadingPromise = this.#fetch();
	}

	#fetch = async () => {
		await new Promise((resolve) => setTimeout(resolve, 500));
		this.projects = await this.service.getProjects();
	};

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
