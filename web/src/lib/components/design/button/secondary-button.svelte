<script lang="ts">
	import type { Snippet } from 'svelte';

	interface PropsType {
		children?: Snippet;
		class?: string;
		disabled?: boolean;
		left?: Snippet;
		loading?: boolean;
		success?: boolean;
		onclick?: () => void;
		right?: Snippet;
	}

	const { children, class: className, disabled, left, onclick, right }: PropsType = $props();

	const _class = [
		'rounded-lg font-semibold text-center transition-all duration-300 cursor-pointer',
		'transform hover:scale-105 hover:-translate-y-1 active:scale-95 active:translate-y-0',
		'disabled:opacity-50 disabled:cursor-not-allowed disabled:transform-none',
		'relative overflow-hidden group',
		'bg-secondary/80 hover:bg-surface text-white border border-primary/30',
		'hover:border-primary/70 hover:shadow-lg hover:shadow-primary/20',
		className
	];
</script>

<button class={_class} {onclick} {disabled}>
	<span
		class="absolute inset-0 bg-gradient-to-r from-transparent via-white/10 to-transparent -translate-x-full group-hover:translate-x-full transition-transform duration-700"
	></span>
	{@render left?.()}
	{@render children?.()}
	{@render right?.()}
</button>
