<script lang="ts">
	import type {
		ProjectCreateSchema,
		ProjectDeleteSchema,
		ProjectUpdateSchema
	} from '$lib/schemas/project.schema';

	import { CirclePlus } from '@lucide/svelte';
	import { goto } from '$app/navigation';
	import UserPageLayout from '$lib/components/layouts/user-page-layout.svelte';
	import Pagination from '$lib/components/pagination.svelte';
	import ResultsInfo from '$lib/components/results-info.svelte';
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import { AppRoute } from '$lib/constants';
	import { m } from '$lib/paraglide/messages';
	import { createDialogController } from '$lib/stores/dialog.svelte';
	import { getProjects } from '$lib/stores/projects.svelte';
	import { toast } from 'svelte-sonner';

	import CreateProjectDialog from './_components/_dialogs/create-project-dialog.svelte';
	import DeleteProjectDialog from './_components/_dialogs/delete-project-dialog.svelte';
	import EditProjectDialog from './_components/_dialogs/edit-project-dialog.svelte';
	import ProjectsTable from './_components/projects-table.svelte';

	const projectStore = getProjects();

	// Dialog handlers
	const createDialog = createDialogController();
	const editDialog = createDialogController<{ id: string; project: ProjectUpdateSchema }>();
	const deleteDialog = createDialogController<ProjectDeleteSchema>();

	// Action handlers
	const handleCreateProject = async (data: ProjectCreateSchema) => {
		await projectStore.add(data);
		toast.success('Project created successfully!');
	};

	const handleUpdateProject = async (id: string, data: ProjectUpdateSchema) => {
		await projectStore.edit(id, data);
		toast.success('Project updated successfully!');
	};

	const handleDeleteProject = async (data: ProjectDeleteSchema) => {
		await projectStore.remove(data);
		toast.success('Project deleted successfully!');
	};
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
				loading={projectStore.loading}
				error={projectStore.error}
				projects={projectStore.projects}
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
		{#if projectStore.pagination.totalItems > 0}
			<Card.Footer>
				<div class="flex w-full flex-col items-center justify-between md:flex-row">
					<ResultsInfo
						amountShown={projectStore.projects.length}
						pagination={projectStore.pagination}
					/>
					<Pagination pagination={projectStore.pagination} />
				</div>
			</Card.Footer>
		{/if}
	</Card.Root>

	<CreateProjectDialog dialogController={createDialog} {handleCreateProject} />
	<EditProjectDialog dialogController={editDialog} {handleUpdateProject} />
	<DeleteProjectDialog dialogController={deleteDialog} {handleDeleteProject} />
</UserPageLayout>
