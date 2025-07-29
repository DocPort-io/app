<script lang="ts">
	import type { VersionSchema } from '$lib/schemas/version.schema';

	import { Calendar, Edit, FolderOpen, User } from '@lucide/svelte';
	import { createQuery } from '@tanstack/svelte-query';
	import { Badge } from '$lib/components/ui/badge';
	import { Button } from '$lib/components/ui/button';
	import { m } from '$lib/paraglide/messages';
	import { getLocale } from '$lib/paraglide/runtime';
	import { createUserQuery } from '$lib/queries/user';
	import { cn } from '$lib/utils';

	type Props = {
		version: VersionSchema;
		latest?: boolean;
		selected?: boolean;
		selectVersion: (version: VersionSchema) => void;
		onEdit?: (version: VersionSchema) => void;
	};

	let { version, latest, selected, selectVersion, onEdit }: Props = $props();

	const createdByQuery = $derived.by(() =>
		createQuery(
			createUserQuery({
				id: version.createdBy ?? ''
			})
		)
	);
</script>

<div
	class={cn('rounded-lg p-4 transition-all', {
		'border-primary bg-primary/5 border-2': selected,
		'hover:bg-accent border': !selected
	})}
>
	<div class="flex flex-col gap-3 md:flex-row md:items-center md:justify-between">
		<div class="space-y-1">
			<div class="flex items-center gap-2">
				<h4 class="text-sm font-semibold">{version.name}</h4>
				{#if latest}
					<Badge variant="outline" class="text-xs">{m.latest()}</Badge>
				{/if}
				{#if selected}
					<Badge variant="default" class="text-xs">{m.selected()}</Badge>
				{/if}
			</div>
			<p class="text-muted-foreground text-sm">
				{version.description || m.no_description_provided()}
			</p>
			<div class="text-muted-foreground flex items-center gap-4 text-xs">
				<span class="flex items-center gap-1">
					<User class="h-3 w-3" />
					{$createdByQuery.data?.name || m.unknown_user()}
				</span>
				<span class="flex items-center gap-1">
					<Calendar class="h-3 w-3" />
					{version.created
						? new Date(version.created).toLocaleDateString(getLocale())
						: m.unknown_date()}
				</span>
			</div>
		</div>
		<div class="flex gap-2">
			{#if onEdit}
				<Button variant="outline" size="sm" class="gap-2" onclick={() => onEdit(version)}>
					<Edit class="h-4 w-4" />
					{m.edit()}
				</Button>
			{/if}
			<Button variant="outline" size="sm" class="gap-2" onclick={() => selectVersion(version)}>
				<FolderOpen class="h-4 w-4" />
				{m.open_version()}
			</Button>
		</div>
	</div>
</div>
