<script lang="ts">
	import type { Snippet } from 'svelte';
	import type { FullAutoFill } from 'svelte/elements';

	interface PropsType {
		loading?: boolean;
		value?: string;
		left?: Snippet | any;
		class?: string;
		type?: 'text' | 'number' | 'email' | 'password';
		id?: string;
		placeholder?: string;
		autocomplete?: FullAutoFill;
	}
	let {
		loading,
		value = $bindable(),
		left,
		class: className,
		type = 'text',
		id,
		autocomplete = 'off',
		placeholder
	}: PropsType = $props();

	const _class = [
		'w-full pl-12 pr-4 py-3 bg-surface/50 border',
		'border-gray-600 rounded-lg text-white placeholder-gray-400 focus:outline-none',
		'focus:border-primary focus:ring-2 focus:ring-primary/50 transition-all duration-300 hover:bg-surface/70',
		className
	];
</script>

<div class="relative group">
	{@render left?.()}
	<input {id} {type} bind:value class={_class} {placeholder} disabled={loading} {autocomplete} />
</div>
