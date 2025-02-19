<script lang="ts">
	import Terminal from 'lucide-svelte/icons/terminal';
	import * as Alert from '$lib/components/ui/alert';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';

	let message = $state('');
</script>

<Label for="message">Message</Label>
<Input
	type="text"
	placeholder="Enter some text"
	class="max-w-xs"
	id="message"
	bind:value={message}
/>

<Alert.Root>
	<Terminal class="h-4 w-4" />
	<Alert.Title>Heads up!</Alert.Title>
	<Alert.Description>
		This is a warning message. Please be careful.
		{message}
	</Alert.Description>
</Alert.Root>
