import type {
	ProjectCreateSchema,
	ProjectSchema,
	ProjectUpdateSchema
} from '$lib/schemas/project.schema';
import type { ListResult } from 'pocketbase';

import { getPocketBase, type TypedPocketBase } from './pocketbase';

export type FindAllOptions = {
	page?: number;
	perPage?: number;
	team: string;
};

export interface IProjectService {
	findAll(options: FindAllOptions): Promise<ListResult<ProjectSchema>>;
	findOne(id: string): Promise<ProjectSchema>;
	create(data: ProjectCreateSchema): Promise<ProjectSchema>;
	update(id: string, data: ProjectUpdateSchema): Promise<ProjectSchema>;
	remove(id: string): Promise<void>;
}

export class ProjectService implements IProjectService {
	constructor(protected readonly pocketbase: TypedPocketBase = getPocketBase()) {}

	async findAll({ page, perPage, team }: FindAllOptions): Promise<ListResult<ProjectSchema>> {
		if (!team) throw new Error('Team is required');
		page = page ?? 1;
		perPage = perPage ?? 50;

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
