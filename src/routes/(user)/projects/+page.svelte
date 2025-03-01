<script lang="ts">
	import UserPageLayout from '$lib/components/layouts/user-page-layout.svelte';
	import type { PageData } from './$types';
	import Ellipsis from 'lucide-svelte/icons/ellipsis';
	import CirclePlus from 'lucide-svelte/icons/circle-plus';
	import File from 'lucide-svelte/icons/file';
	import ListFilter from 'lucide-svelte/icons/list-filter';
	import { Badge } from '$lib/components/ui/badge/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import * as Card from '$lib/components/ui/card/index.js';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js';
	import * as Table from '$lib/components/ui/table/index.js';
	import * as Breadcrumb from '$lib/components/ui/breadcrumb/index.js';
	import { AppRoute } from '$lib/constants';
	import * as m from '$lib/paraglide/messages.js';
	import PocketBase, { type RecordModel } from 'pocketbase';
	import { asyncWritable } from '@square/svelte-store';
	import { projectActions, projects } from '$lib/stores/projects';

	let { data }: { data: PageData } = $props();

	let filterActive = $state(true);
	let filterArchived = $state(false);
	let filterDraft = $state(false);

	// const pb = new PocketBase('http://127.0.0.1:8090');

	// const getProjects = async () => {
	// 	const response = await pb.collection('projects').getList();
	// 	console.log('Fetched projects', response.items);
	// 	return response.items;
	// };

	// function isEqual(project1: RecordModel, project2: RecordModel) {
	// 	// Deep comparison of two project objects
	// 	// This is a simple example; you might need to adjust based on your RecordModel properties
	// 	return JSON.stringify(project1) === JSON.stringify(project2);
	// }

	// const addProject = async (project: RecordModel) => {
	// 	const response = await pb.collection('projects').create(project);
	// 	console.log('Added project', response);
	// 	return response;
	// };

	// const updateProject = async (project: RecordModel) => {
	// 	const response = await pb.collection('projects').update(project.id, project);
	// 	console.log('Updated project', response);
	// 	return response;
	// };

	// const deleteProject = async (projectId: string) => {
	// 	await pb.collection('projects').delete(projectId);
	// 	console.log('Deleted project', projectId);
	// };

	// const updateProjects = async (
	// 	newProjects: RecordModel[],
	// 	_: unknown,
	// 	oldProjects: RecordModel[] = []
	// ) => {
	// 	// Create maps for faster lookups
	// 	const oldProjectsMap = new Map(oldProjects.map((project) => [project.id, project]));
	// 	const newProjectsMap = new Map(newProjects.map((project) => [project.id, project]));

	// 	// Find projects to add (in new but not in old)
	// 	const projectsToAdd = newProjects.filter((project) => !oldProjectsMap.has(project.id));

	// 	// Find projects to delete (in old but not in new)
	// 	const projectsToDelete = oldProjects.filter((project) => !newProjectsMap.has(project.id));

	// 	// Find projects to update (in both but might have changes)
	// 	const projectsToUpdate = newProjects.filter((project) => {
	// 		const oldProject = oldProjectsMap.get(project.id);
	// 		return oldProject && !isEqual(project, oldProject);
	// 	});

	// 	// Initialize the final result array with projects that remain unchanged
	// 	const finalProjects = newProjects.filter((project) => {
	// 		const oldProject = oldProjectsMap.get(project.id);
	// 		return oldProject && isEqual(project, oldProject);
	// 	});

	// 	// Process additions and collect updated objects with potential new IDs
	// 	for (const project of projectsToAdd) {
	// 		const addedProject = await addProject(project);
	// 		finalProjects.push(addedProject); // Use the response which might contain a new ID
	// 	}

	// 	// Process updates and collect updated objects
	// 	for (const project of projectsToUpdate) {
	// 		const updatedProject = await updateProject(project);
	// 		finalProjects.push(updatedProject); // Use the response which might contain updated data
	// 	}

	// 	// Process deletions (no need to add these to the final array)
	// 	for (const project of projectsToDelete) {
	// 		await deleteProject(project.id);
	// 	}

	// 	return finalProjects;
	// };

	// const addXProject = () => {
	// 	console.log('Add project');
	// 	projects.update((projects: RecordModel[]) => {
	// 		return [
	// 			...projects,
	// 			{
	// 				name: 'Project X'
	// 			}
	// 		] as RecordModel[];
	// 	});
	// };

	// const projects = asyncWritable([], getProjects, updateProjects);
</script>

