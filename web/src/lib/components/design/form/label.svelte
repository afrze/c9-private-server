<script lang="ts">
	import type { Snippet } from 'svelte';

	interface PropsType {
		children?: Snippet;
		class?: string;
		for?: string;
	}

	const { children, class: className, for: htmlFor }: PropsType = $props();

	const _class = ['block text-sm font-medium text-gray-300 mb-2', className];
</script>

<label class={_class} for={htmlFor}>{@render children?.()}</label>
