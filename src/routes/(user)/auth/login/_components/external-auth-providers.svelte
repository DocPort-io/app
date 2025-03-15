<script lang="ts">
	import { Fingerprint } from '@lucide/svelte';
	import { Button } from '$lib/components/ui/button';
	import { Separator } from '$lib/components/ui/separator';
	import { m } from '$lib/paraglide/messages';
	import { getUserState } from '$lib/stores/user.svelte';
	import { type AuthProviderInfo } from 'pocketbase';

	const userStore = getUserState();

	let externalAuthProviders = $state<AuthProviderInfo[]>([]);

	$effect(() => {
		(async () => {
			externalAuthProviders = await userStore.getOAuth2Providers();
		})();
	});
</script>

<div class="mt-6">
	<div class="relative">
		<div class="absolute inset-0 flex items-center">
			<Separator />
		</div>
		<div class="relative flex justify-center text-xs uppercase">
			<span class="bg-card text-muted-foreground px-2">{m.warm_bland_deer_bend()}</span>
		</div>
	</div>

	<div class="mt-6 grid grid-cols-1 gap-4">
		{#each externalAuthProviders as externalAuthProvider}
			<Button
				variant="outline"
				class="w-full"
				onclick={() => userStore.signInWithExternalProvider(externalAuthProvider.name)}
			>
				<Fingerprint class="mr-2 h-4 w-4" />
				{externalAuthProvider.displayName}
			</Button>
		{/each}
	</div>
</div>
