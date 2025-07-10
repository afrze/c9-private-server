<script lang="ts">
	import Fire from '$lib/assets/images/fire.png';
</script>

<div class="glow-flame">
    <img src={Fire} alt="purple flame" class="w-[40px] h-auto" />
</div>

<style>
	.glow-flame {
		border-radius: 100%;
		display: inline-block;
		position: relative;
		animation: purple-glow 2.5s ease-in-out infinite alternate;
	}
	.glow-flame img {
		display: block;
		position: relative;
		z-index: 1;
	}
	.glow-flame::before {
		content: '';
		position: absolute;
		inset: -6%;
		border-radius: 50%;
		background: radial-gradient(
			circle,
			var(--primary) 0%,
			rgba(138, 59, 213, 0.35) 40%,
			rgba(138, 59, 213, 0) 70%
		);
		filter: blur(12px);
		z-index: 0;
	}
	@keyframes purple-glow {
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
