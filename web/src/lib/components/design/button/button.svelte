<script lang="ts">
	import { Loader } from '$lib/assets/icons';
	import type { Snippet } from 'svelte';
	import PrimaryButton from './primary-button.svelte';
	import SecondaryButton from './secondary-button.svelte';

	interface PropsType {
		loading?: boolean;
		disabled?: boolean;
		type?: 'primary' | 'secondary';
		class?: string;
		left?: Snippet;
		right?: Snippet;
		children?: Snippet;
		onclick?: () => void;
	}

	const {
		loading,
		type = 'primary',
		class: className,
		children,
		onclick
	}: PropsType = $props();
</script>

{#snippet loader(color: string)}
	<span>
		<Loader {color} />
	</span>
{/snippet}


{#if type === 'primary'}
	<PrimaryButton disabled={loading} class={className} {onclick} {loading} >
		{@render children?.()}
	</PrimaryButton>
{/if}

{#if type === 'secondary'}
	<SecondaryButton disabled={loading} class={className} {onclick} {loading}>
		{@render children?.()}
	</SecondaryButton>
{/if}
