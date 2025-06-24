<script lang="ts">
	import { Folders } from '@lucide/svelte';
	import { createQuery } from '@tanstack/svelte-query';
	import UserPageLayout from '$lib/components/layouts/user-page-layout.svelte';
	import * as Card from '$lib/components/ui/card';
	import { m } from '$lib/paraglide/messages';
	import { createProjectsCountQuery } from '$lib/queries/projects';
	import { getTeamState } from '$lib/stores/team.svelte';

	const teamState = getTeamState();

	const projectsCountQuery = $derived.by(() =>
		createQuery(
			createProjectsCountQuery({
				team: teamState.currentTeam ?? ''
			})
		)
	);
</script>

<UserPageLayout title={m.dashboard()}>
	<div class="grid gap-4 md:grid-cols-2 md:gap-8 lg:grid-cols-4">
		<Card.Root>
			<Card.Header class="flex flex-row items-center justify-between space-y-0 pb-2">
				<Card.Title class="text-sm font-medium">{m.total_projects()}</Card.Title>
				<Folders class="text-muted-foreground h-4 w-4" />
			</Card.Header>
			<Card.Content>
				<div class="text-2xl font-bold">
					{#if $projectsCountQuery.isLoading}
						<span class="animate-pulse">{m.loading()}</span>
					{:else if $projectsCountQuery.isError}
						<span>{m.error()}</span>
					{:else if $projectsCountQuery.data === 0}
						<span>{m.no_projects_yet()}</span>
					{:else}
						{$projectsCountQuery.data}
					{/if}
				</div>
				<p class="text-muted-foreground text-xs">{m.good_job()}</p>
			</Card.Content>
		</Card.Root>
	</div>
</UserPageLayout>
