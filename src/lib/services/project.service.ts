import type {
	ProjectCreateSchema,
	ProjectSchema,
	ProjectUpdateSchema
} from '$lib/schemas/project.schema';

import PocketBase from 'pocketbase';

import type { IProjectService } from './interfaces/project-service.interface';

export class ProjectService implements IProjectService {
	constructor(
		protected readonly pocketBaseClient: PocketBase = new PocketBase('http://127.0.0.1:8090')
	) {}

	async getProjects(): Promise<ProjectSchema[]> {
		const records = await this.pocketBaseClient
			.collection<ProjectSchema>('projects')
			.getList(1, 50, {
				sort: '-created'
			});

		return records.items;
	}

	async createProject(data: ProjectCreateSchema): Promise<ProjectSchema> {
		return await this.pocketBaseClient.collection<ProjectSchema>('projects').create(data);
	}

	async updateProject(id: string, data: ProjectUpdateSchema): Promise<ProjectSchema> {
		return await this.pocketBaseClient.collection<ProjectSchema>('projects').update(id, data);
	}

	async deleteProject(id: string): Promise<void> {
		await this.pocketBaseClient.collection<ProjectSchema>('projects').delete(id);
	}
}
