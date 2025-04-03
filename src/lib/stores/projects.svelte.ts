import type {
	ProjectCreateSchema,
	ProjectDeleteSchema,
	ProjectSchema,
	ProjectUpdateSchema
} from '$lib/schemas/project.schema';
import type { IProjectService } from '$lib/services/interfaces/project-service.interface';

import { ProjectService } from '$lib/services/project.service';
import { getContext, setContext } from 'svelte';
import { toast } from 'svelte-sonner';

import { getTeamState, TeamState } from './team.svelte';

export class ProjectsState {
	projects = $state<ProjectSchema[]>([]);
	totalItems = $state<number>(0);
	totalPages = $state<number>(0);
	currentPage = $state<number>(1);
	perPage = $state<number>(20);
	filters = $state<{ active: boolean; completed: false }>({ active: true, completed: false });
	loading = $state(false);
	error = $state<string | null>(null);

	constructor(
		protected readonly service: IProjectService = new ProjectService(),
		protected readonly teamState: TeamState = getTeamState()
	) {
		$effect(() => {
			if (!this.error) return;
			toast.error('An error occurred', {
				description:
					"Your projects could not be loaded. We're sorry for the inconvenience. Please try again later.",
				duration: 10_000
			});
		});

		$effect(() => {
			void this.getAll();
		});
	}

	async getAll() {
		this.loading = true;
		this.error = null;

		const filterAnd: string[] = [];
		const statusOr = [];

		filterAnd.push(`team='${this.teamState.selectedTeam?.id}'`);

		if (this.filters.active) {
			statusOr.push('active');
		}

		if (this.filters.completed) {
			statusOr.push('completed');
		}

		if (statusOr.length) {
			filterAnd.push(`status='${statusOr.join("' || status='")}'`);
		}

		return this.service
			.getProjects({
				page: this.currentPage,
				perPage: this.perPage,
				filter: filterAnd.join('&&')
			})
			.then(({ items, totalItems, totalPages }) => {
				this.projects = items;
				this.totalItems = totalItems;
				this.totalPages = totalPages;
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
	return setContext(PROJECTS_KEY, new ProjectsState());
};

export const getProjects = () => {
	return getContext<ReturnType<typeof setProjects>>(PROJECTS_KEY);
};
