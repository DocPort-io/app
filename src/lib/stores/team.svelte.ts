import type { TeamSchema } from '$lib/schemas/team.schema';

import { TeamService, type ITeamService } from '$lib/services/team.service';
import { getContext, setContext } from 'svelte';

import { getUserState, type UserState } from './user.svelte';

export interface ITeamState {
	teams: TeamSchema[];
	loading: boolean;
	error: string | null;
	currentTeam: string | null;
	load: () => void;
	selectTeam: (team: TeamSchema) => void;
}

export class TeamState implements ITeamState {
	teams = $state<TeamSchema[]>([]);
	loading = $state(false);
	error = $state<string | null>(null);
	currentTeam = $state<string | null>(null);

	constructor(
		protected readonly service: ITeamService = new TeamService(),
		protected readonly userState: UserState = getUserState()
	) {
		$effect(() => {
			if (this.currentTeam) return;
			if (this.teams.length === 0) return;
			this.selectTeam(this.teams[0]);
		});

		this.#loadTeamFromLocalStorage();

		$effect(() => {
			if (!userState.isValid) return;
			void this.#fetch();
		});
	}

	load() {
		this.#fetch();
	}

	async #fetch() {
		this.loading = true;
		this.error = null;

		return this.service
			.findAll()
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
		this.currentTeam = team.id;
		this.#saveTeamToLocalStorage();
	}

	#saveTeamToLocalStorage() {
		if (!this.currentTeam) {
			localStorage.removeItem('team');
			return;
		}

		localStorage.setItem('team', this.currentTeam);
	}

	#loadTeamFromLocalStorage() {
		const team = localStorage.getItem('team');
		this.currentTeam = team;
	}
}

const TEAM_STATE_KEY = Symbol('TEAM_STATE');

export const setTeamState = () => {
	return setContext(TEAM_STATE_KEY, new TeamState());
};

export const getTeamState = () => {
	return getContext<ReturnType<typeof setTeamState>>(TEAM_STATE_KEY);
};
