<script lang="ts">
	import Favicon from '$lib/assets/images/favicon.ico';

	const { href } = $props();
</script>

<div class="w-full">
	<a {href} target="_blank" class="download-btn w-full">
		<div class="w-8 h-8">
			<img src={Favicon} alt="icon" />
		</div>
		Download the Game
	</a>
</div>

<style>
	:root {
		--gold: #c9a24044;
	}

	.download-btn {
		display: inline-flex;
		align-items: center;
		gap: 0.5rem;
		padding: 0.75rem 1.5rem;
		border-radius: 0.5rem;
		font-weight: 600;
		color: #fff;
		background: var(--gold);
		animation: gold-glow 2s ease-in-out infinite alternate;

		transition: transform 0.3s ease;
		position: relative;
		overflow: hidden;
		cursor: pointer;
		user-select: none;
	}

	.download-btn:hover {
		transform: translateY(-2px) scale(1.05);
	}
	.download-btn:active {
		transform: translateY(0) scale(0.97);
	}

	.download-btn::after {
		content: '';
		position: absolute;
		inset: 0;
		background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.25), transparent);
		transform: translateX(-100%);
		transition: transform 0.7s ease;
		pointer-events: none;
	}
	.download-btn:hover::after {
		transform: translateX(100%);
	}

	@keyframes gold-glow {
		0%,
		100% {
			box-shadow:
				0 0 6px rgba(201, 163, 64, 0.6),
				0 0 12px rgba(201, 163, 64, 0.4);
		}
		50% {
			box-shadow:
				0 0 10px rgba(201, 163, 64, 0.8),
				0 0 24px rgba(201, 163, 64, 0.6);
		}
	}
</style>
