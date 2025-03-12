<script lang="ts">
	import * as Card from '$lib/components/ui/card';
	import { m } from '$lib/paraglide/messages';
	import { getProjects } from '$lib/stores/projects.svelte';

	import ProjectsTable from '../projects-table.svelte';

	const projects = getProjects();
</script>

<Card.Root>
	<Card.Header>
		<Card.Title>{m.projects()}</Card.Title>
		<Card.Description>{m.only_nimble_martin_strive()}</Card.Description>
	</Card.Header>
	<Card.Content>
		<ProjectsTable />
	</Card.Content>
	<Card.Footer>
		<div class="text-muted-foreground text-xs">
			{m.weird_sharp_javelina_sail({
				amount: projects.projects.length,
				start: 1,
				end: projects.projects.length
			})}
		</div>
	</Card.Footer>
</Card.Root>
