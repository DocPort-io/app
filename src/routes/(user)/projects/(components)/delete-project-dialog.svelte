<script lang="ts">
	import type { DialogController } from '$lib/stores/dialog.svelte';

	import { LoaderCircle } from '@lucide/svelte';
	import { Button } from '$lib/components/ui/button';
	import * as Dialog from '$lib/components/ui/dialog';
	import * as Form from '$lib/components/ui/form';
	import { m } from '$lib/paraglide/messages';
	import { createDeleteProjectMutation } from '$lib/queries/projects';
	import { projectDeleteSchema, type ProjectDeleteSchema } from '$lib/schemas/project.schema';
	import { defaults, setError, superForm } from 'sveltekit-superforms';
	import { zod, zodClient } from 'sveltekit-superforms/adapters';

	type Props = {
		dialogController: DialogController<ProjectDeleteSchema>;
	};

	let { dialogController, ...restProps }: Props = $props();

	const deleteMutation = createDeleteProjectMutation();

	const form = $derived(
		superForm(defaults(dialogController.data, zod(projectDeleteSchema)), {
			id: 'delete-project-form',
			SPA: true,
			validators: zodClient(projectDeleteSchema),
			onUpdate: async ({ form }) => {
				if (!form.valid) return;

				await $deleteMutation.mutateAsync(form.data.id, {
					onSuccess: () => {
						dialogController.close();
					},
					onError: () => {
						setError(form, 'Failed to delete project. Please try again.');
					}
				});
			}
		})
	);

	const { enhance, validateForm, submitting, delayed } = $derived(form);

	$effect(() => {
		validateForm({ update: true });
	});
</script>

<Dialog.Root bind:open={dialogController.isOpen} {...restProps}>
	<Dialog.Content class="sm:max-w-[425px]">
		<Dialog.Header>
			<Dialog.Title>{m.delete_project()}</Dialog.Title>
			<Dialog.Description>
				{m.are_you_sure_you_want_to_delete_this_project()}
			</Dialog.Description>
		</Dialog.Header>
		<form method="POST" class="grid gap-4 py-4" use:enhance>
			<Dialog.Footer>
				<Button variant="outline" onclick={() => dialogController.close()} disabled={$submitting}>
					{m.cancel()}
				</Button>
				<Form.Button type="submit" variant="destructive" disabled={$submitting}>
					{#if $delayed}
						<LoaderCircle class="mr-2 h-4 w-4 animate-spin" />
					{/if}
					{#if $submitting}
						Deleting...
					{:else}
						{m.delete()}
					{/if}
				</Form.Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
