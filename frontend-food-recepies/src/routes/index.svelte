<script>
	import { onMount } from 'svelte';

	let result = {
		title: '',
		description: ''
	};

	onMount(async () => {
		const res = await fetch('http://127.0.0.1:8000/food-recepies/api/v1/index');
		result = await res.json();
	});
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
<center>
	<h2>{result.title}</h2>
	<p>{result.description}</p>
</center>
