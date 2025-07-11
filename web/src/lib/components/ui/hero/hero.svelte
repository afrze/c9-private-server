<script lang="ts">
	import Logo from '$lib/assets/images/logo.png';
	import { Modal } from '$lib/components/design/modal';
	import Register from '../register/register.svelte';
	import { Button, DiscordButton, DownloadButton } from '$lib/components/design/button/';
	import { UserPlus } from '@lucide/svelte';
	import Download from '../download/download.svelte';

	let registerModalOpen = false;
	let downloadModalOpen = false;
</script>

{#snippet add_user()}
	<span>
		<UserPlus />
	</span>
{/snippet}

<Modal
	isOpen={registerModalOpen}
	title="Register"
	onClose={() => (registerModalOpen = false)}
	class="max-w-1/4"
>
	<Register />
</Modal>

<Modal
	isOpen={downloadModalOpen}
	title="Download the client"
	onClose={() => (downloadModalOpen = false)}
	class="max-w-1/3"
>
	<Download />
</Modal>

<article class="w-1/2 flex flex-col items-end p-20">
	<div class="w-[320px] transition-all animate-float">
		<img src={Logo} alt="logo" class="w-full h-full" />
	</div>
	<h1 class="text-2xl font-bold pt-5 pb-10 text-gold gold-heading">
		Welcome to Continent of Ninth Dark
	</h1>
	<div class="flex flex-col gap-5">
		<DiscordButton href="https://discord.gg/pHhHyP56M3" />
		<DownloadButton onclick={() => (downloadModalOpen = true)} />
		<Button
			variant="primary"
			class="w-full p-[0.75rem]"
			left={add_user}
			onclick={() => (registerModalOpen = true)}>Create an account</Button
		>
	</div>
</article>

<style>
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
</style>
