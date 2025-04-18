import type {
	ProjectCreateSchema,
	ProjectDeleteSchema,
	ProjectSchema,
	ProjectUpdateSchema
} from '$lib/schemas/project.schema';

import { ProjectService, type IProjectService } from '$lib/services/project.service';
import { getContext, setContext } from 'svelte';
import { toast } from 'svelte-sonner';

import { createPaginationController, type IPaginationController } from './pagination.svelte';
import { getTeamState, TeamState } from './team.svelte';
import { getUserState, type UserState } from './user.svelte';

export interface IProjectsState {
	projects: ProjectSchema[];
	loading: boolean;
	error: string | null;
	pagination: IPaginationController;
	add: (project: ProjectCreateSchema) => Promise<void>;
	edit: (id: string, project: ProjectUpdateSchema) => Promise<void>;
	remove: (project: ProjectDeleteSchema) => Promise<void>;
	getAll: () => Promise<void>;
}

export class ProjectsState implements IProjectsState {
	projects = $state<ProjectSchema[]>([]);
	pagination = createPaginationController({ page: 1, perPage: 25 });
	loading = $state(false);
	error = $state<string | null>(null);

	constructor(
		protected readonly service: IProjectService = new ProjectService(),
		protected readonly teamState: TeamState = getTeamState(),
		protected readonly userState: UserState = getUserState()
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
			if (!userState.isValid) return;
			void this.getAll();
		});
	}

	async getAll() {
		console.log('ProjectsState.getAll()');
		if (!this.teamState.selectedTeam) return;

		this.loading = true;
		this.error = null;

		const { page, perPage } = this.pagination;
		const { id: team } = this.teamState.selectedTeam;

		return this.service
			.getProjects({
				page,
				perPage,
				team
			})
			.then(({ items, totalItems, totalPages }) => {
				this.projects = items;
				this.pagination.totalItems = totalItems;
				this.pagination.totalPages = totalPages;
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
