import type { FileCreateSchema, FileSchema, FileUpdateSchema } from '$lib/schemas/file.schema';
import type { ListResult } from 'pocketbase';

import { getPocketBase, type TypedPocketBase } from './pocketbase';

export type FindAllOptions = {
	page?: number;
	perPage?: number;
	version: string;
};

export interface IFileService {
	findAll(options: FindAllOptions): Promise<ListResult<FileSchema>>;
	findOne(id: string): Promise<FileSchema>;
	create(data: FileCreateSchema): Promise<FileSchema>;
	update(id: string, data: FileUpdateSchema): Promise<FileSchema>;
	remove(id: string): Promise<void>;
}

export class FileService implements IFileService {
	constructor(protected readonly pocketbase: TypedPocketBase = getPocketBase()) {}

	async findAll({ page, perPage, version }: FindAllOptions): Promise<ListResult<FileSchema>> {
		if (!version) throw new Error('Version is required');
		page = page ?? 1;
		perPage = perPage ?? 50;

		return await this.pocketbase.collection('files').getList(page, perPage, {
			sort: '-created',
			filter: this.pocketbase.filter('versions.id ?= {:version}', { version })
		});
	}

	async findOne(id: string): Promise<FileSchema> {
		return await this.pocketbase.collection('files').getOne(id);
	}

	async create(data: FileCreateSchema): Promise<FileSchema> {
		return await this.pocketbase.collection('files').create(data);
	}

	async update(id: string, data: FileUpdateSchema): Promise<FileSchema> {
		return await this.pocketbase.collection('files').update(id, data);
	}

	async remove(id: string): Promise<void> {
		await this.pocketbase.collection('files').delete(id);
	}
}
