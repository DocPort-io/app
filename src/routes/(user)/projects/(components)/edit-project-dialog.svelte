<script lang="ts">
	import type { DialogController } from '$lib/stores/dialog.svelte';

	import { LoaderCircle } from '@lucide/svelte';
	import { Button } from '$lib/components/ui/button';
	import * as Dialog from '$lib/components/ui/dialog';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import {
		Select,
		SelectContent,
		// SelectGroup,
		SelectItem,
		SelectTrigger
	} from '$lib/components/ui/select';
	import Field from '$lib/form/field.svelte';
	import { createForm } from '$lib/form/form.svelte';
	import { m } from '$lib/paraglide/messages';
	import { createUpdateProjectMutation } from '$lib/queries/projects';
	import { projectUpdateSchema, type ProjectUpdateSchema } from '$lib/schemas/project.schema';

	type Props = {
		dialogController: DialogController<{ id: string; project: ProjectUpdateSchema }>;
	};

	let { dialogController, ...restProps }: Props = $props();

	const updateMutation = createUpdateProjectMutation();

	// const form = superForm(defaults($project, zod(schema)), {
	// 	id: 'edit-project-form',
	// 	SPA: true,
	// 	validators: zodClient(schema),
	// 	invalidateAll: false,
	// 	resetForm: false,
	// 	onUpdate: async ({ form }) => {
	// 		if (!form.valid) return;
	// 		if (!dialogController.data?.id) return;

	// 		await $updateMutation.mutateAsync(
	// 			{
	// 				id: dialogController.data.id,
	// 				project: {
	// 					...form.data,
	// 					team: dialogController.data.project.team
	// 				}
	// 			},
	// 			{
	// 				onSuccess: () => {
	// 					dialogController.close();
	// 				},
	// 				onError: () => {
	// 					setError(form, m.failed_to_update_project());
	// 				}
	// 			}
	// 		);
	// 	}
	// });

	const validStatusses = [
		{ value: 'planned', label: m.planned() },
		{ value: 'active', label: m.active() },
		{ value: 'completed', label: m.completed() }
	];

	const form = $derived(
		createForm({
			schema: projectUpdateSchema,
			defaultValues: {
				...dialogController.data?.project
			},
			onSubmit: async ({ data }) => {
				console.log('going to submit');
				console.log(data);
				console.log('submitted');

				await new Promise((resolve) => setTimeout(resolve, 1000));

				await $updateMutation.mutateAsync(
					{
						id: dialogController.data!.id,
						project: data
					},
					{
						onSuccess: () => {
							dialogController.close();
						},
						onError: () => {
							// setError(form, m.failed_to_update_project());
						}
					}
				);
			}
		})
	);

	// $inspect(dialogController.data?.project);
	$inspect(form.fields.name.state);
</script>

<Dialog.Root bind:open={dialogController.isOpen} {...restProps}>
	<Dialog.Content class="sm:max-w-[425px]">
		<Dialog.Header>
			<Dialog.Title>{m.edit_project()}</Dialog.Title>
			<Dialog.Description>{m.update_the_project_details_below()}</Dialog.Description>
		</Dialog.Header>
		<form
			class="grid gap-4 py-4"
			onsubmit={(e) => {
				e.preventDefault();
				e.stopPropagation();
				form.handleSubmit();
			}}
		>
			<Field {form} name="name">
				{#snippet children(field)}
					<Label for={field.name}>{m.name()}</Label>
					<Input
						id={field.name}
						name={field.name}
						value={field.state.value}
						onblur={field.handleBlur}
						placeholder={m.my_awesome_project()}
						oninput={(e) => field.handleChange((e.target as HTMLInputElement).value)}
						disabled={form.state.isSubmitting}
					/>
					<!-- <p class="text-destructive text-sm font-medium">{field.state.meta.errors[0]}</p> -->
				{/snippet}
			</Field>
			<Field {form} name="status">
				{#snippet children(field)}
					<Label for={field.name}>{m.status()}</Label>
					<!-- <Input
						id={field.name}
						name={field.name}
						value={field.state.value}
						onblur={field.handleBlur}
						placeholder={m.my_awesome_project()}
						oninput={(e) => field.handleChange((e.target as HTMLSelectElement).value)}
						disabled={tForm.state.isSubmitting}
					/> -->
					<Select
						name={field.name}
						type="single"
						value={field.state.value}
						onValueChange={(value) =>
							field.handleChange(value as 'active' | 'completed' | 'planned')}
						disabled={form.state.isSubmitting}
					>
						<SelectTrigger>
							{field.state.value
								? validStatusses.find((vs) => vs.value === field.state.value)?.label
								: m.select_a_status_for_the_project_placeholder()}
						</SelectTrigger>
						<SelectContent>
							{#each validStatusses as status (status.value)}
								<SelectItem value={status.value} label={status.label} />
							{/each}
						</SelectContent>
					</Select>
					<!-- <p class="text-destructive text-sm font-medium">{field.state.meta.errors[0]}</p> -->
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
				<Button type="submit" disabled={form.state.isSubmitting}>
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
