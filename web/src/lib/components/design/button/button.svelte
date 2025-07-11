<script lang="ts">
	import type { Snippet } from 'svelte';
	import PrimaryButton from './primary-button.svelte';
	import SecondaryButton from './secondary-button.svelte';

	interface PropsType {
		loading?: boolean;
		disabled?: boolean;
		variant?: 'primary' | 'secondary';
		type?: 'button' | 'submit';
		class?: string;
		left?: Snippet;
		right?: Snippet;
		children?: Snippet;
		onclick?: () => void;
	}

	const {
		loading,
		type = 'button',
		variant = 'primary',
		class: className,
		children,
		onclick,
		left,
		right
	}: PropsType = $props();
</script>

{#if variant === 'primary'}
	<PrimaryButton disabled={loading} class={className} {onclick} {loading} {left} {right} {type}>
		{@render children?.()}
	</PrimaryButton>
{/if}

{#if variant === 'secondary'}
	<SecondaryButton disabled={loading} class={className} {onclick} {loading} {type}>
		{@render children?.()}
	</SecondaryButton>
{/if}
