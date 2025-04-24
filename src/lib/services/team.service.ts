import type { TeamCreateSchema, TeamSchema, TeamUpdateSchema } from '$lib/schemas/team.schema';
import type { ListResult } from 'pocketbase';

import { getPocketBase, type TypedPocketBase } from './pocketbase';

export type FindAllOptions = {
	page?: number;
	perPage?: number;
};

export interface ITeamService {
	findAll(options?: FindAllOptions): Promise<ListResult<TeamSchema>>;
	create(data: TeamCreateSchema): Promise<TeamSchema>;
	update(id: string, data: TeamUpdateSchema): Promise<TeamSchema>;
	remove(id: string): Promise<void>;
}

export class TeamService implements ITeamService {
	constructor(protected readonly pocketbase: TypedPocketBase = getPocketBase()) {}

	async findAll({ page, perPage }: FindAllOptions = {}): Promise<ListResult<TeamSchema>> {
		page = page ?? 1;
		perPage = perPage ?? 50;

		return await this.pocketbase.collection('teams').getList(page, perPage, {
			sort: '-created'
		});
	}

	async create(data: TeamCreateSchema): Promise<TeamSchema> {
		return await this.pocketbase.collection('teams').create(data);
	}

	async update(id: string, data: TeamUpdateSchema): Promise<TeamSchema> {
		return await this.pocketbase.collection('teams').update(id, data);
	}

	async remove(id: string): Promise<void> {
		await this.pocketbase.collection('teams').delete(id);
	}
}
