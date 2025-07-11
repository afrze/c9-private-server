<script lang="ts">
	import Close from '$lib/assets/icons/close.svelte';
	import type { Snippet } from 'svelte';

	interface PropTypes {
		isOpen: boolean;
		title: string;
		onClose: () => void;
		children: Snippet;
		class?: string;
	}

	const {
		isOpen = false,
		title,
		onClose = () => {},
		children,
		class: className
	}: PropTypes = $props();

	const _backdropClass = ['absolute inset-0 bg-black/70 backdrop-blur-sm animate-fade-in'];
	const _cardBaseClass = [
		'card relative p-6 w-full mx-4',
		'transform transition-all duration-300',
		'scale-100 opacity-100 animate-scale-in',
		className
	];
	const _closeButtonClass = [
		'text-gray-400 hover:text-white',
		'transition-all duration-300 p-1',
		'hover:scale-110 hover:rotate-90'
	];
	const _headerClass = ['flex items-center justify-between mb-10'];
</script>

{#if isOpen}
	<div class="fixed inset-0 z-50 flex items-center justify-center">
		<!-- Backdrop -->
		<div
			role="button"
			tabindex="0"
			class={_backdropClass}
			onclick={onClose}
			onkeydown={(e) => {
				if (e.key === 'Escape') onClose();
			}}
		></div>

		<!-- Modal -->
		<div class={_cardBaseClass}>
			<header class={_headerClass}>
				{#if title}
					<h2 class="text-xl font-bold text-white">{title}</h2>
				{/if}

				<button class={_closeButtonClass} onclick={onClose} aria-label="Close modal">
					<Close />
				</button>
			</header>

			{@render children?.()}
		</div>
	</div>
{/if}

<style>
	@keyframes fade-in {
		from {
			opacity: 0;
		}
		to {
			opacity: 1;
		}
	}
	@keyframes scale-in {
		from {
			opacity: 0;
			transform: scale(0.9);
		}
		to {
			opacity: 1;
			transform: scale(1);
		}
	}
	.animate-fade-in {
		animation: fade-in 0.25s ease-out both;
	}
	.animate-scale-in {
		animation: scale-in 0.3s ease-out both;
	}

	/* frosted card style */
	.card {
		border-radius: 1rem;
		background: rgba(38, 32, 77, 0.5); /* bg-surface/50 */
		backdrop-filter: blur(1rem);
		box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.5);
		transition: box-shadow 0.3s ease;
	}
	.card:hover {
		box-shadow: 0 25px 50px -12px rgba(138, 59, 213, 0.2);
	}
</style>