<UserPageLayout>
	<!-- <h1>Test</h1> -->
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
				<span class="sr-only sm:not-sr-only sm:whitespace-nowrap">Export</span>
			</Button>
			<Button size="sm" class="h-8 gap-1" onclick={projectActions.createProject}>
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
						<Table.Head>Name</Table.Head>
						<Table.Head>Status</Table.Head>
						<Table.Head class="hidden md:table-cell">Price</Table.Head>
						<Table.Head class="hidden md:table-cell">Total Sales</Table.Head>
						<Table.Head class="hidden md:table-cell">Created at</Table.Head>
						<Table.Head>
							<span class="sr-only">Actions</span>
						</Table.Head>
					</Table.Row>
				</Table.Header>
				<Table.Body>
					{#await projects.load()}
						<p>Currently loading...</p>
					{:then}
						{#each $projects as project (project.id)}
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
								<Table.Cell>
									<Badge variant="outline">Draft</Badge>
									{#if project._isOptimistic}
										<Badge variant="outline">Saving...</Badge>
									{/if}
								</Table.Cell>
								<Table.Cell class="hidden md:table-cell">$499.99</Table.Cell>
								<Table.Cell class="hidden md:table-cell">25</Table.Cell>
								<Table.Cell class="hidden md:table-cell">2023-07-12 10:42 AM</Table.Cell>
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
											<DropdownMenu.Item>Delete</DropdownMenu.Item>
										</DropdownMenu.Content>
									</DropdownMenu.Root>
								</Table.Cell>
							</Table.Row>
						{/each}
					{:catch error}
						<p>Failed to load projects: {error.message}</p>
					{/await}
					<!-- <Table.Row>
						<Table.Cell class="hidden sm:table-cell">
							<img
								alt="Product example"
								class="aspect-square rounded-md object-cover"
								height="64"
								src="/images/placeholder.svg"
								width="64"
							/>
						</Table.Cell>
						<Table.Cell class="font-medium">Laser Lemonade Machine</Table.Cell>
						<Table.Cell>
							<Badge variant="outline">Draft</Badge>
						</Table.Cell>
						<Table.Cell class="hidden md:table-cell">$499.99</Table.Cell>
						<Table.Cell class="hidden md:table-cell">25</Table.Cell>
						<Table.Cell class="hidden md:table-cell">2023-07-12 10:42 AM</Table.Cell>
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
									<DropdownMenu.Item>Delete</DropdownMenu.Item>
								</DropdownMenu.Content>
							</DropdownMenu.Root>
						</Table.Cell>
					</Table.Row>
					<Table.Row>
						<Table.Cell class="hidden sm:table-cell">
							<img
								alt="Product"
								class="aspect-square rounded-md object-cover"
								height="64"
								src="/images/placeholder.svg"
								width="64"
							/>
						</Table.Cell>
						<Table.Cell class="font-medium">Hypernova Headphones</Table.Cell>
						<Table.Cell>
							<Badge variant="outline">Active</Badge>
						</Table.Cell>
						<Table.Cell class="hidden md:table-cell">$129.99</Table.Cell>
						<Table.Cell class="hidden md:table-cell">100</Table.Cell>
						<Table.Cell class="hidden md:table-cell">2023-10-18 03:21 PM</Table.Cell>
						<Table.Cell>
							<DropdownMenu.Root>
								<DropdownMenu.Trigger asChild let:builder>
									<Button builders={[builder]} aria-haspopup="true" size="icon" variant="ghost">
										<Ellipsis class="h-4 w-4" />
										<span class="sr-only">Toggle menu</span>
									</Button>
								</DropdownMenu.Trigger>
								<DropdownMenu.Content align="end">
									<DropdownMenu.Label>Actions</DropdownMenu.Label>
									<DropdownMenu.Item>Edit</DropdownMenu.Item>
									<DropdownMenu.Item>Delete</DropdownMenu.Item>
								</DropdownMenu.Content>
							</DropdownMenu.Root>
						</Table.Cell>
					</Table.Row>
					<Table.Row>
						<Table.Cell class="hidden sm:table-cell">
							<img
								alt="Product"
								class="aspect-square rounded-md object-cover"
								height="64"
								src="/images/placeholder.svg"
								width="64"
							/>
						</Table.Cell>
						<Table.Cell class="font-medium">AeroGlow Desk Lamp</Table.Cell>
						<Table.Cell>
							<Badge variant="outline">Active</Badge>
						</Table.Cell>
						<Table.Cell class="hidden md:table-cell">$39.99</Table.Cell>
						<Table.Cell class="hidden md:table-cell">50</Table.Cell>
						<Table.Cell class="hidden md:table-cell">2023-11-29 08:15 AM</Table.Cell>
						<Table.Cell>
							<DropdownMenu.Root>
								<DropdownMenu.Trigger asChild let:builder>
									<Button builders={[builder]} aria-haspopup="true" size="icon" variant="ghost">
										<Ellipsis class="h-4 w-4" />
										<span class="sr-only">Toggle menu</span>
									</Button>
								</DropdownMenu.Trigger>
								<DropdownMenu.Content align="end">
									<DropdownMenu.Label>Actions</DropdownMenu.Label>
									<DropdownMenu.Item>Edit</DropdownMenu.Item>
									<DropdownMenu.Item>Delete</DropdownMenu.Item>
								</DropdownMenu.Content>
							</DropdownMenu.Root>
						</Table.Cell>
					</Table.Row>
					<Table.Row>
						<Table.Cell class="hidden sm:table-cell">
							<img
								alt="Product"
								class="aspect-square rounded-md object-cover"
								height="64"
								src="/images/placeholder.svg"
								width="64"
							/>
						</Table.Cell>
						<Table.Cell class="font-medium">TechTonic Energy Drink</Table.Cell>
						<Table.Cell>
							<Badge variant="secondary">Draft</Badge>
						</Table.Cell>
						<Table.Cell class="hidden md:table-cell">$2.99</Table.Cell>
						<Table.Cell class="hidden md:table-cell">0</Table.Cell>
						<Table.Cell class="hidden md:table-cell">2023-12-25 11:59 PM</Table.Cell>
						<Table.Cell>
							<DropdownMenu.Root>
								<DropdownMenu.Trigger asChild let:builder>
									<Button builders={[builder]} aria-haspopup="true" size="icon" variant="ghost">
										<Ellipsis class="h-4 w-4" />
										<span class="sr-only">Toggle menu</span>
									</Button>
								</DropdownMenu.Trigger>
								<DropdownMenu.Content align="end">
									<DropdownMenu.Label>Actions</DropdownMenu.Label>
									<DropdownMenu.Item>Edit</DropdownMenu.Item>
									<DropdownMenu.Item>Delete</DropdownMenu.Item>
								</DropdownMenu.Content>
							</DropdownMenu.Root>
						</Table.Cell>
					</Table.Row>
					<Table.Row>
						<Table.Cell class="hidden sm:table-cell">
							<img
								alt="Product"
								class="aspect-square rounded-md object-cover"
								height="64"
								src="/images/placeholder.svg"
								width="64"
							/>
						</Table.Cell>
						<Table.Cell class="font-medium">Gamer Gear Pro Controller</Table.Cell>
						<Table.Cell>
							<Badge variant="outline">Active</Badge>
						</Table.Cell>
						<Table.Cell class="hidden md:table-cell">$59.99</Table.Cell>
						<Table.Cell class="hidden md:table-cell">75</Table.Cell>
						<Table.Cell class="hidden md:table-cell">2024-01-01 12:00 AM</Table.Cell>
						<Table.Cell>
							<DropdownMenu.Root>
								<DropdownMenu.Trigger asChild let:builder>
									<Button builders={[builder]} aria-haspopup="true" size="icon" variant="ghost">
										<Ellipsis class="h-4 w-4" />
										<span class="sr-only">Toggle menu</span>
									</Button>
								</DropdownMenu.Trigger>
								<DropdownMenu.Content align="end">
									<DropdownMenu.Label>Actions</DropdownMenu.Label>
									<DropdownMenu.Item>Edit</DropdownMenu.Item>
									<DropdownMenu.Item>Delete</DropdownMenu.Item>
								</DropdownMenu.Content>
							</DropdownMenu.Root>
						</Table.Cell>
					</Table.Row>
					<Table.Row>
						<Table.Cell class="hidden sm:table-cell">
							<img
								alt="Product"
								class="aspect-square rounded-md object-cover"
								height="64"
								src="/images/placeholder.svg"
								width="64"
							/>
						</Table.Cell>
						<Table.Cell class="font-medium">Luminous VR Headset</Table.Cell>
						<Table.Cell>
							<Badge variant="outline">Active</Badge>
						</Table.Cell>
						<Table.Cell class="hidden md:table-cell">$199.99</Table.Cell>
						<Table.Cell class="hidden md:table-cell">30</Table.Cell>
						<Table.Cell class="hidden md:table-cell">2024-02-14 02:14 PM</Table.Cell>
						<Table.Cell>
							<DropdownMenu.Root>
								<DropdownMenu.Trigger asChild let:builder>
									<Button builders={[builder]} aria-haspopup="true" size="icon" variant="ghost">
										<Ellipsis class="h-4 w-4" />
										<span class="sr-only">Toggle menu</span>
									</Button>
								</DropdownMenu.Trigger>
								<DropdownMenu.Content align="end">
									<DropdownMenu.Label>Actions</DropdownMenu.Label>
									<DropdownMenu.Item>Edit</DropdownMenu.Item>
									<DropdownMenu.Item>Delete</DropdownMenu.Item>
								</DropdownMenu.Content>
							</DropdownMenu.Root>
						</Table.Cell>
					</Table.Row> -->
				</Table.Body>
			</Table.Root>
		</Card.Content>
		<Card.Footer>
			<div class="text-muted-foreground text-xs">
				Showing <strong>1-10</strong> of <strong>32</strong> products
			</div>
		</Card.Footer>
	</Card.Root>
</UserPageLayout>
