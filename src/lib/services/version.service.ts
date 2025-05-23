import type {
	VersionCreateSchema,
	VersionSchema,
	VersionUpdateSchema
} from '$lib/schemas/version.schema';
import type { ListResult } from 'pocketbase';

import { getPocketBase, type TypedPocketBase } from './pocketbase';

export type FindAllOptions = {
	page?: number;
	perPage?: number;
	project: string;
};

export interface IVersionService {
	findAll(options: FindAllOptions): Promise<ListResult<VersionSchema>>;
	findOne(id: string): Promise<VersionSchema>;
	create(data: VersionCreateSchema): Promise<VersionSchema>;
	update(id: string, data: VersionUpdateSchema): Promise<VersionSchema>;
	remove(id: string): Promise<void>;
}

export class VersionService implements IVersionService {
	constructor(protected readonly pocketbase: TypedPocketBase = getPocketBase()) {}

	async findAll({ page, perPage, project }: FindAllOptions): Promise<ListResult<VersionSchema>> {
		if (!project) throw new Error('Project is required');
		page = page ?? 1;
		perPage = perPage ?? 50;

		return await this.pocketbase.collection('versions').getList(page, perPage, {
			sort: '-created',
			filter: this.pocketbase.filter('project = {:project}', { project })
		});
	}

	async findOne(id: string): Promise<VersionSchema> {
		return await this.pocketbase.collection('versions').getOne(id);
	}

	async create(data: VersionCreateSchema): Promise<VersionSchema> {
		return await this.pocketbase.collection('versions').create(data);
	}

	async update(id: string, data: VersionUpdateSchema): Promise<VersionSchema> {
		return await this.pocketbase.collection('versions').update(id, data);
	}

	async remove(id: string): Promise<void> {
		await this.pocketbase.collection('versions').delete(id);
	}
}
