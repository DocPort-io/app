<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import * as Dialog from '$lib/components/ui/dialog';
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import { m } from '$lib/paraglide/messages';
	import { projectCreateSchema } from '$lib/schemas/project.schema';
	import { getProjects } from '$lib/states/projects.svelte';
	import { defaults, superForm } from 'sveltekit-superforms';
	import { zod, zodClient } from 'sveltekit-superforms/adapters';

	type Props = {
		open?: boolean;
	};

	let { open = $bindable(false), ...restProps }: Props = $props();

	const projectsState = getProjects();

	const form = superForm(defaults(zod(projectCreateSchema)), {
		SPA: true,
		validators: zodClient(projectCreateSchema),
		onUpdate: async ({ form }) => {
			if (!form.valid) return;

			await projectsState.add(form.data);
			open = false;
		}
	});

	const { form: formData, constraints, enhance, validateForm } = form;
</script>

<Dialog.Root bind:open {...restProps}>
	<Dialog.Content class="sm:max-w-[425px]">
		<Dialog.Header>
			<Dialog.Title>{m.weak_weak_bulldog_assure()}</Dialog.Title>
			<Dialog.Description>{m.proof_noisy_monkey_type()}</Dialog.Description>
		</Dialog.Header>
		<form method="POST" class="grid gap-4 py-4" use:enhance>
			<Form.Field {form} name="name">
				<Form.Control>
					{#snippet children({ props }: { props: any })}
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
				<Button type="reset" variant="outline" on:click={() => (open = false)}
					>{m.red_same_flea_clip()}
				</Button>
				<Form.Button type="submit">
					{m.hour_swift_crab_breathe()}
				</Form.Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
