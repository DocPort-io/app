import type {
	ProjectCreateSchema,
	ProjectSchema,
	ProjectUpdateSchema
} from '$lib/schemas/project.schema';

import type { IProjectService } from './interfaces/project-service.interface';

import { getPocketBase } from './pocketbase';

export class ProjectService implements IProjectService {
	constructor(protected readonly pocketBaseClient = getPocketBase()) {}

	async getProjects(): Promise<ProjectSchema[]> {
		const records = await this.pocketBaseClient.collection('projects').getList(1, 50, {
			sort: '-created'
		});

		return records.items;
	}

	async createProject(data: ProjectCreateSchema): Promise<ProjectSchema> {
		return await this.pocketBaseClient.collection('projects').create(data);
	}

	async updateProject(id: string, data: ProjectUpdateSchema): Promise<ProjectSchema> {
		return await this.pocketBaseClient.collection('projects').update(id, data);
	}

	async deleteProject(id: string): Promise<void> {
		await this.pocketBaseClient.collection('projects').delete(id);
	}
}
