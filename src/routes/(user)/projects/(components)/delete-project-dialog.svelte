<script lang="ts">
	import type { DialogController } from '$lib/stores/dialog.svelte';

	import { LoaderCircle } from '@lucide/svelte';
	import { Button } from '$lib/components/ui/button';
	import * as Dialog from '$lib/components/ui/dialog';
	import FormErrors from '$lib/form/form-errors.svelte';
	import { createForm } from '$lib/form/form.svelte';
	import { m } from '$lib/paraglide/messages';
	import { createDeleteProjectMutation } from '$lib/queries/projects';
	import { projectDeleteSchema, type ProjectDeleteSchema } from '$lib/schemas/project.schema';

	type Props = {
		dialogController: DialogController<ProjectDeleteSchema>;
	};

	let { dialogController, ...restProps }: Props = $props();

	const deleteMutation = createDeleteProjectMutation();

	const form = $derived(
		createForm({
			schema: projectDeleteSchema,
			defaultValues: {
				id: dialogController.data!.id
			},
			onSubmit: async ({ data, setError }) => {
				try {
					await $deleteMutation.mutateAsync(data.id);
					dialogController.close();
					form.reset();
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
				<Button
					variant="outline"
					onclick={() => dialogController.close()}
					disabled={form.state.isSubmitting}
				>
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
