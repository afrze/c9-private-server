<script lang="ts">
	import Info from '$lib/assets/icons/info.svelte';
	interface PropsType {
		title?: string;
		id?: string;
		name?: string;
		class?: string;
		type?: 'text' | 'number' | 'email';
		readonly?: boolean;
		disabled?: boolean;
		size?: string;
		maxlength?: number;
		min?: string;
		max?: string;
		pattern?: string;
		placeholder?: string;
		required?: boolean;
		autocomplete?: 'on' | 'off';
		errorMessage?: string;
		value?: string;
		onblur?: () => void;
		onchange?: () => void;
		oninput?: () => void;
		onfocus?: () => void;
	}
	const {
		value,
		title,
		type = 'text',
		readonly,
		disabled,
		max,
		min,
		maxlength,
		pattern,
		placeholder,
		autocomplete,
		required,
		errorMessage,
		onblur,
		onchange,
		oninput,
		onfocus,
		id,
		name,
		class: className
	}: PropsType = $props();

	const _class = [
		'focus:border-primary outline-none rounded-md border-1 border-gray-200 bg-white px-3 py-2 placeholder:text-gray-300 placeholder:text-sm placeholder:text',
		{ 'border-1 border-red-500 focus:border-red-500 focus:outline-red-500': errorMessage },
		className
	];
</script>

<div class="flex flex-col gap-3">
	<input
		{value}
		{title}
		{name}
		{id}
		{disabled}
		{type}
		{readonly}
		{max}
		{min}
		{maxlength}
		{pattern}
		{placeholder}
		{autocomplete}
		{required}
		{onblur}
		{onchange}
		{oninput}
		{onfocus}
		class={_class}
	/>
	{#if errorMessage}
		<span class="flex items-center gap-2 text-red-500"><Info /> {errorMessage}</span>
	{/if}
</div>
