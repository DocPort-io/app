<script lang="ts">
	import { createQuery } from '@tanstack/svelte-query';
	import { Button } from '$lib/components/ui/button';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { createPaginatedTeamsQuery } from '$lib/queries/teams';
	import { getTeamState } from '$lib/stores/team.svelte';
	import { cn } from '$lib/utils';

	const teamState = getTeamState();

	const teams = $derived(createQuery(createPaginatedTeamsQuery()));

	const team = $derived($teams.data?.items.find((team) => team.id === teamState.currentTeam));
</script>

<DropdownMenu.Root>
	<DropdownMenu.Trigger>
		<Button variant="outline" size="sm" class="gap-1" data-testid="team-switcher-button">
			<span data-testid="team-switcher-button-text">{team?.name || 'No team selected'}</span>
			<span class="sr-only">Switch team</span>
		</Button>
	</DropdownMenu.Trigger>
	<DropdownMenu.Content align="end">
		<DropdownMenu.Label>Teams</DropdownMenu.Label>
		<DropdownMenu.Separator />
		{#if $teams.isSuccess}
			{#each $teams.data.items as team}
				<DropdownMenu.Item
					class={cn(team.id === teamState.currentTeam && 'bg-accent')}
					onclick={() => teamState.selectTeam(team)}
					data-testid="team-switcher-item"
				>
					{team.name}
				</DropdownMenu.Item>
			{/each}
		{/if}
	</DropdownMenu.Content>
</DropdownMenu.Root>
