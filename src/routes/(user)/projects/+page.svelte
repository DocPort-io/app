<script lang="ts">
	import UserPageLayout from '$lib/components/layouts/user-page-layout.svelte';
	import CreateProjectDialog from './_components/_dialogs/create-project-dialog.svelte';
	import type { PageProps } from './$types';
	import Ellipsis from 'lucide-svelte/icons/ellipsis';
	import CirclePlus from 'lucide-svelte/icons/circle-plus';
	import File from 'lucide-svelte/icons/file';
	import ListFilter from 'lucide-svelte/icons/list-filter';
	import { Badge } from '$lib/components/ui/badge';
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import * as Table from '$lib/components/ui/table';
	import * as Breadcrumb from '$lib/components/ui/breadcrumb';
	import { AppRoute } from '$lib/constants';
	import { getProjects } from '$lib/states/projects.svelte';
	import { m } from '$lib/paraglide/messages';
	import DeleteProjectDialog from './_components/_dialogs/delete-project-dialog.svelte';
	import type { ProjectDeleteSchema } from '$lib/schemas/project.schema';
	import { getLocale } from '$lib/paraglide/runtime';

	let { data }: PageProps = $props();

	let filterActive = $state(true);
	let filterArchived = $state(false);
	let filterDraft = $state(false);

	let createProjectDialogOpen = $state(false);
	let deleteProjectDialogOpen = $state(false);
	let deleteProject = $state<ProjectDeleteSchema>({ id: '' });

	const projects = getProjects();

	$effect(() => {
		projects.load();
	});
</script>

<UserPageLayout>
	<div class="flex items-center">
		<Breadcrumb.Root class="hidden md:flex">
			<Breadcrumb.List>
				<Breadcrumb.Item>
					<Breadcrumb.Link href={AppRoute.DASHBOARD}>{m.dashboard()}</Breadcrumb.Link>
				</Breadcrumb.Item>
				<Breadcrumb.Separator />
				<Breadcrumb.Item>
					<Breadcrumb.Link href={AppRoute.PROJECTS}>{m.projects()}</Breadcrumb.Link>
				</Breadcrumb.Item>
				<Breadcrumb.Separator />
				<Breadcrumb.Item>
					<Breadcrumb.Page>{m.solid_heroic_poodle_dash()}</Breadcrumb.Page>
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
			<Button size="sm" class="h-8 gap-1" onclick={() => (createProjectDialogOpen = true)}>
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
		<Card.Content>
			<Table.Root>
				<Table.Header>
					<Table.Row>
						<Table.Head class="hidden w-[100px] sm:table-cell">
							<span class="sr-only">Image</span>
						</Table.Head>
						<Table.Head>{m.weak_few_ant_link()}</Table.Head>
						<!-- <Table.Head>Status</Table.Head> -->
						<!-- <Table.Head class="hidden md:table-cell">Price</Table.Head> -->
						<!-- <Table.Head class="hidden md:table-cell">Total Sales</Table.Head> -->
						<Table.Head class="hidden md:table-cell">Created at</Table.Head>
						<Table.Head>
							<span class="sr-only">Actions</span>
						</Table.Head>
					</Table.Row>
				</Table.Header>
				<Table.Body>
					{#await projects.loadingPromise}
						<p>Currently loading...</p>
					{:then}
						{#each projects.projects as project (project.id)}
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
									})}</Table.Cell
								>
								<Table.Cell>
									<DropdownMenu.Root>
										<DropdownMenu.Trigger asChild let:builder>
											<Button aria-haspopup="true" size="icon" variant="ghost" builders={[builder]}>
												<Ellipsis class="h-4 w-4" />
												<span class="sr-only">Toggle menu</span>
											</Button>
										</DropdownMenu.Trigger>
										<DropdownMenu.Content align="end">
											<DropdownMenu.Label>Actions</DropdownMenu.Label>
											<DropdownMenu.Item>Edit</DropdownMenu.Item>
											<DropdownMenu.Item
												onclick={() => {
													deleteProject = project;
													deleteProjectDialogOpen = true;
												}}
											>
												{m.fuzzy_lofty_stork_jest()}
											</DropdownMenu.Item>
										</DropdownMenu.Content>
									</DropdownMenu.Root>
								</Table.Cell>
							</Table.Row>
						{/each}
					{:catch err}
						<p>Failed to load projects: {err instanceof Error ? err.message : err}</p>
					{/await}
				</Table.Body>
			</Table.Root>
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
	<CreateProjectDialog bind:open={createProjectDialogOpen} />
	<DeleteProjectDialog bind:open={deleteProjectDialogOpen} project={deleteProject} />
</UserPageLayout>
