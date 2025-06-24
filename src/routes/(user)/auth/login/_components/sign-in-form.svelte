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
					setError(form, m.invalid_email_or_password());
				} else {
					setError(form, m.an_unexpected_error_occurred());
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
		<Alert.Title data-testid="login-error-title">{m.oops()}</Alert.Title>
		{#each formErrors as formError}
			<Alert.Description data-testid="login-error-description">{formError}</Alert.Description>
		{/each}
	</Alert.Root>
{/if}

<form method="POST" class="grid gap-4 py-4" use:enhance>
	<Form.Field {form} name="email">
		<Form.Control>
			{#snippet children({ props }: { props: object })}
				<Form.Label>{m.email()}</Form.Label>
				<Input
					{...props}
					{...$constraints.email}
					bind:value={$formData.email}
					placeholder={m.name_at_example_dot_com()}
					type="email"
					data-testid="login-email"
				/>
			{/snippet}
		</Form.Control>
		<Form.Description>{m.enter_your_email_address()}</Form.Description>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Field {form} name="password">
		<Form.Control>
			{#snippet children({ props }: { props: object })}
				<div class="flex items-center justify-between">
					<Form.Label>{m.password()}</Form.Label>
					<a href="/auth/reset-password" class="text-primary text-sm hover:underline">
						{m.forgot_password()}
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
		<Form.Description>{m.enter_your_password()}</Form.Description>
		<Form.FieldErrors />
	</Form.Field>

	<Button type="submit" class="w-full" disabled={$submitting} data-testid="login-signin">
		{#if $submitting}
			<span class="mr-2 animate-spin"><LoaderCircle /></span>
			{m.signing_in()}
		{:else}
			<Lock class="mr-2 h-4 w-4" />
			{m.sign_in()}
		{/if}
	</Button>
</form>
