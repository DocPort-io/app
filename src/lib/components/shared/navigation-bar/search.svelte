<script lang="ts">
	import { Search } from '@lucide/svelte';
	import { Input } from '$lib/components/ui/input';
	import { m } from '$lib/paraglide/messages';

	let search = $state('');
</script>

<form class="ml-auto flex-1 sm:flex-initial">
	<div class="relative">
		<Search class="text-muted-foreground absolute top-2.5 left-2.5 h-4 w-4" />
		<Input
			type="search"
			placeholder={m.search()}
			class="pl-8 sm:w-[300px] md:w-[200px] lg:w-[300px]"
			bind:value={search}
		/>
	</div>
</form>
