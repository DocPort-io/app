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
	page: number;
	perPage: number;
	totalItems: number;
	totalPages: number;
};

export interface IProjectService {
	findAll(options: ProjectServiceGetProjectsOptions): Promise<ProjectServiceGetProjectsResult>;
	findOne(id: string): Promise<ProjectSchema>;
	create(data: ProjectCreateSchema): Promise<ProjectSchema>;
	update(id: string, data: ProjectUpdateSchema): Promise<ProjectSchema>;
	remove(id: string): Promise<void>;
}

export class ProjectService implements IProjectService {
	constructor(protected readonly pocketbase: TypedPocketBase = getPocketBase()) {}

	async findAll({
		page = 1,
		perPage = 5,
		team
	}: ProjectServiceGetProjectsOptions): Promise<ProjectServiceGetProjectsResult> {
		return await this.pocketbase.collection('projects').getList(page, perPage, {
			sort: '-created',
			filter: this.pocketbase.filter('team = {:team}', { team })
		});
	}

	async findOne(id: string): Promise<ProjectSchema> {
		return await this.pocketbase.collection('projects').getOne(id);
	}

	async create(data: ProjectCreateSchema): Promise<ProjectSchema> {
		return await this.pocketbase.collection('projects').create(data);
	}

	async update(id: string, data: ProjectUpdateSchema): Promise<ProjectSchema> {
		return await this.pocketbase.collection('projects').update(id, data);
	}

	async remove(id: string): Promise<void> {
		await this.pocketbase.collection('projects').delete(id);
	}
}
