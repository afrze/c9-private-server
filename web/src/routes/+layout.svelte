<script lang="ts">
	import '../app.css';
	let { children } = $props();
</script>

<main class="h-screen">
	{@render children()}
</main>

<style>
	:root {
		--hero-image: url('/bg.png');
		--overlay: hsla(0, 0%, 0%, 0.7);
	}

	main {
		background:
			linear-gradient(var(--overlay), var(--overlay)),
			var(--hero-image) center/cover no-repeat;
		color: #f3f3f3;
		font-family: 'Poppins', sans-serif;
	}
</style>
