<script lang="ts">
	import type { PaginationController } from '$lib/stores/pagination.svelte';

	import { m } from '$lib/paraglide/messages';

	type Props = {
		amountShown: number;
		pagination: PaginationController;
	};

	let { amountShown, pagination }: Props = $props();

	const resultsStart = $derived(
		pagination.totalItems === 0 ? 0 : (pagination.page - 1) * pagination.perPage + 1
	);
	const resultsEnd = $derived(
		pagination.totalItems === 0
			? 0
			: (pagination.page - 1) * pagination.perPage + 1 + amountShown - 1
	);
</script>

<div class="text-muted-foreground text-xs">
	{m.weird_sharp_javelina_sail({
		start: resultsStart,
		end: resultsEnd,
		amount: pagination.totalItems
	})}
</div>
