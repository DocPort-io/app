<script lang="ts">
	import type { VersionSchema } from '$lib/schemas/version.schema';

	import { Calendar, Clock } from '@lucide/svelte';
	import { createQuery } from '@tanstack/svelte-query';
	import { page } from '$app/state';
	import UserPageLayout from '$lib/components/layouts/user-page-layout.svelte';
	import { Badge } from '$lib/components/ui/badge';
	import { Tabs, TabsList, TabsTrigger } from '$lib/components/ui/tabs';
	import { m } from '$lib/paraglide/messages';
	import { getLocale } from '$lib/paraglide/runtime';
	import { createProjectQuery } from '$lib/queries/project';
	import { createPaginatedVersionsQuery } from '$lib/queries/versions';
	import { createDialogController } from '$lib/stores/dialog.svelte';

	import CurrentVersionTab from './(components)/current-version-tab.svelte';
	import EditVersionDialog from './(components)/edit-version-dialog.svelte';
	import VersionsHistoryTab from './(components)/versions-history-tab.svelte';

	const projectQuery = $derived.by(() =>
		createQuery(
			createProjectQuery({
				id: page.params.id
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

	const statusMap = {
		planned: m.planned(),
		active: m.active(),
		completed: m.completed()
	};

	const latestVersion = $derived.by(() => {
		return $latestVersionQuery.data?.items.at(0) ?? null;
	});

	let selectedVersion: VersionSchema | null = $state(null);

	const currentVersion = $derived.by(() => {
		return selectedVersion ?? latestVersion;
	});

	let selectedTab: 'current' | 'history' = $state('current');

	const editVersionDialogController = createDialogController<{ version: VersionSchema }>();

	const selectVersion = (version: VersionSchema) => {
		selectedVersion = version;
		selectedTab = 'current';
	};

	const handleEditVersion = (version: VersionSchema) => {
		editVersionDialogController.data = { version };
		editVersionDialogController.open();
	};
</script>

<UserPageLayout title={$projectQuery.data?.name ?? m.unknown_project()}>
	<div>
		<div class="mb-8">
			<div class="mb-4 flex flex-col gap-2 md:flex-row md:items-center md:justify-between">
				<p class="text-3xl font-bold tracking-tight">{$projectQuery.data?.name}</p>
				{#if $projectQuery.data?.status}
					<Badge variant="outline">
						{statusMap[$projectQuery.data.status]}
					</Badge>
				{/if}
			</div>
			<div class="text-muted-foreground flex gap-4 text-sm md:items-center md:gap-6">
				<div class="flex items-center gap-2">
					<Calendar class="h-4 w-4" />
					{#if $projectQuery.data?.created}
						<span
							>{m.created()}
							{new Date($projectQuery.data?.created).toLocaleDateString(getLocale())}</span
						>
					{:else}
						<span>{m.created_unknown()}</span>
					{/if}
				</div>
				<div class="flex items-center gap-2">
					<Clock class="h-4 w-4" />
					<span>{$latestVersionQuery.data?.totalItems ?? 0} {m.versions()}</span>
				</div>
			</div>
		</div>
		<Tabs class="space-y-2" bind:value={selectedTab}>
			<TabsList>
				<TabsTrigger value="current">{m.current_version()}</TabsTrigger>
				<TabsTrigger value="history">{m.versions_history()}</TabsTrigger>
			</TabsList>

			<CurrentVersionTab {currentVersion} onEditVersion={handleEditVersion} />
			<VersionsHistoryTab {selectVersion} {currentVersion} />
		</Tabs>
	</div>
</UserPageLayout>

<EditVersionDialog dialogController={editVersionDialogController} />
