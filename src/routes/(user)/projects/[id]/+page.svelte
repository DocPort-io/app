<script lang="ts">
	import type { FileSchema } from '$lib/schemas/file.schema';
	import type { VersionSchema } from '$lib/schemas/version.schema';

	import {
		Archive,
		Calendar,
		Clock,
		Download,
		File,
		FileText,
		FolderOpen,
		Image,
		User
	} from '@lucide/svelte';
	import { createQuery } from '@tanstack/svelte-query';
	import { page } from '$app/state';
	import UserPageLayout from '$lib/components/layouts/user-page-layout.svelte';
	import { Badge } from '$lib/components/ui/badge';
	import { Button } from '$lib/components/ui/button';
	import {
		Card,
		CardTitle,
		CardContent,
		CardHeader,
		CardDescription
	} from '$lib/components/ui/card';
	import { ScrollArea } from '$lib/components/ui/scroll-area';
	import { Separator } from '$lib/components/ui/separator';
	import { Tabs, TabsContent, TabsList, TabsTrigger } from '$lib/components/ui/tabs';
	import {
		resolveTextMapping,
		type MimeTypeIconMapping,
		resolveIconMapping
	} from '$lib/file-mappings';
	import { getLocale } from '$lib/paraglide/runtime';
	import { createPaginatedFilesQuery } from '$lib/queries/files';
	import { createProjectQuery } from '$lib/queries/project';
	import { createPaginatedVersionsQuery } from '$lib/queries/versions';
	import { getPocketBase } from '$lib/services/pocketbase';
	import prettyBytes from 'pretty-bytes';

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

	let downloadElement: HTMLAnchorElement;

	const downloadFile = async (record: FileSchema) => {
		const pb = getPocketBase();
		const token = await pb.files.getToken();
		const url = pb.files.getURL(record, record.file, { token });

		downloadElement.href = url;
		downloadElement.download = record.name;
		downloadElement.click();
	};

	const statusMap = {
		planned: 'Planned',
		active: 'Active',
		completed: 'Completed'
	};

	const latestVersion = $derived.by(() => {
		return $versionsQuery.data?.items.at(0) ?? null;
	});

	let selectedVersion: VersionSchema | null = $state(null);

	const currentVersion = $derived.by(() => {
		return selectedVersion ?? latestVersion;
	});

	const otherVersions = $derived.by(() => {
		return $versionsQuery.data?.items.slice(1) ?? [];
	});

	const filesPagination = $state({
		page: 1,
		perPage: 25
	});

	const filesQuery = $derived.by(() =>
		createQuery(
			createPaginatedFilesQuery({
				version: currentVersion?.id,
				page: filesPagination.page,
				perPage: filesPagination.perPage
			})
		)
	);

	let selectedTab: 'current' | 'history' = $state('current');

	const resetFilesPagination = () => {
		filesPagination.page = 1;
	};

	const selectVersion = (version: VersionSchema) => {
		selectedVersion = version;
		resetFilesPagination();
		selectedTab = 'current';
	};
</script>

<!-- svelte-ignore a11y_consider_explicit_label -->
<!-- svelte-ignore a11y_missing_attribute -->
<a bind:this={downloadElement} target="_blank"></a>

