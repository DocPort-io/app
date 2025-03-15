<script lang="ts">
	import type { ProjectDeleteSchema, ProjectUpdateSchema } from '$lib/schemas/project.schema';

	import * as Table from '$lib/components/ui/table';
	import { m } from '$lib/paraglide/messages';
	import { Projects } from '$lib/stores/projects.svelte';

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
			<!-- <Table.Head class="hidden w-[100px] sm:table-cell">
				<span class="sr-only">Image</span>
			</Table.Head> -->
			<Table.Head>{m.weak_few_ant_link()}</Table.Head>
			<!-- <Table.Head>Status</Table.Head> -->
			<!-- <Table.Head class="hidden md:table-cell">Price</Table.Head> -->
			<!-- <Table.Head class="hidden md:table-cell">Total Sales</Table.Head> -->
			<Table.Head class="hidden md:table-cell">Created at</Table.Head>
			<Table.Head>
				<span class="sr-only">Actions</span>
			</Table.Head>
		</Table.Row>
	</Table.Header>
	<Table.Body>
		{#await projectStore.loadingPromise}
			<p>Currently loading...</p>
		{:then}
			{#each projectStore.projects as project (project.id)}
				<ProjectTableRow {project} {handleEditProject} {handleDeleteProject} />
			{/each}
		{:catch err}
			<p>Failed to load projects: {err instanceof Error ? err.message : err}</p>
		{/await}
	</Table.Body>
</Table.Root>
