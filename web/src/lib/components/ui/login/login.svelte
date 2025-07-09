<script lang="ts">
	import Button from '$lib/components/design/button/button.svelte';

	let email = '';
	let password = '';
	let loading = false;
	let errorMsg = '';

	function handleSubmit(e: Event) {
		e.preventDefault();

		if (!email || !password) {
			errorMsg = 'Please fill in all fields';
			return;
		}
		if (!email.includes('@')) {
			errorMsg = 'Please enter a valid email address';
			return;
		}

		errorMsg = '';
		loading = true;
		setTimeout(() => (loading = false), 1200);
	}
</script>

<div class="card p-8 w-full max-w-md">
	<div class="text-center mb-6">
		<span class="icon w-12 h-12 mx-auto mb-4"></span>
		<h2 class="text-2xl font-bold text-white mb-2">Welcome Back</h2>
		<p class="text-gray-300">Sign in to your account</p>
	</div>

	<form on:submit={handleSubmit} class="space-y-4">
		<div>
			<label class="block text-sm font-medium text-gray-300 mb-2" for="email">Email Address</label>
			<div class="relative group">
				<span
					class="icon absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-gray-400 transition-colors duration-200 group-focus-within:text-primary"
				></span>
				<input
					type="text"
					bind:value={email}
					class="w-full pl-10 pr-4 py-3 bg-surface/50 border border-gray-600 rounded-lg text-white placeholder-gray-400 focus:outline-none focus:border-primary focus:ring-2 focus:ring-primary/50 transition-all duration-300 hover:bg-surface/70"
					placeholder="Enter your email"
					disabled={loading}
					name="email"
				/>
			</div>
		</div>

		<div>
			<label class="block text-sm font-medium text-gray-300 mb-2" for="password">Password</label>
			<div class="relative group">
				<span
					class="icon absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-gray-400 transition-colors duration-200 group-focus-within:text-primary"
				></span>
				<input
					type="password"
					bind:value={password}
					class="w-full pl-10 pr-4 py-3 bg-surface/50 border border-gray-600 rounded-lg text-white placeholder-gray-400 focus:outline-none focus:border-primary focus:ring-2 focus:ring-primary/50 transition-all duration-300 hover:bg-surface/70"
					placeholder="Enter your password"
					disabled={loading}
					name="password"
				/>
			</div>
		</div>

		{#if errorMsg}
			<div
				class="text-red-400 text-sm text-center bg-red-400/10 border border-red-400/20 rounded-lg p-3 animate-scale-in"
			>
				{errorMsg}
			</div>
		{/if}

		<Button variant="primary" type="submit" class="w-full justify-center">Login</Button>
	</form>

	<div class="mt-6 text-center">
		<p class="text-gray-400 text-sm animate-fade-in" style="animation-delay:0.5s">
			Use any email and password to demo the login
		</p>
	</div>
</div>

<style>
	:root {
		--surface: #26204d;
		--primary: #8a3bd5;
	}
	.card {
		border-radius: 1rem;
		background: hsla(248, 41%, 21%, 0.6);
		backdrop-filter: blur(1rem);
		box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.5);
		transition: all 0.3s ease;
	}
</style>
