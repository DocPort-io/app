<script lang="ts">
	import type { DialogController } from '$lib/stores/dialog.svelte';

	import { AlertTriangle, LoaderCircle } from '@lucide/svelte';
	import * as Alert from '$lib/components/ui/alert';
	import { Button } from '$lib/components/ui/button';
	import * as Dialog from '$lib/components/ui/dialog';
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import { m } from '$lib/paraglide/messages';
	import { createAddProjectMutation } from '$lib/queries/projects';
	import { projectCreateSchema } from '$lib/schemas/project.schema';
	import { getTeamState } from '$lib/stores/team.svelte';
	import { defaults, setError, superForm } from 'sveltekit-superforms';
	import { zod, zodClient } from 'sveltekit-superforms/adapters';

	type Props = {
		dialogController: DialogController<unknown>;
	};

	let { dialogController, ...restProps }: Props = $props();

	const teamState = getTeamState();

	const addMutation = createAddProjectMutation();

	const form = superForm(defaults(zod(projectCreateSchema)), {
		id: 'create-project-form',
		SPA: true,
		validators: zodClient(projectCreateSchema),
		onUpdate: async ({ form }) => {
			if (!form.valid) return;
			if (!teamState.currentTeam) return setError(form, 'Please select a team first.');

			await $addMutation.mutateAsync(
				{
					...form.data,
					status: 'active',
					team: teamState.currentTeam
				},
				{
					onSuccess: () => {
						dialogController.close();
					},
					onError: () => {
						setError(form, 'Failed to create project. Please try again.');
					}
				}
			);
		}
	});

	const { form: formData, constraints, enhance, submitting, delayed, allErrors } = form;

	let formErrors = $derived(
		$allErrors.filter((error) => error.path === '_errors').flatMap((error) => error.messages)
	);
</script>

<Dialog.Root bind:open={dialogController.isOpen} {...restProps}>
	<Dialog.Content class="sm:max-w-[425px]">
		<Dialog.Header>
			<Dialog.Title>{m.weak_weak_bulldog_assure()}</Dialog.Title>
			<Dialog.Description>{m.proof_noisy_monkey_type()}</Dialog.Description>
		</Dialog.Header>

		{#if formErrors.length > 0}
			<Alert.Root variant="destructive" class="mt-4">
				<AlertTriangle class="mr-2 h-4 w-4" />
				<Alert.Title>Error</Alert.Title>
				<Alert.Description>{formErrors[0]}</Alert.Description>
			</Alert.Root>
		{/if}

		<form method="POST" class="grid gap-4 py-4" use:enhance>
			<Form.Field {form} name="name">
				<Form.Control>
					{#snippet children({ props }: { props: object })}
						<Form.Label>{m.royal_major_impala_charm()}</Form.Label>
						<Input
							{...props}
							{...$constraints.name}
							bind:value={$formData.name}
							placeholder={m.alert_nimble_pug_play()}
							disabled={$submitting}
						/>
					{/snippet}
				</Form.Control>
				<Form.Description>{m.busy_tame_jackdaw_read()}</Form.Description>
				<Form.FieldErrors />
			</Form.Field>

			<Dialog.Footer>
				<Button
					type="reset"
					variant="outline"
					on:click={() => dialogController.close()}
					disabled={$submitting}
				>
					{m.red_same_flea_clip()}
				</Button>
				<Form.Button type="submit" disabled={$submitting}>
					{#if $delayed}
						<LoaderCircle class="mr-2 h-4 w-4 animate-spin" />
					{/if}
					{#if $submitting}
						Creating...
					{:else}
						{m.hour_swift_crab_breathe()}
					{/if}
				</Form.Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
