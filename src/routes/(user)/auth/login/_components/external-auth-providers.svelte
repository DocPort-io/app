<script lang="ts">
	import { Fingerprint } from '@lucide/svelte';
	import { Button } from '$lib/components/ui/button';
	import { Separator } from '$lib/components/ui/separator';
	import { Skeleton } from '$lib/components/ui/skeleton';
	import { useOAuth2Providers } from '$lib/hooks/useOAuth2Providers.svelte';
	import { m } from '$lib/paraglide/messages';
	import { getUserState } from '$lib/stores/user.svelte';

	const userStore = getUserState();
	const oauth2Providers = useOAuth2Providers();
</script>

<div class="mt-6">
	<div class="relative">
		<div class="absolute inset-0 flex items-center">
			<Separator />
		</div>
		<div class="relative flex justify-center text-xs uppercase">
			<span class="bg-card text-muted-foreground px-2">{m.or_continue_with()}</span>
		</div>
	</div>

	<div class="mt-6 grid grid-cols-1 gap-4">
		{#if oauth2Providers.loading}
			<Skeleton class="h-10 w-full" />
		{:else if oauth2Providers.error}
			<Button variant="outline" class="w-full" disabled>
				<Fingerprint class="mr-2 h-4 w-4" />
				{m.external_authentication_not_possible()}
			</Button>
		{:else if oauth2Providers.providers.length === 0}
			<Button variant="outline" class="w-full" disabled>
				<Fingerprint class="mr-2 h-4 w-4" />
				{m.no_external_authentication_available()}
			</Button>
		{:else}
			{#each oauth2Providers.providers as oauth2Provider}
				<Button
					variant="outline"
					class="w-full"
					onclick={() => userStore.signInWithExternalProvider(oauth2Provider.name)}
				>
					<Fingerprint class="mr-2 h-4 w-4" />
					{oauth2Provider.displayName}
				</Button>
			{/each}
		{/if}
	</div>
</div>
