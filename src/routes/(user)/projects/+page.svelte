<script lang="ts">
	import type { ProjectUpdateSchema } from '$lib/schemas/project.schema';

	import { CirclePlus } from '@lucide/svelte';
	import { Ellipsis } from '@lucide/svelte';
	import { createQuery } from '@tanstack/svelte-query';
	import { goto } from '$app/navigation';
	import UserPageLayout from '$lib/components/layouts/user-page-layout.svelte';
	import { Badge } from '$lib/components/ui/badge';
	import { Button } from '$lib/components/ui/button';
	import { buttonVariants } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import {
		Pagination,
		PaginationContent,
		PaginationEllipsis,
		PaginationItem,
		PaginationLink,
		PaginationNextButton,
		PaginationPrevButton
	} from '$lib/components/ui/pagination';
	import { Skeleton } from '$lib/components/ui/skeleton';
	import * as Table from '$lib/components/ui/table';
	import { AppRoute } from '$lib/constants';
	import { m } from '$lib/paraglide/messages';
	import { createPaginatedProjectsQuery } from '$lib/queries/projects';
	import { createDialogController } from '$lib/stores/dialog.svelte';
	import { getTeamState } from '$lib/stores/team.svelte';

	import CreateProjectDialog from './(components)/create-project-dialog.svelte';
	import DeleteProjectDialog from './(components)/delete-project-dialog.svelte';
	import EditProjectDialog from './(components)/edit-project-dialog.svelte';

	// Dialog handlers
	const createDialog = createDialogController();
	const editDialog = createDialogController<{ id: string; project: ProjectUpdateSchema }>();
	const deleteDialog = createDialogController<{ id: string }>();

	const teamState = getTeamState();

	const pagination = $state({
		page: 1,
		perPage: 5
	});

	const projects = $derived.by(() =>
		createQuery(
			createPaginatedProjectsQuery({
				team: teamState.currentTeam ?? '',
				page: pagination.page,
				perPage: pagination.perPage
			})
		)
	);

	const statusMap = {
		planned: 'Planned',
		active: 'Active',
		completed: 'Completed'
	};
</script>

<UserPageLayout title="Projects">
	<Card.Root data-testid="projects-card">
		<Card.Header class="flex flex-row items-center">
			<div class="grid gap-2">
				<Card.Title>{m.projects()}</Card.Title>
				<Card.Description>{m.manage_your_projects()}</Card.Description>
			</div>
			<div class="ml-auto flex items-center gap-2">
				<Button
					size="sm"
					class="h-8 gap-1"
					onclick={() => createDialog.open()}
					data-testid="projects-create-button"
				>
					<CirclePlus class="h-3.5 w-3.5" />
					<span class="sr-only sm:not-sr-only sm:whitespace-nowrap">{m.add_project()}</span>
				</Button>
			</div>
		</Card.Header>
		<Card.Content>
			{#if $projects.isError}
				<p>{$projects.error.message}</p>
			{/if}
			<Table.Root data-testid="projects-table">
				<Table.Header data-testid="projects-table-header">
					<Table.Row>
						<Table.Head class="w-full md:w-2/3">{m.name()}</Table.Head>
						<Table.Head class="hidden md:table-cell md:w-1/3">Status</Table.Head>
						<Table.Head>
							<span class="sr-only">Actions</span>
						</Table.Head>
					</Table.Row>
				</Table.Header>
				<Table.Body data-testid="projects-table-body">
					{#if $projects.isLoading}
						<Table.Row>
							<Table.Cell class="font-medium">
								<Skeleton class="h-4 w-[200px] md:w-[300px]" />
							</Table.Cell>
							<Table.Cell class="hidden md:table-cell">
								<Skeleton class="h-4 w-[200px]" />
							</Table.Cell>
							<Table.Cell></Table.Cell>
						</Table.Row>

						<Table.Row>
							<Table.Cell class="font-medium">
								<Skeleton class="h-4 w-[200px] md:w-[300px]" />
							</Table.Cell>
							<Table.Cell class="hidden md:table-cell">
								<Skeleton class="h-4 w-[200px]" />
							</Table.Cell>
							<Table.Cell></Table.Cell>
						</Table.Row>

						<Table.Row>
							<Table.Cell class="font-medium">
								<Skeleton class="h-4 w-[200px] md:w-[300px]" />
							</Table.Cell>
							<Table.Cell class="hidden md:table-cell">
								<Skeleton class="h-4 w-[200px]" />
							</Table.Cell>
							<Table.Cell></Table.Cell>
						</Table.Row>
					{/if}
					{#if $projects.isError}
						<Table.Row>
							<Table.Cell colspan={4}>
								{$projects.error.message}
							</Table.Cell>
						</Table.Row>
					{/if}
					{#if $projects.isSuccess}
						{#each $projects.data.items as project (project.id)}
							<Table.Row data-testid="projects-table-row">
								<Table.Cell class="font-medium">
									{project.name}
								</Table.Cell>
								<Table.Cell class="hidden md:table-cell">
									<Badge variant="outline">{statusMap[project.status]}</Badge>
								</Table.Cell>
								<Table.Cell>
									<DropdownMenu.Root>
										<DropdownMenu.Trigger
											aria-haspopup="true"
											class={buttonVariants({ size: 'icon', variant: 'ghost' })}
										>
											<Ellipsis class="h-4 w-4" />
											<span class="sr-only">Toggle menu</span>
										</DropdownMenu.Trigger>
										<DropdownMenu.Content align="end">
											<DropdownMenu.Label>{m.actions()}</DropdownMenu.Label>
											<DropdownMenu.Item onclick={() => goto(AppRoute.PROJECT_VIEW(project.id))}>
												{m.view()}
											</DropdownMenu.Item>
											<DropdownMenu.Item
												onclick={() => {
													editDialog.data = { id: project.id, project };
													editDialog.open();
												}}
											>
												{m.edit()}
											</DropdownMenu.Item>
											<DropdownMenu.Item
												onclick={() => {
													deleteDialog.data = project;
													deleteDialog.open();
												}}
											>
												{m.delete()}
											</DropdownMenu.Item>
										</DropdownMenu.Content>
									</DropdownMenu.Root>
								</Table.Cell>
							</Table.Row>
						{/each}
					{/if}
				</Table.Body>
			</Table.Root>
		</Card.Content>
		{#if $projects.isSuccess && $projects.data.totalItems > 0}
			<Card.Footer>
				<Pagination
					count={$projects.data.totalItems}
					perPage={pagination.perPage}
					bind:page={pagination.page}
				>
					{#snippet children({ pages, currentPage, range })}
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
						<p class="text-muted-foreground text-center text-[13px]">
							Showing {range.start} - {range.end} of {$projects.data.totalItems} results
						</p>
					{/snippet}
				</Pagination>
			</Card.Footer>
		{/if}
	</Card.Root>

	<CreateProjectDialog dialogController={createDialog} />
	<EditProjectDialog dialogController={editDialog} />
	<DeleteProjectDialog dialogController={deleteDialog} />
</UserPageLayout>
