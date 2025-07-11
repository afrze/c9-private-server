<script>
	import Mega from '$lib/assets/icons/mega.svelte';

	const { href } = $props();
</script>

<a {href} target="_blank" class="mega-btn btn-base">
	<Mega />
	Mega Cloud
</a>

<style>
	:root {
		--mega: #d90007;
	}

	.btn-base {
		display: inline-flex;
		align-items: center;
		gap: 0.6rem;
		padding: 0.75rem 1.5rem;
		border-radius: 0.5rem;
		font-weight: 600;
		color: #fff;
		position: relative;
		overflow: hidden;
		cursor: pointer;
		user-select: none;
		transition: transform 0.3s ease;
	}

	.btn-base:hover {
		transform: translateY(-2px) scale(1.05);
	}

	.btn-base:active {
		transform: translateY(0) scale(0.97);
	}

	.btn-base::after {
		content: '';
		position: absolute;
		inset: 0;
		background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.25), transparent);
		transform: translateX(-100%);
		transition: transform 0.7s ease;
		pointer-events: none;
	}

	.btn-base:hover::after {
		transform: translateX(100%);
	}

	.mega-btn {
		background: var(--mega);
		animation: mega-glow 2s ease-in-out infinite alternate;
	}

	@keyframes mega-glow {
		0%,
		100% {
			box-shadow:
				0 0 6px rgba(217, 0, 7, 0.6),
				0 0 12px rgba(217, 0, 7, 0.4);
		}
		50% {
			box-shadow:
				0 0 10px rgba(217, 0, 7, 0.8),
				0 0 24px rgba(217, 0, 7, 0.6);
		}
	}
</style>
