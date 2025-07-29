<script lang="ts">
	import type { DialogController } from '$lib/stores/dialog.svelte';

	import { LoaderCircle } from '@lucide/svelte';
	import { Button } from '$lib/components/ui/button';
	import * as Dialog from '$lib/components/ui/dialog';
	import { Input } from '$lib/components/ui/input';
	import { Textarea } from '$lib/components/ui/textarea';
	import FieldErrors from '$lib/form/field-errors.svelte';
	import FieldLabel from '$lib/form/field-label.svelte';
	import Field from '$lib/form/field.svelte';
	import FormErrors from '$lib/form/form-errors.svelte';
	import { createForm } from '$lib/form/form.svelte';
	import { m } from '$lib/paraglide/messages';
	import { createAddVersionMutation } from '$lib/queries/versions';
	import { versionCreateSchema } from '$lib/schemas/version.schema';
	import { getUserState } from '$lib/stores/user.svelte';

	type Props = {
		dialogController: DialogController<{ projectId: string }>;
	};

	let { dialogController }: Props = $props();

	const addVersionMutation = createAddVersionMutation();
	const userState = getUserState();

	const form = $derived(
		createForm({
			schema: versionCreateSchema,
			defaultValues: {
				name: '',
				description: '',
				project: dialogController.data?.projectId ?? '',
				createdBy: userState.userId ?? ''
			},
			onSubmit: async ({ data, setError }) => {
				if (!dialogController.data?.projectId) {
					setError(m.project_id_required());
					return;
				}

				await $addVersionMutation.mutateAsync(
					{
						...data,
						project: dialogController.data.projectId,
						createdBy: userState.userId ?? ''
					},
					{
						onSuccess: () => {
							dialogController.close();
							form.reset();
						},
						onError: () => {
							setError(m.failed_to_create_version());
						}
					}
				);
			}
		})
	);
</script>

<Dialog.Root bind:open={dialogController.isOpen}>
	<Dialog.Content class="sm:max-w-[425px]" data-testid="create-version-dialog">
		<Dialog.Header>
			<Dialog.Title>{m.create_version()}</Dialog.Title>
			<Dialog.Description>{m.create_a_new_version_of_this_project()}</Dialog.Description>
		</Dialog.Header>

		<FormErrors {form} />

		<form {...form.props} class="grid gap-4 py-4">
			<Field {form} name="name">
				{#snippet children({ props, state })}
					<FieldLabel>{m.version_name()}</FieldLabel>
					<Input {...props} bind:value={state.value} disabled={form.state.isSubmitting} />
					<FieldErrors />
				{/snippet}
			</Field>

			<Field {form} name="description">
				{#snippet children({ props, state })}
					<FieldLabel>Description ({m.optional()})</FieldLabel>
					<Textarea
						{...props}
						bind:value={state.value}
						disabled={form.state.isSubmitting}
						rows={3}
					/>
					<FieldErrors />
				{/snippet}
			</Field>

			<Dialog.Footer>
				<Button
					variant="outline"
					onclick={() => {
						form.reset();
						dialogController.close();
					}}
					disabled={form.state.isSubmitting}
				>
					{m.cancel()}
				</Button>
				<Button type="submit" disabled={!form.state.isSubmittable}>
					{#if form.state.isSubmitting}
						<LoaderCircle class="mr-2 h-4 w-4 animate-spin" />
					{/if}
					{#if form.state.isSubmitting}
						{m.creating()}
					{:else}
						{m.create()}
					{/if}
				</Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
