<script lang="ts">
	import { createQuery } from '@tanstack/svelte-query';
	import { page } from '$app/state';
	import UserPageLayout from '$lib/components/layouts/user-page-layout.svelte';
	import * as Card from '$lib/components/ui/card';
	import { m } from '$lib/paraglide/messages';
	import { createProjectQuery } from '$lib/queries/project';

	const projectQuery = $derived.by(() =>
		createQuery(
			createProjectQuery({
				id: page.params.id
			})
		)
	);
</script>

<UserPageLayout title={$projectQuery.data?.name ?? 'Unknown Project'}>
	<Card.Root>
		<Card.Header>
			<Card.Title>{m.projects()}</Card.Title>
			<Card.Description>{m.manage_your_projects()}</Card.Description>
		</Card.Header>
		<Card.Content></Card.Content>
		<Card.Footer></Card.Footer>
	</Card.Root>
</UserPageLayout>
