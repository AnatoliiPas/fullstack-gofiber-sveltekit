<script>
	export let recipes;
</script>

<head>
	<title>Рецепты</title>
</head>
<div>
	{#each recipes as item}
		<div style="margin: 4em;">
			<p>{item.name}</p>
			<a href={`/recipesv2/${item.id}`}>
				<img
					style="width: 200px;height: 200px;border-radius: 15px;"
					src={`http://127.0.0.1:8000/img/${item.image}`}
					alt={item.name}
				/>
			</a>
		</div>
	{/each}
</div>
