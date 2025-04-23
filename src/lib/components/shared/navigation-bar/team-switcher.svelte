<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { getTeamState } from '$lib/stores/team.svelte';
	import { cn } from '$lib/utils';

	const teamState = getTeamState();

	const team = $derived(teamState.teams.find((team) => team.id === teamState.currentTeam));
</script>

<DropdownMenu.Root>
	<DropdownMenu.Trigger asChild let:builder>
		<Button
			variant="outline"
			size="sm"
			builders={[builder]}
			class="gap-1"
			data-testid="team-switcher-button"
		>
			<span data-testid="team-switcher-button-text">{team?.name || 'No team selected'}</span>
			<span class="sr-only">Switch team</span>
		</Button>
	</DropdownMenu.Trigger>
	<DropdownMenu.Content align="end">
		<DropdownMenu.Label>Teams</DropdownMenu.Label>
		<DropdownMenu.Separator />
		{#each teamState.teams as team}
			<DropdownMenu.Item
				class={cn(team.id === teamState.currentTeam && 'bg-accent')}
				on:click={() => teamState.selectTeam(team)}
				data-testid="team-switcher-item"
			>
				{team.name}
			</DropdownMenu.Item>
		{/each}
	</DropdownMenu.Content>
</DropdownMenu.Root>
