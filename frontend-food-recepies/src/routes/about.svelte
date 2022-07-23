<script>
	import { onMount } from 'svelte';

	let result = {
		title: '',
		description: '',
		data: ''
	};

	onMount(async () => {
		const res = await fetch('http://127.0.0.1:8000/food-recepies/api/v1/recipes/about');
		result = await res.json();
	});
</script>

<head>
	<title>О проекте</title>
</head>

<center>
	<h2>{result.title}</h2>
	<p>{result.description}</p>
	<p>{result.data}</p>
</center>
