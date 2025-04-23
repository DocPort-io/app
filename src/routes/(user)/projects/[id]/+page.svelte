<script lang="ts">
	import { CirclePlus, File, ListFilter } from '@lucide/svelte';
	import { createQuery } from '@tanstack/svelte-query';
	import { page } from '$app/state';
	import UserPageLayout from '$lib/components/layouts/user-page-layout.svelte';
	import * as Breadcrumb from '$lib/components/ui/breadcrumb';
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { AppRoute } from '$lib/constants';
	import { m } from '$lib/paraglide/messages';
	import { createProjectQuery } from '$lib/queries/project';

	let filterActive = $state(true);
	let filterArchived = $state(false);
	let filterDraft = $state(false);

	const projectQuery = $derived.by(() =>
		createQuery(
			createProjectQuery({
				id: page.params.id
			})
		)
	);
</script>

<UserPageLayout title="Projects">
	<div class="flex items-center">
		<Breadcrumb.Root class="hidden md:flex">
			<Breadcrumb.List>
				<Breadcrumb.Item>
					<Breadcrumb.Link href={AppRoute.DASHBOARD()}>{m.dashboard()}</Breadcrumb.Link>
				</Breadcrumb.Item>
				<Breadcrumb.Separator />
				<Breadcrumb.Item>
					<Breadcrumb.Link href={AppRoute.PROJECTS()}>{m.projects()}</Breadcrumb.Link>
				</Breadcrumb.Item>
				<Breadcrumb.Separator />
				<Breadcrumb.Item>
					<Breadcrumb.Page>{$projectQuery.data?.name}</Breadcrumb.Page>
				</Breadcrumb.Item>
			</Breadcrumb.List>
		</Breadcrumb.Root>

		<div class="ml-auto flex items-center gap-2">
			<DropdownMenu.Root>
				<DropdownMenu.Trigger asChild let:builder>
					<Button builders={[builder]} variant="outline" size="sm" class="h-8 gap-1">
						<ListFilter class="h-3.5 w-3.5" />
						<span class="sr-only sm:not-sr-only sm:whitespace-nowrap"
							>{m.that_weary_anteater_push()}</span
						>
					</Button>
				</DropdownMenu.Trigger>
				<DropdownMenu.Content align="end">
					<DropdownMenu.Label>{m.vexed_steep_piranha_kick()}</DropdownMenu.Label>
					<DropdownMenu.Separator />
					<DropdownMenu.CheckboxItem bind:checked={filterActive}>
						{m.alive_ok_kangaroo_boil()}
					</DropdownMenu.CheckboxItem>
					<DropdownMenu.CheckboxItem bind:checked={filterDraft}>
						{m.green_white_eagle_scold()}
					</DropdownMenu.CheckboxItem>
					<DropdownMenu.CheckboxItem bind:checked={filterArchived}>
						{m.petty_trick_ladybug_reap()}
					</DropdownMenu.CheckboxItem>
				</DropdownMenu.Content>
			</DropdownMenu.Root>
			<Button size="sm" variant="outline" class="h-8 gap-1">
				<File class="h-3.5 w-3.5" />
				<span class="sr-only sm:not-sr-only sm:whitespace-nowrap"
					>{m.slow_great_gadfly_expand()}</span
				>
			</Button>
			<Button size="sm" class="h-8 gap-1" on:click={() => {}}>
				<CirclePlus class="h-3.5 w-3.5" />
				<span class="sr-only sm:not-sr-only sm:whitespace-nowrap">{m.stout_elegant_jan_flip()}</span
				>
			</Button>
		</div>
	</div>

	<Card.Root>
		<Card.Header>
			<Card.Title>{m.projects()}</Card.Title>
			<Card.Description>{m.only_nimble_martin_strive()}</Card.Description>
		</Card.Header>
		<Card.Content></Card.Content>
		<Card.Footer></Card.Footer>
	</Card.Root>
</UserPageLayout>
