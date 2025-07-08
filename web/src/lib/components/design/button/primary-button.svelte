<script lang="ts">
	import type { Snippet } from 'svelte';

	interface PropsType {
		children?: Snippet;
		class?: string;
		disabled?: boolean;
		loading?: boolean;
		success?: boolean;
		onclick?: () => void;
	}

	const { children, class: className, disabled, onclick }: PropsType = $props();

	const _class = [
		'rounded-lg font-semibold text-center transition-all duration-300 cursor-pointer',
		'transform hover:scale-105 hover:-translate-y-1 active:scale-95 active:translate-y-0',
		'disabled:opacity-50 disabled:cursor-not-allowed disabled:transform-none',
		'relative overflow-hidden group',
		'bg-primary hover:bg-primary/90 text-white shadow-lg shadow-primary/25',
		'hover:shadow-xl hover:shadow-primary/40 animate-glow',
		className
	];
</script>

<button class={_class} {onclick} {disabled}>
	<span
		class="absolute inset-0 bg-gradient-to-r from-transparent via-white/10 to-transparent -translate-x-full group-hover:translate-x-full transition-transform duration-700"
	></span>
	{@render children?.()}
</button>
