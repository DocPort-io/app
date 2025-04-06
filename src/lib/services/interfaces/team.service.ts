import type { TeamCreateSchema, TeamSchema, TeamUpdateSchema } from '$lib/schemas/team.schema';

export interface ITeamService {
	getTeams(): Promise<TeamSchema[]>;
	createTeam(data: TeamCreateSchema): Promise<TeamSchema>;
	updateTeam(id: string, data: TeamUpdateSchema): Promise<TeamSchema>;
	deleteTeam(id: string): Promise<void>;
}
