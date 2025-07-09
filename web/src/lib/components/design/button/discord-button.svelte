<script>
	import Discord from '$lib/assets/icons/discord.svelte';
	const { href } = $props();
</script>

<div class="w-full">
	<a {href} target="_blank" class="discord-btn w-full">
		<Discord color="#FFF" width="32" height="32" />
		Join Discord Channel
	</a>
</div>

<style>
	:root {
		--discord: #5865f2;
	}

	.discord-btn {
		display: inline-flex;
		align-items: center;
		gap: 1rem;
		padding: 0.75rem 1.5rem;
		border-radius: 0.5rem;
		font-weight: 600;
		color: #fff;
		background: var(--discord);
		animation: discord-glow 2s ease-in-out infinite alternate;
		transition:
			transform 0.3s ease,
			background 0.3s ease;
		position: relative;
		overflow: hidden;
		cursor: pointer;
		user-select: none;
	}

	.discord-btn:hover {
		transform: translateY(-2px) scale(1.05);
	}
	.discord-btn:active {
		transform: translateY(0) scale(0.97);
	}

	.discord-btn::after {
		content: '';
		position: absolute;
		inset: 0;
		background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.25), transparent);
		transform: translateX(-100%);
		transition: transform 0.7s ease;
		pointer-events: none;
	}

	.discord-btn:hover::after {
		transform: translateX(100%);
	}

    @keyframes discord-glow {
		0%,
		100% {
			box-shadow:
				0 0 6px rgba(88, 101, 242, 0.6),
				0 0 12px rgba(88, 101, 242, 0.4);
		}
		50% {
			box-shadow:
				0 0 10px rgba(88, 101, 242, 0.8),
				0 0 24px rgba(88, 101, 242, 0.6);
		}
	}
</style>
