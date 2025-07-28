<script lang="ts">
	import type { DialogController } from '$lib/stores/dialog.svelte';

	import { LoaderCircle } from '@lucide/svelte';
	import { Button } from '$lib/components/ui/button';
	import * as Dialog from '$lib/components/ui/dialog';
	import { Input } from '$lib/components/ui/input';
	import { Select, SelectContent, SelectItem, SelectTrigger } from '$lib/components/ui/select';
	import FieldErrors from '$lib/form/field-errors.svelte';
	import FieldLabel from '$lib/form/field-label.svelte';
	import Field from '$lib/form/field.svelte';
	import FormErrors from '$lib/form/form-errors.svelte';
	import { createForm } from '$lib/form/form.svelte';
	import { m } from '$lib/paraglide/messages';
	import { createUpdateProjectMutation } from '$lib/queries/projects';
	import { projectUpdateSchema, type ProjectUpdateSchema } from '$lib/schemas/project.schema';

	type Props = {
		dialogController: DialogController<{ id: string; project: ProjectUpdateSchema }>;
	};

	let { dialogController, ...restProps }: Props = $props();

	const updateMutation = createUpdateProjectMutation();

	/**
	 * Available project status options for the select dropdown
	 */
	const PROJECT_STATUSES = [
		{ value: 'planned', label: m.planned() },
		{ value: 'active', label: m.active() },
		{ value: 'completed', label: m.completed() }
	] as const;

	/**
	 * Create form instance with validation schema and submission handler
	 */
	const form = $derived(
		createForm({
			schema: projectUpdateSchema,
			defaultValues: {
				...dialogController.data?.project
			},
			onSubmit: async ({ data, setError }) => {
				try {
					await $updateMutation.mutateAsync({
						id: dialogController.data!.id,
						project: data
					});
					
					dialogController.close();
				} catch {
					setError('Failed to update project. Please try again.');
				}
			}
		})
	);
</script>

<Dialog.Root bind:open={dialogController.isOpen} {...restProps}>
	<Dialog.Content class="sm:max-w-[425px]">
		<Dialog.Header>
			<Dialog.Title>{m.edit_project()}</Dialog.Title>
			<Dialog.Description>{m.update_the_project_details_below()}</Dialog.Description>
		</Dialog.Header>
		<form class="grid gap-4 py-4" {...form.props}>
			<FormErrors {form} />
			<Field {form} name="name">
				{#snippet children({ props, state })}
					<FieldLabel>{m.name()}</FieldLabel>
					<Input
						{...props}
						bind:value={state.value}
						placeholder={m.my_awesome_project()}
						disabled={form.state.isSubmitting}
					/>
					<FieldErrors />
				{/snippet}
			</Field>
			<Field {form} name="status">
				{#snippet children({ props, state })}
					<FieldLabel>{m.status()}</FieldLabel>
					<Select
						{...props}
						type="single"
						bind:value={state.value}
						disabled={form.state.isSubmitting}
					>
						<SelectTrigger>
							{state.value
								? PROJECT_STATUSES.find((status) => status.value === state.value)?.label
								: m.select_a_status_for_the_project_placeholder()}
						</SelectTrigger>
						<SelectContent>
							{#each PROJECT_STATUSES as status (status.value)}
								<SelectItem value={status.value} label={status.label} />
							{/each}
						</SelectContent>
					</Select>
					<FieldErrors />
				{/snippet}
			</Field>

			<Dialog.Footer>
				<Button
					variant="outline"
					onclick={() => dialogController.close()}
					disabled={form.state.isSubmitting}
				>
					{m.cancel()}
				</Button>
				<Button type="submit" disabled={!form.state.isSubmittable}>
					{#if form.state.isSubmitting}
						<LoaderCircle class="h-4 w-4 animate-spin" />{m.saving()}
					{:else}
						{m.save()}
					{/if}
				</Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
