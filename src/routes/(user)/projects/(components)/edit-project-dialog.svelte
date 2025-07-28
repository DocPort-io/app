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
			onSubmit: async ({ data, state, setError }) => {
				// Early return if form is not valid
				if (!state.isValid) {
					console.error('Form submission blocked: validation failed');
					return;
				}

				console.log('Submitting project update:', data);

				// Simulate network delay (remove in production)
				await new Promise((resolve) => setTimeout(resolve, 1000));

				// Submit the mutation
				await $updateMutation.mutateAsync(
					{
						id: dialogController.data!.id,
						project: data
					},
					{
						onSuccess: () => {
							console.log('Project updated successfully');
							setError('This is a test');
							// dialogController.close();
						},
						onError: (error) => {
							console.error('Failed to update project:', error);
							// TODO: Display user-friendly error message
							// setError(form, m.failed_to_update_project());
						}
					}
				);
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
		<!-- <form method="POST" class="grid gap-4 py-4" use:enhance>
			<Form.Field {form} name="name">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>{m.name()}</Form.Label>
						<Input
							{...props}
							{...$constraints.name}
							bind:value={$formData.name}
							placeholder={m.my_awesome_project()}
							disabled={$submitting}
						/>
					{/snippet}
				</Form.Control>
				<Form.Description>{m.enter_a_meaningful_name_for_your_project()}</Form.Description>
				<Form.FieldErrors />
			</Form.Field>

			<Form.Field {form} name="status">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>{m.status()}</Form.Label>
						<Select
							name={props.name}
							{...$constraints.status}
							bind:value={$formData.status}
							type="single"
						>
							<SelectTrigger {...props}>
								{$formData.status
									? validStatusses.find((vs) => vs.value === $formData.status)?.label
									: m.select_a_status_for_the_project_placeholder()}
							</SelectTrigger>
							<SelectContent>
								<SelectGroup>
									{#each validStatusses as status (status.value)}
										<SelectItem value={status.value} label={status.label} />
									{/each}
								</SelectGroup>
							</SelectContent>
						</Select>
					{/snippet}
				</Form.Control>
				<Form.Description>
					{m.select_a_status_for_the_project()}
				</Form.Description>
				<Form.FieldErrors />
			</Form.Field>

			<Dialog.Footer>
				<Button variant="outline" onclick={() => dialogController.close()} disabled={$submitting}>
					{m.cancel()}
				</Button>
				<Form.Button type="submit" disabled={$submitting}>
					{#if $delayed}
						<LoaderCircle class="mr-2 h-4 w-4 animate-spin" />
					{/if}
					{#if $submitting}
						{m.saving()}
					{:else}
						{m.save()}
					{/if}
				</Form.Button>
			</Dialog.Footer>
		</form> -->
	</Dialog.Content>
</Dialog.Root>
