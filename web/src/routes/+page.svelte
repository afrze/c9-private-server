<script lang="ts">
	import Discord from '$lib/assets/icons/discord.svelte';
	import Logo from '$lib/assets/images/logo.png';
	import Favicon from '$lib/assets/images/favicon.ico';
	import Button from '$lib/components/design/button/button.svelte';
</script>

<svelte:head>
	<title>Continent of Ninth Dark</title>
	<meta name="description" content="Continent of Ninth Dark is a non pay to win Private Server" />
</svelte:head>

<section class="flex h-full w-full items-center">
	<div class="w-1/2 flex flex-col items-end p-20">
		<div class="w-[320px] transition-all animate-float">
			<img src={Logo} alt="logo" class="w-full h-full" />
		</div>
		<h1 class="text-2xl font-bold pt-5 pb-10 text-gold gold-heading">
			Welcome to Continent of Ninth Dark
		</h1>
		<div class="flex flex-col gap-5">
			<div class="w-full">
				<a href="https://discord.gg/pHhHyP56M3" target="_blank" class="discord-btn w-full">
					<Discord color="#FFF" width="32" height="32" />
					Join Discord Channel
				</a>
			</div>
			<div class="w-full">
				<a
					href="https://drive.google.com/file/d/1bzkMMPHjw2hxd0RJc4eOENszfr63tdUF/view?usp=drive_link"
					target="_blank"
					class="download-btn w-full"
				>
					<div class="w-8 h-8">
						<img src={Favicon} alt="icon" />
					</div>
					Download the Game
				</a>
			</div>
			<div class="w-full">
				<Button type="primary" class="w-full p-[0.75rem]">Create an Account</Button>
			</div>
		</div>
	</div>
	<div>
		<h1>login</h1>
	</div>
</section>

<style>
	:root {
		--discord: #5865f2; /* Discord brand purple-blue */
	}

	.animate-float {
		animation: float infinite 5s ease-in-out;
	}

	.gold-heading {
		animation: glow 2s ease-in-out infinite alternate;
		transition: transform 0.3s ease;
	}
	.gold-heading:hover {
		transform: scale(1.01);
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
		position: relative; /* for sweeping light */
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

	.download-btn {
		display: inline-flex;
		align-items: center;
		gap: 0.5rem;
		padding: 0.75rem 1.5rem; /* 12 px × 24 px */
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

	/* lift on hover, press on active */
	.download-btn:hover {
		transform: translateY(-2px) scale(1.05);
	}
	.download-btn:active {
		transform: translateY(0) scale(0.97);
	}

	/* sweeping light */
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

	@keyframes float {
		0%,
		100% {
			transform: translateY(0px);
		}
		50% {
			transform: translateY(-15px);
		}
	}

	@keyframes glow {
		0%,
		100% {
			text-shadow:
				0 0 4px rgba(201, 163, 64, 0.6),
				0 0 10px rgba(201, 163, 64, 0.4);
		}
		50% {
			text-shadow:
				0 0 8px rgba(201, 163, 64, 0.85),
				0 0 20px rgba(201, 163, 64, 0.65);
		}
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
