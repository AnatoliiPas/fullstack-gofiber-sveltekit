/** @type {import('@sveltejs/kit').Load} */
export async function GET() {
    const url = `http://127.0.0.1:8000/food-recepies/api/v1/recipes`
    const res = await fetch(url)
    const { data } = await res.json()

    if (res.ok) {
        return {
            status: 200,
            body: {
                recipes: data
            }
        }
    }

    return {
        status: 404,
        body: {
            error: "невозможно получить рецепты"
        }
    }
}