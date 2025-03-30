<script lang="ts">
	import type {
		ProjectCreateSchema,
		ProjectDeleteSchema,
		ProjectUpdateSchema
	} from '$lib/schemas/project.schema';

	import { CirclePlus, ListFilter } from '@lucide/svelte';
	import { goto } from '$app/navigation';
	import UserPageLayout from '$lib/components/layouts/user-page-layout.svelte';
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import * as Pagination from '$lib/components/ui/pagination';
	import { AppRoute } from '$lib/constants';
	import { m } from '$lib/paraglide/messages';
	import { createDialogController } from '$lib/stores/dialog.svelte';
	import { getProjects } from '$lib/stores/projects.svelte';
	import { toast } from 'svelte-sonner';

	import CreateProjectDialog from './_components/_dialogs/create-project-dialog.svelte';
	import DeleteProjectDialog from './_components/_dialogs/delete-project-dialog.svelte';
	import EditProjectDialog from './_components/_dialogs/edit-project-dialog.svelte';
	import ProjectsTable from './_components/projects-table.svelte';

	let filterActive = $state(true);
	let filterCompleted = $state(false);

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
	<Card.Root>
		<Card.Header class="flex flex-row items-center">
			<div class="grid gap-2">
				<Card.Title>{m.projects()}</Card.Title>
				<Card.Description>{m.only_nimble_martin_strive()}</Card.Description>
			</div>
			<div class="ml-auto flex items-center gap-2">
				<DropdownMenu.Root>
					<DropdownMenu.Trigger asChild let:builder>
						<Button builders={[builder]} variant="outline" size="sm" class="h-8 gap-1">
							<ListFilter class="h-3.5 w-3.5" />
							<span class="sr-only sm:not-sr-only sm:whitespace-nowrap"
								>{m.that_weary_anteater_push()}</span
							>
						</Button>
					</DropdownMenu.Trigger>
					<DropdownMenu.Content align="end">
						<DropdownMenu.Label>{m.vexed_steep_piranha_kick()}</DropdownMenu.Label>
						<DropdownMenu.Separator />
						<DropdownMenu.CheckboxItem bind:checked={filterActive}>
							{m.alive_ok_kangaroo_boil()}
						</DropdownMenu.CheckboxItem>
						<DropdownMenu.CheckboxItem bind:checked={filterCompleted}>
							Completed
						</DropdownMenu.CheckboxItem>
					</DropdownMenu.Content>
				</DropdownMenu.Root>
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
				{projectStore}
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
		<Card.Footer>
			<div class="flex w-full items-center justify-between">
				<div class="text-muted-foreground text-xs">
					{m.weird_sharp_javelina_sail({
						amount: projectStore.projects.length,
						start: 1,
						end: projectStore.projects.length
					})}
				</div>
				<div>
					<Pagination.Root count={100} perPage={10} let:pages let:currentPage>
						<Pagination.Content>
							<Pagination.Item>
								<Pagination.PrevButton />
							</Pagination.Item>
							{#each pages as page (page.key)}
								{#if page.type === 'ellipsis'}
									<Pagination.Item>
										<Pagination.Ellipsis />
									</Pagination.Item>
								{:else}
									<Pagination.Item>
										<Pagination.Link {page} isActive={currentPage == page.value}>
											{page.value}
										</Pagination.Link>
									</Pagination.Item>
								{/if}
							{/each}
							<Pagination.Item>
								<Pagination.NextButton />
							</Pagination.Item>
						</Pagination.Content>
					</Pagination.Root>
				</div>
			</div>
		</Card.Footer>
	</Card.Root>

	<CreateProjectDialog dialogController={createDialog} {handleCreateProject} />
	<EditProjectDialog dialogController={editDialog} {handleUpdateProject} />
	<DeleteProjectDialog dialogController={deleteDialog} {handleDeleteProject} />
</UserPageLayout>
