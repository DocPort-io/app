<script lang="ts">
	import { m } from '$lib/paraglide/messages';

	type Props = {
		results: number;
		page: number;
		perPage: number;
		totalItems: number;
	};

	let { results, page, perPage, totalItems }: Props = $props();

	const resultsStart = $derived(totalItems === 0 ? 0 : (page - 1) * perPage + 1);
	const resultsEnd = $derived(totalItems === 0 ? 0 : (page - 1) * perPage + 1 + results - 1);
</script>

<div class="text-muted-foreground text-xs">
	{m.weird_sharp_javelina_sail({
		start: resultsStart,
		end: resultsEnd,
		amount: totalItems
	})}
</div>
