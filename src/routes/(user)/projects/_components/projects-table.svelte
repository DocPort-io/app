<script lang="ts">
	import type {
		ProjectDeleteSchema,
		ProjectSchema,
		ProjectUpdateSchema
	} from '$lib/schemas/project.schema';

	import * as Table from '$lib/components/ui/table';
	import { m } from '$lib/paraglide/messages';

	import ProjectTableRowSkeleton from './project-table-row-skeleton.svelte';
	import ProjectTableRow from './project-table-row.svelte';

	type Props = {
		loading: boolean;
		error: string | null;
		projects: ProjectSchema[];
		handleViewProject: (project: ProjectSchema) => void;
		handleEditProject: (id: string, project: ProjectUpdateSchema) => void;
		handleDeleteProject: (project: ProjectDeleteSchema) => void;
	};

	let {
		loading,
		error,
		projects,
		handleViewProject,
		handleEditProject,
		handleDeleteProject
	}: Props = $props();
</script>

<Table.Root data-testid="projects-table">
	<Table.Header data-testid="projects-table-header">
		<Table.Row>
			<Table.Head>{m.weak_few_ant_link()}</Table.Head>
			<Table.Head class="hidden md:table-cell">Status</Table.Head>
			<Table.Head class="hidden md:table-cell">Created at</Table.Head>
			<Table.Head>
				<span class="sr-only">Actions</span>
			</Table.Head>
		</Table.Row>
	</Table.Header>
	<Table.Body data-testid="projects-table-body">
		{#if loading}
			<ProjectTableRowSkeleton />
			<ProjectTableRowSkeleton />
			<ProjectTableRowSkeleton />
		{:else if error}
			<Table.Row>
				<Table.Cell colspan={3} class="text-center">
					{error}
				</Table.Cell>
			</Table.Row>
		{:else}
			{#each projects as project (project.id)}
				<ProjectTableRow {project} {handleViewProject} {handleEditProject} {handleDeleteProject} />
			{/each}
		{/if}
	</Table.Body>
</Table.Root>
