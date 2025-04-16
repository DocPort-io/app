import type {
	ProjectCreateSchema,
	ProjectSchema,
	ProjectUpdateSchema
} from '$lib/schemas/project.schema';

import { getPocketBase, type TypedPocketBase } from './pocketbase';

export type ProjectServiceGetProjectsOptions = {
	page?: number;
	perPage?: number;
	team: string;
};

export type ProjectServiceGetProjectsResult = {
	items: ProjectSchema[];
	totalItems: number;
	totalPages: number;
};

export interface IProjectService {
	getProjects(options: ProjectServiceGetProjectsOptions): Promise<ProjectServiceGetProjectsResult>;
	createProject(data: ProjectCreateSchema): Promise<ProjectSchema>;
	updateProject(id: string, data: ProjectUpdateSchema): Promise<ProjectSchema>;
	deleteProject(id: string): Promise<void>;
}

export class ProjectService implements IProjectService {
	constructor(protected readonly pocketbase: TypedPocketBase = getPocketBase()) {}

	async getProjects({
		page = 1,
		perPage = 5,
		team
	}: ProjectServiceGetProjectsOptions): Promise<ProjectServiceGetProjectsResult> {
		const records = await this.pocketbase.collection('projects').getList(page, perPage, {
			sort: '-created',
			filter: this.pocketbase.filter('team = {:team}', { team })
		});

		return {
			items: records.items,
			totalItems: records.totalItems,
			totalPages: records.totalPages
		};
	}

	async createProject(data: ProjectCreateSchema): Promise<ProjectSchema> {
		return await this.pocketbase.collection('projects').create(data);
	}

	async updateProject(id: string, data: ProjectUpdateSchema): Promise<ProjectSchema> {
		return await this.pocketbase.collection('projects').update(id, data);
	}

	async deleteProject(id: string): Promise<void> {
		await this.pocketbase.collection('projects').delete(id);
	}
}
