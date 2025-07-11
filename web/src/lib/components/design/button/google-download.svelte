<script>
	import Google from '$lib/assets/icons/google.svelte';

	const { href } = $props();
</script>

<a {href} target="_blank" class="gdrive-btn btn-base">
	<Google width={24} height={24} />
	Google Drive
</a>

<style>
	:root {
		--gd-gold: #c9a340;
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

	.gdrive-btn {
		background: var(--gd-gold) / 80;
		animation: gd-glow 2s ease-in-out infinite alternate;
	}

	@keyframes gd-glow {
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
