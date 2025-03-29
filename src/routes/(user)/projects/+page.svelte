<script lang="ts">
	import type {
		ProjectCreateSchema,
		ProjectDeleteSchema,
		ProjectUpdateSchema
	} from '$lib/schemas/project.schema';

	import { CirclePlus, File, ListFilter } from '@lucide/svelte';
	import UserPageLayout from '$lib/components/layouts/user-page-layout.svelte';
	import * as Breadcrumb from '$lib/components/ui/breadcrumb';
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { AppRoute } from '$lib/constants';
	import { m } from '$lib/paraglide/messages';
	import { createDialogController } from '$lib/stores/dialog.svelte';
	import { getProjects } from '$lib/stores/projects.svelte';

	import CreateProjectDialog from './_components/_dialogs/create-project-dialog.svelte';
	import DeleteProjectDialog from './_components/_dialogs/delete-project-dialog.svelte';
	import EditProjectDialog from './_components/_dialogs/edit-project-dialog.svelte';
	import ProjectsTable from './_components/projects-table.svelte';

	let filterActive = $state(true);
	let filterArchived = $state(false);
	let filterDraft = $state(false);

	const projectStore = getProjects();

	// Dialog handlers
	const createDialog = createDialogController();
	const editDialog = createDialogController<{ id: string; project: ProjectUpdateSchema }>();
	const deleteDialog = createDialogController<ProjectDeleteSchema>();

	// Action handlers
	const handleCreateProject = async (data: ProjectCreateSchema) => {
		await projectStore.add(data);
	};

	const handleUpdateProject = async (id: string, data: ProjectUpdateSchema) => {
		await projectStore.edit(id, data);
	};

	const handleDeleteProject = async (data: ProjectDeleteSchema) => {
		await projectStore.remove(data);
	};
</script>

<UserPageLayout title="Projects">
	<div class="flex items-center">
		<Breadcrumb.Root class="hidden md:flex">
			<Breadcrumb.List>
				<Breadcrumb.Item>
					<Breadcrumb.Link href={AppRoute.DASHBOARD}>{m.dashboard()}</Breadcrumb.Link>
				</Breadcrumb.Item>
				<Breadcrumb.Separator />
				<Breadcrumb.Item>
					<Breadcrumb.Link href={AppRoute.PROJECTS}>{m.projects()}</Breadcrumb.Link>
				</Breadcrumb.Item>
				<Breadcrumb.Separator />
				<Breadcrumb.Item>
					<Breadcrumb.Page>{m.solid_heroic_poodle_dash()}</Breadcrumb.Page>
				</Breadcrumb.Item>
			</Breadcrumb.List>
		</Breadcrumb.Root>

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
					<DropdownMenu.CheckboxItem bind:checked={filterDraft}>
						{m.green_white_eagle_scold()}
					</DropdownMenu.CheckboxItem>
					<DropdownMenu.CheckboxItem bind:checked={filterArchived}>
						{m.petty_trick_ladybug_reap()}
					</DropdownMenu.CheckboxItem>
				</DropdownMenu.Content>
			</DropdownMenu.Root>
			<Button size="sm" variant="outline" class="h-8 gap-1">
				<File class="h-3.5 w-3.5" />
				<span class="sr-only sm:not-sr-only sm:whitespace-nowrap"
					>{m.slow_great_gadfly_expand()}</span
				>
			</Button>
			<Button size="sm" class="h-8 gap-1" onclick={() => createDialog.open()}>
				<CirclePlus class="h-3.5 w-3.5" />
				<span class="sr-only sm:not-sr-only sm:whitespace-nowrap">{m.stout_elegant_jan_flip()}</span
				>
			</Button>
		</div>
	</div>

	<Card.Root>
		<Card.Header>
			<Card.Title>{m.projects()}</Card.Title>
			<Card.Description>{m.only_nimble_martin_strive()}</Card.Description>
		</Card.Header>
		<Card.Content>
			<ProjectsTable
				{projectStore}
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
			<div class="text-muted-foreground text-xs">
				{m.weird_sharp_javelina_sail({
					amount: projectStore.projects.length,
					start: 1,
					end: projectStore.projects.length
				})}
			</div>
		</Card.Footer>
	</Card.Root>

	<CreateProjectDialog dialogController={createDialog} {handleCreateProject} />
	<EditProjectDialog dialogController={editDialog} {handleUpdateProject} />
	<DeleteProjectDialog dialogController={deleteDialog} {handleDeleteProject} />
</UserPageLayout>
