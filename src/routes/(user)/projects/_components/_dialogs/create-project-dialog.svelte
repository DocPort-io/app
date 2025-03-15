<script lang="ts">
	import type { DialogController } from '$lib/stores/dialog.svelte';

	import { Button } from '$lib/components/ui/button';
	import * as Dialog from '$lib/components/ui/dialog';
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import { m } from '$lib/paraglide/messages';
	import { projectCreateSchema, type ProjectCreateSchema } from '$lib/schemas/project.schema';
	import { defaults, superForm } from 'sveltekit-superforms';
	import { zod, zodClient } from 'sveltekit-superforms/adapters';

	type Props = {
		dialogController: DialogController<unknown>;
		handleCreateProject: (data: ProjectCreateSchema) => unknown;
	};

	let { dialogController, handleCreateProject, ...restProps }: Props = $props();

	const form = superForm(defaults(zod(projectCreateSchema)), {
		id: 'create-project-form',
		SPA: true,
		validators: zodClient(projectCreateSchema),
		onUpdate: async ({ form }) => {
			if (!form.valid) return;
			await handleCreateProject({ ...form.data, team: 'wdqkmp0jbmm5ua5' });
			dialogController.close();
		}
	});

	const { form: formData, constraints, enhance } = form;
</script>

<Dialog.Root bind:open={dialogController.isOpen} {...restProps}>
	<Dialog.Content class="sm:max-w-[425px]">
		<Dialog.Header>
			<Dialog.Title>{m.weak_weak_bulldog_assure()}</Dialog.Title>
			<Dialog.Description>{m.proof_noisy_monkey_type()}</Dialog.Description>
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
				<Button type="reset" variant="outline" on:click={() => dialogController.close()}
					>{m.red_same_flea_clip()}
				</Button>
				<Form.Button type="submit">
					{m.hour_swift_crab_breathe()}
				</Form.Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
