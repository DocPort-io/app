<script lang="ts">
	import type { Snippet } from 'svelte';

	import { createQuery } from '@tanstack/svelte-query';
	import { createPaginatedTeamsQuery } from '$lib/queries/teams';
	import { setTeamState } from '$lib/stores/team.svelte';

	type Props = {
		children: Snippet;
	};

	let { children }: Props = $props();

	const teamState = setTeamState();

	const teams = $derived(createQuery(createPaginatedTeamsQuery()));

	$effect(() => {
		if (teamState.currentTeam !== null) return;
		if ($teams.data === undefined) return;
		if ($teams.data?.items.length === 0) return;

		teamState.selectTeam($teams.data.items[0]);
	});
</script>

{@render children()}
