<script lang="ts">
	import type { DialogController } from '$lib/stores/dialog.svelte';

	import { AlertTriangle, LoaderCircle } from '@lucide/svelte';
	import * as Alert from '$lib/components/ui/alert';
	import { Button } from '$lib/components/ui/button';
	import * as Dialog from '$lib/components/ui/dialog';
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import {
		Select,
		SelectContent,
		SelectGroup,
		SelectItem,
		SelectTrigger
	} from '$lib/components/ui/select';
	import { m } from '$lib/paraglide/messages';
	import { createAddProjectMutation } from '$lib/queries/projects';
	import { projectCreateSchema } from '$lib/schemas/project.schema';
	import { getTeamState } from '$lib/stores/team.svelte';
	import { defaults, setError, superForm } from 'sveltekit-superforms';
	import { zod, zodClient } from 'sveltekit-superforms/adapters';

	type Props = {
		dialogController: DialogController<unknown>;
	};

	let { dialogController, ...restProps }: Props = $props();

	const teamState = getTeamState();

	const addMutation = createAddProjectMutation();

	const form = superForm(defaults(zod(projectCreateSchema)), {
		id: 'create-project-form',
		SPA: true,
		validators: zodClient(projectCreateSchema),
		onUpdate: async ({ form }) => {
			if (!form.valid) return;
			if (!teamState.currentTeam) return setError(form, 'Please select a team first.');

			await $addMutation.mutateAsync(
				{
					...form.data,
					team: teamState.currentTeam
				},
				{
					onSuccess: () => {
						dialogController.close();
					},
					onError: () => {
						setError(form, 'Failed to create project. Please try again.');
					}
				}
			);
		}
	});

	const { form: formData, constraints, enhance, submitting, delayed, allErrors, reset } = form;

	let formErrors = $derived(
		$allErrors.filter((error) => error.path === '_errors').flatMap((error) => error.messages)
	);

	const validStatusses = [
		{ value: 'planned', label: 'Planned' },
		{ value: 'active', label: 'Active' },
		{ value: 'completed', label: 'Completed' }
	];
</script>

<Dialog.Root bind:open={dialogController.isOpen} {...restProps}>
	<Dialog.Content class="sm:max-w-[425px]" data-testid="projects-create-dialog">
		<Dialog.Header>
			<Dialog.Title>{m.create_project()}</Dialog.Title>
			<Dialog.Description>{m.add_a_project_to_your_workspace()}</Dialog.Description>
		</Dialog.Header>

		{#if formErrors.length > 0}
			<Alert.Root variant="destructive" class="mt-4">
				<AlertTriangle class="mr-2 h-4 w-4" />
				<Alert.Title>Error</Alert.Title>
				<Alert.Description>{formErrors[0]}</Alert.Description>
			</Alert.Root>
		{/if}

		<form method="POST" class="grid gap-4 py-4" use:enhance>
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
							data-testid="projects-create-dialog-input-name"
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
							disabled={$submitting}
							type="single"
						>
							<SelectTrigger {...props} data-testid="projects-create-dialog-select-trigger-status">
								{$formData.status
									? validStatusses.find((vs) => vs.value === $formData.status)?.label
									: 'Select a status for the project'}
							</SelectTrigger>
							<SelectContent data-testid="projects-create-dialog-select-content-status">
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
				<Form.Button
					type="submit"
					disabled={$submitting}
					data-testid="projects-create-dialog-button-submit"
				>
					{#if $delayed}
						<LoaderCircle class="mr-2 h-4 w-4 animate-spin" />
					{/if}
					{#if $submitting}
						Creating...
					{:else}
						{m.create()}
					{/if}
				</Form.Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
