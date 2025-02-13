<script lang="ts">
    import { Input } from "$lib/components/ui/input";
	import { Label } from "$lib/components/ui/label";

	let message = $state('');
</script>

<!-- <h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p> -->

<Label for="message">Message</Label>
<Input type="text" placeholder="Enter some text" class="max-w-xs" id="message" />

<sl-alert open>
	<sl-icon slot="icon" name="info-circle"></sl-icon>
	{message}
</sl-alert>

<sl-qr-code label="Scan this code!" value={message}></sl-qr-code>
