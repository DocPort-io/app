<script lang="ts">
	import type {
		ProjectDeleteSchema,
		ProjectSchema,
		ProjectUpdateSchema
	} from '$lib/schemas/project.schema';

	import { Button } from '$lib/components/ui/button';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import * as Table from '$lib/components/ui/table';
	import { m } from '$lib/paraglide/messages';
	import { getLocale } from '$lib/paraglide/runtime';
	import Ellipsis from 'lucide-svelte/icons/ellipsis';

	type Props = {
		project: ProjectSchema;
		handleEditProject: (id: string, project: ProjectUpdateSchema) => void;
		handleDeleteProject: (project: ProjectDeleteSchema) => void;
	};

	let { project, handleEditProject, handleDeleteProject }: Props = $props();
</script>

<Table.Row>
	<Table.Cell class="hidden sm:table-cell">
		<img
			alt="Product example"
			class="aspect-square rounded-md object-cover"
			height="64"
			src="/images/placeholder.svg"
			width="64"
		/>
	</Table.Cell>
	<Table.Cell class="font-medium">{project.name}</Table.Cell>
	<!-- <Table.Cell>
        <Badge variant="outline">Draft</Badge>
        {#if project._isOptimistic}
            <Badge variant="outline">Saving...</Badge>
        {/if}
    </Table.Cell> -->
	<!-- <Table.Cell class="hidden md:table-cell">$499.99</Table.Cell> -->
	<!-- <Table.Cell class="hidden md:table-cell">25</Table.Cell> -->
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
				<DropdownMenu.Item onclick={() => handleEditProject(project.id, project)}>
					{m.lucky_factual_marmot_scoop()}
				</DropdownMenu.Item>
				<DropdownMenu.Item onclick={() => handleDeleteProject(project)}>
					{m.fuzzy_lofty_stork_jest()}
				</DropdownMenu.Item>
			</DropdownMenu.Content>
		</DropdownMenu.Root>
	</Table.Cell>
</Table.Row>
