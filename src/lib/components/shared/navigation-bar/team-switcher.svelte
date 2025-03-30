<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { getTeamState } from '$lib/stores/team.svelte';
	import { cn } from '$lib/utils';

	const teamState = getTeamState();
</script>

<DropdownMenu.Root>
	<DropdownMenu.Trigger asChild let:builder>
		<Button variant="outline" size="sm" builders={[builder]} class="gap-1">
			<span>{teamState.selectedTeam?.name || 'No team selected'}</span>
			<span class="sr-only">Switch team</span>
		</Button>
	</DropdownMenu.Trigger>
	<DropdownMenu.Content align="end">
		<DropdownMenu.Label>Teams</DropdownMenu.Label>
		<DropdownMenu.Separator />
		{#each teamState.teams as team}
			<DropdownMenu.Item
				class={cn(team.id === teamState.selectedTeam?.id && 'bg-accent')}
				on:click={() => teamState.selectTeam(team)}
			>
				{team.name}
			</DropdownMenu.Item>
		{/each}
	</DropdownMenu.Content>
</DropdownMenu.Root>
