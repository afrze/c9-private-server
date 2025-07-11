<script lang="ts">
	import { Button } from '$lib/components/design/button';
	import { Form, Label } from '$lib/components/design/form';
	import Input from '$lib/components/design/form/input.svelte';
	import { User, Lock } from '@lucide/svelte';

	let username = '';
	let password = '';
	let confirmPassword = '';
	let loading = false;
	let errorMsg = '';

	function handleSubmit(e: Event) {
		e.preventDefault();

		if (!user || !password || !user) {
			errorMsg = 'Please fill in all fields';
			return;
		}

		if (password !== confirmPassword) {
			errorMsg = 'Password do not match';
			return;
		}

		errorMsg = '';
		loading = true;
		setTimeout(() => (loading = false), 1200);
		console.log('Submitted', username, password);
	}

	const _userIconClass = [
		'icon absolute left-3 top-1/2 -translate-y-3/5 w-5 h-5',
		'text-gray-400 transition-colors duration-200 group-focus-within:text-primary'
	];

	const _passwordIconClass = [
		'icon absolute left-3 top-1/2 -translate-y-3/5 w-5 h-5',
		'text-gray-400 transition-colors duration-200 group-focus-within:text-primary'
	];

	const _errorMessageClass = [
		'rounded-lg p-3 animate-scale-in absolute top-0 left-0 w-full',
		'text-red-400 text-sm text-center bg-red-400/10 border border-red-400/20'
	];
</script>

{#snippet user()}
	<span class={_userIconClass}>
		<User width={22} height={22} />
	</span>
{/snippet}

{#snippet lock()}
	<span class={_passwordIconClass}>
		<Lock width={22} height={22} />
	</span>
{/snippet}

<Form onsubmit={handleSubmit}>
	<div class="mb-8">
		<Label for="email">Username</Label>
		<Input
			{loading}
			left={user}
			bind:value={username}
			type={'text'}
			id={'email'}
			placeholder={'Enter your username'}
		/>
	</div>

	<div class="mb-8">
		<Label for="password">Password</Label>
		<Input
			{loading}
			left={lock}
			bind:value={password}
			type={'password'}
			id={'password'}
			placeholder={'Enter your password'}
		/>
	</div>

	<div class="mb-8">
		<Label for="confirm-password">Confirm Password</Label>
		<Input
			{loading}
			left={lock}
			bind:value={confirmPassword}
			type={'password'}
			id={'confirm-password'}
			placeholder={'Confirm your password'}
		/>
	</div>

	<Button variant="primary" type="submit" class="w-full justify-center mt-4"
		>Register your account</Button
	>

	{#if errorMsg}
		<div class={_errorMessageClass}>
			{errorMsg}
		</div>
	{/if}
</Form>