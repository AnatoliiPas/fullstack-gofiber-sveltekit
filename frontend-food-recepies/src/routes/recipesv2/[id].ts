/** @type {import('@sveltejs/kit').Load} */
export async function GET(page) {
    console.log(page.params.id)
    const url = `http://localhost:8000/food-recepies/api/v1/recipes/${page.params.id}`
    const res = await fetch(url)
    const { data } = await res.json()

    if (res.ok) {
        return {
            body: {
                recipe: data
            }
        }
    }

    return {
        body: {
            error: "невозможно получить рецепт"
        }
    }
}