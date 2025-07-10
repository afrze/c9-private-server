<script>
	const { onsubmit, children, class: className = '', ...props } = $props();

	const _class = [className];
</script>

<form {onsubmit} class={_class} {...props}>
	{@render children?.()}
</form>
