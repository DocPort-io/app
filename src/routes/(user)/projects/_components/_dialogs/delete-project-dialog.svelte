<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import * as Dialog from '$lib/components/ui/dialog';
	import * as Form from '$lib/components/ui/form';
	import { m } from '$lib/paraglide/messages';
	import { projectDeleteSchema, type ProjectDeleteSchema } from '$lib/schemas/project.schema';
	import { getProjects } from '$lib/states/projects.svelte';
	import { defaults, superForm } from 'sveltekit-superforms';
	import { zod, zodClient } from 'sveltekit-superforms/adapters';

	type Props = {
		open?: boolean;
		project: ProjectDeleteSchema;
	};

	let { open = $bindable(false), project, ...restProps }: Props = $props();

	const projectsState = getProjects();

	const form = $derived(
		superForm(defaults(project, zod(projectDeleteSchema)), {
			SPA: true,
			validators: zodClient(projectDeleteSchema),
			onUpdate: async ({ form }) => {
				if (!form.valid) return;

				await projectsState.remove(form.data);
				open = false;
			}
		})
	);

	$effect(() => {
		form.validateForm({ update: true });
	});
</script>

<Dialog.Root bind:open {...restProps}>
	<Dialog.Content class="sm:max-w-[425px]">
		<Dialog.Header>
			<Dialog.Title>{m.ornate_wise_lemur_dance()}</Dialog.Title>
			<Dialog.Description>
				{m.equal_crisp_crow_edit()}
			</Dialog.Description>
		</Dialog.Header>
		<form method="POST" class="grid gap-4 py-4" use:form.enhance>
			<Dialog.Footer>
				<Button type="reset" variant="outline" on:click={() => (open = false)}
					>{m.red_same_flea_clip()}
				</Button>
				<Form.Button type="submit" variant="destructive">
					{m.mellow_dark_puma_boil()}
				</Form.Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
