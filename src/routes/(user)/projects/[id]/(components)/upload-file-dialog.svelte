<script lang="ts">
	import type { DialogController } from '$lib/stores/dialog.svelte';

	import { LoaderCircle, Upload as UploadIcon } from '@lucide/svelte';
	import { Button } from '$lib/components/ui/button';
	import * as Dialog from '$lib/components/ui/dialog';
	import FieldErrors from '$lib/form/field-errors.svelte';
	import FieldLabel from '$lib/form/field-label.svelte';
	import Field from '$lib/form/field.svelte';
	import FormErrors from '$lib/form/form-errors.svelte';
	import { createForm } from '$lib/form/form.svelte';
	import { m } from '$lib/paraglide/messages';
	import { getLocale } from '$lib/paraglide/runtime';
	import { createUploadFileMutation } from '$lib/queries/files';
	import { fileUploadSchema } from '$lib/schemas/file.schema';
	import prettyBytes from 'pretty-bytes';

	type Props = {
		dialogController: DialogController<{ versionId: string }>;
	};

	let { dialogController }: Props = $props();

	const uploadMutation = createUploadFileMutation();

	const form = $derived(
		createForm({
			schema: fileUploadSchema,
			defaultValues: {
				versions: [dialogController.data?.versionId ?? '']
			},
			onSubmit: async ({ data, setError }) => {
				if (!dialogController.data?.versionId) {
					setError(m.version_id_required());
					return;
				}

				await $uploadMutation.mutateAsync(
					{
						...data
					},
					{
						onSuccess: () => {
							dialogController.close();
							resetForm();
						},
						onError: () => {
							setError(m.failed_to_upload_file());
						}
					}
				);
			}
		})
	);

	let fileInput: HTMLInputElement;
	let isDragOver = $state(false);

	// Derived state for selected file info
	const selectedFile = $derived(form.fields.file.state.value);

	// Centralized file setting function
	const setFile = (file: File | null) => {
		form.setFieldValue('file', file);
	};

	// Centralized reset function
	const resetForm = () => {
		form.reset();
		if (fileInput) {
			fileInput.value = '';
		}
	};

	const handleFileChange = (event: Event) => {
		const target = event.target as HTMLInputElement;
		const files = target.files;
		setFile(files && files.length > 0 ? files[0] : null);
	};

	const handleDragOver = (event: DragEvent) => {
		event.preventDefault();
		event.stopPropagation();
		isDragOver = true;
	};

	const handleDragLeave = (event: DragEvent) => {
		event.preventDefault();
		event.stopPropagation();
		isDragOver = false;
	};

	const handleDrop = (event: DragEvent) => {
		event.preventDefault();
		event.stopPropagation();
		isDragOver = false;

		const files = event.dataTransfer?.files;
		if (files && files.length > 0) {
			setFile(files[0]);
			// Update the file input to stay in sync
			if (fileInput) {
				fileInput.files = files;
			}
		}
	};
</script>

<Dialog.Root bind:open={dialogController.isOpen}>
	<Dialog.Content class="sm:max-w-[425px]" data-testid="upload-file-dialog">
		<Dialog.Header>
			<Dialog.Title>{m.upload_file()}</Dialog.Title>
			<Dialog.Description>{m.upload_a_new_file()}</Dialog.Description>
		</Dialog.Header>

		<FormErrors {form} />

		<form {...form.props} method="POST" enctype="multipart/form-data" class="grid gap-4 py-4">
			<Field {form} name="file">
				{#snippet children({ props })}
					<FieldLabel>{m.select_files()}</FieldLabel>
					<div
						class="rounded-lg border-2 border-dashed p-6 text-center transition-colors {isDragOver
							? 'border-primary bg-primary/5'
							: 'border-gray-300 hover:border-gray-400'}"
						ondragover={handleDragOver}
						ondragleave={handleDragLeave}
						ondrop={handleDrop}
						role="button"
						tabindex="0"
						onclick={() => fileInput?.click()}
						onkeydown={(e) => e.key === 'Enter' && fileInput?.click()}
					>
						<UploadIcon class="mx-auto mb-4 h-12 w-12 text-gray-400" />
						<p class="mb-2 text-sm text-gray-600">
							{m.drag_and_drop_files_here()}
						</p>
						<p class="text-xs text-gray-500">{m.maximum_file_size()}</p>
						{#if selectedFile}
							<div class="mt-4 rounded-md border border-green-200 bg-green-50 p-3">
								<p class="text-sm font-medium text-green-700">
									{selectedFile.name}
								</p>
								<p class="text-xs text-green-600">
									{prettyBytes(selectedFile.size, { locale: getLocale() })} • {selectedFile.type ||
										m.unknown_type()}
								</p>
							</div>
						{/if}
					</div>
					<input
						{...props}
						bind:this={fileInput}
						type="file"
						class="hidden"
						onchange={handleFileChange}
						disabled={form.state.isSubmitting}
					/>
					<FieldErrors />
				{/snippet}
			</Field>

			<Dialog.Footer>
				<Button
					variant="outline"
					onclick={() => {
						resetForm();
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
						{m.uploading()}
					{:else}
						{m.upload()}
					{/if}
				</Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
