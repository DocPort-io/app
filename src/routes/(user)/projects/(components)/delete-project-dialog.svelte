<script lang="ts">
	import type { DialogController } from '$lib/stores/dialog.svelte';

	import { LoaderCircle } from '@lucide/svelte';
	import { Button } from '$lib/components/ui/button';
	import * as Dialog from '$lib/components/ui/dialog';
	import FormErrors from '$lib/form/form-errors.svelte';
	import { createForm } from '$lib/form/form.svelte';
	import { m } from '$lib/paraglide/messages';
	import { createDeleteProjectMutation } from '$lib/queries/projects';
	import { type ProjectDeleteSchema } from '$lib/schemas/project.schema';
	import z from 'zod';

	type Props = {
		dialogController: DialogController<ProjectDeleteSchema>;
	};

	let { dialogController, ...restProps }: Props = $props();

	const deleteMutation = createDeleteProjectMutation();

	const schema = z.object({});

	const form = $derived(
		createForm({
			schema,
			onSubmit: async ({ setError }) => {
				if (!dialogController.data?.id) {
					setError('No project selected for deletion.');
					return;
				}

				try {
					await $deleteMutation.mutateAsync(dialogController.data.id);
					dialogController.close();
				} catch {
					setError('Failed to delete project. Please try again.');
				}
			}
		})
	);
</script>

<Dialog.Root bind:open={dialogController.isOpen} {...restProps}>
	<Dialog.Content class="sm:max-w-[425px]">
		<Dialog.Header>
			<Dialog.Title>{m.delete_project()}</Dialog.Title>
			<Dialog.Description>
				{m.are_you_sure_you_want_to_delete_this_project()}
			</Dialog.Description>
		</Dialog.Header>
		<form class="grid gap-4 py-4" {...form.props}>
			<FormErrors {form} />
			<Dialog.Footer>
				<Button variant="outline" onclick={() => dialogController.close()} disabled={form.state.isSubmitting}>
					{m.cancel()}
				</Button>
				<Button type="submit" variant="destructive" disabled={!form.state.isSubmittable}>
					{#if form.state.isSubmitting}
						<LoaderCircle class="mr-2 h-4 w-4 animate-spin" />
						{m.deleting()}
					{:else}
						{m.delete()}
					{/if}
				</Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
