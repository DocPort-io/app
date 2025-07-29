<script lang="ts">
	import type { VersionSchema } from '$lib/schemas/version.schema';

	import { createQuery } from '@tanstack/svelte-query';
	import { page } from '$app/state';
	import {
		Card,
		CardTitle,
		CardContent,
		CardHeader,
		CardDescription
	} from '$lib/components/ui/card';
	import {
		Pagination,
		PaginationContent,
		PaginationEllipsis,
		PaginationItem,
		PaginationLink,
		PaginationNextButton,
		PaginationPrevButton
	} from '$lib/components/ui/pagination';
	import { Separator } from '$lib/components/ui/separator';
	import { TabsContent } from '$lib/components/ui/tabs';
	import { m } from '$lib/paraglide/messages';
	import { createProjectQuery } from '$lib/queries/project';
	import { createPaginatedVersionsQuery } from '$lib/queries/versions';

	import Version from './version.svelte';

	type Props = {
		currentVersion?: VersionSchema | null;
		selectVersion: (version: VersionSchema) => void;
	};

	let { currentVersion, selectVersion }: Props = $props();

	const projectQuery = $derived.by(() =>
		createQuery(
			createProjectQuery({
				id: page.params.id
			})
		)
	);

	const versionsPagination = $state({
		page: 1,
		perPage: 25
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

	const latestVersionQuery = $derived.by(() =>
		createQuery(
			createPaginatedVersionsQuery({
				project: $projectQuery.data?.id,
				page: 1,
				perPage: 1
			})
		)
	);

	const latestVersion = $derived.by(() => {
		return $latestVersionQuery.data?.items.at(0) ?? null;
	});

	const otherVersions = $derived.by(() => {
		return $versionsQuery.data?.items ?? [];
	});
</script>

<TabsContent value="history" class="space-y-4">
	<Card>
		<CardHeader>
			<CardTitle>{m.version_history()}</CardTitle>
			<CardDescription>{m.browse_through_previous_versions()}</CardDescription>
		</CardHeader>
		<CardContent>
			<div class="space-y-4">
				{#if currentVersion}
					<Version version={currentVersion} selected={true} {selectVersion} />
				{:else}
					<p>{m.no_versions_available()}</p>
				{/if}

				{#if $versionsQuery.data?.totalItems ?? 0 > 0}
					<Separator />
					{#each otherVersions as version (version.id)}
						<Version {version} {selectVersion} latest={version.id === latestVersion?.id} />
					{/each}

					<Pagination
						count={$versionsQuery.data?.totalItems ?? 0}
						perPage={versionsPagination.perPage}
						bind:page={versionsPagination.page}
					>
						{#snippet children({ pages, currentPage })}
							<PaginationContent>
								<PaginationItem class="hidden md:block">
									<PaginationPrevButton />
								</PaginationItem>
								{#each pages as page (page.key)}
									{#if page.type === 'ellipsis'}
										<PaginationItem>
											<PaginationEllipsis />
										</PaginationItem>
									{:else}
										<PaginationItem>
											<PaginationLink {page} isActive={currentPage == page.value}>
												{page.value}
											</PaginationLink>
										</PaginationItem>
									{/if}
								{/each}
								<PaginationItem class="hidden md:block">
									<PaginationNextButton />
								</PaginationItem>
							</PaginationContent>
						{/snippet}
					</Pagination>
				{/if}
			</div>
		</CardContent>
	</Card>
</TabsContent>
