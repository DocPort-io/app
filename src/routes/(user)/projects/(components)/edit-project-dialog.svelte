<script lang="ts">
	import type { DialogController } from '$lib/stores/dialog.svelte';

	import { LoaderCircle } from '@lucide/svelte';
	import { Button } from '$lib/components/ui/button';
	import * as Dialog from '$lib/components/ui/dialog';
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import { m } from '$lib/paraglide/messages';
	import { createUpdateProjectMutation } from '$lib/queries/projects';
	import { projectUpdateSchema, type ProjectUpdateSchema } from '$lib/schemas/project.schema';
	import { defaults, setError, superForm } from 'sveltekit-superforms';
	import { zod, zodClient } from 'sveltekit-superforms/adapters';

	type Props = {
		dialogController: DialogController<{ id: string; project: ProjectUpdateSchema }>;
	};

	let { dialogController, ...restProps }: Props = $props();

	const updateMutation = createUpdateProjectMutation();

	const form = $derived(
		superForm(defaults(dialogController.data?.project, zod(projectUpdateSchema)), {
			id: 'edit-project-form',
			SPA: true,
			validators: zodClient(projectUpdateSchema),
			onUpdate: async ({ form }) => {
				if (!form.valid) return;
				if (!dialogController.data?.id) return;

				await $updateMutation.mutateAsync(
					{ id: dialogController.data.id, project: form.data },
					{
						onSuccess: () => {
							dialogController.close();
						},
						onError: () => {
							setError(form, 'Failed to update project. Please try again.');
						}
					}
				);
			}
		})
	);

	const {
		form: formData,
		constraints,
		enhance,
		validateForm,
		submitting,
		delayed
	} = $derived(form);

	$effect(() => {
		validateForm({ update: true });
	});
</script>

<Dialog.Root bind:open={dialogController.isOpen} {...restProps}>
	<Dialog.Content class="sm:max-w-[425px]">
		<Dialog.Header>
			<Dialog.Title>{m.known_major_hawk_stop()}</Dialog.Title>
			<Dialog.Description>{m.factual_cute_goose_snap()}</Dialog.Description>
		</Dialog.Header>
		<form method="POST" class="grid gap-4 py-4" use:enhance>
			<Form.Field {form} name="name">
				<Form.Control>
					{#snippet children({ props })}
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
				<Button variant="outline" onclick={() => dialogController.close()} disabled={$submitting}>
					{m.red_same_flea_clip()}
				</Button>
				<Form.Button type="submit" disabled={$submitting}>
					{#if $delayed}
						<LoaderCircle class="mr-2 h-4 w-4 animate-spin" />
					{/if}
					{#if $submitting}
						Saving...
					{:else}
						{m.big_male_bear_peek()}
					{/if}
				</Form.Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
