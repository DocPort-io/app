<script lang="ts">
	import type { DialogController } from '$lib/stores/dialog.svelte';

	import { Button } from '$lib/components/ui/button';
	import * as Dialog from '$lib/components/ui/dialog';
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import { m } from '$lib/paraglide/messages';
	import { projectUpdateSchema, type ProjectUpdateSchema } from '$lib/schemas/project.schema';
	import { defaults, superForm } from 'sveltekit-superforms';
	import { zod, zodClient } from 'sveltekit-superforms/adapters';

	type Props = {
		dialogController: DialogController<{ id: string; project: ProjectUpdateSchema }>;
		handleUpdateProject: (id: string, data: ProjectUpdateSchema) => unknown;
	};

	let { dialogController, handleUpdateProject, ...restProps }: Props = $props();

	const form = $derived(
		superForm(defaults(dialogController.data?.project, zod(projectUpdateSchema)), {
			id: 'edit-project-form',
			SPA: true,
			validators: zodClient(projectUpdateSchema),
			onUpdate: async ({ form }) => {
				if (!form.valid) return;
				if (!dialogController.data?.id) return;
				await handleUpdateProject(dialogController.data.id, form.data);
				dialogController.close();
			}
		})
	);

	const { form: formData, constraints, enhance, validateForm } = $derived(form);

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
					{#snippet children({ props }: { props: object })}
						<Form.Label>{m.royal_major_impala_charm()}</Form.Label>
						<Input
							{...props}
							{...$constraints.name}
							bind:value={$formData.name}
							placeholder={m.alert_nimble_pug_play()}
						/>
					{/snippet}
				</Form.Control>
				<Form.Description>{m.busy_tame_jackdaw_read()}</Form.Description>
				<Form.FieldErrors />
			</Form.Field>
			<Dialog.Footer>
				<Button variant="outline" on:click={() => dialogController.close()}>
					{m.red_same_flea_clip()}
				</Button>
				<Form.Button type="submit">
					{m.big_male_bear_peek()}
				</Form.Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
