<script lang="ts">
	import { AlertTriangle, LoaderCircle, Lock } from '@lucide/svelte';
	import * as Alert from '$lib/components/ui/alert';
	import { Button } from '$lib/components/ui/button';
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import { m } from '$lib/paraglide/messages';
	import { signInSchema } from '$lib/schemas/sign-in.schema';
	import { getUserState } from '$lib/stores/user.svelte';
	import { ClientResponseError } from 'pocketbase';
	import { defaults, setError, superForm } from 'sveltekit-superforms';
	import { zod } from 'sveltekit-superforms/adapters';

	const userStore = getUserState();

	const form = superForm(defaults(zod(signInSchema)), {
		id: 'sign-in-form',
		SPA: true,
		validators: zod(signInSchema),
		onUpdate: async ({ form }) => {
			if (!form.valid) return;

			const { email, password } = form.data;

			try {
				await userStore.signIn(email, password);
			} catch (err) {
				if (err instanceof ClientResponseError && err.message === 'Failed to authenticate.') {
					setError(form, m.tasty_seemly_anaconda_kiss());
				} else {
					setError(form, m.key_white_crab_coax());
				}
			}
		}
	});

	const { form: formData, constraints, enhance, submitting, allErrors } = form;

	let formErrors = $derived(
		$allErrors.filter((error) => error.path === '_errors').flatMap((error) => error.messages)
	);
</script>

{#if formErrors.length > 0}
	<Alert.Root variant="destructive" class="mt-4">
		<AlertTriangle class="mr-2 h-4 w-4" />
		<Alert.Title data-testid="login-error-title">Oops!</Alert.Title>
		{#each formErrors as formError}
			<Alert.Description data-testid="login-error-description">{formError}</Alert.Description>
		{/each}
	</Alert.Root>
{/if}

<form method="POST" class="grid gap-4 py-4" use:enhance>
	<Form.Field {form} name="email">
		<Form.Control>
			{#snippet children({ props }: { props: object })}
				<Form.Label>{m.spare_muddy_piranha_feast()}</Form.Label>
				<Input
					{...props}
					{...$constraints.email}
					bind:value={$formData.email}
					placeholder={m.heavy_honest_ladybug_gasp()}
					type="email"
					data-testid="login-email"
				/>
			{/snippet}
		</Form.Control>
		<Form.Description>{m.aloof_quaint_manatee_hurl()}</Form.Description>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Field {form} name="password">
		<Form.Control>
			{#snippet children({ props }: { props: object })}
				<div class="flex items-center justify-between">
					<Form.Label>{m.fine_frail_cobra_reside()}</Form.Label>
					<a href="/auth/reset-password" class="text-primary text-sm hover:underline">
						{m.sea_sea_myna_quell()}
					</a>
				</div>
				<Input
					{...props}
					{...$constraints.password}
					bind:value={$formData.password}
					type="password"
					data-testid="login-password"
				/>
			{/snippet}
		</Form.Control>
		<Form.Description>{m.pink_mean_hyena_express()}</Form.Description>
		<Form.FieldErrors />
	</Form.Field>

	<Button type="submit" class="w-full" disabled={$submitting} data-testid="login-signin">
		{#if $submitting}
			<span class="mr-2 animate-spin"><LoaderCircle /></span>
			{m.salty_best_zebra_race()}
		{:else}
			<Lock class="mr-2 h-4 w-4" />
			{m.actual_zany_grebe_ascend()}
		{/if}
	</Button>
</form>
