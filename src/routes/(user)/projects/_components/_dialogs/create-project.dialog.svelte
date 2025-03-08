<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import * as Dialog from '$lib/components/ui/dialog';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import { m } from '$lib/paraglide/messages';

	type Props = {
		open?: boolean;
		onSubmit?: () => Promise<void>;
	};

	let { open = $bindable(false), onSubmit, ...restProps }: Props = $props();

	let name = $state('');
	let errors = $state<Record<string, string>>({});
	let isSubmitting = $state(false);

	const handleSubmit = async (event: SubmitEvent) => {
		event.preventDefault();
	};
</script>

<Dialog.Root bind:open {...restProps}>
	<Dialog.Content class="sm:max-w-[425px]">
		<Dialog.Header>
			<Dialog.Title>{m.weak_weak_bulldog_assure()}</Dialog.Title>
			<Dialog.Description>{m.proof_noisy_monkey_type()}</Dialog.Description>
		</Dialog.Header>
		<form onsubmit={handleSubmit} class="grid gap-4 py-4">
			<div class="grid gap-2">
				<Label for="name">{m.royal_major_impala_charm()}</Label>
				<Input
					id="name"
					bind:value={name}
					placeholder={m.alert_nimble_pug_play()}
					aria-invalid={errors.name ? 'true' : undefined}
					aria-describedby={errors.name ? 'name-error' : undefined}
				/>
				{#if errors.name}
					<div class="text-destructive col-span-3 col-start-2 text-sm" id="name-error">
						{errors.name}
					</div>
				{/if}
			</div>
			<Dialog.Footer>
				<Button type="button" variant="outline" on:click={() => (open = false)}
					>{m.red_same_flea_clip()}
				</Button>
				<Button type="submit" disabled={isSubmitting}>
					{isSubmitting ? m.fuzzy_inner_fireant_feast() : m.hour_swift_crab_breathe()}
				</Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
