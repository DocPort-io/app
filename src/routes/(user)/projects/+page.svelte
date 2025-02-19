<script lang="ts">
	import UserPageLayout from '$lib/components/layouts/user-page-layout.svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
</script>

<UserPageLayout>
	<h1>Test</h1>
</UserPageLayout>
