import type {
	ProjectCreateSchema,
	ProjectDeleteSchema,
	ProjectSchema,
	ProjectUpdateSchema
} from '$lib/schemas/project.schema';
import type { IProjectService } from '$lib/services/interfaces/project-service.interface';

import { ProjectService } from '$lib/services/project.service';
import { getContext, onMount, setContext } from 'svelte';
import { toast } from 'svelte-sonner';

export class Projects {
	projects = $state<ProjectSchema[]>([]);
	loading = $state(false);
	error = $state<string | null>(null);

	constructor(protected readonly service: IProjectService = new ProjectService()) {
		onMount(() => {
			void this.getAll();
		});

		$effect(() => {
			if (!this.error) return;
			toast.error('An error occurred', {
				description:
					"Your projects could not be loaded. We're sorry for the inconvenience. Please try again later.",
				duration: 10_000
			});
		});
	}

	async getAll() {
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
		await this.getAll();
	}

	async edit(id: string, project: ProjectUpdateSchema) {
		await this.service.updateProject(id, project);
		await this.getAll();
	}

	async remove(project: ProjectDeleteSchema) {
		await this.service.deleteProject(project.id);
		await this.getAll();
	}
}

const PROJECTS_KEY = Symbol('PROJECTS');

export const setProjects = () => {
	return setContext(PROJECTS_KEY, new Projects());
};

export const getProjects = () => {
	return getContext<ReturnType<typeof setProjects>>(PROJECTS_KEY);
};
