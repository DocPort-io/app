<script lang="ts">
	import type {
		ProjectDeleteSchema,
		ProjectSchema,
		ProjectUpdateSchema
	} from '$lib/schemas/project.schema';

	import { Ellipsis } from '@lucide/svelte';
	import { Badge } from '$lib/components/ui/badge';
	import { Button } from '$lib/components/ui/button';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import * as Table from '$lib/components/ui/table';
	import { m } from '$lib/paraglide/messages';
	import { getLocale } from '$lib/paraglide/runtime';

	type Props = {
		project: ProjectSchema;
		handleViewProject: (project: ProjectSchema) => void;
		handleEditProject: (id: string, project: ProjectUpdateSchema) => void;
		handleDeleteProject: (project: ProjectDeleteSchema) => void;
	};

	let { project, handleViewProject, handleEditProject, handleDeleteProject }: Props = $props();

	const statusMap = {
		active: 'Active',
		completed: 'Completed'
	};
</script>

<Table.Row data-testid="projects-table-row">
	<Table.Cell class="font-medium">
		{project.name}
	</Table.Cell>
	<Table.Cell>
		<Badge>{statusMap[project.status]}</Badge>
	</Table.Cell>
	<Table.Cell class="hidden md:table-cell"
		>{new Date(project.created).toLocaleString(getLocale(), {
			dateStyle: 'long',
			timeStyle: 'short'
		})}
	</Table.Cell>
	<Table.Cell>
		<DropdownMenu.Root>
			<DropdownMenu.Trigger asChild let:builder>
				<Button aria-haspopup="true" size="icon" variant="ghost" builders={[builder]}>
					<Ellipsis class="h-4 w-4" />
					<span class="sr-only">Toggle menu</span>
				</Button>
			</DropdownMenu.Trigger>
			<DropdownMenu.Content align="end">
				<DropdownMenu.Label>{m.trite_gaudy_marten_scold()}</DropdownMenu.Label>
				<DropdownMenu.Item on:click={() => handleViewProject(project)}>
					{m.view()}
				</DropdownMenu.Item>
				<DropdownMenu.Item on:click={() => handleEditProject(project.id, project)}>
					{m.lucky_factual_marmot_scoop()}
				</DropdownMenu.Item>
				<DropdownMenu.Item on:click={() => handleDeleteProject(project)}>
					{m.fuzzy_lofty_stork_jest()}
				</DropdownMenu.Item>
			</DropdownMenu.Content>
		</DropdownMenu.Root>
	</Table.Cell>
</Table.Row>
