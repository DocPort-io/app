<script lang="ts">
	import type { DialogController } from '$lib/stores/dialog.svelte';

	import { AlertTriangle, LoaderCircle, Upload as UploadIcon } from '@lucide/svelte';
	import * as Alert from '$lib/components/ui/alert';
	import { Button } from '$lib/components/ui/button';
	import * as Dialog from '$lib/components/ui/dialog';
	import * as Form from '$lib/components/ui/form';
	import { m } from '$lib/paraglide/messages';
	import { getLocale } from '$lib/paraglide/runtime';
	import { createUploadFileMutation } from '$lib/queries/files';
	import { fileUploadSchema } from '$lib/schemas/file.schema';
	import prettyBytes from 'pretty-bytes';
	import { defaults, setError, superForm } from 'sveltekit-superforms';
	import { zod, zodClient } from 'sveltekit-superforms/adapters';

	type Props = {
		dialogController: DialogController<{ versionId: string }>;
	};

	let { dialogController }: Props = $props();

	const uploadMutation = createUploadFileMutation();

	const form = $derived(
		superForm(
			defaults({ versions: [dialogController.data?.versionId ?? ''] }, zod(fileUploadSchema)),
			{
				id: 'upload-file-form',
				SPA: true,
				validators: zodClient(fileUploadSchema),
				onUpdate: async ({ form }) => {
					if (!form.valid) return;
					if (!dialogController.data?.versionId) return setError(form, m.version_id_required());

					await $uploadMutation.mutateAsync(
						{
							...form.data
						},
						{
							onSuccess: () => {
								dialogController.close();
								reset();
							},
							onError: () => {
								setError(form, m.failed_to_upload_file());
							}
						}
					);
				}
			}
		)
	);

	const { form: formData, enhance, submitting, delayed, allErrors, reset } = $derived(form);

	let formErrors = $derived(
		$allErrors.filter((error) => error.path === '_errors').flatMap((error) => error.messages)
	);

	let fileInput: HTMLInputElement;
	let selectedFiles: FileList | null = $state(null);
	let isDragOver = $state(false);

	const handleFileChange = (event: Event) => {
		const target = event.target as HTMLInputElement;
		selectedFiles = target.files;
		if (selectedFiles && selectedFiles.length > 0) {
			$formData.file = selectedFiles[0];
		}
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
			selectedFiles = files;
			$formData.file = files[0];
			// Update the file input
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

		{#if formErrors.length > 0}
			<Alert.Root variant="destructive" class="mt-4">
				<AlertTriangle class="mr-2 h-4 w-4" />
				<Alert.Title>{m.error()}</Alert.Title>
				<Alert.Description>{formErrors[0]}</Alert.Description>
			</Alert.Root>
		{/if}

		<form method="POST" enctype="multipart/form-data" class="grid gap-4 py-4" use:enhance>
			<Form.Field {form} name="file">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>{m.select_files()}</Form.Label>
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
							{#if selectedFiles && selectedFiles.length > 0}
								<div class="mt-4 rounded-md border border-green-200 bg-green-50 p-3">
									<p class="text-sm font-medium text-green-700">
										{selectedFiles[0].name}
									</p>
									<p class="text-xs text-green-600">
										{prettyBytes(selectedFiles[0].size, { locale: getLocale() })} • {selectedFiles[0]
											.type || m.unknown_type()}
									</p>
								</div>
							{/if}
						</div>
						<input
							bind:this={fileInput}
							{...props}
							type="file"
							class="hidden"
							onchange={handleFileChange}
							disabled={$submitting}
						/>
					{/snippet}
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>

			<Dialog.Footer>
				<Button
					variant="outline"
					onclick={() => {
						reset();
						dialogController.close();
					}}
					disabled={$submitting}
				>
					{m.cancel()}
				</Button>
				<Form.Button type="submit" disabled={$submitting || !selectedFiles}>
					{#if $delayed}
						<LoaderCircle class="mr-2 h-4 w-4 animate-spin" />
					{/if}
					{#if $submitting}
						{m.uploading()}
					{:else}
						{m.upload()}
					{/if}
				</Form.Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
