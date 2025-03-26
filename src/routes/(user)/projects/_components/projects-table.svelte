<script lang="ts">
	import type { ProjectDeleteSchema, ProjectUpdateSchema } from '$lib/schemas/project.schema';

	import * as Table from '$lib/components/ui/table';
	import { m } from '$lib/paraglide/messages';
	import { Projects } from '$lib/stores/projects.svelte';

	import ProjectTableRowSkeleton from './project-table-row-skeleton.svelte';
	import ProjectTableRow from './project-table-row.svelte';

	type Props = {
		projectStore: Projects;
		handleEditProject: (id: string, project: ProjectUpdateSchema) => void;
		handleDeleteProject: (project: ProjectDeleteSchema) => void;
	};

	let { projectStore, handleEditProject, handleDeleteProject }: Props = $props();
</script>

<Table.Root>
	<Table.Header>
		<Table.Row>
			<Table.Head>{m.weak_few_ant_link()}</Table.Head>
			<Table.Head class="hidden md:table-cell">Created at</Table.Head>
			<Table.Head>
				<span class="sr-only">Actions</span>
			</Table.Head>
		</Table.Row>
	</Table.Header>
	<Table.Body>
		{#await projectStore.loadingPromise}
			<ProjectTableRowSkeleton />
			<ProjectTableRowSkeleton />
			<ProjectTableRowSkeleton />
		{:then}
			{#each projectStore.projects as project (project.id)}
				<ProjectTableRow {project} {handleEditProject} {handleDeleteProject} />
			{/each}
		{/await}
	</Table.Body>
</Table.Root>
