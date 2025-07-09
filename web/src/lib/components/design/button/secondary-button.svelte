<script lang="ts">
	/* Runes-style props */
	import type { Snippet } from 'svelte';

	interface PropsType {
		children?: Snippet;
		class?: string;
		disabled?: boolean;
		loading?: boolean;
		success?: boolean;
		left?: Snippet;
		right?: Snippet;
		type?: 'submit' | 'button'
		onclick?: () => void;
	}

	const {
		children,
		class: className = '',
		disabled = false,
		onclick,
		left,
		right,
		type = 'button'
	}: PropsType = $props();
</script>

<button class={`primary-btn ${className}`} {onclick} {disabled} {type} aria-disabled={disabled}>
	<span class="beam"></span>
	{@render left?.()}
	{@render children?.()}
	{@render right?.()}
</button>

<style>
	:root {
		--secondary: #26204d;
	}
	.primary-btn {
		display: inline-flex;
		align-items: center;
		gap: .75rem;
		padding: 0.75rem 1.5rem;
		border-radius: 0.5rem;
		font-weight: 600;
		color: #fff;
		background: var(--secondary);

		animation: primary-glow 2s ease-in-out infinite alternate;
		transition: transform 0.3s ease;
		position: relative;
		overflow: hidden;
		cursor: pointer;
		user-select: none;
	}

	.primary-btn:hover {
		transform: translateY(-2px) scale(1.05);
	}
	.primary-btn:active {
		transform: translateY(0) scale(0.97);
	}

	.primary-btn .beam {
		content: '';
		position: absolute;
		inset: 0;
		background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.25), transparent);
		transform: translateX(-100%);
		transition: transform 0.7s ease;
		pointer-events: none;
	}
	.primary-btn:hover .beam {
		transform: translateX(100%);
	}

	.primary-btn[disabled],
	.primary-btn[aria-disabled='true'] {
		opacity: 0.5;
		cursor: not-allowed;
		animation: none;
		transform: none;
	}

	@keyframes primary-glow {
		0%,
		100% {
			box-shadow:
				0 0 6px rgba(138, 59, 213, 0.6),
				0 0 12px rgba(138, 59, 213, 0.4);
		}
		50% {
			box-shadow:
				0 0 10px rgba(138, 59, 213, 0.8),
				0 0 24px rgba(138, 59, 213, 0.6);
		}
	}
</style>
