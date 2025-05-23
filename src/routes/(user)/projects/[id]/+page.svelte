<script lang="ts">
	import type { VersionSchema } from '$lib/schemas/version.schema';

	import { createQuery } from '@tanstack/svelte-query';
	import { page } from '$app/state';
	import UserPageLayout from '$lib/components/layouts/user-page-layout.svelte';
	import { Badge } from '$lib/components/ui/badge';
	import * as Card from '$lib/components/ui/card';
	import { createProjectQuery } from '$lib/queries/project';
	import { createPaginatedVersionsQuery } from '$lib/queries/versions';
	import { getPocketBase } from '$lib/services/pocketbase';

	const projectQuery = $derived.by(() =>
		createQuery(
			createProjectQuery({
				id: page.params.id
			})
		)
	);

	const versionsPagination = $state({
		page: 1,
		perPage: 5
	});

	const versionsQuery = $derived.by(() =>
		createQuery(
			createPaginatedVersionsQuery({
				project: $projectQuery.data?.id,
				page: versionsPagination.page,
				perPage: versionsPagination.perPage
			})
		)
	);

	let downloadElement: HTMLAnchorElement;

	const downloadFile = async (record: VersionSchema, file: string) => {
		const pb = getPocketBase();
		const token = await pb.files.getToken();
		const url = pb.files.getURL(record, file, { token });

		downloadElement.href = url;
		downloadElement.download = file;
		downloadElement.click();
	};
</script>

<!-- svelte-ignore a11y_consider_explicit_label -->
<!-- svelte-ignore a11y_missing_attribute -->
<a bind:this={downloadElement} target="_blank"></a>

<UserPageLayout title={$projectQuery.data?.name ?? 'Unknown Project'}>
	<Card.Root>
		<Card.Header>
			<Card.Title><Badge class="mr-2">Project</Badge>{$projectQuery.data?.name}</Card.Title>
		</Card.Header>
		<Card.Content>
			<ul>
				{#each $versionsQuery.data?.items ?? [] as version}
					<li>
						<p>Name: {version.name}</p>
						<p>Description: {version.description}</p>
						<p>Project: {version.project}</p>
						<p>Files:</p>
						<ul class="ml-4">
							{#each version.files as file}
								<li>
									<p>File name: {file}</p>
									<button onclick={() => downloadFile(version, file)}>Download</button>
								</li>
							{/each}
						</ul>
					</li>
				{/each}
			</ul>
		</Card.Content>
		<Card.Footer></Card.Footer>
	</Card.Root>
</UserPageLayout>
