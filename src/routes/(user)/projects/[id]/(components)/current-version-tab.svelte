<script lang="ts">
	import type { FileSchema } from '$lib/schemas/file.schema';
	import type { VersionSchema } from '$lib/schemas/version.schema';

	import { Archive, Clock, Download, File, FileText, Image, Upload } from '@lucide/svelte';
	import { createQuery } from '@tanstack/svelte-query';
	import { Badge } from '$lib/components/ui/badge';
	import { Button } from '$lib/components/ui/button';
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
	import { TabsContent } from '$lib/components/ui/tabs';
	import {
		resolveTextMapping,
		type MimeTypeIconMapping,
		resolveIconMapping
	} from '$lib/file-mappings';
	import { m } from '$lib/paraglide/messages';
	import { getLocale } from '$lib/paraglide/runtime';
	import { createPaginatedFilesQuery } from '$lib/queries/files';
	import { getPocketBase } from '$lib/services/pocketbase';
	import { createDialogController } from '$lib/stores/dialog.svelte';
	import prettyBytes from 'pretty-bytes';

	import UploadFileDialog from './upload-file-dialog.svelte';

	type Props = {
		currentVersion?: VersionSchema | null;
	};

	let { currentVersion }: Props = $props();

	let downloadElement: HTMLAnchorElement;

	const uploadDialogController = createDialogController<{ versionId: string }>();

	const downloadFile = async (record: FileSchema) => {
		const pb = getPocketBase();
		const token = await pb.files.getToken();
		const url = pb.files.getURL(record, record.file, { token });

		downloadElement.href = url;
		downloadElement.download = record.name;
		downloadElement.click();
	};

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

	$effect(() => {
		if (currentVersion) {
			filesPagination.page = 1;
		}
	});
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
				<div class="flex items-center justify-between">
					<h3 class="text-sm font-medium">{m.files()} ({$filesQuery.data?.items.length})</h3>
					{#if currentVersion}
						<Button
							variant="outline"
							size="sm"
							class="gap-2"
							onclick={() => {
								uploadDialogController.data = { versionId: currentVersion.id };
								uploadDialogController.open();
							}}
						>
							<Upload class="h-4 w-4" />
							{m.upload()}
						</Button>
					{/if}
				</div>
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
							<Button variant="outline" size="sm" class="gap-2" onclick={() => downloadFile(file)}>
								<Download class="h-4 w-4" />
								{m.download()}
							</Button>
						</div>
					{/each}
				</div>
				<Pagination
					count={$filesQuery.data?.totalItems ?? 0}
					perPage={filesPagination.perPage}
					bind:page={filesPagination.page}
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
			</div>
		</CardContent>
	</Card>
</TabsContent>

<UploadFileDialog dialogController={uploadDialogController} />
