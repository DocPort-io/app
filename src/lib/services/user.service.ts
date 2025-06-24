import type { UserCreateSchema, UserSchema, UserUpdateSchema } from '$lib/schemas/user.schema';
import type { ListResult } from 'pocketbase';

import { getPocketBase, type TypedPocketBase } from './pocketbase';

export type FindAllOptions = {
	page?: number;
	perPage?: number;
};

export interface IUserService {
	findAll(options?: FindAllOptions): Promise<ListResult<UserSchema>>;
	findOne(id: string): Promise<UserSchema>;
	create(data: UserCreateSchema): Promise<UserSchema>;
	update(id: string, data: UserUpdateSchema): Promise<UserSchema>;
	remove(id: string): Promise<void>;
}

export class UserService implements IUserService {
	constructor(protected readonly pocketbase: TypedPocketBase = getPocketBase()) {}

	async findAll({ page, perPage }: FindAllOptions = {}): Promise<ListResult<UserSchema>> {
		page = page ?? 1;
		perPage = perPage ?? 50;

		return await this.pocketbase.collection('users').getList(page, perPage, {
			sort: '-created'
		});
	}

	async findOne(id: string): Promise<UserSchema> {
		return await this.pocketbase.collection('users').getOne(id);
	}

	async create(data: UserCreateSchema): Promise<UserSchema> {
		return await this.pocketbase.collection('users').create(data);
	}

	async update(id: string, data: UserUpdateSchema): Promise<UserSchema> {
		return await this.pocketbase.collection('users').update(id, data);
	}

	async remove(id: string): Promise<void> {
		await this.pocketbase.collection('users').delete(id);
	}
}
