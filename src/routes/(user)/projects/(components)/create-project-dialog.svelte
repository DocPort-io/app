<script lang="ts">
	import type { DialogController } from '$lib/stores/dialog.svelte';

	import { LoaderCircle } from '@lucide/svelte';
	import { Button } from '$lib/components/ui/button';
	import * as Dialog from '$lib/components/ui/dialog';
	import { Input } from '$lib/components/ui/input';
	import {
		Select,
		SelectContent,
		SelectGroup,
		SelectItem,
		SelectTrigger
	} from '$lib/components/ui/select';
	import FieldErrors from '$lib/form/field-errors.svelte';
	import FieldLabel from '$lib/form/field-label.svelte';
	import Field from '$lib/form/field.svelte';
	import FormErrors from '$lib/form/form-errors.svelte';
	import { createForm } from '$lib/form/form.svelte';
	import { m } from '$lib/paraglide/messages';
	import { createAddProjectMutation } from '$lib/queries/projects';
	import { projectSchema } from '$lib/schemas/project.schema';
	import { getTeamState } from '$lib/stores/team.svelte';

	type Props = {
		dialogController: DialogController<unknown>;
	};

	let { dialogController, ...restProps }: Props = $props();

	const teamState = getTeamState();

	const addMutation = createAddProjectMutation();

	const schema = projectSchema.pick({
		name: true,
		status: true
	});

	const form = $derived(
		createForm({
			schema,
			defaultValues: {
				status: 'planned'
			},
			onSubmit: async ({ data, setError }) => {
				if (!teamState.currentTeam) {
					setError('Please select a team first.');
					return;
				}

				try {
					await $addMutation.mutateAsync({
						...data,
						team: teamState.currentTeam
					});

					dialogController.close();
				} catch {
					setError('Failed to create project. Please try again.');
				}
			}
		})
	);

	const PROJECT_STATUSES = [
		{ value: 'planned', label: m.planned() },
		{ value: 'active', label: m.active() },
		{ value: 'completed', label: m.completed() }
	] as const;
</script>

<Dialog.Root bind:open={dialogController.isOpen} {...restProps}>
	<Dialog.Content class="sm:max-w-[425px]" data-testid="projects-create-dialog">
		<Dialog.Header>
			<Dialog.Title>{m.create_project()}</Dialog.Title>
			<Dialog.Description>{m.add_a_project_to_your_workspace()}</Dialog.Description>
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
						data-testid="projects-create-dialog-input-name"
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
						<SelectTrigger data-testid="projects-create-dialog-select-trigger-status">
							{state.value
								? PROJECT_STATUSES.find((status) => status.value === state.value)?.label
								: m.select_a_status_for_the_project_placeholder()}
						</SelectTrigger>
						<SelectContent data-testid="projects-create-dialog-select-content-status">
							<SelectGroup>
								{#each PROJECT_STATUSES as status (status.value)}
									<SelectItem value={status.value} label={status.label} />
								{/each}
							</SelectGroup>
						</SelectContent>
					</Select>
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
				<Button
					type="submit"
					disabled={!form.state.isSubmittable}
					data-testid="projects-create-dialog-button-submit"
				>
					{#if form.state.isSubmitting}
						<LoaderCircle class="mr-2 h-4 w-4 animate-spin" />
						{m.creating()}
					{:else}
						{m.create()}
					{/if}
				</Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
