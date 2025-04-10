import type { TeamSchema } from '$lib/schemas/team.schema';
import type { ITeamService } from '$lib/services/interfaces/team.service';

import { TeamService } from '$lib/services/team.service';
import { getContext, setContext } from 'svelte';

export class TeamState {
	teams = $state<TeamSchema[]>([]);
	loading = $state(false);
	error = $state<string | null>(null);
	selectedTeam = $state<TeamSchema | null>(null);

	constructor(protected readonly service: ITeamService = new TeamService()) {
		$effect(() => {
			if (this.selectedTeam) return;
			if (this.teams.length === 0) return;
			this.selectTeam(this.teams[0]);
		});

		this.#loadTeamFromLocalStorage();
	}

	load() {
		this.#fetch();
	}

	async #fetch() {
		this.loading = true;
		this.error = null;

		return this.service
			.getTeams()
			.then((teams) => {
				this.teams = teams;
				this.loading = false;
			})
			.catch((err) => {
				this.error = err instanceof Error ? err.message : err;
				this.loading = false;
			});
	}

	selectTeam(team: TeamSchema) {
		this.selectedTeam = team;
		this.#saveTeamToLocalStorage();
	}

	#saveTeamToLocalStorage() {
		localStorage.setItem('team', JSON.stringify(this.selectedTeam));
	}

	#loadTeamFromLocalStorage() {
		const team = localStorage.getItem('team');
		this.selectedTeam = team ? JSON.parse(team) : null;
	}
}

const TEAM_STATE_KEY = Symbol('TEAM_STATE');

export const setTeamState = () => {
	return setContext(TEAM_STATE_KEY, new TeamState());
};

export const getTeamState = () => {
	return getContext<ReturnType<typeof setTeamState>>(TEAM_STATE_KEY);
};
