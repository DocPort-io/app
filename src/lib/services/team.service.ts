import type { TeamCreateSchema, TeamSchema, TeamUpdateSchema } from '$lib/schemas/team.schema';

import type { ITeamService } from './interfaces/team.service';

import { getPocketBase, type TypedPocketBase } from './pocketbase';

export class TeamService implements ITeamService {
	constructor(protected readonly pocketbase: TypedPocketBase = getPocketBase()) {}

	async getTeams(): Promise<TeamSchema[]> {
		const records = await this.pocketbase.collection('teams').getList(1, 50, {
			sort: '-created'
		});

		return records.items;
	}

	async createTeam(data: TeamCreateSchema): Promise<TeamSchema> {
		return await this.pocketbase.collection('teams').create(data);
	}

	async updateTeam(id: string, data: TeamUpdateSchema): Promise<TeamSchema> {
		return await this.pocketbase.collection('teams').update(id, data);
	}

	async deleteTeam(id: string): Promise<void> {
		await this.pocketbase.collection('teams').delete(id);
	}
}
