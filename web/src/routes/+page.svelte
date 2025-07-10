<script lang="ts">
	import Hero from '$lib/components/ui/hero/hero.svelte';
	import Login from '$lib/components/ui/login/login.svelte';
</script>

<svelte:head>
	<title>Continent of Ninth Dark</title>
	<meta name="description" content="Continent of Ninth Dark is a non pay to win Private Server" />
</svelte:head>

<section class="flex h-full w-full items-center gap-20">
	<Hero />
	<Login />
</section>
