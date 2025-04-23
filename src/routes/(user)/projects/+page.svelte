<script lang="ts">
	import type { ProjectUpdateSchema } from '$lib/schemas/project.schema';

	import { CirclePlus } from '@lucide/svelte';
	import { createQuery } from '@tanstack/svelte-query';
	import { goto } from '$app/navigation';
	import UserPageLayout from '$lib/components/layouts/user-page-layout.svelte';
	import Pagination from '$lib/components/pagination.svelte';
	import ResultsInfo from '$lib/components/results-info.svelte';
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import { AppRoute } from '$lib/constants';
	import { m } from '$lib/paraglide/messages';
	import { createPaginatedProjectsQuery } from '$lib/queries/projects';
	import { createDialogController } from '$lib/stores/dialog.svelte';
	import { createPaginationController } from '$lib/stores/pagination.svelte';
	import { getTeamState } from '$lib/stores/team.svelte';

	import CreateProjectDialog from './_components/_dialogs/create-project-dialog.svelte';
	import DeleteProjectDialog from './_components/_dialogs/delete-project-dialog.svelte';
	import EditProjectDialog from './_components/_dialogs/edit-project-dialog.svelte';
	import ProjectsTable from './_components/projects-table.svelte';

	// Dialog handlers
	const createDialog = createDialogController();
	const editDialog = createDialogController<{ id: string; project: ProjectUpdateSchema }>();
	const deleteDialog = createDialogController<{ id: string }>();

	const teamState = getTeamState();

	const pagination = createPaginationController({
		page: 1,
		perPage: 2
	});

	const projectsQuery = $derived.by(() =>
		createQuery(
			createPaginatedProjectsQuery({
				team: teamState.currentTeam ?? '',
				page: pagination.page,
				perPage: pagination.perPage
			})
		)
	);
</script>

<UserPageLayout title="Projects">
	<Card.Root data-testid="projects-card">
		<Card.Header class="flex flex-row items-center">
			<div class="grid gap-2">
				<Card.Title>{m.projects()}</Card.Title>
				<Card.Description>{m.only_nimble_martin_strive()}</Card.Description>
			</div>
			<div class="ml-auto flex items-center gap-2">
				<Button size="sm" class="h-8 gap-1" on:click={() => createDialog.open()}>
					<CirclePlus class="h-3.5 w-3.5" />
					<span class="sr-only sm:not-sr-only sm:whitespace-nowrap"
						>{m.stout_elegant_jan_flip()}</span
					>
				</Button>
			</div>
		</Card.Header>
		<Card.Content>
			<ProjectsTable
				loading={$projectsQuery.isLoading}
				error={$projectsQuery.error?.message ?? null}
				projects={$projectsQuery.data?.items ?? []}
				handleViewProject={(project) => {
					goto(AppRoute.PROJECT_VIEW(project.id));
				}}
				handleEditProject={(id, project) => {
					editDialog.data = { id, project };
					editDialog.open();
				}}
				handleDeleteProject={(project) => {
					deleteDialog.data = project;
					deleteDialog.open();
				}}
			/>
		</Card.Content>
		{#if $projectsQuery.data?.totalItems && $projectsQuery.data.totalItems > 0}
			<Card.Footer>
				<div class="flex w-full flex-col items-center justify-between md:flex-row">
					<ResultsInfo
						results={$projectsQuery.data?.items.length ?? 0}
						page={$projectsQuery.data?.page ?? 0}
						perPage={$projectsQuery.data?.perPage ?? 0}
						totalItems={$projectsQuery.data?.totalItems ?? 0}
					/>
					<Pagination {pagination} totalItems={$projectsQuery.data?.totalItems ?? 0} />
				</div>
			</Card.Footer>
		{/if}
	</Card.Root>

	<CreateProjectDialog dialogController={createDialog} />
	<EditProjectDialog dialogController={editDialog} />
	<DeleteProjectDialog dialogController={deleteDialog} />
</UserPageLayout>
