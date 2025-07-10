<script lang="ts">
	import type { Snippet } from 'svelte';

	interface PropsType {
		children?: Snippet;
		class?: string;
	}

	const { children, class: className }: PropsType = $props();

	let _class = ['card', className];
</script>

<div class={_class}>
	{@render children?.()}
</div>

<style>
	.card {
		border-radius: 1rem;
		background: hsla(248, 41%, 21%, 0.7);
		backdrop-filter: blur(0.5rem);
		box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.7);
		transition: all 0.3s ease;
	}
</style>
