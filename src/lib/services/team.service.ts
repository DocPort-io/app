import type { TeamCreateSchema, TeamSchema, TeamUpdateSchema } from '$lib/schemas/team.schema';

import { getPocketBase, type TypedPocketBase } from './pocketbase';

export interface ITeamService {
	findAll(): Promise<TeamSchema[]>;
	create(data: TeamCreateSchema): Promise<TeamSchema>;
	update(id: string, data: TeamUpdateSchema): Promise<TeamSchema>;
	remove(id: string): Promise<void>;
}

export class TeamService implements ITeamService {
	constructor(protected readonly pocketbase: TypedPocketBase = getPocketBase()) {}

	async findAll(): Promise<TeamSchema[]> {
		const records = await this.pocketbase.collection('teams').getList(1, 50, {
			sort: '-created'
		});

		return records.items;
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