{#snippet getFileIcon(type: MimeTypeIconMapping)}
	{#if type === 'document'}
		<FileText class="h-4 w-4" />
	{:else if type === 'image'}
		<Image class="h-4 w-4" />
	{:else if type === 'archive'}
		<Archive class="h-4 w-4" />
	{:else}
		<File class="h-4 w-4" />
	{/if}
{/snippet}

<UserPageLayout title={$projectQuery.data?.name ?? 'Unknown Project'}>
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
						<span>Created {new Date($projectQuery.data?.created).toLocaleDateString()}</span>
					{:else}
						<span>Created: Unknown</span>
					{/if}
				</div>
				<div class="flex items-center gap-2">
					<Clock class="h-4 w-4" />
					<span>{$versionsQuery.data?.totalItems ?? 0} version(s)</span>
				</div>
			</div>
		</div>
		<Tabs class="space-y-2" bind:value={selectedTab}>
			<TabsList>
				<TabsTrigger value="current">Current Version</TabsTrigger>
				<TabsTrigger value="history">Versions History</TabsTrigger>
			</TabsList>

			<TabsContent value="current" class="space-y-4">
				<Card>
					<CardHeader>
						<div class="flex items-center justify-between">
							<div>
								<CardTitle>{currentVersion?.name}</CardTitle>
								<CardDescription class="mt-2">{currentVersion?.description}</CardDescription>
							</div>
							<Badge variant="outline" class="flex items-center gap-1">
								<Clock class="h-3 w-3" />
								{currentVersion
									? new Date(currentVersion?.created).toLocaleDateString(getLocale())
									: null}
							</Badge>
						</div>
					</CardHeader>
					<CardContent>
						<div class="space-y-4">
							<div>
								<h3 class="mb-3 text-sm font-medium">Files ({$filesQuery.data?.items.length})</h3>
								<div class="space-y-2">
									{#each $filesQuery.data?.items ?? [] as file}
										<div
											class="bg-card hover:bg-accent flex flex-col gap-3 rounded-lg border p-3 transition-colors md:flex-row md:items-center md:justify-between"
										>
											<div class="flex flex-col gap-3 md:flex-row md:items-center">
												<div class="bg-muted w-min rounded-md p-2">
													{@render getFileIcon(resolveIconMapping(file.type))}
												</div>
												<div>
													<p class="text-sm font-medium break-all">{file.name}</p>
													<p class="text-muted-foreground text-xs">
														{prettyBytes(file.size, { locale: getLocale() })} - {resolveTextMapping(
															file.type
														)}
													</p>
												</div>
											</div>
											<Button
												variant="outline"
												size="sm"
												class="gap-2"
												onclick={() => downloadFile(file)}
											>
												<Download class="h-4 w-4" />
												Download
											</Button>
										</div>
									{/each}
								</div>
							</div>
						</div>
					</CardContent>
				</Card>
			</TabsContent>

			<TabsContent value="history" class="space-y-4">
				<Card>
					<CardHeader>
						<CardTitle>Version History</CardTitle>
						<CardDescription>Browse through previous versions of this project</CardDescription>
					</CardHeader>
					<CardContent>
						<ScrollArea class="h-[400px] pr-4">
							<div class="space-y-4">
								<div class="border-primary bg-primary/5 rounded-lg border-2 p-4 transition-all">
									<div class="flex flex-col gap-3 md:flex-row md:items-center md:justify-between">
										<div class="space-y-1">
											<div class="flex items-center gap-2">
												<h4 class="text-sm font-semibold">{latestVersion?.name}</h4>
												<Badge variant="default" class="text-xs">Latest</Badge>
											</div>
											<p class="text-muted-foreground text-sm">{latestVersion?.description}</p>
											<div class="text-muted-foreground flex items-center gap-4 text-xs">
												<span class="flex items-center gap-1">
													<User class="h-3 w-3" />
													Jonas Claes
												</span>
												<span class="flex items-center gap-1">
													<Calendar class="h-3 w-3" />
													{latestVersion?.created
														? new Date(latestVersion.created).toLocaleDateString()
														: null}
												</span>
											</div>
										</div>
										<Button
											variant="outline"
											size="sm"
											class="gap-2"
											onclick={() => (latestVersion ? selectVersion(latestVersion) : undefined)}
										>
											<FolderOpen class="h-4 w-4" />
											Open Version
										</Button>
									</div>
								</div>

								<Separator />

								{#each otherVersions as version}
									<div class="hover:bg-accent rounded-lg border p-4 transition-all">
										<div class="flex flex-col gap-3 md:flex-row md:items-center md:justify-between">
											<div class="space-y-1">
												<h4 class="text-sm font-semibold">{version.name}</h4>
												<p class="text-muted-foreground text-sm">
													{version.description || 'No description provided.'}
												</p>
												<div class="text-muted-foreground flex items-center gap-4 text-xs">
													<span class="flex items-center gap-1">
														<User class="h-3 w-3" />
														Jonas Claes
													</span>
													<span class="flex items-center gap-1">
														<Calendar class="h-3 w-3" />
														{new Date(version.created).toLocaleDateString()}
													</span>
												</div>
											</div>
											<Button
												variant="outline"
												size="sm"
												class="gap-2"
												onclick={() => selectVersion(version)}
											>
												<FolderOpen class="h-4 w-4" />
												Open Version
											</Button>
										</div>
									</div>
								{/each}
							</div>
						</ScrollArea>
					</CardContent>
				</Card>
			</TabsContent>
		</Tabs>
	</div>
</UserPageLayout>
