<script lang="ts">
	import type { VersionSchema } from '$lib/schemas/version.schema';
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
	import { createUpdateVersionMutation } from '$lib/queries/versions';
	import { versionUpdateSchema } from '$lib/schemas/version.schema';
	import { getUserState } from '$lib/stores/user.svelte';

	type Props = {
		dialogController: DialogController<{ version: VersionSchema }>;
	};

	let { dialogController }: Props = $props();

	const updateVersionMutation = createUpdateVersionMutation();
	const userState = getUserState();

	const form = $derived(
		createForm({
			schema: versionUpdateSchema,
			defaultValues: {
				name: dialogController.data?.version.name ?? '',
				description: dialogController.data?.version.description ?? '',
				project: dialogController.data?.version.project ?? '',
				createdBy: dialogController.data?.version.createdBy ?? userState.userId ?? ''
			},
			onSubmit: async ({ data, setError }) => {
				if (!dialogController.data?.version.id) {
					setError(m.version_id_required());
					return;
				}

				await $updateVersionMutation.mutateAsync(
					{
						id: dialogController.data.version.id,
						version: {
							...data,
							project: dialogController.data.version.project,
							createdBy: dialogController.data.version.createdBy ?? userState.userId ?? ''
						}
					},
					{
						onSuccess: () => {
							dialogController.close();
							form.reset();
						},
						onError: () => {
							setError(m.failed_to_update_version());
						}
					}
				);
			}
		})
	);
</script>

<Dialog.Root bind:open={dialogController.isOpen}>
	<Dialog.Content class="sm:max-w-[425px]" data-testid="edit-version-dialog">
		<Dialog.Header>
			<Dialog.Title>{m.edit_version()}</Dialog.Title>
			<Dialog.Description>{m.update_version_details()}</Dialog.Description>
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
						{m.updating()}
					{:else}
						{m.update()}
					{/if}
				</Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
