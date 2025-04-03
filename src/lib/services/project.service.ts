import type {
	ProjectCreateSchema,
	ProjectSchema,
	ProjectUpdateSchema
} from '$lib/schemas/project.schema';

import type { IProjectService } from './interfaces/project-service.interface';

import { getPocketBase, type TypedPocketBase } from './pocketbase';

export class ProjectService implements IProjectService {
	constructor(protected readonly pocketbase: TypedPocketBase = getPocketBase()) {}

	async getProjects({
		page = 1,
		perPage = 5,
		filter
	}: {
		page?: number;
		perPage?: number;
		filter?: string;
	}): Promise<{ items: ProjectSchema[]; totalItems: number; totalPages: number }> {
		const records = await this.pocketbase.collection('projects').getList(page, perPage, {
			sort: '-created',
			filter
		});

		return {
			items: records.items as ProjectSchema[],
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
