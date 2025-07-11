<script lang="ts">
	import { Lock, User } from '@lucide/svelte';
	import { Button } from '$lib/components/design/button/';
	import { Card } from '$lib/components/design/card';
	import { Form, Label, Input } from '$lib/components/design/form';
	import Flame from '../flame/flame.svelte';

	let username = '';
	let password = '';
	let loading = false;
	let errorMsg = '';

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

	function handleSubmit(e: Event) {
		e.preventDefault();

		if (!username || !password) {
			errorMsg = 'Please fill in all fields';
			return;
		}

		errorMsg = '';
		loading = true;
		setTimeout(() => (loading = false), 1200);
		console.log('Submitted', username, password);
	}
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

<aside class="w-[420px]">
	<Card class="p-8 flex justify-between items-center gap-5 mb-5">
		<Flame />
		<div class="text-center w-full">
			<span class="icon w-12 h-12 mx-auto mb-4"></span>
			<h2 class="text-2xl font-bold text-gold mb-2">Welcome to C9 Dark</h2>
			<p class="text-white">Sign in to your account</p>
		</div>
	</Card>
	<Card class="py-12 px-10 relative">
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

			<Button variant="primary" type="submit" class="w-full justify-center mt-4">Login</Button>

			{#if errorMsg}
				<div class={_errorMessageClass}>
					{errorMsg}
				</div>
			{/if}
		</Form>
	</Card>
</aside>
