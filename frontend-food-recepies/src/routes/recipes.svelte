<script>
	export let recipes;
</script>

<head>
	<title>Рецепты</title>
</head>
<h1 class="text-center">Рецепты</h1>

<div style="">
	{#each recipes as item}
		<div style="margin: 4em;">
			<p>{item.title}</p>
			<a href={`/recipes/${item.id}`}>
				<img src={item.image} alt={item.title} style="border-radius: 15px;" />
			</a>
		</div>
	{/each}
</div>
