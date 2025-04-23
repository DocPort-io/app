<script lang="ts">
	import type { PaginationController } from '$lib/stores/pagination.svelte';

	import * as Pagination from './ui/pagination';

	type Props = {
		pagination: PaginationController;
		totalItems: number;
	};

	let { pagination, totalItems }: Props = $props();
</script>

<div>
	<Pagination.Root
		count={totalItems}
		perPage={pagination.perPage}
		let:pages
		let:currentPage
		bind:page={pagination.page}
	>
		<Pagination.Content>
			<Pagination.Item>
				<Pagination.PrevButton />
			</Pagination.Item>
			{#each pages as page (page.key)}
				{#if page.type === 'ellipsis'}
					<Pagination.Item>
						<Pagination.Ellipsis />
					</Pagination.Item>
				{:else}
					<Pagination.Item>
						<Pagination.Link {page} isActive={currentPage == page.value}>
							{page.value}
						</Pagination.Link>
					</Pagination.Item>
				{/if}
			{/each}
			<Pagination.Item>
				<Pagination.NextButton />
			</Pagination.Item>
		</Pagination.Content>
	</Pagination.Root>
</div>
