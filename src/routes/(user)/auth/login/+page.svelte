<script lang="ts">
	import * as Card from '$lib/components/ui/card';
	import { m } from '$lib/paraglide/messages';

	import ExternalAuthProviders from './_components/external-auth-providers.svelte';
	import SignInForm from './_components/sign-in-form.svelte';
</script>

<div class="flex min-h-screen items-center justify-center bg-slate-50 p-4">
	<Card.Root class="w-full max-w-md">
		<Card.Header>
			<Card.Title class="text-center text-2xl font-bold">{m.welcome_back()}</Card.Title>
			<Card.Description class="text-center">{m.sign_in_to_your_account()}</Card.Description>
		</Card.Header>
		<Card.Content>
			<SignInForm />
			<ExternalAuthProviders />
		</Card.Content>
		<Card.Footer class="flex justify-center">
			<p class="text-muted-foreground text-sm">
				{m.dont_have_an_account()}
				<a href="/register" class="text-primary hover:underline">{m.create_an_account()}</a>
			</p>
		</Card.Footer>
	</Card.Root>
</div>
