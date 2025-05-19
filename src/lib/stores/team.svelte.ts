import type { TeamSchema } from '$lib/schemas/team.schema';

import { getContext, setContext } from 'svelte';

export interface ITeamState {
	currentTeam: string | null;
	selectTeam: (team: TeamSchema) => void;
}

export class TeamState implements ITeamState {
	currentTeam = $state<string | null>(null);

	constructor() {
		this.#loadTeamFromLocalStorage();
	}

	selectTeam(team: TeamSchema) {
		this.currentTeam = team.id;
		this.#saveTeamToLocalStorage();
	}

	#saveTeamToLocalStorage() {
		return this.currentTeam
			? localStorage.setItem('team', this.currentTeam)
			: localStorage.removeItem('team');
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
